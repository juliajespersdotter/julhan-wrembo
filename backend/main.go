package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type joke struct {
	ID        string `json:"id"`
	TELLING   string `json:"telling"`
	PUNCHLINE string `json:"punchline"`
	CATEGORY  string `json:"category"`
}

var jokes = []joke{
	{ID: "1", TELLING: "Jag har börjat säga moucho till mina spanska vänner...", PUNCHLINE: "...Det betyder mycket för dom.", CATEGORY: "HRAPPROVED"},
	{ID: "2", TELLING: "Vad är det för likhet mellan en nyförlöst mamma och någon som kommer i tid till bussen?", PUNCHLINE: "Slapp springa", CATEGORY: "18+"},
}

func getJokes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jokes)
}

func main() {
	router := gin.Default()
	router.GET("/jokes", getJokes)

	router.Run("localhost:8000")
}
