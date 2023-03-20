package mappers

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	ce "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/database/models"
	"code.nurture.farm/platform/CampaignService/zerotouch/golang/metrics"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/opentracing/opentracing-go/log"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
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

const (
	CONST_MEDIA_TYPE            = "media_type"
	CONST_MEDIA_ACCESS_TYPE     = "media_access_type"
	CONST_MEDIA_INFO            = "media_info"
	CONST_DOCUMENT_NAME         = "document_name"
	CONST_MSG                   = "msg"
	CONST_DEFAULT_DOCUMENT_NAME = "file"
	CONST_MEDIA_URL             = "mediaUrl"
)

type Attributes struct {
	ContentMetaData        []ContentMetaData           `json:"content_metadata,omitempty"`
	MediaInfo              []ContentMetaData           `json:"media,omitempty"`
	ChannelAttributes      string                      `json:"channel_attributes,omitempty"`
	MetaData               UserJourneyCampaignMetadata `json:"meta_data,omitempty""`
	ControlGroupPercentage int32                       `json:"control_group_percentage,omitempty""`
	UserMetaDataList       []*fs.UserMetadata          `json:"user_metadata_list,omitempty"`
}

type TargetUserAttributes struct {
	PlaceHolders []PlaceHolder `json:"placeHolders,omitempty"`
}

type PlaceHolder struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type UserJourneyCampaignMetadata struct {
	UserJourneyMetadata string `json:"user_journey_metadata,omitempty"`
	EngagementMetadata  string `json:"engagement_metadata,omitempty"`
}

type ContentMetaData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetTargetUserAttributes(attribs []*common.Attribs) sql.NullString {

	targetUserAttributes := TargetUserAttributes{}
	for _, attrib := range attribs {
		targetUserAttributes.PlaceHolders = append(targetUserAttributes.PlaceHolders, PlaceHolder{
			Key:   attrib.Key,
			Value: attrib.Value,
		})
	}
	attributesBytes, err := json.Marshal(targetUserAttributes)
	if err != nil {
		logger.Error("Error while Marshalling targetUser attributes", zap.Error(err), zap.Any("attribs", attribs))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_MARSHAL_Error_Metrics, err, context.Background())
		return sql.NullString{
			Valid: false,
		}
	}
	return GetNullableString(string(attributesBytes))
}
func GetBitSetOfBloomFilter(bloomFilter *bloom.BloomFilter) sql.RawBytes {
	s := bloomFilter.BitSet().Bytes()
	cr, err := json.Marshal(s)
	if err != nil {
		log.Error(err)
	}
	return cr
}
func GetControlGroupUserAttributes(controlGroupPercentage int, bloomFilter *bloom.BloomFilter) string {
	controlGroupUserAttributes := make(map[string]string)
	controlGroupUserAttributes["control_group_percentage"] = cast.ToString(controlGroupPercentage)
	controlGroupUserAttributes["capacity"] = cast.ToString(bloomFilter.Cap())
	controlGroupUserAttributes["k"] = cast.ToString(bloomFilter.K())
	attributesBytes, err := json.Marshal(controlGroupUserAttributes)
	if err != nil {
		logger.Error("Error while Marshalling control group attributes", zap.Error(err), zap.Any("attribs", controlGroupUserAttributes))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_MARSHAL_Error_Metrics, err, context.Background())
		return ""
	}
	return string(attributesBytes)
}
func GetAttributes(request *fs.AddCampaignRequest, userJourneyMetadata string, engagementMetadata string, userMetadataList []*fs.UserMetadata) sql.NullString {

	attributes := Attributes{}
	for _, contentMetadata := range request.ContentMetadata {
		placeHolder := ContentMetaData{
			Key:   contentMetadata.Key,
			Value: contentMetadata.Value,
		}
		if contentMetadata.Key == "cta" {
			attributes.ContentMetaData = append(attributes.ContentMetaData, ContentMetaData{
				Key:   "message",
				Value: contentMetadata.Value,
			})
			attributes.ContentMetaData = append(attributes.ContentMetaData, ContentMetaData{
				Key:   "link",
				Value: contentMetadata.Value,
			})
		}
		attributes.ContentMetaData = append(attributes.ContentMetaData, placeHolder)
	}
	if request.Media != nil {
		placeHolder := ContentMetaData{
			Key:   CONST_MEDIA_TYPE,
			Value: request.Media.MediaType.String(),
		}
		attributes.MediaInfo = append(attributes.MediaInfo, placeHolder)
		placeHolder = ContentMetaData{
			Key:   CONST_MEDIA_ACCESS_TYPE,
			Value: request.Media.MediaAccessType.String(),
		}
		attributes.MediaInfo = append(attributes.MediaInfo, placeHolder)
		placeHolder = ContentMetaData{
			Key:   CONST_MEDIA_INFO,
			Value: request.Media.MediaInfo,
		}
		attributes.MediaInfo = append(attributes.MediaInfo, placeHolder)

		docName := strings.Split(request.Media.MediaInfo, "/")
		filterDocName := strings.Split(docName[len(docName)-1], ".")
		documentName := filterDocName[0]
		if request.Media.MediaType == common.MediaType_DOCUMENT {
			request.Media.DocumentName = documentName
		}

		placeHolder = ContentMetaData{
			Key:   CONST_DOCUMENT_NAME,
			Value: request.Media.DocumentName,
		}
		attributes.MediaInfo = append(attributes.MediaInfo, placeHolder)
		placeHolder = ContentMetaData{
			Key:   CONST_MSG,
			Value: request.Media.Msg,
		}
		attributes.MediaInfo = append(attributes.MediaInfo, placeHolder)
	}
	if request.ChannelAttributes != nil {
		attributes.ChannelAttributes = request.ChannelAttributes.PushNotificationType.String()
	}
	if userJourneyMetadata != "" || engagementMetadata != "" || userMetadataList != nil {
		attributes.MetaData = UserJourneyCampaignMetadata{}
		if userJourneyMetadata != "" {
			attributes.MetaData.UserJourneyMetadata = userJourneyMetadata
		}
		if engagementMetadata != "" {
			attributes.MetaData.EngagementMetadata = engagementMetadata
		}
		if userMetadataList != nil {
			attributes.UserMetaDataList = userMetadataList
		}
	}
	if request.ControlGroupPercentage != 0 {
		attributes.ControlGroupPercentage = request.ControlGroupPercentage
	}
	attributesBytes, err := json.Marshal(attributes)
	if err != nil {
		logger.Error("Error while Marshalling attributes", zap.Error(err), zap.Any("request", request))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_MARSHAL_Error_Metrics, err, context.Background())
		return sql.NullString{
			Valid: false,
		}
	}

	return GetNullableString(string(attributesBytes))
}

