package metrics

import (
	"context"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	metrics "github.com/nurture-farm/go-common/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	SERVICE_NAME = "CampaignService"
	DATABASE     = "database"
)

var Metrics metrics.MetricWrapper

var (
	AddCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ADDCAMPAIGN",
		Help:       "Sumary metrics for AddCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ADDCAMPAIGN_ERROR",
		Help: "Sumary metrics for AddCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKADDCAMPAIGN",
		Help:       "Sumary metrics for BulkAddCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKADDCAMPAIGN_ERROR",
		Help: "Sumary metrics for BulkAddCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	UpdateCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_UPDATECAMPAIGN",
		Help:       "Sumary metrics for UpdateCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	UpdateCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_UPDATECAMPAIGN_ERROR",
		Help: "Sumary metrics for UpdateCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddCampaignTemplate_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ADDCAMPAIGNTEMPLATE",
		Help:       "Sumary metrics for AddCampaignTemplate",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddCampaignTemplate_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ADDCAMPAIGNTEMPLATE_ERROR",
		Help: "Sumary metrics for AddCampaignTemplate_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddCampaignTemplate_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKADDCAMPAIGNTEMPLATE",
		Help:       "Sumary metrics for BulkAddCampaignTemplate",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddCampaignTemplate_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKADDCAMPAIGNTEMPLATE_ERROR",
		Help: "Sumary metrics for BulkAddCampaignTemplate_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddNewCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ADDNEWCAMPAIGN",
		Help:       "Sumary metrics for AddNewCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddNewCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ADDNEWCAMPAIGN_ERROR",
		Help: "Sumary metrics for AddNewCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddNewCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKADDNEWCAMPAIGN",
		Help:       "Sumary metrics for BulkAddNewCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddNewCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKADDNEWCAMPAIGN_ERROR",
		Help: "Sumary metrics for BulkAddNewCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetUserList_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_GETUSERLIST",
		Help:       "Sumary metrics for GetUserList",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	GetUserList_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETUSERLIST_ERROR",
		Help: "Sumary metrics for GetUserList_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetTentativeList_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETTentativeList_ERROR",
		Help: "Summary metrics for GetTentativeList_Error",
	}, []string{"nservice", "ntype", "nerror"})
)
var (
	FindCampaignById_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FINDCAMPAIGNBYID",
		Help:       "Sumary metrics for FindCampaignById",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindCampaignById_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FINDCAMPAIGNBYID_ERROR",
		Help: "Sumary metrics for FindCampaignById_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	FindCampaignTemplateById_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FINDCAMPAIGNTEMPLATEBYID",
		Help:       "Sumary metrics for FindCampaignTemplateById",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindCampaignTemplateById_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FINDCAMPAIGNTEMPLATEBYID_ERROR",
		Help: "Sumary metrics for FindCampaignTemplateById_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddTargetUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ADDTARGETUSER",
		Help:       "Sumary metrics for AddTargetUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddTargetUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ADDTARGETUSER_ERROR",
		Help: "Sumary metrics for AddTargetUser_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddTargetUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKADDTARGETUSER",
		Help:       "Sumary metrics for BulkAddTargetUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddTargetUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKADDTARGETUSER_ERROR",
		Help: "Sumary metrics for BulkAddTargetUser_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	FindTargetUserById_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FINDTARGETUSERBYID",
		Help:       "Sumary metrics for FindTargetUserById",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindTargetUserById_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FINDTARGETUSERBYID_ERROR",
		Help: "Sumary metrics for FindTargetUserById_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AthenaQuery_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ATHENAQUERY",
		Help:       "Sumary metrics for ATHENAQUERY",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	ATHENAQUERY_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ATHENAQUERY_ERROR",
		Help: "Sumary metrics for ATHENAQUERY_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	ATTRIBUTES_MARSHAL_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ATTRIBUTES_MARSHAL_ERROR",
		Help: "Sumary metrics for ATTRIBUTES_MARSHALY_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	ATTRIBUTES_UNMARSHAL_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ATTRIBUTES_UNMARSHAL_ERROR",
		Help: "Sumary metrics for ATTRIBUTES_UNMARSHAL_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddInactionTargetUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_ADDINACTIONTARGETUSER",
		Help:       "Sumary metrics for AddInactionTargetUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddInactionTargetUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_ADDINACTIONTARGETUSER_ERROR",
		Help: "Sumary metrics for AddInactionTargetUser_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddInactionTargetUser_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_BULKADDINACTIONTARGETUSER",
		Help:       "Sumary metrics for BulkAddInactionTargetUser",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddInactionTargetUser_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_BULKADDINACTIONTARGETUSER_ERROR",
		Help: "Sumary metrics for BulkAddInactionTargetUser_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	FindInactionTargetUserByCampaignId_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_FINDINACTIONTARGETUSERBYCAMPAIGNID",
		Help:       "Sumary metrics for FindInactionTargetUserByCampaignId",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindInactionTargetUserByCampaignId_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_FINDINACTIONTARGETUSERBYCAMPAIGNID_ERROR",
		Help: "Sumary metrics for FindInactionTargetUserByCampaignId_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	TestNewCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_TESTNEWCAMPAIGN",
		Help:       "Sumary metrics for TestNewCapmaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	TestNewCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_TESTNEWCAMPAIGN_ERROR",
		Help: "Sumary metrics for TestNewCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AthenaQueryExecutor_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_ATHENAQUERYEXECUTOR",
		Help:       "Sumary metrics for ATHENAQUERYEXECUTOR",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AthenaQuery_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_ATHENAQUERY_ERROR",
		Help: "Sumary metrics for ATHENAQUERY_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	FilterCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_FILTERCAMPAIGN",
		Help:       "Sumary metrics for FILTERCAMPAIGN",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FilterCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_FILTERCAMPAIGN_ERROR",
		Help: "Sumary metrics for FILTERCAMPAIGN_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	TestCampaignById_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_ES_TESTCAMPAIGNBYID",
		Help:       "Summary metrics for TestCampaignById",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	TestCampaignById_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_ES_TESTCAMPAIGNBYID_ERROR",
		Help: "Sumary metrics for TestCampaignById_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetDynamicDataByKey_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_GETDYNAMICDATABYKEY",
		Help:       "Sumary metrics for GetDynamicDataByKey",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	GetDynamicDataByKey_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETDYNAMICDATABYKEY_ERROR",
		Help: "Sumary metrics for GetDynamicDataByKey_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddDynamicData_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ADDDYNAMICDATA",
		Help:       "Sumary metrics for AddDynamicData",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddDynamicData_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ADDDYNAMICDATA_ERROR",
		Help: "Sumary metrics for AddDynamicData_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddDynamicData_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKADDDYNAMICDATA",
		Help:       "Sumary metrics for BulkAddDynamicData",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddDynamicData_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKADDDYNAMICDATA_ERROR",
		Help: "Sumary metrics for BulkAddDynamicData_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	FindQueryCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FindQueryCampaign",
		Help:       "Sumary metrics for FindQueryCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindQueryCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FindQueryCampaign_Error",
		Help: "Sumary metrics for FindQueryCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	AddQueryCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_AddQueryCampaign",
		Help:       "Sumary metrics for AddQueryCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	AddQueryCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_AddQueryCampaign_Error",
		Help: "Sumary metrics for AddQueryCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	BulkAddQueryCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_BULKAddQueryCampaign",
		Help:       "Sumary metrics for BulkAddQueryCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	BulkAddQueryCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_BULKAddQueryCampaign_Error",
		Help: "Sumary metrics for BulkAddQueryCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	ScheduleUserJourneyCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_ScheduleUserJourneyCampaign",
		Help:       "Sumary metrics for ScheduleUserJourneyCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	ScheduleUserJourneyCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_ScheduleUserJourneyCampaign_Error",
		Help: "Sumary metrics for ScheduleUserJourneyCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetActorDetails_Distinct_ActorType_Namespace_Error = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETACTORDETAILS_DISTINCT_ACTORTYPE_Namespace_Error",
		Help: "Sumary metrics for GetActorDetails_Distinct_ActorType_Namespace_Error",
	}, []string{"nservice", "ntype", "campaignname", "actorytype", "channel"})
)

