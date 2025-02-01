# **2.5 Conversão de Tipos**

Go é uma linguagem **fortemente tipada**, o que significa que não realiza conversões implícitas entre tipos diferentes. Se precisar converter um valor de um tipo para outro, o desenvolvedor deve fazê-lo explicitamente.

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
fmt.Println(y)
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

## **2.5.3 Conversão de `string` para `[]byte` e `rune`**

### **1. De `string` para `[]byte`**

```go
s := "GoLang"
b := []byte(s) // Converte string para slice de bytes
fmt.Println(b) // [71 111 76 97 110 103]
```

📌 Útil ao trabalhar com arquivos ou manipulação binária.

### **2. De `string` para `rune`**

```go
r := []rune("Golang")
fmt.Println(r) // [71 111 108 97 110 103]
```

📌 Útil para lidar com caracteres Unicode.

---

## **2.5.4 Conversão Entre `bool` e `string`**

Para converter `bool` para `string`:

```go
import "strconv"

var ativo bool = true
var str string = strconv.FormatBool(ativo)
fmt.Println(str) // "true"
```

Para converter `string` para `bool`:

```go
b, err := strconv.ParseBool("true")
fmt.Println(b) // true
```

📌 `ParseBool` aceita `"true"`, `"false"`, `"1"` e `"0"`, mas retorna erro para outros valores.

---

## **2.5.5 Casos Especiais e Armadilhas**

1. **Conversão de Inteiro para `string` via `string(x)` NÃO FUNCIONA como esperado!**

```go
x := 65
s := string(x)
fmt.Println(s) // "A" (código Unicode de 65)
```

📌 Para converter corretamente, use `strconv.Itoa(x)`, e não `string(x)`!

2. **Conversão de `bool` para `int` não existe diretamente**:

```go
var b bool = true
// x := int(b) // ERRO!
```

Para contornar:

```go
x := 0
if b {
    x = 1
}
fmt.Println(x) // 1
```

---

## **Conclusão**

Go exige **conversões explícitas** para garantir segurança de tipos e evitar bugs sutis. No próximo capítulo, veremos **estruturas de controle de fluxo**! 🚀
