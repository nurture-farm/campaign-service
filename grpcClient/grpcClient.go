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

package main

import (
	proto "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	CommunicationEngine "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strings"
)

const (
	LOCAL_URL     = ":7800"
	DEV_URL       = "internal-a2d376e6948514a73a514e3584160823-409853754.ap-south-1.elb.amazonaws.com:80"
	STAGE_URL     = "internal-a49872e2b408d43d28ce6000e575e84b-437549400.ap-south-1.elb.amazonaws.com:80"
	DC2_STAGE_URL = "internal-aa4da3981805144d9950c42874804ed8-2091095273.ap-south-1.elb.amazonaws.com:80"
	PROD_URL      = "internal-a85f3a2f550f84f578459537c8ae3304-113299092.ap-south-1.elb.amazonaws.com:80"
)

var ENV = LOCAL_URL

func main() {

	//TestAddNewCampaignWeather()
	//TestAddNewCampaignMandi()
	//TestAddNewCampaign()
	//TestAddNewCampaignVideoFeed()
	//TestCampaign()
	//TestAddCampaign()
	//TestUpdateCampaign()
	//TestAddInactionTargetUser()
	//TestAddInactionTargetUserBulk()
	//TestParseCampaignQueryForId()
	//TestFindInactionTargetUserByCampaignId()
	//TestAthenaQuery()
	//TestTestNewCampaign()
	//TestFilterCampaign()
	//TestTestCampaignById()
	//TestGetDynamicDataByKey()
	//TestAddDynamicData()
	//TestSchceduleUserJourneyCampaign()
	//TestFindUserJourneyCampaignById()
	//TestFilterUserJourneyCampaign()
	TestUserJourneyCampaign()
	//TestAddDynamicDataBulk()
	//TestFindQueryCampaign()
	//TestAddQueryCampaign()
	//TestAddTargetUser()
	//TestFindTargetUserById()
	//TestTestNewCampaign()
	//Testme()
	//TestAddCampaignTemplate()
}

//	func Testme() {
//		var distinctUserTypes []common.ActorType
//		var campaignquery = "select id from farmer"
//		var err error
//		distinctUserTypes, err = getDistinctUserTypes(context.Background(), campaignQuery, -1)
//
// }
func TestUserJourneyCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.UserJourneyCampaignRequest{
		//Status: common.CampaignStatus_HALTED,
		CampaignId:         689,
		ReferenceId:        "78823d89e-13c0-4372-8f78-ad8a8849734d",
		EngagementVertexId: 105,
	}

	response, err := c.ExecuteUserJourneyCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling indUserJourneyCampaignById: %s", err)
	}
	log.Println(response)
}

func TestFindUserJourneyCampaignById() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindUserJourneyCampaignByIdRequest{
		//Status: common.CampaignStatus_HALTED,
		CampaignId: 368,
	}

	response, err := c.ExecuteFindUserJourneyCampaignById(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling indUserJourneyCampaignById: %s", err)
	}
	log.Println(response)
}

func TestFilterUserJourneyCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FilterUserJourneyCampaignRequest{
		//Status: common.CampaignStatus_HALTED,
		//SearchFilter: "260",
		Namespace:  common.NameSpace_FARM,
		PageNumber: 1,
		Limit:      10,
	}

	response, err := c.ExecuteFilterUserJourneyCampaigns(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling ExecuteFilterUserJourneyCampaigns: %s", err)
	}
	log.Println(response)
}

func TestSchceduleUserJourneyCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.ScheduleUserJourneyCampaignRequest{
		Namespace: common.NameSpace_FARM,
		Name:      "Dev_testing_1",
		CreatedByActor: &common.ActorID{
			ActorId:   1234,
			ActorType: common.ActorType_FARMER,
		},
		//CampaignId: 103,
		Campaign: &proto.UserJourneyCampaign{
			UserJourneys: []*proto.UserJourney{
				{
					UserJourneyVertex: &proto.UserJourneyVertex{
						EventMetadata: &proto.EventMetadata{
							EventName: "HOME_SCREEN_LAUNCHED",
						},
						EventType: common.UserJourneyEventType_ACTION,
						Edge: &proto.UserJourneyEdge{
							WaitTime: &proto.WaitTime{
								WaitFor: "24 HOUR",
							},
							UserJourneyVertex: &proto.UserJourneyVertex{
								EventMetadata: &proto.EventMetadata{
									EventName: "ONBOARDING_LAUNCH_OTP_SCREEN",
								},
								EventType: common.UserJourneyEventType_ACTION,
							},
						},
					},
					Operator: common.LogicalOperator_INTERSECTION,
				},
			},
			EngagementStartVertex: &proto.EngagementVertex{
				CommunicationChannel: common.CommunicationChannel_APP_NOTIFICATION,
				TemplateName:         "farmer_booking_reject",
				Placeholders: []string{
					"first_name",
					"id",
					"last_name",
				},
				ContentMetadata: []*common.Attribs{
					{
						Key:   "version_code",
						Value: "dsfd",
					},
					{
						Key:   "deep_link",
						Value: "some_depp_link_value",
					},
				},
				AthenaQuery: "Select * from afsdb.famrers limit 5",
				Edges: []*proto.EngagementEdge{
					//{
					//	WaitTime: &proto.WaitTime{
					//		WaitFor: "65 MINUTE",
					//	},
					//	States: []common.CommunicationState{
					//		common.CommunicationState_CUSTOMER_DELIVERED,
					//		common.CommunicationState_CUSTOMER_UNDELIVERED,
					//	},
					//	Vertex: &proto.EngagementVertex{
					//		CommunicationChannel: common.CommunicationChannel_APP_NOTIFICATION,
					//		TemplateName:         "farmer_booking_schedule",
					//		ContentMetadata: []*common.Attribs{
					//			{
					//				Key: "version_code",
					//				Value: "dsfd",
					//			},
					//			{
					//				Key: "deep_link",
					//				Value: "some_depp_link_value",
					//			},
					//		},
					//	},
					//},
					{
						WaitTime: &proto.WaitTime{
							//WaitFor: ptypes.DurationProto(time.Minute*15),
							WaitTill: "2022-07-17 13:38:00",
						},
						States: []common.CommunicationState{
							common.CommunicationState_CUSTOMER_DELIVERED,
							common.CommunicationState_CUSTOMER_READ,
							common.CommunicationState_VENDOR_DELIVERED,
							common.CommunicationState_VENDOR_UNDELIVERED,
						},
						Vertex: &proto.EngagementVertex{
							CommunicationChannel: common.CommunicationChannel_APP_NOTIFICATION,
							TemplateName:         "farmer_booking_reject",
							Placeholders: []string{
								"first_name",
								"id",
								"last_name",
							},
							ContentMetadata: []*common.Attribs{
								{
									Key:   "version_code",
									Value: "dsfd",
								},
								{
									Key:   "deep_link",
									Value: "some_depp_link_value",
								},
							},
							AthenaQuery: "Select * from afsdb.famrers limit 10",
						},
					},
				},
			},
		},
		UserJourneyMetadata: "{\"errr\":23}",
		EngagementMetadata:  "{\"dwsds\":124,\"ffdggg\":345}",
		TriggerCampaign:     true,
	}
	response, err := c.ExecuteScheduleUserJourneyCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling ExecuteScheduleUserJourneyCampaign: %s", err)
	}
	log.Println(response)
}

func TestTestCampaignById() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.TestCampaignByIdRequest{
		CampaignId:  378,
		AthenaQuery: "SELECT id from afsdb.farmers where mobile_number IN ('8007678365','8871533105')",
	}

	response, err := c.ExecuteTestCampaignById(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddTestCampaignById: %s", err)
	}
	log.Println(response)
}

func TestFilterCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FilterCampaignRequest{
		PageNumber: 1,
		Limit:      10,
		//Status: common.CampaignStatus_HALTED,
	}

	response, err := c.ExecuteFilterCampaigns(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestTestNewCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.TestNewCampaignRequest{
		//Set your request here
		TestCampaignRequest: &proto.TestCampaignRequest{
			Namespace: common.NameSpace_FARM,
			ContentMetadata: []*common.Attribs{
				{
					Key:   "media_type",
					Value: "IMAGE",
				},
				{
					Key:   "media_access_type",
					Value: "PUBLIC_URL",
				},
				{
					Key:   "image_EN_US",
					Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Eng.jpg",
				},
				{
					Key:   "image_HI_IN",
					Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
				},
				{
					Key:   "image_TE",
					Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Mega+Sale+Banner-+13+to+15+December-04+(Telugu).png",
				},
			},
			//ContentMetadata: []*common.Attribs{
			//	//{
			//	//	Key: "link",
			//	//	Value: "https://nurturefarm.page.link/videoFeed",
			//	//},
			//	//{
			//	//	Key: "cta",
			//	//	Value: "https://nurturefarm.page.link/videoFeed",
			//	//},
			//	{
			//		Key:   "messageCode",
			//		Value: "1006",
			//	},
			//	{
			//		Key:   "click_action",
			//		Value: "FLUTTER_NOTIFICATION_CLICK",
			//	},
			//	//{
			//	//	Key: "image",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
			//	//},
			//	//{
			//	//	Key: "image_EN_US",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Eng.jpg",
			//	//},
			//	//{
			//	//	Key: "image_HI_IN",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
			//	//},
			//	//{
			//	//	Key: "image_GU",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Guj.jpg",
			//	//},
			//	//{
			//	//	Key: "image_PA",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Pun.jpg",
			//	//},
			//	//{
			//	//	Key: "image_KN",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Kan.jpg",
			//	//},
			//	//{
			//	//	Key: "image_TA",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Tam.jpg",
			//	//},
			//	//{
			//	//	Key: "image_TE",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Tel.jpg",
			//	//},
			//	//{
			//	//	Key: "image_BN",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Ban.jpg",
			//	//},
			//	//{
			//	//	Key: "image_MR",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Mar.jpg",
			//	//},
			//	//{
			//	//	Key: "image_ML",
			//	//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Mal.jpg",
			//	//},
			//},
			CommunicationChannel: common.CommunicationChannel_WHATSAPP,
			Type:                 common.CampaignQueryType_ATHENA,
			//Query:                "SELECT user_type FROM campaignservicedb.target_users limit 10",
			Query: "select id,'Saaransh' as farmer_first_name, '100' as points_earned from stage_afsdb.farmers where mobile_number in ('8247356888', '6191241553')",
			//Query: "select fo_mobile_number as mobile_number,'FIELD_OFFICER' as user_type, min_time_date, machine_number, operator_full_name, operator_mobile_number, operator_acres, system_acres from (select machine_number, date(from_unixtime(min_time/1000000) + interval '330' minute) as min_time_date, sum(area_by_convexhull) as system_acres from iot_gateway.machine_summary_ro where area_by_convexhull >= 1 and current_date - interval '1' day = date(from_unixtime(min_time/1000000) + interval '330' minute) group by machine_number, date(from_unixtime(min_time/1000000) + interval '330' minute))a left join (select machine_id, date(activity_time) as activity_time , sum(cast(input_quantity as double)) as operator_acres, operator_id from afsdb.booking_spraying_data where current_date - interval '1' day = date(activity_time) and service_type = 'SPRAYING' group by machine_id, date(activity_time), operator_id) b on b.machine_id = a.machine_number and a.min_time_date = b.activity_time left join (select id as operator_id, full_name as operator_full_name, mobile_number as operator_mobile_number, reporting_to from afsdb.operator)x on b.operator_id = x.operator_id join (select full_name as fo_full_name , mobile_number as fo_mobile_number ,id as fo_id from afsdb.field_officer)y on x.reporting_to = y.fo_id where a.system_acres is not null and (a.system_acres > 1.2*b.operator_acres or operator_acres is null) and fo_mobile_number = '9284151797'",
			//ChannelAttributes: &CommunicationEngine.CommunicationChannelAttributes{
			//	PushNotificationType: common.PushNotificationType_NO_PUSH_NOTIFICATION_TYPE,
			//},
			Media: &CommunicationEngine.Media{
				MediaType:       common.MediaType_IMAGE,
				MediaAccessType: common.MediaAccessType_PUBLIC_URL,
			},
		},
		TestCampaignTemplateRequests: []*proto.TestCampaignTemplateRequest{
			{
				TemplateName:        "farmer_registration",
				DistributionPercent: 100,
			},
		},
		//TestTargetUserRequests: []*proto.TestTargetUserRequest{
		//	{
		//		User: &common.ActorID{
		//			ActorId:   6646919,
		//			ActorType: common.ActorType_FARMER,
		//		},
		//		//Attribs: []*common.Attribs{
		//		//	{
		//		//		Key:   "farmer_first_name",
		//		//		Value: "saransh",
		//		//	},
		//		//	{
		//		//		Key:   "points_earned",
		//		//		Value: "200",
		//		//	},
		//		//},
		//	},
		//	//{
		//	//	User: &common.ActorID{
		//	//		ActorId:   803777,
		//	//		ActorType: common.ActorType_FARMER,
		//	//	},
		//	//	Attribs: []*common.Attribs{
		//	//		{
		//	//			Key:   "some_key_3",
		//	//			Value: "some_value_3",
		//	//		},
		//	//		{
		//	//			Key:   "some_key_4",
		//	//			Value: "some_value_4",
		//	//		},
		//	//	},
		//	//},
		//},
	}

	response, err := c.ExecuteTestNewCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAthenaQuery() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AthenaQueryRequest{
		AthenaQuery: "select id from afsdb.farmers where mobile_number = '9453849441'",
	}

	response, err := c.ExecuteAthenaQuery(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddCampaignRequest: %s", err)
	}
	log.Println(response)

}

