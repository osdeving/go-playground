# **12.1 O Pacote `context`**

O **pacote `context`** foi introduzido no Go para fornecer **controle eficiente sobre o tempo de vida de Goroutines** e permitir **propagação de cancelamento e deadlines**.  
Ele resolve um problema crítico em aplicações concorrentes: **como interromper Goroutines de forma segura e evitar vazamentos de memória**?

Nesta seção, exploraremos:

- O que é o `context` e por que ele é essencial em Go
- Como `context` é propagado entre Goroutines
- Estrutura do `context.Context` e seus principais métodos
- Diferença entre `context.Background()` e `context.TODO()`
- Comparação entre `context` e outras técnicas de controle concorrente
- Boas práticas e erros comuns ao utilizá-lo

---

## **12.1.1 O Que é `context` e Por Que Ele É Necessário?**

Sem `context`, a única maneira de cancelar uma Goroutine seria usar **channels** ou **variáveis globais**, o que pode ser propenso a **vazamentos de Goroutines**.

✅ **Exemplo problemático: Goroutine que nunca é cancelada**

```go
package main

import (
    "fmt"
    "time"
)

func worker(stop chan bool) {
    for {
        select {
        case <-stop:
            fmt.Println("Worker finalizado!")
            return
        default:
            fmt.Println("Trabalhando...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    stop := make(chan bool)
    go worker(stop)

    time.Sleep(2 * time.Second)
    stop <- true // Cancela o worker
}
```

📌 **Esse código funciona, mas não é escalável**: se houver muitas Goroutines, precisaremos gerenciar muitos channels.  

✅ **Solução com `context`:**

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
    cancel() // Cancela a Goroutine
}
```

📌 **Agora podemos gerenciar o cancelamento de forma centralizada.**  
📌 **Todas as Goroutines que recebem `ctx` sabem quando devem ser encerradas.**  

✅ **Isso evita vazamento de Goroutines e facilita o controle de concorrência.**  

---

## **12.1.2 Como `context` É Propagado?**

O `context` é **passado como argumento para funções concorrentes**, garantindo que toda a hierarquia de Goroutines possa responder ao cancelamento.

✅ **Fluxo de propagação de `context`**:

```
Main Goroutine -----> Goroutine 1 -----> Goroutine 2
           (context)          (context)          (context)
```

📌 **Se a Goroutine principal cancelar o `context`, todas as Goroutines filhas também serão encerradas.**  

✅ **Exemplo de propagação:**

```go
func process(ctx context.Context) {
    go subProcess(ctx) // Propaga o mesmo contexto
}

func subProcess(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("Subprocesso cancelado!")
    }
}
```

📌 **`ctx.Done()` é um canal fechado quando o contexto é cancelado.**  
📌 **Isso permite encadear cancelamentos de forma automática.**  

✅ **Essa abordagem é essencial para aplicações distribuídas e serviços HTTP.**  

---

## **12.1.3 Estrutura do `context.Context`**

O `context.Context` é uma interface com os seguintes métodos:

| Método | Descrição |
|--------|------------|
| `Done()` | Retorna um canal fechado quando o contexto for cancelado |
| `Err()` | Retorna um erro indicando o motivo do cancelamento |
| `Deadline()` | Retorna o deadline configurado, se houver |
| `Value(key interface{})` | Recupera um valor associado ao contexto |

✅ **Exemplo de uso do `Err()` para verificar cancelamento:**

```go
ctx, cancel := context.WithCancel(context.Background())
cancel()

fmt.Println(ctx.Err()) // context canceled
```

📌 **Isso evita que Goroutines continuem executando código após o cancelamento.**  

---

## **12.1.4 `context.Background()` vs. `context.TODO()`**

O Go fornece dois contextos iniciais que podem ser utilizados:

| Função | Uso Principal |
|--------|--------------|
| `context.Background()` | Contexto base padrão |
| `context.TODO()` | Indica que o contexto ainda não foi decidido |

✅ **Quando usar `context.Background()`?**  
- Para iniciar um contexto raiz em aplicações.  
- Em programas que não precisam de propagação de contexto.  

✅ **Quando usar `context.TODO()`?**  
- Em código onde o contexto será definido no futuro.  
- Durante o desenvolvimento para indicar dependências pendentes.  

📌 **`context.TODO()` é útil para refatoração e transição de código.**  

---

## **12.1.5 Comparação: `context` vs. Outras Técnicas**

| Técnica | Uso Principal | Suporte a Propagação? | Gerenciado Automaticamente? |
|---------|--------------|-----------------|----------------|
| `context` | Cancelamento e tempo de vida de Goroutines | ✅ Sim | ✅ Sim |
| Channels | Comunicação entre Goroutines | ❌ Não | ❌ Não |
| Variáveis globais | Controle manual | ❌ Não | ❌ Não |

📌 **O `context` fornece um mecanismo escalável e eficiente para controle de Goroutines.**  

✅ **Em aplicações HTTP e RPC, `context` é essencial para evitar requisições pendentes indefinidamente.**  

---

## **12.1.6 Boas Práticas**

✔ **Sempre passe `context.Context` como primeiro argumento de funções concorrentes.**  
✔ **Nunca armazene `context.Context` dentro de structs (ele deve ser transitório).**  
✔ **Use `ctx.Done()` para detectar cancelamentos e evitar vazamentos de Goroutines.**  
✔ **Prefira `context.Background()` para criar contextos iniciais e `context.TODO()` para refatorações.**  
✔ **Evite usar `context` para compartilhar dados — prefira channels ou variáveis seguras.**  

---

## **Conclusão**

O **pacote `context`** é um dos recursos mais poderosos do Go para **controle de Goroutines e propagação de cancelamento**.  
No próximo capítulo, exploraremos **`context.WithCancel`**, um método essencial para criar contextos dinâmicos e encadear cancelamentos eficientes! 🚀
