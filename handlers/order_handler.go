package handlers

import (
	"applicationDesignTest/models"
	"applicationDesignTest/services"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	err = oh.orderService.CreateOrder(newOrder)
	if err != nil {
		/**	@todo перехватывать тип ошибки и отдавать соответсвующий код */
		http.Error(w, "Failed to create order "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)
	if err != nil {
		return
	}
}
