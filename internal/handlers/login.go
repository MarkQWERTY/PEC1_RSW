package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmplFile := filepath.Join("web", "templates", "login.html")
		tmpl, err := template.ParseFiles(tmplFile)
		if err != nil {
			http.Error(w, "Error cargando vista de login", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		nombre := r.FormValue("usuario")
		contrasena := r.FormValue("contrasena")

		socio := db.ObtenerSocioPorNombre(nombre)
		if socio != nil && socio.Contrasena == contrasena {
			// Login exitoso, establecer cookie
			http.SetCookie(w, &http.Cookie{
				Name:    "session_user",
				Value:   socio.Nombre,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			})
			http.Redirect(w, r, "/reservas", http.StatusSeeOther)
			return
		}

		// Login fallido
		tmplFile := filepath.Join("web", "templates", "login.html")
		tmpl, _ := template.ParseFiles(tmplFile)
		tmpl.Execute(w, map[string]interface{}{
			"ErrorMsg": "Usuario o contraseña incorrectos",
		})
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_user",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
