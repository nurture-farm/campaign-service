
INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('AddCampaign',NULL,NULL,1,1,1,0,
     'insert into campaigns(namespace,name,description,cron_expression,occurrences,communication_channel,status,type,query,created_by_actorId,created_by_actorType) values(?,?,?,?,?,?,?,?,?,?,?)',
     'namespace:ContractEnum.farm.nurture.core.contracts.common.NameSpace,name:campaigns.name,description:campaigns.description,cron_expression:campaigns.cron_expression,occurrences:campaigns.occurrences,communication_channel:ContractEnum.farm.nurture.core.contracts.common.CommunicationChannel,status:ContractEnum.farm.nurture.core.contracts.common.CampaignStatus,type:ContractEnum.farm.nurture.core.contracts.common.CampaignQueryType,query:campaigns.query,created_by_actorId:campaigns.created_by_actorId,created_by_actorType:campaigns.created_by_actorType',
     0,1,1,0,'I',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('UpdateCampaign',NULL,NULL,1,1,1,0,
     'Update campaigns set name = ?, cron_expression = ?, status = ?, query = ?, updated_by_actorId = ?, updated_by_actorType = ? where id = ?',
     'name:campaigns.name,cron_expression:campaigns.cron_expression,status:ContractEnum.farm.nurture.core.contracts.common.CampaignStatus,query:campaigns.query,updated_by_actorId:campaigns.updated_by_actorId, updated_by_actorType:campaigns.updated_by_actorType,id:campaigns.id',
     NULL,0,NULL,1,1,0,NULL,NULL,NULL,NULL,'U',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('AddCampaignTemplate',NULL,NULL,1,1,1,0,
     'Insert into campaign_templates(campaign_id,template_name,campaign_name,distribution_precent) values (?,?,?,?)',
     'campaign_id:campaign_templates.id,template_name:campaign_templates.template_name,campaign_name:campaign_templates.campaign_name,distribution_precent:campaign_templates.distribution_percent',
     0,1,1,0,'I',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`,
                                             `req_override`,
                                             `mutation`, `status`) VALUES
    ('AddNewCampaign',NULL,NULL,1,1,1,0,
     NULL,
     NULL,
     0,0,0,0,
     'AddCampaignRequest addCampaignRequest = 3; repeated AddCampaignTemplateRequest bulkAddCampaignTemplateRequest = 4; repeated AddTargetUserRequest addTargetUserRequest = 5;',
     'I',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`,
                                             `req_override`, `res_override`,
                                             `mutation`, `status`)
VALUES ('GetUserList',NULL,NULL,1,1,1,0,
        NULL,
        NULL,
        0,0,0,0,
        'farm.nurture.core.contracts.common.CampaignQueryType campaignQueryType = 3;\nstring query = 4;\nfarm.nurture.core.contracts.common.CommunicationChannel communicationChannel = 5;\nrepeated farm.nurture.core.contracts.common.ActorID actorId = 6;\n',
        'repeated farm.nurture.core.contracts.ce.ActorDetails actorDetails = 3;\n',
        'S',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('FindCampaignById', NULL, NULL, 1,1,1,0,
     'select id,namespace,name,description,cron_expression,communication_channel,status,type,query,created_by_actorId,created_by_actorType,updated_by_actorId,updated_by_actorType,version,created_at,updated_at,deleted_at from campaigns where id = ?',
     'id:campaigns.id',
     1,1,1,0,'S',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('FindCampaignTemplateById',NULL,NULL,1,1,1,0,
     'select id,campaign_id,template_name,campaign_name,distribution_percent from campaign_templates where campaign_id = ?',
     'campaign_id:campaigns.id',
     0, 1,1,0,'S',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`) VALUES
    ('AddTargetUser',NULL,NULL,1,1,1,0,
     'insert into target_users(campaign_id,user_id,user_type) values (?,?,?)',
     'campaign_id:campaigns.id,user_id:target_users.user_id,user_type:target_users.user_type',
     0, 1,1,0, 'I',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`)
VALUES ('FindTargetUserById', NULL, NULL, '1', '1', '1', '0',
        'select id,campaign_id,user_id,user_type from target_users where campaign_id = ?',
        'campaign_id:campaigns.id',
        '0', '1', '1', '0', 'S', '0');

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`)
VALUES ('AddInactionTargetUser', NULL, NULL, 1,1,1,0,
        'insert into inaction_target_users(campaign_id, user_id, user_type) values(?,?,?)',
        'campaign_id:campaigns.id,user_id:inaction_target_users.user_id,user_type:inaction_target_users.user_type',
        0, 1, 1, 0, 'I',1);

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`)
VALUES ('FindInactionTargetUserByCampaignId', NULL, NULL, '1', '1', '1', '0',
        'select id,campaign_id,user_id,user_type from inaction_target_users where campaign_id = ?',
        'campaign_id:campaigns.id',
        '0', '1', '1', '0', 'S', '0');

INSERT INTO `campaign_service`.`app_config` (`title`, `req_name`, `res_name`,`decl_req`, `decl_res`, `decl_grpc`, `decl_grapql`,
                                             `sql_stmt`,
                                             `sql_params`,
                                             `sql_uniquekey`, `impl_dao`, `impl_grpc`, `impl_reacrjs`, `mutation`, `status`)
VALUES ('GetDynamicDataByKey', NULL, NULL, '1', '1', '1', '0',
        'SELECT campaign_id, cta_link, media FROM dynamic_media WHERE campaign_id = ? and dynamic_key = ?;',
        'campaign_id: dynamic_media.campaign_id, key: dynamic_media.dynamic_key',
        '0', '1', '1', '0', 'S', 1);

