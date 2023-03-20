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
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	ce "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"code.nurture.farm/platform/CampaignService/core/golang/cache"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"math"
	"strings"
	"time"
)

var namespaceByActorTypeMap = make(map[common.NameSpace]int32)

func GetActorTypeUsingNamespace(namespace common.NameSpace) int32 {
	if len(namespaceByActorTypeMap) == 0 {
		initNamespaceByActorTypeMap()
	}
	return namespaceByActorTypeMap[namespace]
}

func initNamespaceByActorTypeMap() {
	namespaceByActorTypeMap[common.NameSpace_FARM] = common.ActorType_value["FARMER"]
	namespaceByActorTypeMap[common.NameSpace_RETAIL] = common.ActorType_value["RETAILER"]
	namespaceByActorTypeMap[common.NameSpace_NURTURE_PARTNER] = common.ActorType_value["PARTNER"]
}

func GetInQueryWithValues(arr []int64, query string) string {
	ids := ""
	for i, id := range arr {
		ids = ids + cast.ToString(id)
		if i != len(arr)-1 {
			ids = ids + ","
		}
	}
	ret := strings.ReplaceAll(query, "@query@", ids)
	return ret
}

func GetIndexInRow(row []string, fieldName string) int {
	for k, v := range row {
		if v == fieldName {
			return k
		}
	}
	return -1
}

func GetStringValue(index int, arr []string) string {
	if index == -1 {
		return ""
	}
	return arr[index]
}

func GetInt64Value(index int, arr []string) int64 {
	if index == -1 {
		return 0
	}
	return cast.ToInt64(arr[index])
}

func GetWorkflowId(campaignId int64) string {
	return fmt.Sprintf("%s_%d", "CRON_ExecuteCampaignWorkflow", campaignId)
}
func GetUserJourneyWorkflowId(campaignId int64, engagementVertexId int64) string {
	return fmt.Sprintf("%s_%d_%d", "CRON_ExecuteUserJourneyCampaignWorkflow", campaignId, engagementVertexId)
}

func GetUserJourneyWorkflowParentReferenceId(campaignId int64, engagementVertexId int64) string {
	return fmt.Sprintf("%s_%d_%d", "CRON_UserJourneyCampaignWorkflow", campaignId, engagementVertexId)
}

type UserDetail struct {
	Id               int64
	LangKey          string
	MobileNumber     string
	Email            string
	PushToken        string
	ModifiedOn       int64
	AppId            string
	ActorType        string
	SecondaryLangKey string
	Location         string
}

func OutputLog(log string, logLevel string, error error, request *fs.AddNewCampaignRequest) {
	if logLevel == CONST_LOG_LEVEL_INFO {
		if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
			logger.Info(log, zap.Any("addCampaignRequest", request.AddCampaignRequest))
		} else {
			logger.Info(log, zap.Any("request", request))
		}
	} else if logLevel == CONST_LOG_LEVEL_ERRROR {
		if request.AddTargetUserRequests != nil && len(request.AddTargetUserRequests) > 10 {
			logger.Error(log, zap.Error(error), zap.Any("addCampaignRequest", request.AddCampaignRequest))
		} else {
			logger.Error(log, zap.Error(error), zap.Any("request", request))
		}
	}
	return
}

func OutputQueryLog(log string, logLevel string, error error, query string) {

	if logLevel == CONST_LOG_LEVEL_INFO {
		if len(query) > 5000 {
			logger.Info(log)
		} else {
			logger.Info(log, zap.Any("query", query))
		}
	} else if logLevel == CONST_LOG_LEVEL_ERRROR {
		if len(query) > 5000 {
			logger.Error(log, zap.Error(error))
		} else {
			logger.Error(log, zap.Error(error), zap.Any("request", query))
		}
	}
	return
}

