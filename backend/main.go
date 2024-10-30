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

func postJoke(c *gin.Context) {
	var newJoke joke

	if err := c.BindJSON(&newJoke); err != nil {
		return
	}

	jokes = append(jokes, newJoke)
	c.IndentedJSON(http.StatusCreated, newJoke)
}

func main() {
	router := gin.Default()
	router.GET("/jokes", getJokes)
	router.GET("/page", getPage)

	router.POST("/jokes", postJoke)

	router.Run("localhost:8000")
}
