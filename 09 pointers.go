package main

import "fmt"

func main() {
	mostrarPointers()
	// pointersYFunciones()
}

/*
Todos los valores que definimos en nuestro programa (o código) están almacenados en la memoria de la computadora y tiene su
propia dirección única en memoria.
Los pointers (punteros) son variables especiales que mantienen las direcciones en la memoria de los valores.
En Go, declaramos los pointers usando *

Ejemplo: var p *int (Ahora, p is un puntero a un valore int)

Los pointers permiten pasar referencias a valores en tu programa.

Ahora sabemos como definir un puntero, pero ahora ¿Como lo asignamos a una dirección de memoria?
Esto se hace usando el operador &, el cual devuelve la dirección de memoria de una variable.

Ejemplo:
x := 42
p := &x

Ahora p es un puntero y mantiene la dirección de memoria de x. Veamos su valor.

*/

func mostrarPointers() {
	x := 42
	p := &x

	fmt.Println(p)
	fmt.Println(*p) // Sí queremos acceder al underlying value de un puntero, podemos usar el operador *.

	*p = 8
	fmt.Println(*p) // El operador * también puede ser usado para cambiar el valor del puntero que mantiene la dirección de memoria.
	fmt.Println(x)  // Hemos cambiado el valor de x usando el puntero p.

	fmt.Println("-----------------------------------------")

	// Ejercicio
	a := 4
	b := &a
	a += 2
	*b = *b - 1
	fmt.Println(a) // 5
}

// func pointersYFunciones() {
// 	func change(val int)  { // Función que toma una argumento de tipo int y cambia su valor.
// 		val = 8
// 	}
// 	func change_ptr(ptr *int)  { // Función que toma una argumento de tipo pointer y cambia su valor.
// 		*ptr = 8
// 	}

// 	x := 42

// 	change(x)
// 	fmt.Println(x)

// 	change_ptr(&x)
// 	fmt.Println(x)
// }
