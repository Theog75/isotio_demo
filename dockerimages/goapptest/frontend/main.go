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
	Success bool
}

type FileList struct {
	filename string
}

func main() {
	// os.Stderr.WriteString("Starting Front END")
	fmt.Println("Starting Front End")
	tmpl := template.Must(template.ParseFiles("upload.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		} else {
			fmt.Println("method:", r.Method)

			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadFile")
			if err != nil {
				// fmt.Println(err)
				return
			}
			defer file.Close()
			// fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				// fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			var filedata []string

			filedata = FilesList()

			filestatus := Fileupload{true}
			tmpl.Execute(w, filestatus)
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
