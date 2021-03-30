package main

import (
	"Learn/09mysql/sql_init"
	"database/sql"
	"fmt"
)

//注入问题：
//任何时候都不应该自己拼接sql语句,例如：
//select *from user where name='xxx' or 1=1#
//此时可以获取到所以的用户信息，会导致信息泄露
var db *sql.DB

type user struct {
	id   int
	name string
}

// 预处理查询示例
func select_pre(args int) {
	sqlStr := "select * from user where id >= ?"
	stmt, err := db.Prepare(sqlStr) //将语句提前提交至服务器，等待后续的命令
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(args) //输入缺省值，返回命令获取的结果
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s\n", u.id, u.name)
	}
}

func main() {
	dbret, err := sql_init.InitDB() // 调用输出化数据库的函数
	//将初始化数据库得到的db对象赋值给当前的全局变量
	//当然也可直接使用init得到的db对象，只是需要传入crud函数中
	db = dbret
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	select_pre(1)
}
