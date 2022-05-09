package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"sync"
	//"os"
)

var counter int
var mutex = &sync.Mutex{}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", strings.Replace(html.EscapeString(r.URL.Path), "/", "", 1))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}