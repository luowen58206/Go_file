package processes

import (
	"encoding/binary"
	"encoding/json"
	"example.com/m/client/utils"
	"example.com/m/common/message"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

// Login 写一个函数，完成登录
func (this *UserProcess)Login(userId string, userPwd string) (err error) {
	//指定一个协议

	//链接服务器
	conn,err := net.Dial("tcp","0.0.0.0:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//延时关闭
	defer conn.Close()

	//通过conn发消息给服务端
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建一个loginMess 填充消息
	var loginMes message.LoginMessage
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将loginMes序列化
	dataJson,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("jsonMarshal err = ",err)
		return
	}
	//把dataJson 赋给 mes.Data
	mes.Data = string(dataJson)

	//将mes进行序列化
	mesJson,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mesMarshal err = ",err)
		return
	}

	//mesJson 就是我们需要发送的数据
	//先把mesJson的长度发送给servers net Write发送的是一个切片类型的数据
	//必须先将data的长度-》成一个表示长度的byte切片
	var mesJsonLen uint32
	mesJsonLen = uint32(len(mesJson))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],mesJsonLen)

	//发送长度
	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write err = ",err)
		return
	}

	//提示发送数据成功
	fmt.Printf("conn send dataJsonLen success len(dataJson) = %d dataJson = %s\n",len(dataJson),string(dataJson))


	//发送消息本身
	_,err = conn.Write(mesJson)
	if  err != nil{
		fmt.Println("conn.Write(mesJson) err = ",err)
		return
	}

	//这里处理服务器端返回的消息
	var tf  = utils.Transfer{
		Conn: conn,
	}
	mes,err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err = ",err)
		return
	}

	//将mes.Data 部分反序列化成LoginResMes
	var loginResMes message.LoginResMessage
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println

		//这里需要起一个协程，该协程保持和服务端的通讯
		//如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go ServerProcessMessage(conn)

		//显示登录成功的菜单 循环显示i
		for {
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}
	return
}

// Register 注册函数
func (this *UserProcess) Register(userId , userPwd , userName string) (err error) {
	//链接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//延时关闭
	defer conn.Close()

	//通过conn发消息给服务端
	var mes message.Message
	mes.Type = message.RegisterMesType

	//创建一个loginMess 填充消息
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//将loginMes序列化
	dataJson, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("jsonMarshal err = ", err)
		return
	}
	//把dataJson 赋给 mes.Data
	mes.Data = string(dataJson)

	//将mes进行序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mesMarshal err = ", err)
		return
	}

	var tf = utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("register seed info err ", err)
	}

	//mes 就是RegisterResMes
	mes,err = tf.ReadPkg()
	if err != nil {
		fmt.Println("register take over err = ",err)
		return
	}

	var registerResMes message.RegisterRes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("register success")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

//ServerProcessMessage 和服务器端保持此通讯
func ServerProcessMessage(conn net.Conn)  {
	//创建一个transfer实例，不停的读取服务器发送的数据
	var tf = &utils.Transfer{
		Conn: conn,
	}
	for  {
		fmt.Println("Client waiting for server information ")
			mes,err := tf.ReadPkg()
			if err != nil {
				fmt.Println("tf.ReadPkg err = ",err)
				return
			}
			//如果读取到消息
			fmt.Printf("mes = %v \n",mes)
	}
}