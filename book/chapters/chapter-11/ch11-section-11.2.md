# **11.2 `sync.Cond`: Sincronização Baseada em Eventos**

Enquanto `sync.Mutex` e `sync.RWMutex` são usados para **exclusão mútua**, o pacote `sync` também fornece **`sync.Cond`**, que permite sincronizar Goroutines **com base em eventos**.

Nesta seção, exploraremos:

- O que é `sync.Cond` e como funciona
- Diferença entre `sync.Cond` e `sync.Mutex`
- Uso de `sync.Cond` para coordenação de Goroutines
- Estratégias eficientes para evitar espera ativa (busy-waiting)
- Comparação com outras técnicas de sincronização

---

## **11.2.1 O Que é `sync.Cond`?**

`sync.Cond` é um mecanismo que permite que **Goroutines aguardem notificações de eventos**.  
Ele resolve um problema comum em programação concorrente: **como fazer uma Goroutine esperar uma condição específica sem desperdiçar CPU?**

📌 **Enquanto `sync.Mutex` protege seções críticas, `sync.Cond` permite que Goroutines esperem até que um evento aconteça.**

✅ **Fluxo de `sync.Cond`:**

1. Uma Goroutine **aguarda** uma condição ser satisfeita (`Wait()`).
2. Outra Goroutine **sinaliza** (`Signal()`) ou **notifica todas** (`Broadcast()`) quando a condição mudar.
3. A Goroutine despertada reavalia a condição e prossegue se estiver correta.

---

## **11.2.2 Como Criar um `sync.Cond`?**

Criamos um `sync.Cond` usando um `sync.Mutex`:

```go
import "sync"

var cond = sync.NewCond(&sync.Mutex{})
```

📌 **O `sync.Mutex` é obrigatório, pois `sync.Cond` depende de um bloqueio para garantir sincronização segura.**

✅ **Exemplo básico:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var cond = sync.NewCond(&sync.Mutex{})
var pronto = false

func esperarEvento() {
    cond.L.Lock() // Bloqueia antes de aguardar
    for !pronto {
        cond.Wait() // Aguarda o sinal
    }
    fmt.Println("Evento recebido!")
    cond.L.Unlock() // Libera o bloqueio
}

func dispararEvento() {
    time.Sleep(time.Second)
    cond.L.Lock()
    pronto = true
    cond.Signal() // Desperta uma Goroutine
    cond.L.Unlock()
}

func main() {
    go esperarEvento()
    go dispararEvento()

    time.Sleep(2 * time.Second)
}
```

📌 **`cond.Wait()` bloqueia até que `cond.Signal()` ou `cond.Broadcast()` seja chamado.**  
📌 **A verificação `for !pronto` garante que o evento ainda é válido após ser acordado.**  

✅ **Sem `for !pronto`, a Goroutine poderia ser despertada sem que a condição fosse verdadeira (falsa ativação).**

---

## **11.2.3 Diferença Entre `sync.Cond`, `sync.Mutex` e `sync.WaitGroup`**

| Técnica | Uso Principal | Bloqueante? | Melhor Aplicação |
|---------|--------------|------------|-----------------|
| `sync.Mutex` | Proteção de recursos compartilhados | ✅ Sim | Controle de acesso |
| `sync.WaitGroup` | Aguardar Goroutines finalizarem | ✅ Sim | Execução concorrente |
| `sync.Cond` | Sincronização por eventos | ✅ Sim | Espera condicional |

📌 **Use `sync.Cond` quando precisar aguardar um evento específico antes de continuar a execução.**

✅ **Exemplo prático: Um sistema de fila de tarefas**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var cond = sync.NewCond(&sync.Mutex{})
var fila []int

func produtor() {
    for i := 1; i <= 5; i++ {
        cond.L.Lock()
        fila = append(fila, i)
        fmt.Println("Produziu:", i)
        cond.Signal() // Notifica o consumidor
        cond.L.Unlock()
        time.Sleep(time.Second)
    }
}

func consumidor() {
    for i := 1; i <= 5; i++ {
        cond.L.Lock()
        for len(fila) == 0 {
            cond.Wait() // Aguarda novos itens
        }
        item := fila[0]
        fila = fila[1:]
        fmt.Println("Consumiu:", item)
        cond.L.Unlock()
    }
}

func main() {
    go consumidor()
    go produtor()

    time.Sleep(6 * time.Second)
}
```

📌 **O consumidor espera por novas tarefas sem desperdiçar CPU.**  
📌 **O produtor adiciona itens e notifica o consumidor via `Signal()`.**  

✅ **Essa abordagem evita o uso de polling ativo (busy-waiting), tornando o sistema mais eficiente.**

---

## **11.2.4 `Signal()` vs. `Broadcast()`**

- **`Signal()`** → Desperta **uma única** Goroutine esperando em `Wait()`.
- **`Broadcast()`** → Desperta **todas** as Goroutines esperando em `Wait()`.

✅ **Quando usar `Broadcast()`?**  
Quando várias Goroutines precisam ser notificadas ao mesmo tempo.

```go
cond.Broadcast() // Desperta todas as Goroutines esperando o evento
```

✅ **Quando usar `Signal()`?**  
Quando apenas **uma** Goroutine precisa ser notificada.

```go
cond.Signal() // Notifica uma Goroutine aleatória esperando o evento
```

📌 **Se várias Goroutines esperam pelo mesmo evento, `Broadcast()` pode ser mais eficiente.**  

---

## **11.2.5 Erros Comuns ao Usar `sync.Cond`**

❌ **Chamar `Wait()` sem antes bloquear com `Lock()`**

```go
cond.Wait() // ERRO: Deve estar dentro de cond.L.Lock() e cond.L.Unlock()
```

✅ **Sempre envolva `Wait()` dentro de um `Lock() / Unlock()`**

```go
cond.L.Lock()
cond.Wait()
cond.L.Unlock()
```

---

❌ **Esquecer de verificar a condição dentro de um loop**

```go
if !pronto { // ERRO: Pode causar falsa ativação
    cond.Wait()
}
```

✅ **Use um `for` para verificar a condição repetidamente**

```go
for !pronto {
    cond.Wait()
}
```

📌 **Isso protege contra "spurious wakeups" (acordar sem motivo real).**  

---

## **11.2.6 Boas Práticas**

✔ **Use `sync.Cond` quando precisar aguardar um evento antes de continuar.**  
✔ **Sempre use `Signal()` para acordar uma única Goroutine e `Broadcast()` para todas.**  
✔ **Evite busy-waiting utilizando `Wait()` corretamente.**  
✔ **Certifique-se de envolver `Wait()` dentro de um `for`, nunca um `if`.**  
✔ **Use `sync.Cond` com `sync.Mutex` para evitar condições de corrida.**  

---

## **Conclusão**

O **`sync.Cond`** é um mecanismo poderoso para sincronização baseada em eventos, evitando busy-waiting e garantindo eficiência na comunicação entre Goroutines.  
No próximo capítulo, exploraremos **`sync.Once`**, um recurso essencial para inicializações seguras e eficientes em Go! 🚀
