# ************************************************************
# Sequel Ace SQL dump
# 版本号： 3024
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 5.5.5-10.5.9-MariaDB)
# 数据库: trading
# 生成时间: 2021-03-21 16:39:59 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 t_admin_user
# ------------------------------------------------------------

CREATE TABLE `t_admin_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# 转储表 t_order
# ------------------------------------------------------------

CREATE TABLE `t_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0,
  `symbol` varchar(50) NOT NULL DEFAULT '',
  `price` decimal(5,2) NOT NULL DEFAULT 0.00,
  `num` decimal(5,2) NOT NULL DEFAULT 0.00,
  `state` int(11) NOT NULL DEFAULT 0,
  `ctime` timestamp NOT NULL DEFAULT current_timestamp(),
  `mtime` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# 转储表 t_stock
# ------------------------------------------------------------

CREATE TABLE `t_stock` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `symbol` varchar(11) DEFAULT NULL,
  `ctime` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# 转储表 t_user
# ------------------------------------------------------------

CREATE TABLE `t_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `state` int(11) DEFAULT NULL COMMENT '状态',
  `last_login_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '上次登录时间',
  `ctime` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '编辑时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# 转储表 t_user_stock_star
# ------------------------------------------------------------

CREATE TABLE `t_user_stock_star` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `stock` varchar(50) NOT NULL DEFAULT '' COMMENT '股票',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
