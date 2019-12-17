package model

import (
	"encoding/json"
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