func TestParseCampaignQueryForId() {
	campaignQuery := "SELECT id, farmer_name, state from afsdb.farmers where mobile_number ='9971181761'"
	var newQuery string
	var secondIterationStartIndex int
	log.Println(strings.Fields(campaignQuery))
	words := strings.Fields(strings.ToLower(campaignQuery))

	for i, word := range words {
		if strings.Contains(word, "from") {
			secondIterationStartIndex = i
			break
		}
		if strings.Contains(word, "id") {
			newQuery += strings.ReplaceAll(word, ",", " ")
			log.Println(newQuery)
		} else if strings.Contains(word, "select") {
			newQuery += word
		} else {
			continue
		}
		newQuery += " "
	}

	for i := secondIterationStartIndex; i < len(words); i++ {
		newQuery += words[i]
		newQuery += " "
	}

	log.Println(newQuery)
}

func TestAddCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddCampaignRequest{
		//Set your request here
		Namespace:      common.NameSpace_FARM,
		Name:           "testing_multimedia",
		Description:    "Testing_campaign",
		CronExpression: "*/1 * * * *",
		Occurrences:    12,
		ContentMetadata: []*common.Attribs{
			{
				Key:   "image_EN_US",
				Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Steal_the_deal.png",
			},
			{
				Key:   "image_HI_IN",
				Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
			},
			{
				Key:   "image_GU",
				Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Guj.jpg",
			},
		},
		CommunicationChannel: common.CommunicationChannel_WHATSAPP,
		Type:                 common.CampaignQueryType_ATHENA,
		Query:                "SELECT id from afsdb.farmers where mobile_number ='7037033632'",
		CreatedByActor: &common.ActorID{
			ActorId:   413428,
			ActorType: common.ActorType_FARMER,
		},
		Media: &CommunicationEngine.Media{
			MediaType:       common.MediaType_IMAGE,
			MediaAccessType: common.MediaAccessType_PUBLIC_URL,
		},
	}

	response, err := c.ExecuteAddCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestUpdateCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.UpdateCampaignRequest{
		//Set your request here
		Id: 8,
		UpdatedByActor: &common.ActorID{
			ActorId:   413428,
			ActorType: common.ActorType_FARMER,
		},
		AddCampaignRequest: &proto.AddCampaignRequest{
			Status:         common.CampaignStatus_RUNNING,
			Query:          "select profile_id as id from afsdb.farmer_users fu join afsdb.farmer_app_account fa on fu.login = fa.primary_mobile_number where lower(state) in ('rajasthan', 'haryana')",
			CronExpression: "*/5 * * * *",
		},
		AddCampaignTemplateRequests: []*proto.AddCampaignTemplateRequest{
			{
				TemplateName:        "videofeed_campaign_push_notification",
				CampaignName:        "video_feed_campaign",
				DistributionPercent: 65,
			},
			{
				TemplateName:        "some_template_name234",
				CampaignName:        "dwsfsfsf",
				DistributionPercent: 35,
			},
		},
	}

	response, err := c.ExecuteUpdateCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling UpdateCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddCampaignTemplate() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddCampaignTemplateRequest{
		//Set your request here
		CampaignId:          382,
		TemplateName:        "some_template_name",
		CampaignName:        "test_multime_whatsapp",
		DistributionPercent: 100,
	}

	response, err := c.ExecuteAddCampaignTemplate(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddCampaignTemplateRequest: %s", err)
	}
	log.Println(response)
}

