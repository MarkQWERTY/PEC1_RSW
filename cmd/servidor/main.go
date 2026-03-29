package main

import (
	"log"
	"net/http"
	"pec2/internal/handlers"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/css/", fs)
	http.Handle("/img/", fs)
	http.Handle("/scss/", fs)

	http.HandleFunc("/registro", handlers.RegistroHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/css/") || strings.HasPrefix(r.URL.Path, "/img/") || strings.HasPrefix(r.URL.Path, "/scss/") {
			fs.ServeHTTP(w, r)
			return
		}
		handlers.PageHandler(w, r)
	})

	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error iniciando servidor: ", err)
	}
}
