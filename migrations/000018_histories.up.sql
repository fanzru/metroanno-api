CREATE TABLE histories (
    id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id INT(11) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    deleted_at DATETIME,
    CONSTRAINT `fk_questions_histories_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION
);
