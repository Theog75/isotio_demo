package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

// var session *mgo.Session

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

	// Optional. Switch the session to a monotonic behavior.
	updatemongo := mgoSession.DB(os.Getenv("MONGO_DATABASE")).C(titleline.Collection)
	err = updatemongo.Insert(&Titlesmongo{titleline.Titleid, titleline.Ordering, titleline.Title, titleline.Region, titleline.Language, titleline.Types, titleline.Attibutes, titleline.IsOriginalTitle})
	if err != nil {
		fmt.Println("Could not update mongo")
		log.Fatal(err)
	}
}

func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(os.Getenv("MONGO_URL"))
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}

func init() {
	mgoSession := GetMongoSession()
	mgoSession.SetMode(mgo.Monotonic, true)
}
func main() {

	// mgoSession, err = mgo.Dial(os.Getenv("MONGO_URL"))
	defer mgoSession.Close()

	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.POST("/populatetitles", populatetitlesDB)
	}
	router.Run(":8088")
}
