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
	klog.Infof("get one user")
}

func (c *UserController) UpdateOneUser(context *gin.Context) {
	klog.Infof("update user")
}

func (c *UserController) DeleteOneUser(context *gin.Context) {
	klog.Infof("delete user")

}
