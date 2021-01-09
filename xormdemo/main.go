package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"k8s.io/klog"
	"xormdemo/models"
)

func NewDbEngine() *xorm.Engine {
	// url格式：[username]:[password]@tcp([ip]:[port])/[database]?charset=utf8
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		"root",
		"rati0nal",
		"127.0.0.1",
		3306,
		"db02")
	// 创建Engine实例
	engine, err := xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	// 显示sql
	engine.ShowSQL(true)
	return engine
}

const Tony = "tony"

func main() {
	// 创建xorm.Engine实例
	dbEngine := NewDbEngine()
	user := models.PoUser{Username: "tom2"}
	affectedRow, err := dbEngine.InsertOne(user)
	if err != nil {
		panic(err)
	} else {
		klog.Infof("affectedRow %d", affectedRow)
	}

	user = models.PoUser{Username: Tony}
	affectedRow, err = dbEngine.InsertOne(&user)
	if err != nil {
		panic(err)
	} else {
		klog.Infof("affectedRow %d", affectedRow)
	}

	newUser := new(models.PoUser)
	isGotten, err := dbEngine.Where("username = ?", Tony).Get(newUser)
	if err == nil {
		if isGotten {
			klog.Infof("get record:%v", newUser)
			newUser.Username = newUser.Username + "2"
			affectedRow, err = dbEngine.Where("id =?", newUser.Id).Update(newUser)
			klog.Infof("affectedRow:%d, updateErr:%v", affectedRow, err)
		} else {
			klog.Infof("no record by where username = %s", Tony)
		}
	} else {
		klog.Warningf("failed to get record, err:%v", err)
	}

	// 定义一个切片，保存所有用户
	datalist := make([]models.PoUser, 0)
	// 按照id的倒序查询用户
	klog.Info("start to select ")
	err = dbEngine.Desc("id").Find(&datalist)
	// 判断查询是否成功
	if err != nil {
		panic(err)
	}
	// 遍历输出结果
	for i := range datalist {
		fmt.Printf("%v \n", datalist[i])
	}

}
