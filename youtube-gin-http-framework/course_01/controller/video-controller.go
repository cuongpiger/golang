package controller

import (
	"course/entity"
	"course/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{service: service}
}

func (c *controller) FindAll() []entity.Video
func (c *controller) Save(ctx *gin.Context)