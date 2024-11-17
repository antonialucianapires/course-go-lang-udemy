## Anotações sobre Tipos Básicos em Go

Go é uma linguagem com **tipagem estática e forte**. Os tipos das variáveis são definidos em tempo de compilação e não podem ser alterados em tempo de execução. Operações entre tipos incompatíveis exigem conversão explícita, reduzindo erros de programação.

---

### Números Inteiros
**Tipos:** `int8`, `int16`, `int32`, `int64`  
> Podem armazenar números positivos e negativos.

```go
i1 := math.MaxInt64
fmt.Println(i1) // Exibe o maior valor de int64
```

#### Alias de Inteiros
- `byte`:  lias para uint8.
- `rune`: Alias para int32, usado para representar caracteres Unicode.

### Números reais
Tipos: `float32`, `float64`
> float64 é o padrão para literais de ponto flutuante.
```go
var x float32 = 49.99
fmt.Println(x) // Exibe: 49.99
```

### Booleanos
- Tipo: `bool`
> Representa valores true ou false.

```go
bo := true
fmt.Println(bo) // Saída: true
fmt.Println(!bo) // Saída: false
```

### Strings
Definição de texto:
```go
s1 := "Olá meu nome é Gabriel"
fmt.Println(s1 + "!") // Concatenação de strings
fmt.Println(len(s1))  // Comprimento da string
```

#### Strings com Múltiplas Linhas
```go
s2 := `Olá
meu
nome
é
Scoob`
fmt.Println(len(s2)) // Comprimento da string incluindo quebras de linha
```

### Caracteres
Representados como literais rune entre aspas simples.
```go
char := 'a'
fmt.Println(char) // Saída: 97 (valor Unicode do caractere 'a')
```