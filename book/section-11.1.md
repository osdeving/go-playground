# **11.1 Mutexes (`sync.Mutex`, `sync.RWMutex`)**

A sincronização de acesso a recursos compartilhados é um desafio comum na programação concorrente.  
Go oferece mecanismos eficientes para evitar **condições de corrida** e garantir **consistência de dados**, sendo os **Mutexes (`sync.Mutex`)** uma das ferramentas fundamentais.

Nesta seção, exploraremos:

- O que é um Mutex e quando usá-lo
- Diferença entre `sync.Mutex` e `sync.RWMutex`
- Erros comuns ao usar Mutexes e como evitá-los
- Comparação com outras técnicas de sincronização
- Melhores práticas para uso eficiente

---

## **11.1.1 O Que é um Mutex (`sync.Mutex`)?**

Um **Mutex (Mutual Exclusion)** é um bloqueio que garante que apenas **uma Goroutine** pode acessar um recurso de cada vez.

✅ **Exemplo de problema sem Mutex:**

```go
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

📌 **Múltiplas Goroutines acessam `contador` simultaneamente, causando inconsistência.**

✅ **Usando `sync.Mutex` para garantir segurança:**

```go
import "sync"

var contador int
var mutex sync.Mutex

func incrementar() {
    for i := 0; i < 1000; i++ {
        mutex.Lock()   // Bloqueia o acesso ao contador
        contador++
        mutex.Unlock() // Libera o acesso ao contador
    }
}
```

📌 **Agora, apenas uma Goroutine pode modificar `contador` por vez.**

---

## **11.1.2 O Que é `sync.RWMutex`?**

O **`sync.RWMutex`** é uma versão otimizada do `Mutex` que permite:

- **Múltiplas leituras simultâneas (`RLock`)**
- **Escrita exclusiva (`Lock`)**

✅ **Uso eficiente do `sync.RWMutex`:**

```go
var dados string
var mutex sync.RWMutex

func leitor() {
    mutex.RLock()  // Permite múltiplas leituras simultâneas
    fmt.Println("Lendo:", dados)
    mutex.RUnlock()
}

func escritor(novoValor string) {
    mutex.Lock()  // Bloqueia todas as leituras e escritas
    dados = novoValor
    mutex.Unlock()
}
```

📌 **Use `sync.RWMutex` quando houver mais operações de leitura do que escrita!**

---

## **11.1.3 Erros Comuns ao Usar Mutexes**

❌ **Esquecer de liberar o Mutex (`Unlock`)**

```go
mutex.Lock()
contador++
// mutex.Unlock()  // ERRO: Mutex nunca liberado! Deadlock!
```

✅ **Sempre use `defer` para garantir que o Mutex será liberado:**

```go
mutex.Lock()
defer mutex.Unlock()
contador++
```

---

❌ **Chamar `Unlock` sem `Lock` anterior**

```go
var mutex sync.Mutex

mutex.Unlock() // ERRO: Fatal error - Unlock sem Lock!
```

📌 **Nunca chame `Unlock()` sem antes ter chamado `Lock()`.**

✅ **Certifique-se de que o Mutex sempre será adquirido antes da liberação.**

---

## **11.1.4 Comparação: `Mutex` vs. Outras Técnicas de Sincronização**

| Técnica | Uso Principal | Bloqueante? | Performance |
|---------|--------------|------------|------------|
| `sync.Mutex` | Proteção de dados compartilhados | ✅ Sim | ⚡ Alta |
| `sync.RWMutex` | Múltiplas leituras simultâneas | ✅ Sim | ⚡ Muito alta |
| `sync.WaitGroup` | Aguardar Goroutines | ✅ Sim | ⚡ Alta |
| `chan` (Channels) | Comunicação entre Goroutines | ✅ Sim | ⚡ Média |

📌 **Use `Mutex` para acessar recursos compartilhados, `sync.WaitGroup` para esperar Goroutines e Channels para comunicação concorrente.**

---

## **11.1.5 Boas Práticas**

✔ **Use `sync.Mutex` apenas quando necessário — Channels podem ser uma opção melhor.**  
✔ **Prefira `sync.RWMutex` quando houver muitas leituras e poucas escritas.**  
✔ **Sempre use `defer mutex.Unlock()` para evitar deadlocks.**  
✔ **Evite manter o Mutex bloqueado por muito tempo para reduzir contenção.**  

---

## **Conclusão**

Os **Mutexes (`sync.Mutex`, `sync.RWMutex`)** são essenciais para proteger recursos compartilhados em Go.  
No próximo capítulo, exploraremos **`sync.Cond`**, uma ferramenta poderosa para **sincronização baseada em eventos!** 🚀
