package main

import (
	"example.com/m/common/message"
	"example.com/m/server/processes"
	"example.com/m/server/utils"
	"fmt"
	"io"
	"log"
	"net"
)

type Processor struct {
	Conn net.Conn
}


//ServerProcessMes 函数根据客户端发送的消息种类，决定调用那个函数来解决
func (this *Processor) ServerProcessMes(mes *message.Message) (err error)  {
	switch mes.Type {
	case message.LoginMesType:
		//创建一个userProcess实例
		up := &processes.UserProcess{
		Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:

	default:
		fmt.Println("info Type not exist cont deal")
	}
	return
}

func (this *Processor) Process02 () (err error) {

	for  {
		var tf = utils.Transfer{
			Conn: this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit server exit")
				return err
			}else {
				log.Fatalln("readPkg fail err = ", err)
				return err
			}
		}
		//fmt.Println("mess = ",mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	
}