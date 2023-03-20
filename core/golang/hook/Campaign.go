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
	query "code.nurture.farm/platform/CampaignService/core/golang/database"
	"code.nurture.farm/platform/CampaignService/core/golang/hook/aws"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/executor"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/mappers"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/metrics"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

const (
	CONST_ID                      = "id"
	CONST_ATHENA_NIL              = "NIL"
	CONST_EMAIL_ID                = "email_id"
	CONST_MOBILE_NUMBER           = "mobile_number"
	CONST_LANG_KEY                = "lang_key"
	CONST_FCM_TOKEN               = "push_token"
	CONST_MODIFIED_ON             = "modified_on"
	CONST_APP_ID                  = "app_id"
	CONST_ACTOR_TYPE              = "user_type"
	CONST_MESSAGE_EXPIRY_DURATION = time.Hour * 5
	CONST_ACTOR_ID_SYSTEM         = 1000
	CONST_IMAGE                   = "image"
	CONST_UNDERSCORE              = "_"
	CONST_FARMER_APP_ID           = "NF_FARMER"
	CONST_RETAILER_APP_ID         = "NF_RETAILER"
	CONST_PARTNER_APP_ID          = "NF_PARTNER"
	CONST_DEFAULT_AP_ID           = ""
	CONST_ZERO                    = 0
	CONST_ERROR_CODE_NIL          = -1
	CONST_CTA                     = "cta"
	CONST_LOCATION                = "location"
	CONST_CAPACITY                = "capacity"
	CONST_K                       = "k"
	TEXT_EXTENSION                = ".txt"
)

type CampaignInterface interface {
	OnRequest(ctx context.Context, request *fs.CampaignRequest) *fs.CampaignResponse
	OnData(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse
	OnError(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse, err error) *fs.CampaignResponse
	OnResponse(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse
}

type GenericCampaignExecutor struct {
	CampaignInterface CampaignInterface
}

type CampaignController struct {
}

var CampaignExecutor *GenericCampaignExecutor

type trackingData struct {
	CampaignName string `json:"campaignName"`
}

func setCampaignResponse(status common.RequestStatus, errorMessages []string, errorCode common.ErrorCode) *fs.CampaignResponse {
	if errorCode == CONST_ERROR_CODE_NIL {
		return &fs.CampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        status,
				ErrorMessages: errorMessages,
			},
		}
	}
	return &fs.CampaignResponse{
		Status: &common.RequestStatusResult{
			Status:        status,
			ErrorMessages: errorMessages,
			ErrorCode:     errorCode,
		},
	}
}

