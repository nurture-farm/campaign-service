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
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"context"
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	"go.uber.org/zap"
)

type AddCampaignTemplateInterface interface {
	OnRequest(ctx context.Context, request *fs.AddCampaignTemplateRequest) *fs.AddCampaignTemplateResponse
	OnError(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse, err error) *fs.AddCampaignTemplateResponse
	OnResponse(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse) *fs.AddCampaignTemplateResponse
}

type AddCampaignTemplateBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) *fs.BulkAddCampaignTemplateResponse
	OnError(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse, err error) *fs.BulkAddCampaignTemplateResponse
	OnResponse(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse) *fs.BulkAddCampaignTemplateResponse
}

type GenericAddCampaignTemplateExecutor struct {
	AddCampaignTemplateInterface AddCampaignTemplateInterface
}

type GenericAddCampaignTemplateExecutorBulk struct {
	AddCampaignTemplateBulkInterface AddCampaignTemplateBulkInterface
}

type AddCampaignTemplateController struct {
}

type BulkAddCampaignTemplateController struct {
}

var AddCampaignTemplateExecutor *GenericAddCampaignTemplateExecutor
var BulkAddCampaignTemplateExecutor *GenericAddCampaignTemplateExecutorBulk

func (ge *GenericAddCampaignTemplateExecutor) OnRequest(ctx context.Context, request *fs.AddCampaignTemplateRequest) *fs.AddCampaignTemplateResponse {

	logger.Info("AddCampaignTemplateExecutor OnRequest hook started", zap.Any("request", request))

	if request.CampaignId == 0 || request.CampaignName == "" || request.TemplateName == "" || request.DistributionPercent == 0 {
		logger.Error("AddCampaignTemplateExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.AddCampaignTemplateResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("AddCampaignTemplateExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericAddCampaignTemplateExecutor) OnResponse(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse) *fs.AddCampaignTemplateResponse {
	return ge.AddCampaignTemplateInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddCampaignTemplateExecutor) OnError(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse, err error) *fs.AddCampaignTemplateResponse {
	return ge.AddCampaignTemplateInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddCampaignTemplateExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) *fs.BulkAddCampaignTemplateResponse {
	return ge.AddCampaignTemplateBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddCampaignTemplateExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse) *fs.BulkAddCampaignTemplateResponse {
	return ge.AddCampaignTemplateBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddCampaignTemplateExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse, err error) *fs.BulkAddCampaignTemplateResponse {
	return ge.AddCampaignTemplateBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddCampaignTemplateController) OnRequest(ctx context.Context, request *fs.AddCampaignTemplateRequest) *fs.AddCampaignTemplateResponse {
	return nil
}

func (rc *AddCampaignTemplateController) OnResponse(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse) *fs.AddCampaignTemplateResponse {
	return nil
}

func (rc *AddCampaignTemplateController) OnError(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse, err error) *fs.AddCampaignTemplateResponse {
	return nil
}

func (rc *BulkAddCampaignTemplateController) OnRequest(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) *fs.BulkAddCampaignTemplateResponse {
	return nil
}

func (rc *BulkAddCampaignTemplateController) OnResponse(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse) *fs.BulkAddCampaignTemplateResponse {
	return nil
}

func (rc *BulkAddCampaignTemplateController) OnError(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse, err error) *fs.BulkAddCampaignTemplateResponse {
	return nil
}
