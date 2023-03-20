package executor

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"code.nurture.farm/platform/CampaignService/core/golang/database"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/models"
	"context"
	"database/sql"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"strings"
)

type ServiceExecutor interface {
	ExecuteAddCampaignBulk(ctx context.Context, bulkrequest *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error)
	ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error)
	ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest, tx ...dialect.Tx) (*fs.UpdateCampaignResponse, error)
	ExecuteAddControlGroup(ctx context.Context, request *fs.AddControlGroupRequest) (*fs.AddControlGroupResponse, error)
	ExecuteAddCampaignTemplateBulk(ctx context.Context, bulkrequest *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error)
	ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error)
	ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error)
	ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error)
	ExecuteAddTargetUserBulk(ctx context.Context, bulkrequest *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error)
	ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error)
	ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) (*fs.FindTargetUserByIdResponse, error)
	ExecuteAddInactionTargetUserBulk(ctx context.Context, bulkrequest *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error)
	ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error)
	ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) (*fs.FindInactionTargetUserByCampaignIdResponse, error)
	ExecuteFilterCampaigns(ctx context.Context, request *fs.FilterCampaignRequest) (*fs.FilterCampaignResponse, error)
	ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, tx ...dialect.Tx) (*fs.GetDynamicDataByKeyResponse, error)
	ExecuteAddDynamicDataBulk(ctx context.Context, bulkrequest *fs.BulkAddDynamicDataRequest, tx ...dialect.Tx) (*fs.BulkAddDynamicDataResponse, error)
	ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest, tx ...dialect.Tx) (*fs.AddDynamicDataResponse, error)
	ExecuteFindQueryCampaign(ctx context.Context, request *fs.FindQueryCampaignRequest) (*fs.FindQueryCampaignResponse, error)
	ExecuteAddQueryCampaignBulk(ctx context.Context, bulkrequest *fs.BulkAddQueryCampaignRequest) (*fs.BulkAddQueryCampaignResponse, error)
	ExecuteAddQueryCampaign(ctx context.Context, request *fs.AddQueryCampaignRequest) (*fs.AddQueryCampaignResponse, error)
	ExecuteFindControlGroupByCampaignId(ctx context.Context, request *fs.FindControlGroupByCampaignIdRequest) (*fs.FindControlGroupByCampaignIdResponse, error)
}

type GenericExecutor struct {
	ServiceExecutor ServiceExecutor
}

type Executor struct {
}

var RequestExecutor *GenericExecutor

func (se *GenericExecutor) ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, tx ...dialect.Tx) (*fs.GetDynamicDataByKeyResponse, error) {
	return se.ServiceExecutor.ExecuteGetDynamicDataByKey(ctx, request, tx...)
}

func (se *GenericExecutor) ExecuteAddCampaignBulk(ctx context.Context, bulkrequest *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteAddCampaignBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteAddCampaign(ctx, request)
}

func (se *GenericExecutor) ExecuteAddControlGroup(ctx context.Context, request *fs.AddControlGroupRequest) (*fs.AddControlGroupResponse, error) {
	return se.ServiceExecutor.ExecuteAddControlGroup(ctx, request)
}

func (se *GenericExecutor) ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest, tx ...dialect.Tx) (*fs.UpdateCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteUpdateCampaign(ctx, request, tx...)
}

func (se *GenericExecutor) ExecuteAddCampaignTemplateBulk(ctx context.Context, bulkrequest *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error) {
	return se.ServiceExecutor.ExecuteAddCampaignTemplateBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error) {
	return se.ServiceExecutor.ExecuteAddCampaignTemplate(ctx, request)
}

func (se *GenericExecutor) ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error) {
	return se.ServiceExecutor.ExecuteFindCampaignById(ctx, request)
}

func (se *GenericExecutor) ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error) {
	return se.ServiceExecutor.ExecuteFindCampaignTemplateById(ctx, request)
}

