package datamodels

import ()

// Product 商品模型结构体.
type Product struct {
	ID           int64  `json:"id" form:"id" sql:"id" ` // form为自定义标签
	ProductName  string `json:"product_name" form:"product_name" sql:"product_name"`
	ProductNum   int64  `json:"product_num" form:"product_num" sql:"product_num"`
	ProductImage string `json:"product_image" form:"product_image" sql:"product_image"`
	ProductURL   string `json:"product_url" form:"product_url" sql:"product_url"`
}
