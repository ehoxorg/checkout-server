package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
	// const (
	// 	sellersURL = "http://myhost.com/v1/sellers"
	// )

	// q := url.Values{}
	// q.Add("id", "1")
	req, err := http.NewRequest("GET", "/baskets/0/", nil)
	if err != nil {
		t.Fatal(err)
	}
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

func TestAddProductInBasket(t *testing.T) {
	// req, err := http.NewRequest("POST", "/baskets/", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(CreateBasket)
	// handler.ServeHTTP(rr, req)
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	// req2, err2 := http.NewRequest("PUT", "/baskets/0/products/1/", nil)
	// if err2 != nil {
	// 	t.Fatal(err2)
	// }
	// rr2 := httptest.NewRecorder()
	// handler2 := http.HandlerFunc(CreateBasket)
	// handler3 := http.HandlerFunc(AddItemInBasket)
	// handler3.ServeHTTP(rr2, req2)
	// if status2 := rr2.Code; status2 != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status2, http.StatusOK)
	// }
	// if handler2 != nil {
	// }

	// Check the response body is what we expect.
	// fmt.Println(rr2.Body.String())
	// expected := `[{"Total":0,"Items":[]}]`
	// if strings.TrimSuffix(rr2.Body.String(), "\n") != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr2.Body.String(), expected)
	// }
}
