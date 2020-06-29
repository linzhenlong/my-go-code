package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
)

// IUserRepository 用户接口.
type IUserRepository interface {
	Conn() error
	SelectByName(string) (*datamodels.User, error)
	Insert(*datamodels.User) (int64, error)
	//DeleteByID(int64) bool
	//DeleteByName(string) bool
	//Update(*datamodels.User) error
}

// UserManagerRepository .
type UserManagerRepository struct {
	table  string
	myGorm *gorm.DB
}

// NewUserRepository 构造函数.
func NewUserRepository(table string, db *gorm.DB) IUserRepository {
	return &UserManagerRepository{
		table:  table,
		myGorm: db,
	}
}

// Conn 数据库
func (u *UserManagerRepository) Conn() error {
	if u.myGorm == nil {
		myGorm, err := common.NewGorm()
		if err != nil {
			return err
		}
		u.myGorm = myGorm
	}
	if u.table == "" {
		u.table = "ms_user"
	}
	return nil
}

// SelectByName .
func (u *UserManagerRepository) SelectByName(name string) (user *datamodels.User, err error) {
	if len(name) == 0 {
		return
	}
	if err = u.Conn(); err != nil {
		return
	}
	user = &datamodels.User{}
	err = u.myGorm.Debug().Where("user_name=?", name).First(user).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return user, err
}

// Insert 插入用户.
func (u *UserManagerRepository) Insert(user *datamodels.User) (id int64, err error) {
	if err = u.Conn(); err != nil {
		return
	}
	err = u.myGorm.Debug().Create(&user).Error
	return user.ID, err
}
