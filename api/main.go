package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Persona struct {
	Nombre    string
	Apellidos string
}

func handlerRaiz(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hola Mundo API")
}

func handlerUsuarios(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hola usuarios!")
}

func handlerPersonas(response http.ResponseWriter, request *http.Request) {

	persona := Persona{"Esteban", "Paredez"}

	jsResponde, err := json.Marshal(persona)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")

	response.Write(jsResponde)
}

func main() {

	http.HandleFunc("/", handlerRaiz)
	http.HandleFunc("/usuarios", handlerUsuarios)
	http.HandleFunc("/personas", handlerPersonas)

	http.ListenAndServe(":8000", nil)
}
