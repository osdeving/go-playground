# **2.3 Operadores Aritméticos, Lógicos e Comparativos**

Os operadores são fundamentais em Go para realizar cálculos, comparações e operações lógicas. A sintaxe de Go é intuitiva, mas possui algumas regras específicas que diferem de outras linguagens.

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

### **Incremento e Decremento (`++`, `--`)**

Diferente de C e Java, Go **não permite** `x++` ou `x--` em expressões:

```go
x := 5
x++  // Ok
// fmt.Println(x++) // ERRO! Incremento não pode estar dentro de expressões
```

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

Exemplo:

```go
fmt.Println(10 > 5)  // true
fmt.Println(10 == 5) // false
fmt.Println(10 != 5) // true
```

📌 **Os operadores de comparação só podem ser usados em valores do mesmo tipo!**

```go
var a int = 10
var b float64 = 10.0

// fmt.Println(a == b) // ERRO! Tipos diferentes
fmt.Println(float64(a) == b) // true (após conversão)
```

---

## **2.3.3 Operadores Lógicos (`&&`, `||`, `!`)**

Os operadores lógicos são usados para combinar expressões booleanas:

| Operador | Descrição | Exemplo |
|----------|------------|---------|
| `&&` | E lógico (AND) | `(x > 0) && (y > 0)` |
| `||` | OU lógico (OR) | `(x > 0) || (y > 0)` |
| `!`  | Negação (NOT)  | `!(x > 0)` |

Exemplo:

```go
a := true
b := false

fmt.Println(a && b) // false
fmt.Println(a || b) // true
fmt.Println(!a)     // false
```

📌 **Avaliação Curta (Short-Circuiting)**  
Go **não avalia a segunda condição se a primeira já for suficiente para determinar o resultado**:

```go
x := 10
if x > 0 || pesado() {  // pesado() NUNCA será chamado
    fmt.Println("Executou!")
}

func pesado() bool {
    fmt.Println("Função pesada executada!")
    return true
}
```

---

## **2.3.4 Operadores de Atribuição**

Os operadores de atribuição permitem modificar valores diretamente:

| Operador | Exemplo | Equivalente a |
|----------|---------|---------------|
| `+=`  | `x += 5`  | `x = x + 5` |
| `-=`  | `x -= 3`  | `x = x - 3` |
| `*=`  | `x *= 2`  | `x = x * 2` |
| `/=`  | `x /= 4`  | `x = x / 4` |
| `%=`  | `x %= 2`  | `x = x % 2` |

Exemplo:

```go
x := 10
x += 5  // x = 15
x *= 2  // x = 30
```

---

## **2.3.5 Precedência e Ordem de Avaliação**

A ordem de precedência dos operadores em Go segue:

1. `* / %`  
2. `+ -`  
3. `== != < > <= >=`  
4. `&&`  
5. `||`  

Exemplo:

```go
resultado := 10 + 5 * 2 // 10 + (5 * 2) = 20
```

📌 **Use parênteses para tornar expressões mais claras**:

```go
resultado := (10 + 5) * 2 // 30
```

---

## **2.3.6 Diferenças em Relação a Outras Linguagens**

| Característica | Go | C / Java |
|---------------|----|---------|
| `++` e `--` dentro de expressões | ❌ Não permitido | ✅ Permitido |
| Conversão de tipos implícita | ❌ Não ocorre | ✅ Ocorre |
| Avaliação curta (Short-Circuit) | ✅ Sim | ✅ Sim |
| Operador ternário (`? :`) | ❌ Não existe | ✅ Existe |
| Sobrecarga de operadores | ❌ Não permitido | ✅ Possível em C++ |

---

## **Conclusão**

Os operadores de Go são simples, mas seguem regras rígidas de tipagem e precedência. No próximo capítulo, exploraremos **entrada e saída de dados com `fmt`**, incluindo formatação avançada! 🚀