var (
	FindUserJourneyCampaignById_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FindUserJourneyCampaignById",
		Help:       "Sumary metrics for FindUserJourneyCampaignById",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FindUserJourneyCampaignById_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FindUserJourneyCampaignById_Error",
		Help: "Sumary metrics for FindUserJourneyCampaignById_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetActorDetails_Distinct_ActorType_Error = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETACTORDETAILS_DISTINCT_ACTORTYPE_Error",
		Help: "Sumary metrics for GetActorDetails_Distinct_ActorType_Error",
	}, []string{"nservice", "ntype", "campaignname", "actorytype", "channel"})
)

var (
	FilterUserJourneyCampaigns_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_FilterUserJourneyCampaigns",
		Help:       "Sumary metrics for FilterUserJourneyCampaigns",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	FilterUserJourneyCampaigns_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_FilterUserJourneyCampaigns_Error",
		Help: "Sumary metrics for FilterUserJourneyCampaigns_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	Campaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_Campaign",
		Help:       "Summary metrics for Campaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	Campaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_Campaign_Error",
		Help: "Sumary metrics for Campaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	UserJourneyCampaign_Metrics = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "NF_CS_UserJourneyCampaign",
		Help:       "Sumary metrics for UserJourneyCampaign",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"nservice", "nmethod", "ncode"})
)

