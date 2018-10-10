package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if html.EscapeString(r.URL.Path) == "/liran" {
			fmt.Println("Correct")
		}
	})

	log.Fatal(http.ListenAndServe(":8089", nil))

}
