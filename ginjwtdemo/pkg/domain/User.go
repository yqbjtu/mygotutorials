package domain

import "fmt"

/*
  示例而已，因此字段只有几个
*/
type UserCreateReq struct {
	UserName    string
	Email       string
	Phone       string
	LockedState bool
	//etc ..
}

type LoginReq struct {
	UserName string `json:"userName"`
	//实际当中不会以明文传输密码，本工程是示例工程，为简单起见使用明文
	Passwd string `json:"passwd"`
}

// LoginResult 登录结果结构
type LoginResp struct {
	Token string `json:"token"`
}

/*
  示例而已，因此字段只有几个
*/
type User struct {
	UserId int64
	UserCreateReq
	//etc ..
}

func init() {
	user := &User{UserId: 0,
		UserCreateReq: UserCreateReq{
			UserName: "tom"}}
	fmt.Printf("user:%+v\n", user)
}
