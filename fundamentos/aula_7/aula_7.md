## Anotações

### Conversão de int para float64
```go
x := 2.4
y := 2
fmt.Println(x / float64(y)) // Saída: 1.2
```

### Conversão de float64 para int
```go
nota := 6.9
notaFinal := int(nota)
fmt.Println(notaFinal) // Saída: 6
```

### Conversão de int para string
```go
fmt.Println("Teste " + strconv.Itoa(123)) // Saída: Teste 123
```

### Conversão de string para int
Usa a função `strconv.Atoi (ASCII to Integer)`, que retorna dois valores: o número convertido e um erro.
```go
num, err := strconv.Atoi("123")
if err == nil {
    fmt.Println(num - 122) // Saída: 1
}
```

### Conversão de string para bool
Usa a função `strconv.ParseBool`. Aceita strings como "true", "false", "1" e "0".
```go
fmt.Println("Teste " + string(123)) // Saída: Teste {
```