package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pec2/internal/models"
)

var dbFile = filepath.Join("internal", "db", "datos_registro.json")
var resenasFile = filepath.Join("internal", "db", "datos_resenas.json")

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

func ObtenerResenas() []models.Resena {
	var resenas []models.Resena
	file, err := os.ReadFile(resenasFile)
	if err == nil {
		json.Unmarshal(file, &resenas)
	} else {
		// Mock data if file doesn't exist
		resenas = []models.Resena{
			{Autor: "Carlos G.", Puntuacion: 5, Texto: "El mejor gimnasio de la ciudad. Las máquinas son increíbles."},
			{Autor: "María P.", Puntuacion: 4, Texto: "Muy buen ambiente, aunque a veces hay mucha gente en hora punta."},
			{Autor: "Luis R.", Puntuacion: 5, Texto: "Los fisioterapeutas me curaron una lesión de hombro que llevaba meses arrastrando."},
		}
	}
	return resenas
}

func GuardarResena(r models.Resena) error {
	resenas := ObtenerResenas()
	resenas = append(resenas, r)

	data, err := json.MarshalIndent(resenas, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(resenasFile, data, 0644)
}
