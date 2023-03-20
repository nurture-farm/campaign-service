package service

import (
	"code.nurture.farm/platform/CampaignService/core/golang/hook"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/executor"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/metrics"
	"context"
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger = getLogger()

func getLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

const (
	MULTI_REQUEST = "NF_RG_MULTI_REQUEST"
)

func ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) *fs.AddCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddCampaign_Metrics, "AddCampaign", &err, ctx)
	logger.Info("Serving AddCampaign request", zap.Any("request", request))

	onRequestResponse := hook.AddCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddCampaign request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddCampaign(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddCampaign request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddCampaign_Error_Metrics, err, ctx)
		response = &fs.AddCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.AddCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddCampaign request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddCampaignBulk(ctx context.Context, request *fs.BulkAddCampaignRequest) *fs.BulkAddCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddCampaign_Metrics, "AddCampaignBulk", &err, ctx)
	logger.Info("Serving ExecuteAddCampaignBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddCampaignBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddCampaignBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddCampaignBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddCampaign_Error_Metrics, err, ctx)
		response = &fs.BulkAddCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddCampaignBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest) (*fs.UpdateCampaignResponse, *fs.AddNewCampaignRequest) {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.UpdateCampaign_Metrics, "UpdateCampaign", &err, ctx)
	logger.Info("Serving UpdateCampaign request", zap.Any("request", request))

	onRequestResponse, addNewCampaignRequest := hook.UpdateCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteUpdateCampaign request", zap.Any("request", request))
		return response, addNewCampaignRequest
	}

	response, err := executor.RequestExecutor.ExecuteUpdateCampaign(ctx, request)
	if err != nil {
		logger.Error("ExecuteUpdateCampaign request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.UpdateCampaign_Error_Metrics, err, ctx)
		response = &fs.UpdateCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.UpdateCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response, nil
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteUpdateCampaign request served successfully!", zap.Any("request", request))
	return response, nil
}

func ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) *fs.AddCampaignTemplateResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddCampaignTemplate_Metrics, "AddCampaignTemplate", &err, ctx)
	logger.Info("Serving AddCampaignTemplate request", zap.Any("request", request))

	onRequestResponse := hook.AddCampaignTemplateExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddCampaignTemplate request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddCampaignTemplate(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddCampaignTemplate request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddCampaignTemplate_Error_Metrics, err, ctx)
		response = &fs.AddCampaignTemplateResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.AddCampaignTemplateExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddCampaignTemplate request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddCampaignTemplateBulk(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) *fs.BulkAddCampaignTemplateResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddCampaignTemplate_Metrics, "AddCampaignTemplateBulk", &err, ctx)
	logger.Info("Serving ExecuteAddCampaignTemplateBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddCampaignTemplateExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddCampaignTemplateBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddCampaignTemplateBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddCampaignTemplateBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddCampaignTemplate_Error_Metrics, err, ctx)
		response = &fs.BulkAddCampaignTemplateResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddCampaignTemplateExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddCampaignTemplateBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddNewCampaign(ctx context.Context, request *fs.AddNewCampaignRequest) *fs.AddNewCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddNewCampaign_Metrics, "AddNewCampaign", &err, ctx)

	if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
		logger.Info("Serving AddNewCampaign request", zap.Any("addCampaignRequest", request.AddCampaignRequest))
	} else {
		logger.Info("Serving AddNewCampaign request", zap.Any("request", request))
	}

	onRequestResponse := hook.AddNewCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
			logger.Info("Skipping ExecuteAddNewCampaign request", zap.Any("addCampaignRequest", request.AddCampaignRequest))
		} else {
			logger.Info("SSkipping ExecuteAddNewCampaign request", zap.Any("request", request))
		}
		return response
	}

	if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
		logger.Info("ExecuteAddNewCampaign request served successfully!", zap.Any("addCampaignRequest", request.AddCampaignRequest))
	} else {
		logger.Info("SExecuteAddNewCampaign request served successfully!", zap.Any("request", request))
	}
	return nil
}

