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
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type TestNewCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.TestNewCampaignRequest) *fs.TestNewCampaignResponse
	OnError(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse, err error) *fs.TestNewCampaignResponse
	OnResponse(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse) *fs.TestNewCampaignResponse
}

type GenericTestNewCampaignExecutor struct {
	TestNewCampaignInterface TestNewCampaignInterface
}

type TestNewCampaignController struct {
}

var TestNewCampaignExecutor *GenericTestNewCampaignExecutor

func (ge *GenericTestNewCampaignExecutor) OnRequest(ctx context.Context, request *fs.TestNewCampaignRequest) *fs.TestNewCampaignResponse {
	return ge.TestNewCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericTestNewCampaignExecutor) OnResponse(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse) *fs.TestNewCampaignResponse {
	return ge.TestNewCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericTestNewCampaignExecutor) OnError(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse, err error) *fs.TestNewCampaignResponse {
	return ge.TestNewCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *TestNewCampaignController) OnRequest(ctx context.Context, request *fs.TestNewCampaignRequest) *fs.TestNewCampaignResponse {

	logger.Info("TestNewCampaignExecutor OnRequest hook started", zap.Any("request", request))

	if request.TestCampaignRequest == nil || request.TestCampaignRequest.Namespace == common.NameSpace_NO_NAMESPACE ||
		request.TestCampaignRequest.CommunicationChannel == common.CommunicationChannel_NO_CHANNEL ||
		request.TestCampaignRequest.Type == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE ||
		(request.TestCampaignRequest.CommunicationChannel == common.CommunicationChannel_APP_NOTIFICATION && request.TestCampaignRequest.ChannelAttributes == nil) {
		logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid addCampaign Request", zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	if request.TestCampaignTemplateRequests == nil {
		logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid request, TestCampaignTemplateRequests is nil", zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	var distributionPercent int32
	distributionPercent = 0
	for _, templateRequest := range request.TestCampaignTemplateRequests {
		distributionPercent += templateRequest.DistributionPercent
		if templateRequest.TemplateName == "" || templateRequest.DistributionPercent == 0 {
			logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid TestCampaignTemplateRequests Request", zap.Any("request", request))
			return &fs.TestNewCampaignResponse{
				Status: &common.RequestStatusResult{
					Status: common.RequestStatus_BAD_INPUT,
				},
			}
		}
	}
	if distributionPercent != 100 {
		logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid request, distribution percent doesn't add to 100", zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	for _, targtUserRequest := range request.TestTargetUserRequests {
		if targtUserRequest.User == nil || targtUserRequest.User.ActorId == 0 || targtUserRequest.User.ActorType == common.ActorType_NO_ACTOR {
			logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid targtUserRequest Request", zap.Any("targtUserRequest", targtUserRequest))
			return &fs.TestNewCampaignResponse{
				Status: &common.RequestStatusResult{
					Status: common.RequestStatus_BAD_INPUT,
				},
			}
		}
	}

	channel := request.TestCampaignRequest.CommunicationChannel
	campaignQueryType := request.TestCampaignRequest.Type
	campaignQuery := request.TestCampaignRequest.Query

	campaignContentMetaData := request.TestCampaignRequest.ContentMetadata
	channelAttributes := request.TestCampaignRequest.ChannelAttributes
	if (request.TestCampaignRequest.Media != nil) && (request.TestCampaignRequest.Media.MediaType == common.MediaType_DOCUMENT) {
		request.TestCampaignRequest.Media.DocumentName = fetchDocumentName(request.TestCampaignRequest.Media.GetMediaInfo())
	}

	targetUsers := []*common.ActorID{}
	targetUsersPlaceHolderMap := make(map[string][]*common.Attribs)
	for _, val := range request.TestTargetUserRequests {
		targetUsers = append(targetUsers, val.User)
		targetUsersPlaceHolderMap[cast.ToString(val.User.ActorId)] = val.Attribs
	}

	if campaignQueryType == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE ||
		(campaignQueryType == common.CampaignQueryType_ATHENA && campaignQuery == "") ||
		(campaignQueryType == common.CampaignQueryType_USER_LIST && targetUsers == nil) {
		logger.Error("TestNewCampaignExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	actorDetailsData, err := getActorDetailsData(ctx, "CampaignExecutor OnRequest hook", campaignQueryType, channel, campaignQuery, targetUsers, -1)
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in getting actor details", zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        common.RequestStatus_INTERNAL_ERROR,
				ErrorMessages: []string{"Error in getting actor details"},
			},
		}
	}
	actorDetailsData = mapUserListPlaceHolders(targetUsersPlaceHolderMap, actorDetailsData)
	actorDetailMap, actorByPlaceholderMap := getActorDetailsMap(actorDetailsData)

	contentMetaData, imageMap := getContentMetaDataAndImageMap(campaignContentMetaData)
	temp := make(map[int64]UserDetail)
	var stateToLanguageMap = getStatetoLanguageMapping()
	var stateIdToStateMap = getStateIdtoStateMapping()
	templateSelectOption := "default"
	for _, v := range contentMetaData {
		if v.Key == "templateSelectOption" {
			templateSelectOption = v.Value
		}

	}
	if templateSelectOption == "location" {
		for k, v := range actorDetailMap {
			location := v.Location
			if len(location) == 0 {
				continue
			}
			if v.ActorType == "FARMER" {
				secondaryLangKey, ok := stateToLanguageMap[location]
				if ok {
					v.SecondaryLangKey = secondaryLangKey
				}
			} else {
				state, ok := stateIdToStateMap[location]
				if ok {
					secondaryLangKey, ok := stateToLanguageMap[state]
					if ok {
						v.SecondaryLangKey = secondaryLangKey
					}
				}
			}
			temp[k] = v
		}
	}
	for k, v := range temp {
		actorDetailMap[k] = v
	}
	actorDetails, actorIDDetails := getActorDetails(actorDetailMap, channel, -1)
	events := []*ce.CommunicationEvent{}

	campaignMediaList := []*ce.Media{}
	parentReferenceId := GetWorkflowId(-1)
	startIndex := 0
	for _, requestTemplate := range request.TestCampaignTemplateRequests {
		userList := (len(actorDetails) * cast.ToInt(requestTemplate.DistributionPercent)) / 100
		for i := startIndex; i < userList+startIndex; i++ {
			if i >= len(actorDetails) {
				break
			}
			trackingData := trackingData{
				CampaignName: "TEST_CAMPAIGN",
			}
			trackingDataBytes, err := json.Marshal(trackingData)
			if err != nil {
				logger.Error("TestNewCampaignExecutor OnRequest hook, Error in marshalling tracking_data", zap.Any("error", err),
					zap.Any("campaignId", -1), zap.Any("tracking_data", trackingData))
				err = fmt.Errorf("TRACKING_DATA_MARSHAL_ERROR")
				return &fs.TestNewCampaignResponse{
					Status: &common.RequestStatusResult{
						Status:        common.RequestStatus_INTERNAL_ERROR,
						ErrorMessages: []string{"Error in marshalling tracking data"},
					},
				}
			}
			trackingDataPlaceHolder := &ce.Placeholder{Key: "tracking_data", Value: string(trackingDataBytes)}
			contentMetaDataWithTrackingData := contentMetaData
			contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, trackingDataPlaceHolder)
			image := GetImageMetaData(actorDetails[i].LanguageCode, imageMap)
			if image != "" {
				contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, &ce.Placeholder{Key: CONST_IMAGE, Value: image})
			}

			campaignMedia := &ce.Media{}
			if channel == common.CommunicationChannel_WHATSAPP {
				campaignMedia = getWhatsAppMedia(request.TestCampaignRequest.Media, image)
			}
			campaignMediaList = append(campaignMediaList, campaignMedia)
			if channel == common.CommunicationChannel_APP_NOTIFICATION {
				campaignMediaList[i] = nil
			}
			event := &ce.CommunicationEvent{
				TemplateName:         requestTemplate.TemplateName,
				Expiry:               timestamppb.New(timestamppb.Now().AsTime().Add(CONST_MESSAGE_EXPIRY_DURATION)), //check once
				Channel:              []common.CommunicationChannel{channel},
				ContentMetadata:      contentMetaDataWithTrackingData,
				ReceiverActorDetails: actorDetails[i],
				ReceiverActor:        actorIDDetails[i],
				ChannelAttributes:    channelAttributes,
				Media:                campaignMediaList[i],
				CampaignName:         "TEST_CAMPAIGN",
				ParentReferenceId:    parentReferenceId,
				Placeholder:          actorByPlaceholderMap[actorIDDetails[i].ActorId],
			}
			events = append(events, event)
		}
		startIndex += userList
	}
	logger.Info("TestNewCampaignExecutor OnRequest hook", zap.Any("Total number of communication events", len(events)))
	err = sendCommunicationEvent(ctx, "TestNewCampaignExecutor OnRequest hook", events, -1)
	if err != nil {
		logger.Info("TestNewCampaignExecutor OnRequest hook, Error in sending communication event ", zap.Error(err),
			zap.Any("request", request))
		return &fs.TestNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        common.RequestStatus_INTERNAL_ERROR,
				ErrorMessages: []string{"Error in sending communication event"},
			},
		}
	}

	logger.Info("TestNewCampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.TestNewCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},
	}
}

func fetchDocumentName(mediaUrl string) string {
	docName := strings.Split(mediaUrl, "/")
	filterDocName := strings.Split(docName[len(docName)-1], ".")
	documentName := filterDocName[0]
	return documentName
}

func (rc *TestNewCampaignController) OnResponse(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse) *fs.TestNewCampaignResponse {
	return nil
}

func (rc *TestNewCampaignController) OnError(ctx context.Context, request *fs.TestNewCampaignRequest, response *fs.TestNewCampaignResponse, err error) *fs.TestNewCampaignResponse {
	return nil
}
