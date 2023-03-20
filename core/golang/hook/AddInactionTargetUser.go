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

type AddInactionTargetUserInterface interface {
	OnRequest(ctx context.Context, request *fs.AddInactionTargetUserRequest) *fs.AddInactionTargetUserResponse
	OnError(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse, err error) *fs.AddInactionTargetUserResponse
	OnResponse(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse) *fs.AddInactionTargetUserResponse
}

type AddInactionTargetUserBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) *fs.BulkAddInactionTargetUserResponse
	OnError(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse, err error) *fs.BulkAddInactionTargetUserResponse
	OnResponse(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse) *fs.BulkAddInactionTargetUserResponse
}

type GenericAddInactionTargetUserExecutor struct {
	AddInactionTargetUserInterface AddInactionTargetUserInterface
}

type GenericAddInactionTargetUserExecutorBulk struct {
	AddInactionTargetUserBulkInterface AddInactionTargetUserBulkInterface
}

type AddInactionTargetUserController struct {
}

type BulkAddInactionTargetUserController struct {
}

var AddInactionTargetUserExecutor *GenericAddInactionTargetUserExecutor
var BulkAddInactionTargetUserExecutor *GenericAddInactionTargetUserExecutorBulk

func (ge *GenericAddInactionTargetUserExecutor) OnRequest(ctx context.Context, request *fs.AddInactionTargetUserRequest) *fs.AddInactionTargetUserResponse {
	//return ge.AddInactionTargetUserInterface.OnRequest(ctx,request)
	logger.Info("AddInactionTargetUserExecutor OnRequest hook started", zap.Any("request", request))

	if request.User == nil || request.User.ActorId == 0 || request.User.ActorType == common.ActorType_NO_ACTOR {
		logger.Error("AddInactionTargetUserExecutor OnRequest hook, invalid request", zap.Any("request", request))
		return &fs.AddInactionTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}

	logger.Info("AddInactionTargetUserExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericAddInactionTargetUserExecutor) OnResponse(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse) *fs.AddInactionTargetUserResponse {
	return ge.AddInactionTargetUserInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddInactionTargetUserExecutor) OnError(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse, err error) *fs.AddInactionTargetUserResponse {
	return ge.AddInactionTargetUserInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddInactionTargetUserExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) *fs.BulkAddInactionTargetUserResponse {
	return ge.AddInactionTargetUserBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddInactionTargetUserExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse) *fs.BulkAddInactionTargetUserResponse {
	return ge.AddInactionTargetUserBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddInactionTargetUserExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse, err error) *fs.BulkAddInactionTargetUserResponse {
	return ge.AddInactionTargetUserBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddInactionTargetUserController) OnRequest(ctx context.Context, request *fs.AddInactionTargetUserRequest) *fs.AddInactionTargetUserResponse {
	return nil
}

func (rc *AddInactionTargetUserController) OnResponse(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse) *fs.AddInactionTargetUserResponse {
	return nil
}

func (rc *AddInactionTargetUserController) OnError(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse, err error) *fs.AddInactionTargetUserResponse {
	return nil
}

func (rc *BulkAddInactionTargetUserController) OnRequest(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) *fs.BulkAddInactionTargetUserResponse {
	return nil
}

func (rc *BulkAddInactionTargetUserController) OnResponse(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse) *fs.BulkAddInactionTargetUserResponse {
	return nil
}

func (rc *BulkAddInactionTargetUserController) OnError(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse, err error) *fs.BulkAddInactionTargetUserResponse {
	return nil
}
