## Anotações
O pacote fmt em Go é utilizado para entrada e saída formatada, principalmente para imprimir mensagens no console. Ele fornece funções para exibir strings, valores de variáveis e realizar formatações personalizadas em strings.

1. `fmt.Print`
- Exibe os argumentos na saída padrão sem uma nova linha ao final.
- Argumentos são convertidos para strings e concatenados.
```go
fmt.Print("Mesma")
fmt.Print(" linha.")
```

2. `fmt.Println`
- Similar ao `fmt.Print`, mas adiciona uma nova linha (\n) automaticamente no final.
```go
fmt.Println("Nova")
fmt.Println("linha.")
```

3.`fmt.Sprint`
- Converte os valores passados em strings e retorna a string resultante.
- Não imprime diretamente, mas permite criar strings para uso posterior.
```go
xs := fmt.Sprint(x)
fmt.Println("O valor de x é " + xs)
```

4. `fmt.Printf`
- Permite formatação detalhada da saída usando especificadores de formato.
- Os principais especificadores incluem:
    - `%d`: Inteiros decimais.
    - `%f`: Números de ponto flutuante.
    - `%.nf`: Números de ponto flutuante com n casas decimais.
    - `%t`: Booleanos (true ou false).
    - `%s`: Strings.
```go
fmt.Printf("O valor de x é %.2f", x)
fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
```

5. `fmt.Sprintf`
- Semelhante ao `fmt.Printf`, mas retorna a string formatada em vez de exibi-la.
- Útil para criar strings complexas para manipulação ou uso posterior.
```go
mensagem := fmt.Sprintf("O valor de x é %.2f", x)
fmt.Println(mensagem)
```

## Importante
**Erro de concatenação:** Go não permite a concatenação direta de strings e outros tipos:
```go
fmt.Println("O valor de x é " + x) // Erro de compilação
```