CREATE TABLE questions_histories (
    id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `difficulty` VARCHAR(255) NOT NULL,
    `reading_material` TEXT NOT NULL,
    `topic` VARCHAR(255) NOT NULL,
    `random` VARCHAR(255),
    `bloom` VARCHAR(255),
    `graesser` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    `history_id` INT(11) NOT NULL,
    CONSTRAINT `fk_questions_histories_histories` FOREIGN KEY (`history_id`) REFERENCES `histories` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);
