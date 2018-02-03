package main

import (
	"html/template"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, struct {
		GaTrackingId string
	}{
		GaTrackingId: os.Getenv("GA_TRACKING_ID"),
	})
}
