package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

var recipes []Recipe
func init() {
	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()  // generate a new id
	recipe.PublishedAt = time.Now() // set the publishedAt to now
	recipes = append(recipes, recipe)  // save to the local memory
	c.JSON(http.StatusOK, gin.H{"id": recipe})

}

func main() {
	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)

	router.Run()
}
