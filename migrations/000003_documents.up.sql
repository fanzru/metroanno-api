CREATE TABLE `documents` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `subject_id` INT(11) NOT NULL,
  `learning_outcame` LONGTEXT NOT NULL,
  `text_document` LONGTEXT NOT NULL,
  `min_number_of_annotators` INT(11) NOT NULL, 
  `current_number_of_annotators_assigned` INT(11) NOT NULL,
  `min_number_of_questions_per_annotator` INT(11) NOT NULL, 
  `current_total_number_of_questions_annotated` INT(11) NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_documents_subjects` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);