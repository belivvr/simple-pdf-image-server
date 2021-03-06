package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"simple-pdf-image-server/src/handlers"
)

func main() {
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	r.Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, ".jpg") ||
				strings.Contains(r.URL.Path, ".jpeg") ||
				strings.Contains(r.URL.Path, ".png") {
				handlers.ReadImage(w, r)
				return
			}

			if strings.Contains(r.URL.Path, ".pdf") {
				handlers.ReadPDF(w, r)
				return
			}

			handlers.ReadDir(w, r)
		})
	r.Methods("POST").HandlerFunc(handlers.Upload)
	r.Methods("PUT").HandlerFunc(handlers.Mkdir)
	r.Methods("DELETE").HandlerFunc(handlers.Delete)
	r.Methods("VIEW").Path("/pdf").HandlerFunc(handlers.ParsePDF)

	println("server listen in http://localhost:9000")
	http.ListenAndServe(":9000", r)
}
