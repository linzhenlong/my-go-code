package repositories

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
)

var (
	productID        int64
	mysql            *sql.DB
	productInterface IProduct
	gormDB           *gorm.DB
)

func TestMain(m *testing.M) {

	mysql, _ = common.NewMysqlConn()
	productInterface = NewProductManager("product", mysql, gormDB)
	//clearData()
	m.Run()
	//clearData()
}

func clearData() {
	mysql.Exec("truncate product")
}

func TestFlow(t *testing.T) {
	t.SkipNow()
	t.Run("insert", TestInsert)
	t.Run("selectByPK", TestSelectBYkey)
}

func TestInsert(t *testing.T) {
	//t.SkipNow()
	IProduct := NewProductManager("product", mysql, gormDB)

	testProductInfo := &datamodels.Product{
		ProductNum:   199,
		ProductName:  "this is a test",
		ProductImage: "http://www.baidu.com",
		ProductURL:   "http://www.baidu.com",
	}
	id, err := IProduct.Insert(testProductInfo)
	productID = id
	if err != nil {
		t.Fatalf("添加商品失败er	r:%s", err.Error())
	}
	t.Logf("Insert product success proID:%d", productID)
}
func TestSelectBYkey(t *testing.T) {
	//t.SkipNow()
	pro, err := productInterface.SelectByKey(productID)
	t.Log("productID===>", productID)
	if err != nil {
		t.Fatalf("SelectByKey err:%s", err.Error())
	}
	if pro.ProductName != "this is a test" {
		t.Fatalf("SelectByKey pro.ProductName Error")
	}
	t.Logf("SelectByKey success =>%v", pro)

	contents, _ := json.Marshal(pro)
	t.Logf("%s", string(contents))
}

func TestSelectAll(t *testing.T) {
	t.SkipNow()
	proList, _ := productInterface.SelectAll()
	t.Logf("proList:%#v", proList)
	contents, _ := json.Marshal(proList)
	t.Logf("TestSelectAll:%s", string(contents))
}

func TestDelete(t *testing.T) {
	//t.SkipNow()
	res := productInterface.Delete(1)
	t.Logf("TestDelete res:%v", res)
}

func TestUpdate(t *testing.T) {
	//t.SkipNow()
	testProductInfo := &datamodels.Product{
		ID:           17,
		ProductNum:   299,
		ProductName:  "this is a test 18",
		ProductImage: " ",
		ProductURL:   "http://www.baidu.com",
	}
	err := productInterface.Update(testProductInfo)
	t.Logf("TestUpdate err:%v", err)
}
