## Anotações

### Constantes
Valores imutáveis definidos durante a compilação. São usadas para representar valores fixos que não mudam durante a execução do programa.

#### Declaração explícita com tipo:
```go
const PI float64 = 3.14159
```

#### Declaração implícita (tipo inferido):
```go
const Gravidade = 9.8
```

> O compilador infere o tipo como float64.

#### Multideclaração em uma única linha
```go
const a, b, c = 1, 2, 3
```

#### Agrupamento de constantes
O agrupamento é usado para organizar constantes relacionadas.
```go
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)
```

> Constantes não possuem uma forma abreviada como o operador :=

### Variáveis 
Valores mutáveis que podem mudar durante a execução do programa.

#### Declaração explícita com tipo
```go
var x int = 10
```

#### Declaração implícita (tipo inferido)
```go
var y = 20.5
```
> O compilador infere o tipo como float64.

#### Declaração abreviada (forma reduzida)
```go
z := "Hello, Go!"

```

#### Multideclaração de variáveis
```go
var a, b, c = 1, 2, "texto"
```

#### Agrupamento de variáveis
```go
var (
    largura = 100
    altura  = 200
)
```

## Agrupamento: Reutilização de valores
```go
const (
    A = 1
    B      // B terá o valor 1 (mesmo de A)
    C = 2
    D      // D terá o valor 2 (mesmo de C)
)
fmt.Println(A, B, C, D) // Saída: 1 1 2 2
```


## Observações importantes
- Constantes são avaliadas em tempo de compilação, enquanto variáveis são alocadas e manipuladas em tempo de execução.

- Não é permitido declarar variáveis sem uso. O compilador gera erro caso uma variável seja declarada mas não utilizada.