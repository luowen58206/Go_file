package message

const  (
	LoginMesType = "LoginMes"
	LoginResType = "LoginRes"
	RegisterMesType = "RegisterMes"
	RegisterResType = "RegisterRes"
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}

type LoginMessage struct {
	UserId string   `json:"useId"`//用户id
	UserPwd string  `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}

type LoginResMessage struct {
	Code int     `json:"code"`//返回状态码
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {
	User  User `json:"user"`//类型就是User结构体
}

type RegisterRes struct {
	Code int     `json:"code"`//返回状态码 400 该用户被占用 200 成功
	Error string `json:"error"`//返回错误信息
}