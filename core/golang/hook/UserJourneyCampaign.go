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
	ce "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	query "github.com/nurture-farm/campaign-service/core/golang/database"
	"github.com/nurture-farm/campaign-service/core/golang/grpc"
	"github.com/nurture-farm/campaign-service/core/golang/hook/aws"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/metrics"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type UserJourneyCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.UserJourneyCampaignRequest) *fs.UserJourneyCampaignResponse
	OnData(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse
	OnError(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse, err error) *fs.UserJourneyCampaignResponse
	OnResponse(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse
}

type GenericUserJourneyCampaignExecutor struct {
	UserJourneyCampaignInterface UserJourneyCampaignInterface
}

type UserJourneyCampaignController struct {
}

var UserJourneyCampaignExecutor *GenericUserJourneyCampaignExecutor

var languageToLanguageCodeMap = map[common.Language]common.LanguageCode{
	common.Language_NO_LANGUAGE: common.LanguageCode_HI_IN, //default Hindi
	common.Language_ENGLISH:     common.LanguageCode_EN_US,
	common.Language_HINDI:       common.LanguageCode_HI_IN,
	common.Language_GUJARATI:    common.LanguageCode_GU,
	common.Language_PUNJABI:     common.LanguageCode_PA,
	common.Language_KANNADA:     common.LanguageCode_KN,
	common.Language_BENGALI:     common.LanguageCode_BN,
	common.Language_MARATHI:     common.LanguageCode_MR,
	common.Language_MALAYALAM:   common.LanguageCode_ML,
	common.Language_TAMIL:       common.LanguageCode_ML,
	common.Language_TELUGU:      common.LanguageCode_TE,
}

func (ge *GenericUserJourneyCampaignExecutor) OnRequest(ctx context.Context, request *fs.UserJourneyCampaignRequest) *fs.UserJourneyCampaignResponse {
	return ge.UserJourneyCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericUserJourneyCampaignExecutor) OnResponse(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse {
	return ge.UserJourneyCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericUserJourneyCampaignExecutor) OnData(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse {
	return ge.UserJourneyCampaignInterface.OnData(ctx, request, response)
}

func (ge *GenericUserJourneyCampaignExecutor) OnError(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse, err error) *fs.UserJourneyCampaignResponse {
	return ge.UserJourneyCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *UserJourneyCampaignController) OnRequest(ctx context.Context, request *fs.UserJourneyCampaignRequest) *fs.UserJourneyCampaignResponse {

	var err error
	var engagementVertexID int64
	referenceId := request.ReferenceId
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.UserJourneyCampaign_Metrics, "UserJourneyCampaign", &err, ctx)
	logger.Info("UserJourneyCampaignExecutor OnRequest hook started", zap.Any("request", request))

	err = validaterUserJourneyCampaignRequest(request)
	if err != nil {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_BAD_INPUT, err, engagementVertexID, referenceId)
	}

	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, &fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	})
	if err != nil {
		logger.Error("UserJourneyCampaignExecutor OnRequest hook, Error in ExecuteFindCampaignById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexID, referenceId)
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		logger.Error("UserJourneyCampaignExecutor OnRequest hook, Error, could not get Campaign from DB based on campaignId", zap.Any("campaignId", request.CampaignId))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexID, referenceId)
	}

	campaignQueryType := common.CampaignQueryType(common.CampaignQueryType_value[findCampaignByIdResponse.Records.Type])
	if campaignQueryType != common.CampaignQueryType_USER_JOURNEY {
		err = errors.New("INVALID_CAMPAIGN_QUERY_TYPE")
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexID, referenceId)
	}
	return handleUserJourneyCampaign(ctx, request, findCampaignByIdResponse)
}

