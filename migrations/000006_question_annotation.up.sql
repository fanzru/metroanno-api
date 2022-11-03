CREATE TABLE `question_annoatations` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) NOT NULL,
  `document_id` INT(11) NOT NULL,
  `is_learning_outcome_shown` BOOLEAN NOT NULL,
  `question_order` INT(11) NOT NULL,
  `question_type_id` INT(11) NOT NULL,
  `keywords` LONGTEXT NOT NULL,
  `question_text` LONGTEXT NOT NULL,
  `answer_text` LONGTEXT NOT NULL,
  `time_duration` BIGINT(20) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_qa_qt` FOREIGN KEY (`question_type_id`) REFERENCES `question_types` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_qa_documents` FOREIGN KEY (`document_id`) REFERENCES `documents` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_qa_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);