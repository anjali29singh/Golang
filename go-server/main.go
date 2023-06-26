package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// build an API that provide client to get and add album for users

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums returns list of all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func postAlbums(c *gin.Context) {
	var newAlbum album
	//BindJSON to bind the recieved JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)

}
func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")

}
