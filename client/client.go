package client

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
)

type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

var box *packr.Box

func init() {
	box = packr.New("ui", "./out")
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	path = filepath.Join(h.StaticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(box).ServeHTTP(w, r)
}
