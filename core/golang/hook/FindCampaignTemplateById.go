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

type FindCampaignTemplateByIdInterface interface {
	OnRequest(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) *fs.FindCampaignTemplateByIdResponse
	OnData(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse
	OnError(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse, err error) *fs.FindCampaignTemplateByIdResponse
	OnResponse(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse
}

type GenericFindCampaignTemplateByIdExecutor struct {
	FindCampaignTemplateByIdInterface FindCampaignTemplateByIdInterface
}

type FindCampaignTemplateByIdController struct {
}

var FindCampaignTemplateByIdExecutor *GenericFindCampaignTemplateByIdExecutor

func (ge *GenericFindCampaignTemplateByIdExecutor) OnRequest(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) *fs.FindCampaignTemplateByIdResponse {

	logger.Info("FindCampaignTemplateByIdExecutor OnRequest hook started", zap.Any("request", request))

	if request.CampaignId == 0 {
		logger.Error("FindCampaignTemplateByIdExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.FindCampaignTemplateByIdResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("FindCampaignTemplateByIdExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericFindCampaignTemplateByIdExecutor) OnResponse(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
	return ge.FindCampaignTemplateByIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindCampaignTemplateByIdExecutor) OnData(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
	return ge.FindCampaignTemplateByIdInterface.OnData(ctx, request, response)
}

func (ge *GenericFindCampaignTemplateByIdExecutor) OnError(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse, err error) *fs.FindCampaignTemplateByIdResponse {
	return ge.FindCampaignTemplateByIdInterface.OnError(ctx, request, response, err)
}

func (rc *FindCampaignTemplateByIdController) OnRequest(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) *fs.FindCampaignTemplateByIdResponse {
	return nil
}

func (rc *FindCampaignTemplateByIdController) OnResponse(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
	return nil
}

func (rc *FindCampaignTemplateByIdController) OnData(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
	return nil
}

func (rc *FindCampaignTemplateByIdController) OnError(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse, err error) *fs.FindCampaignTemplateByIdResponse {
	return nil
}
