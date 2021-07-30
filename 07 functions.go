package main

import (
	"fmt"
)

/*
Al igual que en la mayoria de lenguajes, las variables declaradas dentro de una función son llamadas variables locales,
y las variables declaradas fuera de una función o en la función main, son llamadas variables globales.
En Go, se denomina como mala practica definir variables globales, es mejor usar variables locales y manejarlas a través de
argumentos en funciones y en return en funciones.
*/

// Funciones para saludar.
func solicitarNombre() string {
	var nombre string
	fmt.Print("[?] Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	return nombre
}

func saludar(nombre string) { // Definiendo la variable del argumento y su tipo de valor, en este caso String.
	fmt.Println("[!] Hola, " + nombre + ".")
}

// Funciones para sumar.
func solicitarNumero() int {
	var num int
	fmt.Print("[?] Ingrese un número: ")
	fmt.Scanln(&num)
	return num
}

func solicitarNumeros() (int, int) {
	var num1 int
	var num2 int
	fmt.Print("[?] Ingrese el primer número: ")
	fmt.Scanln(&num1)
	fmt.Print("[?] Ingrese el segundo número: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func sumar(num1, num2 int) { // Definiendo la variable del argumento y su tipo de valor, en este caso Int.
	fmt.Print("[!] La suma de los números es: ")
	fmt.Println(num1 + num2)
}

func main() { // Llamando la función principal.

	nombre := solicitarNombre()
	saludar(nombre) // Llamando la función con argumentos (nombre = lo que el ingrese usuario)

	num1 := solicitarNumero()
	num2 := solicitarNumero()
	sumar(num1, num2) // Llamando la función con argumentos (num1 y num2 = Lo que ingrese el usuario)

	num1, num2 = solicitarNumeros()
	sumar(num1, num2)

}
