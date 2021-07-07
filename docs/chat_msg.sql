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
-- Table structure for chat_msg
-- ----------------------------
DROP TABLE IF EXISTS "chat_msg";

CREATE TABLE "chat_msg" (
	"id" VARCHAR ( 64 ) NOT NULL,--id
	"content_type" INTEGER ( 10 ) NOT NULL,--内容类型
	"content" TEXT NOT NULL,--内容
	"from_id" VARCHAR ( 64 ) NOT NULL,--fromid
	"to_id" VARCHAR ( 64 ) NOT NULL,--toid
	"ptime" INTEGER ( 10 ) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--创建时间
	"is_with_draw" INTEGER ( 10 ) NOT NULL DEFAULT ( 0 ),--是否撤回
	"is_read" INTEGER ( 10 ) NOT NULL DEFAULT ( 0 ),--是否已读
	"record_id" VARCHAR ( 64 ) NOT NULL,--房间id
	PRIMARY KEY ( "id" ),--主键 id
	FOREIGN KEY ( "record_id" ) REFERENCES "chat_record" ( "id" ) ON DELETE CASCADE ON UPDATE CASCADE,--外键关联 record_id chat_record 的id 级联删除
	UNIQUE ( "id" ASC )--唯一键 id
);
PRAGMA foreign_keys = true;
