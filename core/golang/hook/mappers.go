/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package hook

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"database/sql"
	"encoding/json"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
)

type Attributes struct {
	Placeholders    []string  `json:"placeholders,omitempty"`
	ContentMetadata []KvPairs `json:"content_metadata,omitempty"`
}

type KvPairs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func makeAddUserJourneyVertexVO(vertex *fs.UserJourneyVertex, campaignId int64) *AddUserJourneyVertexVO {
	vertexVO := &AddUserJourneyVertexVO{
		CampaignId: mappers.GetNullableInt64(campaignId),
		EventType:  mappers.GetNullableString(vertex.EventType.String()),
		EventName:  mappers.GetNullableString(vertex.EventMetadata.EventName),
	}
	if vertex.EventMetadata != nil && vertex.EventMetadata.EventFilters != nil {
		vertexVO.Attributes = getJourneyVertexAttributes(vertex)
	}
	if vertex.EventType == common.UserJourneyEventType_INACTION {
		waitTime, _ := convertToWaitTime(vertex.InactionDuration)
		vertexVO.InactionDuration = mappers.GetNullableInt64(waitTime)
		vertexVO.InactionEventName = mappers.GetNullableString(vertex.InactionEventMetadata.EventName)
	}
	return vertexVO
}

func makeAddEngagementVertexVO(vertex *fs.EngagementVertex, campaignId int64) *AddEngagementVertexVO {
	return &AddEngagementVertexVO{
		CampaignId:   mappers.GetNullableInt64(campaignId),
		TemplateName: mappers.GetNullableString(vertex.TemplateName),
		AthenaQuery:  mappers.GetNullableString(vertex.AthenaQuery),
		Attributes:   getEngagementVertexAttributes(vertex),
		Channel:      mappers.GetNullableString(vertex.CommunicationChannel.String()),
	}
}

func makeUpdateCampaignVO(addCampaignReq *fs.AddCampaignRequest, cronExpression string, status string, campaignId int64,
	userJourneyMetadata string, engagementMetadata string, userMetadataList []*fs.UserMetadata) *UpdateCampaignVO {
	return &UpdateCampaignVO{
		CronExpression: mappers.GetNullableString(cronExpression),
		Status:         mappers.GetNullableString(status),
		Attributes:     mappers.GetAttributes(addCampaignReq, userJourneyMetadata, engagementMetadata, userMetadataList),
		CampaignID:     mappers.GetNullableInt64(campaignId),
	}
}

func makeAddEdgeVO(campaignId int64, vertexType string, fromVertexId int64, toVertexId int64, waitTenure *fs.WaitTime, waitDurationType string,
	msgAckStates []common.CommunicationState) *AddEdgeVO {

	states := []string{}
	for _, state := range msgAckStates {
		states = append(states, state.String())
	}
	addEdgeVO := &AddEdgeVO{
		CampaignId:   mappers.GetNullableInt64(campaignId),
		VertexType:   mappers.GetNullableString(vertexType),
		FromVertexId: mappers.GetNullableInt64(fromVertexId),
		ToVertexId:   mappers.GetNullableInt64(toVertexId),
	}
	if waitTenure != nil {
		if waitDurationType == "WAIT_FOR" {
			waitTime, _ := convertToWaitTime(waitTenure.WaitFor)
			addEdgeVO.WaitDuration = mappers.GetNullableInt64(waitTime)
		} else if waitDurationType == "WAIT_TILL" {
			waitFor, _ := convertWaitFor(waitTenure.WaitTill)
			addEdgeVO.WaitTime = mappers.GetNullableDateTime(waitFor - 19800)
		}
	}
	addEdgeVO.WaitType = mappers.GetNullableString(waitDurationType)
	if msgAckStates != nil {
		addEdgeVO.MessageDeliveryStates = mappers.GetNullableString(strings.Join(states[:], ","))
	}
	return addEdgeVO
}

func getJourneyVertexAttributes(vertex *fs.UserJourneyVertex) (res sql.NullString) {

	attributes := &Attributes{
		ContentMetadata: mapContentMetaData(vertex.EventMetadata.EventFilters),
	}
	bytes, err := json.Marshal(attributes)
	if err != nil {
		//TO-DO:add error metrics
		return
	}
	res = mappers.GetNullableString(string(bytes))
	logger.Info("Json Res for properties: " + res.String)
	return
}

