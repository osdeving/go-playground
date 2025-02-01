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
- [Declaração de Variáveis (`var`, `:=`)](section-2.1.md)
- [Tipos Primitivos (`int`, `float64`, `bool`, `string`)](section-2.2.md)
- [Operadores Aritméticos, Lógicos e Comparativos](section-2.3.md)
- [Entrada e Saída com `fmt`](section-2.4.md)
- [Conversão de Tipos](section-2.5.md)

### 🔹 Capítulo 3: Controle de Fluxo
- [Estruturas Condicionais: `if`, `else if`, `switch`](section-3.1.md)
- [Laços de Repetição: `for`, `range`](section-3.2.md)
- [Uso de `break`, `continue`, `goto`](section-3.3.md)
- [Defer, Panic e Recover](section-3.4.md)

### 🔹 Capítulo 4: Funções em Go
- [Declaração e Uso de Funções](section-4.1.md)
- [Parâmetros e Retornos](section-4.2.md)
- [Retornos Nomeados](section-4.3.md)
- [Funções Variádicas](section-4.4.md)
- [Funções Anônimas e Closures](section-4.5.md)
- [Recursão](section-4.6.md)
- [Ponteiros e Funções (`*`, `&`)](section-4.7.md)
- [Entendendo e Recriando Funções Built-in do Go](section-4.8.md)

## 📌 Parte 2: Estruturas de Dados e Manipulação de Memória

### 🔹 Capítulo 5: Arrays, Slices e Strings
- [Declaração e Manipulação de Arrays](section-5.1.md)
- [Slices: Conceito, Capacidade e Expansão](section-5.2.md)
- [Strings e Runas (`rune`)](section-5.3.md)
- [Strings Imutáveis e Manipulação com `strings` e `bytes`](section-5.4.md)
- [Deep Copy vs. Shallow Copy](section-5.5.md)

### 🔹 Capítulo 6: Mapas e Estruturas
- [6.1 Declaração e Manipulação de Mapas (`map[key]value`)](section-6.1.md)
- [6.2 Operações Comuns (`delete`, `len`, `range`)](section-6.2.md)
- [6.3 Structs e Métodos](section-6.3.md)
- [6.4 Campos Opcionais e `omitempty`](section-6.4.md)
- [6.5 Comparação de Structs](section-6.5.md)

### 🔹 Capítulo 7: Ponteiros e Gerenciamento de Memória
- [7.1 Conceito de Ponteiros (`*`, `&`)](section-7.1.md)
- [7.2 Ponteiros para Structs e Funções](section-7.2.md)
- [7.3 O Pacote `unsafe`](section-7.3.md)
- [7.4 Alocação Dinâmica com `new` e `make`](section-7.4.md)
- [7.5 Anatomia do Garbage Collector do Go](section-7.5.md)

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
