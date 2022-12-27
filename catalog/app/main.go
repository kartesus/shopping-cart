package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/products", getAllProducts)
	http.HandleFunc("/products/", getProduct)
	http.ListenAndServe(":8000", nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := ioutil.ReadDir("data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []map[string]any{}
	for _, file := range files {
		b, err := ioutil.ReadFile(fmt.Sprintf("data/%s", file.Name()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var product map[string]any
		err = json.Unmarshal(b, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadFile(fmt.Sprintf("data/product-%d.json", id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var product map[string]any
	err = json.Unmarshal(b, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
