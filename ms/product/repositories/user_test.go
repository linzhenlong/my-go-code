package repositories

import (
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUserInsert(t *testing.T) {
	rand.Seed(time.Now().Unix())
	r := rand.Int63n(100)
	user := datamodels.User{
		NickName: "张三" + strconv.Itoa(int(r)),
		Password: common.GenMd5("123456"),
		UserName: "zhangsan" + strconv.Itoa(int(r)),
	}
	gormdb, _ := common.NewGorm()
	userMgr := NewUserRepository("user", gormdb)
	userID, err := userMgr.Insert(&user)
	if err != nil {
		t.Fatalf("inser err :%v", err)
	}
	t.Logf("userid ==>%d", userID)
}

func TestUserSelectByName(t *testing.T) {
	userName1 := "zhangsan1"
	gormdb, _ := common.NewGorm()
	userMgr := NewUserRepository("user", gormdb)
	user1, err := userMgr.SelectByName(userName1)
	if err != nil {
		t.Errorf("user1 err:%v", err)
	}
	t.Logf("user1 ===>%#v", user1)
	userName2 := "zhangsan64"
	user2, err := userMgr.SelectByName(userName2)
	if err != nil {
		t.Errorf("user2 err:%v", err)
	}
	t.Logf("user2 ===>%#v", user2)

}
