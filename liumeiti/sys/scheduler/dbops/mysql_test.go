package dbops

import (
	"github.com/linzhenlong/my-go-code/liumeiti/sys/scheduler/utils"
	"testing"
)


var (
	videoUUID string
)

func TestMain(m *testing.M) {
	ClearData()
	m.Run()
	ClearData()
}

func ClearData() {
	dbConn.Exec("truncate video_del_rec")
}

func TestFlow(t *testing.T) {
	t.Run("add",TestAdd)
	t.Run("get",TestGet)
	t.Run("del",TestDel)
}

func TestAdd(t *testing.T) {
	videoUUID,_ = utils.NewUUID()
	err := AddVideoDeletionRecord(videoUUID)
	if err != nil {
		t.Fatalf("add error")
	}
	t.Log("TestAdd pass")
}
func TestGet(t *testing.T) {
	res, err := ReadVideoDeletionRecord(3)
	if err !=nil {
		t.Fatalf("TestGet error:%s", err.Error())
	}
	if len(res) == 0 {
		t.Fatalf("TestGet error:%v",res)
	}
	var flag bool
	for _, v := range res {
		if v == videoUUID {
			flag = true
			break
		}
	}
	if flag {
		t.Log("TestGet pass")
	} else {
		t.Fatalf("TestGet error")
	}
}

func TestDel(t *testing.T) {
	err := DelVideoDeletionRecord(videoUUID)
	if err != nil {
		t.Fatalf("TestDel error:%s", err.Error())
	}
	t.Logf("TestDel pass")
}