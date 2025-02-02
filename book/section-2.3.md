# **2.3 Operadores Aritméticos, Lógicos e Comparativos**

>⚡ "Entender os operadores é essencial para construir qualquer programa eficiente. Seja realizando cálculos, comparações ou lógica condicional, cada operador tem seu papel. Dominar sua precedência e comportamento evita armadilhas e torna seu código mais expressivo e seguro." — Go Proverbs



Os operadores são fundamentais em **Go** para realizar cálculos, comparações e operações lógicas. A sintaxe de Go é intuitiva, mas possui algumas regras específicas que diferem de outras linguagens.

---

## **2.3.1 Operadores Aritméticos**

Go suporta os operadores matemáticos tradicionais:

| Operador | Descrição | Exemplo |
|----------|------------|---------|
| `+`  | Adição        | `a + b` |
| `-`  | Subtração     | `a - b` |
| `*`  | Multiplicação | `a * b` |
| `/`  | Divisão       | `a / b` |
| `%`  | Módulo (resto) | `a % b` |

Exemplo:

```go
a := 10
b := 3

fmt.Println(a + b)  // 13
fmt.Println(a - b)  // 7
fmt.Println(a * b)  // 30
fmt.Println(a / b)  // 3 (inteiro, sem casas decimais)
fmt.Println(a % b)  // 1 (resto da divisão)
```

📌 **Divisão entre inteiros descarta a parte decimal!** Para manter precisão, converta para `float64`:

```go
c := float64(a) / float64(b) // 3.333333
```

### **Incremento e Decremento (`++`, `--`)** 🚨

Diferente de C e Java, Go **não permite** `x++` ou `x--` em expressões! Isso pode causar surpresa para desenvolvedores acostumados com outras linguagens.

```go
x := 5
x++  // Ok
// fmt.Println(x++) // ❌ ERRO! Incremento não pode estar dentro de expressões
```

📌 **Motivo:** Essa decisão foi tomada para evitar ambiguidades e efeitos colaterais que surgem quando `++` e `--` são usados dentro de expressões complexas. Em Go, o incremento e decremento devem ser usados como instruções separadas.

---

## **2.3.2 Operadores de Comparação**

Go possui os operadores clássicos de comparação:

| Operador | Descrição | Exemplo |
|----------|------------|---------|
| `==` | Igualdade | `x == y` |
| `!=` | Diferença | `x != y` |
| `>`  | Maior que | `x > y` |
| `<`  | Menor que | `x < y` |
| `>=` | Maior ou igual | `x >= y` |
| `<=` | Menor ou igual | `x <= y` |

📌 **Os operadores de comparação só podem ser usados em valores do mesmo tipo!** Isso evita bugs comuns em linguagens como JavaScript, onde comparações entre tipos podem levar a resultados inesperados.

```go
var a int = 10
var b float64 = 10.0

// fmt.Println(a == b) // ❌ ERRO! Tipos diferentes
fmt.Println(float64(a) == b) // ✅ true (após conversão)
```

---

## **2.3.3 Operadores Lógicos (`&&`, `||`, `!`)**

Os operadores lógicos são usados para combinar expressões booleanas:

| Operador | Descrição | Exemplo |
|----------|------------|---------|
| `&&` | E lógico (AND) | `(x > 0) && (y > 0)` |
| `\|\|` | OU lógico (OR) | `(x > 0) \|\| (y > 0)` |
| `!`  | Negação (NOT)  | `!(x > 0)` |

📌 **Short-circuit evaluation**: Em uma operação `&&`, se a primeira condição for `false`, a segunda não é avaliada. Em `||`, se a primeira for `true`, a segunda não é avaliada.

```go
func expensiveCheck() bool {
    fmt.Println("Executando verificação cara...")
    return true
}

if false && expensiveCheck() {
    fmt.Println("Não será impresso.")
}
```

---

## **2.3.4 Operadores de Atribuição Combinada**

Além das atribuições comuns, Go oferece operadores de atribuição combinada para simplificar expressões:

| Operador | Exemplo | Equivalente a |
|----------|---------|---------------|
| `+=`  | `x += 5`  | `x = x + 5` |
| `-=`  | `x -= 3`  | `x = x - 3` |
| `*=`  | `x *= 2`  | `x = x * 2` |
| `/=`  | `x /= 4`  | `x = x / 4` |
| `%=`  | `x %= 2`  | `x = x % 2` |
| `&=`  | `x &= y`  | `x = x & y` |
| `\|=`  | `x \|= y`  | `x = x \| y` |
| `^=`  | `x ^= y`  | `x = x ^ y` |
| `&^=` | `x &^= y` | `x = x &^ y` |

---

## **Exemplo Prático**

Vamos consolidar tudo que aprendemos até agora em um exemplo prático:

```go
package main

import "fmt"

func main() {
    a, b := 10, 5
    fmt.Println("Operações básicas:")
    fmt.Println("Soma:", a + b)
    fmt.Println("Subtração:", a - b)
    fmt.Println("Multiplicação:", a * b)
    fmt.Println("Divisão:", a / b)
    fmt.Println("Resto:", a % b)

    fmt.Println("\nOperações lógicas:")
    fmt.Println("a > b && a > 0:", a > b && a > 0)
    fmt.Println("a < b || b > 0:", a < b || b > 0)
    fmt.Println("!(a == b):", !(a == b))

    fmt.Println("\nAtribuições combinadas:")
    a += 3
    fmt.Println("a += 3:", a)
    a &= 7
    fmt.Println("a &= 7:", a)
}
```

