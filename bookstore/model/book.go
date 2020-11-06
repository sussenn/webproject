package model

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
	//销量
	Sales int
	//库存
	Stock   int
	ImgPath string
}
