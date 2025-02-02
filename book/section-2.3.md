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

## **Pratique Go**

🎯 Agora que você aprendeu sobre operadores, tente os seguintes desafios:

🔨 **Desafios**:

1️⃣ Implemente uma função que receba dois números inteiros e retorne a soma, subtração, multiplicação e divisão como múltiplos valores de retorno.

2️⃣ Crie um programa que utilize operadores bit a bit (&, |, ^, &^) para manipular bits e converter entre representações binárias e decimais.

3️⃣ Escreva uma função que verifique se um número inteiro é par ou ímpar sem usar operadores de comparação (==, !=, <, >).

4️⃣ Implemente um contador de bits 1 que conte quantos bits estão ativados (1) em um número inteiro sem usar laços (for/range).

5️⃣ Construa um mini interpretador de expressões matemáticas, aceitando entradas como "3 + 5 * 2" e calculando o resultado corretamente, respeitando a precedência de operadores.

6️⃣ Utilize operadores lógicos para criar uma função que simule um controle de acesso: dado um usuário e uma senha corretos, conceda ou negue o acesso.

7️⃣ Implemente um circuito digital básico usando operadores bit a bit, simulando AND, OR, XOR e NOT entre dois números binários.

8️⃣ Faça um programa que determine se um número inteiro é uma potência de 2 usando operadores bit a bit.

9️⃣ Escreva uma função que utilize operadores de shift (<<, >>) para multiplicar e dividir números inteiros por potências de 2 sem usar * ou /.

🔟 Crie uma função que troque o valor de duas variáveis sem utilizar variáveis auxiliares, apenas com operadores matemáticos ou bit a bit.

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

1️⃣ O que acontece ao dividir um número inteiro por outro número inteiro em Go? Como evitar perda de precisão?

2️⃣ Qual é o comportamento do operador % (módulo) para números negativos? -10 % 3 resulta em qual valor?

3️⃣ Por que Go não permite o uso de ++ e -- dentro de expressões?

4️⃣ Como Go lida com short-circuit evaluation nos operadores && e ||?

5️⃣ O que acontece ao comparar tipos diferentes (int e float64)? Como evitar esse problema?

6️⃣ Qual é a precedência correta dos operadores em Go? Quais têm maior prioridade?

7️⃣ Como evitar problemas ao usar operadores bit a bit (&, |, ^) para manipulação de permissões e flags?

8️⃣ Como Go trata a conversão automática de tipos em operações aritméticas? Existe type promotion como em C?

9️⃣ Qual a forma correta de utilizar &^ para limpar um bit específico dentro de um número?

🔟 Em quais cenários o uso de operadores bit a bit pode ser mais eficiente do que operadores matemáticos convencionais?

## **Conclusão**

🚀 **Resumo Final:**

Os operadores em Go são projetados para serem simples e previsíveis, seguindo regras rigorosas de tipagem. A ausência de conversões implícitas reduz erros sutis e melhora a clareza do código. Além disso, a decisão de não permitir ++ e -- dentro de expressões evita ambiguidades.

A compreensão profunda de operadores matemáticos, lógicos e bit a bit é fundamental para escrever código eficiente, especialmente ao lidar com manipulação de bits, sistemas de permissões e otimizações de desempenho.

No próximo capítulo, exploraremos entrada e saída de dados com fmt, incluindo formatação avançada! 🚀

