package router

import (
	mycontroller "ginportdemo/pkg/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.Engine) {
	userController := mycontroller.NewUserController()
	router.GET("/users", userController.GetAllUsers)
	router.GET("/usersfind", userController.FindUsers)
	router.GET("/users/:userId", userController.GetOneUser)
	router.PUT("/users", userController.CreateOneUser)
	router.POST("/users/:userId", userController.UpdateOneUser)
	router.DELETE("/users/:userId", userController.DeleteOneUser)
}
