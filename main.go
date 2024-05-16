package main

import (
	"applicationDesignTest/handlers"
	"applicationDesignTest/models"
	"applicationDesignTest/repositories"
	"applicationDesignTest/services"
	"applicationDesignTest/utils"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"sync"
)

func main() {
	orderService := services.NewOrderService(
		repositories.NewInMemoryRoomAvailabilityRepository(&sync.Mutex{}, models.Availability),
		repositories.NewInMemoryOrderRepository(),
	)
	orderHandler := handlers.NewOrderHandler(orderService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/orders", orderHandler.CreateOrder)

	utils.LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", r)

	if errors.Is(err, http.ErrServerClosed) {
		utils.LogInfo("Server closed")
	} else if err != nil {
		utils.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
