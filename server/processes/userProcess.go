package processes

import (
	"encoding/json"
	"example.com/m/common/message"
	"example.com/m/server/utils"
	"fmt"
	"log"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//ServerProcessLogin 编写一个函数，专门处理登录请求
func (this *UserProcess)ServerProcessLogin(mes *message.Message) (err error) {
	//先从mes 中出去data 并直接反序列化为loginMes
	var  loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		log.Fatalln("json Unmarshal fail err = ",err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResType

	//再声明一个 loginResMes
	var loginResMes message.LoginResMessage

	//如果id = 100 ， 密码 = 123456 合法
	if loginMes.UserId == "100" && loginMes.UserPwd == "tom" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "user not exits"
	}

	//将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json Marshal fail err ",err)
		return
	}

	//将data赋值给 resMes
	resMes.Data = string(data)

	//对 resMes 序列化 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal fail err ",err)
		return
	}

	//发送data 封装到WritePkg
	//我们先创建一个transfer实例
	var tf = utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}