📌 **Este exemplo mostra como aplicar operadores matemáticos, lógicos e de atribuição em um contexto real.**

---

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre operadores, tente os seguintes desafios:

🛠️ **Desafios**:

<details>
  <summary>✅ Implemente uma função que receba dois números inteiros e retorne a soma, subtração, multiplicação e divisão como múltiplos valores de retorno.</summary>
  
  ```go
  func operacoes(a, b int) (int, int, int, float64) {
      return a + b, a - b, a * b, float64(a) / float64(b)
  }
  ```
  
</details>

<details>
  <summary>✅ Crie um programa que utilize operadores bit a bit (&, |, ^, &^) para manipular bits e converter entre representações binárias e decimais.</summary>
  
  ```go
  func manipulaBits(a, b int) {
      fmt.Printf("AND: %b\n", a & b)
      fmt.Printf("OR: %b\n", a | b)
      fmt.Printf("XOR: %b\n", a ^ b)
      fmt.Printf("AND NOT: %b\n", a &^ b)
  }
  ```
  
</details>

<details>
  <summary>✅ Escreva uma função que verifique se um número inteiro é par ou ímpar sem usar operadores de comparação (==, !=, <, >).</summary>
  
  ```go
  func ehPar(n int) bool {
      return n & 1 == 0
  }
  ```
  
</details>

<details>
  <summary>✅ Implemente um contador de bits 1 que conte quantos bits estão ativados (1) em um número inteiro sem usar laços (for/range).</summary>
  
  ```go
  func contarBits(n uint) int {
      return bits.OnesCount(n)
  }
  ```
  
</details>

<details>
  <summary>✅ Construa um mini interpretador de expressões matemáticas, aceitando entradas como "3 + 5 * 2" e calculando o resultado corretamente, respeitando a precedência de operadores.</summary>
  
  ```go
  func calcularExpressao(expr string) (int, error) {
      return eval(expr) // Supondo uma implementação de parser
  }
  ```
  
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>💡 O que acontece ao dividir um número inteiro por outro número inteiro em Go? Como evitar perda de precisão?</summary>
  A divisão entre inteiros descarta a parte decimal. Para evitar isso, converta um dos operandos para `float64` antes da divisão.
</details>

<details>
  <summary>💡 Qual é o comportamento do operador % (módulo) para números negativos? `-10 % 3` resulta em qual valor?</summary>
  O operador `%` segue o sinal do dividendo. `-10 % 3` resulta em `-1`.
</details>

<details>
  <summary>💡 Por que Go não permite o uso de ++ e -- dentro de expressões?</summary>
  Para evitar ambiguidades e efeitos colaterais na avaliação das expressões.
</details>

<details>
  <summary>💡 Como Go lida com short-circuit evaluation nos operadores && e ||?</summary>
  Se a primeira condição de `&&` for falsa, a segunda não é avaliada. Se a primeira de `||` for verdadeira, a segunda também não é avaliada.
</details>

<details>
  <summary>💡 O que acontece ao comparar tipos diferentes (int e float64)? Como evitar esse problema?</summary>
  Go não permite comparação direta entre tipos diferentes. Para evitar erros, converta explicitamente para um tipo comum antes da comparação.
</details>

<details>
  <summary>💡 Qual é a precedência correta dos operadores em Go? Quais têm maior prioridade?</summary>
  Operadores aritméticos (`*`, `/`, `%`) têm maior precedência do que `+` e `-`, seguidos por operadores de comparação e lógicos.
</details>

<details>
  <summary>💡 Como evitar problemas ao usar operadores bit a bit (&, |, ^) para manipulação de permissões e flags?</summary>
  Use **máscaras de bits** e operadores bit a bit corretamente para definir, limpar e verificar flags.
</details>

<details>
  <summary>💡 Como Go trata a conversão automática de tipos em operações aritméticas? Existe type promotion como em C?</summary>
  Go não faz conversão implícita de tipos em operações aritméticas. Todas as operações devem ser feitas entre operandos do mesmo tipo.
</details>

<details>
  <summary>💡 Qual a forma correta de utilizar &^ para limpar um bit específico dentro de um número?</summary>
  `x &^ (1 << pos)` pode ser usado para limpar o bit na posição `pos` dentro de `x`.
</details>

<details>
  <summary>💡 Em quais cenários o uso de operadores bit a bit pode ser mais eficiente do que operadores matemáticos convencionais?</summary>
  Em criptografia, compressão de dados e manipulação de flags de controle, onde operações bit a bit são mais eficientes.
</details>

---



## **Conclusão**

🚀 **Resumo Final:**

Os operadores em Go são projetados para serem simples e previsíveis, seguindo regras rigorosas de tipagem. A ausência de conversões implícitas reduz erros sutis e melhora a clareza do código. Além disso, a decisão de não permitir ++ e -- dentro de expressões evita ambiguidades.

A compreensão profunda de operadores matemáticos, lógicos e bit a bit é fundamental para escrever código eficiente, especialmente ao lidar com manipulação de bits, sistemas de permissões e otimizações de desempenho.

No próximo capítulo, exploraremos entrada e saída de dados com fmt, incluindo formatação avançada! 🚀

