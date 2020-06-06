package main

import (
	"fmt"
	"log"
)

func FilterByType(arr []interface{}, deviceType string, myFunc func(interface{}, string) bool) []interface{} {
	var result []interface{}
	if len(arr) == 0 {
		result = make([]interface{}, 0)
	} else {
		for _, v := range arr {
			if myFunc(v, deviceType) {
				result = append(result, v)
			}
		}
	}

	return result
}

type Device struct {
	Id         int
	Name       string
	DeviceType string
}

func main() {
	var devlist = []int{1, 2, 3, 4, 5}

	var _list []interface{}
	for _, v := range devlist {
		_list = append(_list, v)
	}
	var filteredNumList = FilterByType(_list, "", func(item interface{}, deviceType string) bool {
		if item.(int) > 3 {
			return true
		}
		return false
	})
	log.Println("filteredNumList", filteredNumList)

	var deviceList = []Device{
		{1, "bj01", "磨机"},
		{2, "bj02", "立磨"},
		{3, "xa03", "提升机"},
		{4, "xa04", "立磨"}}
	for _, v := range deviceList {
		_list = append(_list, v)
	}
	var filteredDevlist = FilterByType(_list, "立磨", func(item interface{}, deviceType string) bool {
		dev, ok := item.(Device)
		if ok {
			if dev.DeviceType == deviceType {
				return true
			}
		} else {
			fmt.Printf("non-device item %v\n", item)
		}

		return false
	})
	log.Println("filteredDevlist:", filteredDevlist)
}
