-- Migration Up
ALTER TABLE `projects`
ADD COLUMN `manager_id` bigint(20) unsigned NOT NULL AFTER `id`,
ADD CONSTRAINT `fk_managers_projects` FOREIGN KEY (`manager_id`) REFERENCES `users` (`id`);