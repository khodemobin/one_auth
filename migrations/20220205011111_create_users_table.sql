-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE `users`
(
    `id`                   int(10) unsigned    NOT NULL AUTO_INCREMENT,
    `phone`                varchar(20)         NOT NULL,
    `password`             varchar(255)                 DEFAULT NULL,
    `confirmation_token`   varchar(255)                 DEFAULT NULL,
    `confirmation_sent_at` timestamp           NULL     DEFAULT NULL,
    `confirmed_at`         timestamp           NULL     DEFAULT NULL,
    `recovery_token`       varchar(255)                 DEFAULT NULL,
    `recovery_sent_at`     timestamp           NULL     DEFAULT NULL,
    `phone_change_token`   varchar(255)                 DEFAULT NULL,
    `phone_change`         varchar(255)                 DEFAULT NULL,
    `phone_change_sent_at` timestamp           NULL     DEFAULT NULL,
    `role`                 varchar(255)                 DEFAULT NULL,
    `status`               tinyint(3) unsigned NOT NULL DEFAULT 2,
    `is_super_admin`       tinyint(1)                   DEFAULT 0,
    `last_sign_in_at`      timestamp           NULL     DEFAULT NULL,
    `created_at`           timestamp           NULL     DEFAULT NULL,
    `updated_at`           timestamp           NULL     DEFAULT NULL,
    `deleted_at`           timestamp           NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_phone_unique` (`phone`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb4;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE `users`;
-- +goose StatementEnd