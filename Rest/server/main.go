package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

var database = map[int]Product{
	1: {ID: 1, Name: "DummyData", Price: 10000000},
}

var lastID int = 1

func main() {
	// 1. make mux router	
	mux := http.NewServeMux()


	// 3. mux handle function
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)
	mux.HandleFunc("GET /products/{id}", getProductByID)

	// 4. make server
	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}

	// 5. run server
	server.ListenAndServe()

}

// 2. make handler function
func listProduct(w http.ResponseWriter, r *http.Request){
	var products []Product

	for _, product := range database {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func createProduct(w http.ResponseWriter, r *http.Request){
	var product Product = Product{}
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}
	json.Unmarshal(bodyByte, &product)

	lastID++
	product.ID = lastID
	database[lastID] = product

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("berhasil menambahkan data"))
	
}

func updateProduct(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}

	var product Product = Product{ID: idInt}
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}
	json.Unmarshal(bodyByte, &product)

	database[idInt] = product
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("berhasil mengubah data"))
}

func deleteProduct(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}

	if _, ok := database[idInt]; !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte("data tidak ditemukan"))
		return
	}
	delete(database, idInt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("berhasil menghapus data"))
}

func getProductByID(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}

	var product Product = database[idInt]
	data, err := json.Marshal(product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