func MapContentMetaData(AttributeList string) []*common.Attribs {

	attributes := Attributes{}
	err := json.Unmarshal([]byte(AttributeList), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes, contentMetaData", zap.Error(err), zap.Any("AttributeList", AttributeList))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return nil
	}
	contentMetData := []*common.Attribs{}
	for _, attribute := range attributes.ContentMetaData {
		contentMetData = append(contentMetData, &common.Attribs{Key: attribute.Key, Value: attribute.Value})
	}
	return contentMetData
}

func MapControlGroupAttributes(AttributeList string) map[string]string {
	controlGroupUserAttributes := make(map[string]string)
	err := json.Unmarshal([]byte(AttributeList), &controlGroupUserAttributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes, ControlGroupAttributes ", zap.Error(err), zap.Any("AttributeList", AttributeList))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return nil
	}

	return controlGroupUserAttributes

}

func MapMedia(AttributeList string) *ce.Media {

	logger.Info("Calling map media")
	attributes := Attributes{}
	err := json.Unmarshal([]byte(AttributeList), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes, media", zap.Error(err), zap.Any("AttributeList", AttributeList))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return nil
	}
	if attributes.MediaInfo != nil {
		media := &ce.Media{}
		documentName := CONST_DEFAULT_DOCUMENT_NAME
		for _, attribute := range attributes.MediaInfo {
			if attribute.Key == CONST_MEDIA_TYPE {
				media.MediaType = common.MediaType(common.MediaType_value[attribute.Value])
			} else if attribute.Key == CONST_MEDIA_ACCESS_TYPE {
				media.MediaAccessType = common.MediaAccessType(common.MediaAccessType_value[attribute.Value])
			} else if attribute.Key == CONST_MEDIA_INFO {
				media.MediaInfo = attribute.Value
				docName := strings.Split(attribute.Value, "/")
				filterDocName := strings.Split(docName[len(docName)-1], ".")
				documentName = filterDocName[0]
			} else if attribute.Key == CONST_DOCUMENT_NAME {
				media.DocumentName = attribute.Value
			} else if attribute.Key == CONST_MSG {
				media.Msg = attribute.Value
			} else if attribute.Key == CONST_MEDIA_URL {
				media.MediaAccessType = common.MediaAccessType_PUBLIC_URL
				media.MediaType = common.MediaType_IMAGE
				media.MediaInfo = attribute.Value
			}
		}
		if media.MediaType == common.MediaType_DOCUMENT {
			media.DocumentName = documentName
		}
		logger.Info("Media is {} ", zap.Any("media", media))
		return media
	}
	return nil
}

