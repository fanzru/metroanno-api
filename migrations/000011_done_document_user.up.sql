CREATE TABLE `done_document_user` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) NOT NULL,
  `document_id` INT(11) NOT NULL,
  `done` BOOLEAN NOT NULL,
  PRIMARY KEY (`id`)
);