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

	imprimeSlice()
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

	//
	var matriz [2][2]string

	matriz[0][0] = "Hola"
	matriz[0][1] = "Mundo"
	matriz[1][0] = "Cruso"
	matriz[1][1] = "Go"

	fmt.Println(matriz)

	matriz2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matriz2)
}

func imprimeSlice() {
	var slice1 []string

	slice1 = append(slice1, "Hola", "Slice")

	fmt.Println(slice1)

	fmt.Println(len(slice1))

	slice1 = append(slice1, "NuevoElemento")
	fmt.Println(slice1)
	fmt.Println(len(slice1))

	// Matrices

	matriz := [][]string{{"Hola", "Mundo"}, {"Curso", "Go"}}

	fmt.Println(matriz)

	var matriz2 [][]string

	row1 := []string{"Hola2", "slideMatriz2"}
	row2 := []string{"Hola3", "Go2"}

	matriz2 = append(matriz2, row1)
	matriz2 = append(matriz2, row2)

	fmt.Println(matriz2)
}
