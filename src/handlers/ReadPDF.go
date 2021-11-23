package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gen2brain/go-fitz"
)

func ReadPDF(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len(r.URL.Path)-4:] == ".pdf" {
		numPage, err := getPageNum("./files" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("%d", numPage)))
		return
	}

	paths := strings.Split(r.URL.Path, "/")
	pageNumber, _ := strconv.Atoi(paths[len(paths)-1])
	path := strings.Join(paths[:len(paths)-1], "/")
	png, err := getPagePng("./files"+path, pageNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(png))
}

func getPageNum(filepath string) (int, error) {
	doc, err := fitz.New(filepath)
	if err != nil {
		return 0, err
	}
	defer doc.Close()
	return doc.NumPage(), nil
}

func getPagePng(filepath string, pageNumber int) ([]byte, error) {
	doc, err := fitz.New(filepath)
	if err != nil {
		return nil, err
	}
	defer doc.Close()

	png, err := doc.ImagePNG(pageNumber-1, 300)
	return png, err
}
