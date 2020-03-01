package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "u1"
	password = "wwwwww"
	dbname   = "db1"
)

type User struct {
	id uint16
	username string
	fullName string
	email string
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
	query(db)
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





