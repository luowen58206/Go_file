package main

import (
	"fmt"
	"log"
	"net"
)


////ReadPkg 封装读取客户端发送的消息
//func ReadPkg(conn net.Conn) (mes message.Message,err error)  {
//	buf := make([]byte,8096)
//	fmt.Println("conn start read ")
//	//在conn没有被关闭的情况下 才会阻塞
//	_,err = conn.Read(buf[:4])
//	if  err != nil {
//		//err = errors.New("read pakLen fail error")
//		return
//	}
//
//	//根据buf[:4] 转成一个 uint32
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4])
//
//	//根据 pkgLen读取数据 从conn都pkgLen个字节到buf里面去
//	n,err :=conn.Read(buf[:pkgLen])
//	if uint32(n) != pkgLen || err != nil {
//		//err = errors.New("conn read dataJson error")
//		return
//	}
//
//	//反序列化 -》 message。Message类型
//	err =json.Unmarshal(buf[:pkgLen],&mes)
//	if err != nil {
//		err = errors.New("dataJson unMarshal error")
//		return
//	}
//	return
//}
//
////WritePkg 写入信息
//func WritePkg(conn net.Conn, data []byte) (err error) {
//	//先发送一个长度给对方
//	var dataLen uint32
//	dataLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4],dataLen)
//
//	//发送长度
//	n,err := conn.Write(buf[:4])
//	if n != 4 || err != nil{
//		log.Fatalln("conn.Write err = ",err)
//		return
//	}
//
//	//发送data 本身
//	n,err = conn. Write(data)
//	if uint32(n) != dataLen || err != nil{
//		log.Fatalln("conn.Write err = ",err)
//		return
//	}
//	return
//}




////ServerProcessLogin 编写一个函数，专门处理登录请求
//func ServerProcessLogin(conn net.Conn,mes *message.Message) (err error) {
//	//先从mes 中出去data 并直接反序列化为loginMes
//	var  loginMes message.LoginMessage
//	err = json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err != nil {
//		log.Fatalln("json Unmarshal fail err = ",err)
//		return
//	}
//	//先声明一个resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResType
//
//	//再声明一个 loginResMes
//	var loginResMes message.LoginResMessage
//
//	//如果id = 100 ， 密码 = 123456 合法
//	if loginMes.UserId == "100" && loginMes.UserPwd == "tom" {
//		//合法
//		loginResMes.Code = 200
//	} else {
//		//不合法
//		loginResMes.Code = 500
//		loginResMes.Error = "user not exits"
//	}
//
//	//将 loginResMes 序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json Marshal fail err ",err)
//		return
//	}
//
//	//将data赋值给 resMes
//	resMes.Data = string(data)
//
//	//对 resMes 序列化 准备发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json Marshal fail err ",err)
//		return
//	}
//
//	//发送data 封装到WritePkg
//	err = WritePkg(conn,data)
//	return
//}

////ServerProcessMes 函数根据客户端发送的消息种类，决定调用那个函数来解决
//func ServerProcessMes(conn net.Conn,mes *message.Message) (err error)  {
//	switch mes.Type {
//		case message.LoginMesType:
//			err = ServerProcessLogin(conn,mes)
//		case message.RegisterMesType:
//
//		default:
//			fmt.Println("info Type not exist cont deal")
//	}
//	return
//}

//处理和客户端的通讯
func process(conn net.Conn)  {
	defer conn.Close()

	//创建一个总控
	processor := &Processor{
		Conn:conn,
	}
	err := processor.Process02()
	if err != nil {
		log.Fatalln("processor.Process02() error = ",err)
		return
	}

	//读客户端发送的信息，我们直接封装成一个函数ReadPkg
	//for  {
	//		mes,err := ReadPkg(conn)
	//		if err != nil {
	//			if err == io.EOF {
	//				fmt.Println("client exit server exit")
	//				return
	//			}else {
	//				log.Fatalln("readPkg fail err = ", err)
	//				return
	//			}
	//		}
	//		//fmt.Println("mess = ",mes)
	//		err = ServerProcessMes(conn,&mes)
	//		if err != nil {
	//			return
	//		}
	//}
}

func main()  {
	fmt.Println("server inNew listen in 9999 port")
	listen,err := net.Listen("tcp","0.0.0.0:9999")
	defer listen.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}

	//一旦监听成功，等待客户端链接
	for  {
		fmt.Println("waiting client link")
		conn,err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		//一旦链接成功 启动一个协程和客户端保持通讯
		go process(conn)
	}

}
