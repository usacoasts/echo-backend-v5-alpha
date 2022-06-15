CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL COMMENT 'ユーザ名',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `password` VARCHAR(255) NOT NULL COMMENT 'パスワード',
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
PRIMARY KEY (`id`)
) ENGINE = InnoDB;