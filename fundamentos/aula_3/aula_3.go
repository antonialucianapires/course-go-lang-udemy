package main

import (
	"fmt"
	"math"
)

func main() { 
	const PI float64 = 3.14159
	var raio = 3.2 // inferência do tipo float64

	area := PI * math.Pow(raio, 2) // declaração e inicialização abreviada
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)
}
