# **2.5 Conversão de Tipos**

> "Em Go, tudo tem um tipo bem definido, e nada se converte magicamente. Se quiser mudar um tipo, faça isso de forma explícita e controlada." — Filosofia Go

Go é uma linguagem **fortemente tipada**, o que significa que não realiza conversões implícitas entre tipos diferentes. Isso evita erros sutis e melhora a previsibilidade do código. Nesta seção, veremos como converter valores corretamente entre diferentes tipos, abordando desde números e strings até booleanos e slices de bytes.

---

## **2.5.1 Conversão Entre Tipos Numéricos**

Go não permite operações diretas entre tipos numéricos diferentes. Se tentarmos somar um `int` com um `float64`, por exemplo, teremos um erro de compilação:

```go
var a int = 10
var b float64 = 5.5

// fmt.Println(a + b) // ERRO: Tipos incompatíveis
```

Para resolver isso, devemos **converter explicitamente**:

```go
resultado := float64(a) + b // Correto
fmt.Println(resultado) // 15.5
```

📌 **Regra geral**: use `tipo(valor)` para converter valores.

### **Conversão de Tipos Inteiros**

```go
var x int32 = 100
var y int64 = int64(x) // Conversão explícita
fmt.Println(y) // 100
```

### **Conversão de `float` para `int` (Perda de Precisão)**

```go
var f float64 = 3.99
var i int = int(f)
fmt.Println(i) // 3 (trunca o valor, sem arredondamento)
```

📌 **A conversão de `float` para `int` simplesmente descarta a parte decimal, sem arredondamento!**

Se precisar arredondar, use `math.Round`:

```go
import "math"

var f float64 = 3.99
var i int = int(math.Round(f))
fmt.Println(i) // 4
```

🔹 **Dica**: Sempre considere se a conversão pode levar a perda de precisão antes de usá-la.

---

## **2.5.2 Conversão Entre `string` e Números**

Go não converte números para `string` automaticamente. Para fazer isso, usamos o pacote `strconv`.

### **De Número para `string`**

```go
import "strconv"

var num int = 42
var str string = strconv.Itoa(num) // int → string
fmt.Println(str) // "42"
```

Para `float64`:

```go
var f float64 = 3.14
var str string = strconv.FormatFloat(f, 'f', 2, 64) // float → string
fmt.Println(str) // "3.14"
```

📌 **Explicação de `FormatFloat(f, 'f', 2, 64)`**:
- `'f'` → Formato decimal (`'e'` para notação científica).
- `2` → Número de casas decimais.
- `64` → Precisão do float.

### **De `string` para Número**

Para converter `string` em número:

```go
num, err := strconv.Atoi("42") // string → int
if err != nil {
    fmt.Println("Erro:", err)
}
fmt.Println(num) // 42
```

Para `float64`:

```go
f, err := strconv.ParseFloat("3.14", 64) // string → float64
fmt.Println(f) // 3.14
```

📌 **Sempre trate erros ao converter strings para números!**

```go
num, err := strconv.Atoi("abc") // ERRO!
if err != nil {
    fmt.Println("Erro ao converter:", err)
}
```

---

## **2.5.6 Pratique Go**

🎯 Agora que você aprendeu sobre conversão de tipos, tente os seguintes desafios:

🔨 **Desafios**:

1️⃣ Converta um número inteiro para `string` e concatene-o a outra `string`.

2️⃣ Faça um programa que receba um número em formato de `string` e retorne o dobro desse número.

3️⃣ Converta uma `string` em uma slice de bytes e depois reconverta para `string`.

4️⃣ Escreva um programa que converta um `bool` para `int` e vice-versa sem erro de compilação.

5️⃣ Converta uma `string` contendo um número binário para um inteiro decimal.

6️⃣ Converta uma `string` contendo um número hexadecimal para um inteiro decimal.

7️⃣ Converta uma `string` contendo um número octal para um inteiro decimal.

8️⃣ Teste a conversão de números negativos entre `float64` e `int`.

9️⃣ Tente converter uma `string` vazia para um número e veja o que acontece.

🔷 Crie uma função genérica para conversão de tipos numéricos.

---

## **2.5.7 Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

1️⃣ O que acontece se tentarmos converter `float64` para `int`?

2️⃣ Qual pacote deve ser usado para converter `string` em `int`?

3️⃣ O que acontece se tentarmos converter `bool` diretamente para `int`?

4️⃣ Como garantir que uma conversão `float → int` arredonde corretamente?

5️⃣ Como evitar perda de precisão ao converter `float64` para `string`?

6️⃣ Qual é a forma correta de converter uma `string` para um `rune` em Go?

7️⃣ Como lidar com erros ao converter `string` para número?

8️⃣ Por que Go não permite conversão implícita entre tipos numéricos?

9️⃣ O que `strconv.ParseFloat("3.14abc", 64)` retorna?

🔷 Como converter um número em base diferente (binário, octal, hex) para decimal?

---



## **Conclusão**

🚀 **Resumo Final:**

Go exige **conversões explícitas** para garantir segurança de tipos e evitar bugs sutis. Entender como converter corretamente entre tipos evita problemas comuns e melhora a confiabilidade do código. No próximo capítulo, veremos **estruturas de controle de fluxo**, essenciais para criar lógicas dinâmicas no Go! 🔥

