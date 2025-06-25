-- Migration Down
ALTER TABLE `projects`
DROP COLUMN estimated_hours;

ALTER TABLE `projects`
RENAME COLUMN name TO description;

ALTER TABLE `projects`
RENAME COLUMN unique_id TO name;
