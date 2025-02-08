# **14.4 WebSockets e gRPC**

A comunicação em tempo real é essencial para muitas aplicações modernas, como chats, jogos online e sistemas distribuídos.  
Duas tecnologias populares para comunicação eficiente e de baixa latência são **WebSockets** e **gRPC**.  

Nesta seção, exploraremos:

- O que são WebSockets e como usá-los no Go
- Criando um servidor WebSocket em Go
- Comunicação cliente-servidor com WebSockets
- Introdução ao **gRPC** para comunicação binária otimizada
- Criando um **servidor e cliente gRPC**
- Comparação entre **WebSockets e gRPC**

---

## **14.4.1 Introdução aos WebSockets**

WebSockets são uma tecnologia que permite **conexões bidirecionais persistentes** entre cliente e servidor,  
permitindo a **troca contínua de mensagens** sem necessidade de múltiplas requisições HTTP.

📌 **Vantagens dos WebSockets:**  
✔ **Baixa latência** - Perfeito para aplicações em tempo real.  
✔ **Conexão persistente** - Reduz sobrecarga de conexões repetidas.  
✔ **Comunicação bidirecional** - Cliente e servidor podem enviar mensagens a qualquer momento.  

✅ **Exemplo: Criando um Servidor WebSocket em Go**

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Erro ao atualizar conexão:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Erro ao ler mensagem:", err)
            break
        }
        fmt.Println("Mensagem recebida:", string(msg))
        conn.WriteMessage(messageType, []byte("Recebido: "+string(msg)))
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    fmt.Println("Servidor WebSocket rodando em ws://localhost:8080/ws")
    http.ListenAndServe(":8080", nil)
}
```

📌 **O servidor escuta conexões WebSocket na rota `/ws` e responde às mensagens recebidas.**  

✅ **Testando com JavaScript no navegador:**  

Abra o console (`F12` > Console) e execute:  
```js
let ws = new WebSocket("ws://localhost:8080/ws");
ws.onmessage = (event) => console.log(event.data);
ws.send("Olá, WebSocket!");
```

📌 **O servidor responderá "Recebido: Olá, WebSocket!"**  

---

## **14.4.2 Criando um Cliente WebSocket em Go**

O Go permite criar **clientes WebSocket** para interagir com servidores.  

✅ **Exemplo: Cliente WebSocket em Go**

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gorilla/websocket"
)

func main() {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }
    defer conn.Close()

    for i := 0; i < 5; i++ {
        msg := fmt.Sprintf("Mensagem %d", i+1)
        conn.WriteMessage(websocket.TextMessage, []byte(msg))
        
        _, response, _ := conn.ReadMessage()
        fmt.Println("Resposta do servidor:", string(response))
        time.Sleep(time.Second)
    }
}
```

📌 **O cliente envia mensagens ao servidor e recebe respostas.**  

---

## **14.4.3 Introdução ao gRPC**

O **gRPC** (Google Remote Procedure Call) é um framework de comunicação que utiliza **HTTP/2** e **Protocol Buffers**  
para enviar dados binários compactados de maneira eficiente.

📌 **Por que usar gRPC?**  
✔ **Desempenho superior ao REST (dados binários vs JSON)**  
✔ **Suporte a diversas linguagens (Go, Python, Java, etc.)**  
✔ **Streaming bidirecional nativo**  
✔ **Segurança via TLS integrada**  

✅ **Exemplo: Definição de serviço gRPC (`.proto`)**

```proto
syntax = "proto3";

package chat;

service ChatService {
    rpc SendMessage (Message) returns (Response);
}

message Message {
    string sender = 1;
    string text = 2;
}

message Response {
    string reply = 1;
}
```

📌 **Aqui definimos um serviço `ChatService` com um método `SendMessage()`.**  

---

## **14.4.4 Criando um Servidor gRPC em Go**

Para usar **gRPC**, instale o pacote:  
```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

✅ **Exemplo: Servidor gRPC em Go**

```go
package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"
    pb "chat/proto"
)

type server struct {
    pb.UnimplementedChatServiceServer
}

func (s *server) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Response, error) {
    reply := fmt.Sprintf("Mensagem recebida: %s", msg.Text)
    return &pb.Response{Reply: reply}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()
    pb.RegisterChatServiceServer(s, &server{})

    fmt.Println("Servidor gRPC rodando na porta 50051...")
    s.Serve(lis)
}
```

📌 **O servidor processa mensagens e responde ao cliente.**  

---

## **14.4.5 Criando um Cliente gRPC em Go**

✅ **Exemplo: Cliente gRPC em Go**

```go
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    pb "chat/proto"
)

func main() {
    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    defer conn.Close()

    client := pb.NewChatServiceClient(conn)
    response, _ := client.SendMessage(context.Background(), &pb.Message{Sender: "Alice", Text: "Olá, gRPC!"})

    fmt.Println("Resposta do servidor:", response.Reply)
}
```

📌 **O cliente se conecta ao servidor gRPC e envia mensagens.**  

✅ **Executando o teste:**  
1. **Inicie o servidor gRPC** (`go run server.go`)  
2. **Execute o cliente** (`go run client.go`)  
3. **Veja a resposta processada pelo servidor**  

---

## **14.4.6 Comparação entre WebSockets e gRPC**

| Característica | WebSockets | gRPC |
|---------------|-----------|------|
| **Formato** | Texto (JSON) | Binário (Protobuf) |
| **Protocolo** | HTTP/1.1 | HTTP/2 |
| **Velocidade** | Boa | Excelente |
| **Comunicação** | Bidirecional | Bidirecional |
| **Casos de Uso** | Chats, jogos, eventos em tempo real | Comunicação entre microsserviços |

📌 **Use WebSockets para comunicação em tempo real entre navegadores.**  
📌 **Use gRPC para chamadas eficientes entre serviços backend.**  

---

## **Conclusão**

WebSockets e gRPC oferecem **soluções poderosas para comunicação de baixa latência**.  
No próximo capítulo, exploraremos **como criar APIs RESTful robustas em Go!** 🚀
