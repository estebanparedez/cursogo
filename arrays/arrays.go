package arrays

import "fmt"

// Imprime array muestra como se trabaja com matrices en Go
func ImprimeArray() {

	var array1 [2]string
	array1[0] = "Hola"
	array1[1] = "Mundo"

	array2 := [4]int{1, 2, 3, 4}

	fmt.Println(array1)
	fmt.Println(len(array1))

	fmt.Println(array2)

	// Desde aqui

	var matriz [2][2]string
	matriz[0][0] = "Hola"
	matriz[0][1] = "Mundo"
	matriz[1][0] = "Cruso"
	matriz[1][1] = "Go"

	fmt.Println(matriz)

	matriz2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matriz2)
}

// Imprime array muestra como se trabaja con Slices en Go
func ImprimeArray() {

	var slice1 []string

	slice1 = append(slice1, "Hola", "Slice")

	fmt.Println(slice1)
	fmt.Println(len(slice1))

	slice1 = append(slice1, "NuevoElemento")

	fmt.Println(slice1)
	fmt.Println(len(slice1))

	matriz := [][]string{{"Hola", "Mundo"}, {"Curso", "Go"}}

	fmt.Println(matriz)

	var matriz2 [][]string

	row1 := []string{"Hola2", "slideMatriz2"}
	row2 := []string{"Hola3", "Go2"}

	matriz2 = append(matriz2, row1)
	matriz2 = append(matriz2, row2)

	fmt.Println(matriz2)
}
