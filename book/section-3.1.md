# **3.1 Estruturas Condicionais: `if`, `else if`, `switch`**

O controle de fluxo condicional em Go permite executar diferentes blocos de código com base em condições lógicas. Nesta seção, exploraremos **`if`, `else if`, `switch`**, suas particularidades em Go e como podem ser usadas eficientemente.

---

## **3.1.1 O `if` e `else` em Go**

A estrutura `if` em Go segue um formato semelhante ao de outras linguagens, mas possui peculiaridades importantes:

```go
if condição {
    // Bloco executado se a condição for verdadeira
}
```

Exemplo:

```go
x := 10
if x > 5 {
    fmt.Println("x é maior que 5")
}
```

### **Usando `else` e `else if`**

```go
x := 10

if x < 5 {
    fmt.Println("x é menor que 5")
} else if x == 10 {
    fmt.Println("x é exatamente 10")
} else {
    fmt.Println("x é maior que 5 e diferente de 10")
}
```

📌 **Diferente de algumas linguagens, `if` e `else` em Go não exigem parênteses ao redor da condição!**

```go
// Errado:
if (x > 5) { ... } // 🚫

// Correto:
if x > 5 { ... } // ✅
```

### **Declaração de Variáveis no `if`**

Go permite **declarar variáveis dentro da condição do `if`**, tornando o código mais enxuto:

```go
if y := calcular(); y > 0 {
    fmt.Println("y é positivo:", y)
} else {
    fmt.Println("y é negativo:", y)
}
```

📌 **A variável `y` só existe dentro do escopo do `if` e `else`!**

```go
fmt.Println(y) // ERRO: "y" não existe fora do bloco if
```

---

## **3.1.2 `switch`: Alternativa ao `if-else`**

Em Go, `switch` substitui múltiplas comparações `if-else`, tornando o código mais limpo.

### **Forma básica do `switch`**

```go
dia := "segunda"

switch dia {
case "segunda":
    fmt.Println("Início da semana")
case "sexta":
    fmt.Println("Quase fim de semana!")
case "domingo":
    fmt.Println("Descanso!")
default:
    fmt.Println("Dia normal")
}
```

📌 **Diferente de C e Java, `switch` em Go **NÃO** precisa de `break` em cada `case`!**  
Go **não executa os casos seguintes automaticamente**, a menos que usemos `fallthrough`.

### **Usando `fallthrough` para continuar a execução**

Se quisermos **forçar a execução do próximo caso**, usamos `fallthrough`:

```go
x := 1

switch x {
case 1:
    fmt.Println("Caso 1")
    fallthrough
case 2:
    fmt.Println("Caso 2") // Será executado
}
```

📌 **Atenção! `fallthrough` ignora a condição do próximo `case` e o executa incondicionalmente!**

### **`switch` sem Expressão**

Em Go, um `switch` pode funcionar como um **`if-else` simplificado**, sem expressão inicial:

```go
x := 10

switch {
case x > 10:
    fmt.Println("Maior que 10")
case x == 10:
    fmt.Println("Exatamente 10")
default:
    fmt.Println("Menor que 10")
}
```

📌 **Isso é útil para checar múltiplas condições sem usar `if-else`.**

---

## **3.1.3 `switch` com Tipos (`type switch`)**

Go permite verificar o **tipo dinâmico** de uma variável usando `switch`:

```go
var i interface{} = "texto"

switch v := i.(type) {
case int:
    fmt.Println("É um inteiro:", v)
case string:
    fmt.Println("É uma string:", v)
case bool:
    fmt.Println("É um booleano:", v)
default:
    fmt.Println("Tipo desconhecido")
}
```

📌 **Isso é útil em funções genéricas que lidam com diferentes tipos!**

---

## **3.1.4 Melhorando Performance com `switch`**

Em **casos de múltiplas comparações**, `switch` pode ser **mais rápido** que `if-else`, pois algumas implementações otimizam a avaliação de `case` em tabelas de salto.

**Evite isso:**

```go
if x == 1 {
    ...
} else if x == 2 {
    ...
} else if x == 3 {
    ...
}
```

**Prefira isso:**

```go
switch x {
case 1:
    ...
case 2:
    ...
case 3:
    ...
}
```

📌 **Além de mais rápido, `switch` torna o código mais legível.**

---

## **3.1.5 Casos Especiais e Armadilhas**

1. **Comparação entre tipos diferentes falha**:

```go
var x int = 10
var y float64 = 10.0

// if x == y { ... } // ERRO: Tipos diferentes
```

Sempre converta antes:

```go
if float64(x) == y { ... } // Correto
```

2. **Valores booleanos não são convertidos implicitamente**:

```go
if 1 { ... } // ERRO!
```

Use:

```go
if 1 != 0 { ... } // Correto
```

3. **Omitir `default` pode ser um risco**:

Se não houver `default`, um `switch` pode não executar nenhum bloco, o que pode ser indesejado.

```go
switch dia {
case "segunda":
    fmt.Println("Início da semana")
}
```

Sempre que possível, forneça um `default`:

```go
switch dia {
case "segunda":
    fmt.Println("Início da semana")
default:
    fmt.Println("Dia qualquer")
}
```

---

## **Conclusão**

As estruturas condicionais de Go foram projetadas para **simplicidade, clareza e performance**. O `switch` é **mais poderoso e flexível** do que em muitas outras linguagens, e a exigência de tipagem forte ajuda a evitar bugs sutis.

No próximo capítulo, exploraremos **laços de repetição (`for`, `range`)**, fundamentais para manipulação de coleções e execução iterativa de código. 🚀
