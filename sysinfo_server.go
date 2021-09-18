
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Start by browsing /version or /duration")
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v0.1.0")
}

func handlerDuration(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("systemd-analyze").Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		/* Fix: http.Error doesn't have time to send the error */
		log.Fatal(err)
	}
	/*
	 * out should look like:
	 * Startup finished in 2.731s (kernel) + 23.298s (userspace) = 26.030s
	 */
	reg, _ := regexp.Compile("[[:digit:].]+")
	match := reg.FindAllString(string(out), -1)
	if (len(match) < 3) {
		emsg := "Invalid systemd-analyze output"
		http.Error(w, emsg, http.StatusInternalServerError)
		/* Fix: http.Error doesn't have time to send the error */
		log.Fatal(emsg)
	}
	fmt.Fprintf(w, "%ss", match[2])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", handlerVersion)
	http.HandleFunc("/duration", handlerDuration)

	log.Print("Server ready, endpoints: /version and /duration")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