func ExecuteAddNewCampaignBulk(ctx context.Context, request *fs.BulkAddNewCampaignRequest) *fs.BulkAddNewCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddNewCampaign_Metrics, "AddNewCampaignBulk", &err, ctx)
	logger.Info("Serving ExecuteAddNewCampaignBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddNewCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddNewCampaignBulk request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteAddNewCampaignBulk request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteCampaign(ctx context.Context, request *fs.CampaignRequest) *fs.CampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.Campaign_Metrics, "Campaign", &err, ctx)
	logger.Info("Serving ExecuteCampaign request", zap.Any("request", request))

	onRequestResponse := hook.CampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteCampaign request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteCampaign request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) *fs.FindCampaignByIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindCampaignById_Metrics, "FindCampaignById", &err, ctx)
	logger.Info("Serving FindCampaignById request", zap.Any("request", request))

	onRequestResponse := hook.FindCampaignByIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindCampaignById request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindCampaignById request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindCampaignById_Error_Metrics, err, ctx)
		response = &fs.FindCampaignByIdResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.FindCampaignByIdExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.FindCampaignByIdExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeFindCampaignById request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindCampaignById request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) *fs.FindCampaignTemplateByIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindCampaignTemplateById_Metrics, "FindCampaignTemplateById", &err, ctx)
	logger.Info("Serving FindCampaignTemplateById request", zap.Any("request", request))

	onRequestResponse := hook.FindCampaignTemplateByIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindCampaignTemplateById request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindCampaignTemplateById request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindCampaignTemplateById_Error_Metrics, err, ctx)
		response = &fs.FindCampaignTemplateByIdResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.FindCampaignTemplateByIdExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.FindCampaignTemplateByIdExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeFindCampaignTemplateById request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindCampaignTemplateById request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) *fs.AddTargetUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddTargetUser_Metrics, "AddTargetUser", &err, ctx)
	logger.Info("Serving AddTargetUser request", zap.Any("request", request))

	onRequestResponse := hook.AddTargetUserExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddTargetUser request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddTargetUser(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddTargetUser request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddTargetUser_Error_Metrics, err, ctx)
		response = &fs.AddTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.AddTargetUserExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddTargetUser request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddTargetUserBulk(ctx context.Context, request *fs.BulkAddTargetUserRequest) *fs.BulkAddTargetUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddTargetUser_Metrics, "AddTargetUserBulk", &err, ctx)
	logger.Info("Serving ExecuteAddTargetUserBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddTargetUserExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddTargetUserBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddTargetUserBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddTargetUserBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddTargetUser_Error_Metrics, err, ctx)
		response = &fs.BulkAddTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddTargetUserExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddTargetUserBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) *fs.FindTargetUserByIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindTargetUserById_Metrics, "FindTargetUserById", &err, ctx)
	logger.Info("Serving FindTargetUserById request", zap.Any("request", request))

	onRequestResponse := hook.FindTargetUserByIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindTargetUserById request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteFindTargetUserById(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindTargetUserById request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindTargetUserById_Error_Metrics, err, ctx)
		response = &fs.FindTargetUserByIdResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}
		onErrorResponse := hook.FindTargetUserByIdExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.FindTargetUserByIdExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeFindTargetUserById request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindTargetUserById request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) *fs.AddInactionTargetUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddInactionTargetUser_Metrics, "AddInactionTargetUser", &err, ctx)
	logger.Info("Serving AddInactionTargetUser request", zap.Any("request", request))

	onRequestResponse := hook.AddInactionTargetUserExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddInactionTargetUser request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddInactionTargetUser(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddInactionTargetUser request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddInactionTargetUser_Error_Metrics, err, ctx)
		response = &fs.AddInactionTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.AddInactionTargetUserExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddInactionTargetUser request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddInactionTargetUserBulk(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) *fs.BulkAddInactionTargetUserResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddInactionTargetUser_Metrics, "AddInactionTargetUserBulk", &err, ctx)
	logger.Info("Serving ExecuteAddInactionTargetUserBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddInactionTargetUserExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddInactionTargetUserBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddInactionTargetUserBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddInactionTargetUserBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddInactionTargetUser_Error_Metrics, err, ctx)
		response = &fs.BulkAddInactionTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddInactionTargetUserExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddInactionTargetUserBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) *fs.FindInactionTargetUserByCampaignIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindInactionTargetUserByCampaignId_Metrics, "FindInactionTargetUserByCampaignId", &err, ctx)
	logger.Info("Serving FindInactionTargetUserByCampaignId request", zap.Any("request", request))

	onRequestResponse := hook.FindInactionTargetUserByCampaignIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindInactionTargetUserByCampaignId request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteFindInactionTargetUserByCampaignId(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindInactionTargetUserByCampaignId request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindInactionTargetUserByCampaignId_Error_Metrics, err, ctx)
		response = &fs.FindInactionTargetUserByCampaignIdResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.FindInactionTargetUserByCampaignIdExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.FindInactionTargetUserByCampaignIdExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeFindInactionTargetUserByCampaignId request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindInactionTargetUserByCampaignId request served successfully!", zap.Any("request", request))
	return response
}

