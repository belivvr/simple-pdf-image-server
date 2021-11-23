package handlers

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"net/http"

	"github.com/thoas/go-funk"
)

func ReadDir(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("files" + r.URL.Path)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	names := funk.Map(files, func(file fs.FileInfo) string {
		return file.Name()
	}).([]string)

	json.NewEncoder(w).Encode(names)
}
