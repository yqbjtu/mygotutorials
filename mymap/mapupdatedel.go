package main

import "fmt"


func main() {
	var myCurrentMap map[string]int
	myCurrentMap = make(map[string]int)
	myCurrentMap["red"] = 1
	myCurrentMap["yellow"] = 2
	myCurrentMap["green"] = 3

	for k,v := range myCurrentMap {
		fmt.Printf("key:%s, value:%d \n", k,v)
	}
	fmt.Printf("---遍历结束-----\n")

	blueValue, ok := myCurrentMap["blue"]
	if ok {
		fmt.Printf("key:blue, value is %d \n", blueValue)
	} else {
		fmt.Println("there is no key:blue")
	}

	whiteValue := myCurrentMap["white"]
	fmt.Printf("key:white, value is %d \n", whiteValue)

	for k, _ := range myCurrentMap {
		if k == "yellow" {
			myCurrentMap[k] = 22
		}
	}

	// 修改map后再次遍历
	fmt.Printf("---更新一个key后遍历-----\n")
	for k,v := range myCurrentMap {
		fmt.Printf("key:%s, value:%d \n", k,v)
	}

	for k, _ := range myCurrentMap {
		if k == "yellow" {
			delete(myCurrentMap, k)
		}
	}

	fmt.Printf("---删除一个key后遍历-----\n")
	for k,v := range myCurrentMap {
		fmt.Printf("key:%s, value:%d \n", k,v)
	}
}


/*
执行结果
key:red, value:1
key:yellow, value:2
key:green, value:3
---遍历结束-----
there is no key:blue
key:white, value is 0
---更新一个key后遍历-----
key:red, value:1
key:yellow, value:22
key:green, value:3
---删除一个key后遍历-----
key:red, value:1
key:green, value:3
*/
