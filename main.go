package main

import "fmt"

const holaMundo string = "Hola %s %s, bienvenido al curso de Go"

func main() {

	name := getName()

	lastName := "Paredez"

	var miNumero = suma(40, 60)

	x, y, z := getMultiplesVariables()

	decimal32, decimal64 := getDecimal()

	hola := "Hola"

	fmt.Println("Hola ", name)

	fmt.Printf(holaMundo, name, lastName)

	fmt.Println("")

	fmt.Println(miNumero, x, y, z)

	fmt.Println(decimal32, decimal64)

	fmt.Println(getUnicode())

	fmt.Println(string(hola[0]))

	fmt.Println(len(hola))

	imprimeArray()
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

func getUnicode() string {
	return "你好，世界"
}

func imprimeArray() {
	var array1 [2]string
	array1[0] = "Hola"
	array1[1] = "Mundo"

	array2 := [4]int{1, 2, 3, 4}

	fmt.Println(array1)
	fmt.Println(len(array1))

	fmt.Println(array2)

	var matriz [2][2]string
}
