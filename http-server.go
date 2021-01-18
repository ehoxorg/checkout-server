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
	myRouter.HandleFunc("/baskets/", AllBaskets).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/baskets/", CreateBasket).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/baskets/{basketid}/", GetBasket).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/baskets/{basketid}/", DeleteBasket).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/baskets/{basketid}/products/", GetProductsInBasket).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/baskets/{basketid}/products/{productid}/", AddItemInBasket).Methods("PUT", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
}

func AllBaskets(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		json.NewEncoder(w).Encode(baskets)
	}
}

func CreateBasket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		b := NewBasket()
		baskets = append(baskets, b)
		json.NewEncoder(w).Encode(baskets)
	}
}

func AddItemInBasket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
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
}

func GetBasket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
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
}

func GetProductsInBasket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
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
}

func DeleteBasket(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
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
}

func main() {
	fmt.Println("Checkout Server Started")
	handleRequests()
}
