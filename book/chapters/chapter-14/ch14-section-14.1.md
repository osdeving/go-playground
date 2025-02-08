# **14.1 Comunicação via TCP e UDP (`net`)**

A comunicação em rede é um aspecto fundamental no desenvolvimento de sistemas distribuídos e aplicações web.  
O Go oferece suporte nativo para **TCP** (Transmission Control Protocol) e **UDP** (User Datagram Protocol)  
através do pacote `net`, fornecendo uma interface poderosa para construir servidores e clientes de rede.

Nesta seção, exploraremos:

- Como o TCP e UDP funcionam em Go
- Criando servidores e clientes TCP
- Enviando e recebendo dados via UDP
- Comparação entre TCP e UDP
- Melhores práticas para segurança e desempenho

---

## **14.1.1 Introdução ao TCP e UDP**

📌 **TCP (Transmission Control Protocol)**  
- Conexão orientada (handshake de três vias)  
- Garante entrega ordenada dos pacotes  
- Ideal para HTTP, FTP, bancos de dados e streaming  

📌 **UDP (User Datagram Protocol)**  
- Sem conexão, rápido e leve  
- Não garante entrega ou ordem dos pacotes  
- Utilizado em DNS, VoIP, jogos online  

✅ **Escolha TCP para comunicação confiável** e **UDP para comunicação rápida e leve**.  

---

## **14.1.2 Criando um Servidor TCP em Go**

O protocolo **TCP** garante **comunicação confiável e ordenada** entre cliente e servidor.

✅ **Exemplo: Servidor TCP simples**

```go
package main

import (
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Conexão encerrada:", err)
            return
        }
        fmt.Println("Recebido:", string(buffer[:n]))
        conn.Write([]byte("Mensagem recebida!
")) // Responde ao cliente
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
        go handleConnection(conn) // Trata conexões concorrentes
    }
}
```

📌 **O servidor escuta na porta `8080` e aceita múltiplas conexões via Goroutines.**  
📌 **Cada cliente recebe uma resposta do servidor.**  

✅ **Teste o servidor TCP com Telnet:**  
```sh
telnet localhost 8080
```

---

## **14.1.3 Criando um Cliente TCP em Go**

✅ **Exemplo: Cliente TCP que se conecta ao servidor e envia mensagens**

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Erro ao conectar:", err)
        return
    }
    defer conn.Close()

    message := "Olá, servidor!
"
    conn.Write([]byte(message))

    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)
    fmt.Println("Resposta do servidor:", string(buffer[:n]))
}
```

📌 **O cliente se conecta ao servidor na porta `8080`, envia uma mensagem e recebe uma resposta.**  

✅ **Executando o teste:**  
1. Rode o servidor primeiro (`go run server.go`)  
2. Depois, execute o cliente (`go run client.go`)  
3. Veja a troca de mensagens entre cliente e servidor  

---

## **14.1.4 Criando um Servidor UDP em Go**

O **UDP** é ideal para transmissões rápidas, mas sem garantia de entrega.  

✅ **Exemplo: Servidor UDP que recebe mensagens**

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("Erro ao resolver endereço:", err)
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Erro ao iniciar servidor UDP:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Servidor UDP escutando na porta 8080...")

    buffer := make([]byte, 1024)
    for {
        n, clientAddr, _ := conn.ReadFromUDP(buffer)
        fmt.Println("Recebido de", clientAddr, ":", string(buffer[:n]))
        conn.WriteToUDP([]byte("Mensagem recebida!
"), clientAddr)
    }
}
```

📌 **O servidor UDP recebe pacotes e responde ao remetente.**  

✅ **Testando com Netcat:**  
```sh
echo "Olá UDP" | nc -u -w1 localhost 8080
```

---

## **14.1.5 Criando um Cliente UDP em Go**

✅ **Exemplo: Cliente UDP que envia mensagens**

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Erro ao resolver endereço:", err)
        return
    }

    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("Erro ao conectar UDP:", err)
        return
    }
    defer conn.Close()

    message := "Olá, servidor UDP!"
    conn.Write([]byte(message))

    buffer := make([]byte, 1024)
    n, _, _ := conn.ReadFromUDP(buffer)
    fmt.Println("Resposta do servidor:", string(buffer[:n]))
}
```

📌 **O cliente UDP envia um pacote e recebe uma resposta do servidor.**  

---

## **14.1.6 Comparação entre TCP e UDP**

| Característica | TCP | UDP |
|---------------|-----|-----|
| Confiabilidade | Alta (entrega garantida) | Baixa (sem garantias) |
| Ordem dos Pacotes | Sim | Não |
| Velocidade | Mais lento | Mais rápido |
| Uso Típico | HTTP, FTP, SSH | Jogos online, VoIP, DNS |

✅ **Escolha TCP para aplicações que exigem confiabilidade.**  
✅ **Escolha UDP para transmissões em tempo real e baixa latência.**  

---

## **Conclusão**

O **Go fornece suporte robusto para comunicação via TCP e UDP**, permitindo construir servidores e clientes de alto desempenho.  
No próximo capítulo, exploraremos **como criar um servidor e cliente TCP completos para aplicações reais!** 🚀
