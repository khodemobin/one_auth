-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS `refresh_tokens`
(
    `id`         bigint(20)   NOT NULL AUTO_INCREMENT,
    `user_id`    varchar(255) NOT NULL,
    `token`      varchar(255),
    `revoked`    tinyint(4) DEFAULT 0,
    `created_at` timestamp    NULL,
    `updated_at` timestamp    NULL,
    PRIMARY KEY (`id`),
    KEY `tokens_user_id_foreign` (`user_id`),
    KEY `tokens_token_index` (`token`) USING BTREE,
    CONSTRAINT `tokens_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE `refresh_tokens`;

-- +goose StatementEnd