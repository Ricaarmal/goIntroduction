package main

import "fmt"

func main() {

	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For While
	counter := 0
	for i := 0; i < counter; i++ {
		fmt.Println(i)
		counter++
	}

	//Para casos infinitos
	/*for {
		fmt.Println(counter)
		counter++
	}
	*/

	for j := 10; j == 0; j-- {
		fmt.Println(j)
	}
}
