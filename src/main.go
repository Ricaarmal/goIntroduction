package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	returnValor := a * 2
	return returnValor
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	//Funciones en impresión

	normalFunction("Hola Mundo")
	tripleArgument(2, 3, "Hola")
	fmt.Printf("Valor: %d \n", returnValue(4))
	value1, value2 := doubleReturn(4)
	fmt.Println("Value1 y Value 2: ", value1, value2)
	// En caso de necesitar solo un valor de la función se declara
	// con "_" guion bajo, el acomodo es importante.

	/*
		helloMessage := "Hello"
		worldMessage := "World"

		//salto de linea cualquiera
		fmt.Println(helloMessage, worldMessage)

		// Printf Tipo de dato
		nombre := "Platzi"
		cursos := 500

		fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
		fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

		//Sprintf - Imprime de inmediato variable
		message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
		fmt.Println(message)

		//Conocer cual es el tipo de dato

		fmt.Printf("helloMessage: %T \n", helloMessage)
		fmt.Printf("cursos: %T \n", cursos)

		 Variables y tipo de constantes.
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
	*/

}
