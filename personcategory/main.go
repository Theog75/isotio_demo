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

var mgoSession *mgo.Session

func dbcount(c *gin.Context) {
	updatemongo := mgoSession.DB(os.Getenv("MONGO_DATABASE")).C("actors")
	directors, err := updatemongo.Find(bson.M{"category": "director"}).Count()
	if err != nil {
		panic(err)
	}
	actors, err := updatemongo.Find(bson.M{"category": "actor"}).Count()
	if err != nil {
		panic(err)
	}
	actresses, err := updatemongo.Find(bson.M{"category": "actress"}).Count()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s has won %d games.\n", gamesWon)
	if err != nil {
		fmt.Println("Could not update mongo")
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"directors": directors,
		"actors":    actors,
		"actresses": actresses,
	})
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
	router.Run(":8099")
}
