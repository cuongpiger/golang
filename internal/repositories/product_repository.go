package repositories

import "github.com/cuongpiger/golang/internal/models"

type ProductRepositoryInterface interface {
	Add(product models.Product) error
}