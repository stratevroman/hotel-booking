package handlers

import (
	"applicationDesignTest/models"
	"applicationDesignTest/services"
	"applicationDesignTest/utils"
	"encoding/json"
	"errors"
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
		utils.LogErrorf("Decoding error", err.Error())
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	err = oh.orderService.CreateOrder(newOrder)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrIncorrectDates):
			http.Error(w, "Incorrect date range "+err.Error(), http.StatusBadRequest)
		case errors.Is(err, services.ErrNotAvailableRoom):
			http.Error(w, "Selected room is not available "+err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Failed to create order "+err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newOrder)
	if err != nil {
		return
	}
}
