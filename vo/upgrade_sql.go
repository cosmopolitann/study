package vo

const (
	DBversion = "db-version"
)

var Version = "2"

var UpgradeSql = map[int][]string{
	1: []string{"alter table sys_user add column role varchar(64) not null default(2)",
		"alter table article add column external_href text",
		"INSERT INTO sys_user (id, peer_id, name, phone, sex, ptime, utime, nickname, img) VALUES ('414207114215428096', '', '星河飞天-人工客服1','', 0, 1627444008, 1627444008, '人工客服1', '')"},
}
