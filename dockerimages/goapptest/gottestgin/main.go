package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type controlSignal struct {
	Signal string `json:"signal"`
	State  bool   `json:"state"`
}

// func createTodo(c *gin.Context) {
// 	fmt.Println("testing")
// }
func fetchAllTodo(c *gin.Context) {
	todo := transformedTodo{ID: 5, Title: "test", Completed: true}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func initializeProcess(c *gin.Context) {
	todo := controlSignal{Signal: "Starting process", State: true}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func populateDB(c *gin.Context) {
	todo := controlSignal{Signal: "Populate DB", State: true}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func drawForm(c *gin.Context) {

}

func fetchSingleTodo(c *gin.Context) {}
func updateTodo(c *gin.Context)      {}
func deleteTodo(c *gin.Context)      {}

func main() {
	router := gin.Default()
	v1 := router.Group("/")
	{
		// v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/init", initializeProcess)
		v1.GET("/populate", populateDB)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