func (ge *GenericCampaignExecutor) OnRequest(ctx context.Context, request *fs.CampaignRequest) *fs.CampaignResponse {

	var err error
	defer metrics.Metrics.PushToSummarytMetrics()(metrics.GetUserList_Metrics, "GetUserListExecutor", &err, ctx)
	logger.Info("CampaignExecutor OnRequest hook started", zap.Any("request", request))

	findCampaignByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignById(ctx, &fs.FindCampaignByIdRequest{
		Id: request.CampaignId,
	})
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindCampaignById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindCampaignById"}, CONST_ERROR_CODE_NIL)
	}
	if findCampaignByIdResponse == nil || findCampaignByIdResponse.Records == nil {
		logger.Error("CampaignExecutor OnRequest hook, Error, could not get Campaign from DB based on campaignId", zap.Any("campaignId", request.CampaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindCampaignById"}, CONST_ERROR_CODE_NIL)
	}
	campaignQueryType := common.CampaignQueryType(common.CampaignQueryType_value[findCampaignByIdResponse.Records.Type])
	findCampaignTemplateByIdResponse, err := executor.RequestExecutor.ExecuteFindCampaignTemplateById(ctx, &fs.FindCampaignTemplateByIdRequest{
		CampaignId: request.CampaignId,
	})
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindCampaignTemplateById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindCampaignTemplateById"}, CONST_ERROR_CODE_NIL)
	}
	if findCampaignTemplateByIdResponse == nil || findCampaignTemplateByIdResponse.Records == nil {
		logger.Error("CampaignExecutor OnRequest hook, Error could not get CampaignTemplates from DB based on campaignId", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindCampaignTemplateById"}, CONST_ERROR_CODE_NIL)
	}

	findTargetUserByIdResponse, err := executor.RequestExecutor.ExecuteFindTargetUserById(ctx, &fs.FindTargetUserByIdRequest{
		CampaignId: request.CampaignId,
	})
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindTargetUserById", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindTargetUserById"}, CONST_ERROR_CODE_NIL)
	}

	channel := common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
	campaignQuery := findCampaignByIdResponse.Records.Query

	campaignScheduleType := findCampaignByIdResponse.Records.ScheduleType
	campaignStatus := findCampaignByIdResponse.Records.Status
	var inactionTargetUsers []*common.ActorID
	var inactionPhaseTwoErrorResponse *fs.CampaignResponse

	//note: campaignQueryType will be of type target user List as we want to send PNs to targeted users only (for Inaction Overtime)
	if campaignScheduleType == common.CampaignScheduleType_INACTION_OVER_TIME.String() {
		if campaignStatus == common.CampaignStatus_RUNNING.String() {
			//PHASE-1 OF INACTION CAMPAIGN
			return handleInactionCampaignPhaseOne(ctx, request, findCampaignByIdResponse)
		} else if campaignStatus == common.CampaignStatus_PRE_INACTION.String() {
			//PHASE-2 OF INACTION CAMPAIGN
			inactionTargetUsers, inactionPhaseTwoErrorResponse = handleInactionCampaignPhaseTwo(ctx, request, findCampaignByIdResponse)
			if inactionPhaseTwoErrorResponse != nil {
				return inactionPhaseTwoErrorResponse
			} //todo : response should be only either error or response
		}
	}

	targetUsers := []*common.ActorID{}
	targetUsersPlaceHolderMap := make(map[string][]*common.Attribs)
	//note: campaignQueryType will be of type target user List as we want to send PNs to targeted users only (for Inaction Overtime)
	if campaignScheduleType == common.CampaignScheduleType_INACTION_OVER_TIME.String() && campaignStatus == common.CampaignStatus_PRE_INACTION.String() {
		targetUsers = inactionTargetUsers
		campaignQueryType = common.CampaignQueryType_USER_LIST
		//campaignType = "USER_LIST"
	} else {
		for _, val := range findTargetUserByIdResponse.Records {
			actorID := common.ActorID{
				ActorType: common.ActorType(common.ActorType_value[val.UserType]),
				ActorId:   val.UserId,
			}
			targetUsers = append(targetUsers, &actorID)
			targetUsersPlaceHolderMap[cast.ToString(actorID.ActorId)] = val.Attribs
		}
	}

	if campaignQueryType == common.CampaignQueryType_NO_CAMPAIGN_QUERY_TYPE ||
		(campaignQueryType == common.CampaignQueryType_ATHENA && campaignQuery == "") ||
		(campaignQueryType == common.CampaignQueryType_USER_LIST && targetUsers == nil) {
		logger.Error("CampaignExecutor OnRequest hook, Invalid request", zap.Any("request", request))
		err = fmt.Errorf("INVALID_REQUEST")
		return setCampaignResponse(common.RequestStatus_BAD_INPUT, nil, CONST_ERROR_CODE_NIL)
	}

	var actorDetailsData [][]string
	actorDetailsData, err = getActorDetailsData(ctx, "CampaignExecutor OnRequest hook", campaignQueryType, channel, campaignQuery, targetUsers, request.CampaignId)
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in getting actor details", zap.Any("request", request))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in getting actor details"}, CONST_ERROR_CODE_NIL)
	}

	actorDetailsData = mapUserListPlaceHolders(targetUsersPlaceHolderMap, actorDetailsData)
	controlGroupPercentage := mappers.MapControlGroupPercentage(findCampaignByIdResponse.Records.Attributes)
	if controlGroupPercentage != 0 {
		findControlGroupByCampaignId, err := executor.RequestExecutor.ExecuteFindControlGroupByCampaignId(ctx, &fs.FindControlGroupByCampaignIdRequest{
			CampaignId: request.CampaignId,
		})
		if err != nil {
			logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindControlGroupByCampaignId", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
			return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindControlGroupByCampaignId"}, CONST_ERROR_CODE_NIL)
		}
		if findControlGroupByCampaignId.Records == nil {
			addControlGroup(actorDetailsData, controlGroupPercentage, request.CampaignId)
			findControlGroupByCampaignId, err = executor.RequestExecutor.ExecuteFindControlGroupByCampaignId(ctx, &fs.FindControlGroupByCampaignIdRequest{
				CampaignId: request.CampaignId,
			})
			if err != nil {
				logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindControlGroupByCampaignId", zap.Any("error", err), zap.Any("campaignId", request.CampaignId))
				return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindControlGroupByCampaignId"}, CONST_ERROR_CODE_NIL)
			}
		}
		if findControlGroupByCampaignId != nil && findControlGroupByCampaignId.Records != nil {
			controlGroupBloomFilter, err := getBloomFilter(findControlGroupByCampaignId)
			if err != nil {
				logger.Error("CampaignExecutor OnRequest hook, Error in GettingBloomFilter", zap.Any("error", err), zap.Any("ControlFroup", findControlGroupByCampaignId))
				return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in GettingBloomFilter"}, CONST_ERROR_CODE_NIL)
			}
			if controlGroupBloomFilter != nil {
				actorDetailsData, err = removeControlGroupUserIds(actorDetailsData, controlGroupBloomFilter)
				if err != nil {
					logger.Error("CampaignExecutor OnRequest hook, Error in getting actor details for control group", zap.Any("request", request))
					return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in getting actor details after control group"}, CONST_ERROR_CODE_NIL)
				}
			}
		}
	}
	events, errResponse := getEventsToBeSent(findCampaignTemplateByIdResponse, actorDetailsData, findCampaignByIdResponse, request.CampaignId, 0, ctx)
	if errResponse != nil {
		return errResponse
	}
	logger.Info("CampaignExecutor OnRequest hook ", zap.Any("Total number of communication events", len(events)),
		zap.Any("campaignId", request.CampaignId))

	if len(events) == CONST_ZERO {
		logger.Info("CampaignExecutor OnRequest hook, No/zero events to send to communication engine ",
			zap.Any("request", request))
		return setCampaignResponse(common.RequestStatus_REQUEST_NOT_FULFILLED, nil, CONST_ERROR_CODE_NIL)
	}
	err = sendCommunicationEvent(ctx, "CampaignExecutor OnRequest hook ", events, request.CampaignId)
	if err != nil {
		logger.Info("CampaignExecutor OnRequest hook, Error in sending communication event ", zap.Error(err),
			zap.Any("request", request))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, nil, CONST_ERROR_CODE_NIL)
	}

	response := setCampaignResponse(common.RequestStatus_SUCCESS, nil, CONST_ERROR_CODE_NIL)
	logger.Info("CampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return response
}