func getActorDetailsData(ctx context.Context, logString string, campaignQueryType common.CampaignQueryType,
	channel common.CommunicationChannel, campaignQuery string, targetUsers []*common.ActorID, campaignId int64) (actorDetailsData [][]string, err error) {

	if campaignQueryType == common.CampaignQueryType_ATHENA || campaignQueryType == common.CampaignQueryType_USER_JOURNEY {
		_, actorDetailsData, err = executeGetUserAthenaQuery(ctx, strings.ReplaceAll(getCampaignOuterQuery(channel, campaignQueryType), "@query@", campaignQuery), channel, campaignId)
		if err != nil {
			logger.Error(logString+", Error in executing Athena Query", zap.Any("Query", campaignQuery), zap.Any("error", err))
			err = fmt.Errorf("ATHENA_QUERY_ERROR")
			return nil, err
		}
	} else if campaignQueryType == common.CampaignQueryType_USER_LIST {

		actorIds := []int64{}
		for _, actor := range targetUsers {
			actorIds = append(actorIds, actor.ActorId)
		}
		var actorFarmerDetailsData [][]string
		_, actorFarmerDetailsData, err := executeGetUserAthenaQuery(ctx,
			GetInQueryWithValues(actorIds, getCampaignOuterQuery(channel, campaignQueryType)), channel, campaignId)
		if err != nil {
			logger.Error(logString+", Error in executing Athena QUERY for Userlist campaign",
				zap.Any("Query", campaignQuery), zap.Any("error", err))
			err = fmt.Errorf("USERLIST_ATHENA_QUERY_ERROR")
			return nil, err
		}
		actorDetailsData = append(actorDetailsData, actorFarmerDetailsData...)
	}
	return actorDetailsData, nil
}

func getActorDetailsMap(actorDetailsData [][]string) (map[int64]UserDetail, map[int64][]*ce.Placeholder) {

	var actorByPlaceholderMap = make(map[int64][]*ce.Placeholder)
	var indexFieldMap = make(map[int64]string)

	actorDetailMap := make(map[int64]UserDetail)
	var idIndex, langKeyIndex, mobileNumberIndex, emailIndex, pushTokenIndex, modifiedOnIndex, appIdIndex, actorTypeIndex, locationIndex int
	for i, val := range actorDetailsData {
		if i == 0 {
			idIndex = GetIndexInRow(val, CONST_ID)
			langKeyIndex = GetIndexInRow(val, CONST_LANG_KEY)
			mobileNumberIndex = GetIndexInRow(val, CONST_MOBILE_NUMBER)
			emailIndex = GetIndexInRow(val, CONST_EMAIL_ID)
			pushTokenIndex = GetIndexInRow(val, CONST_FCM_TOKEN)
			modifiedOnIndex = GetIndexInRow(val, CONST_MODIFIED_ON)
			appIdIndex = GetIndexInRow(val, CONST_APP_ID)
			actorTypeIndex = GetIndexInRow(val, CONST_ACTOR_TYPE)
			locationIndex = GetIndexInRow(val, CONST_LOCATION)
			indexFieldMap = GetPlaceholderIndexFieldMap(val)
			continue
		}
		actorId := GetInt64Value(idIndex, val)
		langugageValue := GetStringValue(langKeyIndex, val)
		if langugageValue == "en" {
			langugageValue = "en-us"
		}
		userDetail := UserDetail{
			Id:           actorId,
			LangKey:      langugageValue,
			MobileNumber: GetStringValue(mobileNumberIndex, val),
			Email:        GetStringValue(emailIndex, val),
			PushToken:    GetStringValue(pushTokenIndex, val),
			ModifiedOn:   GetInt64Value(modifiedOnIndex, val),
			AppId:        GetStringValue(appIdIndex, val),
			ActorType:    GetStringValue(actorTypeIndex, val),
			Location:     GetStringValue(locationIndex, val),
		}
		_, ok := actorByPlaceholderMap[actorId]
		if !ok {
			actorByPlaceholderMap[actorId] = append(actorByPlaceholderMap[actorId], GetPlaceholderData(val, indexFieldMap)...)
		}
		if userPreviousDetail, ok := actorDetailMap[userDetail.Id]; !ok {
			actorDetailMap[userDetail.Id] = userDetail
		} else if userPreviousDetail.ModifiedOn < userDetail.ModifiedOn {
			actorDetailMap[userDetail.Id] = userDetail
		}
	}
	return actorDetailMap, actorByPlaceholderMap
}

