package vo

const (
	DBversion = "db-version"
)

var Version = "3"

var UpgradeSql = map[int][]string{
	1: {
		"alter table sys_user add column role varchar(64) not null default(2)",
		"alter table article add column external_href text",
		"update sys_user set name = '小龙客服', nickname = '小龙客服', role = '1' where id = '416418922095452160'",
	},
	2: {
		`CREATE TABLE "user_friend" ( "user_id" STRING(64) NOT NULL, "friend_id" STRING(64) NOT NULL, "friend_nickname" STRING(128) NOT NULL DEFAULT '')`,
		`ALTER TABLE chat_msg ADD "send_state" INTEGER(10) NOT NULL DEFAULT 0`,
		`ALTER TABLE chat_msg ADD "send_fail" VARCHAR(64) NOT NULL DEFAULT ''`,
	},
}
