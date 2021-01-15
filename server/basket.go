package main

type Basket struct {
	Total float64
	Items []Product
}

func NewBasket() Basket {
	b := Basket{Total: 0.00, Items: []Product{}}
	return b
}

func RemoveBasket(index int) {
	baskets = append(baskets[:index], baskets[index+1:]...)
}
