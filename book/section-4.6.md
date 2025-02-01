# **4.6 Recursão**

A **recursão** é uma técnica na qual uma função **chama a si mesma** para resolver um problema, geralmente dividindo-o em partes menores e resolvendo cada uma de forma independente. Em Go, a recursão é suportada nativamente e pode ser usada para **resolver problemas de maneira declarativa**.

Nesta seção, abordaremos:

- Como funciona a recursão em Go
- Casos clássicos de uso da recursão
- Diferenças entre recursão e laços (`for`)
- Problemas comuns e otimizações

---

## **4.6.1 O Que é Recursão?**

Uma função recursiva chama a si mesma para resolver um problema:

```go
func countdown(n int) {
    if n <= 0 {
        fmt.Println("Fim!")
        return
    }
    fmt.Println(n)
    countdown(n - 1) // Chamada recursiva
}

func main() {
    countdown(5)
}
```

Saída:

```
5
4
3
2
1
Fim!
```

📌 **Cada chamada empilha um novo frame na stack, exigindo um caso base (`if`) para evitar loops infinitos.**

---

## **4.6.2 Casos Clássicos de Recursão**

### **1. Fatorial (`n!`)**

O cálculo do fatorial pode ser definido recursivamente:

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 120
}
```

📌 **Fatorial cresce rapidamente, podendo causar estouro de stack (`stack overflow`).**

### **2. Sequência de Fibonacci**

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println(fibonacci(10)) // 55
}
```

📌 **Essa versão é ineficiente (O(2^n)), pois recalcula valores repetidos.**  
✅ **Otimização:** Usar **memoization** ou uma abordagem iterativa.

---

## **4.6.3 Recursão vs. Laços (`for`)**

| Método    | Vantagens | Desvantagens |
|-----------|----------|--------------|
| **Recursão** | Código mais legível para problemas naturalmente recursivos | Pode causar estouro de stack |
| **Iteração (`for`)** | Melhor eficiência de memória e desempenho | Pode ser mais difícil de entender |

✅ **Use recursão para problemas naturalmente recursivos, como árvores e grafos.**  
✅ **Use `for` quando possível para evitar uso excessivo de memória.**

---

## **4.6.4 Recursão em Estruturas de Dados**

### **Exemplo: Percorrendo uma Árvore**

```go
type Node struct {
    Value int
    Left  *Node
    Right *Node
}

func traverse(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.Value)
    traverse(node.Left)
    traverse(node.Right)
}

func main() {
    root := &Node{10, &Node{5, nil, nil}, &Node{20, nil, nil}}
    traverse(root)
}
```

📌 **Árvores são um caso ideal para recursão devido à sua estrutura hierárquica.**

---

## **4.6.5 Problemas Comuns e Otimizações**

❌ **Estouro de Stack (`stack overflow`)**  
✅ **Use `tail recursion` (Go não otimiza isso nativamente)**  
✅ **Transforme em iteração se possível**  

❌ **Desempenho ruim em Fibonacci**  
✅ **Use memoization para evitar recomputações**  

```go
var memo = make(map[int]int)

func fibonacciOptimized(n int) int {
    if n <= 1 {
        return n
    }
    if val, exists := memo[n]; exists {
        return val
    }
    memo[n] = fibonacciOptimized(n-1) + fibonacciOptimized(n-2)
    return memo[n]
}
```

📌 **Agora `fibonacci(50)` roda rapidamente sem recomputações.**

---

## **Conclusão**

A recursão em Go é **poderosa e expressiva**, mas deve ser usada com cuidado para evitar problemas de desempenho e stack overflow. No próximo capítulo, exploraremos **ponteiros e funções**, abordando como evitar cópias desnecessárias de dados! 🚀
