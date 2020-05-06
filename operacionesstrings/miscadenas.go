package operacionesstrings

import (
	"fmt"
	"strings"
)

func OperacionesConStrings() {

	var texto = "Hola Go, Hola Antonio, Hola Mundo"

	var texto2 = "Cadena 2"

	fmt.Println(texto)

	fmt.Println(strings.ToUpper(texto))

	fmt.Println(strings.ToLower(texto))

	fmt.Println(strings.Replace(texto, "Hola", "Adios", -1))

	fmt.Println(strings.Compare(texto, texto2))

	fmt.Println(strings.Split(texto, ","))

	fmt.Println(strings.Contains(texto, "Hola"))

	fmt.Println(strings.Contains(texto, "Coche"))

}
