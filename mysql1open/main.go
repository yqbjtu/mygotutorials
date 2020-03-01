package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

// 定义一个全局对象db
var db *sql.DB

type user struct {
	id uint16
	username string
	usertype int
}
// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:rati0nal@tcp(127.0.0.1:3306)/db01"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询单条数据示例
 func queryRowDemo(userIdFrom, userIdTo int) (time.Duration,error) {
	before := time.Now() //获取当前时间
	sqlStr := "select id, username, usertype from user where id=?"
	for i:=userIdFrom; i< userIdTo; i++ {
		var u user
		// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
		err := db.QueryRow(sqlStr, i).Scan(&u.id, &u.username, &u.usertype)
		if err != nil {
			fmt.Printf("scan failed, %d, err:%v\n", i, err)
		} else {
			fmt.Printf("id:%d name:%s usertype:%d\n", u.id, u.username, u.usertype)
		}
	}

	after := time.Now() //获取当前时间
	fmt.Printf("Time duration:%d \n", after.Sub(before) )
	 return after.Sub(before), nil;
 }

// 查询单条数据示例
func queryRowDemoByPrepare(userIdFrom, userIdTo int) (time.Duration, error) {
	before := time.Now() //获取当前时间
	sqlStr := "select id, username, usertype from user where id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return  0, err
	}
	defer stmt.Close()

	for i:=userIdFrom; i< userIdTo; i++ {
		rows, err := stmt.Query(i)
		if err != nil {
			fmt.Printf("query failed, %d, err:%v\n", i, err)
		}
		defer rows.Close()
		// 循环读取结果集中的数据
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.usertype)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
			} else {
				fmt.Printf("id:%d name:%s usertype:%d\n", u.id, u.username, u.usertype)
			}
		}
	}

	after := time.Now() //获取当前时间
	fmt.Printf("prepared Time duration:%d \n",  after.Sub(before))
	return after.Sub(before), nil;
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	dur1, err := queryRowDemo(0, 20)
	dur2, err := queryRowDemoByPrepare(0, 20)
	fmt.Printf("Final time:%d  %d\n", dur1, dur2)
}