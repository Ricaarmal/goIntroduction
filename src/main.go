package main

import (
	"fmt"
	"log"
	"strconv"
)

func esPar(x int) {
	if x%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("No es par")
	}
}

func correctAuth(val1, val2 string) {
	if val1 == "admin" && val2 == "1234" {
		fmt.Println("Bienvenido")
	} else {
		fmt.Println("Ni siquiera se quien eres!")
	}
}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// with And

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// with Or

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Value: ", value)
	}

	esPar(valor2)
	correctAuth("adminsa", "1234")

	//Switch en Go

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Switch sin condicion

	value1 := 5

	switch {
	case value1 > 100:
		fmt.Println("Es mayor a 100")
	case value1 < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No existe en la condición")
	}

	// KEYWORDS

	//defer bases de datos para terminar el proceso al final. Solo un defer por función
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	//continue y break

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Es 8 Break")
			break
		}
	}
}
