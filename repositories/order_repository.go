package repositories

import "applicationDesignTest/models"

type OrderRepository interface {
	CreateOrder(order models.Order) error
}

type InMemoryOrderRepository struct {
	orders []models.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{}
}

func (r *InMemoryOrderRepository) CreateOrder(order models.Order) error {
	r.orders = append(r.orders, order)
	return nil
}
