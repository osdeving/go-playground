# **2.3 Operadores Aritméticos, Lógicos e Comparativos**

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

## **Conclusão**

Os operadores de Go são simples, mas seguem regras rígidas de tipagem e precedência. O suporte a operadores bit a bit permite manipulação de baixo nível quando necessário. Também vimos operadores adicionais úteis para manipulação de ponteiros e mapas.

No próximo capítulo, exploraremos **entrada e saída de dados com `fmt`**, incluindo formatação avançada! 🚀

