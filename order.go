package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Order struct {
	ID           string   `json:"id"`
	CustomerName string   `json:"customer_name"`
	Items        []string `json:"items"`
	TotalAmount  float64  `json:"total_amount"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

var (
	orders = make(map[string]Order)
	lock   = sync.Mutex{}
)

// and a 201 Created status code. If the JSON decoding fails, it returns a 400 Bad Request error.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
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
func GetOrder(w http.ResponseWriter, r *http.Request) {
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

// UpdateOrder handles the HTTP request to update an existing order by its ID. It retrieves the order from the request parameters,
// decodes the updated order details from the request body, and updates the order in the orders map. If the order does not exist,
// it returns a 404 Not Found error. The function preserves the original creation timestamp and updates the modification timestamp.
// It uses a mutex to ensure thread-safe access to the orders map during the update operation.
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
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

// The function uses a mutex to ensure thread-safe access to the orders map during deletion.
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
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

// ListOrders retrieves all existing orders and returns them as a JSON response. It creates a slice of all orders from the global orders map, ensuring thread-safe access through a mutex lock.
func ListOrders(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	orderList := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, order)
	}
	lock.Unlock()

	json.NewEncoder(w).Encode(orderList)
}

// GetOrderStatus retrieves the status of a specific order by its ID.
// It takes an HTTP request with an order ID parameter, checks if the order exists,
// and returns the order's status as a JSON response. If the order is not found,
// it returns a 404 Not Found error.
//
// Parameters:
//   - w: HTTP response writer to send the response
//   - r: HTTP request containing the order ID in URL parameters
//
// Returns:
//   - JSON response with the order status if found
//   - 404 Not Found error if the order does not exist
//
// Example:
//   GET /orders/123/status will return {"status": "pending"}
func GetOrderStatus(w http.ResponseWriter, r *http.Request) {
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

// main3 sets up and starts the HTTP server for the order management API.
// It configures routes for creating, listing, retrieving, updating, and deleting orders,
// as well as checking order status. The server runs on port 8080 and uses Gorilla Mux for routing.
// If the server fails to start, it will log a fatal error and terminate the application.
func main3() {
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
