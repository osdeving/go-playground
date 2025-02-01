# **10.3 Comunicação entre Goroutines com Channels (`chan`)**

A programação concorrente em Go foi projetada com o princípio **"Não se comunique compartilhando memória; compartilhe memória comunicando-se"**.  
Isso significa que, em vez de sincronizar o acesso a variáveis compartilhadas (usando `Mutex` ou `atomic`), o Go favorece **Channels (`chan`)** como mecanismo primário para comunicação entre Goroutines.

Nesta seção, exploraremos:

- O que são Channels e como funcionam
- Criando e utilizando Channels
- Comunicação síncrona e concorrente entre Goroutines
- Diferenças entre Channels e outras formas de sincronização
- Erros comuns e melhores práticas ao usar Channels

---

## **10.3.1 O Que São Channels?**

Um **Channel (`chan`)** é um meio seguro de **passar dados entre Goroutines**.  
Ele funciona como uma **fila de mensagens**: uma Goroutine pode enviar dados para um Channel e outra pode receber.

✅ **Criando um Channel:**

```go
ch := make(chan int) // Canal de inteiros
```

✅ **Enviando e recebendo dados:**

```go
go func() {
    ch <- 42 // Envia o valor 42 pelo canal
}()

x := <-ch // Recebe o valor do canal
fmt.Println(x) // 42
```

📌 **`ch <- valor` envia um valor ao canal.**  
📌 **`<- ch` recebe um valor do canal.**  

🔎 **Visualização do fluxo de comunicação:**

```
Goroutine 1  ---->  [Channel]  ---->  Goroutine 2
```

---

## **10.3.2 Comunicação Bloqueante e Concorrente**

Os Channels **bloqueiam** automaticamente até que haja alguém para receber os dados:

```go
func main() {
    ch := make(chan string)

    go func() {
        ch <- "Mensagem" // Aguarda até que alguém receba
    }()

    fmt.Println(<-ch) // "Mensagem" (desbloqueia o envio)
}
```

✅ **Isso permite sincronizar Goroutines de forma natural, sem precisar de `Mutex`!**

📌 **Se ninguém estiver recebendo, o envio `ch <- valor` bloqueia a execução.**  
📌 **Se ninguém estiver enviando, a recepção `<-ch` também bloqueia.**  

---

## **10.3.3 Comunicação Entre Múltiplas Goroutines**

Channels são ideais para coordenar múltiplas Goroutines:

```go
func trabalhador(id int, ch chan string) {
    ch <- fmt.Sprintf("Trabalhador %d terminou!", id)
}

func main() {
    ch := make(chan string)

    for i := 1; i <= 3; i++ {
        go trabalhador(i, ch)
    }

    for i := 1; i <= 3; i++ {
        fmt.Println(<-ch) // Aguarda cada trabalhador finalizar
    }
}
```

📌 **Cada `trabalhador` envia um resultado para o Channel, e `main` os coleta sequencialmente.**

✅ **Isso evita a necessidade de `sync.WaitGroup` para esperar Goroutines!**

---

## **10.3.4 Comparação Entre Channels e Outras Técnicas de Sincronização**

| Técnica | Uso Principal | Bloqueante? | Seguro para Concorrência? |
|---------|--------------|------------|-----------------|
| `sync.Mutex` | Proteção de dados compartilhados | ❌ Não | ✅ Sim |
| `sync.WaitGroup` | Aguardar Goroutines finalizarem | ✅ Sim | ✅ Sim |
| **`chan` (Channel)** | Comunicação entre Goroutines | ✅ Sim | ✅ Sim |

📌 **Channels são mais seguros e intuitivos do que `Mutex`, pois evitam acesso direto a memória compartilhada.**

---

## **10.3.5 Erros Comuns ao Usar Channels**

❌ **Esquecer de fechar um Channel (`close()`)**  

```go
ch := make(chan int)

go func() {
    ch <- 10
}()

fmt.Println(<-ch)
fmt.Println(<-ch) // Deadlock! Ninguém mais enviando
```

✅ **Fechar o Channel quando não for mais necessário:**

```go
close(ch)
```

---

❌ **Enviar para um Channel fechado**

```go
ch := make(chan int)
close(ch)
ch <- 10 // Pânico! Canal já fechado
```

✅ **Verifique antes de enviar:**

```go
if _, aberto := <-ch; !aberto {
    fmt.Println("Canal fechado")
}
```

---

## **10.3.6 Boas Práticas**

✔ **Use Channels para comunicação entre Goroutines sempre que possível.**  
✔ **Feche um Channel (`close()`) quando não precisar mais enviar dados.**  
✔ **Evite Channels globais; prefira passá-los como argumentos.**  
✔ **Evite `deadlocks` garantindo que sempre há consumidores ativos.**  

---

## **Conclusão**

Os **Channels (`chan`)** são uma das maiores vantagens do Go para escrever código concorrente seguro e eficiente.  
No próximo capítulo, exploraremos **Channels Buffered e Unbuffered**, aprofundando no controle de fluxo entre Goroutines! 🚀
