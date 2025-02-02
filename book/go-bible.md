# 📖 **A Bíblia de Go**

![Capa do Livro: A Bíblia de Go](go-bible.jpg)

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
- [10.6 Exemplos práticos de Concorrência](section-10.6.md)

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
- [15.1 Frameworks Web (Gin, Echo)](section-15.1.md)
- [15.2 Manipulação de Requisições e Respostas](section-15.2.md)
- [15.3 Middlewares e Autenticação](section-15.3.md)
- [15.4 JWT e OAuth2](section-15.4.md)
- [15.5 Serialização e Desserialização de JSON](section-15.5.md)

### 🔹 Capítulo 16: Trabalhando com Bancos de Dados
- [16.1 Drivers SQL (`database/sql`)](section-16.1.md)
- [16.2 ORM com GORM](section-16.2.md)
- [16.3 Conexão com MongoDB e Redis](section-16.3.md)
- [16.4 Transações e Pool de Conexões](section-16.4.md)

## 📌 Parte 7: Testes, Performance e Segurança

### 🔹 Capítulo 17: Testes em Go
- [17.1 Testes Unitários (`testing`)](section-17.1.md)
- [17.2 Testes de Benchmark](section-17.2.md)
- [17.3 Testes de Integração e Mocks](section-17.3.md)

### 🔹 Capítulo 18: Performance e Profiling
- [18.1 Benchmarks (`go test -bench`)](section-18.1.md)
- [18.2 Uso do `pprof`](section-18.2.md)
- [18.3 Gerenciamento de Memória](section-18.3.md)

### 🔹 Capítulo 19: Segurança e Melhores Práticas
- [19.1 Tratamento de Erros](section-19.1.md)
- [19.2 Proteção contra Data Races](section-19.2.md)
- [19.3 Validação de Entrada](section-19.3.md)
- [19.4 Segurança em APIs REST](section-19.4.md)
- [19.5 Práticas de Desenvolvimento Seguro](section-19.5.md)

## 📌 Parte 8: Deploy, DevOps e Ferramentas

### 🔹 Capítulo 20: Compilação e Deploy
- [20.1 `go build`, `go install`, `go run`](section-20.1.md)
- [20.2 Cross Compilation](section-20.2.md)
- [20.3 Distribuindo Binários Go](section-20.3.md)

### 🔹 Capítulo 21: Docker e Kubernetes
- [21.1 Criando e Otimizando Imagens Docker para Go](section-21.1.md)
- [21.2 Deploy no Kubernetes](section-21.2.md)
- [21.3 ConfigMaps e Secrets](section-21.3.md)

### 🔹 Capítulo 22: Monitoramento e Logging
- [22.1 Monitoramento com Prometheus](section-22.1.md)
- [22.2 Logging com Logrus e Zap](section-22.2.md)
- [22.3 Health Checks e Tracing](section-22.3.md)

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

