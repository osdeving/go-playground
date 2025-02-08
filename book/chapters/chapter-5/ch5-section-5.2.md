# **5.2 Slices: Conceito, Capacidade e Expansão**

Os **slices** são a principal estrutura de dados para armazenar sequências dinâmicas em Go. Diferente dos arrays, que possuem **tamanho fixo**, os slices podem crescer e mudar de tamanho sem precisar de uma nova alocação manual.

Nesta seção, exploraremos:

- O conceito e a estrutura interna dos slices
- Como declarar, inicializar e modificar slices
- Capacidade (`cap`) e crescimento dinâmico
- Como o Go gerencia memória para slices
- Comparação de desempenho com arrays

---

## **5.2.1 O Que São Slices?**

Um **slice** é uma abstração sobre arrays, oferecendo **tamanho dinâmico** e operações convenientes:

```go
var s []int // Declara um slice de inteiros (sem tamanho fixo)
```

📌 **Diferente de arrays, slices não têm um tamanho fixo na declaração.**

Podemos inicializá-los diretamente:

```go
numeros := []int{10, 20, 30} // Slice já inicializado
fmt.Println(numeros) // [10, 20, 30]
```

---

## **5.2.2 Criando Slices com `make()`**

Go permite criar slices usando a função `make()`, que aloca memória dinamicamente:

```go
s := make([]int, 5) // Slice de 5 elementos inicializados com 0
fmt.Println(s) // [0 0 0 0 0]
```

📌 **A função `make()` é útil quando queremos criar um slice com tamanho inicial, mas sem valores predefinidos.**

Podemos especificar **capacidade extra**:

```go
s := make([]int, 3, 5) // Tamanho 3, capacidade 5
fmt.Println(len(s), cap(s)) // 3 5
```

📌 **A capacidade extra permite adicionar elementos sem realocar memória.**

---

## **5.2.3 Acessando e Modificando Slices**

Os elementos são acessados da mesma forma que em arrays:

```go
s := []string{"Go", "Python", "Rust"}
fmt.Println(s[0]) // "Go"

s[1] = "JavaScript"
fmt.Println(s) // ["Go", "JavaScript", "Rust"]
```

📌 **Diferente de arrays, slices podem ser redimensionados dinamicamente.**

---

## **5.2.4 Capacidade (`cap`) e Expansão de Slices**

Todo slice possui:

- **Comprimento (`len`)** → Número de elementos armazenados.
- **Capacidade (`cap`)** → Número máximo de elementos antes da realocação.

```go
s := make([]int, 3, 5)
fmt.Println(len(s), cap(s)) // 3 5
```

Se adicionarmos elementos além da capacidade, o Go cria **automaticamente** um novo array maior:

```go
s = append(s, 10, 20, 30)
fmt.Println(len(s), cap(s)) // 6 10 (nova alocação)
```

📌 **Go dobra a capacidade dos slices automaticamente quando eles crescem.**

---

## **5.2.5 Sub-slices e Compartilhamento de Memória**

Podemos criar **sub-slices** de um slice original:

```go
original := []int{1, 2, 3, 4, 5}
sub := original[1:4] // [2, 3, 4]
```

📌 **O sub-slice compartilha a memória com o original!**

```go
sub[0] = 100
fmt.Println(original) // [1, 100, 3, 4, 5] (o original foi alterado)
```

Se quisermos evitar modificações no slice original, podemos copiar os dados:

```go
copia := append([]int{}, sub...)
```

📌 **Use `append([]T{}, slice...)` para criar uma cópia independente.**

---

## **5.2.6 Comparação de Desempenho: Arrays vs. Slices**

Os slices são geralmente mais eficientes do que arrays fixos porque permitem redimensionamento dinâmico sem realocar manualmente memória.

| Característica      | Arrays  | Slices |
|--------------------|--------|--------|
| Tamanho fixo      | ✅ Sim | ❌ Não |
| Redimensionável   | ❌ Não | ✅ Sim |
| Compartilhamento de Memória | ❌ Não | ✅ Sim |
| Uso mais comum    | ❌ Limitado | ✅ Sim |

📌 **Na prática, slices são usados na maioria dos casos.**

---

## **5.2.7 Melhores Práticas com Slices**

✔ **Use `make()` quando souber o tamanho inicial para evitar realocações desnecessárias.**  
✔ **Evite modificar slices derivados (`s[1:3]`), pois isso pode afetar o original.**  
✔ **Use `append()` de forma inteligente para evitar muitas realocações de memória.**  
✔ **Para copiar slices, use `append([]T{}, slice...)` ou `copy()`.**  

---

## **Conclusão**

Os slices são a estrutura de dados mais flexível e eficiente para armazenar listas dinâmicas em Go. No próximo capítulo, exploraremos **strings e runas (`rune`)**, essenciais para manipulação de texto em Go! 🚀
