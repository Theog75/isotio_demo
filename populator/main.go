package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Titles struct for titles tsv file json string to mongodb
type Titles struct {
	Titleid         string   `json:"titleid" binding:"required"`
	Ordering        int      `json:"ordering"`
	Title           string   `json:"title"`
	Region          string   `json:"region"`
	Language        string   `json:"language"`
	Types           []string `json:"types"`
	Attibutes       []string `json:"attributes"`
	IsOriginalTitle bool     `json:"isOriginalTitle"`
}

func populatetitlesDB(c *gin.Context) {
	var titleline Titles
	c.BindJSON(titleline)
	fmt.Println("GOT JSON" + titleline.Title)
	c.JSON(200, gin.H{"status": "200"})

	// session, err := mgo.Dial("langdb:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	//
	// // Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)
	//
	// updatemongo := session.DB("peopletest").C("people")
	// err = updatemongo.Insert(&titles{s)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func main() {
	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.POST("/populatetitles", populatetitlesDB)
	}
	router.Run(":8088")
}
