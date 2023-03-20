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

type FindTargetUserByIdInterface interface {
	OnRequest(ctx context.Context, request *fs.FindTargetUserByIdRequest) *fs.FindTargetUserByIdResponse
	OnData(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse
	OnError(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse, err error) *fs.FindTargetUserByIdResponse
	OnResponse(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse
}

type GenericFindTargetUserByIdExecutor struct {
	FindTargetUserByIdInterface FindTargetUserByIdInterface
}

type FindTargetUserByIdController struct {
}

var FindTargetUserByIdExecutor *GenericFindTargetUserByIdExecutor

func (ge *GenericFindTargetUserByIdExecutor) OnRequest(ctx context.Context, request *fs.FindTargetUserByIdRequest) *fs.FindTargetUserByIdResponse {

	logger.Info("FindTargetUserByIdExecutor OnRequest hook started", zap.Any("request", request))

	if request.CampaignId == 0 {
		logger.Error("FindTargetUserByIdExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.FindTargetUserByIdResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("FindTargetUserByIdExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericFindTargetUserByIdExecutor) OnResponse(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse {
	return ge.FindTargetUserByIdInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFindTargetUserByIdExecutor) OnData(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse {
	return ge.FindTargetUserByIdInterface.OnData(ctx, request, response)
}

func (ge *GenericFindTargetUserByIdExecutor) OnError(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse, err error) *fs.FindTargetUserByIdResponse {
	return ge.FindTargetUserByIdInterface.OnError(ctx, request, response, err)
}

func (rc *FindTargetUserByIdController) OnRequest(ctx context.Context, request *fs.FindTargetUserByIdRequest) *fs.FindTargetUserByIdResponse {
	return nil
}

func (rc *FindTargetUserByIdController) OnResponse(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse {
	return nil
}

func (rc *FindTargetUserByIdController) OnData(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse) *fs.FindTargetUserByIdResponse {
	return nil
}

func (rc *FindTargetUserByIdController) OnError(ctx context.Context, request *fs.FindTargetUserByIdRequest, response *fs.FindTargetUserByIdResponse, err error) *fs.FindTargetUserByIdResponse {
	return nil
}
