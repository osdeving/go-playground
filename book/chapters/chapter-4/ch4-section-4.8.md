# **4.8 Funções Comuns e Builtins**

Go fornece várias **funções embutidas (built-in functions)** que ajudam em operações do dia a dia, como manipulação de strings, conversão de tipos, cálculos matemáticos e criação de estruturas de dados. Algumas dessas funções são fundamentais e vale a pena **memorizá-las**.

Nesta seção, abordaremos:

- As funções built-in mais usadas em Go
- Implementação simplificada de algumas dessas funções
- Uso de closures para recriar comportamentos comuns
- Aplicações práticas das funções embutidas

---

## **4.8.1 Principais Funções Built-in**

Go possui um conjunto de funções **sempre disponíveis**, sem necessidade de importar pacotes:

| Função   | Descrição |
|----------|-----------|
| `len()`  | Retorna o tamanho de arrays, slices, maps ou strings |
| `cap()`  | Retorna a capacidade de um slice |
| `append()` | Adiciona elementos a um slice |
| `copy()` | Copia elementos entre slices |
| `make()` | Cria slices, maps e channels |
| `new()`  | Aloca memória para um tipo |
| `delete()` | Remove elementos de um map |
| `close()` | Fecha um canal |
| `panic()` | Gera um erro fatal |
| `recover()` | Captura um `panic` |

---

## **4.8.2 Implementando `len()` Simplificado**

A função `len()` retorna o tamanho de um slice ou string. Podemos recriar essa funcionalidade:

```go
func length[T any](s []T) int {
    count := 0
    for range s {
        count++
    }
    return count
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(length(nums)) // 5
}
```

📌 **Go otimiza `len()` internamente, mas essa implementação mostra a lógica por trás.**

---

## **4.8.3 Criando um `append()` Personalizado**

A função `append()` adiciona elementos a um slice e retorna um novo slice:

```go
func appendCustom[T any](s []T, elements ...T) []T {
    return append(s, elements...)
}

func main() {
    nums := []int{1, 2, 3}
    nums = appendCustom(nums, 4, 5)
    fmt.Println(nums) // [1, 2, 3, 4, 5]
}
```

📌 **`append()` realoca o slice se necessário, garantindo espaço para os novos elementos.**

---

## **4.8.4 Funções Built-in com Closures**

Closures podem ser usados para criar funções utilitárias dinâmicas.

### **Criando um `filter()` para slices**

Go não tem `filter()` nativo como Python, mas podemos criá-lo:

```go
func filter[T any](s []T, test func(T) bool) []T {
    result := []T{}
    for _, v := range s {
        if test(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    even := filter(nums, func(n int) bool { return n%2 == 0 })
    fmt.Println(even) // [2, 4]
}
```

📌 **Essa técnica simula a função `filter()` de outras linguagens.**

---

## **4.8.5 Recriando `map()` para Transformação de Slices**

Outra função útil que podemos implementar com closures:

```go
func mapSlice[T any, U any](s []T, transform func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = transform(v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    squared := mapSlice(nums, func(n int) int { return n * n })
    fmt.Println(squared) // [1, 4, 9, 16, 25]
}
```

📌 **`map()` permite transformar todos os elementos de um slice sem criar loops explícitos.**

---

## **4.8.6 Criando um `reduce()`**

A função `reduce()` acumula valores de um slice:

```go
func reduce[T any](s []T, accumulator func(T, T) T, initial T) T {
    result := initial
    for _, v := range s {
        result = accumulator(result, v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    sum := reduce(nums, func(a, b int) int { return a + b }, 0)
    fmt.Println(sum) // 15
}
```

📌 **Isso simula `reduce()` do JavaScript e Python, útil para agregações.**

---

## **4.8.7 Trabalhando com `strings`**

Além das funções embutidas, o pacote `strings` oferece várias utilidades. Podemos recriar algumas:

### **Recriando `strings.ToUpper()`**

```go
func toUpper(s string) string {
    result := []rune(s)
    for i, char := range result {
        if char >= 'a' && char <= 'z' {
            result[i] = char - 32
        }
    }
    return string(result)
}

func main() {
    fmt.Println(toUpper("hello")) // "HELLO"
}
```

📌 **Essa versão converte caracteres manualmente sem usar a função nativa.**

---

## **4.8.8 Comparação com Outras Linguagens**

| Função | Go | Python | JavaScript |
|--------|----|--------|------------|
| `len()` | ✅ | ✅ (`len()`) | ✅ (`.length`) |
| `append()` | ✅ | ✅ (`.append()`) | ✅ (`push()`) |
| `map()` | ❌ (precisa de implementação) | ✅ | ✅ |
| `filter()` | ❌ (precisa de implementação) | ✅ | ✅ |
| `reduce()` | ❌ (precisa de implementação) | ✅ | ✅ |

📌 **Go não tem `map()`, `filter()` e `reduce()` nativos para slices, mas podemos implementá-los.**

---

## **Conclusão**

As funções built-in de Go são otimizadas para eficiência, mas podemos **recriá-las** para entender sua lógica e expandir a funcionalidade da linguagem.

No próximo capítulo, abordaremos **estruturas de dados e manipulação de memória**, explorando como Go gerencia slices, maps e alocações de forma eficiente! 🚀
