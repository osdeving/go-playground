# **2.1 Declaração de Variáveis (`var`, `:=`)**

A declaração de variáveis é um dos conceitos fundamentais em Go. Embora simples à primeira vista, sua sintaxe reflete escolhas de design importantes, como a **leitura left-to-right**, a ausência de declarações complexas como em C e a forma como o modelo de memória influencia seu comportamento.

---

## **2.1.1 Forma Geral de Declaração de Variáveis**

Go permite a declaração de variáveis de duas formas principais:

### **1. Declaração Explícita (`var`)**

A palavra-chave `var` permite declarar variáveis com ou sem inicialização explícita:

```go
var x int // x recebe o valor zero do tipo (0 para int)
var y float64 = 3.14 // y recebe o valor 3.14
var nome string = "Golang" // nome recebe "Golang"
```

Se a variável for declarada sem valor, **Go atribui o zero value do tipo**:

| Tipo        | Zero Value |
|------------|------------|
| `int`      | `0`        |
| `float64`  | `0.0`      |
| `bool`     | `false`    |
| `string`   | `""` (string vazia) |
| `pointer`  | `nil`      |

---

### **2. Inferência com `:=`**

Go permite declarar e inicializar variáveis de forma implícita, inferindo o tipo automaticamente:

```go
x := 42         // int
nome := "Go"    // string
ativo := true   // bool
pi := 3.1415    // float64
```

📌 **Regras Importantes do `:=`**:
- **Só pode ser usado dentro de funções**. Fora delas, use `var`.
- **O tipo é inferido pelo valor atribuído**.
- **A declaração e a atribuição devem ocorrer simultaneamente** (diferente de `var`, que permite declaração sem inicialização).

---

## **2.1.2 A Escolha por Left-to-Right em Go**

Diferente de C, onde a declaração de variáveis pode ser complexa (`int *x, (*y)[10]`), Go segue a leitura **da esquerda para a direita**, reduzindo ambiguidades:

```go
var x, y int     // Dois inteiros
var p *int       // Ponteiro para um inteiro
var f func()     // Variável do tipo função sem parâmetros
```

Em C, uma declaração como `int *a, b;` pode levar a erros, pois `b` não é um ponteiro. Em Go, isso seria escrito de forma clara:

```go
var a *int // Ponteiro para int
var b int  // Inteiro normal
```

📌 **Benefício:** Elimina confusão na leitura e evita declarações crípticas.

---

## **2.1.3 Escopo e Tempo de Vida de Variáveis**

O escopo de uma variável em Go segue as regras padrões de blocos `{}`:

- **Variáveis declaradas em um bloco `{}`** são locais ao bloco e não existem fora dele.
- **Variáveis globais (declaradas fora de funções)** existem enquanto o programa estiver em execução.

Exemplo:

```go
package main

import "fmt"

var global = "Eu sou global" // Variável global

func main() {
    local := "Eu sou local" // Variável local

    if true {
        interna := "Escopo do bloco if"
        fmt.Println(interna) // Ok, visível dentro do bloco
    }

    // fmt.Println(interna) // ERRO: "interna" não existe aqui

    fmt.Println(global) // Ok
    fmt.Println(local)  // Ok
}
```

---

## **2.1.4 Modelo de Memória e Alocação**

Variáveis em Go são armazenadas na **stack (pilha)** ou **heap (espaço de memória dinâmica)**, dependendo do contexto:

### **Stack vs. Heap**

- **Stack:** Usada para variáveis locais e temporárias. Gerenciada automaticamente, com alta eficiência.
- **Heap:** Usada quando a alocação precisa persistir além do escopo da função. O garbage collector do Go gerencia isso.

```go
func exemplo() *int {
    x := 42  // Alocado na stack
    return &x // O Go detecta que `x` precisa ir para a heap
}
```

Aqui, `x` normalmente ficaria na stack, mas como seu endereço é retornado, o Go move `x` para a heap.

---

## **2.1.5 Declaração Múltipla e Atribuição**

Go permite declarar múltiplas variáveis em uma única linha:

```go
var a, b, c int  // Três inteiros
var nome, idade = "Alice", 30 // Inferência de tipo
x, y := 10, 20 // Duas variáveis inferidas
```

E também suporta **troca de valores sem variável auxiliar**:

```go
x, y = y, x // Swap direto
```

Essa abordagem evita código redundante e melhora a clareza.

---

## **2.1.6 Constantes (`const`)**

Além de variáveis mutáveis, Go permite declarar **constantes**, que não podem ser alteradas após a compilação:

