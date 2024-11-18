package main

func somar(a int, b int) int {
	return a + b
}

func imprinmir(valor int) {
	println(valor)
}

func main() {
	resultado := somar(10, 20)
	imprinmir(resultado)
}