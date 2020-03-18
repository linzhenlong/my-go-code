package datamodels

import ()

// Product 商品模型结构体.
type Product struct {
	ID           int64  `json:"id" ms:"id" sql:"id" ` // ms为自定义标签
	ProductName  string `json:"product_name" ms:"product_name" sql:"product_name"`
	ProductNum   int64  `json:"product_num" ms:"product_num" sql:"product_num"`
	ProductImage string `json:"product_image" ms:"product_image" sql:"product_image"`
	ProductURL   string `json:"product_url" ms:"product_url" sql:"product_url"`
}
