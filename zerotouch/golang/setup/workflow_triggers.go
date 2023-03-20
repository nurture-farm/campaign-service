package setup

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	Common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	ggw "github.com/nurture-farm/Contracts/Workflows/GeneralGoWorkflows/Gen/GoGeneralGoWorkflows"
	"code.nurture.farm/platform/CampaignService/core/golang/hook"
	"context"
	"fmt"
	ptypes "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"time"
)

const ExecuteCampaignWorkflowName = "ExecuteCampaignWorkflow"
const ExecuteUserJourneyCampaignWorkflowName = "ExecuteUserJourneyCampaignWorkflow"

func getWorkflowOptions(id string, cronSchedule string, occurence int32, scheduleType Common.CampaignScheduleType, inactionDuration *duration.Duration) client.StartWorkflowOptions {
	workflowOptions := client.StartWorkflowOptions{
		ID:                    id,
		TaskQueue:             "GeneralGoWorker",
		WorkflowTaskTimeout:   60 * time.Minute,
		WorkflowRunTimeout:    180 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
		CronSchedule:          cronSchedule,
		//RetryPolicy: &temporal.RetryPolicy{
		//	InitialInterval:    60 * time.Second,
		//	BackoffCoefficient: 2,
		//	MaximumInterval:    10 * time.Minute,
		//	MaximumAttempts:    6,
		//},
	}
	duration, _ := ptypes.Duration(inactionDuration)
	if occurence > 0 {
		workflowOptions.WorkflowExecutionTimeout = time.Duration(occurence) * time.Hour * 24
	}
	if scheduleType == Common.CampaignScheduleType_INACTION_OVER_TIME {
		workflowOptions.WorkflowExecutionTimeout = duration + time.Hour*3 //adding 3 hours to complete inaction query
		workflowOptions.WorkflowRunTimeout = duration + time.Hour*3       //updated for engagement journey update
	}
	return workflowOptions
}

func triggerExecuteUserJourneyCampaignWorkflow(ctx context.Context, sCampaignService *sCampaignService,
	campaignId int64, cronSchedule string, scheduleAsCron bool, engagementVertexId int64, referenceId string, workflowId string, waitDuration time.Duration) error {

	if scheduleAsCron && cronSchedule == hook.EMPTY {
		return fmt.Errorf("MISSING_CRON_SCHEDULE")
	}
	durationProto := ptypes.DurationProto(waitDuration)
	wfRequest := &ggw.ExecuteUserJourneyCampaignRequest{
		CampaignId:         campaignId,
		EngagementVertexId: engagementVertexId,
		ReferenceId:        referenceId,
		WaitDuration:       durationProto,
	}
	scheduleType := Common.CampaignScheduleType_INACTION_OVER_TIME //in order to support non-first engagement Node timeout issue
	we, err := sCampaignService.wfClient.ExecuteWorkflow(ctx, getWorkflowOptions(workflowId, cronSchedule,
		0, scheduleType, durationProto), ExecuteUserJourneyCampaignWorkflowName, wfRequest)
	if err != nil {
		logger.Error("START ExecuteUserJourneyCampaignWorkflow FAILED: ",
			zap.Error(err), zap.Any("workflowRun", we),
			zap.Any("wfRequest", wfRequest))
		return err
	}
	we.GetRunID()
	logger.Info("ExecuteUserJourneyCampaignWorkflow:::STARTED", zap.Any("workflowId", workflowId),
		zap.Any("wfRequest", wfRequest))
	return nil
}

func cancelExecuteUserJourneyCampaignWorkflow(ctx context.Context, sCampaignService *sCampaignService, workflowId string) {

	err := sCampaignService.wfClient.TerminateWorkflow(context.Background(), workflowId, "", "UserJourneyCampaignUpdated")
	if err != nil {
		logger.Error("CANCEL ExecuteUserJourneyCampaignWorkflow FAILED ", zap.Error(err), zap.Any("workflowId", workflowId))
		return
	}
	logger.Info("CANCEL ExecuteUserJourneyCampaignWorkflow SUCCESS ", zap.Any("workflowId", workflowId))
	return
}

func triggerExecuteCampaignWorkflow(ctx context.Context, request *fs.AddNewCampaignRequest, sCampaignService *sCampaignService, campaignId int64, workflowId string) {

	if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
		logger.Info("Triggering ExecuteCampaignWorkflow for request", zap.Any("addCampaignRequest", request.AddCampaignRequest))
	} else {
		logger.Info("Triggering ExecuteCampaignWorkflow for request", zap.Any("request", request))
	}

	wfRequest := &ggw.ExecuteCampaignRequest{
		CampaignId:           campaignId,
		CampaignScheduleType: request.AddCampaignRequest.CampaignScheduleType,
		InactionDuration:     request.AddCampaignRequest.InactionDuration,
	}
	we, err := sCampaignService.wfClient.ExecuteWorkflow(ctx, getWorkflowOptions(workflowId, request.AddCampaignRequest.CronExpression,
		request.AddCampaignRequest.Occurrences, request.AddCampaignRequest.CampaignScheduleType, request.AddCampaignRequest.InactionDuration),
		ExecuteCampaignWorkflowName, wfRequest)
	if err != nil {
		logger.Error("START ExecuteCampaignWorkflow FAILED: ",
			zap.Error(err), zap.Any("workflowRun", we),
			zap.Any("wfRequest", wfRequest))
		return
	}
	we.GetRunID()
	logger.Info("ExecuteCampaignWorkflow:::STARTED", zap.Any("workflowId", workflowId),
		zap.Any("wfRequest", wfRequest))
	return
}

func cancelExecuteCampaignWorkflow(ctx context.Context, sCampaignService *sCampaignService, workflowId string) {

	err := sCampaignService.wfClient.TerminateWorkflow(context.Background(), workflowId, "", "CampaignUpdated")
	if err != nil {
		logger.Error("CANCEL ExecuteCampaignWorkflow FAILED ", zap.Error(err), zap.Any("workflowId", workflowId))
		return
	}
	logger.Info("CANCEL ExecuteCampaignWorkflow SUCCESS ", zap.Any("workflowId", workflowId))
	return
}
