package model

import (
	"encoding/json"
	"example.com/m/common/message"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// MyUserDao 我们在服务器启动后，就初始化一个userDAO实例
//全局变量 在需要和redis交互是 就直接使用即可
var (
	MyUserDao *UserDao
)


// UserDao 定义一个userDao结构体
//完成对User 的一系列操作
type UserDao struct {
	pool *redis.Pool
}

//NewUserDao 使用工厂模式 创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool : pool,
	}
	return
}

// GetUserById 根据一个userID 返回一个user实例或者error
func (this *UserDao) GetUserById(conn redis.Conn,id string) (user *User, err error)  {

	//通过给定id 去链接池查询id
	res,err := redis.String(conn.Do("hget","users",id))
	if err != nil {
		if err == redis.ErrNil{
			//表示在hash中没有这个id
			err = ErrorUserNotExits
		}
		return
	}
	user = &User{}
	//这里需要将res反序列化成User实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("json.UnMarshal err = ",err)
		return
	}
	fmt.Printf("after unmarshal user[userid] = %s user[userid] = %s user[userid] =%s \n",user.UserId,user.UserPwd,user.UserName)
	return
}

//完成对用户的检验 如果用户的id 和 pwd 都正确 就返回一个user实例 如果id 和 pwd 有错误， 就返回的对应的错误

func (this *UserDao) Login(userId string ,userPwd string) (user *User,err error)  {
	//先从链接池里取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.GetUserById(conn,userId)
	if err != nil {
		return
	}

	//这时证明这个用户是获取到了
	//但是无法证明密码是否正确
	if user.UserPwd != userPwd{
		err = ErrorUserPassword
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error)  {
	//先从链接池里取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.GetUserById(conn,user.UserId)
	if err == nil {
		err = ErrorUSerExits
		return
	}
	//这是说明用户没有重复 就可以入库
	data,err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_,err = conn.Do("hset","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("save register user err = ",err)
		return
	}
	return
}