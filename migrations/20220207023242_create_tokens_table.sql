-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS `tokens`
(
    `id`         bigint(20)       NOT NULL AUTO_INCREMENT,
    `user_id`    int(10) unsigned NOT NULL,
    `token`      varchar(255)          DEFAULT NULL,
    `revoked`    tinyint(1)            DEFAULT NULL,
    `created_at` timestamp        NULL DEFAULT NULL,
    `updated_at` timestamp        NULL DEFAULT NULL,
    `deleted_at` timestamp        NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `tokens_user_id_foreign` (`user_id`),
    CONSTRAINT `tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE `tokens`;
-- +goose StatementEnd
