package middlewares

import (
	"github.com/gin-gonic/gin"
)

/*
The username is 'pragmatic', the password is `reviews`
*/
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"pragmatic": "reviews",
	})
}