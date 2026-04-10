package handlers

import (
<<<<<<< HEAD
	"net/http"
=======
	"html/template"
	"net/http"
	"path/filepath"
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
	"pec2/internal/db"
	"strconv"
	"strings"
)

type ResultadoBusqueda struct {
	Titulo string
	Tipo   string
	URL    string
}

func BuscadorHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var resultados []ResultadoBusqueda

	if query != "" {
		for _, m := range db.Maquinas {
			if strings.Contains(strings.ToLower(m.Nombre), query) || strings.Contains(strings.ToLower(m.Descripcion), query) || strings.Contains(strings.ToLower(m.Marca), query) {
				resultados = append(resultados, ResultadoBusqueda{
					Titulo: m.Nombre + " (" + m.Marca + ")",
					Tipo:   "Maquinaria",
					URL:    "/maquinaria-detalle?id=" + strconv.Itoa(m.ID),
				})
			}
		}
		for _, s := range db.Servicios {
			if strings.Contains(strings.ToLower(s.Nombre), query) || strings.Contains(strings.ToLower(s.DescripcionBreve), query) || strings.Contains(strings.ToLower(s.DescripcionLarga), query) {
				resultados = append(resultados, ResultadoBusqueda{
					Titulo: s.Nombre,
					Tipo:   "Servicio",
					URL:    "/servicio-detalle?id=" + strconv.Itoa(s.ID),
				})
			}
		}
	}

<<<<<<< HEAD
	RenderTemplate(w, "buscar", map[string]interface{}{
=======
	tmplFile := filepath.Join("web", "templates", "buscar.html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.Error(w, "Error cargando página de resultados", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
		"Query":      r.URL.Query().Get("q"),
		"Resultados": resultados,
	})
}
