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

package query

import (
	fs "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"github.com/spf13/cast"
	"time"
)

const (
	QUERY_AddCampaign = "INSERT into campaigns" +
		"(" +
		"namespace,name,description,cron_expression," +
		"occurrences,communication_channel,status,type,schedule_type," +
		"query,inaction_query,inaction_duration,attributes," +
		"created_by_actorId,created_by_actorType" +
		") " +
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	QUERY_UpdateCampaign = "UPDATE campaigns SET " +
		"name = ?, cron_expression = ?, status = ?, " +
		"query = ?,  namespace=?, occurrences = ?, " +
		"communication_channel=?, type=?, description=?, " +
		"schedule_type=?, inaction_query=?, inaction_duration =?, " +
		"attributes=?, updated_by_actorId = ?, updated_by_actorType = ?, " +
		"version = version + 1 WHERE id = ?"
	QUERY_AddCampaignTemplate = "INSERT into campaign_templates" +
		"(" +
		"campaign_id,template_name," +
		"campaign_name,distribution_percent" +
		") " +
		"values (?,?,?,?)"
	QUERY_DeleteCampaignTemplate = "DELETE FROM campaign_templates " +
		"WHERE campaign_id = ?"
	QUERY_AddNewCampaign   = ""
	QUERY_GetUserList      = ""
	QUERY_FindCampaignById = "SELECT id,namespace,name,description,cron_expression," +
		"occurrences,communication_channel,status,type,schedule_type, " +
		"query, inaction_query, inaction_duration, attributes, created_by_actorId," +
		"created_by_actorType,updated_by_actorId,updated_by_actorType,version," +
		"created_at,updated_at,deleted_at " +
		"FROM campaigns " +
		"WHERE id = ?"
	QUERY_FindControlGroupByCampaignId = "SELECT id,campaign_id,attributes,bloom_filter " +
		"FROM control_group " +
		"WHERE campaign_id = ?"
	QUERY_FindCampaignTemplateById = "SELECT id,campaign_id,template_name,campaign_name,distribution_percent " +
		"FROM campaign_templates " +
		"WHERE campaign_id = ?"
	QUERY_AddTargetUser = "INSERT into target_users" +
		"(" +
		"campaign_id,user_id,user_type,attributes" +
		") " +
		"values (?,?,?,?)"
	QUERY_AddControlGroup = "INSERT into control_group" +
		"(" +
		"campaign_id,attributes,bloom_filter" +
		") " +
		"values (?,?,?)"
	QUERY_FindTargetUserById = "SELECT id,campaign_id,user_id,user_type,attributes " +
		"FROM target_users " +
		"WHERE campaign_id = ?"

	QUERY_AddInactionTargetUser = "INSERT into inaction_target_users" +
		"(" +
		"campaign_id, user_id, user_type" +
		") " +
		"values(?,?,?)"
	QUERY_FindInactionTargetUserByCampaignId = "SELECT id,campaign_id,user_id,user_type " +
		"FROM inaction_target_users " +
		"WHERE campaign_id = ?"
	QUERY_FindQueryCampaign                             = "SELECT name, query FROM campaign_query WHERE type = ?;"
	QUERY_AddQueryCampaign                              = "INSERT INTO campaign_query(name, type, query, updated_by) VALUES(?,?,?,?);"
	QUERY_AddUserJourneyVertex                          = "INSERT INTO user_journey_vertices(campaign_id,event_type,event_name,inaction_duration,inaction_event_name,attributes) VALUES(?,?,?,?,?,?)"
	QUERY_AddEdge                                       = "INSERT INTO edges(campaign_id,vertex_type,from_vertex_id,to_vertex_id,wait_duration,wait_time,wait_type,message_delivery_states) VALUES(?,?,?,?,?,?,?,?)"
	QUERY_AddEngagementVertex                           = "INSERT INTO engagement_vertices(campaign_id,template_name,attributes,athena_query,channel) VALUES(?,?,?,?,?)"
	QUERY_FindUserJourneyVerticesByCampaignId           = "SELECT id,campaign_id,event_type,event_name,inaction_duration,inaction_event_name,version,attributes,created_at,updated_at,deleted_at FROM user_journey_vertices WHERE campaign_id = ? AND deleted_at is NULL"
	QUERY_FindEdgesByCampaignId                         = "SELECT id,campaign_id,vertex_type,from_vertex_id,to_vertex_id,wait_duration,wait_time,wait_type,message_delivery_states,version,created_at,updated_at,deleted_at FROM edges WHERE campaign_id = ? AND vertex_type = ? AND deleted_at is NULL"
	QUERY_FindEngagementVerticesByCampaignId            = "SELECT id,campaign_id,template_name,attributes,athena_query,channel,version,created_at,updated_at,deleted_at FROM engagement_vertices WHERE campaign_id = ? AND deleted_at is NULL"
	QUERY_FindEngagementVerticesById                    = "SELECT id,campaign_id,template_name,attributes,athena_query,channel,version,created_at,updated_at,deleted_at FROM engagement_vertices WHERE id = ? AND deleted_at is NULL"
	QUERY_AddUserJourneyTargetUsers                     = "INSERT INTO user_journey_target_users(campaign_id,engagement_vertex_id,reference_id,event_reference_id,actor_contact_id,status) VALUES(?,?,?,?,?,?)"
	QUERY_FindStartEngagementVertexId                   = "SELECT id FROM engagement_vertices WHERE campaign_id = ? AND deleted_at is NULL ORDER BY id LIMIT 1"
	QUERY_FindPreviosEngagementVertexId                 = "SELECT from_vertex_id,message_delivery_states FROM edges WHERE campaign_id = ? AND to_vertex_id = ? AND vertex_type = 'ENGAGEMENT' AND deleted_at is NULL"
	QUERY_FindNextEngaagementVertexId                   = "SELECT to_vertex_id,message_delivery_states,wait_duration,wait_time,wait_type FROM edges WHERE campaign_id = ? AND from_vertex_id = ? AND vertex_type = 'ENGAGEMENT' AND deleted_at is NULL"
	QUERY_GetPreviousEngagementTargetUsers              = "SELECT id,campaign_id,engagement_vertex_id,reference_id,event_reference_id,actor_contact_id,status,created_at,updated_at,deleted_at FROM user_journey_target_users WHERE campaign_id = ? AND engagement_vertex_id = ? AND reference_id = ?"
	QUERY_UpdateCampaign_CronExpression_Status_MetaData = "UPDATE campaigns SET cron_expression = ?,status = ?,attributes = ? where id = ?"
	QUERY_DeleteUserJourneyVerticesByCampaignId         = "UPDATE user_journey_vertices SET deleted_at = CURRENT_TIMESTAMP() WHERE campaign_id = ? AND deleted_at is NULL"
	QUERY_DeleteEngagementVerticesByCampaignId          = "UPDATE engagement_vertices SET deleted_at = CURRENT_TIMESTAMP() WHERE campaign_id = ? AND deleted_at is NULL"
	QUERY_DeleteEdgesByCampaignId                       = "UPDATE edges SET deleted_at = CURRENT_TIMESTAMP() WHERE campaign_id = ? AND vertex_type = ? AND deleted_at is NULL"
)

