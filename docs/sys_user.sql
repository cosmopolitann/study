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
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "sys_user";

CREATE TABLE "sys_user" (
	"id" VARCHAR ( 64 ) NOT NULL,--id
	"peer_id" VARCHAR ( 64 ) NOT NULL,--用户id==peerid
	"name" VARCHAR ( 128 ),--文件名字
	"phone" VARCHAR ( 64 ),--手机号
	"sex" INTEGER ( 10 ) DEFAULT ( 1 ),--性别 0 未知 1男 2 女
	"ptime" INTEGER ( 10 ) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--创建时间
	"utime" INTEGER ( 10 ) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--更新时间
	"nickname" VARCHAR ( 128 ),--昵称
	"img" VARCHAR ( 128 ),--图片
	PRIMARY KEY ( "id" ),--主键索引 id
	UNIQUE ( "id" ASC ) --唯一键id
);
PRAGMA foreign_keys = true;
