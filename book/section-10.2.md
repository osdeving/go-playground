# **10.2 `sync.WaitGroup`**

Em Go, as **Goroutines** são executadas de forma independente, o que pode levar a situações onde o programa principal encerra antes que todas as Goroutines tenham finalizado.  
Para gerenciar essa execução, usamos **`sync.WaitGroup`**, uma estrutura essencial para sincronização concorrente.

Nesta seção, exploraremos:

- O que é `sync.WaitGroup` e quando usá-lo
- Como garantir que todas as Goroutines finalizem corretamente
- Diferenças entre `sync.WaitGroup` e outras abordagens de sincronização
- Cuidados ao usar `sync.WaitGroup`
- Comparação com `Mutex` e `Channels`

---

## **10.2.1 O Que é `sync.WaitGroup`?**

O **`sync.WaitGroup`** é um contador que permite aguardar a finalização de múltiplas Goroutines antes de prosseguir com a execução do código.

✅ **Sem `sync.WaitGroup`, Goroutines podem não executar completamente:**

```go
package main

import (
    "fmt"
    "time"
)

func rotina() {
    fmt.Println("Executando Goroutine")
}

func main() {
    go rotina()
    fmt.Println("Fim do programa")
}
```

📌 **O programa pode encerrar antes que `rotina()` seja executada!**

✅ **Usando `sync.WaitGroup` garantimos que todas as Goroutines terminem antes do `main` encerrar:**

```go
package main

import (
    "fmt"
    "sync"
)

func rotina(wg *sync.WaitGroup) {
    defer wg.Done() // Decrementa o contador ao finalizar
    fmt.Println("Executando Goroutine")
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // Incrementa o contador

    go rotina(&wg)

    wg.Wait() // Aguarda todas as Goroutines finalizarem
    fmt.Println("Fim do programa")
}
```

📌 **Agora o programa espera até que `rotina()` seja concluída antes de encerrar.**

---

## **10.2.2 Como `sync.WaitGroup` Funciona?**

O `sync.WaitGroup` possui **três operações principais**:

| Método | Descrição |
|--------|-----------|
| `Add(n)` | Adiciona `n` ao contador (indica quantas Goroutines devem finalizar) |
| `Done()` | Decrementa o contador (indica que uma Goroutine finalizou) |
| `Wait()` | Bloqueia a execução até que o contador chegue a zero |

✅ **Fluxo básico:**

1️⃣ Chamamos `wg.Add(1)` antes de iniciar cada Goroutine.  
2️⃣ Cada Goroutine chama `wg.Done()` ao finalizar.  
3️⃣ O programa principal usa `wg.Wait()` para aguardar todas as Goroutines.  

---

## **10.2.3 Sincronizando Múltiplas Goroutines**

Podemos usar `sync.WaitGroup` para sincronizar **várias Goroutines**:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func rotina(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Second)
    fmt.Println("Goroutine", id, "finalizou")
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go rotina(i, &wg)
    }

    wg.Wait()
    fmt.Println("Todas as Goroutines finalizaram")
}
```

📌 **O programa aguarda todas as 5 Goroutines finalizarem antes de continuar.**

---

## **10.2.4 Erros Comuns ao Usar `sync.WaitGroup`**

❌ **Esquecer `wg.Add(n)` antes de iniciar as Goroutines**

```go
var wg sync.WaitGroup

go func() {
    wg.Done() // ERRO: wg.Add() nunca foi chamado!
}()

wg.Wait() // Deadlock!
```

📌 **O programa entra em `deadlock` pois `wg.Wait()` nunca é liberado.**

✅ **Sempre chame `wg.Add(n)` antes de iniciar Goroutines!**

---

❌ **Chamar `wg.Done()` mais vezes do que `wg.Add()`**

```go
wg.Add(1)
wg.Done()
wg.Done() // ERRO: Decremento além do limite!
```

📌 **Isso causa um erro fatal de runtime!**

✅ **Garanta que `wg.Done()` seja chamado exatamente `n` vezes.**

---

## **10.2.5 Comparação com Outras Técnicas de Sincronização**

| Técnica | Uso Principal | Quando Usar |
|---------|--------------|-------------|
| `sync.WaitGroup` | Aguardar Goroutines | Quando sabemos quantas Goroutines precisam finalizar |
| `sync.Mutex` | Evitar condições de corrida | Quando múltiplas Goroutines acessam um recurso compartilhado |
| `Channels` | Comunicação concorrente | Quando precisamos enviar e receber dados entre Goroutines |

📌 **`sync.WaitGroup` é ideal para aguardar execuções concorrentes, mas não substitui `Mutex` ou `Channels`.**

✅ **Se precisamos sincronizar acesso a variáveis, `sync.Mutex` pode ser mais apropriado.**

---

## **10.2.6 Boas Práticas**

✔ **Sempre chame `wg.Add(n)` antes de iniciar Goroutines.**  
✔ **Use `defer wg.Done()` para garantir que `Done()` sempre seja chamado.**  
✔ **Evite chamar `wg.Wait()` dentro de uma Goroutine — isso pode causar `deadlock`.**  
✔ **Para cenários complexos, combine `WaitGroup` com `Channels` para maior controle.**  

---

## **Conclusão**

O **`sync.WaitGroup`** é uma ferramenta essencial para gerenciar concorrência em Go.  
No próximo capítulo, exploraremos **`Channels`**, a principal forma de comunicação segura entre Goroutines! 🚀
