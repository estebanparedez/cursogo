package main

import (
	"fmt"

	"github.com/cursogo/maps"
	/*"github.com/cursogo/arrays"
	"github.com/cursogo/operacionesstrings"
	"github.com/cursogo/utilidades"*/)

// const holaMundo string = "Hola %s %s, bienvenido al curso de Go"

func main() {

	/* name := utilidades.GetName() // Aqui usamos nuestra funcion

	hola := "Hola"

	lastName := "Paredez"

	var miNumero = utilidades.Suma(40, 60) // Aqui usamos nuestra funciomn

	decimal32, decimal64 := utilidades.GetDecimal()

	x, y, z := utilidades.GetMultiplesVariables()

	fmt.Printf(holaMundo, name, lastName)

	fmt.Println(miNumero, x, y, z)

	fmt.Println(decimal32, decimal64)

	fmt.Println(utilidades.GetUnicode())

	fmt.Println(string(hola[0]))

	fmt.Println(len(hola))

	arrays.ImprimeArray()

	arrays.ImprimeSlice()

	utilidades.Multiplo5()

	utilidades.Multiplo5switch()

	utilidades.Bucles()

	operacionesstrings.OperacionesConStrings() */

	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetKeyMap("Esteban"))

}