func addControlGroup(actorDetailsData [][]string, controlGroupPercentage int32, campaignId int64) *fs.CampaignResponse {
	userIds, err := extractUserIds(actorDetailsData)
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in  extractUserIds", zap.Any("error", err), zap.Any("campaignId", campaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in extractUserIds"}, CONST_ERROR_CODE_NIL)
	}
	controlGroupUserIds := GenerateRandomUserIds(userIds, controlGroupPercentage)
	key := makeKey(campaignId)
	err = aws.PutObjectInBucket([]byte(arrayToString(controlGroupUserIds, ",")), viper.GetString("s3_bucket_name.control_group"), key)
	if err != nil {
		logger.Error("Error in uploading data to s3", zap.Any("error", err), zap.Any("campaignId", campaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in uploading data to s3"}, CONST_ERROR_CODE_NIL)
	}
	if len(controlGroupUserIds) != 0 {
		bloomFilter := BloomFilter(controlGroupUserIds)
		_, err = executor.RequestExecutor.ExecuteAddControlGroup(context.Background(), &fs.AddControlGroupRequest{
			CampaignId:  campaignId,
			Attributes:  mappers.GetControlGroupUserAttributes(int(controlGroupPercentage), bloomFilter),
			BloomFilter: mappers.GetBitSetOfBloomFilter(bloomFilter),
		})
	} else {
		_, err = executor.RequestExecutor.ExecuteAddControlGroup(context.Background(), &fs.AddControlGroupRequest{
			CampaignId:  campaignId,
			Attributes:  mappers.GetControlGroupUserAttributes(int(controlGroupPercentage), bloom.New(0, 0)),
			BloomFilter: nil,
		})
	}
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in ExecuteFindControlGroupByCampaignId", zap.Any("error", err), zap.Any("campaignId", campaignId))
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in ExecuteFindControlGroupByCampaignId"}, CONST_ERROR_CODE_NIL)
	}
	return nil
}

func makeKey(campaignId int64) string {
	return cast.ToString(campaignId) + TEXT_EXTENSION
}

func arrayToString(a []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func extractUserIds(actorDetailsData [][]string) ([]int64, error) {
	var userIds []int64
	var indexOfId int
	for i, rowData := range actorDetailsData {
		if i == 0 {
			indexOfId = GetIndexOfId(rowData)
			continue
		}
		userId, err := convertToInt64(actorDetailsData[i][indexOfId])
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, userId)
	}
	return userIds, nil
}

func removeControlGroupUserIds(actorDetailsData [][]string, controlGroupBloomFilter *bloom.BloomFilter) ([][]string, error) {
	finalActorDetailsData := [][]string{}
	var indexOfId int
	for i, rowData := range actorDetailsData {
		if i == 0 {
			finalActorDetailsData = append(finalActorDetailsData, rowData)
			indexOfId = GetIndexOfId(rowData)
			continue
		}
		userId, err := convertToInt64(actorDetailsData[i][indexOfId])
		if err != nil {
			return nil, err
		}
		if !IsPresent(controlGroupBloomFilter, userId) {
			finalActorDetailsData = append(finalActorDetailsData, rowData)
		}
	}
	return finalActorDetailsData, nil

}
func mapUserListPlaceHolders(targetUsersPlaceHolderMap map[string][]*common.Attribs, actorDetailsData [][]string) [][]string {

	if len(targetUsersPlaceHolderMap) == 0 {
		return actorDetailsData
	}
	targetUsersDetailsData := [][]string{}
	for index, row := range actorDetailsData {
		if index == 0 {
			placeHolders := getDistinctPlaceHolders(targetUsersPlaceHolderMap)
			row = append(row, placeHolders...)
			targetUsersDetailsData = append(targetUsersDetailsData, row)
			continue
		}
		val, ok := targetUsersPlaceHolderMap[row[0]]
		if ok {
			for _, placeHolderData := range val {
				row = append(row, placeHolderData.Value)
			}
			targetUsersDetailsData = append(targetUsersDetailsData, row)
		}
	}
	return targetUsersDetailsData
}

func getDistinctPlaceHolders(targetUsersPlaceHolderMap map[string][]*common.Attribs) []string {

	placeHolders := []string{}
	for _, value := range targetUsersPlaceHolderMap {
		for _, placeholderName := range value {
			placeHolders = append(placeHolders, placeholderName.Key)
		}
		break
	}
	return placeHolders
}

func mapToFCampaignResponse(request *fs.CampaignRequest, status common.RequestStatus, err error) *fs.CampaignResponse {

	if err != nil {
		switch status {
		case common.RequestStatus_BAD_INPUT:
			logger.Error("CampaignExecutor OnRequest hook, invalid request", zap.Any("request", request), zap.Error(err))
		case common.RequestStatus_INTERNAL_ERROR:
			logger.Error("CampaignExecutor OnRequest hook, internal error", zap.Any("request", request), zap.Error(err))
		}
		return &fs.CampaignResponse{
			Status: &common.RequestStatusResult{
				Status:        status,
				ErrorMessages: []string{err.Error()},
			},
		}
	}
	logger.Info("CampaignExecutor OnRequest hook completed successfully", zap.Any("request", request))
	return &fs.CampaignResponse{
		Status: &common.RequestStatusResult{
			Status: status,
		},
	}
}

func isQueryContainUserType(campaignQuery string) bool {
	res := strings.ToLower(campaignQuery)
	index := strings.Index(res, "from")
	campaignQueryTest := campaignQuery[0:index]
	if strings.Contains(campaignQueryTest, "user_type") {
		return true
	}
	return false
}

func getDistinctUserTypes(ctx context.Context, campaignQuery string, campaignId int64) ([]common.ActorType, error) {

	var distinctActors []common.ActorType
	if !isQueryContainUserType(campaignQuery) {
		return distinctActors, nil
	}
	_, result, err := aws.ExecuteAthenaQuery(ctx, strings.ReplaceAll(query.QUERY_DISTINCT_USERS, "@query@", campaignQuery), campaignId)
	logger.Info("getDistinctUserTypes method, ",
		zap.Any("result", len(result)))
	if err != nil && err.Error() == "ATHENA_QUERY_FAILED" {
		logger.Info("getDistinctUserTypes method, campaign doesn't have distinct actors", zap.Any("camppaignId", campaignId),
			zap.Any("campaignQuery", campaignQuery))
		return distinctActors, nil
	}
	if err != nil {
		logger.Error("Error in getDistinctUserTypes from athena query", zap.Any("query", campaignQuery), zap.Any("error", err))
		return nil, err
	}
	for index, row := range result {
		if index == 0 {
			continue
		}
		distinctActors = append(distinctActors, common.ActorType(common.ActorType_value[strings.ToUpper(row[0])]))
	}
	return distinctActors, nil
}

