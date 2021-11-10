package message

const  (
	LoginMesType = "LoginMes"
	LoginResType = "LoginRes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}

type LoginMessage struct {
	UserId string   `json:"user_id"`//用户id
	UserPwd string  `json:"user_pwd"`//用户密码
	UserName string `json:"user_name"`//用户名
}

type LoginResMessage struct {
	Code int     `json:"code"`//返回状态码 500 表示该用户未注册 200 表示登录成功
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {

}