package services

import (
	"errors"

	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
	"golang.org/x/crypto/bcrypt"
)

// IUserService .
type IUserService interface {
	AddUser(*datamodels.User) (int64, error)
	GetUserByName(string) (*datamodels.User, error)
	IsPwdSuccess(userName string, password string) (user *datamodels.User, isOk bool)
}

// UserService .
type UserService struct {
	repository repositories.IUserRepository
}

// NewUserService .
func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}

// AddUser 添加用户.
func (u *UserService) AddUser(user *datamodels.User) (userID int64, err error) {

	pwdByte, errPwd := GeneratePassword(user.Password)
	if errPwd != nil {
		return userID, errPwd
	}
	user.Password = string(pwdByte)
	return u.repository.Insert(user)
}

// GetUserByName .
func (u *UserService) GetUserByName(userName string) (user *datamodels.User, err error) {
	return u.repository.SelectByName(userName)
}

// IsPwdSuccess .
func (u *UserService) IsPwdSuccess(userName string, password string) (user *datamodels.User, ok bool) {
	user, err := u.repository.SelectByName(userName)
	if err != nil {
		return
	}
	ok, _ = ValidatePassword(password, user.Password)
	if !ok {
		return &datamodels.User{}, false
	}
	return
}

//GeneratePassword 生成密码.
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword 比对密码.
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}