func GetCampaignQueryType(ctx context.Context, campaignId int64) (common.CampaignQueryType, error) {

	var campaignQueryType common.CampaignQueryType
	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, &fs.FindCampaignByIdRequest{
		Id: campaignId,
	})
	if err != nil {
		logger.Error("Error in FindCampaignById", zap.Any("error", err), zap.Any("campaignId", campaignId))
		return campaignQueryType, err
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		logger.Error("Error, invalid response for FindCampaignByID", zap.Any("campaignId", campaignId))
		return campaignQueryType, err
	}
	campaignQueryType = common.CampaignQueryType(common.CampaignQueryType_value[findCampaignByIdResponse.Records.Type])
	return campaignQueryType, nil
}

func validaterUserJourneyCampaignRequest(request *fs.UserJourneyCampaignRequest) (err error) {

	logger.Info("UserJourneyCampaignController OnRequest hook started", zap.Any("request", request))
	if request.CampaignId == 0 || request.EngagementVertexId == 0 || request.ReferenceId == EMPTY {
		err = errors.New(INVALID_REQUEST)
	}
	return
}

func mapToUserJourneyCampaignResponse(request *fs.UserJourneyCampaignRequest, status common.RequestStatus, err error, nextEngagementVertexId int64,
	referenceId string) *fs.UserJourneyCampaignResponse {

	if err != nil {
		switch status {
		case common.RequestStatus_BAD_INPUT:
			logger.Error("UserJourneyCampaignController request failed, invalid request", zap.Any("request", request), zap.Error(err))
		case common.RequestStatus_INTERNAL_ERROR:
			logger.Error("UserJourneyCampaignController request failed, internal error", zap.Any("request", request), zap.Error(err))
		}
		return &fs.UserJourneyCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        status,
				ErrorMessages: []string{err.Error()},
			},
		}
	}
	logger.Info("UserJourneyCampaignController OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.UserJourneyCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: status,
		},
		EngagementVertexId: nextEngagementVertexId,
		ReferenceId:        referenceId,
	}
}

func handleUserJourneyCampaign(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.FindCampaignByIdResponse) *fs.UserJourneyCampaignResponse {

	var engagementVertexId int64
	referenceId := request.ReferenceId
	ujcRequest := &fs.FindUserJourneyCampaignByIdRequest{
		CampaignId: request.CampaignId,
	}
	onRequestResponse := FindUserJourneyCampaignByIdExecutor.OnRequest(ctx, ujcRequest)
	if onRequestResponse == nil || onRequestResponse.Status.Status != common.RequestStatus_SUCCESS {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, errors.New("CAMPAIGN_NOT_FOUND"), engagementVertexId, referenceId)
	}
	logger.Info("handleUserJourneyCampaign, onRequest response", zap.Any("request", request), zap.Any("onRequestResponse", onRequestResponse))
	namespace := common.NameSpace(common.NameSpace_value[response.Records.Namespace])
	userJourneyCmpaign := onRequestResponse.Campaign
	engagementVertexId = request.EngagementVertexId
	engagementStartVertexId, err := GetStartEngagementVertex(ctx, request.CampaignId)
	if err != nil {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	if engagementStartVertexId == engagementVertexId {
		return handleEngagementStartVertex(ctx, request, namespace, userJourneyCmpaign, response)
	} else {
		return handleNonStartUserJourney(ctx, request)
	}
	return mapToUserJourneyCampaignResponse(request, common.RequestStatus_SUCCESS, nil, engagementVertexId, referenceId)
}

