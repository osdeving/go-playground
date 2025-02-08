# **10.6 Exemplos Práticos de Concorrência com `select`**

Agora que entendemos como `select` funciona, vamos explorar alguns **exemplos práticos** onde ele é essencial para gerenciar concorrência em Go.

Nesta seção, veremos:

- Um **servidor concorrente** que lida com múltiplas requisições
- Um **worker pool** para distribuição de tarefas
- Um **sistema de timeout dinâmico**

---

## **10.6.1 Servidor Concorrente com `select`**

Vamos criar um **servidor TCP concorrente** que aceita conexões e responde a cada cliente de forma independente:

```go
package main

import (
    "fmt"
    "net"
    "time"
)

func handleClient(conn net.Conn) {
    defer conn.Close()

    ch := make(chan string)
    
    go func() {
        buffer := make([]byte, 1024)
        _, err := conn.Read(buffer)
        if err == nil {
            ch <- "Recebido: " + string(buffer)
        }
    }()

    select {
    case msg := <-ch:
        conn.Write([]byte(msg))
    case <-time.After(5 * time.Second):
        fmt.Println("Timeout! Nenhuma resposta do cliente.")
    }
}

func main() {
    ln, _ := net.Listen("tcp", ":8080")
    fmt.Println("Servidor ouvindo na porta 8080")

    for {
        conn, _ := ln.Accept()
        go handleClient(conn)
    }
}
```

📌 **O servidor aceita múltiplas conexões simultâneas sem bloqueios!**  
📌 **Cada conexão é tratada com um `select`, garantindo timeout adequado.**  

---

## **10.6.2 Worker Pool para Processamento Concorrente**

Podemos usar `select` para implementar um **pool de workers**, onde múltiplas Goroutines processam tarefas de uma fila:

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
    for task := range tasks {
        fmt.Printf("Worker %d processando tarefa %d
", id, task)
        time.Sleep(time.Second)
        results <- task * 2
    }
}

func main() {
    tasks := make(chan int, 5)
    results := make(chan int, 5)

    for i := 1; i <= 3; i++ {
        go worker(i, tasks, results)
    }

    for i := 1; i <= 5; i++ {
        tasks <- i
    }
    close(tasks)

    for i := 1; i <= 5; i++ {
        fmt.Println("Resultado:", <-results)
    }
}
```

📌 **Distribuímos tarefas entre 3 workers de forma eficiente.**  
📌 **O `close(tasks)` sinaliza que não há mais trabalho a ser enviado.**  

---

## **10.6.3 Timeout Dinâmico para Processamento Assíncrono**

Podemos ajustar **timeouts dinamicamente** usando `select` e `time.After()`:

```go
package main

import (
    "fmt"
    "time"
)

func processar(dados chan int) {
    select {
    case valor := <-dados:
        fmt.Println("Processado:", valor)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout! Nenhum dado recebido.")
    }
}

func main() {
    dados := make(chan int)

    go processar(dados)

    time.Sleep(3 * time.Second) // Simula atraso no envio

    dados <- 42 // Esse dado chega depois do timeout
}
```

📌 **Se os dados demorarem mais de 2 segundos, um timeout ocorre.**  
📌 **Evita que Goroutines fiquem bloqueadas indefinidamente.**  

---

## **Conclusão**

Esses exemplos demonstram como `select` pode ser usado para **escrever sistemas concorrentes robustos e escaláveis**.  
No próximo capítulo, exploraremos **Mutexes e controle avançado de concorrência**, garantindo segurança em ambientes multi-threaded! 🚀
