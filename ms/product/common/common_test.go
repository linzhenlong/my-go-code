package common

import (
	"database/sql"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

type product struct {
	ID           int64  `json:"id" ms:"id" sql:"id" ` // ms为自定义标签
	ProductName  string `json:"product_name" ms:"product_name" sql:"product_name"`
	ProductNum   int64  `json:"product_num" ms:"product_num" sql:"product_num"`
	ProductImage string `json:"product_image" ms:"product_image" sql:"product_image"`
	ProductURL   string `json:"product_url" ms:"product_url" sql:"product_url"`
	CreateDate   string `json:"create_date" ms:"create_date" sql:"create_date"`
}

var (
	mysql  *sql.DB
	myGorm *gorm.DB
)

func TestMain(m *testing.M) {
	mysql, _ = NewMysqlConn()
	myGorm, _ = NewGorm()
	m.Run()
}
func TestGetResultRow(t *testing.T) {
	t.SkipNow()
	mysql, _ := NewMysqlConn()
	row, _ := mysql.Query("select id,product_name,product_num,product_image,product_url  from product where id=1")

	res := GetResultRow(row)
	t.Log(res)
}

func TestGetResultRows(t *testing.T) {
	t.SkipNow()
	mysql, _ := NewMysqlConn()
	rows, _ := mysql.Query("select * from product")
	res := GetResultRows(rows)
	t.Log(res)
}

func TestDataToStruct(t *testing.T) {
	t.SkipNow()
	row, _ := mysql.Query("select * from product where id=1")
	res := GetResultRow(row)
	productRes := &product{}
	DataToStructByTagSQL(res, productRes)
	t.Log(productRes)
}

func TestGorm(t *testing.T) {
	has := myGorm.HasTable("product")
	if has {
		t.Log("has")
	} else {
		t.Log("no has")
	}
}

type Goods struct {
	GoodsModel
	GoodsName  string `json:"goods_name" gorm:"goods_name"`
	GoodsNum   int64  `json:"goods_num" gorm:"goods_num"`
	GoodsImage string `json:"goods_image" gorm:"goods_image"`
	GoodsURL   string `json:"goods_url" gorm:"goods_url"`
}

func (g Goods) TableName() string {
	return "goods"
}

type GoodsModel struct {
	ID        int64 `json:"id" gorm:"id primary_key AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func TestGormCreate(t *testing.T) {
	t.SkipNow()
	/* goods := Good{
		GoodsName:  "this is test goods",
		GoodsNum:   199,
		GoodsImage: "http://img0.pconline.com.cn/pconline/2003/26/13310924_61b2bc9b0f61ce9edaf308ecf9138b3_thumb.png",
		GoodsURL:   "http://www.baidu.com/",
	} */
	//myGorm.DropTable("goods")
	err := myGorm.CreateTable(&Goods{}).GetErrors()
	if len(err) > 0 {
		t.Fatalf("%v", err)
	}
	//myGorm.DropTable("goods")
	//myGorm.DropTable(&Good{})
}

func TestNewRecord(t *testing.T) {
	goods := Goods{
		GoodsName:  "this is test goods",
		GoodsNum:   199,
		GoodsImage: "http://img0.pconline.com.cn/pconline/2003/26/13310924_61b2bc9b0f61ce9edaf308ecf9138b3_thumb.png",
		GoodsURL:   "http://www.baidu.com/",
	}

	err := myGorm.Create(&goods).Error
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%d", goods.ID)
}
