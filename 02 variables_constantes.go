package main

import (
	"fmt"
)

func main() {
	variable := 1234567892423569742
	variable2 := 1232456768987574634
	fmt.Println("La suma de los números es: ")
	fmt.Println(variable * variable2)

	nombre := "Pablo "
	apellido := "Tapias"
	fmt.Println("El nombre completo es: " + nombre + apellido)

	// Tipos de valores para variables

	var xVar int = 42
	var yVar float32 = 1.37
	var nameVar string = "James"
	var onlineVar bool = true

	fmt.Println(nameVar)
	fmt.Println(xVar)
	fmt.Println(yVar)
	fmt.Println(onlineVar)

	// Tipos de valores para constantes

	const xCons int = 50
	const yCons float32 = 3.33
	const nameCons string = "Pablo"
	const onlineCons bool = false

	fmt.Println(nameCons)
	fmt.Println(xCons)
	fmt.Println(yCons)
	fmt.Println(onlineCons)

	/*
		Otra característica interesante de Go son los valores cero:
		Las variables que se declaran sin valor toman el valor cero de su tipo:
		0 para tipos numéricos, falso para el tipo booleano, "" para cadenas.
	*/
}