func (se *GenericExecutor) ExecuteAddTargetUserBulk(ctx context.Context, bulkrequest *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error) {
	return se.ServiceExecutor.ExecuteAddTargetUserBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error) {
	return se.ServiceExecutor.ExecuteAddTargetUser(ctx, request)
}

func (se *GenericExecutor) ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) (*fs.FindTargetUserByIdResponse, error) {
	return se.ServiceExecutor.ExecuteFindTargetUserById(ctx, request)
}

func (se *GenericExecutor) ExecuteAddInactionTargetUserBulk(ctx context.Context, bulkrequest *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error) {
	return se.ServiceExecutor.ExecuteAddInactionTargetUserBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error) {
	return se.ServiceExecutor.ExecuteAddInactionTargetUser(ctx, request)
}

func (se *GenericExecutor) ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) (*fs.FindInactionTargetUserByCampaignIdResponse, error) {
	return se.ServiceExecutor.ExecuteFindInactionTargetUserByCampaignId(ctx, request)
}

func (se *GenericExecutor) ExecuteFilterCampaigns(ctx context.Context, request *fs.FilterCampaignRequest) (*fs.FilterCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteFilterCampaigns(ctx, request)
}

func (se *GenericExecutor) ExecuteAddDynamicDataBulk(ctx context.Context, bulkrequest *fs.BulkAddDynamicDataRequest, tx ...dialect.Tx) (*fs.BulkAddDynamicDataResponse, error) {
	return se.ServiceExecutor.ExecuteAddDynamicDataBulk(ctx, bulkrequest, tx...)
}

func (se *GenericExecutor) ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest, tx ...dialect.Tx) (*fs.AddDynamicDataResponse, error) {
	return se.ServiceExecutor.ExecuteAddDynamicData(ctx, request, tx...)
}

func (se *GenericExecutor) ExecuteFindQueryCampaign(ctx context.Context, request *fs.FindQueryCampaignRequest) (*fs.FindQueryCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteFindQueryCampaign(ctx, request)
}

func (se *GenericExecutor) ExecuteAddQueryCampaignBulk(ctx context.Context, bulkrequest *fs.BulkAddQueryCampaignRequest) (*fs.BulkAddQueryCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteAddQueryCampaignBulk(ctx, bulkrequest)
}

func (se *GenericExecutor) ExecuteAddQueryCampaign(ctx context.Context, request *fs.AddQueryCampaignRequest) (*fs.AddQueryCampaignResponse, error) {
	return se.ServiceExecutor.ExecuteAddQueryCampaign(ctx, request)
}

func (se *GenericExecutor) ExecuteFindControlGroupByCampaignId(ctx context.Context, request *fs.FindControlGroupByCampaignIdRequest) (*fs.FindControlGroupByCampaignIdResponse, error) {
	return se.ServiceExecutor.ExecuteFindControlGroupByCampaignId(ctx, request)
}

func init() {
	RequestExecutor = &GenericExecutor{
		ServiceExecutor: &Executor{},
	}
}

func (se *Executor) ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error) {

	model := mappers.MakeAddCampaignRequestVO(request, "", "", nil)
	args := AddCampaignArgs(model)

	var rows sql.Result
	query := query.QUERY_AddCampaign

	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddCampaignRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddCampaignRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddControlGroup(ctx context.Context, request *fs.AddControlGroupRequest) (*fs.AddControlGroupResponse, error) {

	model := mappers.MakeAddControlGroupRequestVO(int(request.CampaignId), request.Attributes, request.BloomFilter)
	args := AddControlGroupArgs(model)

	var rows sql.Result
	query := query.QUERY_AddControlGroup

	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddControlGroupRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddControlGroupRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddControlGroupResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}
	return response, nil
}

