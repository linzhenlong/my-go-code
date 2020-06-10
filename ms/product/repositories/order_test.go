package repositories

import (
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"testing"
)

var (
	orderID int64
)

func TestOrderInsert(t *testing.T) {
	t.SkipNow()
	myGorm, _ := common.NewGorm()
	orderManager := NewOrderManager("order", myGorm)
	order := datamodels.Order{
		UserID:      11,
		ProductID:   9,
		OrderStatus: datamodels.OrderWait,
	}
	id, err := orderManager.Insert(&order)
	orderID = id
	if err != nil {
		t.Fatalf("insert error:%v", err)
	}

	t.Logf("insert succ orderID:%d", orderID)
}

func TestOrderDel(t *testing.T) {
	t.SkipNow()
	myGorm, _ := common.NewGorm()
	orderManager := NewOrderManager("order", myGorm)
	if orderManager.Delete(orderID) {
		t.Logf("order delete succ orderid:%d", orderID)
	} else {
		t.Fatal("order delete err ")
	}
}

func TestOrderSelectByKey(t *testing.T) {
	t.SkipNow()
	myGorm, _ := common.NewGorm()
	orderManager := NewOrderManager("order", myGorm)
	order, err := orderManager.SelectByKey(orderID)
	if err != nil {
		t.Fatalf("TestSelectByKey err:%v", err)
	}
	t.Logf("order info :%#v", order)
}
func TestOrderUpdate(t *testing.T) {
	t.SkipNow()
	myGorm, _ := common.NewGorm()
	orderManager := NewOrderManager("order", myGorm)
	t.Logf("orderid===>%d", orderID)
	order, err := orderManager.SelectByKey(orderID)
	if err != nil {
		t.Fatalf("TestOrderUpdate TestOrderUpdate err:%v", err)
	}
	t.Logf("update before %v", order)
	order.OrderStatus = datamodels.OrderSuccess
	err = orderManager.Update(order)
	if err != nil {
		t.Fatalf("update err:%v", err)
	}
	order2, _ := orderManager.SelectByKey(order.ID)
	if order2.OrderStatus == datamodels.OrderSuccess {
		t.Log("OrderUpdate success")
		t.Logf("update after %#v", order)
	} else {
		t.Fatalf("OrderUpdate err order:%#v", order2)
	}
}

func TestOrderSelectAll(t *testing.T) {
	t.SkipNow()
	myGorm, _ := common.NewGorm()
	orderManager := NewOrderManager("order", myGorm)
	params := make(map[string]interface{})
	params["orderStatus"] = int64(1)
	params["limit"] = int(10)

	orderStatus, ok := params["orderStatus"].(int64)
	t.Logf("orderStatus:%v--ok:%#v\n", orderStatus, ok)
	res, err := orderManager.SelectAll(params)
	if err != nil {
		t.Fatalf("TestOrderSelectAll err:%v", err)
	}

	t.Logf("all--->%v", res)
	for key, val := range res {
		t.Logf("key:%v--value:%#v\n", key, val)
	}
}

func TestOrderFlow(t *testing.T) {
	t.SkipNow()
	t.Run("add", TestOrderInsert)
	t.Run("select", TestOrderSelectByKey)
	t.Run("update", TestOrderUpdate)
	t.Run("del", TestOrderDel)
	t.Run("all", TestOrderSelectAll)
}
