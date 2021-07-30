/*	El paquete "fmt" también le permite recibir información del usuario del programa.
	Para tomar la entrada, necesitamos usar la función Scanln () y proporcionarle la
	variable que debe contener el valor de entrada:
*/

package main

import "fmt"

func main() {
	// Ejercicio 1
	devolverStringInput()

	// Ejercicio 2
	devolverIntInput()
}

//----------------------------------------------------------------------------------

// Ejercicio 1 (Devolver un valor string ingresado por input)
func devolverStringInput() {
	var entrada string

	fmt.Print("[?] Ingrese un dato de tipo texto: ")
	fmt.Scanln(&entrada)
	fmt.Println("[!] Usted ingresó: " + entrada)
	// fmt.Printf("[!] Usted ingresó: %d\n", entrada) // Una forma de pasar variables con formato (f en python)
}

// Ejercicio 2 (Devolver un valor int multiplicado por 2 que ha sido ingresado por input)
func devolverIntInput() {
	var entrada int

	fmt.Println("")
	fmt.Print("[?] Ingrese un dato de tipo entero (Número): ")
	fmt.Scanln(&entrada)
	fmt.Println("[!] El doble del número", entrada, "es: ", entrada*2)

}

/*
	[!] Tenga en cuenta el & antes del nombre de la variable:
	Se utiliza para devolver la dirección de una variable.

	Si necesitamos tomar un número como entrada, como un entero,
	simplemente podemos declarar el tipo de la variable de entrada como int
	y Go convertirá automáticamente la entrada a ese tipo.
*/
