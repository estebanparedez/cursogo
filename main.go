package main

import "fmt"

const holaMundo string = "Hola %s %s, bienvenido al curso de Go"

func main() {

	name := getName()

	lastName := "Paredez"

	var miNumero = suma(40, 60)

	x, y, z := getMultiplesVariables()

	fmt.Println("Hola ", name)

	fmt.Printf(holaMundo, name, lastName)

	fmt.Println("")

	fmt.Println(miNumero, x, y, z)
}

func getName() string {

	var name string = "Nombre por defecto"
}