func Execute(ctx context.Context, request *fs.MultiRequests) *fs.MultiResponses {

	/*var err error
	defer executor.PushToRequestMetrics()(MULTI_REQUEST,&err,ctx)
	logger.Info("Serving Execute request", zap.Any("request", request))

	response := ExecuteRequestExecutor.onRequest(ctx, request)
	if response != nil {
		err = response.(error)
	}
	if err != nil {
		logger.Error("Execute bad request", zap.Error(err))
		return &fs.MultiResponses{
			Status: &fs.Status{
				Status: fs.StatusCode_INVALID_REQUEST,
			},
		}
	}

	responses := []*fs.Response{}
	errs := executor.Execute(ctx, request)
	for _, err := range errs {
		if err != nil {
			logger.Error("Execute request failed", zap.Error(err))
			response := &fs.Response{
				Status: &fs.Status{
					Status: fs.StatusCode_DB_FAILURE,
				},
			}
			responses = append(responses, response)
		}
	}

	//OnDataLogic can be added here
	//On Respponse logic can be added here

	logger.Info("Execute request served successfully!", zap.Any("request", request))
	return &fs.MultiResponses{
		Status: &fs.Status{
			Status: fs.StatusCode_SUCCESS,
		},
		Response: responses,
	}*/
	return nil
}

func ExecuteTestNewCampaign(ctx context.Context, request *fs.TestNewCampaignRequest) *fs.TestNewCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.TestNewCampaign_Metrics, "TestNewCampaign", &err, ctx)
	logger.Info("Serving ExecuteTestNewCampaign request", zap.Any("request", request))

	onRequestResponse := hook.TestNewCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteTestNewCampaign request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteTestNewCampaign request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteAthenaQuery(ctx context.Context, request *fs.AthenaQueryRequest) *fs.AthenaQueryResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AthenaQuery_Metrics, "AthenaQuery", &err, ctx)
	logger.Info("Serving ExecuteAthenaQuery request", zap.Any("request", request))

	onRequestResponse := hook.AthenaQueryExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAthenaQuery request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteAthenaQuery request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteFilterCampaigns(ctx context.Context, request *fs.FilterCampaignRequest) *fs.FilterCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FilterCampaign_Metrics, "AthenaQuery", &err, ctx)
	logger.Info("Serving ExecuteFilterCampaign request", zap.Any("request", request))

	onRequestResponse := hook.FilterCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFilterCampaign request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteFilterCampaign request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteTestCampaignById(ctx context.Context, request *fs.TestCampaignByIdRequest) *fs.TestCampaignByIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.TestCampaignById_Metrics, "TestCampaignById", &err, ctx)
	logger.Info("Serving ExecuteTestCampaignById request", zap.Any("request", request))

	onRequestResponse := hook.TestCampaignByIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteTestCampaignById request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteTestCampaignById request served successfully!", zap.Any("request", request))
	return nil
}

func ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) *fs.GetDynamicDataByKeyResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.GetDynamicDataByKey_Metrics, "GetDynamicDataByKey", &err, ctx)
	logger.Info("Serving GetDynamicDataByKey request", zap.Any("request", request))

	onRequestResponse := hook.GetDynamicDataByKeyExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteGetDynamicDataByKey request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteGetDynamicDataByKey(ctx, request)
	if err != nil {
		logger.Error("ExecuteGetDynamicDataByKey request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.GetDynamicDataByKey_Error_Metrics, err, ctx)
		response = &fs.GetDynamicDataByKeyResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.GetDynamicDataByKeyExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.GetDynamicDataByKeyExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeGetDynamicDataByKey request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here
	logger.Info("ExecuteGetDynamicDataByKey request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest) *fs.AddDynamicDataResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddDynamicData_Metrics, "AddDynamicData", &err, ctx)
	logger.Info("Serving AddDynamicData request", zap.Any("request", request))

	onRequestResponse := hook.AddDynamicDataExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddDynamicData request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddDynamicData(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddDynamicData request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddDynamicData_Error_Metrics, err, ctx)
		response = &fs.AddDynamicDataResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.AddDynamicDataExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddDynamicData request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddDynamicDataBulk(ctx context.Context, request *fs.BulkAddDynamicDataRequest) *fs.BulkAddDynamicDataResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddDynamicData_Metrics, "AddDynamicDataBulk", &err, ctx)
	logger.Info("Serving ExecuteAddDynamicDataBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddDynamicDataExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddDynamicDataBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddDynamicDataBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddDynamicDataBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddDynamicData_Error_Metrics, err, ctx)
		response = &fs.BulkAddDynamicDataResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddDynamicDataExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddDynamicDataBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteFindQueryCampaign(ctx context.Context, request *fs.FindQueryCampaignRequest) *fs.FindQueryCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindQueryCampaign_Metrics, "FindQueryCampaign", &err, ctx)
	logger.Info("Serving FindQueryCampaign request", zap.Any("request", request))

	onRequestResponse := hook.FindQueryCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindQueryCampaign request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteFindQueryCampaign(ctx, request)
	if err != nil {
		logger.Error("ExecuteFindQueryCampaign request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.FindQueryCampaign_Error_Metrics, err, ctx)
		response = &fs.FindQueryCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.FindQueryCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	onDataResponse := hook.FindQueryCampaignExecutor.OnData(ctx, request, response)
	if onDataResponse != nil {
		response := onDataResponse
		logger.Info("Returning OnData response for ExecuteeFindQueryCampaign request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindQueryCampaign request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddQueryCampaign(ctx context.Context, request *fs.AddQueryCampaignRequest) *fs.AddQueryCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.AddQueryCampaign_Metrics, "AddQueryCampaign", &err, ctx)
	logger.Info("Serving AddQueryCampaign request", zap.Any("request", request))

	onRequestResponse := hook.AddQueryCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddQueryCampaign request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddQueryCampaign(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddQueryCampaign request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddQueryCampaign_Error_Metrics, err, ctx)
		response = &fs.AddQueryCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}
		onErrorResponse := hook.AddQueryCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddQueryCampaign request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteAddQueryCampaignBulk(ctx context.Context, request *fs.BulkAddQueryCampaignRequest) *fs.BulkAddQueryCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.BulkAddQueryCampaign_Metrics, "AddQueryCampaignBulk", &err, ctx)
	logger.Info("Serving ExecuteAddQueryCampaignBulk request", zap.Any("request", request))

	onRequestResponse := hook.BulkAddQueryCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteAddQueryCampaignBulk request", zap.Any("request", request))
		return response
	}

	response, err := executor.RequestExecutor.ExecuteAddQueryCampaignBulk(ctx, request)
	if err != nil {
		logger.Error("ExecuteAddQueryCampaignBulk request failed", zap.Error(err))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.AddQueryCampaign_Error_Metrics, err, ctx)
		response = &fs.BulkAddQueryCampaignResponse{
			Status: &common.RequestStatusResult{
				Status:    common.RequestStatus_INTERNAL_ERROR,
				ErrorCode: common.ErrorCode_DATABASE_ERROR,
			},
		}

		onErrorResponse := hook.BulkAddQueryCampaignExecutor.OnError(ctx, request, nil, err)
		if onErrorResponse != nil {
			response = onErrorResponse
		}
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteAddQueryCampaignBulk request served successfully!", zap.Any("request", request))
	return response
}

func ExecuteScheduleUserJourneyCampaign(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) *fs.ScheduleUserJourneyCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.ScheduleUserJourneyCampaign_Metrics, "ScheduleUserJourneyCampaign", &err, ctx)
	logger.Info("Serving ExecuteScheduleUserJourneyCampaign request", zap.Any("request", request))

	onRequestResponse := hook.ScheduleUserJourneyCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteScheduleUserJourneyCampaign request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteScheduleUserJourneyCampaign request served successfully!", zap.Any("request", request))
	return onRequestResponse
}

func ExecuteFindUserJourneyCampaignById(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) *fs.FindUserJourneyCampaignByIdResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FindUserJourneyCampaignById_Metrics, "FindUserJourneyCampaignById", &err, ctx)
	logger.Info("Serving ExecuteFindUserJourneyCampaignById request", zap.Any("request", request))

	onRequestResponse := hook.FindUserJourneyCampaignByIdExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFindUserJourneyCampaignById request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFindUserJourneyCampaignById request served successfully!", zap.Any("request", request))
	return onRequestResponse
}

func ExecuteFilterUserJourneyCampaigns(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) *fs.FilterUserJourneyCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.FilterUserJourneyCampaigns_Metrics, "ExecuteFilterUserJourneyCampaigns", &err, ctx)
	logger.Info("Serving ExecuteFilterUserJourneyCampaigns request", zap.Any("request", request))

	onRequestResponse := hook.FilterUserJourneyCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteFilterUserJourneyCampaigns request", zap.Any("request", request))
		return response
	}

	//On Respponse logic can be added here

	logger.Info("ExecuteFilterUserJourneyCampaigns request served successfully!", zap.Any("request", request))
	return onRequestResponse
}

func ExecuteUserJourneyCampaign(ctx context.Context, request *fs.UserJourneyCampaignRequest) *fs.UserJourneyCampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.UserJourneyCampaign_Metrics, "UserJourneyCamppaign", &err, ctx)
	logger.Info("Serving ExecuteUserJourneyCampaign request", zap.Any("request", request))

	onRequestResponse := hook.UserJourneyCampaignExecutor.OnRequest(ctx, request)
	if onRequestResponse != nil {
		response := onRequestResponse
		logger.Info("Skipping ExecuteUserJourneyCampaign request", zap.Any("request", request))
		return response
	}

	logger.Info("ExecuteUserJourneyCampaign request served successfully!", zap.Any("request", request))
	return nil
}
