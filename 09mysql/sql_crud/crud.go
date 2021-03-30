package main

import (
	"Learn/09mysql/sql_init"
	_ "Learn/09mysql/sql_init"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义一个全局对象db
var db *sql.DB

type user struct {
	id   int
	name string
}

func insert() {
	//可以忽略掉返回值Result，它存在两个方法
	//Result.LastInsertId() 返回语句更改的id
	//Result.RowsAffected() 返回语句更改的行数
	//但不能忽略err
	_, err := db.Exec("insert into user(name ) value(?)", "刘六")
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 插入数据
func insertDemo() {
	ret, err := db.Exec("insert into user(name) values (?)", "哈哈哈")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	therow, err := ret.RowsAffected() // 返回语句更改的行数
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", therow)
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select * from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s\n", u.id, u.name)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select * from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
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

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set name=? where id = ?"
	ret, err := db.Exec(sqlStr, "李四", 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
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
	//insert()
	//insertDemo()
	//queryRowDemo()
	//queryMultiRowDemo()
	//updateRowDemo()
	//deleteRowDemo()
	queryMultiRowDemo()

}
