package domain

/*
  示例而已，因此字段只有几个
*/
type User struct {
	UserId int64
	UserCreateReq
	//etc ..
}

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
