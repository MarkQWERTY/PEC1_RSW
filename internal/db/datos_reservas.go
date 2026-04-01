package db

import "pec2/internal/models"

var Socios = []models.Socio{
	{
		ID:                1,
		Nombre:            "testuser",
		Contrasena:        "1234",
		SuscripcionActiva: true,
	},
}

var ClasesLista = []models.Clases{
	{ID: 1, NombreClase: "Body Pump", Entrenador: "Carlos", Aforo: 20, Horario: "18:00 - 19:00"},
	{ID: 2, NombreClase: "Spinning", Entrenador: "María", Aforo: 15, Horario: "19:00 - 20:00"},
	{ID: 3, NombreClase: "Yoga Integral", Entrenador: "Ana", Aforo: 12, Horario: "09:00 - 10:00"},
	{ID: 4, NombreClase: "CrossFit WOD", Entrenador: "Jorge", Aforo: 10, Horario: "20:00 - 21:00"},
}

var ReservasActuales = []models.Reserva{}

func ObtenerSocioPorNombre(nombre string) *models.Socio {
	for _, s := range Socios {
		if s.Nombre == nombre {
			return &s
		}
	}
	return nil
}

func CrearReserva(socioID int, actividadID int, fecha string) bool {
	// Verificar aforo simple (como simulación)
	var count int
	for _, r := range ReservasActuales {
		if r.ActividadID == actividadID && r.FechaAsist == fecha {
			count++
		}
	}
	
	for _, c := range ClasesLista {
		if c.ID == actividadID {
			if count >= c.Aforo {
				return false // Aforo completo
			}
			break
		}
	}

	nuevaReserva := models.Reserva{
		ID:          len(ReservasActuales) + 1,
		SocioID:     socioID,
		ActividadID: actividadID,
		FechaAsist:  fecha,
	}
	ReservasActuales = append(ReservasActuales, nuevaReserva)
	return true
}

func ObtenerReservasDeSocio(socioID int) []models.Reserva {
	var result []models.Reserva
	for _, r := range ReservasActuales {
		if r.SocioID == socioID {
			result = append(result, r)
		}
	}
	return result
}
