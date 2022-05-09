package main

import (
	"fmt"
	//"html"
	"log"
	"net/http"
	"strings"

	//"net/url"
	"strconv"
	"sync"
	//"os"
)

var counter int
var mutex = &sync.Mutex{}

// Function to echo the request
func echoString (w http.ResponseWriter, r *http.Request) {
	var echoString = strings.TrimLeft(r.URL.Path, "/echo/")
	fmt.Fprintf(w, "hello %s", echoString)
}

// Function to increment the counter
func incrementCounter (w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}



func main() {
	// Serve static website
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	// Servce echo function
	http.HandleFunc("/echo/", echoString)

	// Serve increment function
	http.HandleFunc("/increment", incrementCounter)

	// Serve the generic hi message
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}