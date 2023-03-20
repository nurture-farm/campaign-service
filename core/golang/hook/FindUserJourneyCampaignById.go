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
	Common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	dbQuery "github.com/nurture-farm/campaign-service/core/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
	"context"
	"errors"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
)

type FindUserJourneyCampaignByIdInterface interface {
	OnRequest(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) *fs.FindUserJourneyCampaignByIdResponse
	OnData(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse
	OnError(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse, err error) *fs.FindUserJourneyCampaignByIdResponse
	OnResponse(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse
}

type GenericFindUserJourneyCampaignByIdExecutor struct {
	FindUserJourneyCampaignByIdInterface FindUserJourneyCampaignByIdInterface
}

type FindUserJourneyCampaignByIdController struct {
}

var FindUserJourneyCampaignByIdExecutor *GenericFindUserJourneyCampaignByIdExecutor

func (ge *GenericFindUserJourneyCampaignByIdExecutor) OnRequest(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) *fs.FindUserJourneyCampaignByIdResponse {
	return ge.FindUserJourneyCampaignByIdInterface.OnRequest(ctx, request)
}

func (ge *GenericFindUserJourneyCampaignByIdExecutor) OnResponse(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse {
	return ge.FindUserJourneyCampaignByIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindUserJourneyCampaignByIdExecutor) OnData(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse {
	return ge.FindUserJourneyCampaignByIdInterface.OnData(ctx, request, response)
}

func (ge *GenericFindUserJourneyCampaignByIdExecutor) OnError(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse, err error) *fs.FindUserJourneyCampaignByIdResponse {
	return ge.FindUserJourneyCampaignByIdInterface.OnError(ctx, request, response, err)
}

func getUserJourneyVertexMap(vertices []FindUserJourneyVertexVO) map[int64]FindUserJourneyVertexVO {

	userJouneyVertexmap := make(map[int64]FindUserJourneyVertexVO)
	for _, vertex := range vertices {
		userJouneyVertexmap[vertex.Id.Int64] = vertex
	}
	return userJouneyVertexmap
}

func getEngagementVerticesMap(vertices []FindEngagementVertexVO) map[int64]FindEngagementVertexVO {

	verticesMap := make(map[int64]FindEngagementVertexVO)
	for _, vertex := range vertices {
		verticesMap[vertex.Id.Int64] = vertex
	}
	return verticesMap
}

func (rc *FindUserJourneyCampaignByIdController) OnRequest(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) *fs.FindUserJourneyCampaignByIdResponse {

	err := validateFindUserJourneyCampaignByIdRequest(request)
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_BAD_INPUT, err, EMPTY, nil, EMPTY, EMPTY)
	}

	campaignName, userJourneyMetadata, engagementMetadata, err := getUserJourneyCampaign(ctx, request)
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	userJourneyVerticesModels, err := getUserJourneyCampaignVertices(ctx, request)
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	userJourneyVerticesMap := getUserJourneyVertexMap(userJourneyVerticesModels)
	userJourneyEdgesModels, err := getUserJourneyCampaignEdges(ctx, request, "USER_JOURNEY")
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	engagementVerticesModels, err := getEngagementVertices(ctx, request)
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	engagementVericesMap := getEngagementVerticesMap(engagementVerticesModels)
	engagementEdgesModels, err := getUserJourneyCampaignEdges(ctx, request, "ENGAGEMENT")
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	userJourneyVertices, err := getUserJourneysProto(ctx, userJourneyEdgesModels, userJourneyVerticesMap)
	if err != nil {
		return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_INTERNAL_ERROR, err, EMPTY, nil, EMPTY, EMPTY)
	}
	engagementVertex, err := getEngagementProto(ctx, engagementEdgesModels, engagementVericesMap, engagementVerticesModels)
	userJourneyCampaign := &fs.UserJourneyCampaign{
		EngagementStartVertex: engagementVertex,
		UserJourneys:          []*fs.UserJourney{},
	}
	for _, userJourney := range userJourneyVertices {
		userJourneyCampaign.UserJourneys = append(userJourneyCampaign.UserJourneys,
			&fs.UserJourney{
				UserJourneyVertex: userJourney,
			})
	}
	return mapToFindUserJourneyCampaignByIdResponse(request, Common.RequestStatus_SUCCESS, nil, campaignName, userJourneyCampaign, userJourneyMetadata, engagementMetadata)
}

func getEngagementProto(ctx context.Context, edges []FindEdgesVO, vertexMap map[int64]FindEngagementVertexVO, engagementVerticesModels []FindEngagementVertexVO) (*fs.EngagementVertex, error) {

	vertexProtoMap := map[int64]*fs.EngagementVertex{}
	vertex, err := getEngagementVertexProto(ctx, edges, vertexMap, 0, vertexProtoMap, engagementVerticesModels)
	if err != nil {
		return nil, err
	}
	return vertex, nil
}

func getEngagementVertexProto(ctx context.Context, edges []FindEdgesVO, vertexMap map[int64]FindEngagementVertexVO, edgeIndex int,
	vertexProtoMap map[int64]*fs.EngagementVertex, engagementVerticesModels []FindEngagementVertexVO) (*fs.EngagementVertex, error) {

	if edgeIndex >= len(edges) {
		var fromVertexProto *fs.EngagementVertex
		var err error
		if edgeIndex < len(engagementVerticesModels) {
			fromVertexProto, _, err = makeEngagementVertexProto(ctx, engagementVerticesModels[edgeIndex].Id.Int64, vertexMap, vertexProtoMap)
			if err != nil {
				return nil, err
			}
		}
		return fromVertexProto, nil
	}
	edge := edges[edgeIndex]

	fromVertexProto, prevVertexProto, err := makeEngagementVertexProto(ctx, edge.FromVertexId.Int64, vertexMap, vertexProtoMap)
	if err != nil {
		return nil, err
	}
	if !prevVertexProto {
		vertexProtoMap[edge.FromVertexId.Int64] = fromVertexProto
	}
	toVertexProto, prevVertexProto, err := makeEngagementVertexProto(ctx, edge.ToVertexId.Int64, vertexMap, vertexProtoMap)
	if err != nil {
		return nil, err
	}
	if !prevVertexProto {
		vertexProtoMap[edge.ToVertexId.Int64] = toVertexProto
	}
	fromVertexProto, err = makeEngagementVertex(fromVertexProto, edge, toVertexProto)
	if err != nil {
		return nil, err
	}
	vertexProtoMap[edge.FromVertexId.Int64] = fromVertexProto
	_, err = getEngagementVertexProto(ctx, edges, vertexMap, edgeIndex+1, vertexProtoMap, engagementVerticesModels)
	if err != nil {
		return nil, err
	}
	return fromVertexProto, nil
}

func makeEngagementVertexProto(ctx context.Context, vertexId int64, vertexMap map[int64]FindEngagementVertexVO,
	vertexProtoMap map[int64]*fs.EngagementVertex) (*fs.EngagementVertex, bool, error) {

	prevVertexProto := false
	var fromVertex FindEngagementVertexVO
	if _, ok := vertexMap[vertexId]; !ok {
		return nil, prevVertexProto, errors.New("MAPPING_ERROR")
	} else {
		fromVertex = vertexMap[vertexId]
	}
	var vertexProto *fs.EngagementVertex
	if protoMessage, ok := vertexProtoMap[vertexId]; ok {
		vertexProto = protoMessage
	}

	attributes, err := MapEngagementVertexAttributes(fromVertex.Attributes.String)
	if err != nil {
		return nil, prevVertexProto, err
	}
	contentMetData := []*Common.Attribs{}
	for _, attribute := range attributes.ContentMetadata {
		contentMetData = append(contentMetData, &Common.Attribs{Key: attribute.Key, Value: attribute.Value})
	}
	if vertexProto == nil {
		vertexProto = &fs.EngagementVertex{
			Id:                   cast.ToString(fromVertex.Id.Int64),
			CommunicationChannel: Common.CommunicationChannel(Common.CommunicationChannel_value[fromVertex.Channel.String]),
			TemplateName:         fromVertex.TemplateName.String,
			Placeholders:         attributes.Placeholders,
			ContentMetadata:      contentMetData,
			AthenaQuery:          fromVertex.AthenaQuery.String,
			Edges:                []*fs.EngagementEdge{},
		}
	} else {
		prevVertexProto = true
	}
	return vertexProto, prevVertexProto, nil
}

func makeEngagementVertex(vertexProto *fs.EngagementVertex, edge FindEdgesVO, toVertexProto *fs.EngagementVertex) (*fs.EngagementVertex, error) {

	var waitFor string
	var waitTill string
	if edge.WaitType.String == "WAIT_FOR" {
		waitFor = convertWaitForToString(edge.WaitDuration.Int64)
	} else if edge.WaitType.String == "WAIT_TILL" {
		waitTill = convertUtcToIst(edge.WaitTime.Time).Format(CONST_TIMESTAMP_LAYOUT)
	}
	vertexProto.Edges = append(vertexProto.Edges,
		&fs.EngagementEdge{
			Id: cast.ToString(edge.Id.Int64),
			WaitTime: &fs.WaitTime{
				WaitFor:  waitFor,
				WaitTill: waitTill,
			},
			EdgeType: Common.CampaignEdgeType_CHECKPOINT,
			States:   getMsgAckStates(edge.MessageDeliveryStates.String),
			Vertex:   toVertexProto,
		})
	return vertexProto, nil
}

func getMsgAckStates(states string) []Common.CommunicationState {

	var msgAckStates []Common.CommunicationState
	for _, state := range strings.Split(states, ",") {
		msgAckStates = append(msgAckStates, Common.CommunicationState(Common.CommunicationState_value[state]))
	}
	return msgAckStates
}

func getUserJourneysProto(ctx context.Context, edges []FindEdgesVO, vertexMap map[int64]FindUserJourneyVertexVO) ([]*fs.UserJourneyVertex, error) {

	var vertices []*fs.UserJourneyVertex
	var index int
	var vertex *fs.UserJourneyVertex
	var err error
	index = 0
	for {
		if index >= len(edges) {
			break
		}
		vertex, index, err = getUserJourneyProto(ctx, edges, index, vertexMap)
		if err != nil {
			return nil, err
		}
		vertices = append(vertices, vertex)
		index++
	}
	return vertices, nil
}

func getUserJourneyProto(ctx context.Context, edges []FindEdgesVO, edgeIndex int, vertexMap map[int64]FindUserJourneyVertexVO) (vertexProto *fs.UserJourneyVertex, index int, err error) {

	if edgeIndex >= len(edges) {
		index = edgeIndex
		var toVertex FindUserJourneyVertexVO
		edge := edges[edgeIndex-1]
		if _, ok := vertexMap[edge.ToVertexId.Int64]; !ok {
			err = errors.New("MAPPING_ERROR")
			return
		} else {
			toVertex = vertexMap[edge.ToVertexId.Int64]
		}
		var eventFilters []*Common.Attribs

		if toVertex.Attributes.Valid {
			attributes, err := MapEngagementVertexAttributes(toVertex.Attributes.String)
			if err != nil {
				return nil, index, err
			}
			for _, attribute := range attributes.ContentMetadata {
				eventFilters = append(eventFilters, &Common.Attribs{Key: attribute.Key, Value: attribute.Value})
			}
		}

		vertexProto = &fs.UserJourneyVertex{
			Id: cast.ToString(toVertex.Id.Int64),
			EventMetadata: &fs.EventMetadata{
				EventName:    toVertex.EventName.String,
				EventFilters: eventFilters,
			},
			EventType:        Common.UserJourneyEventType(Common.UserJourneyEventType_value[toVertex.EventType.String]),
			InactionDuration: convertWaitForToString(toVertex.InactionDuration.Int64),
			InactionEventMetadata: &fs.EventMetadata{
				EventName: toVertex.InactionEventName.String,
			},
		}
		return
	}
	edge := edges[edgeIndex]
	var fromVertex FindUserJourneyVertexVO
	if _, ok := vertexMap[edge.FromVertexId.Int64]; !ok {
		err = errors.New("MAPPING_ERROR")
		return
	} else {
		fromVertex = vertexMap[edge.FromVertexId.Int64]
	}
	var eventFilters []*Common.Attribs

	if fromVertex.Attributes.Valid {
		attributes, err := MapEngagementVertexAttributes(fromVertex.Attributes.String)
		if err != nil {
			return nil, index, err
		}
		for _, attribute := range attributes.ContentMetadata {
			eventFilters = append(eventFilters, &Common.Attribs{Key: attribute.Key, Value: attribute.Value})
		}
	}
	vertexProto = &fs.UserJourneyVertex{
		Id: cast.ToString(fromVertex.Id.Int64),
		EventMetadata: &fs.EventMetadata{
			EventName:    fromVertex.EventName.String,
			EventFilters: eventFilters,
		},
		EventType:        Common.UserJourneyEventType(Common.UserJourneyEventType_value[fromVertex.EventType.String]),
		InactionDuration: convertWaitForToString(fromVertex.InactionDuration.Int64),
		InactionEventMetadata: &fs.EventMetadata{
			EventName: fromVertex.InactionEventName.String,
		},
		Edge: &fs.UserJourneyEdge{
			Id: cast.ToString(edge.Id.Int64),
			WaitTime: &fs.WaitTime{
				WaitFor: convertWaitForToString(edge.WaitDuration.Int64),
			},
			EdgeType: Common.CampaignEdgeType_CHECKPOINT,
		},
	}
	if !edge.ToVertexId.Valid {
		vertexProto.Edge.EdgeType = Common.CampaignEdgeType_EXIT
		index = edgeIndex
		return
	}
	vertexProto.Edge.UserJourneyVertex, index, err = getUserJourneyProto(ctx, edges, edgeIndex+1, vertexMap)
	if err != nil {
		return
	}
	return
}

func getUserJourneyCampaign(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) (string, string, string, error) {

	findCampaignByIdRequest := &fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	}
	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, findCampaignByIdRequest)
	if err != nil {
		return EMPTY, EMPTY, EMPTY, err
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		return EMPTY, EMPTY, EMPTY, errors.New("NOT_FOUND")
	}
	metaData := findCampaignByIdResponse.Records.Attributes
	userJourneyMetadata, engagementMetadata, _ := mappers.MapMetaData(metaData)
	return findCampaignByIdResponse.Records.Name, userJourneyMetadata, engagementMetadata, nil
}