func MapChannelAttributes(channelAttributes string) *ce.CommunicationChannelAttributes {
	attributes := Attributes{}
	err := json.Unmarshal([]byte(channelAttributes), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling channelAttributes", zap.Error(err), zap.Any("channelAttributes", channelAttributes))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return nil
	}
	if attributes.ChannelAttributes != "" {
		return &ce.CommunicationChannelAttributes{PushNotificationType: common.PushNotificationType(common.PushNotificationType_value[attributes.ChannelAttributes])}
	}
	return nil
}

func MapControlGroupPercentage(channelAttributes string) int32 {
	attributes := Attributes{}
	err := json.Unmarshal([]byte(channelAttributes), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling channelAttributes", zap.Error(err), zap.Any("channelAttributes", channelAttributes))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return 0
	}
	return attributes.ControlGroupPercentage
}

func MapMetaData(AttributeList string) (string, string, []*fs.UserMetadata) {

	attributes := Attributes{}
	err := json.Unmarshal([]byte(AttributeList), &attributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes metaData", zap.Error(err), zap.Any("AttributeList", AttributeList))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return "", "", nil
	}
	return attributes.MetaData.UserJourneyMetadata, attributes.MetaData.EngagementMetadata, attributes.UserMetaDataList
}

func MapPlaceholders(attributes string) []*common.Attribs {

	attribs := []*common.Attribs{}
	targetUserAttributes := TargetUserAttributes{}
	err := json.Unmarshal([]byte(attributes), &targetUserAttributes)
	if err != nil {
		logger.Error("Error while Unmarshalling attributes for targetUserAttributes", zap.Error(err), zap.Any("attributes", attributes))
		metrics.Metrics.PushToErrorCounterMetrics()(metrics.ATTRIBUTES_UNMARSHAL_Error_Metrics, err, context.Background())
		return attribs
	}
	for _, attrib := range targetUserAttributes.PlaceHolders {
		attribs = append(attribs, &common.Attribs{
			Key:   attrib.Key,
			Value: attrib.Value,
		})
	}
	return attribs
}

