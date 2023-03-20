package models

import "database/sql"

type AddCampaignRequestVO struct {
	Namespace            sql.NullString
	Name                 sql.NullString
	Description          sql.NullString
	CronExpression       sql.NullString
	Occurrences          sql.NullInt32
	CommunicationChannel sql.NullString
	Status               sql.NullString
	Type                 sql.NullString
	ScheduleType         sql.NullString
	Query                sql.NullString
	InactionQuery        sql.NullString
	InactionDuration     sql.NullInt64
	Attributes           sql.NullString
	CreatedByActorid     sql.NullInt64
	CreatedByActortype   sql.NullString
}
type UpdateCampaignRequestVO struct {
	Name                 sql.NullString
	CronExpression       sql.NullString
	Status               sql.NullString
	Query                sql.NullString
	Namespace            sql.NullString
	Occurrences          sql.NullInt32
	CommunicationChannel sql.NullString
	Type                 sql.NullString
	Description          sql.NullString
	ScheduleType         sql.NullString
	InactionQuery        sql.NullString
	InactionDuration     sql.NullInt64
	Attributes           sql.NullString
	UpdatedByActorid     sql.NullInt64
	UpdatedByActortype   sql.NullString
	Id                   sql.NullInt64
}
type AddCampaignTemplateRequestVO struct {
	CampaignId          sql.NullInt64
	TemplateName        sql.NullString
	CampaignName        sql.NullString
	DistributionPercent sql.NullInt32
}
type DeleteCampaignTemplateRequestVO struct {
	CampaignId sql.NullInt64
}
type FindCampaignByIdResponseVO struct {
	Id                   sql.NullInt64
	Namespace            sql.NullString
	Name                 sql.NullString
	Description          sql.NullString
	CronExpression       sql.NullString
	Occurrences          sql.NullInt32
	CommunicationChannel sql.NullString
	Status               sql.NullString
	Type                 sql.NullString
	ScheduleType         sql.NullString
	Query                sql.NullString
	InactionQuery        sql.NullString
	InactionDuration     sql.NullInt64
	Attributes           sql.NullString
	CreatedByActorid     sql.NullInt64
	CreatedByActortype   sql.NullString
	UpdatedByActorid     sql.NullInt64
	UpdatedByActortype   sql.NullString
	Version              sql.NullInt64
	CreatedAt            sql.NullString
	UpdatedAt            sql.NullString
	DeletedAt            sql.NullString
}
type FindCampaignTemplateByIdResponseVO struct {
	Id                  sql.NullInt64
	CampaignId          sql.NullInt64
	TemplateName        sql.NullString
	CampaignName        sql.NullString
	DistributionPercent sql.NullInt32
}
type AddTargetUserRequestVO struct {
	CampaignId sql.NullInt64
	UserId     sql.NullInt64
	UserType   sql.NullString
	Attributes sql.NullString
}

type AddControlGroupRequestVO struct {
	CampaignId  sql.NullInt64
	Attributes  sql.NullString
	BloomFilter sql.RawBytes
}

type FindControlGroupByCampaignIdRequestV0 struct {
	Id          sql.NullInt64
	CampaignId  sql.NullInt64
	Attributes  sql.NullString
	BloomFilter sql.RawBytes
}

type FindTargetUserByIdResponseVO struct {
	Id         sql.NullInt64
	CampaignId sql.NullInt64
	UserId     sql.NullInt64
	UserType   sql.NullString
	Attributes sql.NullString
}

type AddInactionTargetUserRequestVO struct {
	CampaignId sql.NullInt64
	UserId     sql.NullInt64
	UserType   sql.NullString
}

type FindInactionTargetUserByCampaignIdResponseVO struct {
	Id         sql.NullInt64
	CampaignId sql.NullInt64
	UserId     sql.NullInt64
	UserType   sql.NullString
}

type GetDynamicDataByKeyResponseVO struct {
	CampaignId sql.NullInt64
	DynamicKey sql.NullString
	CtaLink    sql.NullString
	Media      sql.NullString
}

type AddDynamicDataRequestVO struct {
	CampaignId sql.NullInt64
	DynamicKey sql.NullString
	CtaLink    sql.NullString
	Media      sql.NullString
}
type FindQueryCampaignResponseVO struct {
	Name  sql.NullString
	Query sql.NullString
}
type FindQueryCampaignRequestVO struct {
	Type sql.NullString
}
type AddQueryCampaignRequestVO struct {
	Name      sql.NullString
	Type      sql.NullString
	Query     sql.NullString
	UpdatedBy sql.NullString
}
