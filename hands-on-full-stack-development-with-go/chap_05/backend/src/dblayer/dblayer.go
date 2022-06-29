package dblayer

import (
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/models"
)

type DBLayer interface {
	GetAllProducts() (*[]models.Product, error)
	GetPromos() (*[]models.Product, error)
	GetCustomerByID(int) (*models.Customer, error)
	GetProduct(uint) (*models.Product, error)
	AddUser(*models.Customer) (*models.Customer, error)
	SignInUser(username, password string) (*models.Customer, error)
	SignOutUser(int) error
	GetCustomerOrderByID(int) (*[]models.Order, error)
}