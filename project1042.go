package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const servePort string = ":80"
const pin string = "23"
const html string = "<html><script>function toggle(){var xhr = new XMLHttpRequest();xhr.open('GET','api/toggle',true);xhr.send();}</script><body><h1>Project 1042</h1><button type='submit' value='Submit' onClick=toggle()>Toggle</button></body></html>"

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	err := exec.Command("gpioctl", "-t", pin).Run()
	if err != nil {
		fmt.Println("gpio toggle failed", err)
	}
	fmt.Fprintf(w, "toggled...")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format(time.RFC3339))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html)
}

func logRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		fn(w, r)
		log.Println(r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Since(startTime))
	}
}

func main() {
	exec.Command("gpioctl", "-c", pin, "OUT").Run()

	// setup web handlers
	http.HandleFunc("/", logRequest(mainHandler))
	http.HandleFunc("/time", logRequest(timeHandler))
	http.HandleFunc("/api/toggle", logRequest(toggleHandler))

	log.Println("Started Project 1042!")
	log.Println("Listening on", servePort)
	log.Fatal(http.ListenAndServe(servePort, nil))
}
