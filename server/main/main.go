package main

import (
	"example.com/m/server/model"
	"fmt"
	"net"
)

func init()  {
	//当服务器启动时 我们就去初始化我们的redis链接池
	initPool("localhost:6379",16,0,300)
	initUserDao()
}

//处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	//创建一个总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process02()
	if err != nil {
		fmt.Println("processes.Process02() error = ", err)
		return
	}
}

//我们编写一个函数 希望完成对userDAO的初始化任务
func initUserDao()  {

	//这里的pool本事就是一个全局变量
	//这里需要注意初始化的顺序问题
	model.MyUserDao = model.NewUserDao(pool)
}

func main()  {
	fmt.Println("server inNew listen in 9999 port")
	listen,err := net.Listen("tcp","0.0.0.0:9999")
	defer listen.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//一旦监听成功，等待客户端链接
	for  {
		fmt.Println("waiting client link")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}

		//一旦链接成功 启动一个协程和客户端保持通讯
		go process(conn)
	}
}
