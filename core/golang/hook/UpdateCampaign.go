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
	query "code.nurture.farm/platform/CampaignService/core/golang/database"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/executor"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/metrics"
	"context"
	"fmt"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

type UpdateCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.UpdateCampaignRequest) *fs.UpdateCampaignResponse
	OnError(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse, err error) *fs.UpdateCampaignResponse
	OnResponse(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse) *fs.UpdateCampaignResponse
}

type UpdateCampaignBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkUpdateCampaignRequest) *fs.BulkUpdateCampaignResponse
	OnError(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse, err error) *fs.BulkUpdateCampaignResponse
	OnResponse(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse) *fs.BulkUpdateCampaignResponse
}

type GenericUpdateCampaignExecutor struct {
	UpdateCampaignInterface UpdateCampaignInterface
}

type GenericUpdateCampaignExecutorBulk struct {
	UpdateCampaignBulkInterface UpdateCampaignBulkInterface
}

type UpdateCampaignController struct {
}

type BulkUpdateCampaignController struct {
}

var UpdateCampaignExecutor *GenericUpdateCampaignExecutor
var BulkUpdateCampaignExecutor *GenericUpdateCampaignExecutorBulk

func (ge *GenericUpdateCampaignExecutor) OnRequest(ctx context.Context, request *fs.UpdateCampaignRequest) (*fs.UpdateCampaignResponse, *fs.AddNewCampaignRequest) {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.UpdateCampaign_Metrics, "UpdateCampaignExecutor", &err, ctx)

	logger.Info("UpdateCampaignExecutor OnRequest hook started", zap.Any("request", request))
	if request.Id == 0 || request.UpdatedByActor == nil || request.UpdatedByActor.ActorId == 0 ||
		request.UpdatedByActor.ActorType == common.ActorType_NO_ACTOR || (request.AddCampaignRequest.Name == "" && request.AddCampaignRequest.CronExpression == "" &&
		request.AddCampaignRequest.Status == common.CampaignStatus_NO_CAMPAGIN_STATUS && request.AddCampaignRequest.Query == "" && request.AddCampaignRequest.Occurrences == 0) {
		logger.Error("UpdateCampaignExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		err = fmt.Errorf("INVALID_REQUEST")
		return &fs.UpdateCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}, nil
	}

	findCampaignByIdRequest := &fs.FindCampaignByIdRequest{
		Id: request.Id,
	}

	var response fs.UpdateCampaignResponse
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	response.Count = 1

	dbResponse, txErr := updateCampaignTx(ctx, request, findCampaignByIdRequest)

	if txErr != nil {
		response.Status = &common.RequestStatusResult{
			Status:        common.RequestStatus_INTERNAL_ERROR,
			ErrorCode:     common.ErrorCode_DATABASE_ERROR,
			ErrorMessages: []string{"Error in database transaction"},
		}
		response.Count = 0
		response.RecordId = ""
		err = fmt.Errorf("TRANSACTION_ERROR")
		return &response, nil
	}

	findCampaignByIdResponse := dbResponse.(*fs.FindCampaignByIdResponse)
	response.RecordId = cast.ToString(findCampaignByIdResponse.Records.Id)
	findCampaignTemplateRequest := &fs.FindCampaignTemplateByIdRequest{
		CampaignId: findCampaignByIdResponse.Records.Id,
	}

	findCampaignTemplateByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, findCampaignTemplateRequest)
	if err != nil {
		logger.Error("UpdateCampaignExecutor OnRequest hook, Error could not execute FindCampaignTemplateById request", zap.Error(err), zap.Any("request", request))
		err = fmt.Errorf("FIND_TEMPLATE_BY_ID_ERROR")
		response.Status = &common.RequestStatusResult{
			Status:        common.RequestStatus_INTERNAL_ERROR,
			ErrorCode:     common.ErrorCode_DATABASE_ERROR,
			ErrorMessages: []string{"Error in getting template by id"},
		}
		return &response, nil
	}

	var addCampaignTemplateRequests []*fs.AddCampaignTemplateRequest
	if len(request.AddCampaignTemplateRequests) != 0 {
		addCampaignTemplateRequests = request.AddCampaignTemplateRequests
	} else {
		for _, requestTemplate := range findCampaignTemplateByIdResponse.Records {
			addCampaignTemplateRequest := &fs.AddCampaignTemplateRequest{
				CampaignId:          requestTemplate.Id,
				TemplateName:        requestTemplate.TemplateName,
				CampaignName:        requestTemplate.CampaignName,
				DistributionPercent: requestTemplate.DistributionPercent,
			}
			addCampaignTemplateRequests = append(addCampaignTemplateRequests, addCampaignTemplateRequest)
		}
	}

	findTargetUserRequest := &fs.FindTargetUserByIdRequest{
		CampaignId: findCampaignByIdResponse.Records.Id,
	}
	findTargetUserResponse, err := executor.RequestExecutor.ExecuteFindTargetUserById(ctx, findTargetUserRequest)
	if err != nil {
		//no need to add metrics error here
		logger.Error("UpdateCampaignExecutor OnRequest hook, Error could not execute FindTargetUserById request", zap.Error(err), zap.Any("request", request))
	}

	addTargetUserRequests := []*fs.AddTargetUserRequest{}
	for _, targetUser := range findTargetUserResponse.Records {
		addTargetUserRequest := &fs.AddTargetUserRequest{
			CampaignId: targetUser.CampaignId,
			User: &common.ActorID{
				ActorId:   targetUser.UserId,
				ActorType: common.ActorType(common.ActorType_value[targetUser.UserType]),
			},
		}
		addTargetUserRequests = append(addTargetUserRequests, addTargetUserRequest)
	}

	addNewCampaignRequest := &fs.AddNewCampaignRequest{
		AddCampaignRequest:          request.AddCampaignRequest,
		AddCampaignTemplateRequests: addCampaignTemplateRequests,
		AddTargetUserRequests:       addTargetUserRequests,
	}

	logger.Info("UpdateCampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return &response, addNewCampaignRequest
}

