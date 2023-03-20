package setup

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"github.com/nurture-farm/campaign-service/core/golang/hook"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/service"
	"context"
	"github.com/spf13/cast"
	"time"
)

type sCampaignServiceActivities struct {
}

var CampaignServiceActivities *sCampaignServiceActivities = &sCampaignServiceActivities{}

func (fs *sCampaignServiceActivities) ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error) {

	return service.ExecuteAddCampaign(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddCampaignBulk(ctx context.Context, request *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error) {

	return service.ExecuteAddCampaignBulk(ctx, request), nil
}

func (s *sCampaignServiceActivities) ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest) (*fs.UpdateCampaignResponse, error) {

	response, addNewCapaignRequest := service.ExecuteUpdateCampaign(ctx, request)
	if response.Status.Status == common.RequestStatus_SUCCESS {
		campaignQueryType, err := hook.GetCampaignQueryType(ctx, request.Id)
		if err != nil {
			return &fs.UpdateCampaignResponse{
				Status: &common.RequestStatusResult{
					Status: common.RequestStatus_INTERNAL_ERROR,
				},
			}, nil
		}
		workflowId := hook.GetWorkflowId(request.Id)
		if campaignQueryType == common.CampaignQueryType_USER_JOURNEY {
			engagementStartVertexId, err := hook.GetStartEngagementVertex(ctx, request.Id)
			if err != nil {
				return &fs.UpdateCampaignResponse{
					Status: &common.RequestStatusResult{
						Status: common.RequestStatus_INTERNAL_ERROR,
					},
				}, nil
			}
			workflowId = hook.GetUserJourneyWorkflowId(request.Id, engagementStartVertexId)
		}
		cancelExecuteCampaignWorkflow(ctx, CampaignService, workflowId)
		if request.AddCampaignRequest.Status != common.CampaignStatus_HALTED {
			triggerExecuteCampaignWorkflow(ctx, addNewCapaignRequest, CampaignService, cast.ToInt64(response.RecordId), hook.GetWorkflowId(cast.ToInt64(response.RecordId)))
		}
	}
	return response, nil
}

func (fs *sCampaignServiceActivities) ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error) {

	return service.ExecuteAddCampaignTemplate(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddCampaignTemplateBulk(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error) {

	return service.ExecuteAddCampaignTemplateBulk(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddNewCampaign(ctx context.Context, request *fs.AddNewCampaignRequest) (*fs.AddNewCampaignResponse, error) {

	response := service.ExecuteAddNewCampaign(ctx, request)
	if response.Status.Status == common.RequestStatus_SUCCESS {
		triggerExecuteCampaignWorkflow(ctx, request, CampaignService, cast.ToInt64(response.RecordId), hook.GetWorkflowId(cast.ToInt64(response.RecordId)))
	}
	return response, nil
}

func (fs *sCampaignServiceActivities) ExecuteAddNewCampaignBulk(ctx context.Context, request *fs.BulkAddNewCampaignRequest) (*fs.BulkAddNewCampaignResponse, error) {

	return service.ExecuteAddNewCampaignBulk(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteCampaign(ctx context.Context, request *fs.CampaignRequest) (*fs.CampaignResponse, error) {

	return service.ExecuteCampaign(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error) {

	return service.ExecuteFindCampaignById(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error) {

	return service.ExecuteFindCampaignTemplateById(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error) {

	return service.ExecuteAddTargetUser(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddTargetUserBulk(ctx context.Context, request *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error) {

	return service.ExecuteAddTargetUserBulk(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) (*fs.FindTargetUserByIdResponse, error) {

	return service.ExecuteFindTargetUserById(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error) {
	return service.ExecuteAddInactionTargetUser(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddInactionTargetUserBulk(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error) {
	return service.ExecuteAddInactionTargetUserBulk(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) (*fs.GetDynamicDataByKeyResponse, error) {

	return service.ExecuteGetDynamicDataByKey(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest) (*fs.AddDynamicDataResponse, error) {

	return service.ExecuteAddDynamicData(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteAddDynamicDataBulk(ctx context.Context, request *fs.BulkAddDynamicDataRequest) (*fs.BulkAddDynamicDataResponse, error) {

	return service.ExecuteAddDynamicDataBulk(ctx, request), nil
}

func (cs *sCampaignServiceActivities) ExecuteScheduleUserJourneyCampaign(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) (*fs.ScheduleUserJourneyCampaignResponse, error) {

	response := service.ExecuteScheduleUserJourneyCampaign(ctx, request)
	if response.Status.Status == common.RequestStatus_SUCCESS {
		engagementStartVertexId, err := hook.GetStartEngagementVertex(ctx, response.CampaignId)
		if err != nil {
			response.Status = &common.RequestStatusResult{
				Status:        common.RequestStatus_INTERNAL_ERROR,
				ErrorMessages: []string{err.Error()},
			}
			return response, nil
		}
		if !request.TriggerCampaign {
			cancelExecuteCampaignWorkflow(ctx, CampaignService, hook.GetUserJourneyWorkflowId(response.CampaignId, engagementStartVertexId))
		} else {
			err := triggerExecuteUserJourneyCampaignWorkflow(ctx, CampaignService, response.CampaignId, response.CronSchedule, true, engagementStartVertexId,
				response.ReferenceId, hook.GetUserJourneyWorkflowId(response.CampaignId, engagementStartVertexId), time.Second*0)
			if err != nil {
				response.Status = &common.RequestStatusResult{
					Status:        common.RequestStatus_INTERNAL_ERROR,
					ErrorMessages: []string{err.Error()},
				}
			}
		}
	}
	return response, nil
}

func (fs *sCampaignServiceActivities) ExecuteFindUserJourneyCampaignById(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) (*fs.FindUserJourneyCampaignByIdResponse, error) {

	return service.ExecuteFindUserJourneyCampaignById(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteFilterUserJourneyCampaigns(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) (*fs.FilterUserJourneyCampaignResponse, error) {

	return service.ExecuteFilterUserJourneyCampaigns(ctx, request), nil
}

func (fs *sCampaignServiceActivities) ExecuteUserJourneyCampaign(ctx context.Context, request *fs.UserJourneyCampaignRequest) (*fs.UserJourneyCampaignResponse, error) {

	response := service.ExecuteUserJourneyCampaign(ctx, request)
	if response.Status.Status == common.RequestStatus_SUCCESS {
		nextEngagementVertices, err := hook.GetNextEngagementVertices(ctx, request.CampaignId, request.EngagementVertexId)
		if err != nil {
			response.Status = &common.RequestStatusResult{
				Status:        common.RequestStatus_INTERNAL_ERROR,
				ErrorMessages: []string{err.Error()},
			}
			return response, nil
		}
		for _, nextEngagementVertex := range nextEngagementVertices {
			waitDuration := hook.GetWaitDuration(ctx, nextEngagementVertex.WaitDuration.Int64, nextEngagementVertex.WaitTime.Time,
				nextEngagementVertex.WaitType.String)
			err = triggerExecuteUserJourneyCampaignWorkflow(ctx, CampaignService, request.CampaignId, hook.EMPTY, false, nextEngagementVertex.Id.Int64,
				request.ReferenceId, hook.GetUserJourneyWorkflowId(request.CampaignId, nextEngagementVertex.Id.Int64), waitDuration)
			if err != nil {
				response.Status = &common.RequestStatusResult{
					Status:        common.RequestStatus_INTERNAL_ERROR,
					ErrorMessages: []string{err.Error()},
				}
				return response, nil
			}
		}
	}
	return response, nil
}
