package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", putAlbums)
	router.DELETE("/albums/:id", deleteAlbums)

	router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Get all albums"})
}

func postAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Creare albums"})
}

func getAlbumByID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Get an albums"})
}

func putAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Update an albums"})
}

func deleteAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Delete an albums"})
}