func updateCampaignTx(ctx context.Context, request *fs.UpdateCampaignRequest, findCampaignByIdRequest *fs.FindCampaignByIdRequest) (interface{}, error) {

	dbResponse, txErr := executor.Driver.TransactionRunner(ctx, "OnRequestUpdateCampaign", func(ctx context.Context, txName string, tx dialect.Tx) (res executor.TransactionResult, err error) {

		findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, findCampaignByIdRequest)
		if err != nil {
			return nil, err
		}

		if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
			return nil, err
		}

		if request.AddCampaignRequest.Name == "" {
			request.AddCampaignRequest.Name = findCampaignByIdResponse.Records.Name
		}
		if request.AddCampaignRequest.CronExpression == "" {
			request.AddCampaignRequest.CronExpression = findCampaignByIdResponse.Records.CronExpression
		}
		if request.AddCampaignRequest.Query == "" {
			request.AddCampaignRequest.Query = findCampaignByIdResponse.Records.Query
		}
		if request.AddCampaignRequest.Status == common.CampaignStatus_NO_CAMPAGIN_STATUS {
			request.AddCampaignRequest.Status = common.CampaignStatus(common.CampaignStatus_value[findCampaignByIdResponse.Records.Status])
		}
		if request.AddCampaignRequest.Occurrences == 0 {
			request.AddCampaignRequest.Occurrences = findCampaignByIdResponse.Records.Occurrences
		}
		if request.AddCampaignRequest.Description == "" {
			request.AddCampaignRequest.Description = findCampaignByIdResponse.Records.Description
		}
		if request.AddCampaignTemplateRequests == nil && request.AddCampaignRequest.InactionQuery == "" {
			request.AddCampaignRequest.InactionQuery = findCampaignByIdResponse.Records.InactionQuery
			request.AddCampaignRequest.InactionDuration = ptypes.DurationProto(time.Duration(findCampaignByIdResponse.Records.InactionDuration) * time.Second)
		}
		if request.AddCampaignRequest.Namespace == common.NameSpace_NO_NAMESPACE {
			request.AddCampaignRequest.Namespace = common.NameSpace(common.NameSpace_value[findCampaignByIdResponse.Records.Namespace])
		}
		if request.AddCampaignTemplateRequests == nil && request.AddCampaignRequest.ContentMetadata == nil {
			request.AddCampaignRequest.ContentMetadata = mappers.MapContentMetaData(findCampaignByIdResponse.Records.Attributes)
		}

		if request.AddCampaignTemplateRequests == nil && request.AddCampaignRequest.Media == nil {
			request.AddCampaignRequest.Media = mappers.MapMedia(findCampaignByIdResponse.Records.Attributes)
		}

		if request.AddCampaignRequest.CommunicationChannel == common.CommunicationChannel_NO_CHANNEL {
			request.AddCampaignRequest.CommunicationChannel = common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
		}
		if request.AddCampaignRequest.Type == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE {
			request.AddCampaignRequest.Type = common.CampaignQueryType(common.CampaignQueryType_value[findCampaignByIdResponse.Records.Type])
		}
		if request.AddCampaignRequest.CampaignScheduleType == common.CampaignScheduleType_NO_CAMPAIGN_SCHEDULE_TYPE {
			request.AddCampaignRequest.CampaignScheduleType = common.CampaignScheduleType(common.CampaignScheduleType_value[findCampaignByIdResponse.Records.ScheduleType])
		}
		if request.AddCampaignTemplateRequests == nil && request.AddCampaignRequest.ChannelAttributes == nil {
			request.AddCampaignRequest.ChannelAttributes = mappers.MapChannelAttributes(findCampaignByIdResponse.Records.Attributes)
		}
		if request.AddCampaignRequest.CreatedByActor == nil {
			request.AddCampaignRequest.CreatedByActor = &common.ActorID{
				ActorId:   findCampaignByIdResponse.Records.CreatedByActorid,
				ActorType: common.ActorType(common.ActorType_value[findCampaignByIdResponse.Records.CreatedByActortype]),
			}
		}
		if request.AddCampaignTemplateRequests != nil && len(request.AddCampaignTemplateRequests) != 0 {
			deleteCampaignTemplate(ctx, tx, request.Id)
			for _, addCampaignTemplateRequest := range request.AddCampaignTemplateRequests {
				addCampaignTemplateRequest.CampaignId = request.Id
				model := mappers.MakeAddCampaignTemplateRequestVO(addCampaignTemplateRequest)
				args := executor.AddCampaignTemplateArgs(model)
				var rows sql.Result
				query := query.QUERY_AddCampaignTemplate
				err := tx.Exec(ctx, query, args, &rows)
				if err != nil {
					return nil, err
				}
			}
		}
		_, err = executor.RequestExecutor.ExecuteUpdateCampaign(ctx, request, tx)
		if err != nil {
			return nil, err
		}
		return findCampaignByIdResponse, nil
	})
	return dbResponse, txErr
}

