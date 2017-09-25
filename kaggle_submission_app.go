package main

import (
	"fmt"
	"net/http"
	"os"
	"kaggle_submission_app/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/status", status)
	http.HandleFunc("/api/submissions", api.SubmissionsHandleFunc)
	http.HandleFunc("/api/submissions/", api.SubmissionHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome Kaggle Submission Tracker")
}

func echo(w http.ResponseWriter, r *http.Request){
	message := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func status(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}