package model

import "errors"

// 根据业务逻辑需要，自定义错误
var (
	ERROR_USER_NOT_EXISTS = errors.New("用户不存在。。。")
	ERROR_USER_EXISTS     = errors.New("用户已存在。。。")
	ERROR_USER_PWD     = errors.New("用户名或密码错误。。。")
	ERROR_NO_USER_ID_PWD = errors.New("用户id或密码不能为空")
	ERROR_NO_USERNAME = errors.New("用户名不能为空")
	ERROR_OTHER = errors.New("未知错误")
	ERROR_WRITE_REDIS = errors.New("redis 写入失败")


)
