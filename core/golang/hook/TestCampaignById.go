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
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TestCampaignByIdInterface interface {
	OnRequest(ctx context.Context, request *fs.TestCampaignByIdRequest) *fs.TestCampaignByIdResponse
	OnError(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse, err error) *fs.TestCampaignByIdResponse
	OnResponse(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse) *fs.TestCampaignByIdResponse
}

type GenericTestCampaignByIdExecutor struct {
	TestCampaignByIdInterface TestCampaignByIdInterface
}

type TestCampaignByIdController struct {
}

var TestCampaignByIdExecutor *GenericTestCampaignByIdExecutor

const (
	CONST_NO_CAMPAIGN_ID     = 0
	CONST_EMPTY_ATHENA_QUERY = ""
)

func testCampaignByIdInternalError() *fs.TestCampaignByIdResponse {
	return &fs.TestCampaignByIdResponse{
		Status: &common.RequestStatusResult{
			Status:        common.RequestStatus_INTERNAL_ERROR,
			ErrorMessages: []string{"Error in ExecuteFindCampaignById"},
		},
	}
}

func testCampaignByIdBadInputError() *fs.TestCampaignByIdResponse {
	return &fs.TestCampaignByIdResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_BAD_INPUT,
		},
	}
}

func testCampaignByIdSuccess() *fs.TestCampaignByIdResponse {
	return &fs.TestCampaignByIdResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},
	}
}

