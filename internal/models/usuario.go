package models

type Usuario struct {
<<<<<<< HEAD
	ID              int    `json:"id"`
=======
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
	Nombre          string `json:"nombre"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Direccion       string `json:"direccion"`
	Telefono        string `json:"telefono"`
	Correo          string `json:"correo"`
	Documento       string `json:"documento"`
	MetodoPago      string `json:"metodo_pago"`
	NumeroPago      string `json:"numero_pago"`
	Password        string `json:"password"`
}
