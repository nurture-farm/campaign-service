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

import "database/sql"

type BaseVO struct {
	Version   sql.NullInt64
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	DeletedAt sql.NullString
}

type AddUserJourneyVertexVO struct {
	CampaignId        sql.NullInt64
	EventType         sql.NullString
	EventName         sql.NullString
	InactionDuration  sql.NullInt64
	InactionEventName sql.NullString
	Attributes        sql.NullString
}

type AddEngagementVertexVO struct {
	CampaignId   sql.NullInt64
	TemplateName sql.NullString
	Attributes   sql.NullString
	AthenaQuery  sql.NullString
	Channel      sql.NullString
}

type AddEdgeVO struct {
	CampaignId            sql.NullInt64
	VertexType            sql.NullString
	FromVertexId          sql.NullInt64
	ToVertexId            sql.NullInt64
	WaitDuration          sql.NullInt64
	WaitTime              sql.NullTime
	WaitType              sql.NullString
	MessageDeliveryStates sql.NullString
}

type AddUserJourneyPrefixVO struct {
	CampaignId       sql.NullInt64
	PrefixExpression sql.NullString
}

type FindUserJourneyVertexVO struct {
	Id                sql.NullInt64
	CampaignId        sql.NullInt64
	EventType         sql.NullString
	EventName         sql.NullString
	InactionDuration  sql.NullInt64
	InactionEventName sql.NullString
	BaseVO            BaseVO
	Attributes        sql.NullString
}

type FindEdgesVO struct {
	Id                    sql.NullInt64
	CampaignId            sql.NullInt64
	VertexType            sql.NullString
	FromVertexId          sql.NullInt64
	ToVertexId            sql.NullInt64
	WaitDuration          sql.NullInt64
	WaitTime              sql.NullTime
	WaitType              sql.NullString
	MessageDeliveryStates sql.NullString
	BaseVO                BaseVO
}

type FindEngagementVertexVO struct {
	Id           sql.NullInt64
	CampaignId   sql.NullInt64
	TemplateName sql.NullString
	Attributes   sql.NullString
	AthenaQuery  sql.NullString
	Channel      sql.NullString
	BaseVO       BaseVO
}

type FindUserJourneyTargetUsersVO struct {
	Id                 sql.NullInt64
	CampaignId         sql.NullInt64
	EngagementVertexId sql.NullInt64
	ReferenceId        sql.NullString
	EventReferenceId   sql.NullString
	ActorContactId     sql.NullString
	Status             sql.NullString
	BaseVO             BaseVO
}

type AddUserJourneyTargetUsersVO struct {
	CampaignId         sql.NullInt64
	EngagementVertexId sql.NullInt64
	ReferenceId        sql.NullString
	EventReferenceId   sql.NullString
	ActorContactId     sql.NullString
	Status             sql.NullString
}

type FindEngagementStartVertexVO struct {
	ID sql.NullInt64
}

type FindUserJourneyCampaigVO struct {
	Id                sql.NullInt64
	Namespace         sql.NullString
	Name              sql.NullString
	Status            sql.NullString
	CreatedAt         sql.NullTime
	Count             sql.NullInt64
	TargetUsersStatus sql.NullString
}

type FindPreviousEngagementVertexVO struct {
	Id                    sql.NullInt64
	MessageDeliveryStates sql.NullString
}

type FindNextEngagementVertexVO struct {
	Id                    sql.NullInt64
	MessageDeliveryStates sql.NullString
	WaitDuration          sql.NullInt64
	WaitTime              sql.NullTime
	WaitType              sql.NullString
}

type UpdateCampaignVO struct {
	CronExpression sql.NullString
	Status         sql.NullString
	Attributes     sql.NullString
	CampaignID     sql.NullInt64
}

type DeleteEngagementVerticesVO struct {
	CampaignId sql.NullInt64
}

type DeleteUserJourneyVerticesVO struct {
	CampaignId sql.NullInt64
}

type DeleteEdgesVO struct {
	CampaignId sql.NullInt64
	VertexType sql.NullString
}
