package main

import "fmt"

type Address struct {
	province string
	city     string
}

type User struct {
	name string
	age  int
	Address
}

func main() {

	user := &User{
		name: "tom",
		age:  18,
		Address: Address{
			province: "taiwan",
			city:     "taibei",
		},
	}

	fmt.Printf("user:%+v", user)

}
