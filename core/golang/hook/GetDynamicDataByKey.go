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

type GetDynamicDataByKeyInterface interface {
	OnRequest(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) *fs.GetDynamicDataByKeyResponse
	OnData(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse
	OnError(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse, err error) *fs.GetDynamicDataByKeyResponse
	OnResponse(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse
}

type GenericGetDynamicDataByKeyExecutor struct {
	GetDynamicDataByKeyInterface GetDynamicDataByKeyInterface
}

type GetDynamicDataByKeyController struct {
}

var GetDynamicDataByKeyExecutor *GenericGetDynamicDataByKeyExecutor

func (ge *GenericGetDynamicDataByKeyExecutor) OnRequest(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) *fs.GetDynamicDataByKeyResponse {
	logger.Info("GetDynamicDataByKeyExecutor OnRequest hook started", zap.Any("request", request))

	if request.CampaignId == 0 || request.DynamicKey == "" {
		logger.Error("GetDynamicDataByKeyExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.GetDynamicDataByKeyResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	logger.Info("GetDynamicDataByKeyExecuter OnRequest hook completed successfully", zap.Any("request", request))
	return nil
}

func (ge *GenericGetDynamicDataByKeyExecutor) OnResponse(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse {
	return ge.GetDynamicDataByKeyInterface.OnResponse(ctx, request, response)
}

func (ge *GenericGetDynamicDataByKeyExecutor) OnData(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse {
	return ge.GetDynamicDataByKeyInterface.OnData(ctx, request, response)
}

func (ge *GenericGetDynamicDataByKeyExecutor) OnError(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse, err error) *fs.GetDynamicDataByKeyResponse {
	return ge.GetDynamicDataByKeyInterface.OnError(ctx, request, response, err)
}

func (rc *GetDynamicDataByKeyController) OnRequest(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) *fs.GetDynamicDataByKeyResponse {
	return nil
}

func (rc *GetDynamicDataByKeyController) OnResponse(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse {
	return nil
}

func (rc *GetDynamicDataByKeyController) OnData(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse) *fs.GetDynamicDataByKeyResponse {
	return nil
}

func (rc *GetDynamicDataByKeyController) OnError(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, response *fs.GetDynamicDataByKeyResponse, err error) *fs.GetDynamicDataByKeyResponse {
	return nil
}
