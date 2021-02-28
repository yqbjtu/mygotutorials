package controller

import (
	"fmt"
	mydomain "ginjwtdemo/pkg/domain"
	myjwt "ginjwtdemo/pkg/middleware"
	"ginjwtdemo/util"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"net/http"
	"time"
)

type UserController struct {
	// service or some to access DB method
}

func NewUserController() *UserController {
	controller := UserController{}
	return &controller
}

func (ctl *UserController) Login(c *gin.Context) {
	klog.Infof("login to get a token")
	var loginReq mydomain.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err == nil {
		//实际当中需要检查用户名和密码的正确性，这里为了简单起见，hardcode，只要和用户是tom，密码是123456就允许通过
		// check whether username exists and passwd is matched
		if loginReq.UserName == "tom" && loginReq.Passwd == "123456" {
			user := mydomain.User{}
			user.UserName = loginReq.UserName
			user.UserId = 0
			generateToken(c, user, "admin", 30)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败, 用户不存在或者密码不正确",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败." + err.Error(),
		})
	}
}

/*
  此工程为了简单，直接将生成token放在controller中
  有效时间长度，单位是分钟
*/
func generateToken(c *gin.Context, user mydomain.User, roleId string, expiredTimeByMinute int64) {
	j := &myjwt.JWT{
		[]byte(util.SignKey),
	}
	claims := myjwt.CustomClaims{
		user.UserId,
		user.UserName,
		roleId,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),                   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + expiredTimeByMinute*60), // 过期时间 一小时
			Issuer:    "ginjwtdemo",                                      //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	data := mydomain.LoginResp{
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

func (ctl *UserController) CreateOneUser(c *gin.Context) {
	klog.Infof("create one user")
	var req mydomain.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//	user := mydomain.User{UserId: 1, UserName: userName} 会报cannot use promoted field UserCreateReq.UserName in struct literal of type domain.User
	// 这种太繁琐
	//user := mydomain.User{}
	//user.UserName = req.UserName
	//user.UserId = 0

	user := mydomain.User{UserId: 0, UserCreateReq: mydomain.UserCreateReq{UserName: req.UserName}}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
		"msg":    "create user successfully",
	})
}

func (ctl *UserController) GetAllUsers(c *gin.Context) {
	claimsFromContext, _ := c.Get(util.Gin_Context_Key)
	claims := claimsFromContext.(*myjwt.CustomClaims)
	currentUser := claims.UserName
	klog.Infof("get all users, loginUser:%q", currentUser)
	var users []mydomain.User
	for i := 0; i < 3; i++ {
		userName := fmt.Sprintf("tom%d", i)
		user := mydomain.User{UserId: 1}
		user.UserName = userName
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}

func (ctl *UserController) GetOneUser(c *gin.Context) {
	userId := c.Param("userId")
	klog.Infof("get one user by id %q", userId)
}

/*
  // 匹配的url格式:  /usersfind?username=tom&email=test1@163.com
*/
func (ctl *UserController) FindUsers(c *gin.Context) {
	userName := c.DefaultQuery("username", "张三")
	email := c.Query("email")
	// 执行实际搜索，这里只是示例
	c.String(http.StatusOK, "search user by %q %q", userName, email)
}

func (ctl *UserController) UpdateOneUser(c *gin.Context) {
	userId := c.Param("userId")
	klog.Infof("update user by id %q", userId)
}

func (ctl *UserController) DeleteOneUser(c *gin.Context) {
	userId := c.Param("userId")
	klog.Infof("delete user by id %q", userId)

}
