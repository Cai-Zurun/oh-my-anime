CREATE DATABASE `anime_gf`;

USE `anime_gf`;

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '账号',
  `password` varchar(45) NOT NULL COMMENT '密码',
  `nickname` varchar(45) NOT NULL COMMENT '昵称',
  `create_time` timestamp NOT NULL COMMENT '创建时间/注册时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `anime` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '动漫ID',
  `name` varchar(45) NOT NULL COMMENT '名字',
  `link` varchar(2083) NOT NULL COMMENT '链接',
  `type` varchar(45) NOT NULL COMMENT '类型',
  `create_time` timestamp NOT NULL COMMENT '创建时间/注册时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `anime_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '动漫类型ID',
  `type` varchar(45) NOT NULL COMMENT '动漫类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;