package setup

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/service"
	"context"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sCampaignService struct {
	fs.UnimplementedCampaignServiceServer
	sCampaignServiceActivities
	wfClient client.Client
}

var CampaignService *sCampaignService = &sCampaignService{
	wfClient: WorkflowClient,
}

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

var WorkflowClient, WorkerConfig = getWorkflowClient()

func getWorkflowClient() (client.Client, map[string]string) {
	workerConfig := viper.GetStringMapString("worker_config")

	logger.Info("Worker config",
		zap.String("TemporalHostPort", database.TemporalHostPort),
		zap.String("TemporalNamespace", database.TemporalNamespace),
		zap.String("AthenaDBName", database.AthenaDBName))

	c, err := client.NewClient(client.Options{
		Namespace: database.TemporalNamespace,
		HostPort:  database.TemporalHostPort,
	})
	if err != nil {
		logger.Fatal("Unable to create client", zap.Error(err))
	}
	return c, workerConfig
}

func (fs *sCampaignService) Close() {
	fs.wfClient.Close()
}

func (fs *sCampaignService) ExecuteAddCampaign(ctx context.Context, request *fs.AddCampaignRequest) (*fs.AddCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddCampaign(ctx, request)
}

func (fs *sCampaignService) ExecuteAddCampaignBulk(ctx context.Context, request *fs.BulkAddCampaignRequest) (*fs.BulkAddCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddCampaignBulk(ctx, request)
}

func (fs *sCampaignService) ExecuteUpdateCampaign(ctx context.Context, request *fs.UpdateCampaignRequest) (*fs.UpdateCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteUpdateCampaign(ctx, request)
}

func (fs *sCampaignService) ExecuteAddCampaignTemplate(ctx context.Context, request *fs.AddCampaignTemplateRequest) (*fs.AddCampaignTemplateResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddCampaignTemplate(ctx, request)
}

func (fs *sCampaignService) ExecuteAddCampaignTemplateBulk(ctx context.Context, request *fs.BulkAddCampaignTemplateRequest) (*fs.BulkAddCampaignTemplateResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddCampaignTemplateBulk(ctx, request)
}

func (fs *sCampaignService) ExecuteAddNewCampaign(ctx context.Context, request *fs.AddNewCampaignRequest) (*fs.AddNewCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddNewCampaign(ctx, request)
}

func (fs *sCampaignService) ExecuteAddNewCampaignBulk(ctx context.Context, request *fs.BulkAddNewCampaignRequest) (*fs.BulkAddNewCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddNewCampaignBulk(ctx, request)
}

func (fs *sCampaignService) ExecuteCampaign(ctx context.Context, request *fs.CampaignRequest) (*fs.CampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteCampaign(ctx, request)
}

func (fs *sCampaignService) ExecuteFindCampaignById(ctx context.Context, request *fs.FindCampaignByIdRequest) (*fs.FindCampaignByIdResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteFindCampaignById(ctx, request)
}

func (fs *sCampaignService) ExecuteFindCampaignTemplateById(ctx context.Context, request *fs.FindCampaignTemplateByIdRequest) (*fs.FindCampaignTemplateByIdResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteFindCampaignTemplateById(ctx, request)
}

func (fs *sCampaignService) ExecuteAddTargetUser(ctx context.Context, request *fs.AddTargetUserRequest) (*fs.AddTargetUserResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddTargetUser(ctx, request)
}

func (fs *sCampaignService) ExecuteAddTargetUserBulk(ctx context.Context, request *fs.BulkAddTargetUserRequest) (*fs.BulkAddTargetUserResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddTargetUserBulk(ctx, request)
}

func (fs *sCampaignService) ExecuteFindTargetUserById(ctx context.Context, request *fs.FindTargetUserByIdRequest) (*fs.FindTargetUserByIdResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteFindTargetUserById(ctx, request)
}

func (fs *sCampaignService) ExecuteAddInactionTargetUser(ctx context.Context, request *fs.AddInactionTargetUserRequest) (*fs.AddInactionTargetUserResponse, error) {

	return service.ExecuteAddInactionTargetUser(ctx, request), nil
}

