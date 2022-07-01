package main

import (
	"course/controller"
	"course/middlewares"
	"course/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"  // used to print the request and response to the consolog (for debugging)
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

/*
Write entire gin logs into the gin.log file.
*/
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	/* Setup the log writer for the server */
	setupLogOutput()

	/**
	 * instead of using gin.Default, replacing it by these lines of code 
	 */
	server := gin.New()

	server.Use(
		gin.Recovery(), 
		middlewares.Logger(), 
		middlewares.BasicAuth(),
		gindump.Dump(),)
	server.Use(gin.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8000")
}