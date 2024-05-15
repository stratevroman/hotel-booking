package services

import (
	"applicationDesignTest/models"
	"applicationDesignTest/repositories"
	"applicationDesignTest/utils"
	"errors"
)

type OrderService struct {
	availabilityRepo repositories.RoomAvailabilityRepository
	orderRepo        repositories.OrderRepository
}

var ErrNotAvailableRoom = errors.New("hotel room is not available for selected dates")

func NewOrderService(availabilityRepo repositories.RoomAvailabilityRepository, orderRepo repositories.OrderRepository) OrderService {
	return OrderService{availabilityRepo: availabilityRepo, orderRepo: orderRepo}
}

func (os *OrderService) CreateOrder(order models.Order) error {
	daysToBook, err := utils.DaysBetween(order.From, order.To)

	if err != nil {
		utils.LogErrorf(err.Error())
		return err
	}

	for _, day := range daysToBook {
		availability, err := os.availabilityRepo.GetAvailabilityByDate(day)
		if err != nil {
			return err
		}

		if availability == nil || availability.Quota <= 0 {
			return ErrNotAvailableRoom
		}

		if err := os.availabilityRepo.DecreaseQuotaByDate(day); err != nil {
			return err
		}
	}

	err = os.orderRepo.CreateOrder(order)
	if err != nil {
		utils.LogErrorf("Order created error: %v", order)
		return err
	}

	utils.LogInfo("Order successfully created: %v", order)
	return nil
}
