package main

import (
	"example.com/m/client"
	"fmt"
)

var userID string
var userPwd string

func main()  {
	ShowIu()
}

func ShowIu()  {
	//接收用户的选择
	var key int

	//判断是否继续显示菜单
	var loop = true

	for loop {
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
			loop = false
		case 2:
			fmt.Println("register new user")
			loop = false
		case 3:
			fmt.Println("Bye to chat system")
			loop = false
		default:
			fmt.Println("Please input a current num")
		}
	}
	//根据用户的输入 来设置不同的二级菜单的功能
	if key == 1{
	  //说明用户要登陆
		fmt.Print("please input userID:")
		fmt.Scanf("%s\n",&userID)
		fmt.Print("please input userPwd:")
		fmt.Scanf("%s\n",&userPwd)
		//先把登陆的函数写到另外的文件里面 如login，go
		client.Login(userID, userPwd)

	}else {

	}
}
