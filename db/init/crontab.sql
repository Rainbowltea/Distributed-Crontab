
CREATE DATABASE IF NOT EXISTS pink default charset utf8mb4 COLLATE utf8mb4_general_ci;

use distributed;

-- ----------------------------
-- Table structure for execute_snapshot_his
-- ----------------------------
DROP TABLE IF EXISTS `execute_snapshot_his`;
CREATE TABLE `execute_snapshot_his`
(
    `id`            bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `job_id`        varchar(32) NOT NULL,
    `job_name`      varchar(32) NOT NULL,
    `cron`          varchar(255) DEFAULT NULL,
    `ip`            varchar(32)  DEFAULT NULL,
    `state`         tinyint(4) DEFAULT NULL,
    `before_time`   varchar(32) NOT NULL,
    `start_time`    varchar(32) NOT NULL,
    `schedule_time` varchar(32) NOT NULL,
    `end_time`      varchar(32)  DEFAULT NULL,
    `times`         bigint(20) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;