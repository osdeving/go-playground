# **3.3 Uso de `break`, `continue`, `goto`**

Além das estruturas de repetição tradicionais, Go fornece comandos para **controlar o fluxo de execução dentro de loops** e até mesmo saltar diretamente para trechos específicos do código.

---

## **3.3.1 `break`: Interrompendo um Loop**

O comando `break` encerra a execução do loop atual e continua com a próxima instrução após ele.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Sai do loop quando i == 5
    }
    fmt.Println(i)
}
```

Saída:

```
0
1
2
3
4
```

📌 **O `break` pode ser usado em loops `for` tradicionais e em loops com `range`.**

### **Uso em Loops Aninhados**

Se `break` for usado dentro de loops aninhados, ele só interrompe o loop **mais interno**:

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break // Apenas o loop interno é interrompido
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

Saída:

```
i=0, j=0
i=1, j=0
i=2, j=0
```

---

## **3.3.2 `continue`: Pulando uma Iteração**

O `continue` interrompe a iteração **atual** do loop e avança para a próxima.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // Pula a iteração onde i == 2
    }
    fmt.Println(i)
}
```

Saída:

```
0
1
3
4
```

📌 **O `continue` é útil para ignorar certos valores sem interromper o loop completamente.**

### **Uso em Loops `range`**

```go
nums := []int{1, 2, 3, 4, 5}

for _, num := range nums {
    if num%2 == 0 {
        continue // Pula números pares
    }
    fmt.Println(num)
}
```

Saída:

```
1
3
5
```

---

## **3.3.3 `goto`: Saltos no Código**

Go permite o uso de `goto` para pular para um **rótulo específico** dentro da mesma função.

```go
fmt.Println("Início")

goto PULO

fmt.Println("Isso nunca será executado!")

PULO:
fmt.Println("Depois do goto!")
```

Saída:

```
Início
Depois do goto!
```

📌 **O `goto` só pode saltar para rótulos dentro da mesma função.**

### **`goto` vs. `break` e `continue`**

Embora `goto` possa ser usado para sair de loops, **seu uso excessivo é desencorajado** pois pode tornar o código difícil de entender.

```go
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 2 {
            goto FIM
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}

FIM:
fmt.Println("Loop encerrado!")
```

Saída:

```
i=0, j=0
i=0, j=1
Loop encerrado!
```

📌 **Evite `goto` sempre que possível! Prefira `break` e `continue` para controle de fluxo.**

---

## **3.3.4 Rotulando Loops para `break` e `continue`**

Go permite rotular loops para usar `break` e `continue` de forma explícita, útil em loops aninhados.

```go
externo:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break externo // Sai do loop externo
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

Saída:

```
i=0, j=0
```

📌 **Rotular loops evita `flags booleanas` e torna o código mais legível.**

---

## **3.3.5 Comparação com Outras Linguagens**

| Conceito | Go | C / Java |
|----------|----|---------|
| `break` | ✅ Sim | ✅ Sim |
| `continue` | ✅ Sim | ✅ Sim |
| `goto` | ✅ Sim | ⚠️ Desencorajado em Java |

📌 **Go evita a complexidade do `goto` ao fornecer loops estruturados com `break` e `continue`.**

---

## **Conclusão**

Os comandos `break`, `continue` e `goto` permitem **controle fino sobre a execução dos loops**. Embora `goto` seja suportado, **seu uso deve ser evitado** para manter a clareza do código. No próximo capítulo, exploraremos **`defer`, `panic` e `recover`**, recursos fundamentais para lidar com erros e finalização de processos em Go! 🚀