```go
const Pi = 3.1415
const Nome = "Golang"
```

**Diferenças entre `const` e `var`**:

| `const` | `var` |
|---------|--------|
| Valor fixo na compilação | Pode ser alterado |
| Apenas valores literais ou expressões constantes | Pode ser atribuído dinamicamente |
| Melhor para otimização de código | Mais flexível |

---

## **2.1.7 Declarações em Bloco e Múltiplas Variáveis**

Go oferece formas elegantes de declarar múltiplas variáveis, seja em bloco ou em linha única.

### **Declaração em Bloco**

Usando `var()`, podemos agrupar declarações de variáveis de forma organizada:

```go
var (
    nome     string
    idade    int
    altura   float64
    ativo    bool
)
```

Esta sintaxe é especialmente útil para:
- Variáveis globais
- Grupos de variáveis relacionadas
- Melhor legibilidade em declarações múltiplas

### **Declarações Múltiplas em Linha**

Go permite declarar e inicializar múltiplas variáveis em uma única linha:

```go
// Com var e tipos inferidos
var nome, idade, altura, ativo = "Maria", 30, 1.65, true

// Com :=
nome, idade, altura, ativo := "João", 25, 1.75, true
```

⚠️ **Importante**: Não é possível declarar variáveis de tipos diferentes especificando os tipos em uma única linha:

```go
// Isto NÃO funciona:
var nome string, idade int, altura float64  // Erro de sintaxe!

// Iss também NÃO funciona:
var nome string, idade int, altura float64 = "João", 25, 1.75

// Forma correta:
var nome string
var idade int
var altura float64

// Ou usando bloco:
var (
    nome   string
    idade  int
    altura float64
)
```

### **Regras e Boas Práticas**

1. **Declaração em Bloco**:
   - Ideal para variáveis globais
   - Melhora organização do código
   - Facilita manutenção

2. **Declaração Múltipla em Linha**:
   - Útil para variáveis relacionadas
   - Requer inicialização de todas as variáveis
   - Tipos são inferidos dos valores

3. **Quando Usar Cada Uma**:
   ```go
   // Use blocos para variáveis não inicializadas ou globais
   var (
       config  string
       version int
       debug   bool
   )

   // Use linha única para variáveis locais relacionadas
   nome, sobrenome := "João", "Silva"
   largura, altura := 100, 200
   ```

🎯 **Exemplo Prático**:
```go
package main

import "fmt"

// Variáveis globais em bloco
var (
    appName    string = "MinhaApp"
    appVersion int    = 1
    debug      bool   = true
)

func main() {
    // Variáveis locais em linha única
    nome, idade := "Alice", 30
    
    // Múltiplas variáveis com tipos diferentes
    var x, y, msg = 10, 20.5, "teste"
    
    fmt.Println(nome, idade)    // Saída formatada básica
    fmt.Println(x, y, msg)
}
```

📝 **Nota**: A formatação de saída (usando `fmt.Printf`, `fmt.Println`, etc.) será explorada em detalhes na seção 2.4.

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre a declaração de variáveis em Go, tente os seguintes desafios:

🛠️ **Desafios**:

<details>
  <summary>✅ Declare variáveis usando `var` e `:=` e explique a diferença de escopo entre elas.</summary>
  
  `var` pode ser usada tanto dentro quanto fora de funções, enquanto `:=` só pode ser usada dentro de funções e infere o tipo automaticamente.
  
</details>

<details>
  <summary>✅ Tente declarar uma variável com `:=` fora de uma função. O que acontece?</summary>
  
  Um erro de compilação ocorre, pois `:=` só pode ser usado dentro de funções.
  
</details>

<details>
  <summary>✅ Declare uma variável com `var` e tente utilizá-la sem inicializar. Qual valor ela assume?</summary>
  
  Ela assume o **valor zero** do seu tipo. Exemplo: `int` será `0`, `string` será `""`, `bool` será `false`.
  
</details>

<details>
  <summary>✅ Crie uma variável global e acesse-a dentro de uma função. Go permite isso?</summary>
  
  Sim, variáveis globais podem ser acessadas dentro de funções, mas seu uso deve ser evitado para evitar efeitos colaterais.
  
</details>

<details>
  <summary>✅ Faça um programa que tente redefinir uma variável já declarada com `:=` no mesmo bloco. Funciona?</summary>
  
  Não, `:=` só pode ser usado para **declaração nova**. Para reatribuir, use apenas `=`.
  
</details>

