# 📖 **A Bíblia de Go**

![Capa do Livro: A Bíblia de Go](book-cover.webp)

---

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
- [8.1 Métodos Associados a Structs](section-8.1.md)
- [8.2 Receptores (`value receiver` vs `pointer receiver`)](section-8.2.md)
- [8.3 Interfaces e Polimorfismo](section-8.3.md)
- [8.4 Interface `io.Reader` e `io.Writer`](section-8.4.md)
- [8.5 Implementação Implícita de Interfaces](section-8.5.md)

### 🔹 Capítulo 9: Embedding e Composição
- [9.1 Embedding de Structs (Herança Simples)](section-9.1.md)
- [9.2 Implementação de Múltiplas Interfaces](section-9.2.md)
- [9.3 Métodos em Embeddings](section-9.3.md)
- [9.4 Composição vs. Herança em Go](section-9.4.md)

## 📌 Parte 4: Concorrência e Paralelismo

### 🔹 Capítulo 10: Goroutines e Channels
- [10.1 Criando e Executando Goroutines](section-10.1.md)
- [10.2 `sync.WaitGroup`](section-10.2.md)
- [10.3 Comunicação entre Goroutines com Channels (`chan`)](section-10.3.md)
- [10.4 Channels Buffered e Unbuffered](section-10.4.md)
- [10.5 `select` para Multiplexação de Canais](section-10.5.md)

### 🔹 Capítulo 11: Sincronização e Controle de Concorrência
- [11.1 Mutexes (`sync.Mutex`, `sync.RWMutex`)](section-11.1.md)
- [11.2 `sync.Cond`](section-11.2.md)
- [11.3 `sync.Once`](section-11.3.md)
- [11.4 `sync/atomic`](section-11.4.md)
- [11.5 Pool de Goroutines (`sync.Pool`)](section-11.5.md)

### 🔹 Capítulo 12: Context e Cancelamento
- [12.1 O Pacote `context`](section-12.1.md)
- [12.2 `context.WithCancel`](section-12.2.md)
- [12.3 `context.WithDeadline`](section-12.3.md)
- [12.4 `context.WithTimeout`](section-12.4.md)

## 📌 Parte 5: Manipulação de Arquivos e Redes

### 🔹 Capítulo 13: Entrada e Saída de Dados
- [13.1 Manipulação de Arquivos (`os`, `io/ioutil`)](section-13.1.md)
- [13.2 Leitura e Escrita em CSV e JSON](section-13.2.md)
- [13.3 Streaming com `bufio`](section-13.3.md)
- [13.4 Tratamento de Erros (`errors`, `fmt.Errorf`)](section-13.4.md)

### 🔹 Capítulo 14: Programação de Redes
- [14.1 Comunicação via TCP e UDP (`net`)](section-14.1.md)
- [14.2 Criando um Servidor e um Cliente TCP](section-14.2.md)
- [14.3 HTTP com `net/http`](section-14.3.md)
- [14.4 WebSockets e GRPC](section-14.4.md)

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


---

📌 **Esse livro é um guia completo para dominar Go, cobrindo desde os fundamentos até técnicas avançadas.** 🚀

