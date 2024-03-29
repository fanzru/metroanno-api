CREATE TABLE `users` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `type` INT(11) NOT NULL,
  `is_document_annotator` BOOLEAN NOT NULL,
  `is_question_annotator` BOOLEAN NOT NULL,
  `subject_preference` LONGTEXT, 
  `username` VARCHAR(255) NOT NULL UNIQUE,
  `password` VARCHAR(255) NOT NULL,
  `contact` VARCHAR(255) NOT NULL,
  `age` INT(11) NOT NULL,
  `number_of_document_added` INT(11) NOT NULL,
  `number_of_question_annotated` INT(11) NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  PRIMARY KEY (`id`)
);