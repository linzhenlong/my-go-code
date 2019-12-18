package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 声明一个全局的变量
// 在服务器启动时就初始化他.
var  (
	MyUserDao *userDao
)


// 定义一个userDao的结构体，完成对user结构体的各种操作
type userDao struct {
	pool *redis.Pool
	user User
}

// 使用工程模式创建一个userDao的实例
func NewUserDao(pool *redis.Pool)*userDao{
	return &userDao{
		pool:pool,
	}
}


func (userDao *userDao)getUserById(conn redis.Conn, id int)(user *User,err error)  {

	res, err := redis.String(conn.Do("hget","go:users", id))
	if err != nil{
		// 如果返回这个错误，在redis 里不存在
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}
	user = &User{}
	// 把res 反序列化为user实例
	err = json.Unmarshal([]byte(res), user)

	return
}



// 登录校验
func (userDao *userDao)Login(userId int, userPwd string) (user *User ,err error)  {

	// 1.先从userDao的redis连接池中取一个连接
	conn := userDao.pool.Get()
	defer conn.Close()
	user , err = userDao.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册
func (userDao *userDao)Register(user User)(b bool, err error)  {
	fmt.Println("(userDao *userDao)Register",user)
	if user.UserId == 0 || len(user.UserPwd) == 0 {
		err = ERROR_NO_USER_ID_PWD
		return
	}
	// 判断用户是否存在
	isExists, err := userDao.existsById(user)
	if err != nil {
		err = ERROR_OTHER
		return
	}
	if isExists {
		err = ERROR_USER_EXISTS
		return
	}

	// 添加用户

	return userDao.addUser(user)

}

// 添加用户
func (userDao *userDao)addUser(user User)(bool, error) {
	fmt.Println("(userDao *userDao)addUser",user)
	userId := user.UserId
	if  userId == 0 {
		err := ERROR_NO_USER_ID_PWD
		return false, err
	}
	if len(user.UserPwd) == 0 {
		err := ERROR_NO_USER_ID_PWD
		return false, err
	}
	if len(user.UserName) == 0 {
		err := ERROR_NO_USERNAME
		return false, err
	}

	// redis 链接
	conn := userDao.pool.Get()
	defer conn.Close()


	userByte, err := json.Marshal(user)
	if err != nil {
		err = ERROR_OTHER
		return false,err
	}

	// 写redis
	status, err := redis.Bool(conn.Do("hset","go:users", userId,string(userByte)))
	if err != nil {
		err = ERROR_WRITE_REDIS
		return status, err
	}
	return status, err
}

// 判断用户id是否存在
func (userDao *userDao)existsById(user User) (bool,error){
	fmt.Println("(userDao *userDao)existsById",user)
	userId :=user.UserId
	// 如果用户id=0 那么就认为存在改用户不允许注册
	if userId == 0 {
		return true, nil
	}
	conn := userDao.pool.Get()
	defer conn.Close()

	// 去redis 里判断
	return redis.Bool(conn.Do("HEXISTS", "go:users", userId))
}