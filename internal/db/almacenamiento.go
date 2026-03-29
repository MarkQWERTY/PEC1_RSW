package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pec2/internal/models"
)

var dbFile = filepath.Join("internal", "db", "datos_registro.json")

func GuardarUsuario(u models.Usuario) error {
	var usuarios []models.Usuario

	file, err := os.ReadFile(dbFile)
	if err == nil {
		json.Unmarshal(file, &usuarios)
	}

	usuarios = append(usuarios, u)

	data, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbFile, data, 0644)
}
