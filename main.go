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
	http.HandleFunc("/add", abrirAddTemplate)
	log.Println("Servidor corriendo en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	// Renderizar el formulario de inicio de sesión
	plantillas.ExecuteTemplate(w, "login", nil)
}

func agregarpalabra(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "agregarpalabra", nil)
}

func abrirAddTemplate(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "add", nil)
}
