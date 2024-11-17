## Anotações
Em Go, quando uma variável é declarada mas não inicializada, ela assume um valor zero padrão correspondente ao seu tipo. Esses valores zero garantem que não haja dados indefinidos.

em Go

| **Tipo**         | **Exemplo**           | **Valor Zero** | **Descrição**                        |
|-------------------|-----------------------|----------------|--------------------------------------|
| **`int`**        | `var a int`           | `0`            | Número inteiro inicializado como zero. |
| **`float64`**    | `var b float64`       | `0.0`          | Número de ponto flutuante com valor zero. |
| **`bool`**       | `var c bool`          | `false`        | Booleano inicializado como falso.    |
| **`string`**     | `var d string`        | `""` (vazia)   | String vazia, sem caracteres.        |
| **`*int`**       | `var e *int`          | `<nil>`        | Ponteiro inicializado como nulo.    |