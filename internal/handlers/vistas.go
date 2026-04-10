package handlers

import (
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
	case "equipo":
		data = db.EquipoCompleto
	case "servicio-detalle":
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for _, s := range db.Servicios {
			if s.ID == id {
				data = s
				break
			}
		}
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
}
