package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	//	1.读取conf配置信息
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	//	2.组织连接数据库的字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	//	3.连接数据库
	db, err := sql.Open(dbDriver, connUrl)
	//	4，获取数据库的对象，处理连接结果
	if err != nil {
		panic("数据库连接错误，请配置检查")
	}
	//为全局的数据库操作对象赋值
	Db = db
}
