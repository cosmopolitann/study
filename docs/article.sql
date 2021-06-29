/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:26:35
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "cloud_article";

CREATE TABLE "article" (
  "id" VARCHAR (64) NOT NULL,--id
  "user_id" VARCHAR (64) NOT NULL,--用户id
  "accesstory" VARCHAR (128),--附件
  "accesstory_type" INT (10) NOT NULL,--附件类型
  "text" text NOT NULL,--文本
  "tag" varchar (64) NOT NULL,--标签
  "ptime" integer(64) NOT NULL DEFAULT 0,--创建时间
  "play_num" INTEGER NOT NULL DEFAULT (0),--播放次数
  "share_num" INTEGER NOT NULL DEFAULT (0),--分享次数
  "title" varchar(128) NOT NULL,--标题
  "thumbnail" varchar(128) NOT NULL,--缩略图
  "file_name" varchar(128) NOT NULL,--文件名字
  "file_size" varchar(128) NOT NULL,--文件大小
  PRIMARY KEY ("id"),--主键 id
  UNIQUE ("id" ASC)--唯一键 id
);
PRAGMA foreign_keys = true;
