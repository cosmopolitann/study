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
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS "chat_record";

CREATE TABLE "chat_record" (
	"id" VARCHAR ( 64 ) NOT NULL,--id
	"name" VARCHAR ( 64 ) NOT NULL,--聊天对方的名称
	"from_id" VARCHAR ( 64 ) NOT NULL,--fromid
	"ptime" INTEGER ( 10 ) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--创建时间
	"last_msg" TEXT NOT NULL,--最后的消息
	"to_id" VARCHAR ( 64 ) NOT NULL,--toid
	PRIMARY KEY ( "id" ),--主键 id
	UNIQUE ( "id" ASC ) --唯一键 id
);
PRAGMA foreign_keys = true;