func setDynamicMediaBanner(findCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse, imageMap map[common.LanguageCode]string,
	contentMetaData []*ce.Placeholder) ([]*ce.Placeholder, map[common.LanguageCode]string) {
	key1 := time.Now().Format("2006.01.02.15")
	key2 := time.Now().Format("2006.01.02")
	logger.Info("Current Time & Current Date", zap.Any("key1", key1), zap.Any("key2", key2))
	dynamicDataByKey1, err := executor.RequestExecutor.ExecuteGetDynamicDataByKey(context.Background(),
		&fs.GetDynamicDataByKeyRequest{
			CampaignId: findCampaignTemplateByIdResponse.Records[0].CampaignId,
			DynamicKey: key1,
		})
	if err != nil || dynamicDataByKey1 == nil || len(dynamicDataByKey1.Records) == 0 {
		dynamicDataByKey2, err2 := executor.RequestExecutor.ExecuteGetDynamicDataByKey(context.Background(),
			&fs.GetDynamicDataByKeyRequest{
				CampaignId: findCampaignTemplateByIdResponse.Records[0].CampaignId,
				DynamicKey: key2,
			})
		if err2 != nil {
			logger.Error("Error when calling function setDynamicMediaBanner::data Not Found Both Key1 & Key2",
				zap.Error(err2),
				zap.Any("key1", key1), zap.Any("key2", key2),
			)
		} else {
			contentMetaData, imageMap = updateMedia(dynamicDataByKey2, imageMap, contentMetaData)
		}
	} else {
		contentMetaData, imageMap = updateMedia(dynamicDataByKey1, imageMap, contentMetaData)
	}
	return contentMetaData, imageMap
}

func updateMedia(dynamicDataByKey *fs.GetDynamicDataByKeyResponse,
	imageMap map[common.LanguageCode]string,
	contentMetaData []*ce.Placeholder) ([]*ce.Placeholder, map[common.LanguageCode]string) {
	media := make(map[string]string)
	error := json.Unmarshal([]byte(dynamicDataByKey.Records[0].Media), &media)

	if error != nil {
		logger.Error("Error when calling function updateMedia:: Media is not in json format ", zap.Error(error))
	} else {
		logger.Info("Initial ", zap.Any("Content Meta Data", contentMetaData), zap.Any("imageMap", imageMap))
		imageMap = setImageMap(media, imageMap)
		for index, attrib := range contentMetaData {
			attribKey := attrib.Key
			if strings.HasPrefix(attribKey, CONST_CTA) {
				contentMetaData[index].Value = dynamicDataByKey.Records[0].CtaLink
				break
			}
		}
		logger.Info("Updated ", zap.Any("Content Meta Data", contentMetaData), zap.Any("imageMap", imageMap))
	}
	return contentMetaData, imageMap
}

