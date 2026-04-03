package services

import (
	"pec2/internal/db"
	"pec2/internal/models"
)

func RegistrarUsuario(u models.Usuario) error {
	return db.GuardarUsuario(u)
}
