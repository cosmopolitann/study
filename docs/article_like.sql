/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:17:22
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_like
-- ----------------------------
DROP TABLE IF EXISTS "article_like";

CREATE TABLE "main"."article_like" (
  "id" VARCHAR NOT NULL,--id
  "user_id" varchar (64) NOT NULL,--用户名字
  "article_id" VARCHAR (64) NOT NULL,--文章id
  "is_like" INT (10) DEFAULT (0),--是否点赞
  PRIMARY KEY ("id"),--主键索引id
  FOREIGN KEY ("user_id") REFERENCES "sys_user" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,--外键关联user_id sys_user表的id
  FOREIGN KEY ("article_id") REFERENCES "article" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,--外键关联article_id article的 id 级联删除
  UNIQUE ("id" ASC)--唯一键 id
);
PRAGMA foreign_keys = true;
