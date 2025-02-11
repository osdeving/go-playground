# **4.2 Parâmetros e Retornos**

Os parâmetros e os retornos de funções são componentes essenciais em Go, permitindo que funções recebam dados, os processem e retornem resultados. Diferente de algumas linguagens, Go possui algumas características específicas, como **tipagem explícita, múltiplos retornos e retorno nomeado**.

Nesta seção, exploraremos:

- Como declarar e usar parâmetros
- Tipagem explícita e inferência de tipos
- Passagem de parâmetros por valor e por referência
- Múltiplos retornos e como tratá-los
- Boas práticas e otimizações

---

## **4.2.1 Parâmetros em Funções**

Os parâmetros são declarados dentro dos parênteses após o nome da função:

```go
func add(a int, b int) int {
    return a + b
}
```

📌 **Se vários parâmetros forem do mesmo tipo, podemos omitir os tipos intermediários:**

```go
func multiply(a, b int) int { // Mais compacto
    return a * b
}
```

### **Parâmetros Opcionais? Não em Go!**

Diferente de Python e JavaScript, **Go não suporta parâmetros opcionais ou valores padrão**. Alternativas incluem:

- **Usar múltiplas versões da função (overloading não existe em Go).**
- **Passar uma struct contendo os parâmetros.**
- **Utilizar `variadic functions` (ver seção 4.4).**

---

## **4.2.2 Passagem de Parâmetros por Valor e Referência**

Por padrão, **Go passa parâmetros por valor**, criando uma cópia da variável:

```go
func double(x int) {
    x = x * 2 // Modifica apenas a cópia
}

func main() {
    num := 10
    double(num)
    fmt.Println(num) // Ainda é 10
}
```

### **Passagem por Referência com Ponteiros**

Para modificar o valor original, passamos um **ponteiro**:

```go
func doublePointer(x *int) {
    *x = *x * 2 // Modifica o valor original
}

func main() {
    num := 10
    doublePointer(&num)
    fmt.Println(num) // Agora é 20
}
```

📌 **Quando usar passagem por referência?**

- Quando precisar modificar a variável original.
- Para evitar cópias desnecessárias de grandes estruturas (como structs e slices).

---

## **4.2.3 Retorno de Valores**

O tipo de retorno de uma função é declarado após os parâmetros:

```go
func square(x int) int {
    return x * x
}
```

📌 **O retorno deve ser explícito. Não há `implicit return` como em Python.**

### **Funções sem Retorno (`void` em Go)**

```go
func logMessage(msg string) {
    fmt.Println("Log:", msg)
}
```

📌 **Go não usa a palavra `void`. Funções sem retorno simplesmente não declaram um tipo de retorno.**

---

## **4.2.4 Retornando Múltiplos Valores**

Diferente de Java e C, Go suporta **múltiplos retornos** nativamente sem necessidade de structs auxiliares:

```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}
```

📌 **Essa funcionalidade é muito usada para tratamento de erros!**

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

### **Ignorando Retornos**

Caso não precisemos de um valor retornado, usamos `_`:

```go
_, remainder := divide(10, 3)
fmt.Println("Remainder:", remainder)
```

📌 **Isso evita erros do compilador sobre variáveis não usadas.**

---

## **4.2.5 Retornos Nomeados**

Go permite **nomes explícitos para valores de retorno**, tornando o código mais legível:

```go
func userInfo(id int) (name string, age int) {
    if id == 1 {
        name, age = "Alice", 30
    } else {
        name, age = "Unknown", 0
    }
    return // Retorno implícito das variáveis nomeadas
}
```

📌 **Use retornos nomeados com moderação, pois podem reduzir a clareza do código!**

---

## **4.2.6 Tratamento de Erros com Retorno Múltiplo**

Diferente de outras linguagens, **Go não possui exceções (`try/catch`)**, mas sim um padrão de erro explícito:

```go
func openFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func main() {
    file, err := openFile("data.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close() // Garante que o arquivo seja fechado
    fmt.Println("File opened successfully")
}
```

📌 **Esse padrão melhora a previsibilidade e controle sobre erros.**

---

## **4.2.7 Comparação com Outras Linguagens**

| Conceito             | Go  | C   | Java  | Python |
|----------------------|----|-----|-------|--------|
| Passagem por valor   | ✅  | ✅  | ✅    | ✅      |
| Passagem por referência | ✅ (com ponteiros) | ✅ | ✅ (objetos) | ✅ (imutável por padrão) |
| Múltiplos retornos   | ✅  | ❌  | ❌    | ✅      |
| Retorno nomeado      | ✅  | ❌  | ❌    | ❌      |
| Tratamento de erro por retorno | ✅  | ❌  | ❌ (exceptions) | ❌ (exceptions) |

📌 **Go evita exceções e prioriza um fluxo de código mais previsível.**

---

## **Conclusão**

Os parâmetros e retornos em Go foram projetados para **clareza e eficiência**, evitando implicitamente muitos dos problemas de outras linguagens. Os principais pontos são:

- **Passagem de valores por padrão, ponteiros para modificações diretas.**
- **Suporte nativo a múltiplos retornos.**
- **Padrão explícito para manipulação de erros.**
- **Retornos nomeados para melhor legibilidade.**

No próximo capítulo, abordaremos **retornos nomeados**, explorando quando e como usá-los para tornar o código mais expressivo. 🚀
