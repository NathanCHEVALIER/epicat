package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// picture represents data about an epicat picture.
type picture struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Cat string  `json:"cat"`
    Url  string `json:"url"`
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

    router.Run("localhost:8080")
}

// getPictures responds with the list of all pictures as JSON.
func getPictures(c *gin.Context) {
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