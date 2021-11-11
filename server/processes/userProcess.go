package processes

import (
	"encoding/json"
	"example.com/m/common/message"
	"example.com/m/server/model"
	"example.com/m/server/utils"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("json Unmarshal fail err = ",err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResType
	var registerResMes message.RegisterRes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ErrorUSerExits{
			registerResMes.Code = 505
			registerResMes.Error = model.ErrorUSerExits.Error()
		}else {
			registerResMes.Code = 506
			registerResMes.Error = "unknown error"
		}
	}else {
		registerResMes.Code = 200
	}

	//将 loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json Marshal fail err ",err)
		return
	}
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


//ServerProcessLogin 编写一个函数，专门处理登录请求
func (this *UserProcess)ServerProcessLogin(mes *message.Message) (err error) {
	//先从mes 中出去data 并直接反序列化为loginMes
	var  loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json Unmarshal fail err = ",err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResType

	//再声明一个 loginResMes
	var loginResMes message.LoginResMessage

	//我们需要到redis数据库去验证
	//使用model。MyuserDao 到redis去验证
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	//fmt.Println("user = ",user)
	if err != nil {
		if err == model.ErrorUserNotExits {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPassword {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "Server internal error "
		}

	} else {
		loginResMes.Code = 200
		fmt.Println(user,"login success")
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