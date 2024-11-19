package main

import "fmt"

func main() {
	a, b := 10, 5

	// Operadores aritméticos
	fmt.Println(a+b) // Adição: 15
	fmt.Println(a-b) // Subtração: 5
	fmt.Println(a*b) // Multiplicação: 50
	fmt.Println(a/b) // Divisão: 2
	fmt.Println(a%b) // Módulo: 0

	// Operadores de comparação
	fmt.Println(a == b) // Igualdade: false
	fmt.Println(a != b) // Diferença: true
	fmt.Println(a > b)  // Maior que: true
	fmt.Println(a < b)  // Menor que: false	

	// Operadores lógicos
	fmt.Println(a > b && a < b) // E: false
	fmt.Println(a > b || a < b) // OU: true
	fmt.Println(!(a > b))        // NÃO: false

}