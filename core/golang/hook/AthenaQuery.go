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
	"code.nurture.farm/platform/CampaignService/core/golang/hook/aws"
	"context"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type AthenaQueryInterface interface {
	OnRequest(ctx context.Context, request *fs.AthenaQueryRequest) *fs.AthenaQueryResponse
	OnError(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse, err error) *fs.AthenaQueryResponse
	OnResponse(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse) *fs.AthenaQueryResponse
}

type GenericAthenaQueryExecutor struct {
	AthenaQueryInterface AthenaQueryInterface
}

type AthenaQueryController struct {
}

var AthenaQueryExecutor *GenericAthenaQueryExecutor

func (ge *GenericAthenaQueryExecutor) OnRequest(ctx context.Context, request *fs.AthenaQueryRequest) *fs.AthenaQueryResponse {
	return ge.AthenaQueryInterface.OnRequest(ctx, request)
}

func (ge *GenericAthenaQueryExecutor) OnResponse(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse) *fs.AthenaQueryResponse {
	return ge.AthenaQueryInterface.OnResponse(ctx, request, response)
}

func (ge *GenericAthenaQueryExecutor) OnError(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse, err error) *fs.AthenaQueryResponse {
	return ge.AthenaQueryInterface.OnError(ctx, request, response, err)
}

func (rc *AthenaQueryController) OnRequest(ctx context.Context, request *fs.AthenaQueryRequest) *fs.AthenaQueryResponse {

	logger.Info("AddNewCampaignExecutor OnRequest hook started", zap.Any("request", request))
	numRecords, err := executeAthenaQuery(ctx, request.AthenaQuery)
	if err != nil {
		logger.Error("AddNewCampaignExecutor OnRequest hook, Error in running Athena query", zap.Any("query", request.AthenaQuery))
		return &fs.AthenaQueryResponse{
			Status: &common.RequestStatusResult{
				Status:        common.RequestStatus_INTERNAL_ERROR,
				ErrorMessages: []string{"Error in executing athena query"},
			},
		}
	}
	logger.Info("AddNewCampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.AthenaQueryResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},
		Count: numRecords,
	}
}

func executeAthenaQuery(ctx context.Context, query string) (int64, error) {
	var res int64
	_, actorDetailsData, err := aws.ExecuteAthenaQuery(ctx, query, -1)
	logger.Info("executeAthenaQuery method, Total number of records received from Athena:",
		zap.Any("numRecords", len(actorDetailsData)))
	if err != nil {
		logger.Error("Error in running athena query", zap.Any("query", query), zap.Any("error", err))
		return res, err
	}
	res = cast.ToInt64(len(actorDetailsData) - 1) // 1st row will always be columns, so -1
	return res, nil
}

func (rc *AthenaQueryController) OnResponse(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse) *fs.AthenaQueryResponse {
	return nil
}

func (rc *AthenaQueryController) OnError(ctx context.Context, request *fs.AthenaQueryRequest, response *fs.AthenaQueryResponse, err error) *fs.AthenaQueryResponse {
	return nil
}
