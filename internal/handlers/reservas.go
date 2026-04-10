package handlers

import (
	"net/http"
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
	return db.ObtenerSocioPorCorreo(cookie.Value)
}

func ReservasHandler(w http.ResponseWriter, r *http.Request) {
	socio := obtenerSocioLogueado(r)
	if socio == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	misReservas := db.ObtenerReservasDeSocio(socio.ID)
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

	RenderTemplate(w, "reservas", map[string]interface{}{
		"Socio":       socio,
		"Clases":      db.ClasesLista,
		"MisReservas": misReservasDetalle,
	})
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
