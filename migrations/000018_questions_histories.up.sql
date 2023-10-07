CREATE TABLE questions_histories (
    id SERIAL PRIMARY KEY,
    `difficulty` VARCHAR(255) NOT NULL,
    `reading_material` TEXT NOT NULL,
    `topic` VARCHAR(255) NOT NULL,
    `random` VARCHAR(255),
    `bloom` VARCHAR(255),
    `graesser` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    `user_id` INT(11) NOT NULL,
    CONSTRAINT `fk_questions_histories_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);
