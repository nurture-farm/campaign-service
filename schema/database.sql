create database campaign_service;

use campaign_service;

CREATE TABLE `campaigns` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `namespace` varchar(16) NOT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `cron_expression` varchar(32) NOT NULL,
  `occurrences` int DEFAULT NULL,
  `communication_channel` varchar(16) NOT NULL,
  `status` varchar(16) NOT NULL,
  `type` varchar(16) NOT NULL,
  `query` text DEFAULT NULL,
  `attributes` json DEFAULT NULL,
  `created_by_actorId` bigint NOT NULL,
  `created_by_actorType` varchar(16) NOT NULL,
  `updated_by_actorId` bigint DEFAULT NULL,
  `updated_by_actorType` varchar(16) DEFAULT NULL,
  `version` bigint DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `campaign_templates` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `campaign_id` bigint NOT NULL,
  `template_name` varchar(36) NOT NULL,
  `campaign_name` varchar(36) NOT NULL,
  `distribution_percent` int NOT NULL,
  `version` bigint DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `target_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `campaign_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `user_type` varchar(16) NOT NULL,
  `version` bigint DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `campaignservicedb`.`campaign_templates`
CHANGE COLUMN `template_name` `template_name` VARCHAR(64) NOT NULL ,
CHANGE COLUMN `campaign_name` `campaign_name` VARCHAR(64) NOT NULL ;

ALTER TABLE `campaign_service`.`campaigns`
ADD COLUMN `inaction_query` TEXT NULL DEFAULT NULL AFTER `query`,
ADD COLUMN `inaction_duration` BIGINT NULL DEFAULT NULL AFTER `inaction_query`,
ADD COLUMN `schedule_type` VARCHAR(64) NULL DEFAULT NULL AFTER `type`;

CREATE TABLE `inaction_overtime_users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `campaign_id` bigint NOT NULL,
    `user_id` bigint NOT NULL,
    `user_type` varchar(16) NOT NULL,
    `version` bigint DEFAULT 1,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `campaign_query` (
      `id` bigint NOT NULL AUTO_INCREMENT,
      `name` varchar(64) NOT NULL,
      `type` varchar(16) NOT NULL,
      `query` text NOT NULL,
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `updated_by` varchar(64) DEFAULT NULL,
      PRIMARY KEY (`id`),
      KEY `name_type` (`name`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table engagement_vertices;
drop table user_journey_vertices;
drop table edges;
drop table user_journey_target_users;

CREATE TABLE `user_journey_vertices` (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `campaign_id` bigint NOT NULL,
     `event_type` varchar(32) NOT NULL,
     `event_name` text NOT NULL,
     `inaction_duration` bigint DEFAULT NULL,
     `inaction_event_name` varchar(32) DEFAULT NULL,
     `version` bigint DEFAULT 1,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp NULL,
     PRIMARY KEY (`id`),
     FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `engagement_vertices` (
       `id` bigint NOT NULL AUTO_INCREMENT,
       `campaign_id` bigint NOT NULL,
       `template_name` varchar(128) NOT NULL,
       `attributes` json NOT NULL,
       `athena_query` text,
       `channel` varchar(16) NOT NULL,
       `version` bigint DEFAULT 1,
       `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
       `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       `deleted_at` timestamp NULL,
       PRIMARY KEY (`id`),
       FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `edges` (
         `id` bigint NOT NULL AUTO_INCREMENT,
         `campaign_id` bigint NOT NULL,
         `vertex_type` varchar(16) NOT NULL,
         `from_vertex_id` bigint NOT NULL,
         `to_vertex_id` bigint DEFAULT NULL,
         `wait_duration` bigint DEFAULT NULL,
         `wait_time` timestamp DEFAULT NULL,
         `wait_type` varchar(16) NOT NULL,
         `message_delivery_states` set('VENDOR_UNDELIVERED','VENDOR_DELIVERED','CUSTOMER_UNDELIVERED','CUSTOMER_DELIVERED','CUSTOMER_SENT','CUSTOMER_READ'),
         `version` bigint DEFAULT 1,
         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
         `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
         `deleted_at` timestamp NULL,
         PRIMARY KEY (`id`),
         FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_journey_perfix_expressions` (
       `id` bigint NOT NULL AUTO_INCREMENT,
       `campaign_id` bigint NOT NULL,
       `prefix_expression` varchar(256) default null,
       `version` bigint DEFAULT 1,
       `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
       `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       `deleted_at` timestamp NULL,
       PRIMARY KEY (`id`),
       FOREIGN KEY (campaign_id) REFERENCES campaigns(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_journey_target_users` (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `campaign_id` bigint NOT NULL,
     `engagement_vertex_id` bigint NOT NULL,
     `reference_id` varchar(36) NOT NULL,
     `event_reference_id` varchar(36) NOT NULL,
     `actor_contact_id` varchar(512) DEFAULT NULL,
     `status` varchar(16) DEFAULT NULL,
     `version` bigint DEFAULT 1,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp NULL,
     PRIMARY KEY (`id`),
     FOREIGN KEY (campaign_id) REFERENCES campaigns(id),
     FOREIGN KEY (engagement_vertex_id) REFERENCES engagement_vertices(id),
     CONSTRAINT `target_user_index` UNIQUE (`event_reference_id`,`reference_id`,`engagement_vertex_id`,`campaign_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `campaignservicedb`.`campaigns`
    CHANGE COLUMN `cron_expression` `cron_expression` VARCHAR(32) DEFAULT NULL ;

ALTER TABLE `target_users`
    ADD COLUMN `attributes` TEXT NULL DEFAULT NULL AFTER `user_type`;
