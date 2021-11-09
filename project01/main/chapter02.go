package main

import (
	"fmt"
	// "unsafe"
	// "reflect"
	// "strconv"cls
)
func main()  {
	//基本数据类型在内存布局
	var i int = 10
	//求i在内存中的地址
	fmt.Println("i的地址  ",&i)

	//1，ptr是一个指针变量
	//2，ptr的类型是 *int
	//3，ptr 本身的值是&i
	var prt *int = &i
	fmt.Println("\nptr 的值 = %v",prt)
	fmt.Println("\nptr的地址  ",&prt)
	fmt.Println("\n prt指向的值  " ,*prt)
}