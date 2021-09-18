
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"encoding/json"
)

const Version = "v1.2.0"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`Start by browsing /version or /duration
Try also with Header "Accept: application/json"
`)
}

type VersionResp struct {
	Version string `json:"version"`
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	switch accept {
	case "application/json":
		d := VersionResp{Version: Version}
		json.NewEncoder(w).Encode(d)
	default:
		fmt.Fprintf(w, "%s\n", Version)
	}
}

type DurationResp struct {
	Duration string `json:"duration"`
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
	reg, _ := regexp.Compile("[[:digit:].]+[smh]")
	match := reg.FindAllString(string(out), -1)
	if (len(match) < 3) {
		emsg := "Invalid systemd-analyze output"
		http.Error(w, emsg, http.StatusInternalServerError)
		/* Fix: http.Error doesn't have time to send the error */
		log.Fatal(emsg)
	}

	accept := r.Header.Get("Accept")
	switch accept {
	case "application/json":
		d := DurationResp{Duration: match[2]}
		json.NewEncoder(w).Encode(d)
	default:
		fmt.Fprintf(w, "%s\n", match[2])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", handlerVersion)
	http.HandleFunc("/duration", handlerDuration)

	log.Print("Server ready, endpoints: /version and /duration")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

