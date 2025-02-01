# **12.4 `context.WithTimeout`: Cancelamento Baseado em Tempo Relativo**

O **`context.WithTimeout`** é uma variação do `context.WithDeadline`, mas com uma diferença fundamental:  
em vez de definir um **tempo absoluto** para expiração, ele define um **tempo relativo** a partir do momento da criação.

Esse método é essencial para cenários onde o tempo de execução **não pode exceder um limite máximo**, garantindo que tarefas não fiquem rodando indefinidamente.

Nesta seção, exploraremos:

- O que é `context.WithTimeout` e como funciona
- Diferença entre `WithTimeout` e `WithDeadline`
- Aplicação prática para evitar tarefas demoradas
- Cancelamento automático e controle eficiente de Goroutines
- Boas práticas e erros comuns

---

## **12.4.1 O Que é `context.WithTimeout`?**

O `context.WithTimeout` cria um **contexto cancelável após um determinado período de tempo**, independentemente do momento atual.

📌 **Ele é útil quando queremos garantir que uma operação não dure mais do que X segundos, a partir do seu início.**

✅ **Exemplo básico de `context.WithTimeout`**

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker finalizado!")
            return
        default:
            fmt.Println("Trabalhando...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel() // Cancela o contexto ao final

    go worker(ctx)

    time.Sleep(4 * time.Second) // Aguarda para visualizar o cancelamento
}
```

📌 **A Goroutine será finalizada após 3 segundos, independentemente do tempo de início.**  
📌 **Se `worker()` ainda estiver rodando, será interrompido automaticamente.**  

✅ **Isso garante que tarefas concorrentes não ultrapassem um tempo limite aceitável.**  

---

## **12.4.2 Diferença Entre `WithTimeout` e `WithDeadline`**

Ambos os métodos impõem um tempo limite, mas de formas diferentes:

| Método | O que faz? | Melhor Aplicação |
|--------|------------|----------------|
| `context.WithTimeout(ctx, duration)` | Cancela após um tempo relativo | Quando o tempo máximo é baseado no início da execução |
| `context.WithDeadline(ctx, time)` | Cancela no tempo absoluto definido | Quando há um horário fixo para expiração |

✅ **Use `WithTimeout` quando a duração for variável e relativa ao início.**  
✅ **Use `WithDeadline` quando a expiração for baseada em um horário absoluto.**  

✅ **Exemplo comparativo:**

```go
ctx1, _ := context.WithTimeout(context.Background(), 5*time.Second) // Cancela após 5s

deadline := time.Now().Add(5 * time.Second)
ctx2, _ := context.WithDeadline(context.Background(), deadline) // Cancela exatamente às 15:05:30
```

📌 **A escolha depende do cenário da aplicação e da necessidade de controle temporal.**  

---

## **12.4.3 Aplicação Prática: Evitando Requisições Bloqueadas**

O `context.WithTimeout` é amplamente utilizado para **cancelar operações que podem travar indefinidamente**.

✅ **Exemplo: Cancelando uma requisição HTTP automaticamente**

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)
    client := &http.Client{}

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Requisição cancelada:", err)
        return
    }
    defer resp.Body.Close()

    fmt.Println("Requisição concluída com sucesso!")
}
```

📌 **Se o servidor não responder em 2 segundos, a requisição será cancelada automaticamente.**  
📌 **Isso melhora a eficiência do sistema e evita travamentos inesperados.**  

✅ **Esse padrão é essencial em aplicações web e APIs.**  

---

## **12.4.4 Cancelamento Automático com `WithTimeout`**

Uma vantagem do `WithTimeout` é que **não precisamos chamar `cancel()` manualmente**, pois ele **se cancela sozinho ao atingir o tempo limite**.

✅ **Exemplo de cancelamento automático:**

```go
ctx, _ := context.WithTimeout(context.Background(), 3*time.Second) // Sem necessidade de chamar cancel()
```

📌 **Se o tempo for atingido, `ctx.Done()` será fechado automaticamente.**  

✅ **Isso reduz a complexidade do código e evita esquecimentos na lógica de cancelamento.**  

---

## **12.4.5 Erros Comuns ao Usar `context.WithTimeout`**

❌ **Definir um tempo muito curto sem necessidade**

```go
ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond) // ERRO: Pode cancelar antes da tarefa terminar!
```

📌 **Se o tempo for muito curto, pode causar cancelamentos desnecessários.**  
📌 **Ajuste os valores de tempo com base no comportamento real das operações.**  

✅ **Garanta tempos realistas para evitar falhas inesperadas.**  

---

❌ **Esquecer de propagar `ctx` para funções concorrentes**

```go
func process() {
    select {
    case <-ctx.Done(): // ERRO: `ctx` não foi passado como argumento!
    }
}
```

📌 **Se `ctx` não for propagado corretamente, as Goroutines não responderão ao cancelamento.**  

✅ **Sempre passe `ctx` como primeiro argumento das funções concorrentes.**  

```go
func process(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("Processo finalizado!")
    }
}
```

---

## **12.4.6 Boas Práticas**

✔ **Use `context.WithTimeout` para garantir que tarefas não excedam um tempo máximo aceitável.**  
✔ **Escolha `WithTimeout` quando o tempo for relativo ao início e `WithDeadline` para tempos fixos.**  
✔ **Sempre propague `ctx` para funções concorrentes para um cancelamento eficiente.**  
✔ **Defina tempos realistas para evitar falhas inesperadas.**  
✔ **Ao lidar com APIs, bancos de dados e chamadas remotas, `context` é essencial para evitar travamentos.**  

---

## **Conclusão**

O **`context.WithTimeout`** fornece um controle eficiente sobre **o tempo de execução de Goroutines**, garantindo que tarefas concorrentes não rodem por mais tempo que o necessário.  
No próximo capítulo, exploraremos **boas práticas para otimizar o uso de contextos e evitar armadilhas comuns!** 🚀
