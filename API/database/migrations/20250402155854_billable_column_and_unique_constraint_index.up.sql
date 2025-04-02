-- Migration Up
ALTER TABLE `categories`
ADD COLUMN `billable` TINYINT(1) NOT NULL DEFAULT 0 AFTER `description`;

ALTER TABLE `projects`
ADD COLUMN `billable` TINYINT(1) NOT NULL DEFAULT 0 AFTER `status`,
ADD CONSTRAINT `unique_project_name` UNIQUE (`name`);