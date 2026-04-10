package handlers

import (
<<<<<<< HEAD
	"net/http"
=======
	"html/template"
	"net/http"
	"path/filepath"
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
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
<<<<<<< HEAD
	RenderTemplate(w, "calculadora", nil)
=======
	tmplFile := filepath.Join("web", "templates", "calculadora.html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.Error(w, "Error cargando vista de calculadora", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
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

<<<<<<< HEAD
	RenderTemplate(w, "calculadora", map[string]interface{}{
=======
	tmplFile := filepath.Join("web", "templates", "calculadora.html")
	navbarFile := filepath.Join("web", "templates", "navbar.html")
	tmpl, err := template.ParseFiles(tmplFile, navbarFile)
	if err != nil {
		http.Error(w, "Error cargando vista de calculadora", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
>>>>>>> d48f6ffdbdb90e0d503e476e6ffbce582ca54153
		"Resultado": resultado,
		"Datos":     datos,
		"ErrorMsg":  errorMsg,
	})
}