func (fs *sCampaignService) ExecuteAddInactionTargetUserBulk(ctx context.Context, request *fs.BulkAddInactionTargetUserRequest) (*fs.BulkAddInactionTargetUserResponse, error) {

	return service.ExecuteAddInactionTargetUserBulk(ctx, request), nil
}

func (fs *sCampaignService) ExecuteFindInactionTargetUserByCampaignId(ctx context.Context, request *fs.FindInactionTargetUserByCampaignIdRequest) (*fs.FindInactionTargetUserByCampaignIdResponse, error) {

	return service.ExecuteFindInactionTargetUserByCampaignId(ctx, request), nil
}

func (fs *sCampaignService) Execute(ctx context.Context, request *fs.MultiRequests) (*fs.MultiResponses, error) {

	// TO-DO
	return nil, nil
}

func (fs *sCampaignService) ExecuteTestNewCampaign(ctx context.Context, request *fs.TestNewCampaignRequest) (*fs.TestNewCampaignResponse, error) {

	return service.ExecuteTestNewCampaign(ctx, request), nil
}

func (fs *sCampaignService) ExecuteAthenaQuery(ctx context.Context, request *fs.AthenaQueryRequest) (*fs.AthenaQueryResponse, error) {

	return service.ExecuteAthenaQuery(ctx, request), nil
}

func (fs *sCampaignService) ExecuteFilterCampaigns(ctx context.Context, request *fs.FilterCampaignRequest) (*fs.FilterCampaignResponse, error) {

	return service.ExecuteFilterCampaigns(ctx, request), nil
}

func (fs *sCampaignService) ExecuteTestCampaignById(ctx context.Context, request *fs.TestCampaignByIdRequest) (*fs.TestCampaignByIdResponse, error) {

	return service.ExecuteTestCampaignById(ctx, request), nil
}

func (fs *sCampaignService) ExecuteGetDynamicDataByKey(ctx context.Context, request *fs.GetDynamicDataByKeyRequest) (*fs.GetDynamicDataByKeyResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteGetDynamicDataByKey(ctx, request)
}

func (fs *sCampaignService) ExecuteAddDynamicData(ctx context.Context, request *fs.AddDynamicDataRequest) (*fs.AddDynamicDataResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddDynamicData(ctx, request)
}

func (fs *sCampaignService) ExecuteAddDynamicDataBulk(ctx context.Context, request *fs.BulkAddDynamicDataRequest) (*fs.BulkAddDynamicDataResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteAddDynamicDataBulk(ctx, request)
}

func (fs *sCampaignService) ExecuteFindQueryCampaign(ctx context.Context, request *fs.FindQueryCampaignRequest) (*fs.FindQueryCampaignResponse, error) {

	return service.ExecuteFindQueryCampaign(ctx, request), nil
}

func (fs *sCampaignService) ExecuteAddQueryCampaign(ctx context.Context, request *fs.AddQueryCampaignRequest) (*fs.AddQueryCampaignResponse, error) {

	return service.ExecuteAddQueryCampaign(ctx, request), nil
}

func (fs *sCampaignService) ExecuteAddQueryCampaignBulk(ctx context.Context, request *fs.BulkAddQueryCampaignRequest) (*fs.BulkAddQueryCampaignResponse, error) {

	return service.ExecuteAddQueryCampaignBulk(ctx, request), nil
}

func (fs *sCampaignService) ExecuteScheduleUserJourneyCampaign(ctx context.Context, request *fs.ScheduleUserJourneyCampaignRequest) (*fs.ScheduleUserJourneyCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteScheduleUserJourneyCampaign(ctx, request)
}

func (fs *sCampaignService) ExecuteFindUserJourneyCampaignById(ctx context.Context, request *fs.FindUserJourneyCampaignByIdRequest) (*fs.FindUserJourneyCampaignByIdResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteFindUserJourneyCampaignById(ctx, request)
}

func (fs *sCampaignService) ExecuteFilterUserJourneyCampaigns(ctx context.Context, request *fs.FilterUserJourneyCampaignRequest) (*fs.FilterUserJourneyCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteFilterUserJourneyCampaigns(ctx, request)
}

func (fs *sCampaignService) ExecuteUserJourneyCampaign(ctx context.Context, request *fs.UserJourneyCampaignRequest) (*fs.UserJourneyCampaignResponse, error) {

	return fs.sCampaignServiceActivities.ExecuteUserJourneyCampaign(ctx, request)
}
