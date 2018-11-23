package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Nconst         string `bson:"nconst"`
	PrimaryName    string `bson:"primaryname"`
	BirthYear      string `bson:"birthyear"`
	DeathYear      string `bson:"deathyear"`
	KnownForTitles string `bson:"knownfortitles"`
}

type RequestAction struct {
	Searchstring string `json:"searchstring"`
}

var mgoSession *mgo.Session

func dbcount(c *gin.Context) {

	// fmt.Printf("%s", c.Request.Body)
	var people []Person
	nameline := new(RequestAction)
	err := c.BindJSON(&nameline)
	if err != nil {
		fmt.Println("Json Bind Error", err)
		c.AbortWithError(400, err)
		return
	}
	// regexpattern := "`" + nameline.Searchstring + "`"
	regexpattern := nameline.Searchstring

	fmt.Println("RegEXP " + regexpattern)
	updatemongo := mgoSession.DB(os.Getenv("MONGO_DATABASE")).C("names")
	// gamesWon, err := updatemongo.Find(bson.M{}).Count()

	updatemongo.Find(bson.M{"primaryname": bson.M{"$regex": bson.RegEx{regexpattern, ""}}}).Limit(10).All(&people)
	// err := updatemongo.Find(bson.M{"primaryName": bson.M{"$regex": bson.RegEx{Pattern: `/dav/`}}}).All(&people)
	// searchres, err := updatemongo.Find().Limit(100).All()
	if err != nil {
		panic(err)
	}
	fmt.Println("Results......")

	// fmt.Println(dta.Directors)
	for _, pps := range people {
		fmt.Println("Name: " + pps.PrimaryName)

		// fmt.Printf("%v", pps)
	}
	// fmt.Println(string(people))

	// returnresult, _ := json.Marshal(people)
	// fmt.Println(returnresult)
	c.JSON(200, people)
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
	// mgoSession, err = mgo.Dial(os.Getenv("MONGO_URL"))
	defer mgoSession.Close()

	router := gin.Default()
	v1 := router.Group("/")
	{
		v1.POST("/", dbcount)
	}
	router.Run(":8100")
}
