package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Order1 struct {
	ID           string   `json:"id"`
	CustomerName string   `json:"customer_name"`
	Items        []string `json:"items"`
	TotalAmount  float64  `json:"total_amount"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

var (
	orders1 = make(map[string]Order)
	lock1   = sync.Mutex{}
)

// and a 201 Created status code. If JSON decoding fails, it returns a 400 Bad Request error.
func CreateOrder1(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	orders[order.ID] = order
	lock.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// The function uses a mutex to safely access the shared orders map in a concurrent environment.
func GetOrder1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	lock.Lock()
	order, exists := orders[id]
	lock.Unlock()

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// UpdateOrder1 updates an existing order by its ID. It accepts a JSON payload with updated order details and replaces the existing order while preserving the original creation timestamp.
// 
// The function retrieves the order ID from the URL parameters, decodes the request body into an Order struct,
// and updates the order in the global orders map. If the order does not exist, it returns a 404 Not Found error.
// 
// Parameters:
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing the updated order details
//
// Returns:
//   - 200 OK with the updated order JSON if successful
//   - 400 Bad Request if the JSON payload is invalid
//   - 404 Not Found if the order with the specified ID does not exist
//
// Example:
//   PUT /orders/123
//   {
//     "customer_name": "Updated Customer",
//     "items": ["New Item"],
//     "total_amount": 99.99
//   }
func UpdateOrder1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedOrder Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lock.Lock()
	currentOrder, exists := orders[id]
	if exists {
		updatedOrder.CreatedAt = currentOrder.CreatedAt
		updatedOrder.UpdatedAt = updatedOrder.UpdatedAt
		orders[id] = updatedOrder
	}
	lock.Unlock()

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updatedOrder)
}

// Returns a 204 No Content status upon successful deletion.
func DeleteOrder1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	lock.Lock()
	_, exists := orders[id]
	if exists {
		delete(orders, id)
	}
	lock.Unlock()

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListOrders1 retrieves and returns a list of all existing orders. It acquires a lock to safely access the orders map, creates a slice of all orders, and then encodes the list as a JSON response. The function does not filter or modify the orders during retrieval.
func ListOrders1(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	orderList := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, order)
	}
	lock.Unlock()

	json.NewEncoder(w).Encode(orderList)
}

// The function is thread-safe, using a mutex to protect concurrent access to the orders map.
func GetOrderStatus1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	lock.Lock()
	order, exists := orders[id]
	lock.Unlock()

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"status": order.Status}
	json.NewEncoder(w).Encode(response)
}

// main31 sets up and starts an HTTP server for managing orders using Gorilla Mux router.
// It defines routes for various order-related operations including creating, listing, retrieving,
// updating, deleting orders, and checking order status. The server listens on port 8080.
// The function will block and log any fatal errors encountered during server startup.
func main31() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", CreateOrder).Methods("POST")
	r.HandleFunc("/orders", ListOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", GetOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}/status", GetOrderStatus).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
