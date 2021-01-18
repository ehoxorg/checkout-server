package main

import "fmt"

var baskets []Basket

// NOTE: CONSIDERATION TO USE A MAP DATASTRUCTURE AS DECLARED IN THE FOLLOWING VARIABLES numberOfPensInBasket, numberOfTshirtsInBasket
// for constant log(1) complexity while trying to calculate whether the basket is a candidate for an offer (trying to avoid log(n) complexity)
// although with the added expense that the remove basket operator would mean traversing through the keys to decrement the basket indeces
// var numberOfPensInBasket = make(map[int]int)
// var numberOfTshirtsInBasket = make(map[int]int)

func AddProduct(basket *Basket, product Product) {
	// NOTE: Research shows that gorilla/mux starts a new goroutine for each request. https://stackoverflow.com/questions/49975616/golang-rest-api-concurrency
	// Therefore it should be safe to perform concurrent requests on the server.
	// However, the append operation is apparently not thread safe. https://medium.com/@cep21/gos-append-is-not-always-thread-safe-a3034db7975
	// For the time being the code keeps using the append operation in its simple form.
	basket.Items = append(basket.Items, product)
	if product.ProductName == pen.ProductName {
		addPenToBasket(basket)
	} else if product.ProductName == tshirt.ProductName {
		addTshirtToBasket(basket)
	} else {
		basket.Total = basket.Total + product.ProductCost
	}
	fmt.Println(basket.Total)
}

func addPenToBasket(basket *Basket) {
	var totalPens int
	for i := 0; i < len(basket.Items); i++ {
		if basket.Items[i].ProductName == pen.ProductName {
			totalPens++
		}
	}
	//checking to see if there will be an even number of pens in the basket with the one just added
	//if so the total remains as it is
	//otherwise the total is added the cost of a pen
	isPenPair := totalPens%2 == 0
	if !isPenPair {
		basket.Total = basket.Total + pen.ProductCost
	}
}

func addTshirtToBasket(basket *Basket) {
	var totalTshirts int
	for i := 0; i < len(basket.Items); i++ {
		if basket.Items[i].ProductName == tshirt.ProductName {
			totalTshirts++
		}
	}
	//totalTshirts == 3 | deducting for the previous 2 tshirts before adding the reduced new tshirt cost
	//totalTshirts > 3  | adding the reduced cost of a tshirt
	//otherwise 		| adding full cost
	if totalTshirts == 3 {
		costToBeDeducted := tshirt.ProductCost * 0.25 * 2
		deductedTotal := basket.Total - costToBeDeducted
		basket.Total = deductedTotal + tshirt.ProductCost*0.75
	} else if totalTshirts > 3 {
		basket.Total = basket.Total + tshirt.ProductCost*0.75
	} else {
		basket.Total = basket.Total + tshirt.ProductCost
	}
}