func (se *Executor) ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest, tx ...dialect.Tx) (*fs.GetDynamicDataByKeyResponse, error) {

	response := &fs.GetDynamicDataByKeyResponse{}
	var rows = entsql.Rows{}
	args := GetDynamicDataByKeyArgs(request)
	query := query.QUERY_GetDynamicDataByKey

	var err error
	if tx != nil {
		err = tx[0].Query(ctx, query, args, &rows)
	} else {
		err = Driver.GetDriver().Query(ctx, query, args, &rows)
	}
	if err != nil {
		logger.Error("Error could not ExecuteGetDynamicDataByKeyRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.GetDynamicDataByKeyResponseVO{}
		err := rows.Scan(&model.CampaignId, &model.DynamicKey, &model.CtaLink, &model.Media)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteGetDynamicDataByKeyRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeGetDynamicDataByKeyResponseVO(&model))
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteAddCampaignBulk(ctx context.Context, bulkRequest *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error) {

	var args []interface{}
	query := query.QUERY_AddCampaign
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			query += "(?,?,?,?,?,?,?,?,?,?,?)"
		} else {
			query += "(?,?,?,?,?,?,?,?,?,?,?),"
		}
		model := mappers.MakeAddCampaignRequestVO(request, "", "", nil)
		args = append(args, AddCampaignArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddCampaignRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddCampaignRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddCampaignResponse
	for index, _ := range bulkRequest.Requests {
		responses = append(responses, &fs.AddCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},

			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}

func (se *Executor) ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest, tx ...dialect.Tx) (*fs.UpdateCampaignResponse, error) {

	model := mappers.MakeUpdateCampaignRequestVO(request)
	args := UpdateCampaignArgs(model)

	var rows sql.Result
	query := query.QUERY_UpdateCampaign

	var err error
	if tx != nil {
		err = tx[0].Exec(ctx, query, args, &rows)
	} else {
		err = Driver.GetDriver().Exec(ctx, query, args, &rows)
	}
	if err != nil {
		logger.Error("Error could not ExecuteUpdateCampaignRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for UpdateCampaignRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.UpdateCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error) {

	model := mappers.MakeAddCampaignTemplateRequestVO(request)
	args := AddCampaignTemplateArgs(model)

	var rows sql.Result
	query := query.QUERY_AddCampaignTemplate

	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddCampaignTemplateRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddCampaignTemplateRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddCampaignTemplateResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddCampaignTemplateBulk(ctx context.Context, bulkRequest *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error) {

	var args []interface{}
	query := query.QUERY_AddCampaignTemplate
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			query += "(?,?,?,?,?)"
		} else {
			query += "(?,?,?,?,?),"
		}
		model := mappers.MakeAddCampaignTemplateRequestVO(request)
		args = append(args, AddCampaignTemplateArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddCampaignTemplateRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddCampaignTemplateRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddCampaignTemplateResponse
	for index, _ := range bulkRequest.Requests {
		responses = append(responses, &fs.AddCampaignTemplateResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},

			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddCampaignTemplateResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}

func (se *Executor) ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error) {

	response := &fs.FindCampaignByIdResponse{}
	var rows = entsql.Rows{}
	args := FindCampaignByIdArgs(request)
	query := query.QUERY_FindCampaignById

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindCampaignByIdRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindCampaignByIdResponseVO{}
		err := rows.Scan(&model.Id, &model.Namespace, &model.Name, &model.Description, &model.CronExpression, &model.Occurrences, &model.CommunicationChannel, &model.Status, &model.Type, &model.ScheduleType, &model.Query, &model.InactionQuery, &model.InactionDuration, &model.Attributes, &model.CreatedByActorid, &model.CreatedByActortype, &model.UpdatedByActorid, &model.UpdatedByActortype, &model.Version, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindCampaignByIdRequest", zap.Error(err))
			return nil, err
		}

		response.Records = mappers.MakeFindCampaignByIdResponseVO(&model)
		templateRequest := &fs.FindCampaignTemplateByIdRequest{
			CampaignId: response.Records.Id,
		}
		templateResponse, _ := se.ExecuteFindCampaignTemplateById(ctx, templateRequest)
		response.Records.TemplateResponse = templateResponse
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error) {

	response := &fs.FindCampaignTemplateByIdResponse{}
	var rows = entsql.Rows{}
	args := FindCampaignTemplateByIdArgs(request)
	query := query.QUERY_FindCampaignTemplateById

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindCampaignTemplateByIdRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindCampaignTemplateByIdResponseVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.TemplateName, &model.CampaignName, &model.DistributionPercent)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindCampaignTemplateByIdRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindCampaignTemplateByIdResponseVO(&model))
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error) {

	model := mappers.MakeAddTargetUserRequestVO(request)
	args := AddTargetUserArgs(model)

	var rows sql.Result
	query := query.QUERY_AddTargetUser

	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddTargetUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddTargetUserRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddTargetUserResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddTargetUserBulk(ctx context.Context, bulkRequest *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error) {

	var args []interface{}
	query := query.QUERY_AddTargetUser
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			query += "(?,?,?)"
		} else {
			query += "(?,?,?),"
		}
		model := mappers.MakeAddTargetUserRequestVO(request)
		args = append(args, AddTargetUserArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddTargetUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddTargetUserRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddTargetUserResponse
	for index, _ := range bulkRequest.Requests {
		responses = append(responses, &fs.AddTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},

			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddTargetUserResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}

func (se *Executor) ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) (*fs.FindTargetUserByIdResponse, error) {

	response := &fs.FindTargetUserByIdResponse{}
	var rows = entsql.Rows{}
	args := FindTargetUserByIdArgs(request)
	query := query.QUERY_FindTargetUserById

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindTargetUserByIdRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindTargetUserByIdResponseVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.UserId, &model.UserType, &model.Attributes)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindTargetUserByIdRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindTargetUserByIdResponseVO(&model))
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error) {

	model := mappers.MakeAddInactionTargetUserRequestVO(request)
	args := AddInactionTargetUserArgs(model)

	var rows sql.Result
	query := query.QUERY_AddInactionTargetUser

	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddInactionTargetUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddInactionTargetUserRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddInactionTargetUserResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddInactionTargetUserBulk(ctx context.Context, bulkRequest *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error) {

	var args []interface{}
	query := query.QUERY_AddInactionTargetUser
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			query += "(?,?,?)"
		} else {
			query += "(?,?,?),"
		}
		model := mappers.MakeAddInactionTargetUserRequestVO(request)
		args = append(args, AddInactionTargetUserArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddInactionTargetUserRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddInactionTargetUserRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddInactionTargetUserResponse
	for index, _ := range bulkRequest.Requests {
		responses = append(responses, &fs.AddInactionTargetUserResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},

			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddInactionTargetUserResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}

func (se *Executor) ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) (*fs.FindInactionTargetUserByCampaignIdResponse, error) {

	response := &fs.FindInactionTargetUserByCampaignIdResponse{}
	var rows = entsql.Rows{}
	args := FindInactionTargetUserByCampaignIdArgs(request)
	query := query.QUERY_FindInactionTargetUserByCampaignId

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindInactionTargetUserByCampaignIdRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindInactionTargetUserByCampaignIdResponseVO{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.UserId, &model.UserType)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindInactionTargetUserByCampaignIdRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindInactionTargetUserByCampaignIdResponseVO(&model))
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func Execute(ctx context.Context, request *fs.MultiRequests) []error {
	/**
	var response []error
	var err error
	multiRequest := request.Request
	for _, customerRequest := range multiRequest {
		addRequest := customerRequest.ARequest
		switch addRequest.(type) {
		case *fs.Request_ReqAddCustomer:
			modifiedRequest := addRequest.(*fs.Request_ReqAddCustomer)
			err = ExecuteAddCustomer(ctx, modifiedRequest.ReqAddCustomer)
			break
		case *fs.Request_ReqAddCustomerBulk:
			modifiedRequest := addRequest.(*fs.Request_ReqAddCustomerBulk)
			err = ExecuteAddCustomerBulk(ctx, modifiedRequest.ReqAddCustomerBulk)
			break
		default:
			logger.Info("Unkown request type")
			break
		}
		response = append(response, err)
	}
	return response
	*/
	return nil
}

func (se *Executor) ExecuteFilterCampaigns(ctx context.Context, request *fs.FilterCampaignRequest) (*fs.FilterCampaignResponse, error) {

	response := &fs.FilterCampaignResponse{}
	var rows = entsql.Rows{}
	query, args := query.GenerateFilterCampaignsQuery(request)

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFilterCampaigns", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindCampaignByIdResponseVO{}
		err := rows.Scan(&model.Id, &model.Namespace, &model.Name, &model.Description, &model.CronExpression, &model.Occurrences, &model.CommunicationChannel, &model.Status, &model.Type, &model.ScheduleType, &model.Query, &model.InactionQuery, &model.InactionDuration, &model.Attributes, &model.CreatedByActorid, &model.CreatedByActortype, &model.UpdatedByActorid, &model.UpdatedByActortype, &model.Version, &model.CreatedAt, &model.UpdatedAt, &model.DeletedAt)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFilterCampaigns", zap.Error(err))
			return nil, err
		}
		campaignResponse := mappers.MakeFindCampaignByIdResponseVO(&model)
		templateRequest := &fs.FindCampaignTemplateByIdRequest{
			CampaignId: campaignResponse.Id,
		}
		templateResponse, _ := se.ExecuteFindCampaignTemplateById(ctx, templateRequest)
		campaignResponse.TemplateResponse = templateResponse
		response.Records = append(response.Records, campaignResponse)
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest, tx ...dialect.Tx) (*fs.AddDynamicDataResponse, error) {

	model := mappers.MakeAddDynamicDataRequestVO(request)
	args := AddDynamicDataArgs(model)

	var rows sql.Result
	query := query.QUERY_AddDynamicData

	var err error
	if tx != nil {
		err = tx[0].Exec(ctx, query, args, &rows)
	} else {
		err = Driver.GetDriver().Exec(ctx, query, args, &rows)
	}
	if err != nil {
		logger.Error("Error could not ExecuteAddDynamicDataRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddDynamicDataRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddDynamicDataResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddDynamicDataBulk(ctx context.Context, bulkRequest *fs.BulkAddDynamicDataRequest, tx ...dialect.Tx) (*fs.BulkAddDynamicDataResponse, error) {

	var args []interface{}
	query := query.QUERY_AddDynamicData
	if idx := strings.Index(query, "(?"); idx != -1 {
		query = query[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			query += "(?,?,?,?)"
		} else {
			query += "(?,?,?,?),"
		}
		model := mappers.MakeAddDynamicDataRequestVO(request)
		args = append(args, AddDynamicDataArgs(model)...)
	}

	var rows sql.Result
	var err error
	if tx != nil {
		err = tx[0].Exec(ctx, query, args, &rows)
	} else {
		err = Driver.GetDriver().Exec(ctx, query, args, &rows)
	}
	if err != nil {
		logger.Error("Error could not BulkAddDynamicDataRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddDynamicDataRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddDynamicDataResponse
	for index, _ := range bulkRequest.Requests {
		responses = append(responses, &fs.AddDynamicDataResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},

			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddDynamicDataResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}

func (se *Executor) ExecuteFindQueryCampaign(ctx context.Context, request *fs.FindQueryCampaignRequest) (*fs.FindQueryCampaignResponse, error) {

	response := &fs.FindQueryCampaignResponse{}
	var rows = entsql.Rows{}
	model := mappers.MakeFindQueryCampaignRequestVO(request)
	args := FindQueryCampaignArgsReq(model)
	currQuery := query.QUERY_FindQueryCampaign

	err := Driver.GetDriver().Query(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindQueryCampaignRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindQueryCampaignResponseVO{}
		err := rows.Scan(&model.Name, &model.Query)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindQueryCampaignRequest", zap.Error(err))
			return nil, err
		}
		response.Records = append(response.Records, mappers.MakeFindQueryCampaignResponseVO(&model))
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}

func (se *Executor) ExecuteAddQueryCampaign(ctx context.Context, request *fs.AddQueryCampaignRequest) (*fs.AddQueryCampaignResponse, error) {

	model := mappers.MakeAddQueryCampaignRequestVO(request)
	args := AddQueryCampaignArgs(model)

	var rows sql.Result
	currQuery := query.QUERY_AddQueryCampaign

	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteAddQueryCampaignRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for AddQueryCampaignRequest", zap.Error(err))
		return nil, err
	}

	response := &fs.AddQueryCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:    1,
		RecordId: cast.ToString(insertedId),
	}

	return response, nil
}

func (se *Executor) ExecuteAddQueryCampaignBulk(ctx context.Context, bulkRequest *fs.BulkAddQueryCampaignRequest) (*fs.BulkAddQueryCampaignResponse, error) {

	var args []interface{}
	currQuery := query.QUERY_AddQueryCampaign
	if idx := strings.Index(currQuery, "(?"); idx != -1 {
		currQuery = currQuery[:idx]
	}

	for index, request := range bulkRequest.Requests {
		if index == len(bulkRequest.Requests)-1 {
			currQuery += "(?,?,?,?);"
		} else {
			currQuery += "(?,?,?,?);,"
		}
		model := mappers.MakeAddQueryCampaignRequestVO(request)
		args = append(args, AddQueryCampaignArgs(model)...)
	}

	var rows sql.Result
	err := Driver.GetDriver().Exec(ctx, currQuery, args, &rows)
	if err != nil {
		logger.Error("Error could not BulkAddQueryCampaignRequest", zap.Error(err))
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		logger.Error("Error could not get lastInsertedId for BulkAddQueryCampaignRequest", zap.Error(err))
		return nil, err
	}

	var responses []*fs.AddQueryCampaignResponse
	for index := range bulkRequest.Requests {
		responses = append(responses, &fs.AddQueryCampaignResponse{
			Status: &common.RequestStatusResult{
				Status: common.RequestStatus_SUCCESS,
			},
			Count:    1,
			RecordId: cast.ToString(cast.ToInt(insertedId) + index),
		})
	}

	response := &fs.BulkAddQueryCampaignResponse{
		Status: &common.RequestStatusResult{
			Status: common.RequestStatus_SUCCESS,
		},

		Count:     cast.ToInt32(len(bulkRequest.Requests)),
		Responses: responses,
	}

	return response, nil
}
func (se *Executor) ExecuteFindControlGroupByCampaignId(ctx context.Context, request *fs.FindControlGroupByCampaignIdRequest) (*fs.FindControlGroupByCampaignIdResponse, error) {
	response := &fs.FindControlGroupByCampaignIdResponse{}
	var rows = entsql.Rows{}
	args := FindControlGroupByCampaignIdArgs(request)
	query := query.QUERY_FindControlGroupByCampaignId

	err := Driver.GetDriver().Query(ctx, query, args, &rows)
	if err != nil {
		logger.Error("Error could not ExecuteFindControlGroupByCampaignIdRequest", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		model := models.FindControlGroupByCampaignIdRequestV0{}
		err := rows.Scan(&model.Id, &model.CampaignId, &model.Attributes, &model.BloomFilter)
		if err != nil {
			logger.Error("Error while fetching rows for ExecuteFindControlGroupByCampaignIdRequest", zap.Error(err))
			return nil, err
		}

		response.Records = mappers.MakeFindControlGroupByCampaignIdResponse(&model)
	}
	response.Status = &common.RequestStatusResult{
		Status: common.RequestStatus_SUCCESS,
	}
	return response, nil
}
