package service_test

//import (
//	"context"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
//	"testing"
//	"errors"
//	"github.com/prometheus/client_golang/prometheus"
//	"code.nurture.farm/platform/CampaignService/core/golang/hook"
//	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/executor"
//	"code.nurture.farm/platform/CampaignService/zerotouch/golang/service"
//	"code.nurture.farm/platform/CampaignService/zerotouch/golang/metrics"
//)
//
//type ExecutorMock struct {
//	mock.Mock
//}
//
//type MetricsMock struct {
//	mock.Mock
//}
//
//type HookAddCampaignMock struct {
//	mock.Mock
//}
//type HookAddCampaignBulkMock struct {
//	mock.Mock
//}
//type HookUpdateCampaignMock struct {
//	mock.Mock
//}
//type HookAddCampaignTemplateMock struct {
//	mock.Mock
//}
//type HookAddCampaignTemplateBulkMock struct {
//	mock.Mock
//}
//type HookAddNewCampaignMock struct {
//	mock.Mock
//}
//type HookAddNewCampaignBulkMock struct {
//	mock.Mock
//}
//type HookCampaignMock struct {
//	mock.Mock
//}
//type HookFindCampaignByIdMock struct {
//	mock.Mock
//}
//type HookFindCampaignTemplateByIdMock struct {
//	mock.Mock
//}
//type HookAddTargetUserMock struct {
//	mock.Mock
//}
//type HookAddTargetUserBulkMock struct {
//	mock.Mock
//}
//
//type HookAddInactionTargetUserMock struct {
//	mock.Mock
//}
//type HookAddInactionTargetUserBulkMock struct {
//	mock.Mock
//}
//
//type HookFindInactionTargetUserByCampaignIdMock struct {
//	mock.Mock
//}
//
//func (ms *MetricsMock) PushToSummarytMetrics() func(*prometheus.SummaryVec, string, *error, context.Context) {
//	return func(request *prometheus.SummaryVec, methodName string, err *error, ctx context.Context) {
//
//		return
//	}
//}
//func (ms *MetricsMock) PushToErrorCounterMetrics() func(*prometheus.CounterVec, error, context.Context) {
//	return func(request *prometheus.CounterVec, err error, ctx context.Context) {
//		return
//	}
//}
//
//func (rc *HookAddCampaignMock) OnRequest(ctx context.Context, request *fs.AddCampaignRequest) *fs.AddCampaignResponse {
//	var AddCampaignResponse *fs.AddCampaignResponse
//	args := rc.Called(ctx, request)
//	mockedAddCampaignResponse := args.Get(0)
//	if mockedAddCampaignResponse != nil {
//		AddCampaignResponse = mockedAddCampaignResponse.(*fs.AddCampaignResponse)
//	}
//	return AddCampaignResponse
//}
//
//func (rc *HookAddCampaignMock) OnResponse(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse) *fs.AddCampaignResponse {
//	var AddCampaignResponse *fs.AddCampaignResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddCampaignResponse := args.Get(0)
//	if mockedAddCampaignResponse != nil {
//		AddCampaignResponse = mockedAddCampaignResponse.(*fs.AddCampaignResponse)
//	}
//	return AddCampaignResponse
//}
//
//func (rc *HookAddCampaignMock) OnError(ctx context.Context, request *fs.AddCampaignRequest, response *fs.AddCampaignResponse, err error) *fs.AddCampaignResponse {
//	var AddCampaignResponse *fs.AddCampaignResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddCampaignResponse := args.Get(0)
//	if mockedAddCampaignResponse != nil {
//		AddCampaignResponse = mockedAddCampaignResponse.(*fs.AddCampaignResponse)
//	}
//	return AddCampaignResponse
//}
//func (rc *HookAddCampaignBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddCampaignRequest) *fs.BulkAddCampaignResponse {
//	var AddCampaignBulkResponse *fs.BulkAddCampaignResponse
//	args := rc.Called(ctx, request)
//	mockedAddCampaignBulkResponse := args.Get(0)
//	if mockedAddCampaignBulkResponse != nil {
//		AddCampaignBulkResponse = mockedAddCampaignBulkResponse.(*fs.BulkAddCampaignResponse)
//	}
//	return AddCampaignBulkResponse
//}
//
//func (rc *HookAddCampaignBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse) *fs.BulkAddCampaignResponse {
//	var AddCampaignBulkResponse *fs.BulkAddCampaignResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddCampaignBulkResponse := args.Get(0)
//	if mockedAddCampaignBulkResponse != nil {
//		AddCampaignBulkResponse = mockedAddCampaignBulkResponse.(*fs.BulkAddCampaignResponse)
//	}
//	return AddCampaignBulkResponse
//}
//
//func (rc *HookAddCampaignBulkMock) OnError(ctx context.Context, request *fs.BulkAddCampaignRequest, response *fs.BulkAddCampaignResponse, err error) *fs.BulkAddCampaignResponse {
//	var AddCampaignBulkResponse *fs.BulkAddCampaignResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddCampaignBulkResponse := args.Get(0)
//	if mockedAddCampaignBulkResponse != nil {
//		AddCampaignBulkResponse = mockedAddCampaignBulkResponse.(*fs.BulkAddCampaignResponse)
//	}
//	return AddCampaignBulkResponse
//}
//func (rc *HookUpdateCampaignMock) OnRequest(ctx context.Context, request *fs.UpdateCampaignRequest) *fs.UpdateCampaignResponse {
//	var UpdateCampaignResponse *fs.UpdateCampaignResponse
//	args := rc.Called(ctx, request)
//	mockedUpdateCampaignResponse := args.Get(0)
//	if mockedUpdateCampaignResponse != nil {
//		UpdateCampaignResponse = mockedUpdateCampaignResponse.(*fs.UpdateCampaignResponse)
//	}
//	return UpdateCampaignResponse
//}
//
//func (rc *HookUpdateCampaignMock) OnResponse(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse) *fs.UpdateCampaignResponse {
//	var UpdateCampaignResponse *fs.UpdateCampaignResponse
//	args := rc.Called(ctx, request, response)
//	mockedUpdateCampaignResponse := args.Get(0)
//	if mockedUpdateCampaignResponse != nil {
//		UpdateCampaignResponse = mockedUpdateCampaignResponse.(*fs.UpdateCampaignResponse)
//	}
//	return UpdateCampaignResponse
//}
//
//func (rc *HookUpdateCampaignMock) OnError(ctx context.Context, request *fs.UpdateCampaignRequest, response *fs.UpdateCampaignResponse, err error) *fs.UpdateCampaignResponse {
//	var UpdateCampaignResponse *fs.UpdateCampaignResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedUpdateCampaignResponse := args.Get(0)
//	if mockedUpdateCampaignResponse != nil {
//		UpdateCampaignResponse = mockedUpdateCampaignResponse.(*fs.UpdateCampaignResponse)
//	}
//	return UpdateCampaignResponse
//}
//func (rc *HookAddCampaignTemplateMock) OnRequest(ctx context.Context, request *fs.AddCampaignTemplateRequest) *fs.AddCampaignTemplateResponse {
//	var AddCampaignTemplateResponse *fs.AddCampaignTemplateResponse
//	args := rc.Called(ctx, request)
//	mockedAddCampaignTemplateResponse := args.Get(0)
//	if mockedAddCampaignTemplateResponse != nil {
//		AddCampaignTemplateResponse = mockedAddCampaignTemplateResponse.(*fs.AddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateResponse
//}
//
//func (rc *HookAddCampaignTemplateMock) OnResponse(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse) *fs.AddCampaignTemplateResponse {
//	var AddCampaignTemplateResponse *fs.AddCampaignTemplateResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddCampaignTemplateResponse := args.Get(0)
//	if mockedAddCampaignTemplateResponse != nil {
//		AddCampaignTemplateResponse = mockedAddCampaignTemplateResponse.(*fs.AddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateResponse
//}
//
//func (rc *HookAddCampaignTemplateMock) OnError(ctx context.Context, request *fs.AddCampaignTemplateRequest, response *fs.AddCampaignTemplateResponse, err error) *fs.AddCampaignTemplateResponse {
//	var AddCampaignTemplateResponse *fs.AddCampaignTemplateResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddCampaignTemplateResponse := args.Get(0)
//	if mockedAddCampaignTemplateResponse != nil {
//		AddCampaignTemplateResponse = mockedAddCampaignTemplateResponse.(*fs.AddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateResponse
//}
//func (rc *HookAddCampaignTemplateBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) *fs.BulkAddCampaignTemplateResponse {
//	var AddCampaignTemplateBulkResponse *fs.BulkAddCampaignTemplateResponse
//	args := rc.Called(ctx, request)
//	mockedAddCampaignTemplateBulkResponse := args.Get(0)
//	if mockedAddCampaignTemplateBulkResponse != nil {
//		AddCampaignTemplateBulkResponse = mockedAddCampaignTemplateBulkResponse.(*fs.BulkAddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateBulkResponse
//}
//
//func (rc *HookAddCampaignTemplateBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse) *fs.BulkAddCampaignTemplateResponse {
//	var AddCampaignTemplateBulkResponse *fs.BulkAddCampaignTemplateResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddCampaignTemplateBulkResponse := args.Get(0)
//	if mockedAddCampaignTemplateBulkResponse != nil {
//		AddCampaignTemplateBulkResponse = mockedAddCampaignTemplateBulkResponse.(*fs.BulkAddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateBulkResponse
//}
//
//func (rc *HookAddCampaignTemplateBulkMock) OnError(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest, response *fs.BulkAddCampaignTemplateResponse, err error) *fs.BulkAddCampaignTemplateResponse {
//	var AddCampaignTemplateBulkResponse *fs.BulkAddCampaignTemplateResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddCampaignTemplateBulkResponse := args.Get(0)
//	if mockedAddCampaignTemplateBulkResponse != nil {
//		AddCampaignTemplateBulkResponse = mockedAddCampaignTemplateBulkResponse.(*fs.BulkAddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateBulkResponse
//}
//func (rc *HookAddNewCampaignMock) OnRequest(ctx context.Context, request *fs.AddNewCampaignRequest) *fs.AddNewCampaignResponse {
//	var AddNewCampaignResponse *fs.AddNewCampaignResponse
//	args := rc.Called(ctx, request)
//	mockedAddNewCampaignResponse := args.Get(0)
//	if mockedAddNewCampaignResponse != nil {
//		AddNewCampaignResponse = mockedAddNewCampaignResponse.(*fs.AddNewCampaignResponse)
//	}
//	return AddNewCampaignResponse
//}
//
//func (rc *HookAddNewCampaignMock) OnResponse(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse) *fs.AddNewCampaignResponse {
//	var AddNewCampaignResponse *fs.AddNewCampaignResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddNewCampaignResponse := args.Get(0)
//	if mockedAddNewCampaignResponse != nil {
//		AddNewCampaignResponse = mockedAddNewCampaignResponse.(*fs.AddNewCampaignResponse)
//	}
//	return AddNewCampaignResponse
//}
//
//func (rc *HookAddNewCampaignMock) OnError(ctx context.Context, request *fs.AddNewCampaignRequest, response *fs.AddNewCampaignResponse, err error) *fs.AddNewCampaignResponse {
//	var AddNewCampaignResponse *fs.AddNewCampaignResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddNewCampaignResponse := args.Get(0)
//	if mockedAddNewCampaignResponse != nil {
//		AddNewCampaignResponse = mockedAddNewCampaignResponse.(*fs.AddNewCampaignResponse)
//	}
//	return AddNewCampaignResponse
//}
//func (rc *HookAddNewCampaignBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddNewCampaignRequest) *fs.BulkAddNewCampaignResponse {
//	var AddNewCampaignBulkResponse *fs.BulkAddNewCampaignResponse
//	args := rc.Called(ctx, request)
//	mockedAddNewCampaignBulkResponse := args.Get(0)
//	if mockedAddNewCampaignBulkResponse != nil {
//		AddNewCampaignBulkResponse = mockedAddNewCampaignBulkResponse.(*fs.BulkAddNewCampaignResponse)
//	}
//	return AddNewCampaignBulkResponse
//}
//
//func (rc *HookAddNewCampaignBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse) *fs.BulkAddNewCampaignResponse {
//	var AddNewCampaignBulkResponse *fs.BulkAddNewCampaignResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddNewCampaignBulkResponse := args.Get(0)
//	if mockedAddNewCampaignBulkResponse != nil {
//		AddNewCampaignBulkResponse = mockedAddNewCampaignBulkResponse.(*fs.BulkAddNewCampaignResponse)
//	}
//	return AddNewCampaignBulkResponse
//}
//
//func (rc *HookAddNewCampaignBulkMock) OnError(ctx context.Context, request *fs.BulkAddNewCampaignRequest, response *fs.BulkAddNewCampaignResponse, err error) *fs.BulkAddNewCampaignResponse {
//	var AddNewCampaignBulkResponse *fs.BulkAddNewCampaignResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddNewCampaignBulkResponse := args.Get(0)
//	if mockedAddNewCampaignBulkResponse != nil {
//		AddNewCampaignBulkResponse = mockedAddNewCampaignBulkResponse.(*fs.BulkAddNewCampaignResponse)
//	}
//	return AddNewCampaignBulkResponse
//}
//func (rc *HookCampaignMock) OnRequest(ctx context.Context, request *fs.CampaignRequest) *fs.CampaignResponse {
//	var GetUserListResponse *fs.CampaignResponse
//	args := rc.Called(ctx, request)
//	mockedGetUserListResponse := args.Get(0)
//	if mockedGetUserListResponse != nil {
//		GetUserListResponse = mockedGetUserListResponse.(*fs.CampaignResponse)
//	}
//	return GetUserListResponse
//}
//
//func (rc *HookGetUserListMock) OnResponse(ctx context.Context, request *fs.GetUserListRequest, response *fs.GetUserListResponse) *fs.GetUserListResponse {
//	var GetUserListResponse *fs.GetUserListResponse
//	args := rc.Called(ctx, request, response)
//	mockedGetUserListResponse := args.Get(0)
//	if mockedGetUserListResponse != nil {
//		GetUserListResponse = mockedGetUserListResponse.(*fs.GetUserListResponse)
//	}
//	return GetUserListResponse
//}
//
//func (rc *HookGetUserListMock) OnError(ctx context.Context, request *fs.GetUserListRequest, response *fs.GetUserListResponse, err error) *fs.GetUserListResponse {
//	var GetUserListResponse *fs.GetUserListResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedGetUserListResponse := args.Get(0)
//	if mockedGetUserListResponse != nil {
//		GetUserListResponse = mockedGetUserListResponse.(*fs.GetUserListResponse)
//	}
//	return GetUserListResponse
//}
//
//func (rc *HookGetUserListMock) OnData(ctx context.Context, request *fs.GetUserListRequest, response *fs.GetUserListResponse) *fs.GetUserListResponse {
//	var GetUserListResponse *fs.GetUserListResponse
//	args := rc.Called(ctx, request, response)
//	mockedGetUserListResponse := args.Get(0)
//	if mockedGetUserListResponse != nil {
//		GetUserListResponse = mockedGetUserListResponse.(*fs.GetUserListResponse)
//	}
//	return GetUserListResponse
//}
//func (rc *HookFindCampaignByIdMock) OnRequest(ctx context.Context, request *fs.FindCampaignByIdRequest) *fs.FindCampaignByIdResponse {
//	var FindCampaignByIdResponse *fs.FindCampaignByIdResponse
//	args := rc.Called(ctx, request)
//	mockedFindCampaignByIdResponse := args.Get(0)
//	if mockedFindCampaignByIdResponse != nil {
//		FindCampaignByIdResponse = mockedFindCampaignByIdResponse.(*fs.FindCampaignByIdResponse)
//	}
//	return FindCampaignByIdResponse
//}
//
//func (rc *HookFindCampaignByIdMock) OnResponse(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
//	var FindCampaignByIdResponse *fs.FindCampaignByIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindCampaignByIdResponse := args.Get(0)
//	if mockedFindCampaignByIdResponse != nil {
//		FindCampaignByIdResponse = mockedFindCampaignByIdResponse.(*fs.FindCampaignByIdResponse)
//	}
//	return FindCampaignByIdResponse
//}
//
//func (rc *HookFindCampaignByIdMock) OnError(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse, err error) *fs.FindCampaignByIdResponse {
//	var FindCampaignByIdResponse *fs.FindCampaignByIdResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedFindCampaignByIdResponse := args.Get(0)
//	if mockedFindCampaignByIdResponse != nil {
//		FindCampaignByIdResponse = mockedFindCampaignByIdResponse.(*fs.FindCampaignByIdResponse)
//	}
//	return FindCampaignByIdResponse
//}
//
//func (rc *HookFindCampaignByIdMock) OnData(ctx context.Context, request *fs.FindCampaignByIdRequest, response *fs.FindCampaignByIdResponse) *fs.FindCampaignByIdResponse {
//	var FindCampaignByIdResponse *fs.FindCampaignByIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindCampaignByIdResponse := args.Get(0)
//	if mockedFindCampaignByIdResponse != nil {
//		FindCampaignByIdResponse = mockedFindCampaignByIdResponse.(*fs.FindCampaignByIdResponse)
//	}
//	return FindCampaignByIdResponse
//}
//func (rc *HookFindCampaignTemplateByIdMock) OnRequest(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) *fs.FindCampaignTemplateByIdResponse {
//	var FindCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse
//	args := rc.Called(ctx, request)
//	mockedFindCampaignTemplateByIdResponse := args.Get(0)
//	if mockedFindCampaignTemplateByIdResponse != nil {
//		FindCampaignTemplateByIdResponse = mockedFindCampaignTemplateByIdResponse.(*fs.FindCampaignTemplateByIdResponse)
//	}
//	return FindCampaignTemplateByIdResponse
//}
//
//func (rc *HookFindCampaignTemplateByIdMock) OnResponse(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
//	var FindCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindCampaignTemplateByIdResponse := args.Get(0)
//	if mockedFindCampaignTemplateByIdResponse != nil {
//		FindCampaignTemplateByIdResponse = mockedFindCampaignTemplateByIdResponse.(*fs.FindCampaignTemplateByIdResponse)
//	}
//	return FindCampaignTemplateByIdResponse
//}
//
//func (rc *HookFindCampaignTemplateByIdMock) OnError(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse, err error) *fs.FindCampaignTemplateByIdResponse {
//	var FindCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedFindCampaignTemplateByIdResponse := args.Get(0)
//	if mockedFindCampaignTemplateByIdResponse != nil {
//		FindCampaignTemplateByIdResponse = mockedFindCampaignTemplateByIdResponse.(*fs.FindCampaignTemplateByIdResponse)
//	}
//	return FindCampaignTemplateByIdResponse
//}
//
//func (rc *HookFindCampaignTemplateByIdMock) OnData(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest, response *fs.FindCampaignTemplateByIdResponse) *fs.FindCampaignTemplateByIdResponse {
//	var FindCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindCampaignTemplateByIdResponse := args.Get(0)
//	if mockedFindCampaignTemplateByIdResponse != nil {
//		FindCampaignTemplateByIdResponse = mockedFindCampaignTemplateByIdResponse.(*fs.FindCampaignTemplateByIdResponse)
//	}
//	return FindCampaignTemplateByIdResponse
//}
//func (rc *HookAddTargetUserMock) OnRequest(ctx context.Context, request *fs.AddTargetUserRequest) *fs.AddTargetUserResponse {
//	var AddTargetUserResponse *fs.AddTargetUserResponse
//	args := rc.Called(ctx, request)
//	mockedAddTargetUserResponse := args.Get(0)
//	if mockedAddTargetUserResponse != nil {
//		AddTargetUserResponse = mockedAddTargetUserResponse.(*fs.AddTargetUserResponse)
//	}
//	return AddTargetUserResponse
//}
//
//func (rc *HookAddTargetUserMock) OnResponse(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse) *fs.AddTargetUserResponse {
//	var AddTargetUserResponse *fs.AddTargetUserResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddTargetUserResponse := args.Get(0)
//	if mockedAddTargetUserResponse != nil {
//		AddTargetUserResponse = mockedAddTargetUserResponse.(*fs.AddTargetUserResponse)
//	}
//	return AddTargetUserResponse
//}
//
//func (rc *HookAddTargetUserMock) OnError(ctx context.Context, request *fs.AddTargetUserRequest, response *fs.AddTargetUserResponse, err error) *fs.AddTargetUserResponse {
//	var AddTargetUserResponse *fs.AddTargetUserResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddTargetUserResponse := args.Get(0)
//	if mockedAddTargetUserResponse != nil {
//		AddTargetUserResponse = mockedAddTargetUserResponse.(*fs.AddTargetUserResponse)
//	}
//	return AddTargetUserResponse
//}
//func (rc *HookAddTargetUserBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddTargetUserRequest) *fs.BulkAddTargetUserResponse {
//	var AddTargetUserBulkResponse *fs.BulkAddTargetUserResponse
//	args := rc.Called(ctx, request)
//	mockedAddTargetUserBulkResponse := args.Get(0)
//	if mockedAddTargetUserBulkResponse != nil {
//		AddTargetUserBulkResponse = mockedAddTargetUserBulkResponse.(*fs.BulkAddTargetUserResponse)
//	}
//	return AddTargetUserBulkResponse
//}
//
//func (rc *HookAddTargetUserBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse) *fs.BulkAddTargetUserResponse {
//	var AddTargetUserBulkResponse *fs.BulkAddTargetUserResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddTargetUserBulkResponse := args.Get(0)
//	if mockedAddTargetUserBulkResponse != nil {
//		AddTargetUserBulkResponse = mockedAddTargetUserBulkResponse.(*fs.BulkAddTargetUserResponse)
//	}
//	return AddTargetUserBulkResponse
//}
//
//func (rc *HookAddTargetUserBulkMock) OnError(ctx context.Context, request *fs.BulkAddTargetUserRequest, response *fs.BulkAddTargetUserResponse, err error) *fs.BulkAddTargetUserResponse {
//	var AddTargetUserBulkResponse *fs.BulkAddTargetUserResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddTargetUserBulkResponse := args.Get(0)
//	if mockedAddTargetUserBulkResponse != nil {
//		AddTargetUserBulkResponse = mockedAddTargetUserBulkResponse.(*fs.BulkAddTargetUserResponse)
//	}
//	return AddTargetUserBulkResponse
//}
//
//func (rc *HookAddInactionTargetUserMock) OnRequest(ctx context.Context, request *fs.AddInactionTargetUserRequest) *fs.AddInactionTargetUserResponse {
//	var AddInactionTargetUserResponse *fs.AddInactionTargetUserResponse
//	args := rc.Called(ctx, request)
//	mockedAddInactionTargetUserResponse := args.Get(0)
//	if mockedAddInactionTargetUserResponse != nil {
//		AddInactionTargetUserResponse = mockedAddInactionTargetUserResponse.(*fs.AddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserResponse
//}
//
//func (rc *HookAddInactionTargetUserMock) OnResponse(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse) *fs.AddInactionTargetUserResponse {
//	var AddInactionTargetUserResponse *fs.AddInactionTargetUserResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddInactionTargetUserResponse := args.Get(0)
//	if mockedAddInactionTargetUserResponse != nil {
//		AddInactionTargetUserResponse = mockedAddInactionTargetUserResponse.(*fs.AddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserResponse
//}
//
//func (rc *HookAddInactionTargetUserMock) OnError(ctx context.Context, request *fs.AddInactionTargetUserRequest, response *fs.AddInactionTargetUserResponse, err error) *fs.AddInactionTargetUserResponse {
//	var AddInactionTargetUserResponse *fs.AddInactionTargetUserResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddInactionTargetUserResponse := args.Get(0)
//	if mockedAddInactionTargetUserResponse != nil {
//		AddInactionTargetUserResponse = mockedAddInactionTargetUserResponse.(*fs.AddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserResponse
//}
//func (rc *HookAddInactionTargetUserBulkMock) OnRequest(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) *fs.BulkAddInactionTargetUserResponse {
//	var AddInactionTargetUserBulkResponse *fs.BulkAddInactionTargetUserResponse
//	args := rc.Called(ctx, request)
//	mockedAddInactionTargetUserBulkResponse := args.Get(0)
//	if mockedAddInactionTargetUserBulkResponse != nil {
//		AddInactionTargetUserBulkResponse = mockedAddInactionTargetUserBulkResponse.(*fs.BulkAddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserBulkResponse
//}
//
//func (rc *HookAddInactionTargetUserBulkMock) OnResponse(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse) *fs.BulkAddInactionTargetUserResponse {
//	var AddInactionTargetUserBulkResponse *fs.BulkAddInactionTargetUserResponse
//	args := rc.Called(ctx, request, response)
//	mockedAddInactionTargetUserBulkResponse := args.Get(0)
//	if mockedAddInactionTargetUserBulkResponse != nil {
//		AddInactionTargetUserBulkResponse = mockedAddInactionTargetUserBulkResponse.(*fs.BulkAddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserBulkResponse
//}
//
//func (rc *HookAddInactionTargetUserBulkMock) OnError(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest, response *fs.BulkAddInactionTargetUserResponse, err error) *fs.BulkAddInactionTargetUserResponse {
//	var AddInactionTargetUserBulkResponse *fs.BulkAddInactionTargetUserResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedAddInactionTargetUserBulkResponse := args.Get(0)
//	if mockedAddInactionTargetUserBulkResponse != nil {
//		AddInactionTargetUserBulkResponse = mockedAddInactionTargetUserBulkResponse.(*fs.BulkAddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserBulkResponse
//}
//
//
//func (rc *HookFindInactionTargetUserByCampaignIdMock) OnRequest(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) *fs.FindInactionTargetUserByCampaignIdResponse {
//	var FindInactionTargetUserByCampaignIdResponse *fs.FindInactionTargetUserByCampaignIdResponse
//	args := rc.Called(ctx, request)
//	mockedFindInactionTargetUserByCampaignIdResponse := args.Get(0)
//	if mockedFindInactionTargetUserByCampaignIdResponse != nil {
//		FindInactionTargetUserByCampaignIdResponse = mockedFindInactionTargetUserByCampaignIdResponse.(*fs.FindInactionTargetUserByCampaignIdResponse)
//	}
//	return FindInactionTargetUserByCampaignIdResponse
//}
//
//func (rc *HookFindInactionTargetUserByCampaignIdMock) OnResponse(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
//	var FindInactionTargetUserByCampaignIdResponse *fs.FindInactionTargetUserByCampaignIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindInactionTargetUserByCampaignIdResponse := args.Get(0)
//	if mockedFindInactionTargetUserByCampaignIdResponse != nil {
//		FindInactionTargetUserByCampaignIdResponse = mockedFindInactionTargetUserByCampaignIdResponse.(*fs.FindInactionTargetUserByCampaignIdResponse)
//	}
//	return FindInactionTargetUserByCampaignIdResponse
//}
//
//func (rc *HookFindInactionTargetUserByCampaignIdMock) OnError(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse, err error) *fs.FindInactionTargetUserByCampaignIdResponse {
//	var FindInactionTargetUserByCampaignIdResponse *fs.FindInactionTargetUserByCampaignIdResponse
//	args := rc.Called(ctx, request, response, err)
//	mockedFindInactionTargetUserByCampaignIdResponse := args.Get(0)
//	if mockedFindInactionTargetUserByCampaignIdResponse != nil {
//		FindInactionTargetUserByCampaignIdResponse = mockedFindInactionTargetUserByCampaignIdResponse.(*fs.FindInactionTargetUserByCampaignIdResponse)
//	}
//	return FindInactionTargetUserByCampaignIdResponse
//}
//
//func (rc *HookFindInactionTargetUserByCampaignIdMock) OnData(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest, response *fs.FindInactionTargetUserByCampaignIdResponse) *fs.FindInactionTargetUserByCampaignIdResponse {
//	var FindInactionTargetUserByCampaignIdResponse *fs.FindInactionTargetUserByCampaignIdResponse
//	args := rc.Called(ctx, request, response)
//	mockedFindInactionTargetUserByCampaignIdResponse := args.Get(0)
//	if mockedFindInactionTargetUserByCampaignIdResponse != nil {
//		FindInactionTargetUserByCampaignIdResponse = mockedFindInactionTargetUserByCampaignIdResponse.(*fs.FindInactionTargetUserByCampaignIdResponse)
//	}
//	return FindInactionTargetUserByCampaignIdResponse
//}
//
//
//func (se *ExecutorMock) ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) (*fs.FindInactionTargetUserByCampaignIdResponse, error) {
//	var FindInactionTargetUserByCampaignIdResponse *fs.FindInactionTargetUserByCampaignIdResponse
//	args := se.Called(ctx,request)
//	mockedFindInactionTargetUserByCampaignIdResponse := args.Get(0)
//	if mockedFindInactionTargetUserByCampaignIdResponse!=nil{
//		FindInactionTargetUserByCampaignIdResponse = mockedFindInactionTargetUserByCampaignIdResponse.(*fs.FindInactionTargetUserByCampaignIdResponse)
//	}
//	return FindInactionTargetUserByCampaignIdResponse,args.Error(1)
//}
//
//
//
//func (se *ExecutorMock) ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error) {
//    var AddCampaignResponse *fs.AddCampaignResponse
//    args := se.Called(ctx,request)
//    mockedAddCampaignResponse := args.Get(0)
//    if mockedAddCampaignResponse!=nil{
//		AddCampaignResponse = mockedAddCampaignResponse.(*fs.AddCampaignResponse)
//	}
//	return AddCampaignResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddCampaignBulk(ctx context.Context, request *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error) {
//    var AddCampaignResponse *fs.BulkAddCampaignResponse
//    args := se.Called(ctx,request)
//    mockedAddCampaignResponse := args.Get(0)
//    if mockedAddCampaignResponse!=nil{
//		AddCampaignResponse = mockedAddCampaignResponse.(*fs.BulkAddCampaignResponse)
//	}
//	return AddCampaignResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest) (*fs.UpdateCampaignResponse, error) {
//    var UpdateCampaignResponse *fs.UpdateCampaignResponse
//    args := se.Called(ctx,request)
//    mockedUpdateCampaignResponse := args.Get(0)
//    if mockedUpdateCampaignResponse!=nil{
//		UpdateCampaignResponse = mockedUpdateCampaignResponse.(*fs.UpdateCampaignResponse)
//	}
//	return UpdateCampaignResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error) {
//    var AddCampaignTemplateResponse *fs.AddCampaignTemplateResponse
//    args := se.Called(ctx,request)
//    mockedAddCampaignTemplateResponse := args.Get(0)
//    if mockedAddCampaignTemplateResponse!=nil{
//		AddCampaignTemplateResponse = mockedAddCampaignTemplateResponse.(*fs.AddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddCampaignTemplateBulk(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error) {
//    var AddCampaignTemplateResponse *fs.BulkAddCampaignTemplateResponse
//    args := se.Called(ctx,request)
//    mockedAddCampaignTemplateResponse := args.Get(0)
//    if mockedAddCampaignTemplateResponse!=nil{
//		AddCampaignTemplateResponse = mockedAddCampaignTemplateResponse.(*fs.BulkAddCampaignTemplateResponse)
//	}
//	return AddCampaignTemplateResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddNewCampaign(ctx context.Context, request *fs.AddNewCampaignRequest) (*fs.AddNewCampaignResponse, error) {
//    var AddNewCampaignResponse *fs.AddNewCampaignResponse
//    args := se.Called(ctx,request)
//    mockedAddNewCampaignResponse := args.Get(0)
//    if mockedAddNewCampaignResponse!=nil{
//		AddNewCampaignResponse = mockedAddNewCampaignResponse.(*fs.AddNewCampaignResponse)
//	}
//	return AddNewCampaignResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddNewCampaignBulk(ctx context.Context, request *fs.BulkAddNewCampaignRequest) (*fs.BulkAddNewCampaignResponse, error) {
//    var AddNewCampaignResponse *fs.BulkAddNewCampaignResponse
//    args := se.Called(ctx,request)
//    mockedAddNewCampaignResponse := args.Get(0)
//    if mockedAddNewCampaignResponse!=nil{
//		AddNewCampaignResponse = mockedAddNewCampaignResponse.(*fs.BulkAddNewCampaignResponse)
//	}
//	return AddNewCampaignResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteGetUserList(ctx context.Context, request *fs.GetUserListRequest) (*fs.GetUserListResponse, error) {
//    var GetUserListResponse *fs.GetUserListResponse
//    args := se.Called(ctx,request)
//    mockedGetUserListResponse := args.Get(0)
//    if mockedGetUserListResponse!=nil{
//		GetUserListResponse = mockedGetUserListResponse.(*fs.GetUserListResponse)
//	}
//	return GetUserListResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error) {
//    var FindCampaignByIdResponse *fs.FindCampaignByIdResponse
//    args := se.Called(ctx,request)
//    mockedFindCampaignByIdResponse := args.Get(0)
//    if mockedFindCampaignByIdResponse!=nil{
//		FindCampaignByIdResponse = mockedFindCampaignByIdResponse.(*fs.FindCampaignByIdResponse)
//	}
//	return FindCampaignByIdResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error) {
//    var FindCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse
//    args := se.Called(ctx,request)
//    mockedFindCampaignTemplateByIdResponse := args.Get(0)
//    if mockedFindCampaignTemplateByIdResponse!=nil{
//		FindCampaignTemplateByIdResponse = mockedFindCampaignTemplateByIdResponse.(*fs.FindCampaignTemplateByIdResponse)
//	}
//	return FindCampaignTemplateByIdResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error) {
//    var AddTargetUserResponse *fs.AddTargetUserResponse
//    args := se.Called(ctx,request)
//    mockedAddTargetUserResponse := args.Get(0)
//    if mockedAddTargetUserResponse!=nil{
//		AddTargetUserResponse = mockedAddTargetUserResponse.(*fs.AddTargetUserResponse)
//	}
//	return AddTargetUserResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddTargetUserBulk(ctx context.Context, request *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error) {
//    var AddTargetUserResponse *fs.BulkAddTargetUserResponse
//    args := se.Called(ctx,request)
//    mockedAddTargetUserResponse := args.Get(0)
//    if mockedAddTargetUserResponse!=nil{
//		AddTargetUserResponse = mockedAddTargetUserResponse.(*fs.BulkAddTargetUserResponse)
//	}
//	return AddTargetUserResponse,args.Error(1)
//}
//
//func (se *ExecutorMock) ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error) {
//	var AddInactionTargetUserResponse *fs.AddInactionTargetUserResponse
//	args := se.Called(ctx,request)
//	mockedAddInactionTargetUserResponse := args.Get(0)
//	if mockedAddInactionTargetUserResponse!=nil{
//		AddInactionTargetUserResponse = mockedAddInactionTargetUserResponse.(*fs.AddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserResponse,args.Error(1)
//}
//func (se *ExecutorMock) ExecuteAddInactionTargetUserBulk(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error) {
//	var AddInactionTargetUserResponse *fs.BulkAddInactionTargetUserResponse
//	args := se.Called(ctx,request)
//	mockedAddInactionTargetUserResponse := args.Get(0)
//	if mockedAddInactionTargetUserResponse!=nil{
//		AddInactionTargetUserResponse = mockedAddInactionTargetUserResponse.(*fs.BulkAddInactionTargetUserResponse)
//	}
//	return AddInactionTargetUserResponse,args.Error(1)
//}
//
//
//
//
//func TestExecuteAddCampaignBulk(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddCampaignBulkMock{}
//	hook.BulkAddCampaignExecutor = &hook.GenericAddCampaignExecutorBulk{
//		AddCampaignBulkInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.BulkAddCampaignResponse{
//		Status: Status,
//	}
//	request := &fs.BulkAddCampaignRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddCampaignBulk", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddCampaignBulk(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddCampaignBulk", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.BulkAddCampaignResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddCampaignBulk(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddCampaign(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddCampaignMock{}
//	hook.AddCampaignExecutor = &hook.GenericAddCampaignExecutor{
//		AddCampaignInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.AddCampaignResponse{
//		Status: Status,
//	}
//	request := &fs.AddCampaignRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddCampaign", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddCampaign(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddCampaign", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.AddCampaignResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddCampaign(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteUpdateCampaign(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookUpdateCampaignMock{}
//	hook.UpdateCampaignExecutor = &hook.GenericUpdateCampaignExecutor{
//		UpdateCampaignInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.UpdateCampaignResponse{
//		Status: Status,
//	}
//	request := &fs.UpdateCampaignRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteUpdateCampaign", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteUpdateCampaign(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteUpdateCampaign", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.UpdateCampaignResponse)(nil), err).Return(nil)
//	response = service.ExecuteUpdateCampaign(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddCampaignTemplateBulk(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddCampaignTemplateBulkMock{}
//	hook.BulkAddCampaignTemplateExecutor = &hook.GenericAddCampaignTemplateExecutorBulk{
//		AddCampaignTemplateBulkInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.BulkAddCampaignTemplateResponse{
//		Status: Status,
//	}
//	request := &fs.BulkAddCampaignTemplateRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddCampaignTemplateBulk", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddCampaignTemplateBulk(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddCampaignTemplateBulk", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.BulkAddCampaignTemplateResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddCampaignTemplateBulk(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddCampaignTemplate(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddCampaignTemplateMock{}
//	hook.AddCampaignTemplateExecutor = &hook.GenericAddCampaignTemplateExecutor{
//		AddCampaignTemplateInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.AddCampaignTemplateResponse{
//		Status: Status,
//	}
//	request := &fs.AddCampaignTemplateRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddCampaignTemplate", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddCampaignTemplate(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddCampaignTemplate", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.AddCampaignTemplateResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddCampaignTemplate(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteFindCampaignById(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookFindCampaignByIdMock{}
//	hook.FindCampaignByIdExecutor = &hook.GenericFindCampaignByIdExecutor{
//		FindCampaignByIdInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.FindCampaignByIdResponse{
//		Status: Status,
//	}
//	request := &fs.FindCampaignByIdRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteFindCampaignById", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//	hookMock.On("OnData", ctx, request, mockedResponse).Return(nil)
//	response := service.ExecuteFindCampaignById(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteFindCampaignById", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.FindCampaignByIdResponse)(nil), err).Return(nil)
//	response = service.ExecuteFindCampaignById(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteFindCampaignTemplateById(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookFindCampaignTemplateByIdMock{}
//	hook.FindCampaignTemplateByIdExecutor = &hook.GenericFindCampaignTemplateByIdExecutor{
//		FindCampaignTemplateByIdInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.FindCampaignTemplateByIdResponse{
//		Status: Status,
//	}
//	request := &fs.FindCampaignTemplateByIdRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteFindCampaignTemplateById", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//	hookMock.On("OnData", ctx, request, mockedResponse).Return(nil)
//	response := service.ExecuteFindCampaignTemplateById(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteFindCampaignTemplateById", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.FindCampaignTemplateByIdResponse)(nil), err).Return(nil)
//	response = service.ExecuteFindCampaignTemplateById(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddTargetUserBulk(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddTargetUserBulkMock{}
//	hook.BulkAddTargetUserExecutor = &hook.GenericAddTargetUserExecutorBulk{
//		AddTargetUserBulkInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.BulkAddTargetUserResponse{
//		Status: Status,
//	}
//	request := &fs.BulkAddTargetUserRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddTargetUserBulk", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddTargetUserBulk(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddTargetUserBulk", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.BulkAddTargetUserResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddTargetUserBulk(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddTargetUser(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddTargetUserMock{}
//	hook.AddTargetUserExecutor = &hook.GenericAddTargetUserExecutor{
//		AddTargetUserInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//        Status: fs.StatusCode_SUCCESS,
//    }
//
//	mockedResponse :=&fs.AddTargetUserResponse{
//		Status: Status,
//	}
//	request := &fs.AddTargetUserRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddTargetUser", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddTargetUser(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//        Status: fs.StatusCode_DB_FAILURE,
//    }
//	executorMock.On("ExecuteAddTargetUser", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.AddTargetUserResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddTargetUser(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//
//
//func TestExecuteAddInactionTargetUserBulk(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddInactionTargetUserBulkMock{}
//	hook.BulkAddInactionTargetUserExecutor = &hook.GenericAddInactionTargetUserExecutorBulk{
//		AddInactionTargetUserBulkInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//		Status: fs.StatusCode_SUCCESS,
//	}
//
//	mockedResponse :=&fs.BulkAddInactionTargetUserResponse{
//		Status: Status,
//	}
//	request := &fs.BulkAddInactionTargetUserRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddInactionTargetUserBulk", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddInactionTargetUserBulk(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//		Status: fs.StatusCode_DB_FAILURE,
//	}
//	executorMock.On("ExecuteAddInactionTargetUserBulk", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.BulkAddInactionTargetUserResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddInactionTargetUserBulk(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//func TestExecuteAddInactionTargetUser(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookAddInactionTargetUserMock{}
//	hook.AddInactionTargetUserExecutor = &hook.GenericAddInactionTargetUserExecutor{
//		AddInactionTargetUserInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//		Status: fs.StatusCode_SUCCESS,
//	}
//
//	mockedResponse :=&fs.AddInactionTargetUserResponse{
//		Status: Status,
//	}
//	request := &fs.AddInactionTargetUserRequest{}
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteAddInactionTargetUser", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//
//	response := service.ExecuteAddInactionTargetUser(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//		Status: fs.StatusCode_DB_FAILURE,
//	}
//	executorMock.On("ExecuteAddInactionTargetUser", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.AddInactionTargetUserResponse)(nil), err).Return(nil)
//	response = service.ExecuteAddInactionTargetUser(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
//
//
//
//func TestExecuteFindInactionTargetUserByCampaignId(t *testing.T) {
//	executorMock := &ExecutorMock{}
//	executor.RequestExecutor = &executor.GenericExecutor{
//		ServiceExecutor: executorMock,
//	}
//	metricsMock := &MetricsMock{}
//	metrics.Metrics = metricsMock
//	hookMock := &HookFindInactionTargetUserByCampaignIdMock{}
//	hook.FindInactionTargetUserByCampaignIdExecutor = &hook.GenericFindInactionTargetUserByCampaignIdExecutor{
//		FindInactionTargetUserByCampaignIdInterface: hookMock,
//	}
//	ctx := context.Background()
//
//	Status :=  &fs.Status{
//		Status: fs.StatusCode_SUCCESS,
//	}
//
//	mockedResponse :=&fs.FindInactionTargetUserByCampaignIdResponse{
//		Status: Status,
//	}
//	request := &fs.FindInactionTargetUserByCampaignIdRequest{}
//
//
//	metricsMock.On("PushToSummarytMetrics").Return()
//	metricsMock.On("IncrementCounterMetrics").Return()
//	executorMock.On("ExecuteFindInactionTargetUserByCampaignId", ctx, request).Return(mockedResponse,nil).Once()
//	hookMock.On("OnRequest", ctx, request).Return(nil)
//	hookMock.On("OnResponse", ctx, request, mockedResponse).Return(nil)
//	hookMock.On("OnError", ctx, request, mockedResponse, nil).Return(nil)
//	hookMock.On("OnData", ctx, request, mockedResponse).Return(nil)
//	response := service.ExecuteFindInactionTargetUserByCampaignId(ctx,request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, response.Status.Status)
//
//
//	err := errors.New("Some Error")
//	mockedResponse.Status = &fs.Status{
//		Status: fs.StatusCode_DB_FAILURE,
//	}
//	executorMock.On("ExecuteFindInactionTargetUserByCampaignId", ctx, request).Return(nil,err).Once()
//	hookMock.On("OnError", ctx, request, (*fs.FindInactionTargetUserByCampaignIdResponse)(nil), err).Return(nil)
//	response = service.ExecuteFindInactionTargetUserByCampaignId(ctx,request)
//	assert.Equal(fs.StatusCode_DB_FAILURE, response.Status.Status)
//}
