package signal

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Age              int
	Name, Occupation string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	p := Person{
		Age:        25,
		Name:       "Bisshwajit",
		Occupation: "Software Engineer",
	}
	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}