func getUserJourneyCampaignVertices(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) ([]FindUserJourneyVertexVO, error) {

	var userJouneyVerticesModels []FindUserJourneyVertexVO
	var rows = entsql.Rows{}
	args := FindUserJourneyCampaignByIdArgs(request)
	query := dbQuery.QUERY_FindUserJourneyVerticesByCampaignId

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("FindUserJourneyCampaignByIdController, Error could not FindUserJourneyCampaignById", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := FindUserJourneyVertexVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.EventType, &model.EventName, &model.InactionDuration, &model.InactionEventName, &model.BaseVO.Version, &model.Attributes, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
		if err != nil {
			logger.Error("FindUserJourneyCampaignByIdController, Error while fetching rows for FindUserJourneyCampaignById", zap.Error(err))
			return nil, err
		}
		userJouneyVerticesModels = append(userJouneyVerticesModels, model)
	}
	return userJouneyVerticesModels, nil
}

func getUserJourneyCampaignEdges(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, vertexType string) ([]FindEdgesVO, error) {

	var edgesModels []FindEdgesVO
	var rows = entsql.Rows{}
	args := FindUserJourneyCampaignEdgesByIdArgs(request, vertexType)
	query := dbQuery.QUERY_FindEdgesByCampaignId

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("FindUserJourneyCampaignByIdController, Error could not FindUserJourneyCampaignById", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := FindEdgesVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.VertexType, &model.FromVertexId, &model.ToVertexId, &model.WaitDuration, &model.WaitTime, &model.WaitType, &model.MessageDeliveryStates, &model.BaseVO.Version, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
		if err != nil {
			logger.Error("FindUserJourneyCampaignByIdController, Error while fetching rows for FindUserJourneyCampaignById", zap.Error(err))
			return nil, err
		}
		edgesModels = append(edgesModels, model)
	}
	return edgesModels, nil
}

