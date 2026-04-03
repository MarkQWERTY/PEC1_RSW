package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"pec2/internal/db"
	"pec2/internal/models"
	"strconv"
	"time"
)

// Devuelve el socio logueado o nil si no hay cookie/socio
func obtenerSocioLogueado(r *http.Request) *models.Socio {
	cookie, err := r.Cookie("session_user")
	if err != nil {
		return nil
	}
	return db.ObtenerSocioPorNombre(cookie.Value)
}

func ReservasHandler(w http.ResponseWriter, r *http.Request) {
	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmplFile := filepath.Join("web", "templates", "reservas.html")
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, "Error cargando vista de reservas", http.StatusInternalServerError)
		return
	}

	misReservas := db.ObtenerReservasDeSocio(socio.ID)
	// Para facilitar la visualización en la plantilla, pasamos a un map o estructuramos mejor
	type ReservaDetalle struct {
		models.Reserva
		NombreClase string
	}
	
	var misReservasDetalle []ReservaDetalle
	for _, res := range misReservas {
		var nombre string
		for _, c := range db.ClasesLista {
			if c.ID == res.ActividadID {
				nombre = c.NombreClase
				break
			}
		}
		misReservasDetalle = append(misReservasDetalle, ReservaDetalle{
			Reserva:     res,
			NombreClase: nombre,
		})
	}

	data := map[string]interface{}{
		"Socio":       socio,
		"Clases":      db.ClasesLista,
		"MisReservas": misReservasDetalle,
	}

	tmpl.Execute(w, data)
}

func ProcesarReservaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/reservas", http.StatusSeeOther)
		return
	}

	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	claseIDStr := r.FormValue("clase_id")
	claseID, _ := strconv.Atoi(claseIDStr)
	
	// Usamos la fecha de mañana por defecto para simplificar
	fecha := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	
	db.CrearReserva(socio.ID, claseID, fecha)

	http.Redirect(w, r, "/reservas", http.StatusSeeOther)
}