func TestAddNewCampaignVideoFeed() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddNewCampaignRequest{
		//Set your request here
		AddCampaignRequest: &proto.AddCampaignRequest{
			Namespace:      common.NameSpace_FARM,
			Name:           "video_feed_pn_campaign",
			Description:    "pn campaign for the video feed in farm app",
			CronExpression: "14 16 * * *",
			//Occurrences: 1, //add cron library parsing to add ocurrence
			ContentMetadata: []*common.Attribs{
				{
					Key:   "link",
					Value: "https://nrf.page.link/videoFeed",
				},
				{
					Key:   "cta",
					Value: "showDetails?requestUrl=video-feed&module=video_feed_page&titleKey=title.video.feed&source=pn",
				},
				{
					Key:   "messageCode",
					Value: "1006",
				},
				{
					Key:   "click_action",
					Value: "FLUTTER_NOTIFICATION_CLICK",
				},
			},
			CommunicationChannel: common.CommunicationChannel_APP_NOTIFICATION,
			Type:                 common.CampaignQueryType_ATHENA,
			Query:                "SELECT distinct(f.id) FROM afsdb.farmer_mobile_device_details fmdd JOIN afsdb.farmer_app_account faa ON faa.id = fmdd.farmer_user_id AND faa.active =1 AND fmdd .active =1 join afsdb.farmers f ON f.mobile_number = faa.primary_mobile_number AND f.active =1 AND faa.active = 1 WHERE fmdd.version_code >= 58 ",
			CreatedByActor: &common.ActorID{
				ActorId:   413428,
				ActorType: common.ActorType_FARMER,
			},
			ChannelAttributes: &CommunicationEngine.CommunicationChannelAttributes{
				PushNotificationType: common.PushNotificationType_NOTIFICATION,
			},
			//Media: &ce.Media{
			//	MediaType: common.MediaType_VIDEO,
			//	MediaAccessType: common.MediaAccessType_PUBLIC_URL,
			//	MediaInfo: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Vaccine+Magnetic+Powers+%23Shorts.mp4",
			//},
		},
		AddCampaignTemplateRequests: []*proto.AddCampaignTemplateRequest{
			{
				TemplateName:        "videofeed_campaign_push_notification",
				CampaignName:        "video_feed_campaign",
				DistributionPercent: 100,
			},
			//{
			//	TemplateName: "some_template_name234",
			//	CampaignName: "dwsfsfsf",
			//	DistributionPercent: 35,
			//},
		},
		//AddTargetUserRequests: []*proto.AddTargetUserRequest{
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.ActorType_FARMER,
		//		},
		//	},
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.ActorType_RETAILER,
		//		},
		//	},
		//},
	}

	response, err := c.ExecuteAddNewCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddNewCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddNewCampaignRequest{
		//Set your request here
		AddCampaignRequest: &proto.AddCampaignRequest{
			Namespace:      common.NameSpace_FARM,
			Name:           "test_whatsapp2",
			Description:    "PN campaign for meg",
			CronExpression: "20 05 * * *",
			Occurrences:    1, //add cron library parsing to add occurrence
			//ContentMetadata: []*common.Attribs{
			//{
			//	Key: "link",
			//	Value: "nurtureretail://sales-enquiry",
			//},
			//{
			//	Key: "cta",
			//	Value: "nurtureretail://sales-enquiry",
			//},
			//{
			//	Key:   "messageCode",
			//	Value: "1006",
			//},
			//{
			//	Key:   "click_action",
			//	Value: "FLUTTER_NOTIFICATION_CLICK",
			//},
			//{
			//	Key:   "image",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/mega_sale_telugu.png",
			//},
			//{
			//	Key: "image_EN_US",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Steal_the_deal.png",
			//},
			//{
			//	Key: "image_HI_IN",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
			//},
			//{
			//	Key: "image_GU",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Guj.jpg",
			//},
			//{
			//	Key: "image_PA",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Pun.jpg",
			//},
			//{
			//	Key: "image_KN",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Kan.jpg",
			//},
			//{
			//	Key: "image_TA",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Tam.jpg",
			//},
			//{
			//	Key:   "image_TE",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/mega_sale_telugu.png",
			//},
			//{
			//	Key: "image_BN",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Ban.jpg",
			//},
			//{
			//	Key: "image_MR",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Mar.jpg",
			//},
			//{
			//	Key: "image_ML",
			//	Value: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Mal.jpg",
			//},
			//},
			CommunicationChannel: common.CommunicationChannel_WHATSAPP,
			Type:                 common.CampaignQueryType_ATHENA,
			CampaignScheduleType: common.CampaignScheduleType_NO_CAMPAIGN_SCHEDULE_TYPE,
			Query:                "select id from afsdb.farmers where mobile_number = '9993098893'",
			//InactionDuration: ptypes.DurationProto(3*time.Minute),
			CreatedByActor: &common.ActorID{
				ActorId:   413428,
				ActorType: common.ActorType_FARMER,
			},
			//ChannelAttributes: &CommunicationEngine.CommunicationChannelAttributes{
			//	PushNotificationType: common.PushNotificationType_NO_PUSH_NOTIFICATION_TYPE,
			//},
			//Media: &CommunicationEngine.Media{
			//	MediaType: common.MediaType_IMAGE,
			//	MediaAccessType: common.MediaAccessType_DOCUMENT_ID,
			//	MediaInfo: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Rewards_for+post.jpeg",
			//},
		},
		AddCampaignTemplateRequests: []*proto.AddCampaignTemplateRequest{
			{
				TemplateName:        "farmer_booking_reject",
				CampaignName:        "farmer_booking_reject",
				DistributionPercent: 100,
			},
			//{
			//	TemplateName: "some_template_name234",
			//	CampaignName: "dwsfsfsf",
			//	DistributionPercent: 35,
			//},
		},
		//AddTargetUserRequests: []*proto.AddTargetUserRequest{
		//	{
		//		User: &common.ActorID{
		//			ActorId:   413428,
		//			ActorType: common.ActorType_FARMER,
		//		},
		//		Attribs: []*common.Attribs{
		//			{
		//				Key:   "Name",
		//				Value: "Kishan",
		//			},
		//			{
		//				Key:   "farmer_first_name",
		//				Value: "Kishan123",
		//			},
		//		},
		//	},
		//	{
		//		User: &common.ActorID{
		//			ActorId:   803777,
		//			ActorType: common.ActorType_FARMER,
		//		},
		//		Attribs: []*common.Attribs{
		//			{
		//				Key:   "Name",
		//				Value: "Kishan2",
		//			},
		//			{
		//				Key:   "farmer_first_name",
		//				Value: "Kishan456",
		//			},
		//		},
		//	},
		//},
	}

	response, err := c.ExecuteAddNewCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddNewCampaignMandi() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddNewCampaignRequest{
		//Set your request here
		AddCampaignRequest: &proto.AddCampaignRequest{
			Namespace:      common.NameSpace_FARM,
			Name:           "mandi_pn_campaign",
			Description:    "pn campaign for the mandi in farm app",
			CronExpression: "03 14 * * *",
			Occurrences:    1, //add cron library parsing to add ocurrence
			ContentMetadata: []*common.Attribs{
				{
					Key:   "link",
					Value: "https://nurturefarm.page.link/viewMandis",
				},
				{
					Key:   "cta",
					Value: "mandi/v2",
				},
				{
					Key:   "messageCode",
					Value: "1006",
				},
				{
					Key:   "click_action",
					Value: "FLUTTER_NOTIFICATION_CLICK",
				},
			},
			CommunicationChannel: common.CommunicationChannel_APP_NOTIFICATION,
			Type:                 common.CampaignQueryType_ATHENA,
			Query:                "SELECT distinct(f.id) FROM afsdb.farmer_mobile_device_details fmdd JOIN afsdb.farmer_app_account faa ON faa.id = fmdd.farmer_user_id AND faa.active =1 AND fmdd .active =1 join afsdb.farmers f ON f.mobile_number = faa.primary_mobile_number AND f.active =1 AND faa.active = 1 WHERE fmdd.version_code >= 0 ",
			CreatedByActor: &common.ActorID{
				ActorId:   413428,
				ActorType: common.ActorType_FARMER,
			},
			ChannelAttributes: &CommunicationEngine.CommunicationChannelAttributes{
				PushNotificationType: common.PushNotificationType_NOTIFICATION,
			},
			//Media: &ce.Media{
			//	MediaType: common.MediaType_VIDEO,
			//	MediaAccessType: common.MediaAccessType_PUBLIC_URL,
			//	MediaInfo: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Vaccine+Magnetic+Powers+%23Shorts.mp4",
			//},
		},
		AddCampaignTemplateRequests: []*proto.AddCampaignTemplateRequest{
			{
				TemplateName:        "mandi_campaign_push_notification",
				CampaignName:        "mandi_campaign",
				DistributionPercent: 100,
			},
			//{
			//	TemplateName: "some_template_name234",
			//	CampaignName: "dwsfsfsf",
			//	DistributionPercent: 35,
			//},
		},
		//AddTargetUserRequests: []*proto.AddTargetUserRequest{
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.
		//			ActorType_FARMER,
		//		},
		//	},
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.ActorType_RETAILER,
		//		},
		//	},
		//},
	}

	response, err := c.ExecuteAddNewCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddNewCampaignWeather() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddNewCampaignRequest{
		//Set your request here
		AddCampaignRequest: &proto.AddCampaignRequest{
			Namespace:      common.NameSpace_FARM,
			Name:           "test_pn_campaign",
			Description:    "test pn campaign for the app",
			CronExpression: "*/2 * * * *",
			//Occurrences: 1, //add cron library parsing to add ocurrence
			ContentMetadata: []*common.Attribs{
				{
					Key:   "link",
					Value: "https://nrf.page.link/weather",
				},
				{
					Key:   "cta",
					Value: "/showDetails?requestUrl=weather-detail&module=weather_detail_page&title=Today's Weather&source=dashboard.dailyUpdate",
				},
				{
					Key:   "messageCode",
					Value: "1006",
				},
				{
					Key:   "click_action",
					Value: "FLUTTER_NOTIFICATION_CLICK",
				},
			},
			CommunicationChannel: common.CommunicationChannel_SMS,
			Type:                 common.CampaignQueryType_ATHENA,
			Query:                "SELECT id from afsdb.farmers where mobile_number ='8007678365'",
			CreatedByActor: &common.ActorID{
				ActorId:   413428,
				ActorType: common.ActorType_FARMER,
			},
			ChannelAttributes: &CommunicationEngine.CommunicationChannelAttributes{
				PushNotificationType: common.PushNotificationType_NO_PUSH_NOTIFICATION_TYPE,
			},
			//Media: &ce.Media{
			//	MediaType: common.MediaType_VIDEO,
			//	MediaAccessType: common.MediaAccessType_PUBLIC_URL,
			//	MediaInfo: "https://afs-static-content.s3.ap-south-1.amazonaws.com/push_notification_campaign/Vaccine+Magnetic+Powers+%23Shorts.mp4",
			//},
		},
		AddCampaignTemplateRequests: []*proto.AddCampaignTemplateRequest{
			{
				TemplateName:        "farmer_booking_reject",
				CampaignName:        "test_campaign",
				DistributionPercent: 100,
			},
			//{
			//	TemplateName: "some_template_name234",
			//	CampaignName: "dwsfsfsf",
			//	DistributionPercent: 35,
			//},
		},
		//AddTargetUserRequests: []*proto.AddTargetUserRequest{
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.ActorType_FARMER,
		//		},
		//	},
		//	{
		//		User: &common.ActorID{
		//			ActorId: 413428,
		//			ActorType: common.ActorType_RETAILER,
		//		},
		//	},
		//},
	}

	response, err := c.ExecuteAddNewCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddNewCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)
	request := &proto.CampaignRequest{
		CampaignId: 382,
	}
	response, err := c.ExecuteCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling CampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestFindCampaignById() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindCampaignByIdRequest{
		Id: 103,
	}

	response, err := c.ExecuteFindCampaignById(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindCampaignByIdRequest: %s", err)
	}
	log.Println(response)
}

