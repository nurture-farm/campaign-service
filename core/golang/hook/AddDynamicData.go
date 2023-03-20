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

type AddDynamicDataInterface interface {
	OnRequest(ctx context.Context, request *fs.AddDynamicDataRequest) *fs.AddDynamicDataResponse
	OnError(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse, err error) *fs.AddDynamicDataResponse
	OnResponse(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse) *fs.AddDynamicDataResponse
}

type AddDynamicDataBulkInterface interface {
	OnRequest(ctx context.Context, request *fs.BulkAddDynamicDataRequest) *fs.BulkAddDynamicDataResponse
	OnError(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse, err error) *fs.BulkAddDynamicDataResponse
	OnResponse(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse) *fs.BulkAddDynamicDataResponse
}

type GenericAddDynamicDataExecutor struct {
	AddDynamicDataInterface AddDynamicDataInterface
}

type GenericAddDynamicDataExecutorBulk struct {
	AddDynamicDataBulkInterface AddDynamicDataBulkInterface
}

type AddDynamicDataController struct {
}

type BulkAddDynamicDataController struct {
}

var AddDynamicDataExecutor *GenericAddDynamicDataExecutor
var BulkAddDynamicDataExecutor *GenericAddDynamicDataExecutorBulk

func (ge *GenericAddDynamicDataExecutor) OnRequest(ctx context.Context, request *fs.AddDynamicDataRequest) *fs.AddDynamicDataResponse {
	return ge.AddDynamicDataInterface.OnRequest(ctx, request)
}

func (ge *GenericAddDynamicDataExecutor) OnResponse(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse) *fs.AddDynamicDataResponse {
	return ge.AddDynamicDataInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddDynamicDataExecutor) OnError(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse, err error) *fs.AddDynamicDataResponse {
	return ge.AddDynamicDataInterface.OnError(ctx, request, response, err)
}

func (ge *GenericAddDynamicDataExecutorBulk) OnRequest(ctx context.Context, request *fs.BulkAddDynamicDataRequest) *fs.BulkAddDynamicDataResponse {
	return ge.AddDynamicDataBulkInterface.OnRequest(ctx, request)
}

func (ge *GenericAddDynamicDataExecutorBulk) OnResponse(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse) *fs.BulkAddDynamicDataResponse {
	return ge.AddDynamicDataBulkInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAddDynamicDataExecutorBulk) OnError(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse, err error) *fs.BulkAddDynamicDataResponse {
	return ge.AddDynamicDataBulkInterface.OnError(ctx, request, response, err)
}

func (rc *AddDynamicDataController) OnRequest(ctx context.Context, request *fs.AddDynamicDataRequest) *fs.AddDynamicDataResponse {
	logger.Info("AddDynamicDataExecuter OnRequest hook started", zap.Any("request", request))
	if request.CampaignId == 0 || request.DynamicKey == "" {
		logger.Error("AddDynamicDataExecuter OnRequest hook, Invalid request", zap.Any("request", request))
		return &fs.AddDynamicDataResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_BAD_INPUT,
			},
		}
	}
	logger.Info("AddDynamicDataExecuter OnRequest hook completed successfully", zap.Any("request", request))
	return nil

}

func (rc *AddDynamicDataController) OnResponse(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse) *fs.AddDynamicDataResponse {
	return nil
}

func (rc *AddDynamicDataController) OnError(ctx context.Context, request *fs.AddDynamicDataRequest, response *fs.AddDynamicDataResponse, err error) *fs.AddDynamicDataResponse {
	return nil
}

func (rc *BulkAddDynamicDataController) OnRequest(ctx context.Context, request *fs.BulkAddDynamicDataRequest) *fs.BulkAddDynamicDataResponse {
	return nil
}

func (rc *BulkAddDynamicDataController) OnResponse(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse) *fs.BulkAddDynamicDataResponse {
	return nil
}

func (rc *BulkAddDynamicDataController) OnError(ctx context.Context, request *fs.BulkAddDynamicDataRequest, response *fs.BulkAddDynamicDataResponse, err error) *fs.BulkAddDynamicDataResponse {
	return nil
}
