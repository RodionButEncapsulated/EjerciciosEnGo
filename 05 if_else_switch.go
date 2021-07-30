package main // Declarando el Package

import ( // Importando los librerías a utilizar
	"fmt"
	"strconv" // Para convertir valores de tipo int en str.
)

func main() { // Llamando la función principal.

	revisarEdadConIfElse() // Llamando función para revisar la edad con if y else.
	revisarEdadConSwitch() // Llamando función para revisar la edad con switch.
}

/*
-----------------------------IF ELSE-----------------------------
Una declaración if else en Go es igual que en cualquier lenguaje común.
*/

func revisarEdadConIfElse() {
	var edad int // Declarando la variable edad como valor entero.

	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)              // Solicitando el valor de la variable edad.
	edad_str := strconv.Itoa(edad) // Creando una variable con el valor de la variable edad pero en tipo str.

	// NOTA: El código de las instrucciones if y else debe incluirse entre llaves {} y puede incluir varias líneas de código.
	if edad > 0 {
		if edad >= 18 {
			fmt.Println("Sí tienes " + edad_str + " eres mayor de edad!")
		} else {
			fmt.Println("Sí tienes " + edad_str + " eres menor de edad!")
		}
	} else {
		fmt.Println("[!] El valor ingresado debe ser numérico.")
	}
	/*	No hay ternario en Go, por lo que deberá usar una declaración if completa incluso para las condiciones básicas.
		NOTA: A veces necesita una variable solo para las declaraciones if / else.
		Para esto, la instrucción if en Go puede comenzar con una declaración de variable antes de la condición:

		if x := 42; x >= 18 {
			fmt.Println("Eres mayor de edad.")
		} else {
			fmt.Println("Eres menor de edad.")
		}

	*/
}

/*
-----------------------------SWITCH-----------------------------
Una declaración switch es una forma más corta de escribir una secuencia de declaraciones if / else.
*/

func revisarEdadConSwitch() {
	var edad int

	fmt.Println("Ingrese su edad: ")
	fmt.Scan(&edad)
	edad_str := strconv.Itoa(edad)

	switch {
	case edad >= 18:
		fmt.Println("Sí tienes " + edad_str + " eres mayor de edad.")
	case edad < 18:
		fmt.Println("Sí tienes " + edad_str + " eres menor de edad.")
	}
}

/*
En otros lenguajes de programación, como C ++ o Java, cada caso debe tener una instrucción break para detener la ejecución de la instrucción switch.
En Go, la declaración break no es necesaria, ya que los case de switch se evalúan de arriba a abajo, deteniéndose cuando un caso tiene éxito.
*/
