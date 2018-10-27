// Package of main frontend
package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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
	// os.Stderr.WriteString("Starting Front END")
	fmt.Println("Starting Front End")
	tmpl := template.Must(template.ParseFiles("upload.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			fmt.Println("file upload ok")
			return
		} else {
			fmt.Println("method:", r.Method)

			r.ParseMultipartForm(2000 << 20)
			file, handler, err := r.FormFile("uploadFile")
			if err != nil {
				fmt.Println(err)
				return
			}

			collection := r.FormValue("collection")
			if err != nil {
				fmt.Println(err)
				return
			}

			defer file.Close()
			// fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("Error opening file:", handler.Filename, err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			var filedata []string

			filedata = filesList()
			fmt.Printf("%v", filedata)
			dt := Fileupload{
				Success:   true,
				FilesList: filedata,
			}
			// filestatus := Fileupload{true, filedata}
			tmpl.Execute(w, dt)
			fmt.Println(handler.Filename)
			readUploadedFile(handler.Filename, collection)

		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ping")

		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":8088", nil)
}

func filesList() []string {
	var fileList []string
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	return fileList
}

func readUploadedFile(filename string, collection string) {
	fmt.Println("Reading file" + filename)
	file, err := os.Open("/" + filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	//
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	// csvData, err := reader.ReadAll()
	// var tsvlines []string
	// var tsvline string

	// fmt.Println()reflect.TypeOf("%T", csvData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")

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
	url := os.Getenv("POPULATOR_URL")
	fmt.Println(s[1] + " " + s[2] + " " + collection)
	fmt.Println(len(s))
	// url := "http://restapi3.apiary.io/notes"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
