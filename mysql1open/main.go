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
			fmt.Printf("scan failed,  %d, err:%v\n", i, err)
		} else {
			fmt.Printf("id:%d name:%s usertype:%d\n", u.id, u.username, u.usertype)
		}
	}

	after := time.Now() //获取当前时间
	dur := after.Sub(before)
	fmt.Printf("Time duration:%d \n", dur )
	fmt.Printf("[%v]Seconds [%v]ms\n", dur.Seconds(), dur.Milliseconds())
	return dur, nil;
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

	dur := after.Sub(before)
	fmt.Printf("prepared Time duration:%d \n", dur )
	fmt.Printf("[%v]Seconds [%v]ms\n", dur.Seconds(), dur.Milliseconds())
	return dur, nil;
}

// 预处理插入示例
func prepareInsertDemo(username string, usertype int) {
	sqlStr := "insert into user(username, usertype) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare insert failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(username, usertype)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theLastId, err := result.LastInsertId() // 新插入数据的id
	fmt.Println("insert successfully. id:", theLastId)
}

// 更新数据
func updateRowDemo(username string, active int) {
	//sqlStr := "update user set active=? where username = ?"
	//ret, err := db.Exec(sqlStr, active, username)

	sqlStr := "update user set active=? where username = ?"

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare update failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(active,username)
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update successfully, affected rows:%d\n", n)
}

// 更新数据
func deleteRowDemo(username string, active int) {
	sqlStr := "delete from user where username = ? and active=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare delete failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(username, active)
	if err != nil {
		fmt.Printf("delete exec failed, err:%v\n", err)
		return
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete successfully, affected rows:%d\n", n)
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	queryRowDemo(0, 5)
	//dur1, err := queryRowDemo(0, 20)
	//dur2, err := queryRowDemoByPrepare(0, 20)
	//fmt.Printf("Final time:%d  %d\n", dur1, dur2)

	prepareInsertDemo("nihaoGo",1)

	updateRowDemo("nihaoGo",1)
	//deleteRowDemo("nihaoGo",1)
	transactionDemo(19)
}

// 事务操作示例
func transactionDemo(userId int) {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("Failed to begin for trans, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set active=30 where id=?"
	_, err = tx.Exec(sqlStr1, userId)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("Failed to exec sql1, err:%v\n", err)
		return
	}
	sqlStr2 := "Update user1 set active=40 where id=?"
	_, err = tx.Exec(sqlStr2, userId)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("Failed to exec sql2, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("Failed to commit, err:%s\n", err)
		return
	}
	fmt.Println("trans successes!")
}
