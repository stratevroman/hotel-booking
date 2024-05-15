// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
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
	err := http.ListenAndServe(":8081", r)

	if errors.Is(err, http.ErrServerClosed) {
		utils.LogInfo("Server closed")
	} else if err != nil {
		utils.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
