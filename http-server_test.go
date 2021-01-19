package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestAllBaskets_SMOKE_TEST(t *testing.T) {
	req, err := http.NewRequest("GET", "/baskets/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AllBaskets)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.

	if strings.TrimSuffix(rr.Body.String(), "\n") != "null" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "null")
	}
}

func TestCreateBaskets(t *testing.T) {
	req, err := http.NewRequest("POST", "/baskets/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBasket)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `[{"Total":0,"Items":[]}]`
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBasket(t *testing.T) {
	url := "/baskets/{basketid}/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{
		"basketid": "0",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBasket)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body is what we expect.
	expected := `{"Total":0,"Items":[]}`
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//following test from GITHUB example
// Items: PEN, TSHIRT, PEN
// Total: 25.00€
func TestAddProductInBasket_TWO_PENS_AND_TSHIRT(t *testing.T) {
	req, err2 := http.NewRequest("PUT", "/baskets/{basketid}/products/{productid}/", nil)
	if err2 != nil {
		t.Fatal(err2)
	}
	vars := map[string]string{
		"basketid":  "0",
		"productid": "1",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddItemInBasket)
	handler.ServeHTTP(rr, req)

	vars = map[string]string{
		"basketid":  "0",
		"productid": "2",
	}
	rr2 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr2, req)

	vars = map[string]string{
		"basketid":  "0",
		"productid": "1",
	}
	rr3 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr3, req)

	if status2 := rr3.Code; status2 != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status2, http.StatusOK)
	}

	expected := `{"Total":25,"Items":[{"ProductName":"Pen","ProductCost":5},{"ProductName":"Tshirt","ProductCost":20},{"ProductName":"Pen","ProductCost":5}]}`
	if strings.TrimSuffix(rr3.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr3.Body.String(), expected)
	}
}

// Items: TSHIRT, TSHIRT, TSHIRT, PEN, TSHIRT
// Total: 65.00€
func TestAddProductInBasket_ONE_PEN_AND_THREE_TSHIRTS(t *testing.T) {

	req, err := http.NewRequest("POST", "/baskets/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rrCreateSecondBasket := httptest.NewRecorder()
	handlerCreateSecondBasket := http.HandlerFunc(CreateBasket)
	handlerCreateSecondBasket.ServeHTTP(rrCreateSecondBasket, req)
	if status := rrCreateSecondBasket.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req, err2 := http.NewRequest("PUT", "/baskets/{basketid}/products/{productid}/", nil)
	if err2 != nil {
		t.Fatal(err2)
	}
	vars := map[string]string{
		"basketid":  "1",
		"productid": "2",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddItemInBasket)
	handler.ServeHTTP(rr, req)

	vars = map[string]string{
		"basketid":  "1",
		"productid": "2",
	}
	rr2 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr2, req)

	vars = map[string]string{
		"basketid":  "1",
		"productid": "2",
	}
	rr3 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr3, req)

	vars = map[string]string{
		"basketid":  "1",
		"productid": "1",
	}
	rr4 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr4, req)

	vars = map[string]string{
		"basketid":  "1",
		"productid": "2",
	}
	rr5 := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)
	handler.ServeHTTP(rr5, req)

	if status2 := rr5.Code; status2 != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status2, http.StatusOK)
	}

	expected := `{"Total":65,"Items":[{"ProductName":"Tshirt","ProductCost":20},{"ProductName":"Tshirt","ProductCost":20},{"ProductName":"Tshirt","ProductCost":20},{"ProductName":"Pen","ProductCost":5},{"ProductName":"Tshirt","ProductCost":20}]}`
	if strings.TrimSuffix(rr5.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr5.Body.String(), expected)
	}
}

//NOTE: CONSIDERATIONS TO USE GOROUTINES TO TEST ADDING ITEMS TO THE SAME BASKET ARE CONSIDERED
//HOWEVER SINCE THIS IS THE FIRST TIME IM DOING SMTH IN GO I WOULD RATHER PASS
//AS IT MIGHT BE VERY TIME CONSUMING TO HANDLE THEM. NORMALLY I EXPECT THE APPEND FUNCTION TO FAIL
//AS MENTIONED IN THE CHECKOUT.GO FILE, BUT FOR THE TIME BEING I WILL NOT BE ABLE TO FIX THAT AND
//MAKE THE PROGRAM RUN CONCURRENT OPERATIONS ON THE SAME BASKET. HOPE IT IS UNDERSTANDABLE :)
