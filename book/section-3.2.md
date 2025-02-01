# **3.2 Laços de Repetição: `for`, `range`**

Go utiliza apenas uma estrutura de repetição: **`for`**. No entanto, sua sintaxe é flexível o suficiente para cobrir diferentes cenários, incluindo loops tradicionais, iterações sobre coleções e repetições indefinidas.

---

## **3.2.1 Estrutura Básica do `for`**

A forma mais comum do `for` em Go segue o padrão de três expressões:

```go
for inicialização; condição; incremento {
    // Código a ser repetido
}
```

Exemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println("Iteração:", i)
}
```

📌 **Diferente de C e Java, Go não suporta `while` e `do-while`, pois `for` cobre todos esses casos.**

### **Equivalente ao `while`**

Podemos usar `for` sem a inicialização e incremento, criando um loop estilo `while`:

```go
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}
```

### **Loop Infinito**

Se omitirmos todas as expressões, teremos um loop infinito:

```go
for {
    fmt.Println("Rodando para sempre...")
}
```

📌 **Loop infinito é útil para servidores e processos que nunca devem encerrar.**

---

## **3.2.2 Iterando sobre Arrays, Slices e Mapas com `range`**

Go fornece a palavra-chave `range` para percorrer **arrays, slices, strings, mapas e canais** de forma simplificada.

### **Iterando sobre um Slice**

```go
numeros := []int{10, 20, 30}

for indice, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
}
```

📌 **Se o índice não for necessário, use `_` para ignorá-lo:**

```go
for _, valor := range numeros {
    fmt.Println("Valor:", valor)
}
```

### **Iterando sobre um Mapa**

```go
alunos := map[string]int{"Alice": 20, "Bob": 25}

for nome, idade := range alunos {
    fmt.Printf("%s tem %d anos\n", nome, idade)
}
```

### **Iterando sobre uma String (`rune` por `rune`)**

Strings em Go são codificadas em **UTF-8**. Usando `range`, podemos percorrer os caracteres:

```go
s := "GoLang"

for i, r := range s {
    fmt.Printf("Índice: %d, Caractere: %c\n", i, r)
}
```

📌 **Isso é útil para manipulação correta de strings Unicode!**

---

## **3.2.3 Uso de `break` e `continue`**

### **Interrompendo o Loop com `break`**

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Sai do loop quando i == 5
    }
    fmt.Println(i)
}
```

### **Pulando uma Iteração com `continue`**

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // Pula a iteração quando i == 2
    }
    fmt.Println(i)
}
```

📌 **`break` e `continue` funcionam tanto em loops normais quanto com `range`.**

---

## **3.2.4 Rotulando Loops para Controle Avançado**

Go permite **rotular loops** para controlar `break` e `continue` em loops aninhados:

```go
externo:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break externo // Sai do loop externo
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

📌 **Isso evita a necessidade de usar flags booleanas para sair de loops aninhados.**

---

## **3.2.5 Comparação com Outras Linguagens**

| Conceito | Go | C / Java |
|----------|----|---------|
| `while` loop | ❌ Não existe | ✅ Existe |
| `for` tradicional | ✅ Sim | ✅ Sim |
| `for-each` (`range`) | ✅ Sim | ✅ Sim (`foreach` em Java) |
| `break` e `continue` | ✅ Sim | ✅ Sim |
| Rotulagem de loops | ✅ Sim | ❌ Não existe em Java |

📌 **A ausência de `while` e `do-while` simplifica a sintaxe e reduz redundância.**

---

## **Conclusão**

Os loops em Go são poderosos e flexíveis, cobrindo todos os casos com apenas `for`. O uso de `range` torna a iteração sobre coleções mais simples e eficiente. No próximo capítulo, exploraremos **`break`, `continue` e `goto`**, aprofundando o controle de fluxo! 🚀
