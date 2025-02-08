# **3.4 Defer, Panic e Recover**

Go fornece três mecanismos especiais para controle de fluxo em situações específicas: **`defer`**, **`panic`** e **`recover`**. Eles são essenciais para garantir a **finalização de recursos**, **manipulação de erros inesperados** e **recuperação de falhas** sem comprometer a execução do programa.

---

## **3.4.1 `defer`: Execução Adiada**

O comando `defer` **atrasará** a execução de uma função até que a função que a contém retorne. Isso é útil para **fechar arquivos, liberar conexões ou limpar memória**, garantindo que essas operações ocorram independentemente de erros.

### **Sintaxe Básica**

```go
func main() {
    defer fmt.Println("Isso será impresso por último")
    fmt.Println("Executando...")
}
```

Saída:

```
Executando...
Isso será impresso por último
```

📌 **Go empilha os `defer`, executando-os em ordem LIFO (Last In, First Out)**:

```go
func main() {
    defer fmt.Println("1º defer")
    defer fmt.Println("2º defer")
    defer fmt.Println("3º defer")
    fmt.Println("Finalizando função")
}
```

Saída:

```
Finalizando função
3º defer
2º defer
1º defer
```

### **Uso Comum: Fechamento de Arquivos**

```go
func main() {
    arquivo, err := os.Open("dados.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer arquivo.Close() // Garante o fechamento do arquivo
}
```

📌 **Mesmo que ocorra um erro, `defer` será executado antes do retorno da função.**

---

## **3.4.2 `panic`: Interrompendo a Execução**

`panic` é usado para gerar um erro fatal e interromper a execução do programa.

### **Criando um `panic`**

```go
func main() {
    fmt.Println("Antes do panic")
    panic("Erro crítico!") // Interrompe a execução
    fmt.Println("Isso nunca será executado")
}
```

Saída:

```
Antes do panic
panic: Erro crítico!
```

📌 **Um `panic` causa a finalização do programa, mas executa os `defer` antes de encerrar.**

### **`panic` com `defer`**

```go
func main() {
    defer fmt.Println("Isso será executado antes do fechamento")
    panic("Erro inesperado!")
}
```

Saída:

```
Isso será executado antes do fechamento
panic: Erro inesperado!
```

📌 **Isso garante que recursos sejam liberados antes da falha.**

---

## **3.4.3 `recover`: Capturando um `panic`**

O `recover` permite capturar um `panic` e evitar que o programa seja encerrado abruptamente.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do erro:", r)
        }
    }()

    fmt.Println("Iniciando")
    panic("Falha grave!") // Disparando um panic
    fmt.Println("Isso nunca será executado")
}
```

Saída:

```
Iniciando
Recuperado do erro: Falha grave!
```

📌 **Se `recover()` for chamado dentro de `defer`, ele captura o erro e impede o fechamento do programa.**

### **Manipulando `panic` e retornando à execução normal**

```go
func podeFalhar() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Erro tratado:", r)
        }
    }()

    panic("Erro crítico!")
    fmt.Println("Isso não será impresso")
}

func main() {
    fmt.Println("Executando...")
    podeFalhar()
    fmt.Println("Execução continua após recover")
}
```

Saída:

```
Executando...
Erro tratado: Erro crítico!
Execução continua após recover
```

📌 **Isso é útil para capturar erros, logá-los e continuar a execução do programa.**

---

## **3.4.4 Comparação entre `defer`, `panic` e `recover`**

| Comando  | Função |
|----------|--------|
| `defer`  | Atrasar execução até o final da função |
| `panic`  | Interromper execução imediatamente |
| `recover` | Capturar um `panic` e evitar o encerramento do programa |

📌 **Geralmente, `panic` e `recover` são usados para erros críticos, enquanto `defer` é mais comum para limpeza de recursos.**

---

## **3.4.5 Casos Especiais e Boas Práticas**

1. **Evite usar `panic` para erros comuns** 🚫  
   - Prefira retornar erros em vez de interromper o programa.

```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

2. **Use `defer` para fechar conexões** ✅  
   - Isso evita vazamento de memória e recursos abertos.

```go
func salvarDados() {
    conn := conectarBanco()
    defer conn.Fechar() // Garante que o banco seja fechado
}
```

3. **Use `recover` apenas onde necessário** 🚨  
   - Capturar `panic` indiscriminadamente pode esconder erros sérios.

---

## **Conclusão**

Os comandos `defer`, `panic` e `recover` fornecem um mecanismo robusto para **controle de fluxo e manipulação de erros**. `defer` é amplamente utilizado para **finalização de recursos**, enquanto `panic` e `recover` são úteis para **tratar falhas críticas**.

No próximo capítulo, exploraremos **estruturas de dados e manipulação de memória**, aprofundando a modelagem de dados em Go! 🚀
