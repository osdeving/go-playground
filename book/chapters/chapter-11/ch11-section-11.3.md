# **11.3 `sync.Once`: Inicialização Segura em Go**

Em alguns cenários, é necessário garantir que **um trecho de código seja executado apenas uma vez**, independentemente do número de Goroutines concorrentes.  
Para isso, o Go fornece o **`sync.Once`**, um mecanismo eficiente para inicializações seguras e execução única de código crítico.

Nesta seção, exploraremos:

- O que é `sync.Once` e como funciona
- Diferença entre `sync.Once` e `sync.Mutex`
- Casos de uso comuns, como inicialização de singletons
- Comparação com técnicas manuais de sincronização
- Boas práticas para evitar erros ao usá-lo

---

## **11.3.1 O Que é `sync.Once`?**

O `sync.Once` garante que um bloco de código seja executado **exatamente uma vez**, mesmo quando múltiplas Goroutines tentam acessá-lo simultaneamente.  

📌 **Importante:** Após a primeira execução, chamadas subsequentes para `Do()` **não executam novamente** a função registrada.  

✅ **Exemplo básico de `sync.Once`:**

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func inicializar() {
    fmt.Println("Executando apenas uma vez!")
}

func main() {
    for i := 0; i < 5; i++ {
        go once.Do(inicializar) // Apenas a primeira Goroutine executa
    }
}
```

📌 **Mesmo com 5 chamadas concorrentes, `inicializar()` só será executado uma vez!**

✅ **Isso é útil para inicializar conexões, caches e configurações globais de forma segura.**  

---

## **11.3.2 `sync.Once` vs. `sync.Mutex`**

Muitos desenvolvedores inicialmente usam `sync.Mutex` para garantir inicialização única:

```go
var mutex sync.Mutex
var inicializado bool

func inicializar() {
    mutex.Lock()
    defer mutex.Unlock()

    if !inicializado {
        fmt.Println("Executando apenas uma vez!")
        inicializado = true
    }
}
```

📌 **O problema desse código é que `mutex.Lock()` pode ser chamado várias vezes.**  

✅ **Com `sync.Once`, esse problema desaparece:**

```go
var once sync.Once

func inicializar() {
    once.Do(func() {
        fmt.Println("Executando apenas uma vez!")
    })
}
```

📌 **O código fica mais limpo, seguro e evita verificações manuais.**  

✅ **`sync.Once` é a melhor escolha para inicialização única!**

---

## **11.3.3 Quando Usar `sync.Once`?**

`sync.Once` é ideal para:

1️⃣ **Inicializar singletons** (exemplo: conexão com banco de dados)  
2️⃣ **Criar configurações globais**  
3️⃣ **Carregar arquivos de configuração uma única vez**  
4️⃣ **Inicializar pools de recursos compartilhados**  

✅ **Exemplo: Inicialização segura de um pool de conexões**

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var dbConnection string

func connectDatabase() {
    once.Do(func() {
        dbConnection = "Conexão estabelecida"
        fmt.Println("Banco de dados conectado!")
    })
}

func main() {
    go connectDatabase()
    go connectDatabase()

    fmt.Println(dbConnection) // Garantido que foi inicializado
}
```

📌 **Mesmo com múltiplas chamadas, `connectDatabase()` só executa uma vez.**  

✅ **Isso evita bugs onde múltiplas conexões seriam criadas desnecessariamente.**  

---

## **11.3.4 `sync.Once` e Goroutines Concorrentes**

Se várias Goroutines chamarem `once.Do()` simultaneamente, o Go garante que apenas **uma** delas executará a função, enquanto as demais aguardarão a finalização.  

✅ **Exemplo de execução segura:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var once sync.Once

func tarefa() {
    fmt.Println("Executando tarefa única!")
}

func main() {
    for i := 0; i < 3; i++ {
        go once.Do(tarefa)
    }

    time.Sleep(time.Second) // Aguarda a execução
}
```

📌 **Não importa quantas Goroutines chamem `once.Do()`, apenas uma executará `tarefa()`.**  

✅ **Go lida automaticamente com concorrência, evitando condições de corrida.**

---

## **11.3.5 Erros Comuns ao Usar `sync.Once`**

❌ **Chamar `once.Do()` com funções que retornam valores**

```go
var once sync.Once

func inicializar() string {
    return "Erro! Função com retorno"
}

// once.Do(inicializar) // ERRO: sync.Once.Do não aceita funções com retorno
```

✅ **`sync.Once` aceita apenas funções sem retorno.**  
📌 **Se precisar armazenar um valor, use variáveis globais.**  

---

❌ **Reutilizar `sync.Once` após a primeira execução**

```go
var once sync.Once

once.Do(func() {
    fmt.Println("Executando...")
})

// once = sync.Once{} // ERRO: Resetar `sync.Once` manualmente pode causar problemas!
```

✅ **Se precisar repetir a inicialização, use outro mecanismo como `sync.Mutex`.**  

---

## **11.3.6 Comparação: `sync.Once` vs. Outras Técnicas**

| Técnica | Uso Principal | Executa Apenas Uma Vez? | Bloqueante? | Simples de Usar? |
|---------|--------------|-----------------|------------|-----------------|
| `sync.Once` | Inicialização única | ✅ Sim | ✅ Sim | ✅ Sim |
| `sync.Mutex` | Exclusão mútua | ❌ Não | ✅ Sim | ❌ Não |
| `init()` | Execução automática | ✅ Sim | ❌ Não | ✅ Sim |
| `sync.Atomic` | Operações atômicas | ❌ Não | ❌ Não | ✅ Sim |

📌 **Use `sync.Once` sempre que precisar de inicialização única concorrente.**  

✅ **Se precisar de inicialização automática, `init()` pode ser uma alternativa melhor.**  

---

## **11.3.7 Boas Práticas**

✔ **Use `sync.Once` para inicializações únicas em ambiente concorrente.**  
✔ **Evite funções com retorno dentro de `once.Do()`.**  
✔ **Se precisar reexecutar código, `sync.Once` não é a melhor escolha.**  
✔ **Combine `sync.Once` com variáveis globais para armazenar valores iniciais.**  

---

## **Conclusão**

O **`sync.Once`** é uma ferramenta essencial para garantir que blocos de código sejam executados **apenas uma vez** em ambientes concorrentes.  
No próximo capítulo, exploraremos **`sync/atomic`**, um poderoso recurso para operações atômicas e manipulação segura de memória em Go! 🚀
