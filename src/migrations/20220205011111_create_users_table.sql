-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE `users`
(
    `id`              varchar(255)        NOT NULL,
    `phone`           varchar(20)         NULL,
    `email`           varchar(100)        NULL,
    `username`        varchar(100)        NULL,
    `password`        varchar(255)        NULL,
    `otp_key`         varchar(255)        NULL,
    `otp_value`       varchar(255)        NULL,
    `confirmed_at`    timestamp           NULL,
    `is_active`       tinyint(1) unsigned NOT NULL DEFAULT 0,
    `last_sign_in_at` timestamp           NULL,
    `created_at`      timestamp           NULL,
    `updated_at`      timestamp           NULL,
    `deleted_at`      timestamp           NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_phone_unique` (`phone`),
    UNIQUE KEY `users_email_unique` (`email`),
    UNIQUE KEY `users_username_unique` (`username`),
    KEY `users_is_active_index` (`is_active`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE `users`;

-- +goose StatementEnd