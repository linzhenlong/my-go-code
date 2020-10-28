package repositories

import (
	"database/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"log"
)

// 1.定义接口

// IProduct 接口.
type IProduct interface {
	Conn() error                                    // 数据库连接
	Insert(*datamodels.Product) (int64, error)      // 添加商品
	Delete(int64) bool                              // 删除商品
	Update(*datamodels.Product) error               // 更新
	SelectByKey(int64) (*datamodels.Product, error) // 查询
	SelectAll() ([]*datamodels.Product, error)      // 获取所有商品
	SelectAllByParams(map[string]interface{}) ([]*datamodels.Product, error)
	GetTotal(map[string]interface{}) int64
}

// 2.实现接口

// ProductManager 商品manager.
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
	gormCoon  *gorm.DB
}

// NewProductManager 构造方法 这里的方法返回值是IProduct return 里面返回的是ProductManager
// 也就是说，如果ProductManager实现了IProduct里的方法就OK了，起到自检的效果.
func NewProductManager(table string, db *sql.DB, gormDB *gorm.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
		gormCoon:  gormDB,
	}
}

// Conn 实现数据库连接的方法
func (p *ProductManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.gormCoon == nil {
		gorm, err := common.NewGorm()
		if err != nil {
			return err
		}
		p.gormCoon = gorm
	}
	if p.table == "" {
		p.table = "product"
	}
	return
}

// Insert 实现插入方法
func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	/* sql := "insert into " + p.table + "(product_name,product_num,product_image,product_url) values(?,?,?,?)"

	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductURL)
	if err != nil {
		return
	}
	return result.LastInsertId() */
	p.gormCoon.NewRecord(product)
	err = p.gormCoon.Create(&product).Error
	if err != nil {
		return
	}

	return product.ID, err
}

// Delete 实现删除方法.
func (p *ProductManager) Delete(id int64) bool {
	// 判断数据库连接是否存在
	if err := p.Conn(); err != nil {
		// 数据库连不上就直接删除失败吧.
		return false
	}
	if id == 0 {
		return false
	}
	/* sql := "Delete FORM" + p.table + " WHERE id=? limit 1"
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return false
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return false
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return affectedRows > 0 */
	product := datamodels.Product{
		ID: id,
	}
	err := p.gormCoon.Delete(&product).Error
	if err != nil {
		return false
	}
	return true
}

// Update 商品更新方法.
func (p *ProductManager) Update(product *datamodels.Product) error {
	// 判断数据库连接是否存在
	if err := p.Conn(); err != nil {
		return err
	}
	if product.ID == 0 {
		return errors.New("productID 不能为空")
	}
	pro := datamodels.Product{
		ID: product.ID,
	}

	log.Printf("%#v", product)
	productMap := make(map[string]interface{})
	productMap["product_name"] = product.ProductName
	productMap["product_num"] = product.ProductName
	productMap["product_image"] = product.ProductImage
	productMap["product_url"] = product.ProductURL
	err := p.gormCoon.Debug().Model(&pro).Updates(productMap).Error
	return err
	/* sql := "update " + p.table + " set product_name=?,product_num=?,product_image=?,product_url=? where id=? limit 1"
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductURL, product.ID)

	if err != nil {
		return err
	}
	return nil */
}

// SelectByKey 通过productID 获取product信息.
func (p *ProductManager) SelectByKey(productID int64) (product *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}

	pro := &datamodels.Product{}
	err = p.gormCoon.First(&pro, productID).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return pro, err
	}
	return pro, nil
}

// SelectAll 获取所有的商品列表.
func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	err = p.gormCoon.Find(&products).Error
	return
}

// GetTotal 获取总数.
func (p *ProductManager) GetTotal(params map[string]interface{}) int64 {
	db := p.gormCoon.Debug().Model(&datamodels.Product{})
	productName, ok := params["product_name"].(string)
	if ok {
		db = db.Where("product_name like %?%", productName)
	}
	var total int64
	db.Count(&total)
	return total
}

// SelectAllByParams .
func (p *ProductManager) SelectAllByParams(params map[string]interface{}) (products []*datamodels.Product, err error) {
	db := p.gormCoon.Debug().Model(&datamodels.Product{})
	productName, ok := params["product_name"].(string)
	if ok {
		db = db.Where("product_name like %?%", productName)
	}
	limit, ok := params["limit"].(int)
	if ok {
		db = db.Limit(limit)
	}
	offset, ok := params["offset"].(int)
	if ok {
		db = db.Offset(offset)
	}
	err = db.Find(&products).Error
	return
}
