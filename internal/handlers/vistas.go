package handlers

import (
<<<<<<< HEAD
	"log"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"strconv"
	"strings"
)

func GetWebPath(relative ...string) string {
	webBase := "web"
	// Si no existe ./web, probar ../../web (si corremos desde cmd/servidor)
	if _, err := http.Dir(webBase).Open("."); err != nil {
		webBase = "../../web"
	}
	parts := append([]string{webBase}, relative...)
	return filepath.Join(parts...)
}

=======
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"pec2/internal/db"
)

>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
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

<<<<<<< HEAD
=======
	tmplFile := filepath.Join("web", "templates", page+".html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
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
<<<<<<< HEAD
	case "equipo":
		data = db.EquipoCompleto
=======
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
	case "servicio-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, s := range db.Servicios {
			if s.ID == id {
				data = s
				break
			}
		}
<<<<<<< HEAD
	case "tienda/tramitar":
		// Si cae aquí es que no se ha registrado el handler específico o falta reiniciar
		cookie, err := r.Cookie("session_user")
		if err == nil {
			usuario := db.ObtenerUsuarioPorCorreo(cookie.Value)
			if usuario != nil {
				data = map[string]interface{}{"Usuario": usuario}
				page = "tramitar-pedido"
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		} else {
			http.Redirect(w, r, "/login?return=/tienda/tramitar", http.StatusSeeOther)
			return
		}
	}

	RenderTemplate(w, page, data)
=======
	}

	tmpl.Execute(w, data)
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
}
