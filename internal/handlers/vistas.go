package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[1:]
	if page == "" || page == "index.html" {
		page = "index"
	}

	// Quitar .html si el usuario o un enlace navega directamente con la extensión
	page = strings.TrimSuffix(page, ".html")

	if strings.Contains(page, ".") {
		http.NotFound(w, r)
		return
	}

	tmplFile := filepath.Join("web", "templates", page+".html")
	
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl.Execute(w, nil)
}
