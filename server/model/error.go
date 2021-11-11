package model

import "errors"

//自定义一些错误


var (
	ErrorUserNotExits  = errors.New("user not exits")
	ErrorUSerExits = errors.New("user yet exits")
	ErrorUserPassword = errors.New("wrong password")
)
