package services

import "pec2/internal/models"

type CalculadoraService struct{}

func (s *CalculadoraService) CalcularFFMI(datos *models.DatosFFMI) float64 {
	if !datos.ValidarDatos() {
		return 0.0
	}

	masaMagra := datos.Peso * (1.0 - (datos.IndiceGrasaCorporal / 100.0))
	ffmi := masaMagra / (datos.Altura * datos.Altura)

	// FFMI normalizado
	ffmiNormalizado := ffmi + 6.1*(1.8-datos.Altura)
	
	datos.Resultado = ffmiNormalizado
	return ffmiNormalizado
}

