package main

import "fmt"

func main() {
	i := 1 //ocupa 64bits
	var p *int = nil //ponteiro para inteiro

	p = &i //pegando o endereço da variável i

	*p++ //desreferenciando (pegando o valor) do ponteiro e incrementando

	i++ //incrementando a variável i

	//Go não tem aritmética de ponteiros

	fmt.Println(p, *p, i, &i)
}