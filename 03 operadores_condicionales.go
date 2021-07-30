/*
Los operadores relacionales se utilizan para comparar dos valores y devolver un bool como resultado:
verdadero cuando se cumple la condición de comparación y falso cuando no.
*/

package main

import "fmt"

func main() {

	// Definiendo los valores de las variables
	x := 42
	y := 8

	//-----------------------------------------------------

	// OPERADORES CONDICIONALES

	// Igual a
	fmt.Println(x == y)

	// No igual a
	fmt.Println(x != y)

	// Mayor que
	fmt.Println(x > y)

	// Mayor o igual que
	fmt.Println(x >= y)

	// Menor que
	fmt.Println(x < y)

	// Menor o igual que
	fmt.Println(x <= y)

	//-----------------------------------------------------

	// OPERADORES LÓGICOS

	// Lógica AND (&&)
	fmt.Println(x != y && x <= y)

	// Lógica OR (||)
	fmt.Println(x != y || x <= y)

	// Lógica NOT (!)
	fmt.Println(!(x > y))

	// Puedes combinar varias condiciones y usar paréntesis para agruparlas, como: ((x> 0 && x <100) || x == 42)
}
