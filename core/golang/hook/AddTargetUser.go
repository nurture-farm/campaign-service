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

type AddTargetUserInterface interface {
	OnRequest(ctx context.Context, request *fs.AddTargetUserRequest) *fs.AddTargetUserResponse
	OnError(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse, err error) *fs.AddTargetUserResponse
	OnResponse(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse) *fs.AddTargetUserResponse
}

type AddTargetUserBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddTargetUserRequest) *fs.BulkAddTargetUserResponse
	OnError(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse, err error) *fs.BulkAddTargetUserResponse
	OnResponse(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse) *fs.BulkAddTargetUserResponse
}

type GenericAddTargetUserExecutor struct {
	AddTargetUserInterface AddTargetUserInterface
}

type GenericAddTargetUserExecutorBulk struct {
	AddTargetUserBulkInterface AddTargetUserBulkInterface
}

type AddTargetUserController struct {
}

type BulkAddTargetUserController struct {
}

var AddTargetUserExecutor *GenericAddTargetUserExecutor
var BulkAddTargetUserExecutor *GenericAddTargetUserExecutorBulk

func (ge *GenericAddTargetUserExecutor) OnRequest(ctx context.Context, request *fs.AddTargetUserRequest) *fs.AddTargetUserResponse {

	logger.Info("AddTargetUserExecutor OnRequest hook started", zap.Any("request", request))

	if request.User == nil || request.User.ActorId == 0 || request.User.ActorType == common.ActorType_NO_ACTOR {
		logger.Error("AddTargetUserExecutor OnRequest hook, invalid request", zap.Any("request", request))
		return &fs.AddTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("AddTargetUserExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericAddTargetUserExecutor) OnResponse(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse) *fs.AddTargetUserResponse {
	return ge.AddTargetUserInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddTargetUserExecutor) OnError(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse, err error) *fs.AddTargetUserResponse {
	return ge.AddTargetUserInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddTargetUserExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddTargetUserRequest) *fs.BulkAddTargetUserResponse {
	return ge.AddTargetUserBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddTargetUserExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse) *fs.BulkAddTargetUserResponse {
	return ge.AddTargetUserBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddTargetUserExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse, err error) *fs.BulkAddTargetUserResponse {
	return ge.AddTargetUserBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddTargetUserController) OnRequest(ctx context.Context, request *fs.AddTargetUserRequest) *fs.AddTargetUserResponse {
	return nil
}

func (rc *AddTargetUserController) OnResponse(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse) *fs.AddTargetUserResponse {
	return nil
}

func (rc *AddTargetUserController) OnError(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse, err error) *fs.AddTargetUserResponse {
	return nil
}

func (rc *BulkAddTargetUserController) OnRequest(ctx context.Context, request *fs.BulkAddTargetUserRequest) *fs.BulkAddTargetUserResponse {
	return nil
}

func (rc *BulkAddTargetUserController) OnResponse(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse) *fs.BulkAddTargetUserResponse {
	return nil
}

func (rc *BulkAddTargetUserController) OnError(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse, err error) *fs.BulkAddTargetUserResponse {
	return nil
}
