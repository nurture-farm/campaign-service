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

type FindCampaignByIdInterface interface {
	OnRequest(ctx context.Context, request *fs.FindCampaignByIdRequest) *fs.FindCampaignByIdResponse
	OnData(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse
	OnError(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse, err error) *fs.FindCampaignByIdResponse
	OnResponse(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse
}

type GenericFindCampaignByIdExecutor struct {
	FindCampaignByIdInterface FindCampaignByIdInterface
}

type FindCampaignByIdController struct {
}

var FindCampaignByIdExecutor *GenericFindCampaignByIdExecutor

func (ge *GenericFindCampaignByIdExecutor) OnRequest(ctx context.Context, request *fs.FindCampaignByIdRequest) *fs.FindCampaignByIdResponse {

	logger.Info("FindCampaignByIdExecutor OnRequest hook started", zap.Any("request", request))

	if request.Id == 0 {
		logger.Error("FindCampaignByIdExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.FindCampaignByIdResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("FindCampaignByIdExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericFindCampaignByIdExecutor) OnResponse(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
	return ge.FindCampaignByIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindCampaignByIdExecutor) OnData(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
	return ge.FindCampaignByIdInterface.OnData(ctx, request, response)
}

func (ge *GenericFindCampaignByIdExecutor) OnError(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse, err error) *fs.FindCampaignByIdResponse {
	return ge.FindCampaignByIdInterface.OnError(ctx, request, response, err)
}

func (rc *FindCampaignByIdController) OnRequest(ctx context.Context, request *fs.FindCampaignByIdRequest) *fs.FindCampaignByIdResponse {
	return nil
}

func (rc *FindCampaignByIdController) OnResponse(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
	return nil
}

func (rc *FindCampaignByIdController) OnData(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
	return nil
}

func (rc *FindCampaignByIdController) OnError(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse, err error) *fs.FindCampaignByIdResponse {
	return nil
}
