package executor_test

//
//import (
//	"context"
//	"database/sql"
//	"errors"
//	"testing"
//
//	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
//	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/executor"
//	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/mappers"
//	"github.com/DATA-DOG/go-sqlmock"
//	entsql "github.com/facebook/ent/dialect/sql"
//	"github.com/stretchr/testify/assert"
//)
//
//var Mock sqlmock.Sqlmock
//
//func init() {
//	var db *sql.DB
//	var err error
//	db, Mock, err = sqlmock.New()
//	if err != nil {
//		panic(err)
//	}
//	executor.Driver.Driver = entsql.OpenDB("mysql", db)
//}
//
//func TestExecuteAddCampaign(t *testing.T) {
//   request := &fs.AddCampaignRequest{}
//   model := mappers.MakeAddCampaignRequestVO(request)
//	args := executor.AddCampaignArgs(model)
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9],args[10]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddCampaign(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9],args[10]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddCampaign(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteAddCampaignBulk(t *testing.T) {
//   request := &fs.BulkAddCampaignRequest{}
//
//   req := &fs.AddCampaignRequest{}
//	request.Requests = append(request.Requests, req)
//	request.Requests = append(request.Requests, req)
//	var args []interface{}
//	for _, r := range request.Requests {
//		model := mappers.MakeAddCampaignRequestVO(r)
//		args = append(args, executor.AddCampaignArgs(model)...)
//	}
//
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9],args[10],args[11],args[12],args[13],args[14],args[15],args[16],args[17],args[18],args[19],args[20],args[21]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddCampaignBulk(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9],args[10],args[11],args[12],args[13],args[14],args[15],args[16],args[17],args[18],args[19],args[20],args[21]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddCampaignBulk(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteUpdateCampaign(t *testing.T) {
//   request := &fs.UpdateCampaignRequest{}
//   model := mappers.MakeUpdateCampaignRequestVO(request)
//	args := executor.UpdateCampaignArgs(model)
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteUpdateCampaign(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteUpdateCampaign(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteAddCampaignTemplate(t *testing.T) {
//   request := &fs.AddCampaignTemplateRequest{}
//   model := mappers.MakeAddCampaignTemplateRequestVO(request)
//	args := executor.AddCampaignTemplateArgs(model)
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddCampaignTemplate(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3],args[4]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddCampaignTemplate(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteAddCampaignTemplateBulk(t *testing.T) {
//   request := &fs.BulkAddCampaignTemplateRequest{}
//
//   req := &fs.AddCampaignTemplateRequest{}
//	request.Requests = append(request.Requests, req)
//	request.Requests = append(request.Requests, req)
//	var args []interface{}
//	for _, r := range request.Requests {
//		model := mappers.MakeAddCampaignTemplateRequestVO(r)
//		args = append(args, executor.AddCampaignTemplateArgs(model)...)
//	}
//
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddCampaignTemplateBulk(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddCampaignTemplateBulk(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteFindCampaignById(t *testing.T) {
//   request := &fs.FindCampaignByIdRequest{}
//   rows := sqlmock.NewRows([]string{" "," "," "," "," "," "," "," "," "," "," "," "," "," "," "," "," "}).
//   AddRow(nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
//   ctx := context.Background()
//
//   Mock.ExpectQuery("select").WillReturnRows(rows)
//   resp, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectQuery("select ").WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteFindCampaignById(ctx, request)
//	assert.Equal(er, err)
//}
//
//func TestExecuteFindCampaignTemplateById(t *testing.T) {
//   request := &fs.FindCampaignTemplateByIdRequest{}
//   rows := sqlmock.NewRows([]string{" "," "," "," "," "}).
//   AddRow(nil,nil,nil,nil,nil)
//   ctx := context.Background()
//
//   Mock.ExpectQuery("select").WillReturnRows(rows)
//   resp, err := executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectQuery("select ").WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, request)
//	assert.Equal(er, err)
//}
//
//func TestExecuteAddTargetUser(t *testing.T) {
//   request := &fs.AddTargetUserRequest{}
//   model := mappers.MakeAddTargetUserRequestVO(request)
//	args := executor.AddTargetUserArgs(model)
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddTargetUser(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddTargetUser(ctx, request)
//   assert.Equal(er, err)
//}
//
//func TestExecuteAddTargetUserBulk(t *testing.T) {
//   request := &fs.BulkAddTargetUserRequest{}
//
//   req := &fs.AddTargetUserRequest{}
//	request.Requests = append(request.Requests, req)
//	request.Requests = append(request.Requests, req)
//	var args []interface{}
//	for _, r := range request.Requests {
//		model := mappers.MakeAddTargetUserRequestVO(r)
//		args = append(args, executor.AddTargetUserArgs(model)...)
//	}
//
//   ctx := context.Background()
//
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5]).WillReturnResult(sqlmock.NewResult(1, 1))
//   resp, err := executor.RequestExecutor.ExecuteAddTargetUserBulk(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//   er := errors.New("Some Error")
//   Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5]).WillReturnError(er)
//   resp, err = executor.RequestExecutor.ExecuteAddTargetUserBulk(ctx, request)
//   assert.Equal(er, err)
//}
//
//
//func TestExecuteAddInactionTargetUser(t *testing.T) {
//	request := &fs.AddInactionTargetUserRequest{}
//	model := mappers.MakeAddInactionTargetUserRequestVO(request)
//	args := executor.AddInactionTargetUserArgs(model)
//	ctx := context.Background()
//
//	Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2]).WillReturnResult(sqlmock.NewResult(1, 1))
//	resp, err := executor.RequestExecutor.ExecuteAddInactionTargetUser(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//	er := errors.New("Some Error")
//	Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2]).WillReturnError(er)
//	resp, err = executor.RequestExecutor.ExecuteAddInactionTargetUser(ctx, request)
//	assert.Equal(er, err)
//}
//
//func TestExecuteAddInactionTargetUserBulk(t *testing.T) {
//	request := &fs.BulkAddInactionTargetUserRequest{}
//
//	req := &fs.AddInactionTargetUserRequest{}
//	request.Requests = append(request.Requests, req)
//	request.Requests = append(request.Requests, req)
//	var args []interface{}
//	for _, r := range request.Requests {
//		model := mappers.MakeAddInactionTargetUserRequestVO(r)
//		args = append(args, executor.AddInactionTargetUserArgs(model)...)
//	}
//
//	ctx := context.Background()
//
//	Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5]).WillReturnResult(sqlmock.NewResult(1, 1))
//	resp, err := executor.RequestExecutor.ExecuteAddInactionTargetUserBulk(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//	er := errors.New("Some Error")
//	Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5]).WillReturnError(er)
//	resp, err = executor.RequestExecutor.ExecuteAddInactionTargetUserBulk(ctx, request)
//	assert.Equal(er, err)
//}
//
//
//func TestExecuteFindInactionTargetUserByCampaignId(t *testing.T) {
//	request := &fs.FindInactionTargetUserByCampaignIdRequest{}
//	rows := sqlmock.NewRows([]string{" "," "," "," "}).
//		AddRow(nil,nil,nil,nil)
//	ctx := context.Background()
//
//	Mock.ExpectQuery("select").WillReturnRows(rows)
//	resp, err := executor.RequestExecutor.ExecuteFindInactionTargetUserByCampaignId(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//	er := errors.New("Some Error")
//	Mock.ExpectQuery("select ").WillReturnError(er)
//	resp, err = executor.RequestExecutor.ExecuteFindInactionTargetUserByCampaignId(ctx, request)
//	assert.Equal(er, err)
//}
//
//
//func TestExecuteAddDynamicData(t *testing.T) {
//	request := &fs.AddDynamicDataRequest{}
//	model := mappers.MakeAddDynamicDataRequestVO(request)
//	args := executor.AddDynamicDataArgs(model)
//	ctx := context.Background()
//
//	Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3]).WillReturnResult(sqlmock.NewResult(1, 1))
//	resp, err := executor.RequestExecutor.ExecuteAddDynamicData(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//	er := errors.New("Some Error")
//	Mock.ExpectExec("insert into").WithArgs(args[0],args[1],args[2],args[3]).WillReturnError(er)
//	resp, err = executor.RequestExecutor.ExecuteAddDynamicData(ctx, request)
//	assert.Equal(er, err)
//}
//
//func TestExecuteAddDynamicDataBulk(t *testing.T) {
//	request := &fs.BulkAddDynamicDataRequest{}
//
//	req := &fs.AddDynamicDataRequest{}
//	request.Requests = append(request.Requests, req)
//	request.Requests = append(request.Requests, req)
//	var args []interface{}
//	for _, r := range request.Requests {
//		model := mappers.MakeAddDynamicDataRequestVO(r)
//		args = append(args, executor.AddDynamicDataArgs(model)...)
//	}
//
//	ctx := context.Background()
//
//	Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7]).WillReturnResult(sqlmock.NewResult(1, 1))
//	resp, err := executor.RequestExecutor.ExecuteAddDynamicDataBulk(ctx, request)
//	assert := assert.New(t)
//	assert.Equal(fs.StatusCode_SUCCESS, resp.Status.Status)
//
//	er := errors.New("Some Error")
//	Mock.ExpectExec("insert into ").WithArgs(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7]).WillReturnError(er)
//	resp, err = executor.RequestExecutor.ExecuteAddDynamicDataBulk(ctx, request)
//	assert.Equal(er, err)
//}
//
