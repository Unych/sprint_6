package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	formFileName = "index.html"
	fileField    = "myFile"
	maxFormSize  = 10 << 20
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, formFileName)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxFormSize); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile(fileField)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := fmt.Sprintf("%d%s", time.Now().UTC().UnixNano(), filepath.Ext(header.Filename))
	if err = os.WriteFile(name, []byte(result), 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(result))
}
