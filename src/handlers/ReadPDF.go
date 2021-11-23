package handlers

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func ReadPDF(w http.ResponseWriter, r *http.Request) {
	var JEPG_OPTIONS = &jpeg.Options{Quality: jpeg.DefaultQuality}
	var TMP_DIR, _ = ioutil.TempDir(os.TempDir(), "fitz")
	var tmpFileName = func(i int) string {
		return filepath.Join(TMP_DIR, fmt.Sprintf("test%02d.jpg", i))
	}

	doc, err := fitz.New("./files" + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result [][]byte

	for i := 0; i < doc.NumPage(); i++ {
		img, _ := doc.Image(i)
		f, _ := os.Create(tmpFileName(i))

		jpeg.Encode(f, img, JEPG_OPTIONS)
		f.Close()

		contents, _ := ioutil.ReadFile(tmpFileName(i))
		result = append(result, contents)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