func deleteCampaignTemplate(ctx context.Context, tx dialect.Tx, campaignId int64) error {
	model := mappers.MakeCampaignTemplateRequestVO(campaignId)
	args := executor.DeleteCampaignTemplateArgs(model)
	var rows sql.Result
	err := tx.Exec(ctx, query.QUERY_DeleteCampaignTemplate, args, &rows)
	if err != nil {
		logger.Error("UpdateCampaign OnRequest hook, Error could not delete CampaignTemplate",
			zap.Error(err), zap.Any("campaignId", campaignId))
		return err
	}
	return nil
}

func (ge *GenericUpdateCampaignExecutor) OnResponse(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse) *fs.UpdateCampaignResponse {
	return ge.UpdateCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericUpdateCampaignExecutor) OnError(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse, err error) *fs.UpdateCampaignResponse {
	return ge.UpdateCampaignInterface.OnError(ctx, request, response, err)
}

func (ge *GenericUpdateCampaignExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkUpdateCampaignRequest) *fs.BulkUpdateCampaignResponse {
	return ge.UpdateCampaignBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericUpdateCampaignExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse) *fs.BulkUpdateCampaignResponse {
	return ge.UpdateCampaignBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericUpdateCampaignExecutorBulk) OnError(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse, err error) *fs.BulkUpdateCampaignResponse {
	return ge.UpdateCampaignBulkInterface.OnError(ctx, request, response, err)
}

func (rc *UpdateCampaignController) OnRequest(ctx context.Context, request *fs.UpdateCampaignRequest) *fs.UpdateCampaignResponse {
	return nil
}

func (rc *UpdateCampaignController) OnResponse(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse) *fs.UpdateCampaignResponse {
	return nil
}

func (rc *UpdateCampaignController) OnError(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse, err error) *fs.UpdateCampaignResponse {
	return nil
}

func (rc *BulkUpdateCampaignController) OnRequest(ctx context.Context, request *fs.BulkUpdateCampaignRequest) *fs.BulkUpdateCampaignResponse {
	return nil
}

func (rc *BulkUpdateCampaignController) OnResponse(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse) *fs.BulkUpdateCampaignResponse {
	return nil
}

func (rc *BulkUpdateCampaignController) OnError(ctx context.Context, request *fs.BulkUpdateCampaignRequest, response *fs.BulkUpdateCampaignResponse, err error) *fs.BulkUpdateCampaignResponse {
	return nil
}
