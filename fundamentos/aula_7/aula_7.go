package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// fmt.Println(x / y) // Error: invalid operation: x / y (mismatched types float64 and int)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste " + string(123)) // 123 é o código da tabela Unicode

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") //retorna dois valores, o segundo é o erro
	fmt.Println(num - 122)

	// string para bool
	b, _ := strconv.ParseBool("true") //ele aceita "1" e "0" também
	if b {
		fmt.Println("Verdadeiro")
	}
}