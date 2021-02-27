package controller

import (
	"fmt"
	mydomain "ginportdemo/pkg/domain"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"net/http"
)

type UserController struct {
	// service or some to access DB method
}

func NewUserController() *UserController {
	controller := UserController{}
	return &controller
}

func (c *UserController) CreateOneUser(context *gin.Context) {
	klog.Infof("create one user")
}

func (c *UserController) GetAllUsers(context *gin.Context) {
	klog.Infof("get all user")
	//H is a shortcut for map[string]interface{}

	var users []mydomain.User
	for i := 0; i < 3; i++ {
		userName := fmt.Sprintf("tom%d", i)
		user := mydomain.User{UserId: 1, UserName: userName}
		users = append(users, user)
	}
	context.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}

func (c *UserController) GetOneUser(context *gin.Context) {
	userId := context.Param("userId")
	klog.Infof("get one user by id %v", userId)
}

/*
  // 匹配的url格式:  /users/find?username=tom&email=test1@163.com
*/
func (c *UserController) FindUsers(context *gin.Context) {
	userId := context.Param("userId")
	klog.Infof("get one user by id %v", userId)

	userName := context.DefaultQuery("username", "张三")
	email := context.Query("email")
	// 执行实际搜索，这里只是示例
	context.String(http.StatusOK, "search user by %s %s", userName, email)
}

func (c *UserController) UpdateOneUser(context *gin.Context) {
	userId := context.Param("userId")
	klog.Infof("update user by id %v", userId)
}

func (c *UserController) DeleteOneUser(context *gin.Context) {
	userId := context.Param("userId")
	klog.Infof("delete user by id %v", userId)

}