func MakeAddCampaignRequestVO(request *fs.AddCampaignRequest, userJourneyMetadata string, engagementMetadata string, userMetadataList []*fs.UserMetadata) *models.AddCampaignRequestVO {

	return &models.AddCampaignRequestVO{
		Namespace:            GetNullableString(request.Namespace.String()),
		Name:                 GetNullableString(request.Name),
		Description:          GetNullableString(request.Description),
		CronExpression:       GetNullableString(request.CronExpression),
		Occurrences:          GetNullableInt32(request.Occurrences),
		CommunicationChannel: GetNullableString(request.CommunicationChannel.String()),
		Status:               GetNullableString(request.Status.String()),
		Type:                 GetNullableString(request.Type.String()),
		ScheduleType:         GetNullableString(request.CampaignScheduleType.String()),
		Query:                GetNullableString(request.Query),
		InactionQuery:        GetNullableString(request.InactionQuery),
		InactionDuration:     GetNullableDuration(request.InactionDuration),
		Attributes:           GetAttributes(request, userJourneyMetadata, engagementMetadata, userMetadataList),
		CreatedByActorid:     GetNullableInt64(request.CreatedByActor.ActorId),
		CreatedByActortype:   GetNullableString(request.CreatedByActor.ActorType.String()),
	}

}
func MakeUpdateCampaignRequestVO(request *fs.UpdateCampaignRequest) *models.UpdateCampaignRequestVO {

	return &models.UpdateCampaignRequestVO{
		Name:                 GetNullableString(request.AddCampaignRequest.Name),
		CronExpression:       GetNullableString(request.AddCampaignRequest.CronExpression),
		Status:               GetNullableString(request.AddCampaignRequest.Status.String()),
		Query:                GetNullableString(request.AddCampaignRequest.Query),
		Occurrences:          GetNullableInt32(request.AddCampaignRequest.Occurrences),
		UpdatedByActorid:     GetNullableInt64(request.UpdatedByActor.ActorId),
		UpdatedByActortype:   GetNullableString(request.UpdatedByActor.ActorType.String()),
		Id:                   GetNullableInt64(request.Id),
		Namespace:            GetNullableString(request.AddCampaignRequest.Namespace.String()),
		Description:          GetNullableString(request.AddCampaignRequest.Description),
		CommunicationChannel: GetNullableString(request.AddCampaignRequest.CommunicationChannel.String()),
		Type:                 GetNullableString(request.AddCampaignRequest.Type.String()),
		ScheduleType:         GetNullableString(request.AddCampaignRequest.CampaignScheduleType.String()),
		InactionQuery:        GetNullableString(request.AddCampaignRequest.InactionQuery),
		InactionDuration:     GetNullableDuration(request.AddCampaignRequest.InactionDuration),
		Attributes:           GetAttributes(request.AddCampaignRequest, "", "", nil),
	}

}
func MakeAddCampaignTemplateRequestVO(request *fs.AddCampaignTemplateRequest) *models.AddCampaignTemplateRequestVO {

	return &models.AddCampaignTemplateRequestVO{
		CampaignId:          GetNullableInt64(request.CampaignId),
		TemplateName:        GetNullableString(request.TemplateName),
		CampaignName:        GetNullableString(request.CampaignName),
		DistributionPercent: GetNullableInt32(request.DistributionPercent),
	}

}
func MakeCampaignTemplateRequestVO(campaignId int64) *models.DeleteCampaignTemplateRequestVO {

	return &models.DeleteCampaignTemplateRequestVO{
		CampaignId: GetNullableInt64(campaignId),
	}
}
func MakeFindCampaignByIdResponseVO(model *models.FindCampaignByIdResponseVO) *fs.FindCampaignByIdResponseRecord {

	return &fs.FindCampaignByIdResponseRecord{
		Id:                   model.Id.Int64,
		Namespace:            model.Namespace.String,
		Name:                 model.Name.String,
		Description:          model.Description.String,
		CronExpression:       model.CronExpression.String,
		Occurrences:          model.Occurrences.Int32,
		CommunicationChannel: model.CommunicationChannel.String,
		Status:               model.Status.String,
		Type:                 model.Type.String,
		ScheduleType:         model.ScheduleType.String,
		Query:                model.Query.String,
		InactionQuery:        model.InactionQuery.String,
		InactionDuration:     model.InactionDuration.Int64,
		Attributes:           model.Attributes.String,
		CreatedByActorid:     model.CreatedByActorid.Int64,
		CreatedByActortype:   model.CreatedByActortype.String,
		UpdatedByActorid:     model.UpdatedByActorid.Int64,
		UpdatedByActortype:   model.UpdatedByActortype.String,
		Version:              model.Version.Int64,
		CreatedAt:            GetUnixTime(model.CreatedAt.String),
		UpdatedAt:            GetUnixTime(model.UpdatedAt.String),
		DeletedAt:            GetUnixTime(model.DeletedAt.String),
	}

}
func MakeFindControlGroupByCampaignIdResponse(model *models.FindControlGroupByCampaignIdRequestV0) *fs.FindControlGroupByCampaignIdResponseRecord {

	return &fs.FindControlGroupByCampaignIdResponseRecord{
		Id:          model.Id.Int64,
		CampaignId:  model.CampaignId.Int64,
		Attributes:  model.Attributes.String,
		BloomFilter: model.BloomFilter,
	}

}
func MakeFindCampaignTemplateByIdResponseVO(model *models.FindCampaignTemplateByIdResponseVO) *fs.FindCampaignTemplateByIdResponseRecord {

	return &fs.FindCampaignTemplateByIdResponseRecord{
		Id:                  model.Id.Int64,
		CampaignId:          model.CampaignId.Int64,
		TemplateName:        model.TemplateName.String,
		CampaignName:        model.CampaignName.String,
		DistributionPercent: model.DistributionPercent.Int32,
	}

}
func MakeAddTargetUserRequestVO(request *fs.AddTargetUserRequest) *models.AddTargetUserRequestVO {

	return &models.AddTargetUserRequestVO{
		CampaignId: GetNullableInt64(request.CampaignId),
		UserId:     GetNullableInt64(request.User.ActorId),
		UserType:   GetNullableString(request.User.ActorType.String()),
		Attributes: GetTargetUserAttributes(request.Attribs),
	}

}
func MakeAddControlGroupRequestVO(campaignId int, attributes string, bloomFilter []byte) *models.AddControlGroupRequestVO {
	return &models.AddControlGroupRequestVO{
		CampaignId:  GetNullableInt64(int64(campaignId)),
		Attributes:  GetNullableString(attributes),
		BloomFilter: bloomFilter,
	}
}
func MakeFindTargetUserByIdResponseVO(model *models.FindTargetUserByIdResponseVO) *fs.FindTargetUserByIdResponseRecord {

	return &fs.FindTargetUserByIdResponseRecord{
		Id:         model.Id.Int64,
		CampaignId: model.CampaignId.Int64,
		UserId:     model.UserId.Int64,
		UserType:   model.UserType.String,
		Attribs:    MapPlaceholders(model.Attributes.String),
	}
}

