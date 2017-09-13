package skeleton

import (
	"encoding/json"
	"net/http"
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
}
