package refactoring

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CatagoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CatagoryName string `json:"categoryName"`
}

func GetAllProducts([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil

}

func AddProduct() (Product, error) {
	product := Product{ProductName: "microphone", CatagoryId: 1, UnitPrice: 200.00}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println("JSON'e dönüştürme hatası:", err)
		return
	}

	response, err := http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	return productResponse, nil

}