var (
	QUERY_GetDynamicDataByKey = "SELECT campaign_id, dynamic_key, cta_link, media FROM dynamic_media WHERE campaign_id = ? and dynamic_key = ?;"

	QUERY_AddDynamicData                       = "insert into dynamic_media(campaign_id, dynamic_key, cta_link, media) values(?,?,?,?)"
	QUERY_DISTINCT_USERS                       = "SELECT DISTINCT(user_type) FROM (@query@)"
	QUERY_ATHENA_SMS_ATTRIBUTES                = ""
	QUERY_ATHENA_APP_NOTIFICATION_ATTRIBUTES   = ""
	QUERY_ATHENA_WHATSAPP_ATTRIBUTES           = ""
	QUERY_ATHENA_EMAIL_ATTRIBUTES              = ""
	QUERY_USERLIST_SMS_ATTRIBUTES              = ""
	QUERY_USERLIST_APP_NOTIFICATION_ATTRIBUTES = ""
	QUERY_USERLIST_WHATSAPP_ATTRIBUTES         = ""
	QUERY_USERLIST_EMAIL_ATTRIBUTES            = ""
)

func GenerateFilterUserJourneyCampaignsQuery(request *fs.FilterUserJourneyCampaignRequest, startTime time.Time, endTime time.Time) (string, []interface{}) {

	query := "SELECT c.id as id,c.namespace as namespace,c.name as name,c.status as status,c.created_at as created_at," +
		"count(ujtu.id) as count,ujtu.status as target_users_status FROM campaigns c left join user_journey_target_users ujtu " +
		"on c.id = ujtu.campaign_id WHERE TRUE "
	var args []interface{}
	var searchFilterId int64
	var searchFilterName string
	searchFilterId, err := cast.ToInt64E(request.SearchFilter)
	if err != nil {
		searchFilterName = cast.ToString(request.SearchFilter)
	}

	if searchFilterId != 0 {
		query += " AND c.id = " + cast.ToString(searchFilterId)
	}
	if searchFilterName != "" {
		query += " AND c.name like '%" + searchFilterName + "%'"
	}
	if request.Status != common.CampaignStatus_NO_CAMPAGIN_STATUS {
		query += " AND c.status = ? "
		args = append(args, request.Status.String())
	}
	if request.Namespace != common.NameSpace_NO_NAMESPACE {
		query += " AND c.namespace = ? "
		args = append(args, request.Namespace.String())
	}
	if !startTime.IsZero() {
		query += " AND c.created_at >= ? "
		args = append(args, startTime)
	}
	if !startTime.IsZero() {
		query += " AND c.created_at <= ? "
		args = append(args, endTime)
	}
	query += " AND c.type = 'USER_JOURNEY' AND c.deleted_at is null GROUP BY c.id, ujtu.status ORDER BY c.updated_at DESC "
	query += " LIMIT " + cast.ToString(request.Limit) + " OFFSET " + cast.ToString(request.Limit*(request.PageNumber-1))
	return query, args
}

func GenerateFilterCampaignsQuery(request *fs.FilterCampaignRequest) (string, []interface{}) {

	query := "SELECT id,namespace,name,description,cron_expression," +
		"occurrences,communication_channel,status,type," +
		"schedule_type, query, inaction_query, inaction_duration, " +
		"attributes, created_by_actorId,created_by_actorType," +
		"updated_by_actorId,updated_by_actorType,version,created_at," +
		"updated_at,deleted_at " +
		"FROM campaigns " +
		"WHERE TYPE != 'USER_JOURNEY'"
	var args []interface{}

	if request.Name != "" {
		query += "AND name like '%" + request.Name + "%'"
	}
	if request.Status != common.CampaignStatus_NO_CAMPAGIN_STATUS {
		query += "AND status = ? "
		args = append(args, request.Status.String())
	}
	if request.Description != "" {
		query += "AND description like '%" + request.Description + "%'"
	}
	query += "AND deleted_at is null "
	query += "ORDER BY ID DESC"
	query += " LIMIT " + cast.ToString(request.Limit) + " OFFSET " + cast.ToString(request.Limit*(request.PageNumber-1))
	return query, args
}
