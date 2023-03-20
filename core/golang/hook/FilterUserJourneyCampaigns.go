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
	query "github.com/nurture-farm/campaign-service/core/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
	"context"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"time"
)

type FilterUserJourneyCampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) *fs.FilterUserJourneyCampaignResponse
	OnData(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse
	OnError(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse, err error) *fs.FilterUserJourneyCampaignResponse
	OnResponse(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse
}

type GenericFilterUserJourneyCampaignExecutor struct {
	FilterUserJourneyCampaignInterface FilterUserJourneyCampaignInterface
}

type FilterUserJourneyCampaignController struct {
}

var FilterUserJourneyCampaignExecutor *GenericFilterUserJourneyCampaignExecutor

func (ge *GenericFilterUserJourneyCampaignExecutor) OnRequest(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) *fs.FilterUserJourneyCampaignResponse {
	return ge.FilterUserJourneyCampaignInterface.OnRequest(ctx, request)
}

func (ge *GenericFilterUserJourneyCampaignExecutor) OnResponse(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse {
	return ge.FilterUserJourneyCampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericFilterUserJourneyCampaignExecutor) OnData(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse {
	return ge.FilterUserJourneyCampaignInterface.OnData(ctx, request, response)
}

func (ge *GenericFilterUserJourneyCampaignExecutor) OnError(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse, err error) *fs.FilterUserJourneyCampaignResponse {
	return ge.FilterUserJourneyCampaignInterface.OnError(ctx, request, response, err)
}

func (rc *FilterUserJourneyCampaignController) OnRequest(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) *fs.FilterUserJourneyCampaignResponse {

	startTime, endTime, err := validateFilterUserJourneyCampaignRequest(request)
	if err != nil {
		return mapToFilterUserJourneyCampaignResponse(request, common.RequestStatus_BAD_INPUT, err, nil)
	}
	if request.PageNumber == 0 {
		request.PageNumber = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}
	var rows = entsql.Rows{}
	query, args := query.GenerateFilterUserJourneyCampaignsQuery(request, startTime, endTime)
	err = executor.Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("FilterUserJourneyCampaignController, Error in executing query", zap.Any("request", request),
			zap.Any("query", query))
		return mapToFilterUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, nil)
	}
	models := []*FindUserJourneyCampaigVO{}
	for rows.Next() {
		model := FindUserJourneyCampaigVO{}
		err := rows.Scan(&model.Id, &model.Namespace, &model.Name, &model.Status, &model.CreatedAt, &model.Count, &model.TargetUsersStatus)
		if err != nil {
			logger.Error("FilterUserJourneyCampaignController, Error in scanning rows", zap.Error(err),
				zap.Any("request", request), zap.Any("model", model))
			return mapToFilterUserJourneyCampaignResponse(request, common.RequestStatus_INTERNAL_ERROR, err, nil)
		}
		models = append(models, &model)
	}
	records := MakeFindUserJourneyCampaignVO(models)
	var status common.RequestStatus
	if records == nil || len(records) == 0 {
		status = common.RequestStatus_NOT_FOUND
	} else {
		status = common.RequestStatus_SUCCESS
	}
	return mapToFilterUserJourneyCampaignResponse(request, status, nil, records)
}

func (rc *FilterUserJourneyCampaignController) OnResponse(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse {
	return nil
}

func (rc *FilterUserJourneyCampaignController) OnData(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse) *fs.FilterUserJourneyCampaignResponse {
	return nil
}

func (rc *FilterUserJourneyCampaignController) OnError(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest, response *fs.FilterUserJourneyCampaignResponse, err error) *fs.FilterUserJourneyCampaignResponse {
	return nil
}

func validateFilterUserJourneyCampaignRequest(request *fs.FilterUserJourneyCampaignRequest) (startTime time.Time, endTime time.Time, err error) {

	logger.Info("FilterUserJourneyCampaignController OnRequest hook started", zap.Any("request", request))
	startTime, err = parseTimeStamp(request.StartTime)
	if err != nil {
		startTime = time.Time{}
		err = nil
	}
	if !startTime.IsZero() {
		startTime = convertIstToUtc(startTime)
	}
	endTime, err = parseTimeStamp(request.EndTime)
	if err != nil {
		endTime = time.Time{}
		err = nil
	}
	if !endTime.IsZero() {
		endTime = convertIstToUtc(endTime)
	}
	//if request.SearchFilter == EMPTY && request.Status == common.CampaignStatus_NO_CAMPAGIN_STATUS && request.Namespace == common.NameSpace_NO_NAMESPACE &&
	//	startTime.IsZero() && endTime.IsZero() {
	//	err = errors.New(INVALID_REQUEST)
	//}
	return
}

func mapToFilterUserJourneyCampaignResponse(request *fs.FilterUserJourneyCampaignRequest, status common.RequestStatus, err error, records []*fs.FilterUserJourneyCampaignResponseRecord) *fs.FilterUserJourneyCampaignResponse {

	if err != nil {
		switch status {
		case common.RequestStatus_BAD_INPUT:
			logger.Error("FilterUserJourneyCampaignController request failed, invalid request", zap.Any("request", request), zap.Error(err))
		case common.RequestStatus_INTERNAL_ERROR:
			logger.Error("FilterUserJourneyCampaignController request failed, internal error", zap.Any("request", request), zap.Error(err))
		}
		return &fs.FilterUserJourneyCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        status,
				ErrorMessages: []string{err.Error()},
			},
		}
	}
	logger.Info("FilterUserJourneyCampaignController OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.FilterUserJourneyCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: status,
		},
		Count:   cast.ToInt32(len(records)),
		Records: records,
	}
}
