package main

import "fmt"

const holaMundo string = "Hola %s %s, bienvenido al curso de Go"

func main() {

	name := getName()

	lastName := "Paredez"

	var miNumero = suma(40, 60)

	x, y, z := getMultiplesVariables()

	decimal32, decimal64 := getDecimal()

	fmt.Println("Hola ", name)

	fmt.Printf(holaMundo, name, lastName)

	fmt.Println("")

	fmt.Println(miNumero, x, y, z)

	fmt.Println(decimal32, decimal64)
}

func getName() string {

	var name string = "Nombre por defecto"

	fmt.Print("Ingresa tu nombre: ")

	fmt.Scanf("%s", &name)

	return name

}

func getMultiplesVariables() (int, int32, int64) {
	return 1, 225252525, 252525252242352525
}

func suma(a int, b int) int {
	return a + b
}

func getDecimal() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
