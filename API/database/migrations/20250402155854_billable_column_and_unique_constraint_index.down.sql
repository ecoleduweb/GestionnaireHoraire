-- Migration Down
ALTER TABLE `categories`
DROP COLUMN `billable`;

ALTER TABLE `projects`
DROP COLUMN `billable`,
DROP INDEX `unique_project_name`;