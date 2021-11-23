package handlers

import (
	"net/http"
	"os"
)

func Mkdir(w http.ResponseWriter, r *http.Request) {
	if os.MkdirAll("./files"+r.URL.Path, 0777) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}
