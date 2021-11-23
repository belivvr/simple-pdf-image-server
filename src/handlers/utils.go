package handlers

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func docToImageList(doc *fitz.Document) (result [][]byte) {
	var JEPG_OPTIONS = &jpeg.Options{Quality: jpeg.DefaultQuality}
	var TMP_DIR, _ = ioutil.TempDir(os.TempDir(), "fitz")
	var tmpFileName = func(i int) string {
		return filepath.Join(TMP_DIR, fmt.Sprintf("test%02d.jpg", i))
	}

	for i := 0; i < doc.NumPage(); i++ {
		img, _ := doc.Image(i)
		f, _ := os.Create(tmpFileName(i))

		jpeg.Encode(f, img, JEPG_OPTIONS)
		f.Close()

		contents, _ := ioutil.ReadFile(tmpFileName(i))
		result = append(result, contents)
	}
	return result
}