var (
	UserJourneyCampaign_Error_Metrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_UserJourneyCampaign_Error",
		Help: "Sumary metrics for UserJourneyCampaign_Error",
	}, []string{"nservice", "ntype", "nerror"})
)

var (
	GetActorDetails_Distinct_ActorType_Channel_Error = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "NF_CS_GETACTORDETAILS_DISTINCT_ACTORTYPE_CHANNEL_Error",
		Help: "Sumary metrics for GetActorDetails_Distinct_ActorType_Channel_Error",
	}, []string{"nservice", "ntype", "campaignname", "actorytype", "channel"})
)

func init() {
	Metrics = &metrics.Helper{
		SERVICE_NAME: SERVICE_NAME,
		DATABASE:     DATABASE,
	}
	prometheus.MustRegister(AddCampaign_Metrics)
	prometheus.MustRegister(AddCampaign_Error_Metrics)
	prometheus.MustRegister(BulkAddCampaign_Metrics)
	prometheus.MustRegister(BulkAddCampaign_Error_Metrics)
	prometheus.MustRegister(UpdateCampaign_Metrics)
	prometheus.MustRegister(UpdateCampaign_Error_Metrics)
	prometheus.MustRegister(AddCampaignTemplate_Metrics)
	prometheus.MustRegister(AddCampaignTemplate_Error_Metrics)
	prometheus.MustRegister(BulkAddCampaignTemplate_Metrics)
	prometheus.MustRegister(BulkAddCampaignTemplate_Error_Metrics)
	prometheus.MustRegister(AddNewCampaign_Metrics)
	prometheus.MustRegister(AddNewCampaign_Error_Metrics)
	prometheus.MustRegister(BulkAddNewCampaign_Metrics)
	prometheus.MustRegister(BulkAddNewCampaign_Error_Metrics)
	prometheus.MustRegister(GetUserList_Metrics)
	prometheus.MustRegister(GetUserList_Error_Metrics)
	prometheus.MustRegister(GetTentativeList_Error_Metrics)
	prometheus.MustRegister(FindCampaignById_Metrics)
	prometheus.MustRegister(FindCampaignById_Error_Metrics)
	prometheus.MustRegister(FindCampaignTemplateById_Metrics)
	prometheus.MustRegister(FindCampaignTemplateById_Error_Metrics)
	prometheus.MustRegister(AddTargetUser_Metrics)
	prometheus.MustRegister(AddTargetUser_Error_Metrics)
	prometheus.MustRegister(BulkAddTargetUser_Metrics)
	prometheus.MustRegister(BulkAddTargetUser_Error_Metrics)
	prometheus.MustRegister(FindTargetUserById_Metrics)
	prometheus.MustRegister(FindTargetUserById_Error_Metrics)
	prometheus.MustRegister(AthenaQuery_Metrics)
	prometheus.MustRegister(ATHENAQUERY_Error_Metrics)
	prometheus.MustRegister(ATTRIBUTES_MARSHAL_Error_Metrics)
	prometheus.MustRegister(ATTRIBUTES_UNMARSHAL_Error_Metrics)
	prometheus.MustRegister(AddInactionTargetUser_Metrics)
	prometheus.MustRegister(AddInactionTargetUser_Error_Metrics)
	prometheus.MustRegister(BulkAddInactionTargetUser_Metrics)
	prometheus.MustRegister(BulkAddInactionTargetUser_Error_Metrics)
	prometheus.MustRegister(FindInactionTargetUserByCampaignId_Metrics)
	prometheus.MustRegister(FindInactionTargetUserByCampaignId_Error_Metrics)
	prometheus.MustRegister(TestNewCampaign_Metrics)
	prometheus.MustRegister(TestNewCampaign_Error_Metrics)
	prometheus.MustRegister(AthenaQueryExecutor_Metrics)
	prometheus.MustRegister(AthenaQuery_Error_Metrics)
	prometheus.MustRegister(FilterCampaign_Metrics)
	prometheus.MustRegister(FilterCampaign_Error_Metrics)
	prometheus.MustRegister(TestCampaignById_Metrics)
	prometheus.MustRegister(TestCampaignById_Error_Metrics)
	prometheus.MustRegister(GetDynamicDataByKey_Metrics)
	prometheus.MustRegister(GetDynamicDataByKey_Error_Metrics)
	prometheus.MustRegister(AddDynamicData_Metrics)
	prometheus.MustRegister(AddDynamicData_Error_Metrics)
	prometheus.MustRegister(BulkAddDynamicData_Metrics)
	prometheus.MustRegister(BulkAddDynamicData_Error_Metrics)
	prometheus.MustRegister(FindQueryCampaign_Metrics)
	prometheus.MustRegister(FindQueryCampaign_Error_Metrics)
	prometheus.MustRegister(AddQueryCampaign_Metrics)
	prometheus.MustRegister(AddQueryCampaign_Error_Metrics)
	prometheus.MustRegister(BulkAddQueryCampaign_Metrics)
	prometheus.MustRegister(BulkAddQueryCampaign_Error_Metrics)
	prometheus.MustRegister(ScheduleUserJourneyCampaign_Metrics)
	prometheus.MustRegister(ScheduleUserJourneyCampaign_Error_Metrics)
	prometheus.MustRegister(FindUserJourneyCampaignById_Metrics)
	prometheus.MustRegister(FindUserJourneyCampaignById_Error_Metrics)
	prometheus.MustRegister(FilterUserJourneyCampaigns_Metrics)
	prometheus.MustRegister(FilterUserJourneyCampaigns_Error_Metrics)
	prometheus.MustRegister(Campaign_Metrics)
	prometheus.MustRegister(Campaign_Error_Metrics)
	prometheus.MustRegister(UserJourneyCampaign_Metrics)
	prometheus.MustRegister(UserJourneyCampaign_Error_Metrics)
	prometheus.MustRegister(GetActorDetails_Distinct_ActorType_Namespace_Error)
	prometheus.MustRegister(GetActorDetails_Distinct_ActorType_Error)
	prometheus.MustRegister(GetActorDetails_Distinct_ActorType_Channel_Error)
}

func PushToDistinctActorErrorCounterMetrics() func(*prometheus.CounterVec, context.Context, string, common.ActorType, common.CommunicationChannel) {
	return func(request *prometheus.CounterVec, ctx context.Context, campaignName string, actorType common.ActorType,
		channel common.CommunicationChannel) {
		request.WithLabelValues(SERVICE_NAME, DATABASE, campaignName, actorType.String(), channel.String()).Inc()
	}
}
