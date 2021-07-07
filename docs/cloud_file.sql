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
-- Table structure for cloud_file
-- ----------------------------
DROP TABLE IF EXISTS "cloud_file";

CREATE TABLE "cloud_file" (
	"id" VARCHAR ( 64 ) NOT NULL,--id
	"user_id" VARCHAR ( 64 ) NOT NULL,--用户id
	"file_name" VARCHAR ( 128 ),--文件名字
	"parent_id" VARCHAR ( 64 ) NOT NULL,--父id
	"ptime" INTEGER ( 10 ) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--创建时间
	"file_cid" VARCHAR ( 128 ),--文件cid
	"file_size" INTEGER ( 10 ),--文件大小
	"file_type" INTEGER ( 10 ),--文件类型
	"is_folder" INTEGER ( 10 ) NOT NULL DEFAULT ( 0 ),--是否是文件 还是文件夹 0 文件 1文件夹
	"thumbnail" VARCHAR ( 128 ) NOT NULL,--缩略图
	PRIMARY KEY ( "id" ),--主键 id
	FOREIGN KEY ( "user_id" ) REFERENCES "sys_user" ( "id" ) ON DELETE NO ACTION ON UPDATE NO ACTION,--外键关联user_id sys_user id
	UNIQUE ( "id" ASC ) --唯一键 id
);
PRAGMA foreign_keys = true;