func TestFindCampaignTemplateById() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindCampaignTemplateByIdRequest{
		//Set your request here
	}

	response, err := c.ExecuteFindCampaignTemplateById(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindCampaignTemplateByIdRequest: %s", err)
	}
	log.Println(response)
}

func TestAddTargetUser() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddTargetUserRequest{
		//Set your request here
		CampaignId: 143,
		User: &common.ActorID{
			ActorId:   23443,
			ActorType: common.ActorType_RETAILER,
		},
		Attribs: []*common.Attribs{
			{
				Key:   "farmer_first_name",
				Value: "Kishan",
			},
		},
	}

	response, err := c.ExecuteAddTargetUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddTargetUserRequest: %s", err)
	}
	log.Println(response)
}

func TestAddInactionTargetUser() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddInactionTargetUserRequest{
		//Set your request here
		CampaignId: 1,
		User: &common.ActorID{
			ActorId:   23443,
			ActorType: common.ActorType_RETAILER,
		},
	}

	response, err := c.ExecuteAddInactionTargetUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddInactionTargetUserRequest: %s", err)
	}
	log.Println(response)
}

func TestAddInactionTargetUserBulk() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)
	req1 := &proto.AddInactionTargetUserRequest{
		CampaignId: 5,
		User: &common.ActorID{
			ActorId:   23446,
			ActorType: common.ActorType_FARMER,
		},
	}
	req2 := &proto.AddInactionTargetUserRequest{
		CampaignId: 4,
		User: &common.ActorID{
			ActorId:   23445,
			ActorType: common.ActorType_FARMER,
		},
	}
	request := &proto.BulkAddInactionTargetUserRequest{}
	request.Requests = append(request.Requests, req1)
	request.Requests = append(request.Requests, req2)
	//bulkRequest := &proto.BulkAddInactionTargetUserRequest{
	//	Requests: AddInactionTargetUserRequests,
	//}

	response, err := c.ExecuteAddInactionTargetUserBulk(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddInactionTargetUserRequestBulk: %s", err)
	}
	log.Println(response)
}