func (ge *GenericTestCampaignByIdExecutor) OnRequest(ctx context.Context, request *fs.TestCampaignByIdRequest) *fs.TestCampaignByIdResponse {
	logger.Info("TestCampaignByIdExecutor OnRequest hook started", zap.Any("request", request))
	if request.CampaignId == CONST_NO_CAMPAIGN_ID || request.AthenaQuery == CONST_EMPTY_ATHENA_QUERY {
		logger.Error("TestCampaignByIdExecutor OnRequest hook, Invalid TestCampaignById Request", zap.Any("request", request))
		return testCampaignByIdBadInputError()
	}

	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, &fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	})
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindCampaignById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return testCampaignByIdInternalError()
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		logger.Error("CampaignExecutor OnRequest hook, Error, could not get test campaign from DB based on campaignId", zap.Any("campaignId", request.CampaignId))
		return testCampaignByIdInternalError()
	}

	findCampaignTemplateById, err := executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, &fs.FindCampaignTemplateByIdRequest{
		CampaignId: request.CampaignId,
	})
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindCampaignTemplateById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return testCampaignByIdInternalError()
	}
	if findCampaignTemplateById == nil || findCampaignTemplateById.Records == nil {
		logger.Error("CampaignExecutor OnRequest hook, Error could not get CampaignTemplates from DB based on campaignId", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return testCampaignByIdInternalError()
	}

	events, err := createCommunicationEvent(ctx, request, findCampaignByIdResponse, findCampaignTemplateById)
	if err != nil {
		return testCampaignByIdInternalError()
	}

	logger.Info("TestCampaignByIdExecutor OnRequest hook", zap.Any("Total number of communication events", len(events)))
	err = sendCommunicationEvent(ctx, "TestCampaignByIdExecutor OnRequest hook", events, -1)
	if err != nil {
		logger.Info("TestCampaignByIdExecutor OnRequest hook, Error in sending communication event ", zap.Error(err),
			zap.Any("request", request))
		return testCampaignByIdInternalError()
	}

	logger.Info("TestCampaignByIdExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return testCampaignByIdSuccess()
}

func createCommunicationEvent(ctx context.Context, request *fs.TestCampaignByIdRequest, findCampaignByIdResponse *fs.FindCampaignByIdResponse, findCampaignTemplateById *fs.FindCampaignTemplateByIdResponse) ([]*ce.CommunicationEvent, error) {

	communicationChannel := common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
	campaignQueryType := common.CampaignQueryType(common.CampaignQueryType_value[findCampaignByIdResponse.Records.Type])
	campaignQuery := request.AthenaQuery
	campaignContentMetaData := mappers.MapContentMetaData(findCampaignByIdResponse.Records.Attributes)
	channelAttributes := mappers.MapChannelAttributes(findCampaignByIdResponse.Records.Attributes)
	campaignMedia := mappers.MapMedia(findCampaignByIdResponse.Records.Attributes)
	var targetUsers []*common.ActorID // todo -- currently keeping this as empty (No communication will be sent in case of USER_LIST type)
	var events []*ce.CommunicationEvent

	actorDetailsData, err := getActorDetailsData(ctx, "TestCampaignByIdExecutor OnRequest hook", campaignQueryType, communicationChannel, campaignQuery, targetUsers, request.CampaignId)
	if err != nil {
		logger.Error("TestCampaignByIdExecutor OnRequest hook, Error in getting actor details", zap.Any("request", request))
		err = fmt.Errorf("INTERNAL_ERROR")
		return events, err
	}

	actorDetailMap, actorByPlaceholderMap := getActorDetailsMap(actorDetailsData)
	actorDetails, actorIDDetails := getActorDetails(actorDetailMap, communicationChannel, -1)
	contentMetaData, imageMap := getContentMetaDataAndImageMap(campaignContentMetaData)
	events = []*ce.CommunicationEvent{}

	parentReferenceId := GetWorkflowId(-1)
	startIndex := 0
	for _, requestTemplate := range findCampaignTemplateById.Records {
		userList := (len(actorDetails) * cast.ToInt(requestTemplate.DistributionPercent)) / 100
		for i := startIndex; i < userList+startIndex; i++ {
			if i >= len(actorDetails) {
				break
			}
			trackingData := trackingData{
				CampaignName: "TEST_EXISTING_CAMPAIGN_BY_ID",
			}
			trackingDataBytes, err := json.Marshal(trackingData)
			if err != nil {
				logger.Error("TestCampaignByIdExecutor OnRequest hook, createCommunicationEvent Error in marshalling tracking_data", zap.Any("error", err),
					zap.Any("campaignId", request.CampaignId), zap.Any("tracking_data", trackingData))
				err = fmt.Errorf("TRACKING_DATA_MARSHAL_ERROR")
				return events, err
			}
			trackingDataPlaceHolder := &ce.Placeholder{Key: "tracking_data", Value: string(trackingDataBytes)}
			contentMetaDataWithTrackingData := contentMetaData
			contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, trackingDataPlaceHolder)
			image := GetImageMetaData(actorDetails[i].LanguageCode, imageMap)
			if image != "" {
				contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, &ce.Placeholder{Key: CONST_IMAGE, Value: image})
			}
			event := &ce.CommunicationEvent{
				TemplateName:         requestTemplate.TemplateName,
				Expiry:               timestamppb.New(timestamppb.Now().AsTime().Add(CONST_MESSAGE_EXPIRY_DURATION)), //check once
				Channel:              []common.CommunicationChannel{communicationChannel},
				ContentMetadata:      contentMetaDataWithTrackingData,
				ReceiverActorDetails: actorDetails[i],
				ReceiverActor:        actorIDDetails[i],
				ChannelAttributes:    channelAttributes,
				Media:                campaignMedia,
				CampaignName:         "TEST_EXISTING_CAMPAIGN_BY_ID",
				ParentReferenceId:    parentReferenceId,
				Placeholder:          actorByPlaceholderMap[actorIDDetails[i].ActorId],
			}
			events = append(events, event)
		}
		startIndex += userList
	}
	return events, nil
}

func (ge *GenericTestCampaignByIdExecutor) OnResponse(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse) *fs.TestCampaignByIdResponse {
	return ge.TestCampaignByIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericTestCampaignByIdExecutor) OnError(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse, err error) *fs.TestCampaignByIdResponse {
	return ge.TestCampaignByIdInterface.OnError(ctx, request, response, err)
}

func (rc *TestCampaignByIdController) OnRequest(ctx context.Context, request *fs.TestCampaignByIdRequest) *fs.TestCampaignByIdResponse {
	return nil
}

func (rc *TestCampaignByIdController) OnResponse(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse) *fs.TestCampaignByIdResponse {
	return nil
}

func (rc *TestCampaignByIdController) OnError(ctx context.Context, request *fs.TestCampaignByIdRequest, response *fs.TestCampaignByIdResponse, err error) *fs.TestCampaignByIdResponse {
	return nil
}
