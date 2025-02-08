# **10.1 Criando e Executando Goroutines**

A **concorrência** é um dos pilares centrais do Go, e **Goroutines** são a base para escrever programas concorrentes de forma eficiente.  
Diferente de **threads** tradicionais, Goroutines são extremamente leves e permitem escalabilidade massiva sem a complexidade da programação paralela convencional.

Nesta seção, exploraremos:

- O que são Goroutines e como elas funcionam
- Criando e executando Goroutines
- Como Goroutines são agendadas pelo runtime de Go
- Comparação entre Goroutines e Threads tradicionais
- Erros comuns e boas práticas ao utilizar Goroutines

---

## **10.1.1 O Que São Goroutines?**

Uma **Goroutine** é uma **função que executa de forma independente e concorrente**, gerenciada pelo runtime do Go.  
Diferente de threads tradicionais, uma Goroutine consome menos recursos e pode ser escalada em grande número sem penalidades significativas de desempenho.

✅ **Criando uma Goroutine** é simples:

```go
package main

import (
    "fmt"
    "time"
)

func mensagem() {
    fmt.Println("Executando Goroutine!")
}

func main() {
    go mensagem() // Executa a função de forma concorrente
    time.Sleep(time.Second) // Espera para permitir execução
}
```

📌 **A palavra-chave `go` inicia uma Goroutine.**  
📌 **A execução do programa principal não aguarda a Goroutine finalizar.**  

🔎 **Sem o `time.Sleep()`, o programa pode encerrar antes da Goroutine executar!**

---

## **10.1.2 Agendamento de Goroutines**

Goroutines são gerenciadas pelo **scheduler do Go**, que decide quais Goroutines devem rodar em quais threads do sistema operacional.

📌 **Diferente de threads, Goroutines são multiplexadas em um pool de threads do SO.**  
📌 **Isso significa que podemos criar milhares de Goroutines sem criar milhares de threads.**

✅ **Exemplo de múltiplas Goroutines:**

```go
func imprimirMensagem(mensagem string) {
    for i := 0; i < 5; i++ {
        fmt.Println(mensagem, i)
    }
}

func main() {
    go imprimirMensagem("Goroutine 1")
    go imprimirMensagem("Goroutine 2")

    time.Sleep(time.Second) // Espera execução das Goroutines
}
```

🔎 **Como o scheduler pode alternar Goroutines, a ordem de execução pode variar.**

---

## **10.1.3 Goroutines vs. Threads**

| Característica | Goroutines (Go) | Threads (Java, C++) |
|---------------|----------------|---------------------|
| Criação Leve  | ✅ Sim | ❌ Custo alto |
| Agendamento   | ✅ Cooperativo | ❌ Preemptivo |
| Comunicação   | ✅ Channels | ❌ Mutexes e Locks |
| Stack Inicial | ✅ Pequena (~2KB) | ❌ Grande (1MB ou mais) |
| Quantidade Suportada | ✅ Milhares/Milhões | ❌ Limitado pelo SO |

📌 **Go utiliza um runtime próprio para gerenciar Goroutines, evitando overhead do SO.**

✅ **O runtime do Go pode pausar e alternar Goroutines conforme necessário, otimizando o uso do processador.**

---

## **10.1.4 Controle e Sincronização**

Como Goroutines executam de forma concorrente, precisamos de **mecanismos de sincronização** para evitar problemas como **condições de corrida**.

Exemplo de **condição de corrida**:

```go
contador := 0

for i := 0; i < 1000; i++ {
    go func() { contador++ }() // Acesso concorrente à variável
}

fmt.Println("Contador:", contador) // Resultado imprevisível!
```

📌 **Múltiplas Goroutines acessam `contador` ao mesmo tempo, causando comportamento indeterminado.**

✅ **No Capítulo 11, exploraremos `sync.Mutex` e `sync.WaitGroup` para evitar esses problemas.**

---

## **10.1.5 Melhorando a Escalabilidade**

Em Go, podemos aumentar a eficiência ajustando o número de threads disponíveis para o runtime:

```go
import "runtime"

runtime.GOMAXPROCS(4) // Define 4 threads para execução das Goroutines
```

📌 **Isso pode melhorar a performance em sistemas multicore, mas nem sempre é necessário.**

✅ **O runtime do Go gerencia isso automaticamente na maioria dos casos.**

---

## **10.1.6 Boas Práticas**

✔ **Sempre gerencie a finalização das Goroutines (`sync.WaitGroup`, `channels`).**  
✔ **Evite concorrência desnecessária para reduzir complexidade.**  
✔ **Prefira `channels` para comunicação entre Goroutines em vez de locks (`mutex`).**  
✔ **Use `runtime.NumGoroutine()` para monitorar Goroutines ativas.**  

✅ **Exemplo de monitoramento:**

```go
fmt.Println("Goroutines ativas:", runtime.NumGoroutine())
```

---

## **Conclusão**

As **Goroutines** são uma das maiores vantagens do Go para escrever código concorrente de forma eficiente.  
No próximo capítulo, exploraremos **`sync.WaitGroup`**, uma ferramenta essencial para aguardar a finalização de múltiplas Goroutines! 🚀
