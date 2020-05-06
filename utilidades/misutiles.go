package utilidades

import (
	"errors"
	"fmt"
)

func GetName() string {

	var name string = "Nombre por defecto"

	fmt.Print("Ingresa tu nombre: ")

	fmt.Scanf("%s", &name)

	return name

}

func GetMultiplesVariables() (int, int32, int64) {
	return 1, 225252525, 252525252242352525
}

func Suma(a interface{}, b interface{}) (int, error) {

	switch a.(type) {
	case string:
		return 0, errors.New("El valor A NO es un entero")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B NO es un entero")
	}

	return a.(int) + b.(int), nil
}

func GetDecimal() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func GetUnicode() string {
	return "你好，世界"
}

func GetNumero() int {
	var numero = 0

	fmt.Println("ingresa un numero")
	fmt.Scanf("%d", &numero)

	return numero
}

func Multiplo5() {

	if numero := GetNumero(); numero%5 == 0 {
		fmt.Println("Es multiplo de 5")
	} else if numero := GetNumero(); numero%3 == 0 {
		fmt.Println("Es multiplo de 3")
	} else {
		fmt.Println("No es multiplo de 3 ni 5")
	}
}

func Multiplo5switch() {
	var numero = 0

	fmt.Println("Ingresa un numero:")
	fmt.Scanf("%d", &numero)

	switch numero {
	case 2:
		fmt.Println("El numero es 2")
	default:
		fmt.Println("El numero NO es 2")

	}

	switch numero {
	case numero%5 == 0:
		fmt.Println("Es multiplo de 5")
	case numero%3 == 0:
		fmt.Println("El numero es 3")
	default:
		fmt.Println("Es multiplo de 3 ni de 5")

	}

}

func Bucles() {

	for i := 0; i < 3; i++ {
		fmt.Println("Bucle For: ", i)
	}

	j := 50

	for j > 0 {
		j -= 10
		fmt.Println("While ", j)
	}

	k := 500

	for {
		k -= 1
		if k == 0 {
			fmt.Println("Termino el bucle infinito")
			break
		}
	}

}
