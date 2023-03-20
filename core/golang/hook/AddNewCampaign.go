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
	query "github.com/nurture-farm/campaign-service/core/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/metrics"
	"context"
	"fmt"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type AddNewCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.AddNewCampaignRequest) *fs.AddNewCampaignResponse
	OnError(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse, err error) *fs.AddNewCampaignResponse
	OnResponse(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse) *fs.AddNewCampaignResponse
}

type AddNewCampaignBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddNewCampaignRequest) *fs.BulkAddNewCampaignResponse
	OnError(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse, err error) *fs.BulkAddNewCampaignResponse
	OnResponse(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse) *fs.BulkAddNewCampaignResponse
}

type GenericAddNewCampaignExecutor struct {
	AddNewCampaignInterface AddNewCampaignInterface
}

type GenericAddNewCampaignExecutorBulk struct {
	AddNewCampaignBulkInterface AddNewCampaignBulkInterface
}

type AddNewCampaignController struct {
}

type BulkAddNewCampaignController struct {
}

var AddNewCampaignExecutor *GenericAddNewCampaignExecutor
var BulkAddNewCampaignExecutor *GenericAddNewCampaignExecutorBulk

const (
	CONST_LOG_LEVEL_INFO   = "INFO"
	CONST_LOG_LEVEL_ERRROR = "ERROR"
)

