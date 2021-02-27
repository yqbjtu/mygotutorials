package domain

/*
  示例而已，因此字段只有几个
*/
type User struct {
	UserId      int64
	UserName    string
	Email       string
	Phone       string
	LockedState bool
	//etc ..
}
