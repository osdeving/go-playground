# **2.2 Tipos Primitivos (`int`, `float64`, `bool`, `string`)**

Os tipos primitivos em Go são os blocos fundamentais para armazenar e manipular dados. Diferente de linguagens como Python e JavaScript, Go **é estaticamente tipado**, o que significa que cada variável tem um tipo fixo determinado em tempo de compilação.

---

## **2.2.1 Visão Geral dos Tipos Primitivos**

Os principais tipos primitivos em Go são:

| Tipo       | Descrição |
|------------|------------|
| `int`      | Números inteiros com largura dependente da arquitetura (32 ou 64 bits) |
| `int8` a `int64` | Versões específicas de inteiros com tamanho fixo |
| `uint`, `uint8` a `uint64` | Inteiros sem sinal |
| `float32`, `float64` | Números de ponto flutuante |
| `bool`     | Representa valores `true` ou `false` |
| `string`   | Cadeia de caracteres UTF-8 |

---

## **2.2.2 Inteiros (`int`, `uint`, `int8` a `int64`)**

Go oferece diferentes tamanhos de inteiros:

```go
var a int = 42       // Inteiro padrão (depende da arquitetura)
var b int8 = -128    // Ocupa 8 bits (mín: -128, máx: 127)
var c uint16 = 65535 // Sem sinal, ocupa 16 bits
var d int64 = 9223372036854775807 // Suporta grandes valores
```

### **Escolha do Tipo de Inteiro**
- Use **`int`** para valores inteiros comuns (o compilador otimiza para `int32` ou `int64` conforme necessário).
- Use **`intX` e `uintX`** para controle fino de memória ou interoperabilidade com estruturas binárias.

### **Conversão entre Tipos Inteiros**

Go não realiza **conversão implícita** entre tipos diferentes:

```go
var x int32 = 100
var y int64 = int64(x) // Conversão explícita necessária
```

📌 **Evite misturar tipos diferentes em operações matemáticas para evitar erros de compilação.**

---

## **2.2.3 Números de Ponto Flutuante (`float32`, `float64`)**

Go suporta apenas dois tipos de números de ponto flutuante:

```go
var f1 float32 = 3.14
var f2 float64 = 2.718281828459045
```

### **Precisão dos Tipos Float**
- **`float32`**: Menos preciso, ocupa 4 bytes.
- **`float64`**: Mais preciso, ocupa 8 bytes (**padrão recomendado**).

Por padrão, o Go assume **`float64`** para valores de ponto flutuante:

```go
pi := 3.1415926535 // float64 por padrão
```

---

## **2.2.4 Booleanos (`bool`)**

O tipo `bool` representa valores lógicos `true` ou `false`:

```go
var ativo bool = true
desativado := false // Inferência automática
```

📌 **Go não permite conversões implícitas para `bool`**, então expressões como estas são inválidas:

```go
var x int = 10
if x { // ERRO: "x" não é booleano
    fmt.Println("Inválido!")
}
```

Para verificar se um número é diferente de zero, faça:

```go
if x != 0 { // Correto
    fmt.Println("Número não é zero!")
}
```

---

## **2.2.5 Strings (`string`)**

Go utiliza **strings imutáveis** codificadas em **UTF-8**.

```go
var nome string = "Golang"
mensagem := "Aprendendo Go!"
```

### **Caracteres em Go (`rune`)**

Diferente de outras linguagens, **Go não tem um tipo `char`**, mas permite representar caracteres como `rune`:

```go
var letra rune = 'G' // Representado por aspas simples
fmt.Println(letra)   // Saída: 71 (código Unicode de 'G')
```

📌 **Strings são imutáveis**: não é possível modificar um caractere diretamente:

```go
s := "GoLang"
s[0] = 'X' // ERRO: Strings são imutáveis
```

Se precisar modificar uma string, converta para **`[]rune`**:

```go
s := []rune("GoLang")
s[0] = 'X'
fmt.Println(string(s)) // "XoLang"
```

### **Concatenação de Strings**

```go
s1 := "Hello"
s2 := "World"
resultado := s1 + " " + s2 // "Hello World"
```

---

## **2.2.6 Zero Values e Inicialização**

Go atribui **zero values** automaticamente a variáveis não inicializadas:

| Tipo       | Zero Value |
|------------|------------|
| `int`      | `0` |
| `float64`  | `0.0` |
| `bool`     | `false` |
| `string`   | `""` (vazia) |

Exemplo:

```go
var x int  // 0
var y bool // false
var z string // ""
```


---

## **Pratique Go**

🎯 Agora que você aprendeu sobre os tipos primitivos em Go, tente os seguintes desafios:

🛠️ **Desafios**:

<details>
  <summary>✅ Crie um programa que declare variáveis de todos os tipos primitivos (`int`, `float64`, `bool`, `string`) e exiba seus valores iniciais.</summary>

  ```go
  package main
  import "fmt"
  
  func main() {
      var inteiro int
      var flutuante float64
      var booleano bool
      var texto string
  
      fmt.Println("int:", inteiro)
      fmt.Println("float64:", flutuante)
      fmt.Println("bool:", booleano)
      fmt.Println("string:", texto)
  }
  ```
