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
	svg, err := getPageSvg("./files"+path, pageNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(svg))
}

func getPageNum(filepath string) (int, error) {
	doc, err := fitz.New(filepath)
	if err != nil {
		return 0, err
	}
	defer doc.Close()
	return doc.NumPage(), nil
}

func getPageSvg(filepath string, pageNumber int) (string, error) {
	doc, err := fitz.New(filepath)
	if err != nil {
		return "", err
	}
	defer doc.Close()

	svg, err := doc.SVG(pageNumber - 1)
	return svg, err
}
