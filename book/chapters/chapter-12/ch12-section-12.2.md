# **12.2 `context.WithCancel`: Cancelamento de Goroutines**

O **`context.WithCancel`** é uma das formas mais simples de criar um **contexto cancelável** em Go.  
Ele permite que um **contexto pai** crie um **contexto filho**, que pode ser **cancelado dinamicamente**, interrompendo todas as Goroutines associadas a ele.

Nesta seção, exploraremos:

- O que é `context.WithCancel` e como funciona
- Cancelamento hierárquico de Goroutines
- Uso prático em sistemas concorrentes
- Erros comuns e como evitá-los
- Comparação com outras abordagens de cancelamento

---

## **12.2.1 O Que é `context.WithCancel`?**

O `context.WithCancel` permite criar um contexto que pode ser **cancelado manualmente** através da função `cancel()`.  
Isso garante que todas as Goroutines que compartilham esse contexto possam ser **finalizadas corretamente**, evitando **vazamento de memória** e **execuções desnecessárias**.

✅ **Exemplo básico de `context.WithCancel`**

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
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx)

    time.Sleep(2 * time.Second)
    cancel() // Cancela todas as Goroutines associadas ao contexto

    time.Sleep(time.Second) // Tempo extra para visualizar o cancelamento
}
```

📌 **Quando `cancel()` é chamado, todas as Goroutines ouvindo `ctx.Done()` são finalizadas.**  

✅ **Isso evita vazamentos e melhora a eficiência da aplicação.**  

---

## **12.2.2 Cancelamento Hierárquico de Goroutines**

O `context.WithCancel` permite que um **contexto pai gere vários contextos filhos**.  
Quando o pai é cancelado, **todos os filhos também são automaticamente cancelados**.

✅ **Exemplo de cancelamento encadeado:**

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func process(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Processo %d finalizado!
", id)
            return
        default:
            fmt.Printf("Processo %d rodando...
", id)
            time.Sleep(time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go process(ctx, 1)
    go process(ctx, 2)
    go process(ctx, 3)

    time.Sleep(3 * time.Second)
    cancel() // Cancela todos os processos

    time.Sleep(1 * time.Second) // Tempo para visualizar a finalização
}
```

📌 **Todas as Goroutines são encerradas automaticamente quando `cancel()` é chamado.**  
📌 **Isso evita que processos concorrentes fiquem rodando indefinidamente.**  

✅ **Esse padrão é amplamente utilizado em servidores web e sistemas distribuídos.**  

---

## **12.2.3 Erros Comuns ao Usar `context.WithCancel`**

❌ **Esquecer de chamar `cancel()`**

```go
ctx, _ := context.WithCancel(context.Background()) // ERRO: `cancel()` nunca é chamado!
```

📌 **Se `cancel()` não for chamado, Goroutines associadas ao contexto podem nunca ser finalizadas.**  

✅ **Sempre chame `cancel()` para evitar vazamento de Goroutines!**  

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Garante que `cancel()` será chamado
```

---

❌ **Chamar `cancel()` antes das Goroutines iniciarem**

```go
ctx, cancel := context.WithCancel(context.Background())
cancel() // Cancela imediatamente antes de qualquer Goroutine rodar

go worker(ctx) // Nunca será executado corretamente!
```

📌 **Se `cancel()` for chamado cedo demais, as Goroutines nem chegarão a rodar.**  

✅ **Garanta que `cancel()` só seja chamado no momento apropriado.**  

---

## **12.2.4 Comparação: `context.WithCancel` vs. Outras Técnicas**

| Técnica | Propaga Cancelamento? | Melhoria na Eficiência? | Uso Principal |
|---------|------------------|-----------------|--------------|
| `context.WithCancel` | ✅ Sim | ✅ Sim | Cancelamento de Goroutines |
| `sync.WaitGroup` | ❌ Não | ❌ Não | Aguardar Goroutines finalizarem |
| Channels | ⚠️ Parcial | ⚠️ Média | Comunicação entre Goroutines |
| Variáveis Globais | ❌ Não | ❌ Não | Controle de execução manual |

📌 **`context.WithCancel` é a abordagem mais escalável para cancelamento concorrente.**  

✅ **Use `sync.WaitGroup` quando apenas precisar aguardar finalização, sem cancelamento antecipado.**  

---

## **12.2.5 Boas Práticas**

✔ **Sempre passe `context.Context` como primeiro argumento de funções concorrentes.**  
✔ **Use `ctx.Done()` para detectar cancelamentos de forma eficiente.**  
✔ **Sempre chame `cancel()` para evitar vazamento de Goroutines.**  
✔ **Prefira `context.WithCancel` em sistemas onde o cancelamento precisa ser propagado.**  
✔ **Combine `sync.WaitGroup` com `context.WithCancel` quando precisar aguardar a finalização de múltiplas Goroutines.**  

---

## **Conclusão**

O **`context.WithCancel`** é um mecanismo essencial para **cancelamento eficiente de Goroutines** e controle concorrente.  
No próximo capítulo, exploraremos **`context.WithDeadline`**, que adiciona um limite de tempo para execução de Goroutines! 🚀
