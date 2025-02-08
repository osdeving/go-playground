# **5.3 Strings e Runas (`rune`)**

As **strings** são um dos tipos mais usados em qualquer linguagem de programação, e Go traz algumas peculiaridades importantes na forma como as trata. Além disso, a linguagem possui um tipo especial chamado **`rune`**, que representa caracteres Unicode de maneira mais eficiente.

Nesta seção, exploraremos:

- Como Go trata strings internamente
- Diferença entre `string` e `rune`
- Percorrendo e manipulando strings corretamente
- Como lidar com caracteres Unicode
- Comparação de strings com outras linguagens

---

## **5.3.1 Strings em Go: Conceito e Imutabilidade**

Em Go, **strings são imutáveis**, ou seja, não podem ser modificadas após a criação.

```go
s := "Hello"
s[0] = 'h' // ERRO! Strings são imutáveis.
```

### **Declaração de Strings**

```go
var str1 string = "Go é incrível!"
str2 := "Go suporta Unicode 😊"
```

### **Escape Sequences**

Go suporta caracteres especiais:

```go
s := "Linha 1\nLinha 2"
fmt.Println(s)
```

Saída:

```
Linha 1
Linha 2
```

---

## **5.3.2 Strings e UTF-8: O Que São `rune`?**

Go usa **UTF-8** para armazenar strings. Cada caractere pode ocupar **1 a 4 bytes**.

O tipo `rune` representa **um único caractere Unicode**, armazenado como um número inteiro.

```go
var char rune = 'A'
fmt.Println(char) // 65 (código ASCII de 'A')
```

📌 **Diferente de `byte`, um `rune` pode armazenar caracteres internacionais.**

Exemplo:

```go
char := 'á'
fmt.Println(char)  // 225 (código Unicode de 'á')
```

---

## **5.3.3 Convertendo Strings em `rune` e `byte`**

Podemos converter uma string em `rune` para percorrer corretamente caracteres Unicode:

```go
s := "Golang! 😀"
runes := []rune(s)

fmt.Println(len(s))      // 10 (bytes)
fmt.Println(len(runes))  // 8 (caracteres reais)
```

📌 **Sempre use `[]rune(s)` para contar caracteres corretamente em Unicode!**

---

## **5.3.4 Iterando Sobre Strings**

### **1. Usando `for` Tradicional (Byte a Byte)**

```go
s := "Go言語"

for i := 0; i < len(s); i++ {
    fmt.Printf("Byte %d: %x\n", i, s[i])
}
```

📌 **Isso percorre a string por bytes, podendo cortar caracteres UTF-8.**

### **2. Usando `range` para `rune`**

```go
s := "Go言語"

for i, r := range s {
    fmt.Printf("Posição: %d, Rune: %c\n", i, r)
}
```

Saída:

```
Posição: 0, Rune: G
Posição: 1, Rune: o
Posição: 2, Rune: 言
Posição: 5, Rune: 語
```

📌 **O índice pode pular valores devido à codificação UTF-8!**

---

## **5.3.5 Manipulação de Strings**

### **Concatenando Strings**

A concatenação pode ser feita com `+`:

```go
s1 := "Go"
s2 := "Lang"
s3 := s1 + s2

fmt.Println(s3) // "GoLang"
```

📌 **Evite concatenar muitas strings com `+`, pois isso cria várias cópias na memória.**  
✅ **Prefira `strings.Builder` para eficiência:**

```go
var sb strings.Builder
sb.WriteString("Go")
sb.WriteString("Lang")

fmt.Println(sb.String()) // "GoLang"
```

---

## **5.3.6 Comparação de Strings**

Em Go, strings podem ser comparadas diretamente:

```go
fmt.Println("golang" == "golang") // true
fmt.Println("go" < "golang")      // true (ordem lexicográfica)
```

📌 **A comparação segue a ordem Unicode dos caracteres.**

---

## **5.3.7 Substrings em Go**

Go permite fatiar strings usando índices:

```go
s := "Golang"
sub := s[0:3] // "Gol"
```

📌 **Isso retorna um slice de `byte`, não de `rune`!**  
✅ **Para Unicode, converta para `rune`:**

```go
runes := []rune("Go言語")
sub := string(runes[0:2]) // "Go"
```

---

## **5.3.8 Principais Funções do Pacote `strings`**

| Função | Descrição |
|--------|-----------|
| `strings.Contains(s, "Go")` | Verifica se a string contém um valor |
| `strings.ToUpper(s)` | Converte para maiúsculas |
| `strings.ToLower(s)` | Converte para minúsculas |
| `strings.Replace(s, "Go", "Rust", -1)` | Substitui substrings |
| `strings.Split(s, ",")` | Divide uma string por um separador |
| `strings.TrimSpace(s)` | Remove espaços extras |

Exemplo:

```go
s := " Golang "
fmt.Println(strings.TrimSpace(s)) // "Golang"
```

📌 **Essas funções facilitam a manipulação de strings sem criar loops manuais.**

---

## **5.3.9 Comparação com Outras Linguagens**

| Característica       | Go  | C  | Java  | Python |
|----------------------|----|----|-------|--------|
| Strings Imutáveis   | ✅  | ❌ | ✅     | ✅      |
| Suporte UTF-8       | ✅  | ❌ | ✅     | ✅      |
| Rune (Unicode Char) | ✅  | ❌ | ❌     | ✅ (`ord()`) |
| Concatenar com `+`  | ✅  | ✅ | ❌ (`StringBuilder`) | ✅      |
| Contar Caracteres   | ❌ (`len(s) conta bytes`) | ❌ | ✅ | ✅      |

📌 **Go trata strings de forma eficiente e integrada com Unicode, sem precisar de bibliotecas externas.**

---

## **Conclusão**

As strings em Go são eficientes e bem integradas com UTF-8. O uso correto de `rune` e `strings.Builder` pode melhorar a manipulação e evitar alocações desnecessárias.

No próximo capítulo, exploraremos **strings imutáveis e manipulação avançada com `bytes`!** 🚀