func (ge *GenericAddNewCampaignExecutor) OnRequest(ctx context.Context, request *fs.AddNewCampaignRequest) *fs.AddNewCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddNewCampaign_Metrics, "AddNewCampaignExecutor", &err, ctx)
	OutputLog("AddNewCampaignExecutor OnRequest hook started", CONST_LOG_LEVEL_INFO, nil, request)

	if request.AddCampaignRequest == nil || request.AddCampaignRequest.Namespace == common.NameSpace_NO_NAMESPACE ||
		request.AddCampaignRequest.Name == "" || request.AddCampaignRequest.CronExpression == "" ||
		request.AddCampaignRequest.CommunicationChannel == common.CommunicationChannel_NO_CHANNEL ||
		request.AddCampaignRequest.Type == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE ||
		request.AddCampaignRequest.CreatedByActor == nil || request.AddCampaignRequest.CreatedByActor.ActorId == 0 ||
		request.AddCampaignRequest.CreatedByActor.ActorType == common.ActorType_NO_ACTOR ||
		(request.AddCampaignRequest.CommunicationChannel == common.CommunicationChannel_APP_NOTIFICATION && request.AddCampaignRequest.ChannelAttributes == nil) ||
		(request.AddCampaignRequest.CampaignScheduleType == common.CampaignScheduleType_INACTION_OVER_TIME && request.AddCampaignRequest.InactionQuery == "") {
		OutputLog("AddNewCampaignExecutor OnRequest hook, Invalid addCampaign Request", CONST_LOG_LEVEL_ERRROR, nil, request)
		err = fmt.Errorf("INVALID_REQUEST")
		return &fs.AddNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	if request.AddCampaignTemplateRequests == nil {
		OutputLog("AddNewCampaignExecutor OnRequest hook, Invalid request, AddCampaignTemplate is nil", CONST_LOG_LEVEL_ERRROR, err, request)
		err = fmt.Errorf("INVALID_REQUEST")
		return &fs.AddNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	var distributionPercent int32
	distributionPercent = 0
	for _, addTemplateRequest := range request.AddCampaignTemplateRequests {
		distributionPercent += addTemplateRequest.DistributionPercent
		if addTemplateRequest.CampaignName == "" || addTemplateRequest.TemplateName == "" || addTemplateRequest.DistributionPercent == 0 {
			OutputLog("AddNewCampaignExecutor OnRequest hook, Invalid addCampaingTemplate Request", CONST_LOG_LEVEL_ERRROR, err, request)
			err = fmt.Errorf("INVALID_REQUEST")
			return &fs.AddNewCampaignResponse{
				Status: &common.RequestStatusResult{
					Status: common.RequestStatus_BAD_INPUT,
				},
			}
		}
	}
	if distributionPercent != 100 {
		OutputLog("AddNewCampaignExecutor OnRequest hook, Invalid request, distribution percent doesn't add to 100", CONST_LOG_LEVEL_ERRROR, err, request)
		err = fmt.Errorf("INVALID_REQUEST")
		return &fs.AddNewCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	for _, addTargtUserRequest := range request.AddTargetUserRequests {
		if addTargtUserRequest.User == nil || addTargtUserRequest.User.ActorId == 0 || addTargtUserRequest.User.ActorType == common.ActorType_NO_ACTOR {
			logger.Error("AddNewCampaignExecutor OnRequest hook, Invalid addTargetUser Request", zap.Any("addTargtUserRequest", addTargtUserRequest))
			err = fmt.Errorf("INVALID_REQUEST")
			return &fs.AddNewCampaignResponse{
				Status: &common.RequestStatusResult{
					Status: common.RequestStatus_BAD_INPUT,
				},
			}
		}
	}

	addNewCampaignRequest := request
	var response fs.AddNewCampaignResponse
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	request.AddCampaignRequest.Status = common.CampaignStatus_RUNNING

	response.Count = 1
	_, txErr := executor.Driver.TransactionRunner(ctx, "OnRequestAddNewCampaign", func(ctx context.Context, txName string, tx dialect.Tx) (res executor.TransactionResult, err error) {

		model := mappers.MakeAddCampaignRequestVO(addNewCampaignRequest.AddCampaignRequest, "", "", nil)
		args := executor.AddCampaignArgs(model)
		var rows sql.Result
		err = tx.Exec(ctx, query.QUERY_AddCampaign, args, &rows)
		if err != nil {
			OutputLog("AddNewCampaignExecutor OnRequest hook, Error could not execute AddCampaign request", CONST_LOG_LEVEL_ERRROR, err, request)
			return nil, err
		}
		insertedId, err := rows.LastInsertId()
		if err != nil {
			OutputLog("AddNewCampaignExecutor OnRequest hook, Error could not get lastInsertedId for AddCampaign", CONST_LOG_LEVEL_ERRROR, err, request)
			return nil, err
		}
		response.RecordId = cast.ToString(insertedId)

		for _, addCampaignTemplateRequest := range addNewCampaignRequest.AddCampaignTemplateRequests {

			addCampaignTemplateRequest.CampaignId = insertedId
			model := mappers.MakeAddCampaignTemplateRequestVO(addCampaignTemplateRequest)
			args := executor.AddCampaignTemplateArgs(model)

			var rows sql.Result
			query := query.QUERY_AddCampaignTemplate

			err := tx.Exec(ctx, query, args, &rows)
			if err != nil {
				OutputLog("AddNewCampaignExecutor OnRequest hook, Error could not execute AddCampaignTempalte request", CONST_LOG_LEVEL_ERRROR, err, request)
				return nil, err
			}
		}

		for _, addTargetUserRequest := range addNewCampaignRequest.AddTargetUserRequests {

			addTargetUserRequest.CampaignId = insertedId
			model := mappers.MakeAddTargetUserRequestVO(addTargetUserRequest)
			args := executor.AddTargetUserArgs(model)

			var rows sql.Result
			query := query.QUERY_AddTargetUser

			err := tx.Exec(ctx, query, args, &rows)
			if err != nil {
				OutputLog("AddNewCampaignExecutor OnRequest hook, Error could not execute AddTargetUser request", CONST_LOG_LEVEL_ERRROR, err, request)
				return nil, err
			}
		}

		return nil, nil
	})

	if txErr != nil {
		response.Status = &common.RequestStatusResult{
			Status:        common.RequestStatus_INTERNAL_ERROR,
			ErrorCode:     common.ErrorCode_DATABASE_ERROR,
			ErrorMessages: []string{"Error in database transaction"},
		}
		response.Count = 0
		response.RecordId = ""
		err = fmt.Errorf("TRANSACTION_ERROR")
		return &response
	}

	OutputLog("AddNewCampaignExecutor OnRequest hook completed successfully", CONST_LOG_LEVEL_INFO, nil, request)
	return &response
}

func (ge *GenericAddNewCampaignExecutor) OnResponse(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse) *fs.AddNewCampaignResponse {
	return ge.AddNewCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddNewCampaignExecutor) OnError(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse, err error) *fs.AddNewCampaignResponse {
	return ge.AddNewCampaignInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddNewCampaignExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddNewCampaignRequest) *fs.BulkAddNewCampaignResponse {
	return ge.AddNewCampaignBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddNewCampaignExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse) *fs.BulkAddNewCampaignResponse {
	return ge.AddNewCampaignBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddNewCampaignExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse, err error) *fs.BulkAddNewCampaignResponse {
	return ge.AddNewCampaignBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddNewCampaignController) OnRequest(ctx context.Context, request *fs.AddNewCampaignRequest) *fs.AddNewCampaignResponse {
	return nil
}

func (rc *AddNewCampaignController) OnResponse(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse) *fs.AddNewCampaignResponse {
	return nil
}

func (rc *AddNewCampaignController) OnError(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse, err error) *fs.AddNewCampaignResponse {
	return nil
}

func (rc *BulkAddNewCampaignController) OnRequest(ctx context.Context, request *fs.BulkAddNewCampaignRequest) *fs.BulkAddNewCampaignResponse {
	return nil
}

func (rc *BulkAddNewCampaignController) OnResponse(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse) *fs.BulkAddNewCampaignResponse {
	return nil
}

func (rc *BulkAddNewCampaignController) OnError(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse, err error) *fs.BulkAddNewCampaignResponse {
	return nil
}
