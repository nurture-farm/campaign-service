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
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"context"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type FilterCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.FilterCampaignRequest) *fs.FilterCampaignResponse
	OnError(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse, err error) *fs.FilterCampaignResponse
	OnResponse(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse) *fs.FilterCampaignResponse
}

type GenericFilterCampaignExecutor struct {
	FilterCampaignInterface FilterCampaignInterface
}

type FilterCampaignController struct {
}

var FilterCampaignExecutor *GenericFilterCampaignExecutor

func (ge *GenericFilterCampaignExecutor) OnRequest(ctx context.Context, request *fs.FilterCampaignRequest) *fs.FilterCampaignResponse {
	return ge.FilterCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericFilterCampaignExecutor) OnResponse(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse) *fs.FilterCampaignResponse {
	return ge.FilterCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFilterCampaignExecutor) OnError(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse, err error) *fs.FilterCampaignResponse {
	return ge.FilterCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *FilterCampaignController) OnRequest(ctx context.Context, request *fs.FilterCampaignRequest) *fs.FilterCampaignResponse {

	logger.Info("TestNewCampaignExecutor OnRequest hook started", zap.Any("request", request))

	filterCampaignResponse, err := executor.RequestExecutor.ExecuteFilterCampaigns(ctx, request)
	if err != nil {
		logger.Error("TestNewCampaignExecutor OnRequest hook, Error in ExecuteFilterCampaigns", zap.Any("error", err), zap.Any("request", request))
		return &fs.FilterCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_INTERNAL_ERROR,
			},
		}
	}

	logger.Info("TestNewCampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.FilterCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},
		Count:   cast.ToInt32(len(filterCampaignResponse.Records)),
		Records: filterCampaignResponse.Records,
	}
}

func (rc *FilterCampaignController) OnResponse(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse) *fs.FilterCampaignResponse {
	return nil
}

func (rc *FilterCampaignController) OnError(ctx context.Context, request *fs.FilterCampaignRequest, response *fs.FilterCampaignResponse, err error) *fs.FilterCampaignResponse {
	return nil
}
