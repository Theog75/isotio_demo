// Package of main frontend
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Dbcount is the number of items in the database
type Dbcount struct {
	count string
}

type Professions struct {
	Directors int `json:"directors"`
	Actors    int `json:"actors"`
	Actresses int `json:"actresses"`
}

type Person struct {
	Nconst         string `json:"Nconst"`
	PrimaryName    string `json:"PrimaryName"`
	BirthYear      string `json:"BirthYear"`
	DeathYear      string `json:"DeathYear"`
	KnownForTitles string `json:"KnownForTitles"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/img", "./img")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fe1.tmpl", gin.H{
			"title":         "RedHat Istio Demo",
			"version":       "1.0.17",
			"TotalTitles":   "Total Movie titles in DB",
			"ActorCounter":  "Crew persons in DB",
			"CategoryCount": "Professions",
		})
	})

	router.POST("/dbcount", func(c *gin.Context) {
		url := os.Getenv("DBCOUNTER_URL")
		jsonValue, _ := json.Marshal("countitems")
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		// req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		res := string(body)
		c.JSON(200, res)
	})

	router.POST("/actorcount", func(c *gin.Context) {
		url := os.Getenv("ACTORCOUNTER_URL")
		jsonValue, _ := json.Marshal("countitems")
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		// req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		res := string(body)
		c.JSON(200, res)
	})

	router.POST("/personcategory", func(c *gin.Context) {
		url := os.Getenv("PERSONCATEGORY_URL")
		jsonValue, _ := json.Marshal("countitems")
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		// req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var dta = new(Professions)

		v := json.Unmarshal([]byte(body), &dta)
		if v != nil {
			panic(v)
		}
		// fmt.Println(dta.Directors)
		c.String(200, "<div class='proflien'>Actors: "+strconv.Itoa(dta.Actors)+"</div>"+"<div class='proflien'>Actresses: "+strconv.Itoa(dta.Actresses)+"</div>"+"<div class='proflien'>Directors: "+strconv.Itoa(dta.Directors)+"</div>")
	})

	router.POST("/searchperson", func(c *gin.Context) {
		searchstring := c.PostForm("searchstring")
		jsonStr := []byte(`{"searchstring": "` + searchstring + `"}`)
		url := os.Getenv("SEARCHPERSON_URL")
		// jsonValue, _ := json.Marshal(jsonStr)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var dta []Person

		v := json.Unmarshal([]byte(body), &dta)
		if v != nil {
			panic(v)
		}
		// fmt.Println("printing responses....")
		// fmt.Println(string(body))
		var searchResp string
		for _, pps := range dta {
			fmt.Println("Got Actor Name: " + pps.PrimaryName)
			searchResp = searchResp + "<div class='sres' id='" + pps.Nconst + "'>" + pps.PrimaryName + "</div>"
			// fmt.Printf("%v", pps)
		}
		c.String(200, searchResp)
	})

	router.Run(":8080")
}
