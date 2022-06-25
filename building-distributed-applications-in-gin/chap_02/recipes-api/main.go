package main

import (
	"encoding/json"
	"io/ioutil"
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
	file, _ := ioutil.ReadFile("recipes.json")  // read the json file into file variable
	json.Unmarshal([]byte(file), &recipes)  // save file to recipes slice
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
	
	c.JSON(http.StatusOK, gin.H{"id": recipe})  // return response to the client
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}

func main() {
	router := gin.Default()

	router.GET("/recipes", ListRecipesHandler)
	router.POST("/recipes", NewRecipeHandler)

	router.Run()
}
