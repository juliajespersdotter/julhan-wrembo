package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataChunk struct {
	DATA     []joke       `json:"data"`
	TYPE     string       `json:"type"`
	METADATA pageMetaData `json:"pagemetadata"`
}

type joke struct {
	ID        string `json:"id"`
	TELLING   string `json:"telling"`
	PUNCHLINE string `json:"punchline"`
	CATEGORY  string `json:"category"`
}

type pageMetaData struct {
	ID          string `json:"id"`
	PAGETITLE   string `json:"pagetitle"`
	DESCRIPTION string `json:"description"`
	LOGO        string `json:"logo"`
}

var jokes = []joke{
	{ID: "1", TELLING: "Jag har börjat säga moucho till mina spanska vänner...", PUNCHLINE: "...Det betyder mycket för dom.", CATEGORY: "HRAPPROVED"},
	{ID: "2", TELLING: "Vad är det för likhet mellan en nyförlöst mamma och någon som kommer i tid till bussen?", PUNCHLINE: "Slapp springa", CATEGORY: "18+"},
}

var pageMetaDataObject = pageMetaData{
	ID:          "P1",
	PAGETITLE:   "Världens bästa skämt",
	DESCRIPTION: "Pagetitlen beskriver bra",
	LOGO:        "https://logoipsum.com/artwork/327",
}

var dataChunkObject = dataChunk{
	DATA:     jokes,
	TYPE:     "PAGE",
	METADATA: pageMetaDataObject,
}

func getJokes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jokes)
}

func getPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dataChunkObject)
}

// postJoke adds a new joke to the list of jokes
func postJoke(c *gin.Context) {
	var newJoke joke

	if err := c.BindJSON(&newJoke); err != nil {
		return
	}

	jokes = append(jokes, newJoke)
	c.IndentedJSON(http.StatusCreated, newJoke)
}

// getJokeByID locates the joke whose ID value matches the id
func getJokeByID(c *gin.Context) {
	id := c.Param("id")

	// loop over jokes to find the id that mathes parameter
	for _, j := range jokes {
		if j.ID == id {
			c.IndentedJSON(http.StatusOK, j)
			return
		}
	}

	    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "joke not found"})

}

// Get all categories
func getCategories(c *gin.Context) {
    categories := make(map[string]bool)
    for _, joke := range jokes {
        categories[joke.CATEGORY] = true
    }
    
    uniqueCategories := make([]string, 0, len(categories))
    for category := range categories {
        uniqueCategories = append(uniqueCategories, category)
    }
    
    c.IndentedJSON(http.StatusOK, uniqueCategories)
}

// Get jokes by category
func getJokesByCategory(c *gin.Context) {
    category := c.Param("category")
    
    var categoryJokes []joke
    for _, joke := range jokes {
        if joke.CATEGORY == category {
            categoryJokes = append(categoryJokes, joke)
        }
    }
    
    c.IndentedJSON(http.StatusOK, categoryJokes)
}

// create categories
func createCategory(c *gin.Context) {
    var newJoke joke
    
    if err := c.BindJSON(&newJoke); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid joke format"})
        return
    }
    
    jokes = append(jokes, newJoke)
    c.IndentedJSON(http.StatusCreated, newJoke)
}

func main() {
    router := gin.Default()
	// GET general
    router.GET("/jokes", getJokes)
    router.GET("/page", getPage)
    router.GET("/categories", getCategories)
	
	// GET SPECIFIC
	router.GET("/jokes/:id", getJokeByID)
    router.GET("/jokes/category/:category", getJokesByCategory)

	// POST
    router.POST("/categories", createCategory)
    router.POST("/jokes", postJoke)

    router.Run("localhost:8000")
}

