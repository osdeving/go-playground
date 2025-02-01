# **10.4 Channels Buffered e Unbuffered**

Os **Channels** são um dos mecanismos mais poderosos do Go para comunicação concorrente.  
No capítulo anterior, vimos **Channels Unbuffered**, que bloqueiam a execução até que haja um receptor disponível.  
Agora, exploraremos **Channels Buffered**, que permitem armazenar múltiplos valores antes de serem recebidos.

Nesta seção, abordaremos:

- Diferença entre **Channels Unbuffered e Buffered**
- Criando e utilizando Channels Buffered
- Controle de fluxo e sincronização eficiente
- Como evitar bloqueios indesejados
- Comparação com filas tradicionais de mensagens

---

## **10.4.1 Diferença Entre Channels Buffered e Unbuffered**

| Tipo de Channel | Bloqueia no Envio? | Bloqueia na Leitura? | Capacidade |
|----------------|----------------|----------------|------------|
| **Unbuffered** | ✅ Sim (até que alguém leia) | ✅ Sim (até que alguém envie) | 0 |
| **Buffered** | ❌ Não (até encher) | ✅ Sim (até que haja dados) | `N` valores |

✅ **Exemplo de Channel Unbuffered (bloqueia até receber):**

```go
ch := make(chan int) // Sem buffer

go func() {
    ch <- 10 // Bloqueia até alguém ler
}()

fmt.Println(<-ch) // 10
```

✅ **Exemplo de Channel Buffered (não bloqueia até encher):**

```go
ch := make(chan int, 3) // Buffer de tamanho 3

ch <- 1 // OK
ch <- 2 // OK
ch <- 3 // OK
// ch <- 4 // Bloqueia! Buffer cheio

fmt.Println(<-ch) // 1
```

📌 **O Channel Buffered permite armazenar valores até atingir sua capacidade.**

✅ **Isso permite maior eficiência, reduzindo bloqueios desnecessários.**

---

## **10.4.2 Como Channels Buffered Melhoram a Performance?**

Os Channels Buffered ajudam a **desacoplar o envio e recebimento**:

1️⃣ **Sem buffer:**  
   - Cada envio precisa de um receptor pronto (sincronização rígida).  
   - Útil quando a ordem de execução importa.  

2️⃣ **Com buffer:**  
   - O produtor pode enviar vários valores sem esperar.  
   - O consumidor pode processar os valores em paralelo.  
   - Útil para pipelines de dados.  

✅ **Exemplo com produtores e consumidores:**

```go
ch := make(chan string, 2)

go func() {
    ch <- "Processando 1"
    ch <- "Processando 2"
    fmt.Println("Dados enviados")
}()

time.Sleep(time.Second) // Simulando atraso no consumidor

fmt.Println(<-ch) // "Processando 1"
fmt.Println(<-ch) // "Processando 2"
```

📌 **O produtor não ficou bloqueado, pois havia espaço no buffer.**  
📌 **O consumidor processou os dados quando ficou disponível.**  

---

## **10.4.3 Evitando Deadlocks e Bloqueios**

Se um Channel Buffered estiver **cheio**, o envio bloqueia até que haja espaço disponível:

```go
ch := make(chan int, 2)

ch <- 1
ch <- 2
// ch <- 3 // Bloqueia aqui! Nenhum consumidor disponível
```

✅ **Para evitar deadlocks:**

1. **Leia os valores antes do buffer encher.**  
2. **Feche o canal quando terminar (`close()`).**  
3. **Use `select` para evitar bloqueios.**  

✅ **Evitando bloqueios com `select`**:

```go
select {
case ch <- 10:
    fmt.Println("Valor enviado")
default:
    fmt.Println("Canal cheio, evitando bloqueio!")
}
```

📌 **Se `ch` estiver cheio, a execução continua sem bloquear.**  

---

## **10.4.4 Como Saber Se um Canal Está Fechado?**

Podemos verificar se um canal foi fechado ao tentar receber um valor:

```go
ch := make(chan int, 2)
close(ch)

valor, aberto := <-ch
fmt.Println(valor, aberto) // 0, false (canal fechado)
```

📌 **Se um canal estiver fechado, a leitura retorna o valor padrão do tipo (`0` para `int`, `""` para `string`).**  

✅ **Nunca envie para um canal fechado:**

```go
ch := make(chan int)
close(ch)
// ch <- 10 // Pânico! Canal fechado
```

📌 **O envio para um canal fechado gera um `panic` e encerra o programa.**

---

## **10.4.5 Comparação: Channels vs. Outras Estruturas de Comunicação**

| Técnica | Uso Principal | Bloqueante? | Controle de Fluxo |
|---------|--------------|------------|-----------------|
| **Channel Unbuffered** | Comunicação sincronizada | ✅ Sim | ❌ Não |
| **Channel Buffered** | Comunicação assíncrona | ❌ Não (até encher) | ✅ Sim |
| **Fila (Queue)** | Processamento assíncrono | ❌ Não | ✅ Sim |
| **Mutex (`sync.Mutex`)** | Controle de acesso | ❌ Não | ❌ Não |

📌 **Channels Buffered funcionam como uma **fila de mensagens**, garantindo fluxo controlado entre Goroutines.**

✅ **Se precisar processar mensagens em lote, um Buffer é mais eficiente.**

---

## **10.4.6 Boas Práticas**

✔ **Use Channels Unbuffered para sincronização estrita.**  
✔ **Use Channels Buffered para desacoplar produtores e consumidores.**  
✔ **Sempre feche o Channel (`close()`) quando terminar o envio.**  
✔ **Evite deadlocks garantindo que há consumidores ativos.**  
✔ **Use `select` para evitar bloqueios desnecessários.**  

---

## **Conclusão**

Os **Channels Buffered** aumentam a eficiência ao permitir a comunicação assíncrona entre Goroutines.  
No próximo capítulo, exploraremos o uso do **`select` para multiplexação de canais**, permitindo processar múltiplas comunicações concorrentes! 🚀
