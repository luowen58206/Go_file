package processes

import (
	"fmt"
	"os"
)

//ShowMenu 显示登录成功后的界面...
func ShowMenu()  {
	fmt.Println("---------Congratulations XXX  on successful login--------")
	fmt.Println("---------1 Display list of online users-----------")
	fmt.Println("---------2 Seed Message--------------")
	fmt.Println("---------3 Message history-----------")
	fmt.Println("---------1 Exit System---------------")
	fmt.Println("---------Please choose 1 ~ 4---------")
	var key int
	fmt.Scanf("%d\n",&key)

	switch key {
	case 1:
		fmt.Println("Display list of online users")
	case 2:
		fmt.Println("Seed Message")
	case 3:
		fmt.Println("Message history")
	case 4:
		fmt.Println("Exit System")
		os.Exit(0)
	default:
		fmt.Println("Your input is wrong, please re-enter")
	}
}
