package service


import (
	"github.com/Abhilash3108/E-commerce-with-go/internal/model"
)
type OrderRepository interface {
	Save(order *model.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.repo.Save(order)
}