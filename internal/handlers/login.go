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
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
<<<<<<< HEAD
		RenderTemplate(w, "login", nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		correo := r.FormValue("correo")
		contrasena := r.FormValue("contrasena")

		socio := db.ObtenerSocioPorCorreo(correo)
		if socio != nil && socio.Contrasena == contrasena {
			// Login exitoso
			http.SetCookie(w, &http.Cookie{
				Name:    "session_user",
				Value:   correo,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			})
			http.SetCookie(w, &http.Cookie{
				Name:    "session_nombre",
=======
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
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
				Value:   socio.Nombre,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			})
			http.Redirect(w, r, "/reservas", http.StatusSeeOther)
			return
		}

		// Login fallido
<<<<<<< HEAD
		RenderTemplate(w, "login", map[string]interface{}{
			"ErrorMsg": "Credenciales inválidas. Por favor, inténtelo de nuevo.",
=======
		tmplFile := filepath.Join("web", "templates", "login.html")
		tmpl, _ := template.ParseFiles(tmplFile)
		tmpl.Execute(w, map[string]interface{}{
			"ErrorMsg": "Usuario o contraseña incorrectos",
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
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
<<<<<<< HEAD
	http.SetCookie(w, &http.Cookie{
		Name:    "session_nombre",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
=======
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
