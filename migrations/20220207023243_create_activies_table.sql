-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

CREATE TABLE `activies` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `action` varchar(20) NOT NULL,
  `ip` varchar(12) DEFAULT NULL,
  `path` varchar(100) DEFAULT NULL,
  `operation` varchar(100) DEFAULT NULL,
  `version` varchar(10) DEFAULT NULL,
  `headers` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

DROP TABLE `activies`;

-- +goose StatementEnd