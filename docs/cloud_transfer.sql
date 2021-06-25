/*
 Navicat Premium Data Transfer

 Source Server         : test-cloud
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/06/2021 13:28:16
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for user_download1
-- ----------------------------
DROP TABLE IF EXISTS "cloud_transfer";

CREATE TABLE "cloud_transfer" (
  "id" VARCHAR NOT NULL,--id
  "user_id" varchar (64) NOT NULL,--用户id
  "file_name" varchar (128) NOT NULL,--文件名字
  "ptime" integer(64) NOT NULL DEFAULT 0,--创建时间
  "file_cid" varchar (64) NOT NULL,--文件cid
  "file_size" int (10) NOT NULL,--文件大小
  "down_path" varchar (128) NOT NULL,--下载路径
  "file_type" int (10) NOT NULL,--文件类型
  "transfer_type" integer(10) NOT NULL,--传输类型 0 上传 1 下载
  "upload_parent_id" varchar(128) NOT NULL,--上传的父id
  "upload_file_id" varchar(128) NOT NULL,--上传文件id
  PRIMARY KEY ("id"),--主键id
  FOREIGN KEY ("user_id") REFERENCES "sys_user" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,--外键关联 user_id sys_user 的id
  UNIQUE ("id" ASC)--唯一键id
);
PRAGMA foreign_keys = true;
