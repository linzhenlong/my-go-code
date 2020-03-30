package repositories

import (
	"database/sql"
	"errors"
	"log"

	"github.com/linzhenlong/my-go-code/ms/product/common"
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
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
}

// 2.实现接口

// ProductManager 商品manager.
type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

// NewProductManager 构造方法 这里的方法返回值是IProduct return 里面返回的是ProductManager
// 也就是说，如果ProductManager实现了IProduct里的方法就OK了，起到自检的效果.
func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
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
	sql := "insert into " + p.table + "(product_name,product_num,product_image,product_url) values(?,?,?,?)"

	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return
	}
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductURL)
	if err != nil {
		return
	}
	return result.LastInsertId()
}

// Delete 实现删除方法.
func (p *ProductManager) Delete(id int64) bool {
	// 判断数据库连接是否存在
	if err := p.Conn(); err != nil {
		// 数据库连不上就直接删除失败吧.
		return false
	}
	sql := "Delete FORM" + p.table + " WHERE id=? limit 1"
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
	return affectedRows > 0
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
	log.Printf("%#v", product)
	
	sql := "update " + p.table + " set product_name=?,product_num=?,product_image=?,product_url=? where id=? limit 1"
	stmt, err := p.mysqlConn.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductURL, product.ID)

	if err != nil {
		return err
	}
	return nil
}

// SelectByKey 通过productID 获取product信息.
func (p *ProductManager) SelectByKey(productID int64) (product *datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	sqlTemplate := "select id,product_name,product_num,product_image,product_url from " + p.table + " where id=?"
	stmt, err := p.mysqlConn.Prepare(sqlTemplate)
	defer stmt.Close()
	if err != nil {
		return
	}
	log.Printf("lzltest")
	var (
		productName  string
		productNum   int64
		productImage string
		productURL   string
		ID           int64
	)
	err = stmt.QueryRow(productID).Scan(&ID, &productName, &productNum, &productImage, &productURL)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	product = &datamodels.Product{
		ID:           ID,
		ProductName:  productName,
		ProductNum:   productNum,
		ProductImage: productImage,
		ProductURL:   productURL,
	}
	return
}

// SelectAll 获取所有的商品列表.
func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err = p.Conn(); err != nil {
		return
	}
	stmt, err := p.mysqlConn.Prepare("select id,product_name,product_num,product_image,product_url from " + p.table)
	if err != nil {
		return
	}
	res, err := stmt.Query()
	if err != nil && err != sql.ErrNoRows {
		return
	}
	for res.Next() {
		var (
			productName  string
			productNum   int64
			productImage string
			productURL   string
			ID           int64
		)
		err = res.Scan(&ID, &productName, &productNum, &productImage, &productURL)
		if err != nil && err != sql.ErrNoRows {
			continue
		}
		product := &datamodels.Product{
			ID:           ID,
			ProductName:  productName,
			ProductNum:   productNum,
			ProductImage: productImage,
			ProductURL:   productURL,
		}
		products = append(products, product)
	}
	return
}
