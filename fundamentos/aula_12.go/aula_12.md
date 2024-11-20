## Anotações
Os operadores lógicos em Go retornam valores booleanos baseados em condições:

- `&&` (E lógico): Verdadeiro se ambas as condições forem verdadeiras.
- `||` (OU lógico): Verdadeiro se pelo menos uma das condições for verdadeira.
- `!` (Não lógico): Inverte o valor lógico (verdadeiro vira falso, e vice-versa).

Em Go, não existe um operador lógico XOR (ou exclusivo) embutido diretamente. No entanto, você pode implementar o comportamento do XOR utilizando outros operadores, como `!=` ou combinando `&&`, `||`, e `!`.