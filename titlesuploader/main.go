package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person struct for testing
type Person struct {
	Name  string
	Phone string
}

func uploadToMongo(c *gin.Context) {

}

func main() {

	router := gin.Default()
	v1 := router.Group("/")
	{
		// v1.POST("/", createTodo)
		// v1.GET("/", deleteTodo)
		// v1.GET("/init", deleteTodo)
		v1.POST("/uploadToMongo", uploadToMongo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run(":9000")

	session, err := mgo.Dial("langdb:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("peopletest").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
