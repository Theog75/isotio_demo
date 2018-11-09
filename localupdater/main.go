// Package of main frontend
package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Fileupload struct {
	Success   bool
	FilesList []string
}

type Pong struct {
	Respo string
	Stat  string
}

func main() {
	// fmt.Println(os.Args[1])
	// os.Stderr.WriteString("Starting Front END")
	fmt.Println("Starting population")
	readUploadedFile(os.Args[1], os.Args[2])
}

// func filesList() []string {
// 	var fileList []string
// 	files, err := ioutil.ReadDir("./")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	for _, f := range files {
// 		fileList = append(fileList, f.Name())
// 	}
// 	return fileList
// }

func readUploadedFile(filename string, collection string) {
	fmt.Println("Reading file" + filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		fmt.Println("Sending Json: " + scanner.Text())
		sendDataToMongo(s, collection)
	}

	// for _, each := range csvData {
	// s := strings.Split(each, "\t")
	// fmt.Print(each)
	// fmt.Println(reflect.TypeOf(each))
	// tsvlines = append(tsvlines, each...)
	// fmt.Print(tsvlines[0] + " " + tsvlines[1])
	// }
	// jsonData, err := json.Marshal(tsvline)
}

func sendDataToMongo(s []string, collection string) {
	var jsonStr map[string]string
	var uri string
	if collection == "titles" {
		uri = "/populatetitles"
	} else if collection == "actors" {
		uri = "/populatecrews"
	} else if collection == "names" {
		uri = "/populatenames"
	} else {
		panic("Unknown collection")
	}
	url := os.Getenv("POPULATOR_URL") + uri
	// fmt.Println(s[1] + " " + s[2] + " " + collection)
	// fmt.Println(len(s))
	// url := "http://restapi3.apiary.io/notes"
	if collection == "titles" {
		jsonStr = map[string]string{"collection": collection, "titleid": s[0], "ordering": s[1], "title": s[2], "region": s[3], "language": s[4], "types": s[5], "attributes": s[6], "isOriginalTitle": s[7]}
	} else if collection == "actors" {
		jsonStr = map[string]string{"collection": collection, "titleid": s[0], "ordering": s[1], "nconst": s[2], "category": s[3], "job": s[4], "carachters": s[5]}
	} else if collection == "names" {
		jsonStr = map[string]string{"collection": collection, "nconst": s[0], "primaryName": s[1], "birthYear": s[2], "deathYear": s[3], "primaryProfession": s[4], "knownForTitles": s[5]}

	}
	jsonValue, _ := json.Marshal(jsonStr)
	// fmt.Println("URL:>", url)
	// var jsonStr = []byte(`{"titleid": ` + s[0] + `,"ordering": ` + s[1] + `,"title": ` + s[2] + `,"region":` + s[3] + `,"language": ` + s[4] + `,"types": ` + s[4] + `,"attribute": ` + s[5] + `,"isOriginalTitle": ` + s[6] + `}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
