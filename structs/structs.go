package structs

import "fmt"

type Esteban string

type Persona struct {
	Nombre             string
	Apellidos          string
	DocumentoIdentidad string
	Telefono           []string
	Direccion          string
	Edad               int
}

type Casa struct {
	NumeroCasa int
	Personas   []Persona
}

func (c Casa) GetNumeroCasa() {
	fmt.Println("El Numero de la casa es: ", c.NumeroCasa)
}
func (c Casa) GetPersonaCasa() {
	fmt.Println("Las personas de la casa son: ", c.Personas)
}
