package router

import (
	mycontroller "ginjwtdemo/pkg/controller"
	myjwt "ginjwtdemo/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigRouter(router *gin.Engine) {
	userController := mycontroller.NewUserController()
	router.Use(myjwt.JWTAuth())
	//router.Static("/static", )
    router.LoadHTMLFiles("./web/index.html")
    router.GET("/", index)

	router.POST("/login", userController.Login)
	router.GET("/users", userController.GetAllUsers)
	router.GET("/usersfind", userController.FindUsers)
	router.GET("/users/:userId", userController.GetOneUser)
	router.PUT("/users", userController.CreateOneUser)
	router.POST("/users/:userId", userController.UpdateOneUser)
	router.DELETE("/users/:userId", userController.DeleteOneUser)
}

func index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": demo主页",
    })
}