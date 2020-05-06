package structs

import "fmt"

type Esteban string

type Persona struct {
	Nombres            string
	Apellidos          string
	DocumentoIdentidad string
	Telefonos          []string
	Direccion          string
	Edad               int
}

type Casa struct {
	NumeroCasa int
	Personas   []Persona
}

func SoyDefer() {
	fmt.Println("Defer, esto quiero que se ejecute al final")
}

func (c Casa) GetNumeroCasa() {
	fmt.Println("El Numero de la casa es: ", c.NumeroCasa)
}
func (c Casa) GetPersonaCasa() {
	defer SoyDefer()
	fmt.Println("Las personas de la casa son: ", c.Personas)
}
