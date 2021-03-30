package sql_init

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (db *sql.DB, err error) {
	// DSN格式为：用户名：密码@tcp(ip:端口)/数据库名称
	dsn := "root:123123@tcp(127.0.0.1:3306)/ngucdata"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return db, nil
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return db, nil
	}
	return db, nil
}