//TO-DO: add prefix expression mapping

func getEngagementVertexAttributes(vertex *fs.EngagementVertex) (res sql.NullString) {

	attributes := &Attributes{
		Placeholders:    vertex.Placeholders,
		ContentMetadata: mapContentMetaData(vertex.ContentMetadata),
	}
	bytes, err := json.Marshal(attributes)
	if err != nil {
		//TO-DO:add error metrics
		return
	}
	res = mappers.GetNullableString(string(bytes))
	return
}

func MapEngagementVertexAttributes(AttributeList string) (*Attributes, error) {

	attributes := Attributes{}
	err := json.Unmarshal([]byte(AttributeList), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes, EngagementVertexAttributes", zap.Error(err), zap.Any("AttributeList", AttributeList))
		return nil, err
	}
	return &attributes, nil
}

func mapContentMetaData(contentMetaData []*common.Attribs) []KvPairs {

	kvpairs := []KvPairs{}
	for _, attrib := range contentMetaData {
		kvpairs = append(kvpairs, KvPairs{
			Key:   attrib.Key,
			Value: attrib.Value,
		})
	}
	return kvpairs
}

func makeUserJourneyTargetUsersVO(campaignId int64, engagementVertexId int64, referenceId string,
	eventReferenceId string, actorContactId string, status string) *AddUserJourneyTargetUsersVO {

	return &AddUserJourneyTargetUsersVO{
		CampaignId:         mappers.GetNullableInt64(campaignId),
		EngagementVertexId: mappers.GetNullableInt64(engagementVertexId),
		ReferenceId:        mappers.GetNullableString(referenceId),
		EventReferenceId:   mappers.GetNullableString(eventReferenceId),
		ActorContactId:     mappers.GetNullableString(actorContactId),
		Status:             mappers.GetNullableString(status),
	}
}

func MakeFindUserJourneyCampaignVO(models []*FindUserJourneyCampaigVO) []*fs.FilterUserJourneyCampaignResponseRecord {

	keys := []int{}
	records := []*fs.FilterUserJourneyCampaignResponseRecord{}
	campaignIdToRecordsMap := make(map[int64]*fs.FilterUserJourneyCampaignResponseRecord)
	for _, model := range models {
		record, ok := campaignIdToRecordsMap[model.Id.Int64]
		if !ok {
			keys = append(keys, cast.ToInt(model.Id.Int64))
			campaignIdToRecordsMap[model.Id.Int64] = &fs.FilterUserJourneyCampaignResponseRecord{
				Id:             model.Id.Int64,
				Namespace:      model.Namespace.String,
				Name:           model.Name.String,
				Status:         model.Status.String,
				StartTime:      convertIstToUtc(model.CreatedAt.Time).Format("2006-01-02 15:04:05"),
				Qualifications: model.Count.Int64,
			}
		} else {
			record.Qualifications += model.Count.Int64
			if model.TargetUsersStatus.String == "CONVERTED" {
				var conversionPercent float64
				conversionPercent = 0
				if model.Count.Int64 > 0 {
					conversionPercent = (cast.ToFloat64(model.Count.Int64) / cast.ToFloat64(record.Qualifications)) * 100
				}
				record.Conversions = conversionPercent
			}
			campaignIdToRecordsMap[model.Id.Int64] = record
		}
	}
	for _, key := range keys {
		record := campaignIdToRecordsMap[cast.ToInt64(key)]
		if record.Status == CONST_DRAFTED {
			record.TriggerAction = true
		}
		records = append(records, record)
	}
	return records
}

func MakeDeleteEngagementVerticesVO(campaignId int64) *DeleteEngagementVerticesVO {

	return &DeleteEngagementVerticesVO{
		CampaignId: mappers.GetNullableInt64(campaignId),
	}
}

func MakeDeletUserJourneyVerticesVO(campaignId int64) *DeleteUserJourneyVerticesVO {

	return &DeleteUserJourneyVerticesVO{
		CampaignId: mappers.GetNullableInt64(campaignId),
	}
}

func MakeDeleteEdgesVO(campaignId int64, vertexType string) *DeleteEdgesVO {

	return &DeleteEdgesVO{
		CampaignId: mappers.GetNullableInt64(campaignId),
		VertexType: mappers.GetNullableString(vertexType),
	}
}
