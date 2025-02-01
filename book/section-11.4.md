# **11.4 `sync/atomic`: Operações Atômicas e Segurança de Memória**

A manipulação de variáveis compartilhadas em ambientes concorrentes pode levar a **condições de corrida**.  
Quando `sync.Mutex` e `sync.RWMutex` são opções pesadas, podemos recorrer ao **pacote `sync/atomic`**, que permite manipular variáveis **de forma segura e sem bloqueios**.

Nesta seção, exploraremos:

- O que é `sync/atomic` e como funciona
- Diferença entre `sync/atomic` e `sync.Mutex`
- Operações atômicas disponíveis em Go
- Casos de uso para otimizar concorrência
- Melhores práticas e erros comuns ao utilizar `sync/atomic`

---

## **11.4.1 O Que é `sync/atomic`?**

O pacote `sync/atomic` fornece **operações atômicas** que garantem que leituras e escritas em variáveis compartilhadas sejam **indivisíveis**,  
ou seja, não podem ser interrompidas por outras Goroutines durante a execução.

✅ **Exemplo de condição de corrida sem `sync/atomic`:**

```go
package main

import (
    "fmt"
    "time"
)

var contador int

func incrementar() {
    for i := 0; i < 1000; i++ {
        contador++ // Condição de corrida!
    }
}

func main() {
    go incrementar()
    go incrementar()

    time.Sleep(time.Second)
    fmt.Println("Contador:", contador) // Resultado imprevisível!
}
```

📌 **Duas Goroutines podem modificar `contador` simultaneamente, gerando um resultado incorreto.**  

✅ **Corrigindo com `sync/atomic`:**

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

var contador int64

func incrementar() {
    for i := 0; i < 1000; i++ {
        atomic.AddInt64(&contador, 1) // Operação atômica segura
    }
}

func main() {
    go incrementar()
    go incrementar()

    time.Sleep(time.Second)
    fmt.Println("Contador:", atomic.LoadInt64(&contador)) // Sempre correto!
}
```

📌 **Agora, `contador` é atualizado de forma segura, sem condições de corrida.**  

✅ **As operações atômicas garantem que as variáveis não sejam corrompidas por concorrência.**  

---

## **11.4.2 `sync/atomic` vs. `sync.Mutex`**

| Característica | `sync/atomic` | `sync.Mutex` |
|---------------|--------------|-------------|
| **Bloqueia outras Goroutines?** | ❌ Não | ✅ Sim |
| **Performance** | ⚡ Alta | 🐢 Média |
| **Uso de CPU** | ✅ Baixo | ❌ Pode causar contenção |
| **Complexidade** | ✅ Simples | ❌ Maior |
| **Ideal para...** | Contadores, flags | Estruturas complexas |

📌 **Use `sync/atomic` para operações simples (contadores, flags, indicadores de status).**  
📌 **Use `sync.Mutex` para proteger dados mais complexos (structs, listas encadeadas).**  

✅ **Se precisar modificar um único valor numérico, `sync/atomic` é mais rápido!**  

---

## **11.4.3 Principais Funções do `sync/atomic`**

O pacote `sync/atomic` oferece funções para manipulação atômica de inteiros, ponteiros e booleanos.

| Função | Descrição |
|--------|------------|
| `atomic.AddInt64(&x, n)` | Incrementa `x` de forma atômica |
| `atomic.LoadInt64(&x)` | Lê `x` de forma segura |
| `atomic.StoreInt64(&x, n)` | Define `x` para `n` de forma atômica |
| `atomic.CompareAndSwapInt64(&x, old, new)` | Atualiza `x` se `x == old` |

✅ **Exemplo: Contador seguro com `sync/atomic`**

```go
package main

import (
    "fmt"
    "sync/atomic"
)

var contador int64

func incrementar() {
    atomic.AddInt64(&contador, 1)
}

func main() {
    incrementar()
    fmt.Println("Valor do contador:", atomic.LoadInt64(&contador))
}
```

📌 **O `atomic.LoadInt64()` garante que a leitura seja consistente.**  

---

## **11.4.4 Compare-And-Swap (CAS) com `sync/atomic`**

O **Compare-And-Swap (CAS)** é um mecanismo eficiente para atualização de valores sem bloqueios.

✅ **Exemplo de `CompareAndSwapInt64()`:**

```go
var status int64

func atualizarStatus(novoStatus int64) {
    if atomic.CompareAndSwapInt64(&status, 0, novoStatus) {
        fmt.Println("Status atualizado com sucesso!")
    } else {
        fmt.Println("Já foi atualizado!")
    }
}
```

📌 **Se `status` for `0`, ele será atualizado para `novoStatus`.**  
📌 **Se `status` já foi alterado, a função falha sem modificar nada.**  

✅ **Isso evita operações duplicadas e melhora a performance sem precisar de locks.**  

---

## **11.4.5 Erros Comuns ao Usar `sync/atomic`**

❌ **Usar `sync/atomic` em estruturas complexas**

```go
var dados map[string]int
atomic.AddInt64(&dados["chave"], 1) // ERRO: `sync/atomic` só funciona com inteiros, ponteiros e booleanos!
```

✅ **Para estruturas de dados, use `sync.Mutex`.**  

---

❌ **Achar que `sync/atomic` substitui Mutexes completamente**

```go
type Conta struct {
    saldo int64
}

func depositar(c *Conta, valor int64) {
    atomic.AddInt64(&c.saldo, valor) // ERRO: Pode haver inconsistências na struct!
}
```

✅ **Se precisar modificar múltiplos campos de uma struct, use `sync.Mutex`.**  

---

## **11.4.6 Boas Práticas**

✔ **Use `sync/atomic` apenas para valores numéricos ou flags booleanas.**  
✔ **Para operações mais complexas, `sync.Mutex` pode ser necessário.**  
✔ **Utilize `CompareAndSwap()` para evitar operações concorrentes duplicadas.**  
✔ **Evite usar `sync/atomic` com estruturas de dados não suportadas.**  
✔ **Use `atomic.Load()` para garantir leituras consistentes em variáveis compartilhadas.**  

---

## **Conclusão**

O **pacote `sync/atomic`** fornece operações atômicas eficientes para manipulação segura de variáveis concorrentes sem bloqueios.  
No próximo capítulo, exploraremos **`sync.Pool`**, um recurso avançado para gerenciamento eficiente de alocação de memória! 🚀