func setImageMap(media map[string]string, imageMap map[common.LanguageCode]string) map[common.LanguageCode]string {
	for attribKey, value := range media {
		if strings.HasPrefix(attribKey, CONST_IMAGE) {
			if attribKey == CONST_IMAGE {
				imageMap[common.LanguageCode_NO_LANGUAGE_CODE] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_EN_US.String() {
				imageMap[common.LanguageCode_EN_US] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_HI_IN.String() {
				imageMap[common.LanguageCode_HI_IN] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_GU.String() {
				imageMap[common.LanguageCode_GU] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_PA.String() {
				imageMap[common.LanguageCode_PA] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_KA.String() {
				imageMap[common.LanguageCode_KA] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_TA.String() {
				imageMap[common.LanguageCode_TA] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_TE.String() {
				imageMap[common.LanguageCode_TE] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_BN.String() {
				imageMap[common.LanguageCode_BN] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_MR.String() {
				imageMap[common.LanguageCode_MR] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_ML.String() {
				imageMap[common.LanguageCode_ML] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_KN.String() {
				imageMap[common.LanguageCode_KN] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_OD.String() {
				imageMap[common.LanguageCode_OD] = value
			} else if attribKey == CONST_IMAGE+CONST_UNDERSCORE+common.LanguageCode_PU.String() {
				imageMap[common.LanguageCode_PU] = value
			}
			continue
		}
	}
	return imageMap
}
func getStatetoLanguageMapping() map[string]string {
	var stateToLanguageMap = make(map[string]string)
	stateToLanguageMap["Gujarat"] = "gu"
	stateToLanguageMap["Punjab"] = "pa"
	stateToLanguageMap["Karnataka"] = "kn"
	stateToLanguageMap["Tamil Nadu"] = "ta"
	stateToLanguageMap["Telangana"] = "te"
	stateToLanguageMap["Andhra Pradesh"] = "te"
	stateToLanguageMap["West Bengal"] = "bn"
	stateToLanguageMap["Maharashtra"] = "mr"
	stateToLanguageMap["Kerala"] = "ml"
	return stateToLanguageMap
}

func getStateIdtoStateMapping() map[string]string {

	stateIdToStateMap := map[string]string{
		"695":  "Punjab",
		"947":  "Maharashtra",
		"2704": "Gujarat",
		"5276": "Telangana",
		"2710": "West Bengal",
		"2713": "Karnataka",
		"2714": "Tamil Nadu",
		"2712": "Kerala",
		"2473": "Andhra Pradesh",
	}
	return stateIdToStateMap

}
func getEventsToBeSent(findCampaignTemplateByIdResponse *fs.FindCampaignTemplateByIdResponse, actorDetailsData [][]string, findCampaignByIdResponse *fs.FindCampaignByIdResponse, campaignId int64,
	engagementVertexId int64, ctx context.Context) ([]*ce.CommunicationEvent, *fs.CampaignResponse) {
	channel := common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
	campaignContentMetaData := mappers.MapContentMetaData(findCampaignByIdResponse.Records.Attributes)
	channelAttributes := mappers.MapChannelAttributes(findCampaignByIdResponse.Records.Attributes)
	campaignMediaInitial := mappers.MapMedia(findCampaignByIdResponse.Records.Attributes)
	campaignMediaList := []*ce.Media{}
	var stateToLanguageMap = getStatetoLanguageMapping()
	var stateIdToStateMap = getStateIdtoStateMapping()
	contentMetaData, imageMap := getContentMetadata(campaignContentMetaData)
	campaignScheduleType := findCampaignByIdResponse.Records.ScheduleType
	actorDetailMap, actorByPlaceholderMap := getActorDetailsMap(actorDetailsData)
	templateSelectOption := "default"
	tentativeSegmentSize := ""
	for _, v := range contentMetaData {
		if v.Key == "templateSelectOption" {
			templateSelectOption = v.Value
		}
		if v.Key == "tentativeSize" {
			tentativeSegmentSize = v.Value
		}

	}
	temp := make(map[int64]UserDetail)
	if templateSelectOption == "location" {
		for k, v := range actorDetailMap {
			location := v.Location
			if location == "" {
				continue
			}
			if v.ActorType == "FARMER" {
				secondaryLangKey, ok := stateToLanguageMap[location]
				if ok {
					v.SecondaryLangKey = secondaryLangKey
				}
			} else {
				state, ok := stateIdToStateMap[location]
				if ok {
					secondaryLangKey, ok := stateToLanguageMap[state]
					if ok {
						v.SecondaryLangKey = secondaryLangKey
					}
				}
			}
			temp[k] = v
		}
	}
	for k, v := range temp {
		actorDetailMap[k] = v
	}

	actorDetails, actorIDDetails := getActorDetails(actorDetailMap, channel, campaignId)

	events := []*ce.CommunicationEvent{}
	if campaignScheduleType == common.CampaignScheduleType_DYNAMIC_MEDIA_TYPE.String() {
		contentMetaData, imageMap = setDynamicMediaBanner(findCampaignTemplateByIdResponse, imageMap, contentMetaData)
	}
	if tentativeSegmentSize != "" {
		segmentSize := cast.ToInt(tentativeSegmentSize)
		if len(actorDetails) > segmentSize {
			error := fmt.Errorf("User Segment Size: %d, Tentative Segment Size: %d, CampaignId: %d", len(actorDetails), segmentSize, campaignId)
			metrics.Metrics.PushToErrorCounterMetrics()(metrics.GetTentativeList_Error_Metrics, error, ctx)
			return events, setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"User segment is more than tentative segment size"}, CONST_ERROR_CODE_NIL)
		}
	}
	var parentReferenceId string
	if strings.ToUpper(findCampaignByIdResponse.Records.Type) == "USER_JOURNEY" {
		parentReferenceId = GetUserJourneyWorkflowParentReferenceId(campaignId, engagementVertexId)
	} else {
		parentReferenceId = GetWorkflowId(campaignId)
	}
	//distributing templates among actors based on distribution percentage (A/B testing purpose)
	startIndex := 0
	for _, requestTemplate := range findCampaignTemplateByIdResponse.Records {
		userList := (len(actorDetails) * cast.ToInt(requestTemplate.DistributionPercent)) / 100
		for i := startIndex; i < userList+startIndex; i++ {
			if i >= len(actorDetails) {
				break
			}
			trackingData := trackingData{
				CampaignName: requestTemplate.CampaignName,
			}
			trackingDataBytes, err := json.Marshal(trackingData)
			if err != nil {
				logger.Error("CampaignExecutor OnRequest hook, Error in marshalling tracking_data", zap.Any("error", err),
					zap.Any("campaignId", campaignId), zap.Any("tracking_data", trackingData))
				err = fmt.Errorf("TRACKING_DATA_MARSHAL_ERROR")
				return events, setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in marshalling tracking data"}, CONST_ERROR_CODE_NIL)
			}
			trackingDataPlaceHolder := &ce.Placeholder{Key: "tracking_data", Value: string(trackingDataBytes)}
			contentMetaDataWithTrackingData := contentMetaData
			contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, trackingDataPlaceHolder)
			image := GetImageMetaData(actorDetails[i].LanguageCode, imageMap)
			if image != "" {
				contentMetaDataWithTrackingData = append(contentMetaDataWithTrackingData, &ce.Placeholder{Key: CONST_IMAGE, Value: image})
			}
			campaignMedia := &ce.Media{}
			if channel == common.CommunicationChannel_WHATSAPP {
				campaignMedia = getWhatsAppMedia(campaignMediaInitial, image)
			}
			campaignMediaList = append(campaignMediaList, campaignMedia)
			if channel == common.CommunicationChannel_APP_NOTIFICATION {
				campaignMediaList[i] = nil
			}
			event := &ce.CommunicationEvent{
				TemplateName:         requestTemplate.TemplateName,
				Expiry:               timestamppb.New(timestamppb.Now().AsTime().Add(CONST_MESSAGE_EXPIRY_DURATION)), //check once
				Channel:              []common.CommunicationChannel{channel},
				ContentMetadata:      contentMetaDataWithTrackingData,
				ReceiverActor:        actorIDDetails[i],
				ReceiverActorDetails: actorDetails[i],
				ChannelAttributes:    channelAttributes,
				Media:                campaignMediaList[i],
				CampaignName:         requestTemplate.CampaignName,
				ParentReferenceId:    parentReferenceId,
				Placeholder:          actorByPlaceholderMap[actorIDDetails[i].ActorId],
				ReferenceId:          uuid.New().String(),
			}
			events = append(events, event)
		}
		startIndex += userList
	}
	return events, nil
}

func getWhatsAppMedia(campaignMediaInitial *ce.Media, image string) *ce.Media {
	campaignMedia := &ce.Media{}
	if campaignMediaInitial != nil {
		campaignMedia.MediaInfo = campaignMediaInitial.GetMediaInfo()
		campaignMedia.MediaAccessType = campaignMediaInitial.GetMediaAccessType()
		campaignMedia.DocumentName = campaignMediaInitial.GetDocumentName()
		campaignMedia.MediaType = campaignMediaInitial.GetMediaType()
		if image != "" {
			campaignMedia.MediaInfo = image
			if campaignMedia.MediaType == common.MediaType_DOCUMENT {
				campaignMedia.DocumentName = fetchDocumentName(image)
			}
		}
		return campaignMedia
	}
	return nil
}
func getContentMetadata(campaignContentMetaData []*common.Attribs) ([]*ce.Placeholder, map[common.LanguageCode]string) {
	contentMetaData := []*ce.Placeholder{}
	imageMap := make(map[common.LanguageCode]string)
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
	}
	return contentMetaData, imageMap
}

/*
*
Logic for Inaction queries

	Check if schedule type is INACTION_OVER_TIME
	Check the state of the campaign_id:
		if RUNNING --> Phase1: fetch users from Athena(campaignQuery), insert users into DB, update the campaign status to PRE_INACTION
		if PRE_INACTION --> Phase2: fetch users from new table and run inactionQuery to get users at second event, find inaction target users and send PNS to them, update status back to running

*
*/
func handleInactionCampaignPhaseOne(ctx context.Context, request *fs.CampaignRequest, findCampaignByIdResponse *fs.FindCampaignByIdResponse) *fs.CampaignResponse {
	var err error
	channel := common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
	campaignQuery := findCampaignByIdResponse.Records.Query
	nameSpace := common.NameSpace(common.NameSpace_value[findCampaignByIdResponse.Records.Namespace])
	var inactionCampaignUsersAtFirstEvent [][]string

	//PHASE-1 OF INACTION CAMPAIGN
	//fetch the users, and store them in DB in inaction_target_users table as current state
	_, inactionCampaignUsersAtFirstEvent, err = executeGetUserAthenaQuery(ctx, campaignQuery, channel, request.CampaignId)
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in executing Athena QUERY_InactionOvertime_Phase1", zap.Any("Query", campaignQuery), zap.Any("error", err))
		err = fmt.Errorf("ATHENA_QUERY_ERROR")
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in executing Athena QUERY_InactionOvertime_Phase1"}, CONST_ERROR_CODE_NIL)
	}
	//Insert request to add inaction target users to DB(Phase1)
	actorType := GetActorTypeUsingNamespace(nameSpace) //todo :- this is temp logic for getting actorType, needs a better logic
	addInactionTargetUserRequestList := getAddInactionTargetUserRequestList(inactionCampaignUsersAtFirstEvent, actorType, request.CampaignId)
	bulkRequest := &fs.BulkAddInactionTargetUserRequest{
		Requests: addInactionTargetUserRequestList,
	}
	response, err := executor.RequestExecutor.ExecuteAddInactionTargetUserBulk(ctx, bulkRequest)
	if err != nil {
		logger.Info("CampaignExecutor OnRequest hook, Error in ExecuteAddInactionTargetUserBulk request", zap.Any("request", request), zap.Any("response", response))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.BulkAddInactionTargetUser_Error_Metrics, err, ctx)
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error in executing Athena QUERY_InactionOvertime_Phase1"}, common.ErrorCode_DATABASE_ERROR)
	}
	logger.Info("Added Users in inaction_target_users table. ", zap.Any("response", response))

	// Update the campaign status from running to PRE_INACTION for the given campaignId
	updateCampaignResponse := updateCampaignStatus(ctx, request.CampaignId, common.CampaignStatus_RUNNING)
	if updateCampaignResponse != nil {
		return updateCampaignResponse
	}
	inactionPhaseOneResponse := setCampaignResponse(common.RequestStatus_SUCCESS, nil, CONST_ERROR_CODE_NIL)
	logger.Info("CampaignExecutor OnRequest hook (Inaction Phase One) completed successfully", zap.Any("request", request))
	return inactionPhaseOneResponse
}

func getAddInactionTargetUserRequestList(inactionCampaignUsersAtFirstEvent [][]string, actorType int32, campaignId int64) []*fs.AddInactionTargetUserRequest {
	var addInactionTargetUserRequestList []*fs.AddInactionTargetUserRequest
	var actorIdIndex int

	for i, val := range inactionCampaignUsersAtFirstEvent {
		if i == 0 {
			actorIdIndex = GetIndexInRow(val, CONST_ID)
			continue
		}
		actorId := GetInt64Value(actorIdIndex, val)
		addInactionTargetUserRequest := &fs.AddInactionTargetUserRequest{
			//Set your request here
			CampaignId: campaignId,
			User: &common.ActorID{
				ActorId:   actorId,
				ActorType: common.ActorType(actorType),
			},
		}
		addInactionTargetUserRequestList = append(addInactionTargetUserRequestList, addInactionTargetUserRequest)
	}
	return addInactionTargetUserRequestList
}

func handleInactionCampaignPhaseTwo(ctx context.Context, request *fs.CampaignRequest, findCampaignByIdResponse *fs.FindCampaignByIdResponse) ([]*common.ActorID, *fs.CampaignResponse) {

	var err error
	channel := common.CommunicationChannel(common.CommunicationChannel_value[findCampaignByIdResponse.Records.CommunicationChannel])
	campaignInactionQuery := findCampaignByIdResponse.Records.InactionQuery
	var inactionCampaignUsersAtSecondEvent [][]string
	var inactionCampaignUsersAtSecondEventMap = make(map[int64]bool)
	var inactionTargetUsers []*common.ActorID

	//PHASE-2 OF INACTION CAMPAIGN
	totalInactionCampaignUsersInDB, errResponse := fetchInactionCampaignUsersData(ctx, request)
	if errResponse != nil {
		return nil, errResponse
	}
	_, inactionCampaignUsersAtSecondEvent, err = executeGetUserAthenaQuery(ctx, campaignInactionQuery, channel, request.CampaignId)
	if err != nil {
		logger.Error("CampaignExecutor OnRequest hook, Error in executing Athena QUERY_InactionOvertime_Phase2", zap.Any("Query", campaignInactionQuery), zap.Any("error", err))
		err = fmt.Errorf("ATHENA_QUERY_ERROR")
		return nil, setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, nil, CONST_ERROR_CODE_NIL)
	}
	var idIndex int
	for i, val := range inactionCampaignUsersAtSecondEvent {
		if i == 0 {
			idIndex = GetIndexInRow(val, CONST_ID)
			continue
		}
		actorId := GetInt64Value(idIndex, val)
		inactionCampaignUsersAtSecondEventMap[actorId] = true
	}
	inactionTargetUsers = findInactionTargetUsers(inactionCampaignUsersAtSecondEventMap, totalInactionCampaignUsersInDB)

	// Update the campaign status from PRE_INACTION to RUNNING again for the given campaignId
	updateCampaignResponse := updateCampaignStatus(ctx, request.CampaignId, common.CampaignStatus_RUNNING)
	if updateCampaignResponse != nil {
		return nil, updateCampaignResponse
	}
	return inactionTargetUsers, nil
}

func findInactionTargetUsers(inactionCampaignUsersAtSecondEventMap map[int64]bool, totalInactionCampaignUsersInDB []*common.ActorID) []*common.ActorID {
	// Lets Suppose A = inactionCampaignUsersAtFirstEvent
	// B = inactionCampaignUsersAtSecondEvent
	// Inaction Target Users -->  A - (A common B)
	// Converted Users -->  (A common B)
	var inactionTargetUsers []*common.ActorID
	for _, user := range totalInactionCampaignUsersInDB {
		if _, isCommonActor := inactionCampaignUsersAtSecondEventMap[user.ActorId]; !isCommonActor {
			inactionTargetUsers = append(inactionTargetUsers, &common.ActorID{
				ActorType: user.ActorType,
				ActorId:   user.ActorId,
			})
		}
	}
	return inactionTargetUsers
}

func fetchInactionCampaignUsersData(ctx context.Context, request *fs.CampaignRequest) ([]*common.ActorID, *fs.CampaignResponse) {
	var err error
	findInactionTargetUserByCampaignIdRequest := &fs.FindInactionTargetUserByCampaignIdRequest{
		//Set your request here
		CampaignId: request.CampaignId,
	}
	response := FindInactionTargetUserByCampaignIdExecutor.OnRequest(ctx, findInactionTargetUserByCampaignIdRequest)
	if response != nil {
		logger.Error("FindInactionTargetUserByCampaignId request failed", zap.Any("response ", response))
		err = fmt.Errorf("FindInactionTargetUserByCampaignId_ERROR")
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.UpdateCampaign_Error_Metrics, err, ctx)
		return nil, setCampaignResponse(common.RequestStatus_BAD_INPUT, nil, CONST_ERROR_CODE_NIL)
	}
	findInactionTargetUserByCampaignIdResponse, err := executor.RequestExecutor.ExecuteFindInactionTargetUserByCampaignId(ctx, findInactionTargetUserByCampaignIdRequest)
	if err != nil {
		logger.Error("FindInactionTargetUserByCampaignId request failed", zap.Any("response ", response))
		err = fmt.Errorf("FindInactionTargetUserByCampaignId_ERROR")
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.UpdateCampaign_Error_Metrics, err, ctx)
		return nil, setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, []string{"Error while finding inaction target user by Id"}, common.ErrorCode_DATABASE_ERROR)
	}

	if findInactionTargetUserByCampaignIdResponse == nil || findInactionTargetUserByCampaignIdResponse.Records == nil {
		logger.Error("FindInactionTargetUserByCampaignId request, no records found for the CampaignId : ", zap.Any("campaignId ", request.CampaignId))
	}
	var totalInactionCampaignUsers []*common.ActorID
	for _, record := range findInactionTargetUserByCampaignIdResponse.Records {
		User := &common.ActorID{
			ActorId:   record.UserId,
			ActorType: common.ActorType(common.ActorType_value[record.UserType]),
		}
		totalInactionCampaignUsers = append(totalInactionCampaignUsers, User)
	}
	return totalInactionCampaignUsers, nil
}

func GetPlaceholderData(val []string, indexFieldMap map[int64]string) []*ce.Placeholder {

	var placeholderDataList []*ce.Placeholder

	for fieldIndex, fieldValue := range val {
		placeholderData := &ce.Placeholder{
			Key:   indexFieldMap[int64(fieldIndex)],
			Value: fieldValue,
		}
		if placeholderData.Key != CONST_ID && placeholderData.Key != CONST_APP_ID && placeholderData.Key != CONST_FCM_TOKEN && placeholderData.Key != CONST_EMAIL_ID && placeholderData.Key != CONST_LANG_KEY && placeholderData.Key != CONST_MOBILE_NUMBER && placeholderData.Key != CONST_MODIFIED_ON {
			placeholderDataList = append(placeholderDataList, placeholderData)
		}
	}

	return placeholderDataList
}

func GetImageMetaData(languageCode common.LanguageCode, imageMap map[common.LanguageCode]string) string {

	if value, ok := imageMap[languageCode]; ok {
		return value
	}
	if value, ok := imageMap[common.LanguageCode_NO_LANGUAGE_CODE]; ok {
		return value
	}
	return ""
}

func GetPlaceholderIndexFieldMap(val []string) map[int64]string {
	indexFieldMap := make(map[int64]string)
	for fieldIndex, fieldName := range val {
		indexFieldMap[int64(fieldIndex)] = fieldName
	}
	return indexFieldMap
}

func getOuterQuery(channel common.CommunicationChannel, campaignType common.CampaignQueryType) (outerQuery string) {
	if campaignType == common.CampaignQueryType_ATHENA || campaignType == common.CampaignQueryType_USER_JOURNEY {
		switch channel {
		case common.CommunicationChannel_SMS:
			outerQuery = query.QUERY_ATHENA_SMS_ATTRIBUTES
		case common.CommunicationChannel_APP_NOTIFICATION:
			outerQuery = query.QUERY_ATHENA_APP_NOTIFICATION_ATTRIBUTES
		case common.CommunicationChannel_WHATSAPP:
			outerQuery = query.QUERY_ATHENA_WHATSAPP_ATTRIBUTES
		case common.CommunicationChannel_EMAIL:
			outerQuery = query.QUERY_ATHENA_EMAIL_ATTRIBUTES
		}
	} else if campaignType == common.CampaignQueryType_USER_LIST {
		switch channel {
		case common.CommunicationChannel_SMS:
			outerQuery = query.QUERY_USERLIST_SMS_ATTRIBUTES
		case common.CommunicationChannel_APP_NOTIFICATION:
			outerQuery = query.QUERY_USERLIST_APP_NOTIFICATION_ATTRIBUTES
		case common.CommunicationChannel_WHATSAPP:
			outerQuery = query.QUERY_USERLIST_WHATSAPP_ATTRIBUTES
		case common.CommunicationChannel_EMAIL:
			outerQuery = query.QUERY_USERLIST_EMAIL_ATTRIBUTES
		}
	}
	return
}

func getCampaignOuterQuery(channel common.CommunicationChannel, campaignType common.CampaignQueryType) (outerQuery string) {
	outerQuery = getOuterQuery(channel, campaignType)
	return
}

func executeGetUserAthenaQuery(ctx context.Context, query string, channel common.CommunicationChannel, campaignId int64) (map[string]string, [][]string, error) {
	actorDetailsColumnType, actorDetailsData, err := aws.ExecuteAthenaQuery(ctx, query, campaignId)
	logger.Info("executeGetUserAthenaQuery method, Total number of records received from Athena:",
		zap.Any("actorDetailsData", len(actorDetailsData)))
	if err != nil {
		logger.Error("Error in getting ActorDetails from athena query", zap.Any("query", query), zap.Any("error", err))
		return nil, nil, err
	}
	appId := CONST_DEFAULT_AP_ID
	actorType := common.ActorType_NO_ACTOR
	if channel == common.CommunicationChannel_APP_NOTIFICATION {
		for i, _ := range actorDetailsData {
			if i == 0 {
				actorDetailsData[i] = append(actorDetailsData[i], CONST_APP_ID)
			} else {
				actorDetailsData[i] = append(actorDetailsData[i], appId)
			}
		}
	}

	for i, _ := range actorDetailsData {
		if i == 0 {
			actorDetailsData[i] = append(actorDetailsData[i], CONST_ACTOR_TYPE)
		} else {
			actorDetailsData[i] = append(actorDetailsData[i], actorType.String())
		}
	}

	return actorDetailsColumnType, actorDetailsData, nil
}

func updateCampaignStatus(ctx context.Context, campaignId int64, campaignStatus common.CampaignStatus) *fs.CampaignResponse {
	updateRequest := &fs.UpdateCampaignRequest{
		//Set your request here
		Id: campaignId,
		UpdatedByActor: &common.ActorID{
			ActorId:   CONST_ACTOR_ID_SYSTEM,
			ActorType: common.ActorType_SYSTEM,
		},
		AddCampaignRequest: &fs.AddCampaignRequest{
			Status: campaignStatus,
		},
	}
	updateResponse, _ := UpdateCampaignExecutor.OnRequest(ctx, updateRequest)
	if updateResponse.Status.Status == common.RequestStatus_INTERNAL_ERROR {
		logger.Error("ExecuteUpdateCampaign request failed", zap.Any("response ", updateResponse))
		err := fmt.Errorf("ExecuteUpdateCampaign_ERROR")
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.UpdateCampaign_Error_Metrics, err, ctx)
		return setCampaignResponse(common.RequestStatus_INTERNAL_ERROR, nil, common.ErrorCode_DATABASE_ERROR)
	}
	return nil
}

func (ge *GenericCampaignExecutor) OnResponse(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse {
	return ge.CampaignInterface.OnResponse(ctx, request, response)
}

func (ge *GenericCampaignExecutor) OnData(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse {
	return ge.CampaignInterface.OnData(ctx, request, response)
}

func (ge *GenericCampaignExecutor) OnError(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse, err error) *fs.CampaignResponse {
	return ge.CampaignInterface.OnError(ctx, request, response, err)
}

func (rc *CampaignController) OnRequest(ctx context.Context, request *fs.CampaignRequest) *fs.CampaignResponse {
	return nil
}

func (rc *CampaignController) OnResponse(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse {
	return nil
}

func (rc *CampaignController) OnData(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse) *fs.CampaignResponse {
	return nil
}

func (rc *CampaignController) OnError(ctx context.Context, request *fs.CampaignRequest, response *fs.CampaignResponse, err error) *fs.CampaignResponse {
	return nil
}
