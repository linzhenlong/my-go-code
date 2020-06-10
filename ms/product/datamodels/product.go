package datamodels

// Product 商品模型结构体.
type Product struct {
	ID           int64  `json:"id" form:"id" gorm:"id" ` // form为自定义标签
	ProductName  string `json:"product_name" form:"product_name" gorm:"product_name"`
	ProductNum   int64  `json:"product_num" form:"product_num" gorm:"product_num"`
	ProductImage string `json:"product_image" form:"product_image" gorm:"product_image"`
	ProductURL   string `json:"product_url" form:"product_url" gorm:"product_url"`
}

// TableName 覆盖一下gorm的默认表名.
func (p Product) TableName() string {
	return "product"
}
