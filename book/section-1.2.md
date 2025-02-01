# 🎯 **1.2 Filosofia do Go**

A filosofia da linguagem **Go** foi moldada para resolver desafios práticos enfrentados no desenvolvimento de sistemas distribuídos, grandes bases de código e alta concorrência. Seus princípios fundamentais priorizam **simplicidade, eficiência e concorrência estruturada**.

---

## 🧩 **1. Simplicidade**

O design do Go busca **remover complexidades desnecessárias**. Diferente de linguagens como C++ e Java, Go **elimina características que historicamente tornaram código difícil de manter**:

- **Sem herança tradicional** → Favorece **composição sobre herança**, evitando hierarquias profundas de classes.
- **Sem exceções tradicionais (`try/catch`)** → Prefere **erros explícitos** via `error` como retorno.
- **Inferência de tipos** → Menos código boilerplate sem comprometer a segurança de tipos.
- **Estruturas sintáticas enxutas** → Go usa apenas um laço de repetição (`for`), evitando múltiplas variações complexas.

🌟 **Exemplo: Composição ao invés de Herança**
```go
package main

import "fmt"

type Engine struct {
    Power int
}

type Car struct {
    Engine // Composição ao invés de herança
    Model string
}

func main() {
    myCar := Car{Engine{Power: 150}, "GoCar"}
    fmt.Println(myCar.Model, "tem potência de", myCar.Power, "HP")
}
```

📌 **O código é mais simples e modular sem precisar de classes e herança complexa.**  

---

## ⚡ **2. Eficiência**

Go foi projetado para **compilar rapidamente, ser leve e escalável**:

- 🚀 **Compilação extremamente rápida**, reduzindo ciclos de desenvolvimento.
- 🚲 **Garbage Collection (GC) otimizado**, minimizando pausas na execução.
- ✅ **Sistema de tipos estáticos**, capturando erros em tempo de compilação.

📌 **Comparativo de tempos de compilação**
| Linguagem | Código Médio | Tempo de Compilação |
|-----------|--------------|--------------------|
| **C++** | 10.000 linhas | ⏳ 20-60s |
| **Java** | 10.000 linhas | ⏳ 10-30s |
| **Go** | 10.000 linhas | ⚡ 1-3s |

🌟 **Exemplo: Go elimina dependências externas e recompila rapidamente**
```go
package main

import "fmt"

func main() {
    fmt.Println("Go compila rápido!")
}
```

📌 **Compilar e rodar rapidamente:**  
```sh
go run main.go
```

---

## 🔄 **3. Concorrência Estruturada**

Go implementa um **modelo de concorrência robusto**, baseado no princípio:

> ❝ *"Do not communicate by sharing memory; instead, share memory by communicating."* ❞

📌 **Recursos principais de concorrência em Go:**
- 🏃 **Goroutines** → Threads leves gerenciadas pelo runtime de Go.
- 📺 **Canais (Channels)** → Mecanismo seguro para comunicação entre goroutines.
- 🛠 **`select` statement** → Multiplexação eficiente de múltiplos canais.

🌟 **Exemplo: Criando múltiplas goroutines**
```go
package main

import (
    "fmt"
    "time"
)

func say(msg string) {
    for i := 0; i < 3; i++ {
        time.Sleep(time.Millisecond * 500)
        fmt.Println(msg)
    }
}

func main() {
    go say("Hello") // Goroutine 1
    go say("Go")    // Goroutine 2
    time.Sleep(time.Second * 2) // Aguarda execuções
}
```

📌 **Esse código roda duas funções `say()` simultaneamente, sem criar threads manualmente.**

---

## 🌟 **Conclusão**

A concepção do Go foi impulsionada pela necessidade de **uma linguagem prática, produtiva e eficiente**.  
Ele combina **concorrência simplificada, compilação rápida e sintaxe enxuta**, tornando-se ideal para **infraestrutura de cloud computing e aplicações escaláveis**.

🛠️ **No próximo capítulo**, veremos a **sintaxe básica do Go**, explorando **declaração de variáveis, tipos primitivos e operadores fundamentais**. 🚀

