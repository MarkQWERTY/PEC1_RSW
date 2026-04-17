package handlers

import (
	"net/http"
	"pec2/internal/db"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	returnPath := r.URL.Query().Get("return")

	if r.Method == http.MethodGet {
		RenderTemplate(w, "login", map[string]interface{}{
			"ReturnPath": returnPath,
		})
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		correo := r.FormValue("correo")
		contrasena := r.FormValue("contrasena")
		returnPath = r.FormValue("return") // Leer desde el input hidden

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

			// Redirigir al sitio original si existe, si no a /reservas
			if returnPath != "" {
				http.Redirect(w, r, returnPath, http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/reservas", http.StatusSeeOther)
			}
			return
		}

		// Login fallido
		RenderTemplate(w, "login", map[string]interface{}{
			"ErrorMsg":   "Credenciales inválidas. Por favor, inténtelo de nuevo.",
			"ReturnPath": returnPath,
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

