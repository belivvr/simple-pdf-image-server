package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	file, err := ioutil.ReadFile("./files/" + r.URL.Path)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode([][]byte{file})
}
