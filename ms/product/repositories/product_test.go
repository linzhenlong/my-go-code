package repositories

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
)

var (
	productID        int64
	mysql            *sql.DB
	productInterface IProduct
)

func TestMain(m *testing.M) {

	mysql, _ = common.NewMysqlConn()
	productInterface = NewProductManager("product", mysql)
	clearData()
	m.Run()
	clearData()
}

func clearData() {
	mysql.Exec("truncate product")
}

func TestFlow(t *testing.T) {
	t.Run("insert", TestInsert)
	t.Run("selectByPK", TestSelectBYkey)
}

func TestInsert(t *testing.T) {
	IProduct := NewProductManager("product", mysql)

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
	t.SkipNow()
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
	proList, _ := productInterface.SelectAll()
	contents, _ := json.Marshal(proList)
	t.Logf("%s", string(contents))
}
