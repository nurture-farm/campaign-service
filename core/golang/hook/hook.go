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
	"code.nurture.farm/platform/CampaignService/core/golang/cache"
	query "code.nurture.farm/platform/CampaignService/core/golang/database"
	"context"
	"github.com/spf13/viper"
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

type ExecuteRequestController struct {
}

var ExecuteRequestExecutor *ExecuteRequestController
var EVENT_DB_NAME string
var EVENT_TABLE_NAME string

func init() {
	AddCampaignExecutor = &GenericAddCampaignExecutor{
		AddCampaignInterface: &AddCampaignController{},
	}
	BulkAddCampaignExecutor = &GenericAddCampaignExecutorBulk{
		AddCampaignBulkInterface: &BulkAddCampaignController{},
	}
	UpdateCampaignExecutor = &GenericUpdateCampaignExecutor{
		UpdateCampaignInterface: &UpdateCampaignController{},
	}
	BulkUpdateCampaignExecutor = &GenericUpdateCampaignExecutorBulk{
		UpdateCampaignBulkInterface: &BulkUpdateCampaignController{},
	}
	AddCampaignTemplateExecutor = &GenericAddCampaignTemplateExecutor{
		AddCampaignTemplateInterface: &AddCampaignTemplateController{},
	}
	BulkAddCampaignTemplateExecutor = &GenericAddCampaignTemplateExecutorBulk{
		AddCampaignTemplateBulkInterface: &BulkAddCampaignTemplateController{},
	}
	AddNewCampaignExecutor = &GenericAddNewCampaignExecutor{
		AddNewCampaignInterface: &AddNewCampaignController{},
	}
	BulkAddNewCampaignExecutor = &GenericAddNewCampaignExecutorBulk{
		AddNewCampaignBulkInterface: &BulkAddNewCampaignController{},
	}
	CampaignExecutor = &GenericCampaignExecutor{
		CampaignInterface: &CampaignController{},
	}
	FindCampaignByIdExecutor = &GenericFindCampaignByIdExecutor{
		FindCampaignByIdInterface: &FindCampaignByIdController{},
	}
	FindCampaignTemplateByIdExecutor = &GenericFindCampaignTemplateByIdExecutor{
		FindCampaignTemplateByIdInterface: &FindCampaignTemplateByIdController{},
	}
	AddTargetUserExecutor = &GenericAddTargetUserExecutor{
		AddTargetUserInterface: &AddTargetUserController{},
	}
	BulkAddTargetUserExecutor = &GenericAddTargetUserExecutorBulk{
		AddTargetUserBulkInterface: &BulkAddTargetUserController{},
	}
	FindTargetUserByIdExecutor = &GenericFindTargetUserByIdExecutor{
		FindTargetUserByIdInterface: &FindTargetUserByIdController{},
	}
	AddInactionTargetUserExecutor = &GenericAddInactionTargetUserExecutor{
		AddInactionTargetUserInterface: &AddInactionTargetUserController{},
	}
	BulkAddInactionTargetUserExecutor = &GenericAddInactionTargetUserExecutorBulk{
		AddInactionTargetUserBulkInterface: &BulkAddInactionTargetUserController{},
	}
	FindInactionTargetUserByCampaignIdExecutor = &GenericFindInactionTargetUserByCampaignIdExecutor{
		FindInactionTargetUserByCampaignIdInterface: &FindInactionTargetUserByCampaignIdController{},
	}

	TestNewCampaignExecutor = &GenericTestNewCampaignExecutor{
		TestNewCampaignInterface: &TestNewCampaignController{},
	}
	AthenaQueryExecutor = &GenericAthenaQueryExecutor{
		AthenaQueryInterface: &AthenaQueryController{},
	}
	FilterCampaignExecutor = &GenericFilterCampaignExecutor{
		FilterCampaignInterface: &FilterCampaignController{},
	}
	GetDynamicDataByKeyExecutor = &GenericGetDynamicDataByKeyExecutor{
		GetDynamicDataByKeyInterface: &GetDynamicDataByKeyController{},
	}
	AddDynamicDataExecutor = &GenericAddDynamicDataExecutor{
		AddDynamicDataInterface: &AddDynamicDataController{},
	}
	BulkAddDynamicDataExecutor = &GenericAddDynamicDataExecutorBulk{
		AddDynamicDataBulkInterface: &BulkAddDynamicDataController{},
	}
	FindQueryCampaignExecutor = &GenericFindQueryCampaignExecutor{
		FindQueryCampaignInterface: &FindQueryCampaignController{},
	}
	AddQueryCampaignExecutor = &GenericAddQueryCampaignExecutor{
		AddQueryCampaignInterface: &AddQueryCampaignController{},
	}
	BulkAddQueryCampaignExecutor = &GenericAddQueryCampaignExecutorBulk{
		AddQueryCampaignBulkInterface: &BulkAddQueryCampaignController{},
	}
	ScheduleUserJourneyCampaignExecutor = &GenericScheduleUserJourneyCampaignExecutor{
		ScheduleUserJourneyCampaignInterface: &ScheduleUserJourneyCampaignController{},
	}
	FindUserJourneyCampaignByIdExecutor = &GenericFindUserJourneyCampaignByIdExecutor{
		FindUserJourneyCampaignByIdInterface: &FindUserJourneyCampaignByIdController{},
	}
	FilterUserJourneyCampaignExecutor = &GenericFilterUserJourneyCampaignExecutor{
		FilterUserJourneyCampaignInterface: &FilterUserJourneyCampaignController{},
	}
	UserJourneyCampaignExecutor = &GenericUserJourneyCampaignExecutor{
		UserJourneyCampaignInterface: &UserJourneyCampaignController{},
	}
	ExecuteRequestExecutor = &ExecuteRequestController{}
}

func PreStartUpHook() {
	//This will run on application boot up before gRPC server starts
	EVENT_DB_NAME = viper.GetString(EVENTS_DATABASE_NAME)
	EVENT_TABLE_NAME = viper.GetString(EVENTS_TABLE_NAME)
	query.QUERY_ATHENA_SMS_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.athena_db_queries.sms_query")
	query.QUERY_ATHENA_APP_NOTIFICATION_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.athena_db_queries.app_notification_query")
	query.QUERY_ATHENA_WHATSAPP_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.athena_db_queries.whatsapp_query")
	query.QUERY_ATHENA_EMAIL_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.athena_db_queries.email_query")
	query.QUERY_USERLIST_SMS_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.user_list_queries.sms_query")
	query.QUERY_USERLIST_APP_NOTIFICATION_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.user_list_queries.app_notification_query")
	query.QUERY_USERLIST_WHATSAPP_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.user_list_queries.whatsapp_query")
	query.QUERY_USERLIST_EMAIL_ATTRIBUTES = viper.GetString("user_channel_attribute_queries.user_list_queries.email_query")
	cache.InitCache()
}

func PostStartUpHook() {
	//This will run on application boot up after gRPC server starts
}

func (rc *ExecuteRequestController) OnRequest(ctx context.Context, request interface{}) interface{} {

	return nil
}

func (rc *ExecuteRequestController) OnResponse(ctx context.Context, request interface{}, response interface{}) interface{} {

	return nil
}

func (rc *ExecuteRequestController) OnError(ctx context.Context, request interface{}, response interface{}, err error) interface{} {

	return nil
}
