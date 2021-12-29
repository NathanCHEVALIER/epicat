package main

// picture represents data about an epicat picture.
type picture struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Cat string  `json:"cat"`
    Url  float64 `json:"url"`
}

// pictures slice to seed data.
var pictures = []picture{
    {ID: "1", Title: "", Cat: "Hoppy", Url: ""},
    {ID: "2", Title: "", Cat: "Hoppy", Url: ""},
    {ID: "3", Title: "", Cat: "Hoppy", Url: ""},
}

// getPictures responds with the list of all pictures as JSON.
func getPictures(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pictures)
}