func handleEngagementStartVertex(ctx context.Context, request *fs.UserJourneyCampaignRequest, namespace common.NameSpace,
	userJourneyCampaign *fs.UserJourneyCampaign, response *fs.FindCampaignByIdResponse) *fs.UserJourneyCampaignResponse {

	var engagementVertexId int64
	referenceId := request.ReferenceId
	engagementVertex, err := getEngagementVertex(ctx, request.EngagementVertexId)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error in getting engagementVertex", zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	channel := common.CommunicationChannel(common.CommunicationChannel_value[engagementVertex.Channel.String])
	queries, err := MakeUserJourneyCampaignQuery(namespace.String(), common.NameSpace_value[namespace.String()], userJourneyCampaign, response)
	if err != nil {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	logger.Info("handleEngagementStartVertex, Queries to get user segment", zap.Any("request", request),
		zap.Any("len(queries)", len(queries)), zap.Any("queries", queries))
	var actorDetails [][]string
	actorDetailsMap := make(map[string]string)
	for indexQuery, query := range queries {
		actorDetailsForQueryMap := make(map[string]string)
		actorDetailsData, err := getActorDetailsData(ctx, "CampaignExecutor OnRequest hook", common.CampaignQueryType_USER_JOURNEY, channel, query, nil, request.CampaignId)
		if err != nil {
			logger.Error("UserJourneyCampaignController OnRequest hook, Error in getting actor details", zap.Any("request", request))
			return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
		}
		for indexActorData, row := range actorDetailsData {
			if indexActorData == 0 {
				continue
			}
			if indexQuery == 0 {
				actorDetailsMap[row[0]] = row[0]
			}
			actorDetailsForQueryMap[row[0]] = row[0]
		}
		actorDetails = actorDetailsData
		break
		if indexQuery > 0 {
			for key, _ := range actorDetailsMap {
				_, ok := actorDetailsForQueryMap[key]
				if !ok {
					delete(actorDetailsMap, key)
				}
			}
		}
	}
	//actorDetails = append(actorDetails, []string{"user_id"})
	//for key, _ := range actorDetailsMap {
	//	actorDetails = append(actorDetails, []string{key})
	//}
	findCampaignTemplateByIdResponse := &fs.FindCampaignTemplateByIdResponse{
		Attribs: response.Attribs,
		Records: []*fs.FindCampaignTemplateByIdResponseRecord{
			{
				CampaignId:          request.CampaignId,
				TemplateName:        engagementVertex.TemplateName.String,
				CampaignName:        response.Records.Name,
				DistributionPercent: 100,
			},
		},
	}
	//targetUsers, err := getUserJourneyTargetUsers(ctx, request.CampaignId, request.EngagementVertexId)
	//if err != nil {
	//	return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err)
	//}
	//relevantTrargetUsers := getRelevantTargetUsers(actorDetails, targetUsers)
	//if err != nil {
	//	return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err)
	//}
	response.Records.CommunicationChannel = engagementVertex.Channel.String
	events, errResponse := getEventsToBeSent(findCampaignTemplateByIdResponse, actorDetails, response, request.CampaignId, request.EngagementVertexId, ctx)
	if errResponse != nil {
		logger.Error("UserJourneyCampaignController onRequest hook, Error getting events to be sent", zap.Error(err),
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	if len(events) == CONST_ZERO {
		logger.Info("UserJourneyCampaignController OnRequest hook, No/zero events to send to communication engine for User Journey Campaign ",
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}

	mediaUrl, fileName, mediaType := getMediaDetails(engagementVertex)
	engagementMetaData := mappers.MapContentMetaData(engagementVertex.Attributes.String)
	contentMetadata, imageMap := getContentMetadata(engagementMetaData)
	cta := getDeepLinkURL(contentMetadata)

	addPlaceholders(ctx, engagementVertex, events, request) //in case template has dynamic place holders

	trackingData := trackingData{
		CampaignName: response.Records.Name,
	}
	trackingDataBytes, err := json.Marshal(trackingData)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error in marshalling tracking_data",
			zap.Any("tracking_data", trackingData), zap.Error(err),
			zap.Any("request", request))
		err = fmt.Errorf("TRACKING_DATA_MARSHAL_ERROR")
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	trackingDataPlaceHolder := &ce.Placeholder{Key: "tracking_data", Value: string(trackingDataBytes)}

	for _, event := range events {
		if cta != nil {
			event.ContentMetadata = append(event.ContentMetadata, cta)
		}
		event.ContentMetadata = append(event.ContentMetadata, trackingDataPlaceHolder)
		if mediaType == common.MediaType_NO_MEDIA_TYPE {
			break
		}
		image := GetImageMetaData(event.ReceiverActorDetails.LanguageCode, imageMap)
		if image != "" {
			mediaUrl = image
		}
		event.Media = &ce.Media{
			MediaType:       mediaType,
			MediaAccessType: common.MediaAccessType_PUBLIC_URL,
			MediaInfo:       mediaUrl,
			DocumentName:    fileName,
		}

		if mediaType == common.MediaType_IMAGE {
			event.ContentMetadata = []*ce.Placeholder{
				{
					Key:   "image",
					Value: mediaUrl,
				},
			}
		}
	}

	logger.Info("handleEngagementStartVertex, events to be sent", zap.Any("request", request), zap.Any("events", events))
	err = sendCommunicationEvent(ctx, "CampaignExecutor OnRequest hook ", events, request.CampaignId)
	if err != nil {
		logger.Info("UserJourneyCampaignController OnRequest hook, Error in sending communication event ", zap.Error(err),
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	nextEngagementVertices, err := GetNextEngagementVertices(ctx, request.CampaignId, request.EngagementVertexId)
	if err != nil {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	isLastEngagementVertex := false
	if nextEngagementVertices == nil || len(nextEngagementVertices) == 0 {
		isLastEngagementVertex = true
	}
	err = persistRelevantUsers(ctx, request.CampaignId, engagementVertex.Id.Int64, request.ReferenceId, events, isLastEngagementVertex)
	if err != nil {
		logger.Info("CampaignExecutor OnRequest hook, Error in persisting target users", zap.Error(err),
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	return mapToUserJourneyCampaignResponse(request, common.RequestStatus_SUCCESS, nil, engagementVertexId, referenceId)
}

func getDeepLinkURL(placeholders []*ce.Placeholder) *ce.Placeholder {

	var res *ce.Placeholder
	for _, placeholder := range placeholders {
		if placeholder.Key == "deepLinkURL" {
			res = &ce.Placeholder{
				Key:   "cta",
				Value: placeholder.Value,
			}
		}
	}
	return res
}

func getMediaDetails(engagementVertex FindEngagementVertexVO) (string, string, common.MediaType) {
	mediaUrl, fileName := getMediaUrl(engagementVertex)
	if len(fileName) == 0 {
		return mediaUrl, fileName, common.MediaType_NO_MEDIA_TYPE
	}
	fileNamePart := strings.Split(fileName, ".")
	switch fileNamePart[1] {
	case "png":
		fallthrough
	case "jpeg":
		fallthrough
	case "jpg":
		return mediaUrl, fileName, common.MediaType_IMAGE
	case "pdf":
		fallthrough
	case "doc":
		fallthrough
	case "docx":
		return mediaUrl, fileName, common.MediaType_DOCUMENT
	case "mp4":
		return mediaUrl, fileName, common.MediaType_VIDEO
	}
	return mediaUrl, fileName, common.MediaType_NO_MEDIA_TYPE
}

func getMediaUrl(engagementVertex FindEngagementVertexVO) (string, string) {
	var mediaUrl, fileName string
	if engagementVertex.Attributes.Valid {
		engagementMetaData := mappers.MapContentMetaData(engagementVertex.Attributes.String)
		for _, attribute := range engagementMetaData {
			if attribute.Key == "mediaUrl" {
				mediaUrl = attribute.Value
			}
			if attribute.Key == "fileName" {
				fileName = attribute.Value
			}
		}
	}
	return mediaUrl, fileName
}

func handleNonStartUserJourney(ctx context.Context, request *fs.UserJourneyCampaignRequest) *fs.UserJourneyCampaignResponse {

	var engagementVertexId int64
	referenceId := request.ReferenceId
	previousEngagementVertexVO, err := getPreviousEngagementVertex(ctx, request.CampaignId, request.EngagementVertexId)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error in getPreviousEngagementVertex", zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	previosEngagementTargetUsers, err := getPreviousEngagementTargetUsers(ctx, request.CampaignId, previousEngagementVertexVO.Id.Int64, request.ReferenceId)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error in getPreviousEngagementTargetUsers", zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	currentEngagementVertex, err := getEngagementVertex(ctx, request.EngagementVertexId)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error in getEngagementVertex, for current Vertex", zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	ceStateMap := make(map[common.CommunicationState]string)
	for _, state := range strings.Split(previousEngagementVertexVO.MessageDeliveryStates.String, ",") {
		stateProto := common.CommunicationState(common.CommunicationState_value[state])
		ceStateMap[stateProto] = state
	}
	var events []*ce.CommunicationEvent
	for _, previosEngagementTargetUser := range previosEngagementTargetUsers {
		response, err := grpc.CommunicationEnginePlatformGrpcClient.SearchMessageAcknowledgement(ctx, previosEngagementTargetUser.EventReferenceId.String)
		if err != nil {
			logger.Error("UserJourneyCampaignController OnRequest hook, Error in SearchMessageAcknowledgement", zap.Error(err),
				zap.Any("request", request), zap.Any("previousEngagementTargetUser", previosEngagementTargetUser))
			return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
		}
		if response.MessageAcknowledgements == nil || len(response.MessageAcknowledgements) == 0 {
			logger.Error("UserJourneyCampaignController OnRequest hook, Error in SearchMessageAcknowledgement, invalid response", zap.Any("request", request),
				zap.Any("previousEngagementTargetUser", previosEngagementTargetUser))
			continue
		}
		if _, ok := ceStateMap[common.CommunicationState(common.CommunicationState_value[response.MessageAcknowledgements[0].State.String()])]; !ok {
			continue
		}
		parentReferenceId := GetUserJourneyWorkflowParentReferenceId(request.CampaignId, request.EngagementVertexId)
		messageAcknowledgement := response.MessageAcknowledgements[0]
		actorDetails := &ce.ActorDetails{
			LanguageCode: languageToLanguageCodeMap[messageAcknowledgement.Language],
		}
		if messageAcknowledgement.Channel == common.CommunicationChannel_SMS {
			actorDetails.MobileNumber = messageAcknowledgement.ActorContactId
		} else if messageAcknowledgement.Channel == common.CommunicationChannel_APP_NOTIFICATION {
			actorDetails.FcmToken = messageAcknowledgement.ActorContactId
			actorContactIdSplit := strings.Split(previosEngagementTargetUser.ActorContactId.String, UNDERSCORE)
			if len(actorContactIdSplit) > 1 {
				actorDetails.AppType = common.AppType(common.AppType_value[actorContactIdSplit[0]])
				prefix := actorContactIdSplit[1]
				if prefix == "NF" && len(actorContactIdSplit) > 2 {
					prefix += UNDERSCORE + actorContactIdSplit[2]
					if len(actorContactIdSplit) > 3 && actorContactIdSplit[3] == "IOS" {
						prefix += UNDERSCORE + UNDERSCORE + actorContactIdSplit[3]
					}
				}
				actorDetails.AppId = common.AppID(common.AppID_value[prefix])
			}
		} else if messageAcknowledgement.Channel == common.CommunicationChannel_WHATSAPP {
			actorDetails.MobileNumber = messageAcknowledgement.ActorContactId
		} else if messageAcknowledgement.Channel == common.CommunicationChannel_EMAIL {
			actorDetails.MobileNumber = messageAcknowledgement.ActorContactId
		}

		mediaUrl, fileName, mediaType := getMediaDetails(currentEngagementVertex)
		engagementMetaData := mappers.MapContentMetaData(currentEngagementVertex.Attributes.String)
		contentMetadata, imageMap := getContentMetadata(engagementMetaData)
		cta := getDeepLinkURL(contentMetadata)

		image := GetImageMetaData(actorDetails.LanguageCode, imageMap)
		if image != "" {
			mediaUrl = image
		}

		trackingData := trackingData{
			CampaignName: messageAcknowledgement.CampaignName,
		}
		trackingDataBytes, err := json.Marshal(trackingData)
		if err != nil {
			logger.Error("UserJourneyCampaignController OnRequest hook, Error in marshalling tracking_data",
				zap.Any("tracking_data", trackingData), zap.Error(err),
				zap.Any("request", request), zap.Any("previousEngagementTargetUser", previosEngagementTargetUser))
			err = fmt.Errorf("TRACKING_DATA_MARSHAL_ERROR")
			return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
		}
		trackingDataPlaceHolder := &ce.Placeholder{Key: "tracking_data", Value: string(trackingDataBytes)}

		event := &ce.CommunicationEvent{
			TemplateName:         currentEngagementVertex.TemplateName.String,
			Expiry:               timestamppb.New(timestamppb.Now().AsTime().Add(CONST_MESSAGE_EXPIRY_DURATION)), //check once
			Channel:              []common.CommunicationChannel{messageAcknowledgement.Channel},
			ReceiverActorDetails: actorDetails,
			ChannelAttributes:    mappers.MapChannelAttributes(currentEngagementVertex.Attributes.String),
			Media: &ce.Media{
				MediaType:       mediaType,
				MediaAccessType: common.MediaAccessType_PUBLIC_URL,
				MediaInfo:       mediaUrl,
				DocumentName:    fileName,
			},
			CampaignName:      messageAcknowledgement.CampaignName,
			ParentReferenceId: parentReferenceId,
			ReferenceId:       uuid.New().String(),
			//Placeholder:          actorByPlaceholderMap[actorIDDetails[i].ActorId],
		}
		if mediaType == common.MediaType_IMAGE {
			event.ContentMetadata = []*ce.Placeholder{
				{
					Key:   "image",
					Value: mediaUrl,
				},
			}
		}
		event.ContentMetadata = append(event.ContentMetadata, trackingDataPlaceHolder)
		if cta != nil {
			event.ContentMetadata = append(event.ContentMetadata, cta)
		}
		events = append(events, event)
	}
	if len(events) == CONST_ZERO {
		logger.Info("UserJourneyCampaignController OnRequest hook, No/zero events to send to communication engine for User Journey Campaign, handleNonStartUserJourney ",
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}

	addPlaceholders(ctx, currentEngagementVertex, events, request)

	err = sendCommunicationEvent(ctx, "CampaignExecutor OnRequest hook, handleNonStartUserJourney ", events, request.CampaignId)
	if err != nil {
		logger.Info("UserJourneyCampaignController OnRequest hook, Error in sending communication event, handleNonStartUserJourney ", zap.Error(err),
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	nextEngagementVertices, err := GetNextEngagementVertices(ctx, request.CampaignId, request.EngagementVertexId)
	if err != nil {
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	isLastEngagementVertex := false
	if nextEngagementVertices == nil || len(nextEngagementVertices) == 0 {
		isLastEngagementVertex = true
	}
	err = persistRelevantUsers(ctx, request.CampaignId, currentEngagementVertex.Id.Int64, request.ReferenceId, events, isLastEngagementVertex)
	if err != nil {
		logger.Info("CampaignExecutor OnRequest hook, Error in persisting target users", zap.Error(err),
			zap.Any("request", request))
		return mapToUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, engagementVertexId, referenceId)
	}
	return mapToUserJourneyCampaignResponse(request, common.RequestStatus_SUCCESS, nil, engagementVertexId, referenceId)
}

func addPlaceholders(ctx context.Context, engagementVertex FindEngagementVertexVO, events []*ce.CommunicationEvent, request *fs.UserJourneyCampaignRequest) {
	if !engagementVertex.AthenaQuery.Valid || len(engagementVertex.AthenaQuery.String) == 0 {
		return
	}

	userIdCsv := getUserIdCsv(events)
	_, rowData, err := aws.ExecuteAthenaQuery(ctx, "select * from ("+engagementVertex.AthenaQuery.String+") where id in ("+userIdCsv+")", request.CampaignId)
	if err != nil {
		logger.Error("Error in executing engagement athena query", zap.Any("request", request), zap.Any("query", engagementVertex.AthenaQuery.String))
		return
	}
	_, actorByPlaceholderMap := getActorDetailsMap(rowData)
	for _, event := range events {
		event.Placeholder = actorByPlaceholderMap[event.ReceiverActor.ActorId]
	}
}

func getUserIdCsv(events []*ce.CommunicationEvent) string {
	if events == nil || len(events) == 0 {
		return ""
	}

	var sb strings.Builder
	for index, event := range events {
		sb.WriteString(cast.ToString(event.ReceiverActor.ActorId))
		if index < len(events)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func getPreviousEngagementTargetUsers(ctx context.Context, campaignId int64, engagementVertexId int64, referenceId string) ([]FindUserJourneyTargetUsersVO, error) {

	var models []FindUserJourneyTargetUsersVO
	var rows = entsql.Rows{}
	args := GetUserJourneyTargetUsersArgs(campaignId, engagementVertexId, referenceId)
	query := query.QUERY_GetPreviousEngagementTargetUsers

	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("UserJourneyCampaignController OnRequest hook, Error could not getPreviousEngagementTargetUsers", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := FindUserJourneyTargetUsersVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.EngagementVertexId, &model.ReferenceId, &model.EventReferenceId, &model.ActorContactId, &model.Status, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
		if err != nil {
			logger.Error("UserJourneyCampaignController OnRequest hook, Error while fetching rows for getPreviousEngagementTargetUsers", zap.Error(err))
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

func GetNextEngagementVertices(ctx context.Context, campaignId int64, engagementVertexId int64) ([]FindNextEngagementVertexVO, error) {

	var vertices []FindNextEngagementVertexVO
	var rows = entsql.Rows{}
	args := FindPreviousEngagementVertexArgs(campaignId, engagementVertexId)
	query := query.QUERY_FindNextEngaagementVertexId
	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not getNextEngagementVertices", zap.Error(err))
		return vertices, err
	}
	for rows.Next() {
		model := FindNextEngagementVertexVO{}
		err := rows.Scan(&model.Id, &model.MessageDeliveryStates, &model.WaitDuration, &model.WaitTime, &model.WaitType)
		if err != nil {
			logger.Error("Error could not getNextEngagementVertices, Error while fetching rows", zap.Error(err))
			return vertices, err
		}
		vertices = append(vertices, model)
	}
	return vertices, nil
}

func GetWaitDuration(ctx context.Context, waitDuration int64, waitTime time.Time, waitType string) time.Duration {

	var res time.Duration
	if waitType == "WAIT_FOR" {
		res = time.Duration(waitDuration) * time.Second
	} else if waitType == "WAIT_TILL" {
		res = waitTime.Sub(time.Now())
	}
	return res
}

func getPreviousEngagementVertex(ctx context.Context, campaignId int64, engagementVertexId int64) (FindPreviousEngagementVertexVO, error) {

	var vertex FindPreviousEngagementVertexVO
	var rows = entsql.Rows{}
	args := FindPreviousEngagementVertexArgs(campaignId, engagementVertexId)
	query := query.QUERY_FindPreviosEngagementVertexId
	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not getPreviousEngagementVertex", zap.Error(err))
		return vertex, err
	}
	for rows.Next() {
		model := FindPreviousEngagementVertexVO{}
		err := rows.Scan(&model.Id, &model.MessageDeliveryStates)
		if err != nil {
			logger.Error("Error could not getPreviousEngagementVertex, Error while fetching rows", zap.Error(err))
			return vertex, err
		}
		vertex = model
	}
	return vertex, nil
}

func persistRelevantUsers(ctx context.Context, campaignId int64, engagementVertexId int64, referenceId string,
	communicationEvents []*ce.CommunicationEvent, isLastEngagementVertex bool) error {
	var args []interface{}
	query := query.QUERY_AddUserJourneyTargetUsers
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}
	status := "QUALIFIED"
	if isLastEngagementVertex {
		status = "CONVERTED"
	}

	for index, event := range communicationEvents {
		if index == len(communicationEvents)-1 {
			query += "(?,?,?,?,?,?)"
		} else {
			query += "(?,?,?,?,?,?),"
		}
		model := makeUserJourneyTargetUsersVO(campaignId, engagementVertexId, referenceId, event.ReferenceId, getActorContactId(ctx, event),
			status)
		args = append(args, AddUserJourneyTargetUsersArgs(model)...)
	}

	var rows sql.Result
	err := executor.Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not persistRelevantUsers", zap.Error(err))
		return err
	}
	return err
}

func getActorContactId(ctx context.Context, communicationEvent *ce.CommunicationEvent) string {

	var actorContactId string
	receiverActor := communicationEvent.ReceiverActor
	receiverActorDetails := communicationEvent.ReceiverActorDetails
	if receiverActorDetails != nil {
		if receiverActorDetails.MobileNumber != EMPTY {
			actorContactId = receiverActorDetails.MobileNumber
		} else if receiverActorDetails.FcmToken != EMPTY {
			actorContactId = receiverActorDetails.AppType.String() + UNDERSCORE + receiverActorDetails.AppId.String() + UNDERSCORE + receiverActorDetails.FcmToken
		} else if receiverActorDetails.EmailId != EMPTY {
			actorContactId = receiverActorDetails.EmailId
		}
	} else if receiverActor != nil {
		actorContactId = communicationEvent.ReceiverActor.ActorType.String() + UNDERSCORE + cast.ToString(communicationEvent.ReceiverActor.ActorId)
	}
	return actorContactId
}

func getRelevantTargetUsers(actorDetails [][]string, targetUsers []FindUserJourneyTargetUsersVO) [][]string {

	var relevantTargetUsers [][]string
	actorDetailsIdMap := make(map[string][]string)
	for i := 1; i < len(actorDetails); i++ {
		actorDetailsIdMap[actorDetails[i][0]] = actorDetails[i]
	}
	for _, targetUser := range targetUsers {
		key := cast.ToString(targetUser.Id)
		if _, ok := actorDetailsIdMap[key]; ok {
			delete(actorDetailsIdMap, key)
		}
	}
	for _, relevantUser := range actorDetailsIdMap {
		relevantTargetUsers = append(relevantTargetUsers, relevantUser)
	}
	return relevantTargetUsers
}

//func getUserJouneyTargetUsers(ctx context.Context,) ([]FindUserJourneyTargetUsersVO, error) {
//
//	var userJourneyTargetUsers []FindUserJourneyTargetUsersVO
//	var rows = entsql.Rows{}
//	args := FindUserJourneyCampaignByIdArgs(request)
//	query := dbQuery.QUERY_FindUserJourneyVerticesByCampaignId
//
//	err := executor.Driver.GetDriver().Query(ctx, query, args, &rows)
//	if err != nil {
//		logger.Error("FindUserJourneyCampaignByIdController, Error could not FindUserJourneyCampaignById", zap.Error(err))
//		return nil, err
//	}
//	for rows.Next() {
//		model := FindUserJourneyVertexVO{}
//		err := rows.Scan(&model.Id, &model.CampaignId, &model.EventType, &model.EventName, &model.InactionDuration, &model.InactionEventName, &model.BaseVO.Version, &model.BaseVO.CreatedAt, &model.BaseVO.UpdatedAt, &model.BaseVO.DeletedAt)
//		if err != nil {
//			logger.Error("FindUserJourneyCampaignByIdController, Error while fetching rows for FindUserJourneyCampaignById", zap.Error(err))
//			return nil, err
//		}
//		userJouneyVerticesModels = append(userJouneyVerticesModels, model)
//	}
//	return userJouneyVerticesModels, nil
//}

func (rc *UserJourneyCampaignController) OnResponse(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse {
	return nil
}

func (rc *UserJourneyCampaignController) OnData(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse) *fs.UserJourneyCampaignResponse {
	return nil
}

func (rc *UserJourneyCampaignController) OnError(ctx context.Context, request *fs.UserJourneyCampaignRequest, response *fs.UserJourneyCampaignResponse, err error) *fs.UserJourneyCampaignResponse {
	return nil
}
