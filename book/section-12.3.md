# **12.3 `context.WithDeadline`: Controle de Tempo de Execução**

O **`context.WithDeadline`** permite definir um **tempo limite absoluto** para a execução de uma Goroutine.  
Isso é fundamental para evitar **tarefas bloqueadas indefinidamente** e garantir que operações concorrentes **não ultrapassem um tempo máximo aceitável**.

Nesta seção, exploraremos:

- O que é `context.WithDeadline` e como funciona
- Diferença entre `WithDeadline` e `WithTimeout`
- Uso prático para evitar Goroutines bloqueadas
- Cancelamento automático baseado em tempo
- Boas práticas e erros comuns

---

## **12.3.1 O Que é `context.WithDeadline`?**

O `context.WithDeadline` cria um contexto que **expira automaticamente em um tempo absoluto predefinido**.  
Isso significa que, **independentemente do que estiver acontecendo**, o contexto será cancelado no momento exato especificado.

✅ **Exemplo básico de `context.WithDeadline`**

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
    deadline := time.Now().Add(3 * time.Second) // Define o tempo limite absoluto
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel() // Cancela o contexto após o deadline

    go worker(ctx)

    time.Sleep(4 * time.Second) // Aguarda para visualizar o cancelamento
}
```

📌 **A Goroutine será finalizada exatamente após 3 segundos.**  
📌 **Não importa se o processamento ainda não terminou, o contexto será cancelado automaticamente.**  

✅ **Isso garante que processos longos não fiquem rodando além do tempo esperado.**  

---

## **12.3.2 Diferença Entre `WithDeadline` e `WithTimeout`**

Ambos os métodos fornecem cancelamento baseado em tempo, mas de formas diferentes:

| Método | O que faz? | Melhor Aplicação |
|--------|------------|----------------|
| `context.WithDeadline(ctx, time)` | Cancela no tempo exato definido | Quando há um horário absoluto para expiração |
| `context.WithTimeout(ctx, duration)` | Cancela após um tempo relativo | Quando um tempo máximo de execução é definido |

✅ **Use `WithDeadline` quando o cancelamento for baseado em um tempo específico.**  
✅ **Use `WithTimeout` quando o cancelamento for relativo a quando começou.**  

✅ **Exemplo comparativo:**

```go
deadline := time.Now().Add(5 * time.Second)
ctx1, _ := context.WithDeadline(context.Background(), deadline) // Expira às 15:05:30

ctx2, _ := context.WithTimeout(context.Background(), 5*time.Second) // Expira 5s após a criação
```

📌 **A escolha entre `WithDeadline` e `WithTimeout` depende do cenário da aplicação.**  

---

## **12.3.3 Aplicação Prática: Cancelamento de Requisições HTTP**

Em aplicações web, `context.WithDeadline` é extremamente útil para **evitar requisições demoradas**.

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
    deadline := time.Now().Add(2 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
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
📌 **Isso evita bloqueios indesejados em APIs e melhora a experiência do usuário.**  

✅ **Esse padrão é amplamente utilizado em servidores web e microservices.**  

---

## **12.3.4 Cancelamento Automático com `WithDeadline`**

Uma vantagem do `WithDeadline` é que **não precisamos chamar `cancel()` manualmente**, pois ele **se cancela automaticamente ao atingir o tempo limite**.

✅ **Exemplo de cancelamento automático:**

```go
deadline := time.Now().Add(3 * time.Second)
ctx, _ := context.WithDeadline(context.Background(), deadline) // Sem necessidade de chamar cancel()
```

📌 **Se o tempo for atingido, `ctx.Done()` será fechado automaticamente.**  

✅ **Isso reduz a complexidade e evita esquecimentos no código.**  

---

## **12.3.5 Erros Comuns ao Usar `context.WithDeadline`**

❌ **Definir prazos muito curtos sem necessidade**

```go
deadline := time.Now().Add(50 * time.Millisecond) // ERRO: Pode cancelar antes da tarefa terminar!
```

📌 **Se o deadline for muito curto, pode causar cancelamentos prematuros.**  
📌 **Ajuste o tempo conforme a necessidade do processamento.**  

✅ **Defina tempos realistas para evitar falhas desnecessárias.**  

---

❌ **Achar que `WithDeadline` substitui `WithCancel` completamente**

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
cancel() // Cancela imediatamente!
```

📌 **Se `cancel()` for chamado antes do tempo, o contexto será cancelado antes do deadline.**  

✅ **Use `WithCancel` para cancelamentos manuais e `WithDeadline` para cancelamentos automáticos.**  

---

## **12.3.6 Boas Práticas**

✔ **Use `context.WithDeadline` quando precisar de um cancelamento baseado em tempo absoluto.**  
✔ **Ajuste os deadlines com valores realistas para evitar cancelamentos prematuros.**  
✔ **Sempre propague `ctx` para funções concorrentes para um controle eficiente.**  
✔ **Combine `context.WithDeadline` com `context.WithTimeout` quando necessário.**  
✔ **Para evitar requisições bloqueadas, sempre use `context` ao lidar com HTTP, DBs e RPCs.**  

---

## **Conclusão**

O **`context.WithDeadline`** é um recurso essencial para **garantir que Goroutines não rodem por mais tempo que o permitido**.  
No próximo capítulo, exploraremos **`context.WithTimeout`**, que fornece uma abordagem mais flexível para cancelamento baseado em tempo relativo! 🚀
