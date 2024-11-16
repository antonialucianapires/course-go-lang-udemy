## Anotações

`go` mostra opções de comandos disponíveis
`go help` explica o que cada comando faz
`go version` mostra a versão do Go instalada
`godoc -http=6060` abre a documentação local do Go
`go env` mostra variáveis de ambiente e configurações (ex.: arquitetura)
`go doc cmd/vet` mostra a documentação do comando vet
`go vet app.go` verifica possíveis problemas no código
`go build app.go` compila o código
`go run` compila e executa diretamente

### Workspace

No contexto da linguagem Go, um **workspace** é um diretório que organiza o código-fonte, dependências e artefatos compilados de um projeto. Tradicionalmente, o workspace é estruturado com três subdiretórios principais:

1. **`src`**: contém o código-fonte dos projetos e pacotes.
2. **`pkg`**: armazena os arquivos compilados dos pacotes.
3. **`bin`**: guarda os executáveis gerados.

O caminho do workspace é definido pela variável de ambiente `GOPATH`. Com a introdução dos módulos em Go 1.11, o uso do `GOPATH` tornou-se opcional, permitindo que os projetos sejam gerenciados fora dessa estrutura tradicional. No entanto, o conceito de workspace ainda é relevante, especialmente ao trabalhar com múltiplos módulos simultaneamente.