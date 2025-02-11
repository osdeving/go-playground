# **4.1 Declaração e Uso de Funções**

<div style="text-align: right; border-left: 4px solid #ccc; padding-left: 10px; font-style: italic;">
    <strong>❝ Está funcionando? Nem rela! ❞</strong> <br> Provérbio Chinês <br><br>
</div>


Funções são blocos fundamentais para **organização, reutilização e abstração de código**. Em Go, funções são **primeira classe**, o que significa que podem ser atribuídas a variáveis, passadas como argumentos e retornadas de outras funções.

Nesta seção, exploraremos:

- A sintaxe básica de funções
- Diferenças entre funções em Go e outras linguagens
- Melhores práticas para eficiência e organização do código
- Exemplos realistas de uso

---

## **4.1.1 Estrutura de uma Função em Go**

Uma função em Go segue a estrutura:

```go
func functionName(parameters) returnType {
    // Corpo da função
    return value
}
```

Exemplo básico:

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(10, 20)
    fmt.Println("Sum:", sum) // Sum: 30
}
```

📌 **Regras importantes sobre funções em Go**:

1. **Os tipos dos parâmetros devem ser explicitamente declarados.**  
   - Exceção: Se múltiplos parâmetros forem do mesmo tipo, podemos omitir o tipo dos primeiros.

   ```go
   func multiply(a, b int) int { // Correto
       return a * b
   }
   ```

2. **O tipo de retorno deve ser declarado.**  
   - Se a função não retorna nada, omitimos o tipo (`func doSomething()`).

3. **O retorno deve ser explícito (`return`), exceto para funções que não retornam nada (`func doSomething()`) ou funções que retornam um erro (`func doSomething() error`).**

---

## **4.1.2 Funções sem Retorno**

Funções podem ser usadas apenas para executar ações sem efetivamente retornar valores:

```go
func logMessage(message string) {
    fmt.Println("Log:", message)
}
```

📌 **Go não usa a palavra `void`. Funções sem retorno simplesmente não declaram um tipo de retorno.**

---

## **4.1.3 Chamando Funções e Passando Argumentos**

### **Passagem por Valor**

Por padrão, **Go passa os argumentos por valor**, ou seja, uma cópia do valor é enviada para a função:

```go
func double(x int) {
    x = x * 2 // Isso NÃO altera o valor original
}

func main() {
    num := 10
    double(num)
    fmt.Println(num) // Ainda é 10
}
```

Para modificar o valor original, devemos passar um **ponteiro**.

### **Passagem por Referência usando Ponteiros**

```go
func doublePointer(x *int) {
    *x = *x * 2 // Agora alteramos diretamente o valor
}

func main() {
    num := 10
    doublePointer(&num)
    fmt.Println(num) // Agora é 20
}
```

Veremos mais sobre ponteiros nas próximas seções e no capítulo 7 vamos estudar ponteiros em detalhes.

---

## **4.1.4 Retornando Múltiplos Valores**

Go permite que uma função retorne múltiplos valores, evitando a necessidade de criar estruturas auxiliares:

```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}
```

📌 **Com isso podemos retornar erros sem precisar de exceções (esse recurso será explicado melhor na seção 4.2).**

Veja um exemplo de uma função que tenta buscar um usuário e retorna um erro caso não exista:

```go
func findUser(id int) (string, error) {
    if id == 42 {
        return "John Doe", nil
    }
    return "", fmt.Errorf("User not found")
}

func main() {
    user, err := findUser(10)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("User:", user)
}
```

---

## **4.1.5 Funções são Objetos de Primeira Classe (Higher-Order Functions)**

Objetos de primeira classe são objetos que podem ser atribuídos a variáveis, passados como argumentos e retornados de outras funções. Uma vez que, em Go, funções são objetos de primeira classe, elas podem ser **atribuídas a variáveis**, **passadas como argumentos** e **retornadas de outras funções**, o que facilita **programação funcional**.

### **Passando Funções como Parâmetro**

No exemplo abaixo, a função `applyOperation` recebe uma função como argumento e a aplica aos argumentos `a` e `b`:

```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    result := applyOperation(10, 5, add)

    fmt.Println("Result:", result) // Result: 15
}
```

### **Retornando uma Função**

No exemplo abaixo, a função `multiplier` retorna uma função que multiplica seu argumento por um fator:

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    fmt.Println(double(5)) // 10
}
```

📌 **Isso possibilita gerar funções dinâmicas com diferentes comportamentos. P.ex.: `triple := multiplier(3)`**

---

## **4.1.6 Funções Inline e Uso de `func()`**

Go permite a criação de **funções anônimas** que podem ser usadas diretamente dentro de blocos de código:

```go
result := func(a, b int) int {
    return a + b
}(3, 4)

fmt.Println(result) // 7
```

📌 **Útil para executar lógicas simples sem precisar nomear uma função.**

---

## **4.1.7 Comparação com Outras Linguagens**

| Conceito            | Go | C | JavaScript | Python |
|---------------------|----|---|------------|--------|
| Funções nomeadas    | ✅  | ✅ | ✅          | ✅      |
| Retorno múltiplo    | ✅  | ❌ | ✅ (array) | ✅      |
| Ponteiros           | ✅  | ✅ | ❌         | ❌      |
| Funções anônimas    | ✅  | ❌ | ✅ (arrow) | ✅      |
| Passagem por valor  | ✅  | ✅ | ❌ (obj ref) | ✅      |

📌 **Diferente de C e Java, Go tem suporte nativo para múltiplos retornos e funções anônimas.**

---

## **Conclusão**

Funções em Go são **poderosas e flexíveis**, suportando:
- **Passagem de argumentos por valor e referência**
- **Retorno de múltiplos valores**
- **Funções como primeira classe**
- **Uso de funções anônimas e closures**

No próximo capítulo, exploraremos **parâmetros e retornos**, abordando técnicas avançadas para manipulação de valores em funções. 🚀
