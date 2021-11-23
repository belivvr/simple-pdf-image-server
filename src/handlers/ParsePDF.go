package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gen2brain/go-fitz"
)

func ParsePDF(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	doc, err := fitz.NewFromReader(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer doc.Close()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docToImageList(doc))
}
