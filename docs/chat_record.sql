/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 14:34:35
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS "chat_record";

CREATE TABLE "chat_record" (
  "id" VARCHAR (64) NOT NULL,--id
  "name" varchar (64) NOT NULL,--聊天对方的名称
  "from_id" VARCHAR (64) NOT NULL,--fromid
  "ptime" integer(64) NOT NULL DEFAULT 0,--创建时间
  "last_msg" TEXT NOT NULL,--最后的消息
  "to_id" varchar(64) NOT NULL,--toid
  PRIMARY KEY ("id"),--主键 id
  UNIQUE ("id" ASC)--唯一键 id
);
PRAGMA foreign_keys = true;
