package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JSONTime struct {
	Time string `json:"time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")
	currentTime := JSONTime{Time: time.Now().Format(time.RFC3339)} //  Get the current time and format it
	jsonData, err := json.Marshal(currentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/current-time", getCurrentTime)
	fmt.Println("Server started on port 8795")
	http.ListenAndServe(":8795", nil)
  }