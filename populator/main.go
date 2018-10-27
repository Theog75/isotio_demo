package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

// Titles struct for titles tsv file json string to mongodb
type Titles struct {
	Collection      string `json:"collection"`
	Titleid         string `json:"titleid" binding:"required"`
	Ordering        string `json:"ordering"`
	Title           string `json:"title"`
	Region          string `json:"region"`
	Language        string `json:"language"`
	Types           string `json:"types"`
	Attibutes       string `json:"attributes"`
	IsOriginalTitle string `json:"isOriginalTitle"`
}

type Titlesmongo struct {
	Titleid         string
	Ordering        string
	Title           string
	Region          string
	Language        string
	Types           string
	Attibutes       string
	IsOriginalTitle string
}

func populatetitlesDB(c *gin.Context) {
	titleline := new(Titles)
	err := c.BindJSON(&titleline)
	if err != nil {
		fmt.Println("Main Error", err)
		c.AbortWithError(400, err)
		return
	}
	c.String(200, fmt.Sprintf("%#v", titleline))

	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	updatemongo := session.DB("imdb").C(titleline.Collection)
	err = updatemongo.Insert(&Titlesmongo{titleline.Titleid, titleline.Ordering, titleline.Title, titleline.Region, titleline.Language, titleline.Types, titleline.Attibutes, titleline.IsOriginalTitle})
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.POST("/populatetitles", populatetitlesDB)
	}
	router.Run(":8088")
}
