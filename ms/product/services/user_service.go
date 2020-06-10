package services

import (
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
)

// IUserService .
type IUserService interface {
	AddUser(*datamodels.User) (int64, error)
	GetUserByName(string) (*datamodels.User, error)
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
	return u.repository.Insert(user)
}

// GetUserByName .
func (u *UserService) GetUserByName(userName string) (user *datamodels.User, err error) {
	return u.repository.SelectByName(userName)
}
