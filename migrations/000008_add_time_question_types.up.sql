ALTER TABLE question_types
ADD COLUMN `created_at` DATETIME NOT NULL,
ADD COLUMN  `deleted_at` DATETIME;