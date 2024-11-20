package main

import "fmt"

func main() {
	x := 1
	y := 2

	//Operador de incremento e decremento. Apenas pós-fixado
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)
	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	//Operador ternário não existe em Go
}