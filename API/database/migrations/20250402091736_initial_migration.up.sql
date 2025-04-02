-- Migration Up
CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(50) NOT NULL,
    `last_name` varchar(50) NOT NULL,
    `email` varchar(255) NOT NULL UNIQUE,
    `role` int NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `categories` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `description` text NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `projects` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `description` text NOT NULL,
    `status` TINYINT(1) NOT NULL DEFAULT 1,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `end_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `activities` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `start_date` datetime(3) DEFAULT NULL,
  `end_date` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  `project_id` bigint(20) unsigned NOT NULL,
  `category_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_activities_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_activities_projects` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`),
  CONSTRAINT `fk_activities_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);
-- Modification de la table categories pour ajouter les colonnes user_id et project_id
ALTER TABLE `categories` 
ADD COLUMN `user_id` bigint(20) unsigned NOT NULL,
ADD COLUMN `project_id` bigint(20) unsigned NOT NULL,
ADD CONSTRAINT `fk_categories_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
ADD CONSTRAINT `fk_categories_projects` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

-- Création d'un index unique sur name et project_id dans la table categories
CREATE UNIQUE INDEX idx_name_project ON `categories` (`name`, `project_id`);