CREATE TABLE `question_types` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `question_type` VARCHAR(255) NOT NULL UNIQUE,
  `description` LONGTEXT NOT NULL,
  PRIMARY KEY (`id`)
);