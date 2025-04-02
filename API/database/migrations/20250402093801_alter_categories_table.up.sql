-- Migration Up
-- Modification de la table categories pour ajouter les colonnes user_id et project_id
ALTER TABLE `categories` 
ADD COLUMN `user_id` bigint(20) unsigned NOT NULL,
ADD COLUMN `project_id` bigint(20) unsigned NOT NULL,
ADD CONSTRAINT `fk_categories_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
ADD CONSTRAINT `fk_categories_projects` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

-- Création d'un index unique sur name et project_id dans la table categories
CREATE UNIQUE INDEX idx_name_project ON `categories` (`name`, `project_id`);