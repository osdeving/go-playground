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

Variáveis em Go são armazenadas na **stack (pilha)** ou **heap (montante de memória dinâmica)**, dependendo do contexto:

### **Stack vs. Heap**

- **Stack (Pilha):** Usada para variáveis locais e temporárias. Gerenciada automaticamente, com alta eficiência.
- **Heap (Montante):** Usada quando a alocação precisa persistir além do escopo da função. O garbage collector do Go gerencia isso.

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

## **Conclusão**

A declaração de variáveis em Go é direta, mas embute decisões importantes como:
- **Simplicidade de leitura (left-to-right)**.
- **Redução de complexidade em declarações comparado a C**.
- **Inferência de tipos com `:=`, mas restrita ao escopo local**.
- **Gerenciamento automático de memória entre stack e heap**.

No próximo capítulo, exploraremos os **tipos primitivos** e como eles influenciam o desempenho e a manipulação de dados em Go. 🚀
