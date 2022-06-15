CREATE TABLE IF NOT EXISTS `api_user_tokens` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `token` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT user_id
     FOREIGN KEY (user_id)
     REFERENCES users (id)
     ON DELETE CASCADE
     /**ON UPDATE CASCADE ON DELETE CASCADE*/
    ) ENGINE = InnoDB;