func getEngagementVertex(ctx context.Context, engagementVertexId int64) (FindEngagementVertexVO, error) {

	var vertex FindEngagementVertexVO
	var rows = entsql.Rows{}
	args := FindEngagementVertexByIdArgs(engagementVertexId)
	query := dbQuery.QUERY_FindEngagementVerticesById

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not getEngagementVertices", zap.Error(err))
		return vertex, err
	}
	for rows.Next() {
		model := FindEngagementVertexVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.TemplateName, &model.Attributes, &model.AthenaQuery, &model.Channel, &model.BaseVO.Version, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
		if err != nil {
			logger.Error("Error could not getEngagementVertices, Error while fetching rows", zap.Error(err))
			return vertex, err
		}
		vertex = model
	}
	return vertex, nil
}

func getEngagementVertices(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) ([]FindEngagementVertexVO, error) {

	var vertices []FindEngagementVertexVO
	var rows = entsql.Rows{}
	args := FindUserJourneyCampaignByIdArgs(request)
	query := dbQuery.QUERY_FindEngagementVerticesByCampaignId

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("FindUserJourneyCampaignByIdController, Error could not FindUserJourneyCampaignById", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := FindEngagementVertexVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.TemplateName, &model.Attributes, &model.AthenaQuery, &model.Channel, &model.BaseVO.Version, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
		if err != nil {
			logger.Error("FindUserJourneyCampaignByIdController, Error while fetching rows for FindUserJourneyCampaignById", zap.Error(err))
			return nil, err
		}
		vertices = append(vertices, model)
	}
	return vertices, nil

}

