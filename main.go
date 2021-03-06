package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// picture represents data about an epicat picture.
type picture struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Cat   string `json:"cat"`
	Url   string `json:"url"`
}

// pictures slice to seed data.
var pictures = []picture{
	{ID: "1", Title: "", Cat: "Hoppy", Url: ""},
	{ID: "2", Title: "", Cat: "Hoppy", Url: ""},
	{ID: "3", Title: "", Cat: "Hoppy", Url: ""},
}

func main() {
	router := gin.Default()
	router.GET("/pictures", getPictures)
	router.GET("/pictures/:id", getPictureByID)
	router.GET("/pickpic", getRdmPicture)

	router.Run("localhost:8080")
}

// getPictures responds with the list of all pictures
func getPictures(c *gin.Context) {
	if len(pictures) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Picture"})
		return
	}

	c.IndentedJSON(http.StatusOK, pictures)
}

// getPictureByID locates the picture whose ID value matches the id
// parameter sent by the client, then returns that picture as a response.
func getPictureByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of pictures
	for _, a := range pictures {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Picture not found"})
}

// getRdmPicture pick a random picture in pictures
func getRdmPicture(c *gin.Context) {
	sz := len(pictures)

	if sz == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Picture"})
		return
	}

	c.IndentedJSON(http.StatusOK, pictures[rand.Intn(sz)])
	return
}
