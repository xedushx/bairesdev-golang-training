package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id    string  `json:"Id"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type ProductInventory struct {
	Product  Product `json:"Product"`
	Quantity int     `json:"Quantity"`
}

var inventory []ProductInventory

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createProductInventory(w http.ResponseWriter, r *http.Request) {
	var newProductInventory ProductInventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Product Inventory product and quantity only in order to update")
	}

	json.Unmarshal(reqBody, &newProductInventory)
	inventory = append(inventory, newProductInventory)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProductInventory)
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]

	for _, singleProduct := range inventory {
		if singleProduct.Product.Id == productID {
			json.NewEncoder(w).Encode(singleProduct)
		}
	}
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(inventory)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]
	var updatedProduct ProductInventory

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product code, name and price only in order to update")
	}
	json.Unmarshal(reqBody, &updatedProduct)

	for i, inventoryDb := range inventory {
		if inventoryDb.Product.Id == productID {
			inventoryDb.Product.Code = updatedProduct.Product.Code
			inventoryDb.Product.Name = updatedProduct.Product.Name
			inventoryDb.Product.Price = updatedProduct.Product.Price
			inventoryDb.Quantity = updatedProduct.Quantity
			inventory[i] = inventoryDb
			json.NewEncoder(w).Encode(inventoryDb)
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]

	for i, singleProduct := range inventory {
		if singleProduct.Product.Id == productID {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Fprintf(w, "The product with ID %v has been deleted successfully", productID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/inventory", createProductInventory).Methods("POST")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/product/{id}", getOneProduct).Methods("GET")
	router.HandleFunc("/inventory/product/{id}", updateProduct).Methods("PATCH")
	router.HandleFunc("/inventory/product/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
