package rest

import (
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/dblayer"
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

// ___________________________________________________________________________________ Handler class
type Handler struct {
	db dblayer.DBLayer
}

/**
 * Constructor
 */
func NewHandler() (*Handler, error) {
	// this creates a new pointer to the Handler object
	return new(Handler), nil
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {  // if the db is not initialized then return
		return
	}

	// if db is not nil
	products, err := h.db.GetAllProducts()
	
	// there is an error while getting the products from the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// successfully get all products from the database
	c.JSON(http.StatusOK, &products)
}

/**
 * Get all promotions from the database
 */
func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {  // there is no database connection
		return
	}

	promos, err := h.db.GetPromos()

	// there is an error while getting the promotions from the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// successfully get all promotions from the database
	c.JSON(http.StatusOK, &promos)
}

/**
 * Loggin in a user
 */
func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {  // can not connect to the database
		return
	}

	var customer *models.Customer
	err := c.ShouldBindJSON(customer)  // read data from json to `customer`

	// json data is not valid
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// json data is valid
	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	
	// there is an error while signing in the user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// successfully signed in the user
	c.JSON(http.StatusOK, &customer)
}

/**
 * Add a new user
 */
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {  // can not connect to the database
		return
	}

	var customer *models.Customer
	err := c.ShouldBindJSON(customer)  // read data from json to `customer`

	if err != nil {  // json data is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err = h.db.AddUser(customer)

	// there is an error while adding the user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// successfully added the user
	c.JSON(http.StatusOK, &customer)
}

/**
 * Sign out a user
 */
func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {  // can not connect to the database
		return
	}

	p := c.Param("id")  // get the id from the url
	id, err := strconv.Atoi(p)  // convert id to int data type

	if err != nil {  // id is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserById(id)

	if err != nil {  // there is an error while signing out the user
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

/**
 * Get the orders of a user
 */
func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {  // can not connect to the database
		return
	}

	p := c.Param("id")  // get the id from the url
	id, err := strconv.Atoi(p)  // convert id to int data type

	if err != nil {  // id is not valid
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.db.GetCustomerOrdersByID(id)  // get all orders of this user
	if err != nil {  // there is an error while getting the orders
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// successfully get all orders of this user
	c.JSON(http.StatusOK, &orders)
}

/**
 * User charges their order via credit card
 */
func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {  // can not connect to the database
		return
	}
}