</details>

<details>
  <summary>✅ Declare uma variável do tipo `int`, atribua um valor e converta para `float64`.</summary>

  ```go
  var x int = 42
  var y float64 = float64(x)
  fmt.Println(y) // 42.0
  ```
</details>

<details>
  <summary>✅ Escreva um programa que peça ao usuário para inserir um número decimal (`float64`) e depois converta para um número inteiro (`int`).</summary>
  
  ```go
  package main
  import (
      "fmt"
  )
  
  func main() {
      var num float64
      fmt.Print("Digite um número decimal: ")
      fmt.Scan(&num)
  
      inteiro := int(num)
      fmt.Println("Valor inteiro:", inteiro)
  }
  ```
</details>

<details>
  <summary>✅ Leia um valor booleano (`true` ou `false`) do usuário e inverta seu valor.</summary>
  
  ```go
  package main
  import "fmt"
  
  func main() {
      var valor bool
      fmt.Print("Digite true ou false: ")
      fmt.Scan(&valor)
      fmt.Println("Valor invertido:", !valor)
  }
  ```
</details>

<details>
  <summary>✅ Converta um número (`int`) em uma string e concatene com outra string.</summary>
  
  ```go
  import "strconv"
  
  var numero int = 100
  var texto string = "O valor "+ strconv.Itoa(numero)
  fmt.Println(texto) // "O valor 100"
  ```
</details>

<details>
  <summary>✅ Converta uma `string` contendo um número para `int` e realize operações matemáticas.</summary>
  
  ```go
  import "strconv"
  
  var strNum string = "50"
  num, _ := strconv.Atoi(strNum)
  fmt.Println(num + 10) // 60
  ```
</details>

<details>
  <summary>✅ Declare uma variável `string`, transforme todos os caracteres em maiúsculas e exiba o resultado.</summary>
  
  ```go
  import (
      "fmt"
      "strings"
  )
  
  func main() {
      texto := "golang"
      fmt.Println(strings.ToUpper(texto)) // "GOLANG"
  }
  ```
</details>

<details>
  <summary>✅ Crie um programa que armazene um número como `int`, o converta para binário e exiba sua representação binária.</summary>
  
  ```go
  package main
  import "fmt"
  
  func main() {
      var numero int = 42
      fmt.Printf("Binário: %b\n", numero) // "Binário: 101010"
  }
  ```
</details>

<details>
  <summary>✅ Faça um programa que utilize `reflect.TypeOf` para exibir o tipo de cada variável declarada.</summary>
  
  ```go
  import (
      "fmt"
      "reflect"
  )
  
  func main() {
      var x int = 10
      fmt.Println("Tipo de x:", reflect.TypeOf(x)) // "int"
  }
  ```
</details>

<details>
  <summary>✅ Escreva um programa que leia um nome e um número, formatando a saída como: `"O nome inserido foi <nome> e o número foi <número>"`.</summary>
  
  ```go
  package main
  import "fmt"
  
  func main() {
      var nome string
      var numero int
  
      fmt.Print("Digite seu nome: ")
      fmt.Scan(&nome)
      fmt.Print("Digite um número: ")
      fmt.Scan(&numero)
  
      fmt.Printf("O nome inserido foi %s e o número foi %d\n", nome, numero)
  }
  ```
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>
    💡 Qual a diferença entre `int`, `int32` e `int64`?
  </summary>

  O tamanho de `int` depende da arquitetura do sistema, enquanto `int32` e `int64` possuem tamanhos fixos de 32 e 64 bits, respectivamente.

</details>

<details>
  <summary>
    💡 O que acontece se tentarmos armazenar um número negativo em uma variável do tipo `uint`?
  </summary>
  
  O compilador gera um erro, pois `uint` não aceita valores negativos.

</details>

<details>
  <summary>
    💡 Como Go trata números de ponto flutuante (`float32` vs `float64`)?
  </summary>

  `float64` tem maior precisão do que `float32`, e Go usa `float64` como padrão em operações de ponto flutuante.

</details>

<details>
  <summary>
    💡 O que acontece ao converter um `float64` para `int`? Existe arredondamento?
  </summary>
  
  O valor decimal é truncado (não arredondado), removendo a parte decimal.

</details>

<details>
  <summary>
    💡 Como verificar o tipo de uma variável em tempo de execução?
  </summary>
  
  Usando `reflect.TypeOf(variavel)`.

</details>

<details>
  <summary>
    💡 Qual a diferença entre uma `string` e um slice de `byte` (`[]byte`)?
  </summary>

  `string` é imutável e `[]byte` permite modificação dos caracteres.
  
</details>

---




## **Conclusão**

🚀 **Resumo Final:**

Os tipos primitivos de Go são simples, mas altamente otimizados para eficiência e segurança. Seu modelo de tipagem estática reduz erros e melhora o desempenho. No próximo capítulo, exploraremos os **operadores e expressões em Go**! 🚀