func getActorDetails(actorDetailMap map[int64]UserDetail, channel common.CommunicationChannel, campaignId int64) (actorDetails []*ce.ActorDetails, actorIDDetails []*common.ActorID) {

	for _, actorDetail := range actorDetailMap {
		ceActorDetail := &ce.ActorDetails{
			EmailId:               actorDetail.Email,
			MobileNumber:          actorDetail.MobileNumber,
			LanguageCode:          common.LanguageCode(common.LanguageCode_value[strings.ToUpper(strings.ReplaceAll(actorDetail.LangKey, "-", "_"))]),
			SecondaryLanguageCode: common.LanguageCode(common.LanguageCode_value[strings.ToUpper(strings.ReplaceAll(actorDetail.SecondaryLangKey, "-", "_"))]),
			FcmToken:              actorDetail.PushToken,
			AppId:                 common.AppID(common.AppID_value[actorDetail.AppId]),
			AppType:               common.AppType_ANDROID,
		}
		if common.CommunicationChannel_EMAIL == channel && ceActorDetail.EmailId == CONST_ATHENA_NIL {
			logger.Info("Email not available for user", zap.Any("ActorDetail", actorDetail), zap.Any("campaignId", campaignId))
			continue
		}
		if (common.CommunicationChannel_SMS == channel || common.CommunicationChannel_WHATSAPP ==
			channel) && actorDetail.MobileNumber == CONST_ATHENA_NIL {
			logger.Info("Mobile number not available for user", zap.Any("ActorDetail", actorDetail), zap.Any("campaignId", campaignId))
			continue
		}
		if common.CommunicationChannel_APP_NOTIFICATION == channel && ceActorDetail.FcmToken == CONST_ATHENA_NIL {
			logger.Info("FCM token not available for user", zap.Any("ActorDetail", actorDetail), zap.Any("campaignId", campaignId))
			continue
		}
		actorDetails = append(actorDetails, ceActorDetail)
		actorIDDetails = append(actorIDDetails, &common.ActorID{
			ActorType: common.ActorType(common.ActorType_value[actorDetail.ActorType]),
			ActorId:   actorDetail.Id,
		})
	}
	return actorDetails, actorIDDetails
}

