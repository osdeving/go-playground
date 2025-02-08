# 📖 **A Bíblia de Go**

![Capa do Livro: A Bíblia de Go](go-bible.jpg)

---

# 📖 A Bíblia de Go – Sumário Completo

## 📌 Parte 1: Fundamentos da Linguagem

### 🔹 Capítulo 1: Introdução ao Go
- [História e Motivação](chapters/chapter-1/ch1-section-1.1.md)
- [Filosofia do Go](chapters/chapter-1/ch1-section-1.2.md)
- [Diferenças entre Go e outras linguagens (C, Java, Python)](chapters/chapter-1/ch1-section-1.3.md)
- [Instalação e Configuração do Ambiente](chapters/chapter-1/ch1-section-1.4.md)
- [Estrutura de um Programa Go](chapters/chapter-1/ch1-section-1.5.md)
- [O Primeiro Programa: "Hello, World!"](chapters/chapter-1/ch1-section-1.6.md)

### 🔹 Capítulo 2: Sintaxe Básica
- [Declaração de Variáveis (`var`, `:=`)](chapters/chapter-2/ch2-section-2.1.md)
- [Tipos Primitivos (`int`, `float64`, `bool`, `string`)](chapters/chapter-2/ch2-section-2.2.md)
- [Operadores Aritméticos, Lógicos e Comparativos](chapters/chapter-2/ch2-section-2.3.md)
- [Entrada e Saída com `fmt`](chapters/chapter-2/ch2-section-2.4.md)
- [Conversão de Tipos](chapters/chapter-2/ch2-section-2.5.md)

### 🔹 Capítulo 3: Controle de Fluxo
- [Estruturas Condicionais: `if`, `else if`, `switch`](chapters/chapter-3/sections/section-3.1.md)
- [Laços de Repetição: `for`, `range`](chapters/chapter-3/sections/section-3.2.md)
- [Uso de `break`, `continue`, `goto`](chapters/chapter-3/sections/section-3.3.md)
- [Defer, Panic e Recover](chapters/chapter-3/sections/section-3.4.md)

### 🔹 Capítulo 4: Funções em Go
- [Declaração e Uso de Funções](chapters/chapter-4/sections/section-4.1.md)
- [Parâmetros e Retornos](chapters/chapter-4/sections/section-4.2.md)
- [Retornos Nomeados](chapters/chapter-4/sections/section-4.3.md)
- [Funções Variádicas](chapters/chapter-4/sections/section-4.4.md)
- [Funções Anônimas e Closures](chapters/chapter-4/sections/section-4.5.md)
- [Recursão](chapters/chapter-4/sections/section-4.6.md)
- [Ponteiros e Funções (`*`, `&`)](chapters/chapter-4/sections/section-4.7.md)
- [Entendendo e Recriando Funções Built-in do Go](chapters/chapter-4/sections/section-4.8.md)

## 📌 Parte 2: Estruturas de Dados e Manipulação de Memória

### 🔹 Capítulo 5: Arrays, Slices e Strings
- [Declaração e Manipulação de Arrays](chapters/chapter-5/sections/section-5.1.md)
- [Slices: Conceito, Capacidade e Expansão](chapters/chapter-5/sections/section-5.2.md)
- [Strings e Runas (`rune`)](chapters/chapter-5/sections/section-5.3.md)
- [Strings Imutáveis e Manipulação com `strings` e `bytes`](chapters/chapter-5/sections/section-5.4.md)
- [Deep Copy vs. Shallow Copy](chapters/chapter-5/sections/section-5.5.md)

### 🔹 Capítulo 6: Mapas e Estruturas
- [6.1 Declaração e Manipulação de Mapas (`map[key]value`)](chapters/chapter-6/sections/section-6.1.md)
- [6.2 Operações Comuns (`delete`, `len`, `range`)](chapters/chapter-6/sections/section-6.2.md)
- [6.3 Structs e Métodos](chapters/chapter-6/sections/section-6.3.md)
- [6.4 Campos Opcionais e `omitempty`](chapters/chapter-6/sections/section-6.4.md)
- [6.5 Comparação de Structs](chapters/chapter-6/sections/section-6.5.md)

