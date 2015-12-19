package main

import (
	"fmt"
	"github.com/stensonb/project1042/Godeps/_workspace/src/github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
	"time"
)

// pin GPIO2 is physical pin 3 on the pi
var pin = rpio.Pin(2)

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	pin.Toggle()
	fmt.Fprintf(w, "toggled...")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format(time.RFC3339))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Project 1042</h1><form action='toggle' method='get'><button type='submit' value='Submit'>Toggle</button></form></body></html>")
}

func logRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		fn(w, r)
		log.Println(r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(startTime))
	}
}

func main() {

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	// setup web handlers
	http.HandleFunc("/", logRequest(mainHandler))
	http.HandleFunc("/time", logRequest(timeHandler))
	http.HandleFunc("/api/toggle", logRequest(toggleHandler))

	log.Println("Started Project 1042!")
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