func MakeAddInactionTargetUserRequestVO(request *fs.AddInactionTargetUserRequest) *models.AddInactionTargetUserRequestVO {
	return &models.AddInactionTargetUserRequestVO{
		CampaignId: GetNullableInt64(request.CampaignId),
		UserId:     GetNullableInt64(request.User.ActorId),
		UserType:   GetNullableString(request.User.ActorType.String()),
	}
}

func MakeFindInactionTargetUserByCampaignIdResponseVO(model *models.FindInactionTargetUserByCampaignIdResponseVO) *fs.FindInactionTargetUserByCampaignIdResponseRecord {

	return &fs.FindInactionTargetUserByCampaignIdResponseRecord{
		Id:         model.Id.Int64,
		CampaignId: model.CampaignId.Int64,
		UserId:     model.UserId.Int64,
		UserType:   model.UserType.String,
	}

}

func MakeGetDynamicDataByKeyResponseVO(model *models.GetDynamicDataByKeyResponseVO) *fs.GetDynamicDataByKeyResponseRecord {

	return &fs.GetDynamicDataByKeyResponseRecord{
		CampaignId: model.CampaignId.Int64,
		DynamicKey: model.DynamicKey.String,
		CtaLink:    model.CtaLink.String,
		Media:      model.Media.String,
	}

}

func MakeAddDynamicDataRequestVO(request *fs.AddDynamicDataRequest) *models.AddDynamicDataRequestVO {

	return &models.AddDynamicDataRequestVO{
		CampaignId: GetNullableInt64(request.CampaignId),
		DynamicKey: GetNullableString(request.DynamicKey),
		CtaLink:    GetNullableString(request.CtaLink),
		Media:      GetNullableString(request.Media),
	}

}

func MakeFindQueryCampaignResponseVO(model *models.FindQueryCampaignResponseVO) *fs.FindQueryCampaignResponseRecord {

	return &fs.FindQueryCampaignResponseRecord{
		Name:  model.Name.String,
		Query: model.Query.String,
	}

}
func MakeFindQueryCampaignRequestVO(request *fs.FindQueryCampaignRequest) *models.FindQueryCampaignRequestVO {

	return &models.FindQueryCampaignRequestVO{
		Type: GetNullableString(request.Type),
	}

}
func MakeAddQueryCampaignRequestVO(request *fs.AddQueryCampaignRequest) *models.AddQueryCampaignRequestVO {

	return &models.AddQueryCampaignRequestVO{
		Name:      GetNullableString(request.Name),
		Type:      GetNullableString(request.Type),
		Query:     GetNullableString(request.Query),
		UpdatedBy: GetNullableString(request.UpdatedBy),
	}

}
func GetNullableInt32(nullableInt int32) sql.NullInt32 {

	var result sql.NullInt32
	if nullableInt != 0 {
		result = sql.NullInt32{Int32: nullableInt, Valid: true}
	} else {
		result = sql.NullInt32{}
	}
	return result
}

