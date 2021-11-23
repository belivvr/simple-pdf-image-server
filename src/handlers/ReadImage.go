package handlers

import (
	"io/ioutil"
	"net/http"
)

func ReadImage(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./files/" + r.URL.Path)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.Write(file)
}
