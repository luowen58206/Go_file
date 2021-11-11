package main

import (
	"example.com/m/client/processes"
	"fmt"
)

var userId string
var userPwd string
var userName string

func main()  {
	ShowIu()
}

func ShowIu()  {
	//接收用户的选择
	var key int

	//判断是否继续显示菜单
	//var loop = true

	for true {
		fmt.Println("------Welcome to log in to the chat system------")
		fmt.Println("\t\t\t 1: log in system")
		fmt.Println("\t\t\t 2: register")
		fmt.Println("\t\t\t 3: exit")
		fmt.Println("\t\t\t  please choose 1~3: ")
		fmt.Println("-----------------------------")

		fmt.Scanf("%d\n",&key)

		switch key {
		case 1:
			fmt.Println("Login to chat system")
			fmt.Print("please input userID:")
			fmt.Scanf("%s\n",&userId)
			fmt.Print("please input userPwd:")
			fmt.Scanf("%s\n",&userPwd)

			// 完成登录
			//创建一个UserProcess
			up := &processes.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("register new user")
			fmt.Print("please input a new userId:")
			fmt.Scanf("%s\n",&userId)
			fmt.Print("please input a new password:")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Print("please input a new name:")
			fmt.Scanf("%s\n",&userName)

			//调用实例完成注册请求
			up := &processes.UserProcess{}
			up.Register(userId,userPwd,userName)
			//loop = false
		case 3:
			fmt.Println("Bye to chat system")
			//loop = false
		default:
			fmt.Println("Please input a current num")
		}
	}
}
