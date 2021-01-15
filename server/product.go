package main

type Product struct {
	ProductName string
	ProductCost float64
}

var pen = Product{"Pen", 5.00}
var tshirt = Product{"Tshirt", 20.00}
var mug = Product{"Mug", 7.50}

func GetProductById(id int) Product {
	switch id {
	case 1:
		return pen
	case 2:
		return tshirt
	default:
		return mug
	}
}
