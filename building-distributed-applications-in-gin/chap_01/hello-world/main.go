package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)


func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")  // get the `name` parameter from the URL
	c.JSON(200, gin.H{
		"message": "Hello, " + name,
	})
}

func randomName(c *gin.Context) {
	person := Person{
		FirstName: "Cường",
		LastName: "Dương",
	}
	c.XML(200, person)
}

func main() {
	router := gin.Default()
	
	/**
	 * Method: GET
	 * URL: /
	 * File request: my_requests/get.py
	 */
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	
	/**
	 * Method: GET
	 * URL: /random-name/
	 * File request: my_requests/get_random_name.py
	 */
	router.GET("/random-name", randomName)

	/**
	 * Method: GET
	 * URL: /<str:name>/
	 * File request: my_requests/get_name.py
	 */
	router.GET("/:name", IndexHandler)

	
	router.Run()
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	FirstName string `xml:"firstName,attr"`
	LastName string `xml:"lastName,attr"`
}