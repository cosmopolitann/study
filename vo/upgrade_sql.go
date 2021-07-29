package vo

const (
	DBversion = "db-version"
)

var Version = "2"

var UpgradeSql = map[int][]string{
	1: []string{"alter table sys_user add column role varchar(64) not null default(2)",
		"alter table article add column external_href text",
		"update sys_user set name = '小龙客服', nickname = '小龙客服', role = '1' where id = '416418922095452160'"},
}
