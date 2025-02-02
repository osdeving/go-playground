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


## **Pratique Go**

🎯 Agora que você aprendeu sobre os tipos primitivos em Go, tente os seguintes desafios:

🔨 **Desafios**:

✅ Crie um programa que declare variáveis de todos os tipos primitivos (`int`, `float64`, `bool`, `string`) e exiba seus valores iniciais.  

✅ Declare uma variável do tipo `int`, atribua um valor e converta para `float64`.  

✅ Escreva um programa que peça ao usuário para inserir um número decimal (`float64`) e depois converta para um número inteiro (`int`).  

✅ Leia um valor booleano (`true` ou `false`) do usuário e inverta seu valor.  

✅ Converta um número (`int`) em uma string e concatene com outra string.  

✅ Converta uma `string` contendo um número para `int` e realize operações matemáticas.  

✅ Declare uma variável `string`, transforme todos os caracteres em maiúsculas e exiba o resultado.  

✅ Crie um programa que armazene um número como `int`, o converta para binário e exiba sua representação binária.  

✅ Faça um programa que utilize `reflect.TypeOf` para exibir o tipo de cada variável declarada.  

✅ Escreva um programa que leia um nome e um número, formatando a saída como:  
   `"O nome inserido foi <nome> e o número foi <número>"`.  


## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

💡 Qual a diferença entre `int`, `int32` e `int64`?  

💡 O que acontece se tentarmos armazenar um número negativo em uma variável do tipo `uint`?  

💡 Como Go trata números de ponto flutuante (`float32` vs `float64`)?  

💡 O que acontece ao converter um `float64` para `int`? Existe arredondamento?  

💡 Como verificar o tipo de uma variável em tempo de execução?  

💡 Qual a diferença entre uma `string` e um slice de `byte` (`[]byte`)?  

💡 O que acontece ao tentar concatenar uma `string` e um `int` diretamente?  

💡 Como transformar um número decimal em uma string contendo sua representação binária?  

💡 Qual o valor padrão de `bool` em Go?  

💡 O que acontece se tentarmos converter uma `string` vazia em um número?  


## **Conclusão**

🚀 **Resumo Final:**

Os tipos primitivos de Go são simples, mas altamente otimizados para eficiência e segurança. Seu modelo de tipagem estática reduz erros e melhora o desempenho. No próximo capítulo, exploraremos os **operadores e expressões em Go**! 🚀
