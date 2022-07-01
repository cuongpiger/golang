package controller

import (
	"course/entity"
	"course/service"
	"course/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

/*
Using to check if a `title` of a video "is cool" or not. */
var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)  // register validator for this controller.

	return &controller{service: service}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}

	// Call validator to check the new Video object
	if err := validate.Struct(video); err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}