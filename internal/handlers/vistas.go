package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"pec2/internal/db"
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
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
	log.Printf("Usuario accedió a la página: %s", page)

	var data interface{}

	switch page {
	case "index":
		data = db.ObtenerResenas()
	case "maquinaria":
		data = db.Maquinas
	case "maquinaria-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, m := range db.Maquinas {
			if m.ID == id {
				data = m
				break
			}
		}
	case "servicios":
		data = db.Servicios
	case "servicio-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, s := range db.Servicios {
			if s.ID == id {
				data = s
				break
			}
		}
	}

	tmpl.Execute(w, data)
}
