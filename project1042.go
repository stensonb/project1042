package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const pin string = "23"

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	err := exec.Command("gpioctl", "-t", pin).Run()
	if err != nil {
		fmt.Println("gopio toggle failed", err)
	}
	fmt.Fprintf(w, "toggled...")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format(time.RFC3339))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><script>function toggle(){var xhr = new XMLHttpRequest();xhr.open('GET','api/toggle',true);xhr.send();}</script><body><h1>Project 1042</h1><button type='submit' value='Submit' onClick=toggle()>Toggle</button></body></html>")
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
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
