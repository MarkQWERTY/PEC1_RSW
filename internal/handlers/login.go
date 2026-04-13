package handlers

import (
	"net/http"
	"pec2/internal/db"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
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
				Value:   socio.Nombre,
				Expires: time.Now().Add(24 * time.Hour),
				Path:    "/",
			})
			http.Redirect(w, r, "/reservas", http.StatusSeeOther)
			return
		}

		// Login fallido
		RenderTemplate(w, "login", map[string]interface{}{
			"ErrorMsg": "Credenciales inválidas. Por favor, inténtelo de nuevo.",
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
	http.SetCookie(w, &http.Cookie{
		Name:    "session_nombre",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

