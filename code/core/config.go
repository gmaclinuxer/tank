package core

const (
	//用户身份的cookie字段名
	COOKIE_AUTH_KEY = "_ak"

	//数据库表前缀 tank200表示当前应用版本是tank:2.0.x版，数据库结构发生变化必然是中型升级
	TABLE_PREFIX = "tank20_"

	//当前版本
	VERSION = "2.0.0"
)

type Config interface {

	//是否已经安装
	IsInstalled() bool
	//启动端口
	GetServerPort() int
	//获取mysql链接
	GetMysqlUrl() string

	//文件存放路径
	GetMatterPath() string
	//完成安装过程，主要是要将配置写入到文件中
	FinishInstall(mysqlPort int, mysqlHost string, mysqlSchema string, mysqlUsername string, mysqlPassword string)
}