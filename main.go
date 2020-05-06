package main

import (
	"fmt"

	"github.com/cursogo/structs"
	/*"github.com/cursogo/maps"
	"github.com/cursogo/arrays"
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

	operacionesstrings.OperacionesConStrings() 

	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetKeyMap("Esteban"))*/

	var miVariable structs.Esteban = "Mi propio tipo de dato"

	fmt.Println(miVariable)

	esteban := structs.Persona{
		Nombre:             "Esteban",
		Apellidos:          "Paredez",
		DocumentoIdentidad: "XXX",
		Telefono:           []string{"111", "334"},
		Direccion:          "direcc",
		Edad:               30,
	}

	fmt.Println(esteban)

	maria := structs.Persona {
		Nombre:             "Maria",
		Apellidos:          "Alvarado",
		DocumentoIdentidad: "YYY",
		Telefono:           []string{"222", "555"},
		Direccion:          "ZZZZ",
		Edad:               33,
	}

	fmt.Println(maria)

	jorge := new(structs.Persona)
	jorge.Nombre = "Jorge"
	jorge.Apellidos= "Apellidos"
	jorge.DocumentoIdentidad= "MMM"
	jorge.Telefono=[]string{"222", "6666"}
	jorge.Direccion= "AAAA"
	jorge.Edad=40

	fmt.Println(jorge)

	casa := structs.Casa {
		NumeroCasa:  : 1,
		Personas : []structs.Persona{antonio, maria, *jorge},
	}

	fmt.Println()

	casa.GetNumeroCasa()

	casa.GetPersonaCasa()

}
