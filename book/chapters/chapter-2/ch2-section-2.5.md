# **2.5 Conversão de Tipos**

> "Type systems are the most cost effective unit tests that exist. They are a scaffold that lets you refactor fearlessly." — Steve Yegge, ex-engenheiro do Google e Amazon

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

## **Pratique Go**

🎯 Agora que você aprendeu sobre conversão de tipos, tente os seguintes desafios:

🔨 **Desafios**:

<details>
  <summary>1️⃣ Converta um número inteiro para `string` e concatene-o a outra `string`.</summary>
  
  ```go
  num := 42
  str := "O resultado é: " + strconv.Itoa(num)
  fmt.Println(str) // "O resultado é: 42"
  ```
</details>

<details>
  <summary>2️⃣ Faça um programa que receba um número em formato de `string` e retorne o dobro desse número.</summary>
  
  ```go
  input := "21"
  num, _ := strconv.Atoi(input)
  fmt.Println(num * 2) // 42
  ```
</details>

<details>
  <summary>3️⃣ Converta uma `string` em uma slice de bytes e depois reconverta para `string`.</summary>
  
  ```go
  s := "GoLang"
  b := []byte(s)
  s2 := string(b)
  fmt.Println(s2) // "GoLang"
  ```
</details>

<details>
  <summary>4️⃣ Escreva um programa que converta um `bool` para `int` e vice-versa sem erro de compilação.</summary>
  
  ```go
  var b bool = true
  var i int
  if b {
      i = 1
  } else {
      i = 0
  }
  fmt.Println(i) // 1
  ```
</details>

<details>
  <summary>5️⃣ Converta uma `string` contendo um número binário para um inteiro decimal.</summary>
  
  ```go
  bin := "1010"
  num, _ := strconv.ParseInt(bin, 2, 64)
  fmt.Println(num) // 10
  ```
</details>

<details>
  <summary>6️⃣ Converta uma `string` contendo um número hexadecimal para um inteiro decimal.</summary>
  
  ```go
  hex := "1A"
  num, _ := strconv.ParseInt(hex, 16, 64)
  fmt.Println(num) // 26
  ```
</details>

<details>
  <summary>7️⃣ Converta uma `string` contendo um número octal para um inteiro decimal.</summary>
  
  ```go
  oct := "12"
  num, _ := strconv.ParseInt(oct, 8, 64)
  fmt.Println(num) // 10
  ```
</details>

<details>
  <summary>8️⃣ Teste a conversão de números negativos entre `float64` e `int`.</summary>
  
  ```go
  f := -3.99
  i := int(f)
  fmt.Println(i) // -3 (sem arredondamento)
  ```
</details>

<details>
  <summary>9️⃣ Tente converter uma `string` vazia para um número e veja o que acontece.</summary>
  
  ```go
  num, err := strconv.Atoi("")
  fmt.Println(num, err) // 0, erro
  ```
</details>

<details>
  <summary>🔢 Crie uma função genérica para conversão de tipos numéricos.</summary>
  
  ```go
  func convert[T any](val T) string {
      return fmt.Sprintf("%v", val)
  }
  fmt.Println(convert(42))   // "42"
  fmt.Println(convert(3.14)) // "3.14"
  ```
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>1️⃣ O que acontece se tentarmos converter `float64` para `int`?</summary>
  O valor será truncado, removendo a parte decimal.
</details>

<details>
  <summary>2️⃣ Qual pacote deve ser usado para converter `string` em `int`?</summary>
  O pacote `strconv` fornece `strconv.Atoi` e `strconv.ParseInt`.
</details>

<details>
  <summary>3️⃣ O que acontece se tentarmos converter `bool` diretamente para `int`?</summary>
  Go não permite essa conversão diretamente. É necessário usar uma estrutura condicional.
</details>

<details>
  <summary>4️⃣ Como garantir que uma conversão `float → int` arredonde corretamente?</summary>
  Use `math.Round()` antes de converter.
</details>

<details>
  <summary>5️⃣ Como evitar perda de precisão ao converter `float64` para `string`?</summary>
  Use `strconv.FormatFloat(f, 'f', -1, 64)`.
</details>

<details>
  <summary>6️⃣ Qual é a forma correta de converter uma `string` para um `rune` em Go?</summary>
  Use `runes := []rune("texto")`.
</details>

<details>
  <summary>7️⃣ Como lidar com erros ao converter `string` para número?</summary>
  Sempre verifique o segundo valor de retorno (`err`) das funções `strconv.Atoi` e `strconv.ParseInt`.
</details>

<details>
  <summary>8️⃣ Por que Go não permite conversão implícita entre tipos numéricos?</summary>
  Para evitar erros sutis de perda de precisão e comportamento inesperado.
</details>

<details>
  <summary>9️⃣ O que `strconv.ParseFloat("3.14abc", 64)` retorna?</summary>
  Retorna um erro, pois o valor não é um float válido.
</details>

<details>
  <summary>🔢 Como converter um número em base diferente (binário, octal, hex) para decimal?</summary>
  Use `strconv.ParseInt(valor, base, 64)`, especificando a base adequada (2, 8, 16).
</details>

---





## **Conclusão**

🚀 **Resumo Final:**

Go exige **conversões explícitas** para garantir segurança de tipos e evitar bugs sutis. Entender como converter corretamente entre tipos evita problemas comuns e melhora a confiabilidade do código. No próximo capítulo, veremos **estruturas de controle de fluxo**, essenciais para criar lógicas dinâmicas no Go! 🔥

