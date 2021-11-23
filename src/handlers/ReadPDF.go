package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gen2brain/go-fitz"
)

func ReadPDF(w http.ResponseWriter, r *http.Request) {
	doc, err := fitz.New("./files" + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer doc.Close()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docToImageList(doc))
}
