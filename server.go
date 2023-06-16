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
	currentTime := JSONTime{Time: time.Now().Format(time.RFC3339)}
	jsonData, err := json.Marshal(currentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {}