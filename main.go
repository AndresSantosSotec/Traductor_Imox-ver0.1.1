// main.go

package main

import (
	"log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/agregarpalabra", agregarpalabra)
	log.Println("Servidor corriendo ")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func agregarpalabra(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "agregarpalabra", nil) // Corregir el nombre de la plantilla aquí
}