### 🔹 Capítulo 7: Ponteiros e Gerenciamento de Memória
- [7.1 Conceito de Ponteiros (`*`, `&`)](chapters/chapter-7/sections/section-7.1.md)
- [7.2 Ponteiros para Structs e Funções](chapters/chapter-7/sections/section-7.2.md)
- [7.3 O Pacote `unsafe`](chapters/chapter-7/sections/section-7.3.md)
- [7.4 Alocação Dinâmica com `new` e `make`](chapters/chapter-7/sections/section-7.4.md)
- [7.5 Anatomia do Garbage Collector do Go](chapters/chapter-7/sections/section-7.5.md)

## 📌 Parte 3: Programação Orientada a Objetos em Go

### 🔹 Capítulo 8: Métodos e Interfaces
- [8.1 Métodos Associados a Structs](chapters/chapter-8/sections/section-8.1.md)
- [8.2 Receptores (`value receiver` vs `pointer receiver`)](chapters/chapter-8/sections/section-8.2.md)
- [8.3 Interfaces e Polimorfismo](chapters/chapter-8/sections/section-8.3.md)
- [8.4 Interface `io.Reader` e `io.Writer`](chapters/chapter-8/sections/section-8.4.md)
- [8.5 Implementação Implícita de Interfaces](chapters/chapter-8/sections/section-8.5.md)

### 🔹 Capítulo 9: Embedding e Composição
- [9.1 Embedding de Structs (Herança Simples)](chapters/chapter-9/sections/section-9.1.md)
- [9.2 Implementação de Múltiplas Interfaces](chapters/chapter-9/sections/section-9.2.md)
- [9.3 Métodos em Embeddings](chapters/chapter-9/sections/section-9.3.md)
- [9.4 Composição vs. Herança em Go](chapters/chapter-9/sections/section-9.4.md)

## 📌 Parte 4: Concorrência e Paralelismo

### 🔹 Capítulo 10: Goroutines e Channels
- [10.1 Criando e Executando Goroutines](chapters/chapter-10/sections/section-10.1.md)
- [10.2 `sync.WaitGroup`](chapters/chapter-10/sections/section-10.2.md)
- [10.3 Comunicação entre Goroutines com Channels (`chan`)](chapters/chapter-10/sections/section-10.3.md)
- [10.4 Channels Buffered e Unbuffered](chapters/chapter-10/sections/section-10.4.md)
- [10.5 `select` para Multiplexação de Canais](chapters/chapter-10/sections/section-10.5.md)
- [10.6 Exemplos práticos de Concorrência](chapters/chapter-10/sections/section-10.6.md)

### 🔹 Capítulo 11: Sincronização e Controle de Concorrência
- [11.1 Mutexes (`sync.Mutex`, `sync.RWMutex`)](chapters/chapter-11/sections/section-11.1.md)
- [11.2 `sync.Cond`](chapters/chapter-11/sections/section-11.2.md)
- [11.3 `sync.Once`](chapters/chapter-11/sections/section-11.3.md)
- [11.4 `sync/atomic`](chapters/chapter-11/sections/section-11.4.md)
- [11.5 Pool de Goroutines (`sync.Pool`)](chapters/chapter-11/sections/section-11.5.md)

### 🔹 Capítulo 12: Context e Cancelamento
- [12.1 O Pacote `context`](chapters/chapter-12/sections/section-12.1.md)
- [12.2 `context.WithCancel`](chapters/chapter-12/sections/section-12.2.md)
- [12.3 `context.WithDeadline`](chapters/chapter-12/sections/section-12.3.md)
- [12.4 `context.WithTimeout`](chapters/chapter-12/sections/section-12.4.md)

## 📌 Parte 5: Manipulação de Arquivos e Redes

### 🔹 Capítulo 13: Entrada e Saída de Dados
- [13.1 Manipulação de Arquivos (`os`, `io/ioutil`)](chapters/chapter-13/sections/section-13.1.md)
- [13.2 Leitura e Escrita em CSV e JSON](chapters/chapter-13/sections/section-13.2.md)
- [13.3 Streaming com `bufio`](chapters/chapter-13/sections/section-13.3.md)
- [13.4 Tratamento de Erros (`errors`, `fmt.Errorf`)](chapters/chapter-13/sections/section-13.4.md)

