package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/baskets/", allBaskets).Methods("GET")
	myRouter.HandleFunc("/baskets/", newBasket).Methods("POST")
	myRouter.HandleFunc("/baskets/{basketid}/", getBasket).Methods("GET")
	myRouter.HandleFunc("/baskets/{basketid}/", deleteBasket).Methods("DELETE")
	myRouter.HandleFunc("/baskets/{basketid}/products/", getProductsInBasket).Methods("GET")
	myRouter.HandleFunc("/baskets/{basketid}/products/{productid}/", addItemInBasket).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func allBaskets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(baskets)
}

func newBasket(w http.ResponseWriter, r *http.Request) {
	b := NewBasket()
	baskets = append(baskets, b)
	json.NewEncoder(w).Encode(baskets)
}

func addItemInBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	basketId, err1 := strconv.Atoi(vars["basketid"])
	productId, err2 := strconv.Atoi(vars["productid"])
	b := &baskets[basketId]
	if err1 != nil {
		// handle error
		fmt.Println(err1)
		os.Exit(2)
	}
	if err2 != nil {
		// handle error
		fmt.Println(err2)
		os.Exit(2)
	}
	item := GetProductById(productId)
	AddProduct(b, item)
	json.NewEncoder(w).Encode(b)
}

func getBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	basketId, err1 := strconv.Atoi(vars["basketid"])
	b := baskets[basketId]
	if err1 != nil {
		// handle error
		fmt.Println(err1)
		os.Exit(2)
	}
	json.NewEncoder(w).Encode(b)
}

func getProductsInBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	basketId, err1 := strconv.Atoi(vars["basketid"])
	i := baskets[basketId].Items
	if err1 != nil {
		// handle error
		fmt.Println(err1)
		os.Exit(2)
	}
	json.NewEncoder(w).Encode(i)
}

func deleteBasket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	basketId, err1 := strconv.Atoi(vars["basketid"])
	if err1 != nil {
		// handle error
		fmt.Println(err1)
		os.Exit(2)
	}
	RemoveBasket(basketId)
	json.NewEncoder(w).Encode(baskets)
}

func main() {
	fmt.Println("Checkout Server Started")
	handleRequests()
}
