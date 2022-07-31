package controllers

import (
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// var event models.Event

	// json.NewDecoder(r.Body).Decode(&event)

	// return the event in JSON format
	// json.NewEncoder(w).Encode(event)

	// send to service
	// events.Create(&event)

	// json.NewEncoder(w).Encode(event)

	w.Write([]byte("Hello World!"))
}