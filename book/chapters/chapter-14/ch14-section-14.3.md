# **14.3 HTTP com `net/http`**

O protocolo **HTTP (HyperText Transfer Protocol)** é a base da comunicação na web, permitindo a transferência de dados entre clientes e servidores.  
No Go, a biblioteca padrão `net/http` fornece uma API robusta e eficiente para criar servidores e clientes HTTP sem a necessidade de bibliotecas externas.

Nesta seção, exploraremos:

- Criando um **servidor HTTP básico** em Go
- Manipulação de **rotas, query parameters e request body**
- Criando um **cliente HTTP para consumir APIs**
- Middleware, Headers e Manipulação de Cookies
- Boas práticas para **performance e segurança**

---

## **14.3.1 Criando um Servidor HTTP em Go**

A biblioteca `net/http` facilita a criação de servidores HTTP em Go, permitindo definir rotas e lidar com requisições.

✅ **Exemplo: Criando um servidor HTTP básico**

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Olá! Você acessou: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

📌 **`http.HandleFunc()` registra um handler para a rota `/`**  
📌 **`http.ListenAndServe()` inicia o servidor na porta `8080`**  

✅ **Testando o servidor:**  

Abra um navegador e acesse:  
```
http://localhost:8080
```

O servidor responderá com **"Olá! Você acessou: /"**.

---

## **14.3.2 Rotas e Query Parameters**

O Go permite extrair **query parameters** das requisições HTTP para manipular dados dinamicamente.

✅ **Exemplo: Extraindo parâmetros da URL**

```go
package main

import (
    "fmt"
    "net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Visitante"
    }
    fmt.Fprintf(w, "Olá, %s!", name)
}

func main() {
    http.HandleFunc("/hello", queryHandler)
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

📌 **Acesse `http://localhost:8080/hello?name=Alice` para ver a resposta personalizada.**  

✅ **Saída esperada:**  
```
Olá, Alice!
```

---

## **14.3.3 Lendo JSON no Request Body**

APIs modernas frequentemente recebem dados em **JSON** via **POST**.  
O Go permite **desserializar JSON** facilmente para structs.

✅ **Exemplo: Manipulando JSON no request body**

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var user User
    body, _ := io.ReadAll(r.Body)
    json.Unmarshal(body, &user)

    fmt.Fprintf(w, "Usuário recebido: %s (%s)", user.Name, user.Email)
}

func main() {
    http.HandleFunc("/user", jsonHandler)
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

✅ **Teste com `curl` enviando JSON:**  
```sh
curl -X POST http://localhost:8080/user -d '{"name": "Alice", "email": "alice@example.com"}' -H "Content-Type: application/json"
```

📌 **O servidor processa o JSON e retorna uma resposta formatada.**  

---

## **14.3.4 Criando um Cliente HTTP em Go**

O Go permite consumir APIs HTTP com o pacote `net/http`.

✅ **Exemplo: Fazendo uma requisição HTTP GET**

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    resp, err := http.Get("https://api.github.com")
    if err != nil {
        fmt.Println("Erro na requisição:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
}
```

📌 **`http.Get()` faz uma requisição GET e retorna a resposta.**  
📌 **`io.ReadAll(resp.Body)` lê a resposta do servidor.**  

✅ **Fazendo uma requisição POST**

```go
http.Post("https://example.com/api", "application/json", bytes.NewBuffer([]byte(`{"key":"value"}`)))
```

📌 **Use `http.Post()` para enviar dados ao servidor.**  

---

## **14.3.5 Middleware, Headers e Cookies**

O Go permite manipular **headers HTTP** e implementar **middlewares** para autenticação e logging.

✅ **Exemplo: Middleware de Logging**

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Requisição recebida:", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Bem-vindo ao servidor!"))
    })

    server := http.Server{
        Addr:    ":8080",
        Handler: loggingMiddleware(mux),
    }
    server.ListenAndServe()
}
```

📌 **O middleware intercepta todas as requisições e registra logs.**  

✅ **Manipulando Cookies**

```go
http.SetCookie(w, &http.Cookie{Name: "session", Value: "1234", HttpOnly: true})
```

📌 **Use `http.SetCookie()` para armazenar informações no cliente.**  

---

## **14.3.6 Boas Práticas para Performance e Segurança**

✔ **Evite carregar arquivos estáticos diretamente no código, use `http.FileServer`.**  
✔ **Sempre feche `r.Body.Close()` ao processar requisições.**  
✔ **Use `http.TimeoutHandler` para evitar requisições que travam o servidor.**  
✔ **Ative `TLS` com `http.ListenAndServeTLS()` para segurança.**  

✅ **Exemplo: Servidor HTTP seguro com TLS**

```go
http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
```

📌 **Isso ativa HTTPS usando um certificado SSL.**  

---

## **Conclusão**

O **Go simplifica a criação de servidores e clientes HTTP** com `net/http`, permitindo a construção de APIs robustas e eficientes.  
No próximo capítulo, veremos **como integrar WebSockets e GRPC para comunicação em tempo real!** 🚀
