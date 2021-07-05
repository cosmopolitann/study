/*
 Navicat Premium Data Transfer

 Source Server         : xiaolong
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/07/2021 11:16:34
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_like
-- ----------------------------
DROP TABLE IF EXISTS "article_like";

CREATE TABLE article_like (
	"id" VARCHAR ( 64 ) NOT NULL,--id
	"user_id" VARCHAR ( 64 ) NOT NULL,--用户名字
	"article_id" VARCHAR ( 64 ) NOT NULL,--文章id
	"is_like" INTEGER ( 10 ) DEFAULT ( 0 ),--是否点赞
  PRIMARY KEY ( "id" ), --主键索引id
  FOREIGN KEY ( "article_id" ) REFERENCES "article" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,--外键关联article_id article的 id 级联删除
  UNIQUE ( "id" ASC ) --唯一键 id
);
PRAGMA foreign_keys = true;
