package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tecnologer/inventory/src/Inventory"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Server started...")

	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/product", CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	list, err := Inventory.GetProducts()
	if err != nil {
		printError("Error al obtener la lista de productos.", w)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		printError(fmt.Sprintf("Error al obtener el producto %s", params["id"]), w)
		return
	}

	product, err := Inventory.GetProduct(id)

	if err != nil {
		printError(fmt.Sprintf("Error al obtener el producto %s", params["id"]), w)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Inventory.Product
	_ = json.NewDecoder(r.Body).Decode(&p)

	if exists, _ := Inventory.ExistsProduct(p.ID); exists {
		printError(fmt.Sprintf("Error el producto con ID = %d, ya existe", p.ID), w)
		return
	}

	_, err := Inventory.InsertNewProduct(p)

	if err != nil {
		printError(fmt.Sprintf("Error al insertar el producto %d", p.ID), w)
		return
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		printError(fmt.Sprintf("Error id = %s es invalido", params["id"]), w)
		return
	}

	err = Inventory.DeleteProduct(id)

	if err != nil {
		printError(fmt.Sprintf("Error al eliminar el producto con id = %s.", params["id"]), w)
		return
	}
}

func printError(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Error"
	}

	w.Write([]byte(msg))
}
