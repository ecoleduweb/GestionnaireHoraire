-- Migration Down
-- Pour annuler les modifications (si nécessaire)
ALTER TABLE `categories` 
DROP FOREIGN KEY `fk_categories_projects`,
DROP FOREIGN KEY `fk_categories_users`,
DROP COLUMN `project_id`,
DROP COLUMN `user_id`;

DROP INDEX idx_name_project ON `categories`;