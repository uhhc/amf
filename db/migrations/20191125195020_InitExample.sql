-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `info_examples` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `example_id` varchar(36) NOT NULL DEFAULT '' COMMENT '样例ID，格式为 uuid',
  `example_name` varchar(32) NOT NULL DEFAULT '' COMMENT '样例名称',
  `status` varchar(16) NOT NULL DEFAULT '' COMMENT '样例状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS `info_examples`;
