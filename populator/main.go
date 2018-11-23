package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

type Crew struct {
	Collection string `json:"collection"`
	Titleid    string `json:"tconst" binding:"required"`
	Ordering   string `json:"ordering"`
	Nconst     string `json:"nconst"`
	Category   string `json:"category"`
	Job        string `json:"job"`
	Characters string `json:"carachters"`
}

type Crewmongo struct {
	Titleid    string
	Ordering   string
	Nconst     string
	Category   string
	Job        string
	Characters string
}

type Names struct {
	Collection        string `json:"collection"`
	Nconst            string `json:"nconst" binding:"required"`
	PrimaryName       string `json:"primaryName"`
	BirthYear         string `json:"birthYear"`
	DeathYear         string `json:"deathYear"`
	PrimaryProfession string `json:"primaryProfession"`
	KnownForTitles    string `json:"knownForTitles"`
}

type Namesmongo struct {
	Nconst            string
	PrimaryName       string
	BirthYear         string
	DeathYear         string
	PrimaryProfession string
	KnownForTitles    string
}

func populatenamesDB(c *gin.Context) {
	nameline := new(Names)
	err := c.BindJSON(&nameline)
	if err != nil {
		fmt.Println("Main Error", err)
		c.AbortWithError(400, err)
		return
	}
	c.String(200, fmt.Sprintf("%#v", nameline))

	// fmt.Println("Collection: " + nameline.Collection)
	// Optional. Switch the session to a monotonic behavior.
	updatemongo := mgoSession.DB(os.Getenv("MONGO_DATABASE")).C(nameline.Collection)
	err = updatemongo.Insert(&Namesmongo{nameline.Nconst, nameline.PrimaryName, nameline.BirthYear, nameline.DeathYear, nameline.PrimaryProfession, nameline.KnownForTitles})
	if err != nil {
		fmt.Println("Could not update  mongo collection names " + nameline.Collection)
		// log.Fatal(err)
	} else {
		fmt.Println("Updated: " + nameline.PrimaryName)
	}
}

func populatecrewsDB(c *gin.Context) {
	crewline := new(Crew)
	err := c.BindJSON(&crewline)
	if err != nil {
		fmt.Println("Main Error", err)
		c.AbortWithError(400, err)
		return
	}
	c.String(200, fmt.Sprintf("%#v", crewline))

	// Optional. Switch the session to a monotonic behavior.
	updatemongo := mgoSession.DB(os.Getenv("MONGO_DATABASE")).C(crewline.Collection)
	err = updatemongo.Insert(&Crewmongo{crewline.Titleid, crewline.Ordering, crewline.Nconst, crewline.Category, crewline.Job, crewline.Characters})
	if err != nil {
		fmt.Println("Could not update mongo collection actors: " + crewline.Collection)
		// log.Fatal(err)
	} else {
		fmt.Println("Updated: " + crewline.Titleid)
	}
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
		fmt.Println("Could not update mongo mongo collection titles: " + titleline.Collection)
		// log.Fatal(err)
	} else {
		fmt.Println("Updated: " + titleline.Title)
	}
}

// func GetMongoSessionOld() *mgo.Session {
// 	if mgoSession == nil {
// 		var err error
// 		mgoSession, err = mgo.Dial(os.Getenv("MONGO_URL"))
// 		if err != nil {
// 			log.Fatal("Failed to start the Mongo session")
// 		}
// 	}
// 	return mgoSession.Clone()
// }
// GetMongoSession is an alternative function
func GetMongoSession() *mgo.Session {
	fmt.Println(os.Getenv("MONGO_URL") + " " + os.Getenv("MONGO_DATABASE") + " " + os.Getenv("MONGO_USER") + " " + os.Getenv("MONGO_PASSWORD"))
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{(os.Getenv("MONGO_URL"))},
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
		}
	}
	return mgoSession.Clone()
}

func init() {
	mgoSession := GetMongoSession()
	mgoSession.SetMode(mgo.Monotonic, true)
}
func main() {
	fmt.Println("starting populator V1.1.2")
	// mgoSession, err = mgo.Dial(os.Getenv("MONGO_URL"))
	defer mgoSession.Close()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	fmt.Println()
	v1 := router.Group("/populatetitles")
	{
		v1.POST("/", populatetitlesDB)
	}
	v2 := router.Group("/populatecrews")
	{
		v2.POST("/", populatecrewsDB)
	}
	v3 := router.Group("/populatenames")
	{
		v3.POST("/", populatenamesDB)
	}

	router.Run(":8088")
}
