package main

import "fmt"

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Id int
}

type Category struct {
	Name string
	Id int
}

type Product struct {
	Goods
	Brand
	Category
}

func (b *Brand)SetBrand(name string,id int) {
	b.Name = name
	b.Id = id
}
func (product *Product)Set(name string ,price float64)  {
	product.Price = price
	product.Goods.Name = name
}

func (product *Product)ProductInfo() {
	fmt.Printf("商品名称是:%s\n",product.Goods.Name)
	fmt.Printf("商品价格是:%.2f\n",product.Goods.Price)
	fmt.Printf("品牌名称是:%s\n",product.Brand.Name)
	fmt.Printf("品牌id是:%d\n",product.Brand.Id)
	fmt.Printf("分类名称是:%s\n",product.Category.Name)
	fmt.Printf("分类id是:%d\n",product.Category.Id)
}

func main() {
	product  := Product{}
	product.Brand.SetBrand("松下",1)
	product.Category.Name = "家用电器"
	product.Category.Id = 1
	product.Set("电视机1",5999)
	product.ProductInfo()
}
