// Package of main frontend
package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Fileupload struct {
	Success   bool
	FilesList []string
}

func main() {
	// os.Stderr.WriteString("Starting Front END")
	fmt.Println("Starting Front End")
	tmpl := template.Must(template.ParseFiles("upload.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
<<<<<<< HEAD:dockerimages/goapptest/frontend/main.go
			fmt.Println("file upload ok")
=======
			fmt.Println("Present form")
>>>>>>> 1eda71df4baa9582e36a61bc4d66ca9133521953:frontend/main.go
			return
		} else {
			fmt.Println("method:", r.Method)

			r.ParseMultipartForm(2000 << 20)
			file, handler, err := r.FormFile("uploadFile")
			if err != nil {
<<<<<<< HEAD:dockerimages/goapptest/frontend/main.go
				fmt.Println(err)
=======
				fmt.Println("Got File", err)
>>>>>>> 1eda71df4baa9582e36a61bc4d66ca9133521953:frontend/main.go
				return
			}
			defer file.Close()
			// fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
<<<<<<< HEAD:dockerimages/goapptest/frontend/main.go
				fmt.Println(err)
=======
				fmt.Println("Error opening file:", handler.Filename, err)
>>>>>>> 1eda71df4baa9582e36a61bc4d66ca9133521953:frontend/main.go
				return
			}
			defer f.Close()
			io.Copy(f, file)

			var filedata []string

			filedata = FilesList()
			fmt.Printf("%v", filedata)
			dt := Fileupload{
				Success:   true,
				FilesList: filedata,
			}
			// filestatus := Fileupload{true, filedata}
			tmpl.Execute(w, dt)
		}
	})

	http.ListenAndServe(":8088", nil)
}

func FilesList() []string {
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
