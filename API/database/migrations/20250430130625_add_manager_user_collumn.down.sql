-- Migration Down
ALTER TABLE `projects`
DROP FOREIGN KEY `fk_managers_projects`,
DROP COLUMN `manager_id`;