func GetNullableInt32s(nullableInts []int32) []sql.NullInt32 {

	var result []sql.NullInt32
	if len(nullableInts) > 0 {
		for _, nullableInt := range nullableInts {
			result = append(result, sql.NullInt32{Int32: nullableInt, Valid: true})
		}
	} else {
		result = []sql.NullInt32{}
	}
	return result
}

func GetNullableInt64(nullableInt int64) sql.NullInt64 {

	var result sql.NullInt64
	if nullableInt != 0 {
		result = sql.NullInt64{Int64: nullableInt, Valid: true}
	} else {
		result = sql.NullInt64{}
	}
	return result
}

func GetNullableInt64s(nullableInts []int64) []sql.NullInt64 {

	var result []sql.NullInt64
	if len(nullableInts) > 0 {
		for _, nullableInt := range nullableInts {
			result = append(result, sql.NullInt64{Int64: nullableInt, Valid: true})
		}
	} else {
		result = []sql.NullInt64{}
	}
	return result
}

func GetNullableFloat64(nullableFloat float64) sql.NullFloat64 {

	var result sql.NullFloat64
	if nullableFloat != 0.0 {
		result = sql.NullFloat64{Float64: nullableFloat, Valid: true}
	} else {
		result = sql.NullFloat64{}
	}
	return result
}

func GetNullableString(nullableString string) sql.NullString {

	var result sql.NullString
	if len(nullableString) > 0 {
		result = sql.NullString{String: nullableString, Valid: true}
	} else {
		result = sql.NullString{}
	}
	return result
}

func GetNullableStrings(nullableStrings []string) []sql.NullString {

	var result []sql.NullString
	if len(nullableStrings) > 0 {
		for _, nullableString := range nullableStrings {
			result = append(result, sql.NullString{String: nullableString, Valid: true})
		}
	} else {
		result = []sql.NullString{}
	}
	return result
}

func GetNullableDateTime(timeStamp int64) sql.NullTime {

	gTime := time.Unix(timeStamp, 0)
	pTime, _ := ptypes.TimestampProto(gTime)

	var result sql.NullTime
	parsedAllottedTime, err := ptypes.Timestamp(pTime)
	if err != nil {
		result = sql.NullTime{}
	} else {
		result = sql.NullTime{Time: parsedAllottedTime, Valid: true}
	}
	return result
}

func GetNullableTimestamp(timeStamp int64) sql.NullString {

	gTime := time.Unix(timeStamp, 0)
	pTime, _ := ptypes.TimestampProto(gTime)

	var result sql.NullString
	parsedAllottedTime, err := ptypes.Timestamp(pTime)
	if err != nil {
		result = sql.NullString{}
	} else if parsedAllottedTime.IsZero() {
		result = sql.NullString{}
	} else {
		result = sql.NullString{String: parsedAllottedTime.Format("2006-01-02 15:04:05"), Valid: true}
	}
	return result
}

func GetNullableBool(boolean bool) sql.NullBool {

	var result sql.NullBool
	result = sql.NullBool{Bool: boolean, Valid: true}
	return result
}

func GetNullableDuration(nullableTimeStamp *duration.Duration) sql.NullInt64 {

	var result sql.NullInt64
	parsedAllottedTime, err := ptypes.Duration(nullableTimeStamp)
	if err != nil {
		result = sql.NullInt64{}
	} else {
		result = sql.NullInt64{Int64: cast.ToInt64(parsedAllottedTime.Seconds()), Valid: true}
	}
	return result
}

func GetNullableTimestampFromProtoTime(nullableTimeStamp *timestamp.Timestamp) sql.NullTime {

	var result sql.NullTime
	parsedAllottedTime, err := ptypes.Timestamp(nullableTimeStamp)
	if err != nil {
		result = sql.NullTime{}
	} else {
		result = sql.NullTime{Time: parsedAllottedTime, Valid: true}
	}
	return result
}

func GetUnixTime(timestamp string) int64 {

	layout := "2006-01-02T15:04:05Z"
	time, _ := time.Parse(layout, timestamp)
	return time.Unix()
}
