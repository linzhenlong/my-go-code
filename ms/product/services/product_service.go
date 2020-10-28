package services

import (
	"github.com/linzhenlong/my-go-code/ms/product/datamodels"
	"github.com/linzhenlong/my-go-code/ms/product/repositories"
)

// IProductService product service 接口.
type IProductService interface {
	GetProductByID(int64) (*datamodels.Product, error)
	GetAllProduct() ([]*datamodels.Product, error)
	DeleteProductByID(int64) bool
	InsertProduct(*datamodels.Product) (int64, error)
	UpdateProduct(*datamodels.Product) error
	GetTotal(map[string]interface{}) int64
	SelectAllByParams(map[string]interface{}) ([]*datamodels.Product, error)
}

// ProductService 结构体.
type ProductService struct {
	productRepository repositories.IProduct
}

// NewProductService 构造方法.
func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{
		productRepository: repository,
	}
}

// GetProductByID 获取商品详情.
func (p *ProductService) GetProductByID(productID int64) (productInfo *datamodels.Product, err error) {
	productInfo, err = p.productRepository.SelectByKey(productID)
	return
}

// GetAllProduct 获取商品列表.
func (p *ProductService) GetAllProduct() (productList []*datamodels.Product, err error) {
	return p.productRepository.SelectAll()
}

// DeleteProductByID 删除商品.
func (p *ProductService) DeleteProductByID(productID int64) bool {
	return p.productRepository.Delete(productID)
}

// InsertProduct 插入商品.
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

// UpdateProduct 插入商品.
func (p *ProductService) UpdateProduct(product *datamodels.Product) error {
	return p.productRepository.Update(product)
}

// GetTotal .
func (p *ProductService) GetTotal(params map[string]interface{}) int64 {
	return p.productRepository.GetTotal(params)
}

// SelectAllByParams .
func (p *ProductService) SelectAllByParams(params map[string]interface{}) ([]*datamodels.Product, error) {
	return p.productRepository.SelectAllByParams(params)
}
