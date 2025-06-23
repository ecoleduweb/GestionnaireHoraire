-- Migration Up
ALTER TABLE `projects`
RENAME COLUMN name TO unique_id;

ALTER TABLE `projects`
RENAME COLUMN description TO name;

ALTER TABLE `projects`
ADD COLUMN estimated_hours INT NULL;