<details>
  <summary>✅ Declare várias variáveis de tipos diferentes na mesma linha e atribua valores a elas.</summary>
  
  ```go
  var x, y, z = 10, "hello", true
  a, b, c := 3.14, 'A', 42
  ```
  
</details>

<details>
  <summary>✅ Crie uma constante (`const`) e tente alterá-la em tempo de execução. O que acontece?</summary>
  
  Constantes não podem ser modificadas após a compilação. Tentar reatribuí-las resultará em erro de compilação.
  
</details>

<details>
  <summary>✅ Utilize `reflect.TypeOf` para verificar dinamicamente o tipo de uma variável.</summary>
  
  ```go
  import "fmt"
  import "reflect"
  
  var x = 42
  fmt.Println(reflect.TypeOf(x)) // Output: int
  ```
  
</details>

<details>
  <summary>✅ Declare uma variável `string`, converta-a para `[]byte` e depois reconverta para `string`.</summary>
  
  ```go
  s := "GoLang"
  b := []byte(s)
  s2 := string(b)
  fmt.Println(s2) // GoLang
  ```
  
</details>

<details>
  <summary>✅ Crie um programa que utilize `var` e `:=` dentro de loops e funções aninhadas para analisar o escopo das variáveis.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var x = "fora"
      fmt.Println("Escopo externo:", x)
      
      func() {
          x := "dentro"
          fmt.Println("Escopo interno:", x)
      }()
      
      fmt.Println("Escopo externo novamente:", x)
  }
  ```
  
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>💡 Qual a diferença fundamental entre `var` e `:=` na declaração de variáveis?</summary>
  
  `var` pode ser usada em qualquer escopo e permite declaração sem inicialização, enquanto `:=` só pode ser usada dentro de funções e exige valor inicial.
  
</details>

<details>
  <summary>💡 O que acontece se tentarmos usar `:=` fora de uma função?</summary>
  
  Um erro de compilação ocorre porque `:=` é válido apenas dentro de funções.
  
</details>

<details>
  <summary>💡 Como Go trata variáveis não inicializadas? Elas possuem um valor padrão?</summary>
  
  Sim, Go atribui o **valor zero** do tipo à variável: `int` é `0`, `string` é `""`, `bool` é `false`, etc.
  
</details>

<details>
  <summary>💡 É possível reatribuir uma variável declarada com `:=` usando `myvar := novovalor` dentro do mesmo escopo?</summary>
  
  Não, `:=` só pode ser usada para **declaração nova**. Para reatribuir, use apenas `=`.
  
</details>

<details>
  <summary>💡 Qual a diferença entre `var x int` e `x := 0`? Alguma dessas abordagens é mais eficiente?</summary>
  
  `var x int` declara `x` com valor `0` implicitamente, enquanto `x := 0` infere o tipo. Em termos de desempenho, são equivalentes.
  
</details>

<details>
  <summary>💡 `var` pode ser usada dentro de uma função? E `:=` pode ser usada fora de uma função?</summary>
  
  `var` pode ser usada em qualquer lugar, inclusive fora de funções. `:=` só pode ser usada dentro de funções.
  
</details>

<details>
  <summary>💡 O que acontece ao declarar duas variáveis com o mesmo nome em escopos diferentes?</summary>
  
  A variável mais próxima ao escopo atual é usada, ocultando a variável externa.
  
</details>

<details>
  <summary>💡 Como Go diferencia variáveis locais e globais quando possuem o mesmo nome?</summary>
  
  A variável local tem precedência dentro da função, ocultando a global. Para acessar a global, use um nome diferente ou um pacote.
  
</details>

<details>
  <summary>💡 `const` pode ser declarada usando `:=`? Por quê?</summary>
  
  Não, pois `:=` é usado apenas para declaração de variáveis mutáveis, enquanto `const` deve ser definida com `const`.
  
</details>

<details>
  <summary>💡 Como podemos declarar múltiplas variáveis de tipos diferentes em uma única linha?</summary>
  
  ```go
  var x, y, z = 10, "hello", true
  ```
  
</details>

---

## **Conclusão**

🚀 **Resumo Final:**

A declaração de variáveis em Go é direta, mas reflete decisões importantes dos criadores como:
- **Simplicidade de leitura (left-to-right)**.
- **Redução de complexidade em declarações comparado a C**.
- **Inferência de tipos com `:=`, mas restrita ao escopo local**.
- **Gerenciamento automático de memória entre stack e heap**.

No próximo capítulo, exploraremos os **tipos primitivos** e como eles influenciam o desempenho e a manipulação de dados em Go. 🚀
