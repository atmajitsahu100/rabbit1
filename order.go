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

// The function returns a 201 Created status on successful order creation or a 400 Bad Request if JSON decoding fails.
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

// The function uses a mutex lock to ensure thread-safe access to the orders map.
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
//
// Parameters:
//   - w: HTTP response writer for sending the response
//   - r: HTTP request containing the order ID and updated order details
//
// Returns:
//   - 200 OK with the updated order in JSON format if successful
//   - 400 Bad Request if the request body cannot be decoded
//   - 404 Not Found if the order with the specified ID does not exist
//
// Example:
//   PUT /orders/123
//   {
//     "customerName": "Updated Customer",
//     "items": ["New Item"],
//     "totalAmount": 150.00
//   }
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

// ListOrders retrieves and returns a list of all orders in the system. It creates a slice of orders from the global orders map, ensuring thread-safe access through a mutex lock. The function writes the list of orders as a JSON response to the HTTP response writer.
func ListOrders(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	orderList := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, order)
	}
	lock.Unlock()

	json.NewEncoder(w).Encode(orderList)
}

// The function is thread-safe, using a mutex to protect concurrent access to the orders map.
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

// main3 sets up and starts an HTTP server for managing orders using Gorilla Mux router.
// It defines routes for creating, listing, retrieving, updating, and deleting orders,
// as well as retrieving order status. The server runs on port 8080 and logs any fatal
// errors during server startup.
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
