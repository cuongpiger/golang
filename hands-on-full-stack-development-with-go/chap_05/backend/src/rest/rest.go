package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	// get products
	r.GET("/products", h.GetProducts)

	// get promos
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/user"); {
	  	userGroup.POST("/:id/signout", h.SignOut)
	  	userGroup.GET("/:id/orders", h.GetOrders)
	}
  
	usersGroup := r.Group("/users"); {
		// post user sign in
		usersGroup.POST("/signin", h.SignIn)

		// add a user
		usersGroup.POST("", h.AddUser)

		// post purchase charge
		usersGroup.POST("/charge", h.Charge)
	}
  
	return r.Run(address)
}

func RunAPI(address string) error {
	h, err := NewHandler()

	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}