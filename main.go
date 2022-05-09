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
	fmt.Fprintf(w, "hello %s", strings.TrimPrefix(r.URL.Path, "/"))
}

// Function to increment the counter
func incrementCounter (w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}



func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

	log.Fatal(http.ListenAndServe("localhost:8081", nil))

}