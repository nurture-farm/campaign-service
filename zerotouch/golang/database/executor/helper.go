package executor

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	"github.com/nurture-farm/campaign-service/zerotouch/golang/database/models"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = getLogger()

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

func AddCampaignArgs(model *models.AddCampaignRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.Namespace)
	args = append(args, model.Name)
	args = append(args, model.Description)
	args = append(args, model.CronExpression)
	args = append(args, model.Occurrences)
	args = append(args, model.CommunicationChannel)
	args = append(args, model.Status)
	args = append(args, model.Type)
	args = append(args, model.ScheduleType)
	args = append(args, model.Query)
	args = append(args, model.InactionQuery)
	args = append(args, model.InactionDuration)
	args = append(args, model.Attributes)
	args = append(args, model.CreatedByActorid)
	args = append(args, model.CreatedByActortype)

	return args
}
func UpdateCampaignArgs(model *models.UpdateCampaignRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.Name)
	args = append(args, model.CronExpression)
	args = append(args, model.Status)
	args = append(args, model.Query)
	args = append(args, model.Namespace)
	args = append(args, model.Occurrences)
	args = append(args, model.CommunicationChannel)
	args = append(args, model.Type)
	args = append(args, model.Description)
	args = append(args, model.ScheduleType)
	args = append(args, model.InactionQuery)
	args = append(args, model.InactionDuration)
	args = append(args, model.Attributes)
	args = append(args, model.UpdatedByActorid)
	args = append(args, model.UpdatedByActortype)
	args = append(args, model.Id)

	return args
}
func AddCampaignTemplateArgs(model *models.AddCampaignTemplateRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.TemplateName)
	args = append(args, model.CampaignName)
	args = append(args, model.DistributionPercent)

	return args
}
func DeleteCampaignTemplateArgs(model *models.DeleteCampaignTemplateRequestVO) []interface{} {
	var args []interface{}
	args = append(args, model.CampaignId)
	return args
}

func FindCampaignByIdArgs(request *fs.FindCampaignByIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.Id)

	return args
}
func FindControlGroupByCampaignIdArgs(request *fs.FindControlGroupByCampaignIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)

	return args
}
func FindCampaignTemplateByIdArgs(request *fs.FindCampaignTemplateByIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)

	return args
}
func AddTargetUserArgs(model *models.AddTargetUserRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.UserId)
	args = append(args, model.UserType)
	args = append(args, model.Attributes)

	return args
}
func AddControlGroupArgs(model *models.AddControlGroupRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.Attributes)
	args = append(args, model.BloomFilter)
	//args = append(args, model.BloomFilterText)

	return args
}
func FindTargetUserByIdArgs(request *fs.FindTargetUserByIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)

	return args
}

func AddInactionTargetUserArgs(model *models.AddInactionTargetUserRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.UserId)
	args = append(args, model.UserType)

	return args
}

func FindInactionTargetUserByCampaignIdArgs(request *fs.FindInactionTargetUserByCampaignIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)

	return args
}

func GetDynamicDataByKeyArgs(request *fs.GetDynamicDataByKeyRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)
	args = append(args, request.DynamicKey)

	return args
}

func AddDynamicDataArgs(model *models.AddDynamicDataRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.DynamicKey)
	args = append(args, model.CtaLink)
	args = append(args, model.Media)

	return args
}

func FindQueryCampaignArgs(request *fs.FindQueryCampaignRequest) []interface{} {

	var args []interface{}
	args = append(args, request.Type)

	return args
}

func FindQueryCampaignArgsReq(model *models.FindQueryCampaignRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.Type)

	return args
}
func AddQueryCampaignArgs(model *models.AddQueryCampaignRequestVO) []interface{} {

	var args []interface{}
	args = append(args, model.Name)
	args = append(args, model.Type)
	args = append(args, model.Query)
	args = append(args, model.UpdatedBy)

	return args
}
