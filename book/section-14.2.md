# **14.2 Criando um Servidor e um Cliente TCP**

A comunicação baseada no protocolo **TCP (Transmission Control Protocol)** é um dos fundamentos das redes modernas.  
O TCP oferece uma conexão confiável, garantindo a entrega dos pacotes e a ordem dos dados transmitidos.  

Nesta seção, abordaremos:

- Criando um **servidor TCP** que aceita múltiplas conexões simultâneas
- Desenvolvendo um **cliente TCP** para interagir com o servidor
- Estratégias para **manter conexões ativas e seguras**
- Tratamento de **erros e desconexões inesperadas**
- **Boas práticas para servidores TCP escaláveis**

---

## **14.2.1 Criando um Servidor TCP**

O primeiro passo para uma comunicação TCP é criar um **servidor TCP** que escuta conexões na rede.  

✅ **Exemplo: Criando um Servidor TCP em Go**

```go
package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

// Função que lida com a comunicação com cada cliente
func handleConnection(conn net.Conn) {
    defer conn.Close()

    fmt.Println("Nova conexão:", conn.RemoteAddr())

    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Conexão encerrada:", conn.RemoteAddr())
            return
        }

        fmt.Printf("Mensagem recebida: %s", message)
        conn.Write([]byte("Mensagem recebida: " + strings.ToUpper(message) + "\n"))
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Erro ao iniciar servidor:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Servidor TCP rodando na porta 8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar conexão:", err)
            continue
        }
        go handleConnection(conn) // Processa cada cliente em uma goroutine
    }
}
```

📌 **O servidor escuta na porta `8080` e aceita múltiplas conexões via Goroutines.**  
📌 **Cada mensagem recebida é transformada em maiúsculas e enviada de volta ao cliente.**  

✅ **Para testar, use Telnet:**  

```sh
telnet localhost 8080
```

Digite mensagens e veja como o servidor responde.  

---

## **14.2.2 Criando um Cliente TCP**

O **cliente TCP** precisa estabelecer uma conexão com o servidor e trocar mensagens de maneira eficiente.  

✅ **Exemplo: Criando um Cliente TCP em Go**

```go
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Erro ao conectar:", err)
        return
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Digite uma mensagem: ")
        text, _ := reader.ReadString('\n')

        conn.Write([]byte(text)) // Envia mensagem ao servidor

        response, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Resposta do servidor:", response)
    }
}
```

📌 **O cliente lê mensagens do terminal e as envia ao servidor.**  
📌 **A resposta do servidor é exibida na tela.**  

✅ **Executando o teste:**  
1. Inicie o servidor (`go run server.go`)  
2. Execute o cliente (`go run client.go`)  
3. Digite mensagens no cliente e veja a resposta do servidor  

---

## **14.2.3 Tratando Conexões de Múltiplos Clientes**

No exemplo anterior, cada cliente é processado em uma **Goroutine separada**.  
Isso permite que o servidor lide com **múltiplas conexões simultâneas** sem bloqueios.

📌 **Melhoria: Gerenciando múltiplos clientes com um mapa de conexões**

```go
var clients = make(map[net.Conn]bool)

func handleClient(conn net.Conn) {
    defer conn.Close()
    clients[conn] = true

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        message := scanner.Text()
        fmt.Println("Recebido:", message)
    }

    delete(clients, conn)
    fmt.Println("Cliente desconectado:", conn.RemoteAddr())
}
```

📌 **O mapa `clients` mantém uma lista de conexões ativas, útil para implementar broadcast.**  

---

## **14.2.4 Lidando com Erros e Desconexões**

Uma conexão TCP pode ser encerrada a qualquer momento pelo cliente ou por problemas na rede.  
É essencial tratar esses cenários corretamente.

📌 **Dicas para tratar desconexões:**  
✔ **Sempre verifique `err` após `conn.Read()`**  
✔ **Utilize `defer conn.Close()` para liberar recursos**  
✔ **Evite pânico (`panic`) em erros inesperados**  
✔ **Implemente timeout de conexão com `SetDeadline()`**  

✅ **Exemplo: Definindo um timeout para evitar clientes inativos**

```go
conn.SetDeadline(time.Now().Add(30 * time.Second))
```

📌 **Isso garante que conexões inativas sejam fechadas automaticamente após 30 segundos.**  

---

## **14.2.5 Comparação entre Diferentes Abordagens**

| Abordagem | Vantagens | Desvantagens |
|-----------|----------|-------------|
| **Servidor Single-Thread** | Simplicidade, fácil implementação | Bloqueia ao lidar com múltiplos clientes |
| **Servidor Multi-Thread (Goroutines)** | Alta escalabilidade, suporta milhares de conexões | Consumo de memória maior |
| **Servidor com Pool de Conexões** | Melhor gerenciamento de recursos | Implementação mais complexa |

📌 **Para sistemas de alta escala, recomenda-se um balanceador de carga e múltiplas instâncias do servidor.**  

---

## **Conclusão**

O **Go fornece um excelente suporte para servidores e clientes TCP**, permitindo construir aplicações robustas e escaláveis.  
No próximo capítulo, veremos **como criar aplicações HTTP usando `net/http`, o que facilita a comunicação entre sistemas distribuídos!** 🚀
