package client

import (
	"encoding/binary"
	"encoding/json"
	"example.com/m/common/message"
	"example.com/m/server/utils"
	"fmt"
	"log"
	"net"
)

// Login 写一个函数，完成登录
func Login(userId string, userPwd string) (err error) {
	//指定一个协议

	//链接服务器
	conn,err := net.Dial("tcp","0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln("jsonMarshal err = ",err)
		return
	}
	//把dataJson 赋给 mes.Data
	mes.Data = string(dataJson)

	//将mes进行序列化
	mesJson,err := json.Marshal(mes)
	if err != nil {
		log.Fatalln("mesMarshal err = ",err)
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
		log.Fatalln("conn.Write err = ",err)
		return
	}

	//提示发送数据成功
	fmt.Printf("conn send dataJsonLen success len(dataJson) = %d dataJson = %s",len(dataJson),string(dataJson))


	//发送消息本身
	_,err = conn.Write(mesJson)
	if  err != nil{
		log.Fatalln("conn.Write(mesJson) err = ",err)
		return
	}

	//这里处理服务器端返回的消息
	var tf  = utils.Transfer{
		Conn: conn,
	}
	mes,err = tf.ReadPkg()
	if err != nil {
		log.Fatalln("readPkg(conn) err = ",err)
		return
	}

	//将mes.Data 部分反序列化成LoginResMes
	var loginResMes message.LoginResMessage
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("login success")
	}else if loginResMes.Code == 500{
		fmt.Println(loginResMes.Error)
	}
	return
}
