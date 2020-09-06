-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `info_examples` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `example_id` varchar(36) NOT NULL DEFAULT '',
  `example_name` varchar(32) NOT NULL DEFAULT '',
  `status` varchar(16) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS `info_examples`;
