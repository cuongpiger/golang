package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
This function is called after the server received a request.
It prints to the console some information about the request.
Called by the server.
*/
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,)
	})
}