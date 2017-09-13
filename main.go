package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(struct {
		Greeting string `json:"greeting"`
	}{"hello world"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func init() {
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))
}

func main() {
	addr := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(addr, nil))
}
