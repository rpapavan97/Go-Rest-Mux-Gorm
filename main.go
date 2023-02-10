package main

import (
	"log"
	"net/http"
	"sellerapp/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", routes.CreateOrder).Methods("POST")
	// Read
	router.HandleFunc("/orders/{order_id}", routes.GetOrder).Methods("GET")
	// Read-all
	router.HandleFunc("/orders", routes.GetOrders).Methods("GET")
	// Update
	router.HandleFunc("/orders/{order_id}", routes.UpdateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{order_id}", routes.DeleteOrder).Methods("DELETE")
	// Initialize db connection
	routes.InitDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
