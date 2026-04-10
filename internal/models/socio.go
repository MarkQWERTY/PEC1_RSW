package models

type Socio struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Contrasena        string `json:"contrasena"`
	SuscripcionActiva bool   `json:"suscripcion_activa"`
}

type Clases struct {
	ID          int    `json:"id"`
	NombreClase string `json:"nombre_clase"`
	Entrenador  string `json:"entrenador"`
	Aforo       int    `json:"aforo"`
	Horario     string `json:"horario"`
<<<<<<< HEAD
	Descripcion string `json:"descripcion"`
	Lugar       string `json:"lugar"`
=======
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
}

type Reserva struct {
	ID          int    `json:"id"`
	SocioID     int    `json:"socio_id"`
	ActividadID int    `json:"actividad_id"`
	FechaAsist  string `json:"fecha_asist"`
}
<<<<<<< HEAD

type MiembroEquipo struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Cargo       string `json:"cargo"`
	Descripcion string `json:"descripcion"`
	Imagen      string `json:"imagen"`
}
=======
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
