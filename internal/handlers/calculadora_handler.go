package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"pec2/internal/models"
	"pec2/internal/services"
	"strconv"
)

func CalculadoraHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		MostrarFormulario(w, r)
	} else if r.Method == http.MethodPost {
		ProcesarCalculadora(w, r)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	tmplFile := filepath.Join("web", "templates", "calculadora.html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.Error(w, "Error cargando vista de calculadora", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func ProcesarCalculadora(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	alturaStr := r.FormValue("altura")
	pesoStr := r.FormValue("peso")
	grasaStr := r.FormValue("grasa")

	altura, _ := strconv.ParseFloat(alturaStr, 64)
	peso, _ := strconv.ParseFloat(pesoStr, 64)
	grasa, _ := strconv.ParseFloat(grasaStr, 64)

	datos := models.DatosFFMI{
		Altura:              altura,
		Peso:                peso,
		IndiceGrasaCorporal: grasa,
	}

	var resultado float64
	var errorMsg string

	if datos.ValidarDatos() {
		calcService := services.CalculadoraService{}
		resultado = calcService.CalcularFFMI(&datos)
	} else {
		errorMsg = "Por favor, introduce datos válidos."
	}

	tmplFile := filepath.Join("web", "templates", "calculadora.html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.Error(w, "Error cargando vista de calculadora", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Resultado": resultado,
		"Datos":     datos,
		"ErrorMsg":  errorMsg,
	})
}
