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
	"context"
	"go.uber.org/zap"
)

type AddCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.AddCampaignRequest) *fs.AddCampaignResponse
	OnError(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse, err error) *fs.AddCampaignResponse
	OnResponse(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse) *fs.AddCampaignResponse
}

type AddCampaignBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddCampaignRequest) *fs.BulkAddCampaignResponse
	OnError(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse, err error) *fs.BulkAddCampaignResponse
	OnResponse(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse) *fs.BulkAddCampaignResponse
}

type GenericAddCampaignExecutor struct {
	AddCampaignInterface AddCampaignInterface
}

type GenericAddCampaignExecutorBulk struct {
	AddCampaignBulkInterface AddCampaignBulkInterface
}

type AddCampaignController struct {
}

type BulkAddCampaignController struct {
}

var AddCampaignExecutor *GenericAddCampaignExecutor
var BulkAddCampaignExecutor *GenericAddCampaignExecutorBulk

func (ge *GenericAddCampaignExecutor) OnRequest(ctx context.Context, request *fs.AddCampaignRequest) *fs.AddCampaignResponse {

	logger.Info("AddCampaignExecutor OnRequest hook started", zap.Any("request", request))

	if request.Namespace == common.NameSpace_NO_NAMESPACE || request.Name == "" || request.CronExpression == "" ||
		request.CommunicationChannel == common.CommunicationChannel_NO_CHANNEL || request.Type == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE ||
		request.CreatedByActor == nil || request.CreatedByActor.ActorId == 0 || request.CreatedByActor.ActorType == common.ActorType_NO_ACTOR {
		logger.Error("AddCampaignExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.AddCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	request.Status = common.CampaignStatus_RUNNING

	logger.Info("AddCampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericAddCampaignExecutor) OnResponse(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse) *fs.AddCampaignResponse {
	return ge.AddCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddCampaignExecutor) OnError(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse, err error) *fs.AddCampaignResponse {
	return ge.AddCampaignInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddCampaignExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddCampaignRequest) *fs.BulkAddCampaignResponse {
	return ge.AddCampaignBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddCampaignExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse) *fs.BulkAddCampaignResponse {
	return ge.AddCampaignBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddCampaignExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse, err error) *fs.BulkAddCampaignResponse {
	return ge.AddCampaignBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddCampaignController) OnRequest(ctx context.Context, request *fs.AddCampaignRequest) *fs.AddCampaignResponse {
	return nil
}

func (rc *AddCampaignController) OnResponse(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse) *fs.AddCampaignResponse {
	return nil
}

func (rc *AddCampaignController) OnError(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse, err error) *fs.AddCampaignResponse {
	return nil
}

func (rc *BulkAddCampaignController) OnRequest(ctx context.Context, request *fs.BulkAddCampaignRequest) *fs.BulkAddCampaignResponse {
	return nil
}

func (rc *BulkAddCampaignController) OnResponse(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse) *fs.BulkAddCampaignResponse {
	return nil
}

func (rc *BulkAddCampaignController) OnError(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse, err error) *fs.BulkAddCampaignResponse {
	return nil
}
