package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.1416
	const pi2 = 3.14
	fmt.Println("pi:", pi)
	fmt.Println("p2:", pi2)

	//Declaracion de variables enteras solo 3 formas.
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base", base)
	fmt.Println("altura", altura)
	fmt.Println("area", area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

	//Suma
	var x int = 12
	y := 14

	resultado := x + y
	fmt.Println("Resultado:", resultado)

	//Multiplicacion
	resultado = x * y
	fmt.Println("Multiplicacion", resultado)

	//division
	resultado = x / y
	fmt.Println("Division", resultado)

}
