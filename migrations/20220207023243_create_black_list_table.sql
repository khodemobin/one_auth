-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE `black_list`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `token`      varchar(20)         NOT NULL,
    `created_at` timestamp           NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE `black_list`;

-- +goose StatementEnd