func getContentMetaDataAndImageMap(campaignContentMetaData []*common.Attribs) (contentMetaData []*ce.Placeholder, imageMap map[common.LanguageCode]string) {

	imageMap = make(map[common.LanguageCode]string)
	for _, attrib := range campaignContentMetaData {
		attribKey := attrib.Key
		if strings.HasPrefix(attribKey, CONST_IMAGE) {
			if attribKey == CONST_IMAGE {
				imageMap[common.LanguageCode_NO_LANGUAGE_CODE] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_EN_US.String() {
				imageMap[common.LanguageCode_EN_US] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_HI_IN.String() {
				imageMap[common.LanguageCode_HI_IN] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_GU.String() {
				imageMap[common.LanguageCode_GU] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_PA.String() {
				imageMap[common.LanguageCode_PA] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_KA.String() {
				imageMap[common.LanguageCode_KA] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_TA.String() {
				imageMap[common.LanguageCode_TA] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_TE.String() {
				imageMap[common.LanguageCode_TE] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_BN.String() {
				imageMap[common.LanguageCode_BN] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_MR.String() {
				imageMap[common.LanguageCode_MR] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_ML.String() {
				imageMap[common.LanguageCode_ML] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_KN.String() {
				imageMap[common.LanguageCode_KN] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_OD.String() {
				imageMap[common.LanguageCode_OD] = attrib.Value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_PU.String() {
				imageMap[common.LanguageCode_PU] = attrib.Value
			}
			continue
		}
		placeHolder := &ce.Placeholder{
			Key:   attrib.Key,
			Value: attrib.Value,
		}
		contentMetaData = append(contentMetaData, placeHolder)
		if attrib.Key == "cta" {
			contentMetaData = append(contentMetaData, &ce.Placeholder{
				Key:   "message",
				Value: attrib.Value,
			})
			contentMetaData = append(contentMetaData, &ce.Placeholder{
				Key:   "link",
				Value: attrib.Value,
			})
		}
	}
	return contentMetaData, imageMap
}

func sendCommunicationEvent(ctx context.Context, logString string, events []*ce.CommunicationEvent, campaingId int64) error {

	iterations := len(events) / 1000
	if len(events)%1000 > 0 {
		iterations++
	}
	logger.Info(logString+", total number of events", zap.Any("numEvents", len(events)),
		zap.Any("campaignId", campaingId))

	if len(events) == 0 {
		err := fmt.Errorf("NO_EVENTS_TO_SEND_TO_COMM_ENGINE_ERROR")
		return err
	}
	//logger.Info(logString+", events Data", zap.Any("Events", events))
	logger.Info(logString+", total number of iterations", zap.Any("iterations", iterations),
		zap.Any("campaignId", campaingId))
	for i := 1; i <= iterations; i++ {
		err := SendMessage(ctx, &ce.BulkCommunicationEvent{CommunicationEvents: events[1000*(i-1) : cast.ToInt(math.Min(cast.ToFloat64(len(events)), cast.ToFloat64(1000*(i))))]})
		if err != nil {
			logger.Error("CampaignExecutor OnRequest hook ,Error in SendBulkCommunication", zap.Any("error", err),
				zap.Any("campaignId", campaingId))
			//metrics.Metrics.PushToErrorCounterMetrics()(metrics.ExecuteCampaignWorkflow_SendBulkCommunicationError_Metrics, fmt.Errorf("ERROR_SENDBULKCOMMUNICATION"), ctx)
			return err
		}
		logger.Info("CampaignExecutor OnRequest hook , sent bulk communication for iteration", zap.Any("iteration", i),
			zap.Any("campaignId", campaingId))
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func convertWaitForToString(waitFor int64) string {

	var res string
	if waitFor == 0 {
		return res
	}
	if waitFor%(60*60*24*30) == 0 {
		res = cast.ToString(waitFor/(60*60*24*30)) + " MONTH"
	} else if waitFor%(60*60*24*7) == 0 {
		res = cast.ToString(waitFor/(60*60*24*7)) + " WEEK"
	} else if waitFor%(60*60*24) == 0 {
		res = cast.ToString(waitFor/(60*60*24)) + " DAY"
	} else if waitFor%(60*60) == 0 {
		res = cast.ToString(waitFor/(60*60)) + " HOUR"
	} else if waitFor%(60) == 0 {
		res = cast.ToString(waitFor/(60)) + " MINUTE"
	}
	return res
}

func convertWaitFor(waitFor string) (int64, error) {

	time, err := parseTimeStamp(waitFor)
	if err != nil {
		logger.Error("Error in converting WaitFor", zap.Error(err), zap.Any("waitFor", waitFor))
		return 0, err
	}
	return time.Unix(), nil
}

func parseTimeStamp(timestamp string) (time.Time, error) {

	if timestamp == EMPTY {
		return time.Time{}, nil
	}

	layout := CONST_TIMESTAMP_LAYOUT
	time, err := time.Parse(layout, timestamp)

	if err != nil {
		logger.Error("Error in parsing time", zap.Error(err), zap.Any("timestamp", timestamp))
		return time, err
	}
	return time, nil
}

func convertToWaitTime(waitTime string) (int64, error) {

	var numeric int64
	numericString, unit, err := getTimeNumbericAndUnit(waitTime)
	if err != nil {
		logger.Error("Error in spliting waitTime", zap.Any("waitTime", waitTime), zap.Error(err))
		return numeric, err
	}
	unit = strings.ToUpper(unit)
	switch unit {
	case "MINUTE":
		numeric = cast.ToInt64(numericString) * 60
	case "HOUR":
		numeric = cast.ToInt64(numericString) * 60 * 60
	case "DAY":
		numeric = cast.ToInt64(numericString) * 60 * 60 * 24
	case "WEEK":
		numeric = cast.ToInt64(numericString) * 60 * 60 * 24 * 7
	case "MONTH":
		numeric = cast.ToInt64(numericString) * 60 * 60 * 24 * 30
	}
	return numeric, nil
}

func getTimeNumbericAndUnit(time string) (string, string, error) {

	waitTimeSplit := strings.Split(time, " ")
	if len(waitTimeSplit) < 2 {
		return EMPTY, EMPTY, fmt.Errorf("TIME_SPLIT_ERROR")
	}
	return waitTimeSplit[0], waitTimeSplit[1], nil
}

func addPropertiesFilter(filters [][]*common.Attribs, query string) string {
	if len(filters) > 0 {
		for index, filter := range filters {
			if len(filter) == 0 || len(filter)%3 != 0 {
				continue
			}
			i := 0
			for i < len(filter) {
				query += " and json_extract_scalar(event_meta_data" + cast.ToString(index+1) + ", '$[" + string('"') + filter[i+1].Value + string('"') + "]') in (" + addQuotesToCsvValues(filter[i+2].Value) + ")"
				i += 3
			}
		}
	}
	return query
}

func addQuotesToCsvValues(csv string) string {
	values := strings.Split(csv, ",")

	var sb strings.Builder

	for index, value := range values {
		sb.WriteString("'")
		sb.WriteString(value)
		sb.WriteString("'")

		if index < len(values)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func getTimeBetweenExecutions(cronExp string) int64 {
	args := strings.Split(strings.Split(cronExp, " ")[0], "/")[1]
	return cast.ToInt64(args) * 60 * 1000
}

func MakeUserJourneyCampaignQuery(namespace string, namespaceValue int32, userJourneyCampaign *fs.UserJourneyCampaign, response *fs.FindCampaignByIdResponse) ([]string, error) {
	var queries []string
	var index int32
	currentTime := time.Now().Unix() * 1000
	timeBetweenExecutions := getTimeBetweenExecutions(response.Records.CronExpression)
	userFilterQueries := GetUserFilterQueries(namespace, response)
	for _, userJourney := range userJourneyCampaign.UserJourneys {
		query := "SELECT ue1.user_id as id from "
		inactionQuery := query
		var eventNames []string
		var filters [][]*common.Attribs
		var edgeWaitDuration []int64
		var eventNamesWithInactionEventName []string
		var edgeWaitDurationWithInactionWaitDuration []int64
		var node interface{}
		var prevVertex *fs.UserJourneyVertex
		var EdgeIndex int32 //1-indexed
		inactionFlag := false
		EdgeIndex = 1
		node = userJourney.UserJourneyVertex
	loop:
		for {
			switch node.(type) {
			case *fs.UserJourneyVertex:
				currVertex := node.(*fs.UserJourneyVertex)
				if currVertex == nil {
					break loop
				}
				eventName := currVertex.EventMetadata.EventName
				filter := currVertex.EventMetadata.EventFilters
				eventNames = append(eventNames, eventName)
				filters = append(filters, filter)
				index, err := cache.EventCahce.GetEventIndex(namespace, eventName)
				if err != nil {
					return queries, err
				}
				if prevVertex != nil {
					EdgeIndex++
				}
				if EdgeIndex > 1 {
					query += " join "
				}
				startTimePlaceHolder := "@start_time@"
				if currVertex.Edge == nil || currVertex.Edge.EdgeType == common.CampaignEdgeType_EXIT {
					startTimePlaceHolder = "@last_vertex_start_time@"
				}
				query += "(select user_id, event_names, event_times, event_meta_data from " + EVENT_DB_NAME + "." + EVENT_TABLE_NAME + "  where namespace = " + cast.ToString(namespaceValue) + " and start_time >= " + startTimePlaceHolder + " and end_time <= @end_time@ and event_bits_" + cast.ToString((index/1000)+1) + "[" + cast.ToString(index) + "] = true and @partition_query@)"
				if EdgeIndex > 1 {
					query = strings.ReplaceAll(query, "@end_time@", "@non_start_end_time@")
				} else {
					query = strings.ReplaceAll(query, "@end_time@", "@start_end_time@")
				}
				query += " as ue" + cast.ToString(EdgeIndex)
				if EdgeIndex > 1 {
					query += " on ue" + cast.ToString(EdgeIndex-1) + ".user_id = ue" + cast.ToString(EdgeIndex) + ".user_id"
				}
				if currVertex.EventType == common.UserJourneyEventType_INACTION {
					inactionFlag = true
					EdgeIndex++
					inactionQuery = query
					inactionQuery += " join "
					eventNamesWithInactionEventName = eventNames
					edgeWaitDurationWithInactionWaitDuration = edgeWaitDuration
					inactionEventName := currVertex.InactionEventMetadata.EventName
					eventNamesWithInactionEventName = append(eventNames, inactionEventName)
					inactionEventIndex, err := cache.EventCahce.GetEventIndex(namespace, inactionEventName)
					if err != nil {
						return queries, err
					}
					inactionQuery += " (select user_id, event_names, event_times, event_meta_data from " + EVENT_DB_NAME + "." + EVENT_TABLE_NAME + "  where namespace = " + cast.ToString(namespaceValue) + " and start_time >= @start_time@ and end_time <= @end_time@ and event_bits_" + cast.ToString((inactionEventIndex/1000)+1) + "[" + cast.ToString(inactionEventIndex) + "] = true and @partition_query@)"
					inactionQuery = strings.ReplaceAll(inactionQuery, "@end_time@", "@start_end_time@")
					inactionQuery += " as ue" + cast.ToString(EdgeIndex)
					inactionQuery += " on ue" + cast.ToString(EdgeIndex-1) + ".user_id = ue" + cast.ToString(EdgeIndex) + ".user_id"
					inactionWaitDuration, _ := convertToWaitTime(currVertex.InactionDuration)
					edgeWaitDurationWithInactionWaitDuration = append(edgeWaitDuration, inactionWaitDuration)
				}
				prevVertex = currVertex
				node = currVertex.Edge
			case *fs.UserJourneyEdge:
				edge := node.(*fs.UserJourneyEdge)
				if edge == nil || edge.EdgeType == common.CampaignEdgeType_EXIT {
					break loop
				}
				waitFor, _ := convertToWaitTime(edge.WaitTime.WaitFor)
				edgeWaitDuration = append(edgeWaitDuration, waitFor)
				node = edge.UserJourneyVertex
			}
		}
		query = addFilterConditionUserJourneyCampaignQuery(currentTime, eventNames, edgeWaitDuration, query, filters, timeBetweenExecutions)
		if inactionFlag {
			inactionQuery = addFilterConditionUserJourneyCampaignQuery(currentTime, eventNamesWithInactionEventName, edgeWaitDurationWithInactionWaitDuration, inactionQuery, filters, timeBetweenExecutions)
			query = "SELECT id FROM (" + query + ") WHERE id not in (" + inactionQuery + ")"
			if userFilterQueries != nil && len(userFilterQueries[index]) > 0 {
				query += " AND id in ( " + userFilterQueries[index] + " )"
			}
		} else {
			if userFilterQueries != nil && len(userFilterQueries[index]) > 0 {
				query = " SELECT id FROM (" + query + ") WHERE id in ( " + userFilterQueries[index] + " )"
			}
		}
		queries = append(queries, query)
		index += 1
	}
	return queries, nil
}

func GetUserFilterQueries(namespace string, response *fs.FindCampaignByIdResponse) []string {
	attributes := response.Records.Attributes
	_, _, userMetadataList := mappers.MapMetaData(attributes)
	var userFilterQueries []string
	for index := 0; index < len(userMetadataList); index++ {
		userFilterQuery := ""
		if userMetadataList[index] != nil {
			userFilterQuery = ""
			isFirst := true
			userFilters := userMetadataList[index].UserFilters
			for userFilterIndex := 0; userFilterIndex < len(userFilters); userFilterIndex++ {
				if userFilters[userFilterIndex].Key == "STATE" {
					if isFirst {
						userFilterQuery += " where "
						isFirst = false
					} else {
						userFilterQuery += " and "
					}
					userFilterQuery += getQueryForKey(userFilters[userFilterIndex].Key, userFilters[userFilterIndex].Value)
				}
				if userFilters[userFilterIndex].Key == "CROP" {
					if isFirst {
						userFilterQuery += " where "
						isFirst = false
					} else {
						userFilterQuery += " and "
					}
					userFilterQuery += getQueryForKey(userFilters[userFilterIndex].Key, userFilters[userFilterIndex].Value)
				}
			}
		}
		userFilterQueries = append(userFilterQueries, userFilterQuery)
	}
	return userFilterQueries

}

func getQueryForKey(key string, value string) string {
	query := ""
	if key == "STATE" {
		query += "LOWER(fu.state) in ("
		states := strings.Split(value, ",")
		for index := 0; index < len(states); index++ {
			if index != 0 {
				query += ","
			}

			query += "'" + strings.ToLower(states[index]) + "'"
		}
		query += ")"
	}
	if key == "CROP" {
		query += "LOWER(c.name) in ("
		crops := strings.Split(value, ",")
		for index := 0; index < len(crops); index++ {
			if index != 0 {
				query += ","
			}
			query += "'" + strings.ToLower(crops[index]) + "'"
		}
		query += ")"
	}
	return query
}
func getPartitionQuery(startTime time.Time, endTime time.Time) string {
	startPartQuery := " ((year = " + cast.ToString(startTime.Year()) + " AND ((month = " + cast.ToString(int(startTime.Month())) + " AND ((day = " + cast.ToString(startTime.Day()) + " AND hour >= " + cast.ToString(startTime.Hour()) + ") OR (day > " + cast.ToString(startTime.Day()) + "))) OR month > " + cast.ToString(int(startTime.Month())) + ")) OR year > " + cast.ToString(startTime.Year()) + ")"
	endPartQuery := " ((year = " + cast.ToString(endTime.Year()) + " AND ((month = " + cast.ToString(int(endTime.Month())) + " AND ((day = " + cast.ToString(endTime.Day()) + " AND hour <= " + cast.ToString(endTime.Hour()) + ") OR (day < " + cast.ToString(endTime.Day()) + "))) OR month < " + cast.ToString(int(endTime.Month())) + ")) OR year < " + cast.ToString(endTime.Year()) + ")"
	return startPartQuery + " AND " + endPartQuery
}

func getStartTime(currentTime int64, timeSinceLastExecution int64, totalWaitTime int64) int64 {
	windowTime := viper.GetInt64("flink_windowing_time")
	startTime := currentTime - timeSinceLastExecution - totalWaitTime - windowTime
	windowTimeAligner := startTime % windowTime
	return startTime - windowTimeAligner
}

func getLastVertexStartTime(currentTime int64, timeSinceLastExecution int64) int64 {
	windowTime := viper.GetInt64("flink_windowing_time")
	startTime := currentTime - timeSinceLastExecution - windowTime
	windowTimeAligner := startTime % windowTime
	return startTime - windowTimeAligner
}

func addFilterConditionUserJourneyCampaignQuery(currentTime int64, eventNames []string, edgeWaitDuration []int64, query string, filters [][]*common.Attribs, timeSinceLastExecution int64) string {

	var totalWaitTime int64
	for _, edgeWaitTime := range edgeWaitDuration {
		totalWaitTime += edgeWaitTime
	}
	totalWaitTime = totalWaitTime * 1000
	startTime := getStartTime(currentTime, timeSinceLastExecution, totalWaitTime)
	startTimeString := cast.ToString(startTime)
	startEndTimeString := cast.ToString(currentTime)
	loc, _ := time.LoadLocation("UTC")
	startTimeUTC := time.Unix((startTime)/1000, 0).In(loc)
	endTimeUTC := time.Unix(currentTime/1000, 0).In(loc)
	query = strings.ReplaceAll(query, "@partition_query@", cast.ToString(getPartitionQuery(startTimeUTC, endTimeUTC)))
	query = strings.ReplaceAll(query, "@start_time@", startTimeString)
	query = strings.ReplaceAll(query, "@last_vertex_start_time@", cast.ToString(getLastVertexStartTime(currentTime, timeSinceLastExecution)))
	query = strings.ReplaceAll(query, "@non_start_end_time@", cast.ToString(currentTime))
	query = strings.ReplaceAll(query, "@start_end_time@", cast.ToString(currentTime))
	if len(eventNames) > 0 {
		for index, _ := range eventNames {
			query += ", UNNEST(ue" + cast.ToString(index+1) + ".event_names,ue" + cast.ToString(index+1) + ".event_times,ue" + cast.ToString(index+1) + ".event_meta_data) AS t (event_name" + cast.ToString(index+1) + ",event_time" + cast.ToString(index+1) + ",event_meta_data" + cast.ToString(index+1) + ")"
		}
	}
	query += " where true"
	if len(eventNames) > 0 {
		for index, eventNaame := range eventNames {
			query += " and event_name" + cast.ToString(index+1) + " = '" + cast.ToString(eventNaame) + "'"
		}
	}
	query = addPropertiesFilter(filters, query)
	if len(edgeWaitDuration) > 0 {
		for index, waitDuration := range edgeWaitDuration {
			if index == 0 {
				query += " and event_time" + cast.ToString(index+1) + " >= " + startTimeString + " and event_time" + cast.ToString(index+1) + " <= " + startEndTimeString
			}
			query += " and event_time" + cast.ToString(index+2) + " - event_time" + cast.ToString(index+1) + " <= " + cast.ToString(waitDuration*1000)
			if index == len(edgeWaitDuration)-1 {
				query += " group by ue1.user_id "
			}
		}
	}
	return query
}

func convertUtcToIst(Time time.Time) time.Time {

	Time = Time.Add(5 * time.Hour)
	Time = Time.Add(30 * time.Minute)
	return Time
}

func convertIstToUtc(Time time.Time) time.Time {

	Time = Time.Add(-5 * time.Hour)
	Time = Time.Add(-30 * time.Minute)
	return Time
}

func AddUserJourneyVertexArgs(model *AddUserJourneyVertexVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.EventType)
	args = append(args, model.EventName)
	args = append(args, model.InactionDuration)
	args = append(args, model.InactionEventName)
	args = append(args, model.Attributes)
	return args
}

func AddEngagementVertexArgs(model *AddEngagementVertexVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.TemplateName)
	args = append(args, model.Attributes)
	args = append(args, model.AthenaQuery)
	args = append(args, model.Channel)
	return args
}

func AddEdgeArgs(model *AddEdgeVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.VertexType)
	args = append(args, model.FromVertexId)
	args = append(args, model.ToVertexId)
	args = append(args, model.WaitDuration)
	args = append(args, model.WaitTime)
	args = append(args, model.WaitType)
	args = append(args, model.MessageDeliveryStates)
	return args
}

func FindNextEngagementVerticesArgs(campaignId int64, engagementVertexId int64) []interface{} {

	var args []interface{}
	args = append(args, campaignId)
	args = append(args, engagementVertexId)
	return args
}

func FindPreviousEngagementVertexArgs(campaignId int64, engagementVertexId int64) []interface{} {

	var args []interface{}
	args = append(args, campaignId)
	args = append(args, engagementVertexId)
	return args
}

func FindEngagementVertexByIdArgs(engagementVertexId int64) []interface{} {

	var args []interface{}
	args = append(args, engagementVertexId)
	return args
}

func FindEngagementStartVertexByIdArgs(campaignId int64) []interface{} {

	var args []interface{}
	args = append(args, campaignId)
	return args
}

func FindUserJourneyCampaignByIdArgs(request *fs.FindUserJourneyCampaignByIdRequest) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)
	return args
}

func FindUserJourneyCampaignEdgesByIdArgs(request *fs.FindUserJourneyCampaignByIdRequest, vertexType string) []interface{} {

	var args []interface{}
	args = append(args, request.CampaignId)
	args = append(args, vertexType)
	return args
}

func GetUserJourneyTargetUsersArgs(campaignId int64, engagementsVertexId int64, referenceId string) []interface{} {

	var args []interface{}
	args = append(args, campaignId)
	args = append(args, engagementsVertexId)
	args = append(args, referenceId)
	return args
}

func AddUserJourneyTargetUsersArgs(model *AddUserJourneyTargetUsersVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.EngagementVertexId)
	args = append(args, model.ReferenceId)
	args = append(args, model.EventReferenceId)
	args = append(args, model.ActorContactId)
	args = append(args, model.Status)
	return args
}

func UpdateCampaignArgs(model *UpdateCampaignVO) []interface{} {

	var args []interface{}
	args = append(args, model.CronExpression)
	args = append(args, model.Status)
	args = append(args, model.Attributes)
	args = append(args, model.CampaignID)

	return args
}

func DeleteEngagementVerticesArgs(model *DeleteEngagementVerticesVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)

	return args
}

func DeleteUserJourneyVerticesArgs(model *DeleteUserJourneyVerticesVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)

	return args
}

func DeleteEdgesVOArgs(model *DeleteEdgesVO) []interface{} {

	var args []interface{}
	args = append(args, model.CampaignId)
	args = append(args, model.VertexType)

	return args
}
