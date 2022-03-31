package chat

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)



type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		target_file_path := filepath.Join("templates", t.filename)
		t.templ = template.Must(template.ParseFiles(target_file_path))
	})

	err := t.templ.Execute(w, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
	}
}