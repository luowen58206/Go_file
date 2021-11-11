package message


// User 定义一个用户结构体
type User struct {
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的json串 和 结构体字段对应 tag 名字一致

	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}