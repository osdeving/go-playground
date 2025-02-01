# 📖 A Bíblia de Go – Sumário Completo

## 📌 Parte 1: Fundamentos da Linguagem

### 🔹 Capítulo 1: Introdução ao Go
- [História e Motivação](section-1.1.md)
- [Filosofia do Go](section-1.2.md)
- [Diferenças entre Go e outras linguagens (C, Java, Python)](section-1.3.md)
- [Instalação e Configuração do Ambiente](section-1.4.md)
- [Estrutura de um Programa Go](section-1.5.md)
- [O Primeiro Programa: "Hello, World!"](section-1.6.md)

### 🔹 Capítulo 2: Sintaxe Básica
- Declaração de Variáveis (`var`, `:=`)
- Tipos Primitivos (`int`, `float64`, `bool`, `string`)
- Operadores Aritméticos, Lógicos e Comparativos
- Entrada e Saída com `fmt`
- Conversão de Tipos

### 🔹 Capítulo 3: Controle de Fluxo
- Estruturas Condicionais: `if`, `else if`, `switch`
- Laços de Repetição: `for`, `range`
- Uso de `break`, `continue`, `goto`
- Defer, Panic e Recover

### 🔹 Capítulo 4: Funções em Go
- Declaração e Uso de Funções
- Parâmetros e Retornos
- Retornos Nomeados
- Funções Variádicas
- Funções Anônimas e Closures
- Recursão
- Ponteiros e Funções (`*`, `&`)

## 📌 Parte 2: Estruturas de Dados e Manipulação de Memória

### 🔹 Capítulo 5: Arrays, Slices e Strings
- Declaração e Manipulação de Arrays
- Slices: Conceito, Capacidade e Expansão
- Strings e Runas (`rune`)
- Strings Imutáveis e Manipulação com `strings` e `bytes`
- Deep Copy vs. Shallow Copy

### 🔹 Capítulo 6: Mapas e Estruturas
- Declaração e Manipulação de Mapas (`map[key]value`)
- Operações Comuns (`delete`, `len`, `range`)
- Structs e Métodos
- Campos Opcionais e `omitempty`
- Comparação de Structs

### 🔹 Capítulo 7: Ponteiros e Gerenciamento de Memória
- Conceito de Ponteiros (`*`, `&`)
- Ponteiros para Structs e Funções
- O Pacote `unsafe`
- Alocação Dinâmica com `new` e `make`
- Anatomia do Garbage Collector do Go

## 📌 Parte 3: Programação Orientada a Objetos em Go

### 🔹 Capítulo 8: Métodos e Interfaces
- Métodos Associados a Structs
- Receptores (`value receiver` vs `pointer receiver`)
- Interfaces e Polimorfismo
- Interface `io.Reader` e `io.Writer`
- Implementação Implícita de Interfaces

### 🔹 Capítulo 9: Embedding e Composição
- Embedding de Structs (Herança Simples)
- Implementação de Múltiplas Interfaces
- Métodos em Embeddings
- Composição vs. Herança em Go

## 📌 Parte 4: Concorrência e Paralelismo

### 🔹 Capítulo 10: Goroutines e Channels
- Criando e Executando Goroutines
- `sync.WaitGroup`
- Comunicação entre Goroutines com Channels (`chan`)
- Channels Buffered e Unbuffered
- `select` para Multiplexação de Canais

### 🔹 Capítulo 11: Sincronização e Controle de Concorrência
- Mutexes (`sync.Mutex`, `sync.RWMutex`)
- `sync.Cond`
- `sync.Once`
- `sync/atomic`
- Pool de Goroutines (`sync.Pool`)

### 🔹 Capítulo 12: Context e Cancelamento
- O Pacote `context`
- `context.WithCancel`
- `context.WithDeadline`
- `context.WithTimeout`

## 📌 Parte 5: Manipulação de Arquivos e Redes

### 🔹 Capítulo 13: Entrada e Saída de Dados
- Manipulação de Arquivos (`os`, `io/ioutil`)
- Leitura e Escrita em CSV e JSON
- Streaming com `bufio`
- Tratamento de Erros (`errors`, `fmt.Errorf`)

### 🔹 Capítulo 14: Programação de Redes
- Comunicação via TCP e UDP (`net`)
- Criando um Servidor e um Cliente TCP
- HTTP com `net/http`
- WebSockets e GRPC

## 📌 Parte 6: Desenvolvimento Web e APIs

### 🔹 Capítulo 15: Criando APIs RESTful
- Frameworks Web (Gin, Echo)
- Manipulação de Requisições e Respostas
- Middlewares e Autenticação
- JWT e OAuth2
- Serialização e Desserialização de JSON

### 🔹 Capítulo 16: Trabalhando com Bancos de Dados
- Drivers SQL (`database/sql`)
- ORM com GORM
- Conexão com MongoDB e Redis
- Transações e Pool de Conexões

## 📌 Parte 7: Testes, Performance e Segurança

### 🔹 Capítulo 17: Testes em Go
- Testes Unitários (`testing`)
- Testes de Benchmark
- Testes de Integração e Mocks

### 🔹 Capítulo 18: Performance e Profiling
- Benchmarks (`go test -bench`)
- Uso do `pprof`
- Gerenciamento de Memória

### 🔹 Capítulo 19: Segurança e Melhores Práticas
- Tratamento de Erros
- Proteção contra Data Races
- Validação de Entrada
- Segurança em APIs REST
- Práticas de Desenvolvimento Seguro

## 📌 Parte 8: Deploy, DevOps e Ferramentas

### 🔹 Capítulo 20: Compilação e Deploy
- `go build`, `go install`, `go run`
- Cross Compilation
- Distribuindo Binários Go

### 🔹 Capítulo 21: Docker e Kubernetes
- Criando e Otimizando Imagens Docker para Go
- Deploy no Kubernetes
- ConfigMaps e Secrets

### 🔹 Capítulo 22: Monitoramento e Logging
- Monitoramento com Prometheus
- Logging com Logrus e Zap
- Health Checks e Tracing

---

## 📌 Apêndices

### 🔹 Apêndice A: Certificação Go
- Estrutura do Exame
- Questões Simuladas
- Dicas para a Prova

### 🔹 Apêndice B: Recursos e Bibliotecas
- Frameworks e Ferramentas Essenciais
- Repositórios Importantes no GitHub
- Comunidade Go e Fóruns

### 🔹 Apêndice C: Estudos de Caso
- Arquiteturas Reais de Projetos em Go
- Aplicações Escaláveis em Produção