func (rc *FindUserJourneyCampaignByIdController) OnResponse(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse {
	return nil
}

func (rc *FindUserJourneyCampaignByIdController) OnData(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse) *fs.FindUserJourneyCampaignByIdResponse {
	return nil
}

func (rc *FindUserJourneyCampaignByIdController) OnError(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest, response *fs.FindUserJourneyCampaignByIdResponse, err error) *fs.FindUserJourneyCampaignByIdResponse {
	return nil
}

func validateFindUserJourneyCampaignByIdRequest(request *fs.FindUserJourneyCampaignByIdRequest) (err error) {

	logger.Info("FindUserJourneyCampaignByIdController OnRequest hook started", zap.Any("request", request))
	if request.CampaignId == 0 {
		err = errors.New(INVALID_REQUEST)
	}
	return
}

func mapToFindUserJourneyCampaignByIdResponse(request *fs.FindUserJourneyCampaignByIdRequest, status Common.RequestStatus, err error, campaignName string, campaign *fs.UserJourneyCampaign,
	userJourneyMetadata string, engagementMetadata string) *fs.FindUserJourneyCampaignByIdResponse {

	if err != nil {
		switch status {
		case Common.RequestStatus_BAD_INPUT:
			logger.Error("FindUserJourneyCampaignByIdController request failed, invalid request", zap.Any("request", request), zap.Error(err))
		case Common.RequestStatus_INTERNAL_ERROR:
			logger.Error("FindUserJourneyCampaignByIdController request failed, internal error", zap.Any("request", request), zap.Error(err))
		}
		return &fs.FindUserJourneyCampaignByIdResponse{
			Status: &Common.RequestStatusResult{
				Status:        status,
				ErrorMessages: []string{err.Error()},
			},
		}
	}
	logger.Info("FindUserJourneyCampaignByIdController OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.FindUserJourneyCampaignByIdResponse{
		Status: &Common.RequestStatusResult{
			Status: status,
		},
		Name:                campaignName,
		Campaign:            campaign,
		UserJourneyMetadata: userJourneyMetadata,
		EngagementMetadata:  engagementMetadata,
	}
}