### 🔹 Capítulo 14: Programação de Redes
- [14.1 Comunicação via TCP e UDP (`net`)](chapters/chapter-14/sections/section-14.1.md)
- [14.2 Criando um Servidor e um Cliente TCP](chapters/chapter-14/sections/section-14.2.md)
- [14.3 HTTP com `net/http`](chapters/chapter-14/sections/section-14.3.md)
- [14.4 WebSockets e GRPC](chapters/chapter-14/sections/section-14.4.md)

## 📌 Parte 6: Desenvolvimento Web e APIs

### 🔹 Capítulo 15: Criando APIs RESTful
- [15.1 Frameworks Web (Gin, Echo)](chapters/chapter-15/sections/section-15.1.md)
- [15.2 Manipulação de Requisições e Respostas](chapters/chapter-15/sections/section-15.2.md)
- [15.3 Middlewares e Autenticação](chapters/chapter-15/sections/section-15.3.md)
- [15.4 JWT e OAuth2](chapters/chapter-15/sections/section-15.4.md)
- [15.5 Serialização e Desserialização de JSON](chapters/chapter-15/sections/section-15.5.md)

### 🔹 Capítulo 16: Trabalhando com Bancos de Dados
- [16.1 Drivers SQL (`database/sql`)](chapters/chapter-16/sections/section-16.1.md)
- [16.2 ORM com GORM](chapters/chapter-16/sections/section-16.2.md)
- [16.3 Conexão com MongoDB e Redis](chapters/chapter-16/sections/section-16.3.md)
- [16.4 Transações e Pool de Conexões](chapters/chapter-16/sections/section-16.4.md)

## 📌 Parte 7: Testes, Performance e Segurança

### 🔹 Capítulo 17: Testes em Go
- [17.1 Testes Unitários (`testing`)](chapters/chapter-17/sections/section-17.1.md)
- [17.2 Testes de Benchmark](chapters/chapter-17/sections/section-17.2.md)
- [17.3 Testes de Integração e Mocks](chapters/chapter-17/sections/section-17.3.md)

### 🔹 Capítulo 18: Performance e Profiling
- [18.1 Benchmarks (`go test -bench`)](chapters/chapter-18/sections/section-18.1.md)
- [18.2 Uso do `pprof`](chapters/chapter-18/sections/section-18.2.md)
- [18.3 Gerenciamento de Memória](chapters/chapter-18/sections/section-18.3.md)

### 🔹 Capítulo 19: Segurança e Melhores Práticas
- [19.1 Tratamento de Erros](chapters/chapter-19/sections/section-19.1.md)
- [19.2 Proteção contra Data Races](chapters/chapter-19/sections/section-19.2.md)
- [19.3 Validação de Entrada](chapters/chapter-19/sections/section-19.3.md)
- [19.4 Segurança em APIs REST](chapters/chapter-19/sections/section-19.4.md)
- [19.5 Práticas de Desenvolvimento Seguro](chapters/chapter-19/sections/section-19.5.md)

## 📌 Parte 8: Deploy, DevOps e Ferramentas

### 🔹 Capítulo 20: Compilação e Deploy
- [20.1 `go build`, `go install`, `go run`](chapters/chapter-20/sections/section-20.1.md)
- [20.2 Cross Compilation](chapters/chapter-20/sections/section-20.2.md)
- [20.3 Distribuindo Binários Go](chapters/chapter-20/sections/section-20.3.md)

### 🔹 Capítulo 21: Docker e Kubernetes
- [21.1 Criando e Otimizando Imagens Docker para Go](chapters/chapter-21/sections/section-21.1.md)
- [21.2 Deploy no Kubernetes](chapters/chapter-21/sections/section-21.2.md)
- [21.3 ConfigMaps e Secrets](chapters/chapter-21/sections/section-21.3.md)

### 🔹 Capítulo 22: Monitoramento e Logging
- [Monitoramento com Prometheus](chapters/chapter-22/ch22-section-22.1.md)
- [Logging com Logrus e Zap](chapters/chapter-22/ch22-section-22.2.md)
- [Health Checks e Tracing](chapters/chapter-22/ch22-section-22.3.md)

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

