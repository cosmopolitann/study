/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:42:02
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "sys_user";

CREATE TABLE "sys_user" (
  "id" VARCHAR (64) NOT NULL,--id
  "peer_id" VARCHAR (64) NOT NULL,--用户id==peerid
  "name" VARCHAR (128),--文件名字
  "phone" VARCHAR (64),--手机号
  "sex" INT (10) DEFAULT (1),--性别 0 未知 1男 2 女
  "ptime" integer(64) NOT NULL DEFAULT 0,--创建时间
  "utime" integer(64) NOT NULL DEFAULT 0,--更新时间
  "nickname" VARCHAR (128),--昵称
  "img" VARCHAR (128),--图片
  PRIMARY KEY ("id"),--主键索引 id
  UNIQUE ("id" ASC)--唯一键id
);
PRAGMA foreign_keys = true;