func TestFindInactionTargetUserByCampaignId() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindInactionTargetUserByCampaignIdRequest{
		//Set your request here
		CampaignId: 35,
	}

	response, err := c.ExecuteFindInactionTargetUserByCampaignId(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindInactionTargetUserByCampaignIdRequest: %s", err)
	}
	log.Println(response)
}

func TestGetDynamicDataByKey() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.GetDynamicDataByKeyRequest{
		//Set your request here
		CampaignId: 254,
		DynamicKey: "2022.02.17",
	}

	response, err := c.ExecuteGetDynamicDataByKey(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling GetDynamicDataByKeyRequest: %s", err)
	}
	log.Println(response)
}

func TestAddDynamicData() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddDynamicDataRequest{
		CampaignId: 1,
		DynamicKey: "2022.03.03",
		CtaLink:    "cta_link_data",
		Media:      "{ \"image\" : \"image_link\"}",
		//Set your request here
	}

	response, err := c.ExecuteAddDynamicData(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddDynamicDataRequest: %s", err)
	}
	log.Println(response)
}

func TestAddDynamicDataBulk() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)
	req1 := &proto.AddDynamicDataRequest{
		CampaignId: 257,
		DynamicKey: "2022.03.02",
		CtaLink:    "cta_link_data_bulk_1",
		Media:      "{ \"image\" : \"image_link_bulk_1\"}",
		//Set your request here
	}
	req2 := &proto.AddDynamicDataRequest{
		CampaignId: 2,
		DynamicKey: "2022.03.02",
		CtaLink:    "cta_link_data_bulk_2",
		Media:      "{ \"image\" : \"image_link_bulk_2\"}",
		//Set your request here
	}
	request := &proto.BulkAddDynamicDataRequest{}
	request.Requests = append(request.Requests, req1)
	request.Requests = append(request.Requests, req2)

	response, err := c.ExecuteAddDynamicDataBulk(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling BulkAddDynamicDataRequest: %s", err)
	}
	log.Println(response)
}

func TestFindQueryCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindQueryCampaignRequest{
		//Set your request here
		Type: "BUSINESS_DEFINED",
	}

	response, err := c.ExecuteFindQueryCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling FindQueryCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddQueryCampaign() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.AddQueryCampaignRequest{
		//Set your request here
		Name:      "TEST_QUERY_2",
		Query:     "SELECT id from afsdb.farmer where mobile_number = '9993098893'",
		Type:      "BUSINESS_DEFINED",
		UpdatedBy: "honey@nurture.farm",
	}

	response, err := c.ExecuteAddQueryCampaign(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling AddQueryCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestAddQueryCampaignBulk() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.BulkAddQueryCampaignRequest{
		//Set your request here

	}

	response, err := c.ExecuteAddQueryCampaignBulk(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling BulkAddQueryCampaignRequest: %s", err)
	}
	log.Println(response)
}

func TestFindTargetUserById() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ENV, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewCampaignServiceClient(conn)

	request := &proto.FindTargetUserByIdRequest{
		//Set your request here
		CampaignId: 143,
	}

	response, err := c.ExecuteFindTargetUserById(context.Background(), request)
	if err != nil {
		log.Fatalf("Error when calling TestFindTargetUserById: %s", err)
	}
	log.Println(response)
}
