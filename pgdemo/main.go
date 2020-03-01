package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "u1"
	password = "rati0nal"
	dbname   = "db1"
)

type User struct {
	id uint16
	username string
	fullName string
	email string
}

func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//query(db)
	insertUser(db)
	updateUser(db, 2, "newMail@fb.com")
	deleteUser(db,3)
}


func query(db *sql.DB){
	var u User

	rows,err:=db.Query(" select * from  yqschema1.user where name=$1","Paul")

	if err!= nil{
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next(){
		err:= rows.Scan(&u.id, &u.username, &u.fullName, &u.email)
		if err!= nil{
			fmt.Println(err)
		} else {
			fmt.Printf("id:%d name:%s fullName:%s, email:%s\n", u.id, u.username, u.fullName, u.email)
		}
	}

	err = rows.Err()
	if err!= nil{
		fmt.Println(err)
	}

	fmt.Println("query end")
}

//
func insertUser(db *sql.DB)  {
	stmt,err := db.Prepare("insert into yqschema1.user(name,fullname,email) values($1,$2,$3)")
	if err != nil {
		log.Fatal(err)
	}
	_,err = stmt.Exec("user01","full user01","user01@qq.com")
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("insert into user successfully")
	}
}

func updateUser(db *sql.DB, userId int, email string) {
	stmt,err := db.Prepare("UPDATE  yqschema1.user set email=$1 WHERE  id=$2")
	if err != nil {
		log.Fatal(err)
	}
	result,err := stmt.Exec(email, userId)
	if err != nil {
		log.Fatal(err)
	}else {
		n, _ := result.RowsAffected() // 操作影响的行数
		fmt.Println("udpate user successfully. RowsAffected:", n)
	}

}

func deleteUser(db *sql.DB, userId int) {
	stmt,err := db.Prepare("DELETE FROM yqschema1.user WHERE  id=$1")
	if err != nil {
		log.Fatal("failed to prepare delete sql", err)
	}

	result,err := stmt.Exec(userId)
	if err != nil {
		log.Fatal(err)
	}else {
		n, _ := result.RowsAffected() // 操作影响的行数
		fmt.Println("delete form user successfully. RowsAffected:", n)
	}
}



