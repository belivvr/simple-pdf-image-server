package handlers

import (
	"net/http"
	"os"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if os.RemoveAll("./files/"+r.URL.Path) != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("OK"))
}
