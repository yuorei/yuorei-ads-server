-- Create "users" table
CREATE TABLE `users` (
 `user_id` varchar(255) NOT NULL,
 `username` varchar(255) NOT NULL,
 `email` varchar(255) NOT NULL,
 `hashed_password` char(60) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`user_id`),
 UNIQUE INDEX `email` (`email`),
 UNIQUE INDEX `username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "campaigns" table
CREATE TABLE `campaigns` (
 `campaign_id` varchar(255) NOT NULL,
 `user_id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 `budget` int NOT NULL,
 `start_date` timestamp NOT NULL,
 `end_date` timestamp NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 `is_approval` bool NULL,
 PRIMARY KEY (`campaign_id`),
 INDEX `user_id` (`user_id`),
 CONSTRAINT `campaigns_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "ad_groups" table
CREATE TABLE `ad_groups` (
 `ad_group_id` varchar(255) NOT NULL,
 `campaign_id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 `is_approval` bool NULL,
 PRIMARY KEY (`ad_group_id`),
 INDEX `campaign_id` (`campaign_id`),
 CONSTRAINT `ad_groups_ibfk_1` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`campaign_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "ads" table
CREATE TABLE `ads` (
 `ad_id` varchar(255) NOT NULL,
 `ad_group_id` varchar(255) NOT NULL,
 `type` varchar(255) NOT NULL,
 `content` text NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 `is_approval` bool NULL,
 PRIMARY KEY (`ad_id`),
 INDEX `ad_group_id` (`ad_group_id`),
 CONSTRAINT `ads_ibfk_1` FOREIGN KEY (`ad_group_id`) REFERENCES `ad_groups` (`ad_group_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "impressions" table
CREATE TABLE `impressions` (
 `impression_id` varchar(255) NOT NULL,
 `ad_id` varchar(255) NOT NULL,
 `date` date NOT NULL,
 `impressions` int NOT NULL,
 `clicks` int NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`impression_id`),
 INDEX `ad_id` (`ad_id`),
 CONSTRAINT `impressions_ibfk_1` FOREIGN KEY (`ad_id`) REFERENCES `ads` (`ad_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "organizations" table
CREATE TABLE `organizations` (
 `organization_id` varchar(255) NOT NULL,
 `organization_name` varchar(255) NOT NULL,
 `representative_name` varchar(255) NOT NULL,
 `representative_email` varchar(255) NOT NULL,
 `purpose` varchar(255) NOT NULL,
 `category` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`organization_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "organizations_users" table
CREATE TABLE `organizations_users` (
 `organization_id` varchar(255) NOT NULL,
 `user_id` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`organization_id`, `user_id`),
 INDEX `fk_user_id` (`user_id`),
 CONSTRAINT `fk_organization_id` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`organization_id`) ON UPDATE NO ACTION ON DELETE CASCADE,
 CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "targeting" table
CREATE TABLE `targeting` (
 `targeting_id` varchar(255) NOT NULL,
 `ad_id` varchar(255) NOT NULL,
 `type` varchar(255) NOT NULL,
 `value` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`targeting_id`),
 INDEX `ad_id` (`ad_id`),
 CONSTRAINT `targeting_ibfk_1` FOREIGN KEY (`ad_id`) REFERENCES `ads` (`ad_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "roles" table
CREATE TABLE `roles` (
 `role_id` varchar(255) NOT NULL,
 `name` varchar(255) NOT NULL,
 `description` text NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`role_id`),
 UNIQUE INDEX `name` (`name`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "user_roles" table
CREATE TABLE `user_roles` (
 `user_id` varchar(255) NOT NULL,
 `role_id` varchar(255) NOT NULL,
 `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
 `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 `deleted_at` timestamp NULL,
 PRIMARY KEY (`user_id`, `role_id`),
 INDEX `role_id` (`role_id`),
 CONSTRAINT `user_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
 CONSTRAINT `user_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
