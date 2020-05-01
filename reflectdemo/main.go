package main

import (
	"fmt"
	"reflect"
)

type Category struct {
	ID string
	Name string
	Description  string
}

func main() {
	var category Category
	var i int32 = -5

	

	// 类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型
	// 种类（Kind）指的是对象归属的品种. 在go语言中是固定定义好的，可以通过查看Kind定义知道全部Kind
	typeInfo, kindInfo := reflectTypeAndKind(category)
	fmt.Printf("struct type reflect info.            name:%s, kind:%s \n", typeInfo, kindInfo)
	typeOfCagegory := reflect.TypeOf(category)
	typeInfo, kindInfo = reflectTypeAndKind(typeOfCagegory)
	fmt.Printf("type reflect info of reflect.        name:%s, kind:%s \n", typeInfo, kindInfo)
	typeInfo, kindInfo = reflectTypeAndKind(i)
	fmt.Printf("primitive type reflect info.         name:%s, kind:%v \n", typeInfo, kindInfo)


	myFunc := reflectTypeAndKind
	typeInfo, kindInfo = reflectTypeAndKind(myFunc)
	fmt.Printf("func type reflect info.              name:%s, kind:%v \n", typeInfo, kindInfo)
}

func reflectTypeAndKind(x interface{}) (typeInfo string, kindInfo reflect.Kind) {
	reflectObject := reflect.TypeOf(x)
	return reflectObject.Name(), reflectObject.Kind()
}


