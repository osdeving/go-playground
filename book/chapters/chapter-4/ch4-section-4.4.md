# **4.4 Funções Variádicas**

Funções variádicas permitem passar um **número variável de argumentos** para uma função. Esse recurso é útil quando não sabemos de antemão quantos valores serão fornecidos. Em Go, funções variádicas são implementadas usando **`...` (ellipsis notation)**.

Nesta seção, exploraremos:

- Como declarar e usar funções variádicas
- Como manipular os argumentos dentro da função
- O uso de `variadic` e `non-variadic` parameters juntos
- Eficiência e melhores práticas

---

## **4.4.1 Definição de Funções Variádicas**

A sintaxe para criar uma função variádica em Go é:

```go
func functionName(param ...tipo) retorno {}
```

### **Exemplo Simples**

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))       // 6
    fmt.Println(sum(10, 20, 30, 40)) // 100
}
```

📌 **O parâmetro `numbers` é tratado como um `slice` dentro da função.**

---

## **4.4.2 Misturando Parâmetros Normais e Variádicos**

Podemos combinar **parâmetros fixos** com **parâmetros variádicos**, desde que o variádico seja o último:

```go
func printNames(prefix string, names ...string) {
    for _, name := range names {
        fmt.Println(prefix, name)
    }
}

func main() {
    printNames("Hello,", "Alice", "Bob", "Charlie")
}
```

📌 **O primeiro parâmetro (`prefix`) é obrigatório, os demais são opcionais.**

---

## **4.4.3 Passando Slices como Argumentos Variádicos**

Como funções variádicas esperam um **slice**, podemos passar um **slice existente** usando `...`:

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    valores := []int{1, 2, 3, 4}
    fmt.Println(sum(valores...)) // Passa um slice para uma função variádica
}
```

📌 **Sem `...`, Go tratará o slice como um único argumento inválido.**

---

## **4.4.4 Funções Variádicas com Diferentes Tipos**

Se precisarmos de múltiplos tipos, podemos usar `interface{}`:

```go
func logValues(values ...interface{}) {
    for _, v := range values {
        fmt.Println(v)
    }
}

func main() {
    logValues(1, "Hello", true, 3.14)
}
```

📌 **Isso é útil para logs genéricos, mas evita tipagem forte.**

---

## **4.4.5 Eficiência e Melhor Práticas**

✔ **Evite o uso excessivo de `interface{}`**: reduz a segurança de tipos.  
✔ **Prefira slices quando possível**: evita a necessidade de conversão.  
✔ **Evite grandes alocações em funções variádicas**: cada chamada cria um novo slice.  

---

## **Conclusão**

Funções variádicas tornam o código mais flexível, permitindo lidar com um número dinâmico de argumentos. No próximo capítulo, exploraremos **funções anônimas e closures**! 🚀
