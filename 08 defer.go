package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // defer stackea los resultados de la función definida en la misma linea hasta que todas las demás funciones son ejecutadas.
	}
	fmt.Println("end")
}
