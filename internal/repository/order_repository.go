package repository


import (
	"github.com/Abhilash3108/E-commerce-with-go/internal/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
    DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
    return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *model.Order) error {
    return r.DB.Create(order).Error
}