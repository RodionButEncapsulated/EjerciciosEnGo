package main

import (
	"fmt"
)

/*
La única construcción de bucle en Go es for, que tiene tres componentes: init, condition, post statement.
Por Ejemplo:

for i:=0; i<5; i++{
	fmt.Println("i")
}

El INIT y el POST pueden ser omitidos, por ejemplo:

sum := 1
res := 0
for sum <= 1000 {
	res += sum
	sum++
}
fmt.Println(res)


*/

func main() { // Llamando la función principal
	fmt.Println("--------- Ejemplo 1 ---------")
	loopForEjemplo1()
	fmt.Println("--------- Ejemplo 2 ---------")
	loopForEjemplo2()
}

func loopForEjemplo1() {
	for i := 0; i < 5; i++ { // El INIT declara desde donde comienza el ciclo (0), el CONDITION indica cuando acaba (cuando sea mayor a 5) y el POST indica la función a hacer en cada ciclo (aumentar el valor de i)
		fmt.Println(i)
	}
}

func loopForEjemplo2() { // Tenga en cuenta que con solo la condición, for se vuelve similar a los bucles while disponibles en otros lenguajes de programación.
	sum := 1
	res := 0
	for sum <= 1000 {
		res += sum
		sum++
	}
	fmt.Println(res)
}
