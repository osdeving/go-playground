# **5.4 Strings Imutáveis e Manipulação com `strings` e `bytes`**

Em Go, as **strings são imutáveis**, ou seja, não podem ser alteradas diretamente após a criação. Essa característica pode gerar desafios ao manipular grandes volumes de texto, exigindo abordagens mais eficientes para otimizar a performance.

Nesta seção, exploraremos:

- Por que strings são imutáveis em Go
- Alternativas eficientes para modificar strings
- Uso do pacote `strings` para manipulação avançada
- Como `bytes.Buffer` e `strings.Builder` evitam alocações excessivas
- Casos de uso e boas práticas

---

## **5.4.1 Por Que Strings São Imutáveis?**

Strings em Go são representadas internamente como **slices de bytes (`[]byte`)**:

```go
type string struct {
    array *byte // Ponteiro para os dados
    len   int   // Tamanho da string
}
```

Essa estrutura torna as strings **rápidas para comparação e seguras para concorrência**, mas **ineficientes para modificações frequentes**.

📌 **Qualquer alteração em uma string cria uma nova cópia na memória!**

```go
s := "Go"
s = s + "lang" // Uma nova string é alocada!
```

💡 **Se precisar modificar uma string frequentemente, use `bytes.Buffer` ou `strings.Builder`.**

---

## **5.4.2 Convertendo Strings em `[]byte` e `[]rune`**

Podemos converter uma string para um slice de bytes ou runas para modificá-la:

```go
s := "Golang"

b := []byte(s) // Converte para `[]byte`
b[0] = 'J'     // Modifica o primeiro caractere
s = string(b)  // Converte de volta para string

fmt.Println(s) // "Jolang"
```

📌 **Isso funciona, mas é ineficiente para modificações frequentes.**  
✅ **Prefira `strings.Builder` para concatenações complexas.**

---

## **5.4.3 Uso do Pacote `strings`**

O pacote `strings` oferece funções para manipulação eficiente de strings:

| Função | Descrição |
|--------|-----------|
| `strings.Contains(s, "Go")` | Verifica se a string contém um valor |
| `strings.HasPrefix(s, "Go")` | Verifica se a string começa com um valor |
| `strings.HasSuffix(s, "Lang")` | Verifica se a string termina com um valor |
| `strings.Split(s, ",")` | Divide uma string em um slice |
| `strings.Replace(s, "Go", "Rust", -1)` | Substitui substrings |
| `strings.TrimSpace(s)` | Remove espaços extras |

Exemplo:

```go
s := "  Go é incrível!  "
fmt.Println(strings.TrimSpace(s)) // "Go é incrível!"
```

📌 **Isso evita manipulação manual de índices e alocações desnecessárias.**

---

## **5.4.4 Concatenando Strings de Forma Eficiente**

A concatenação com `+` pode ser custosa, pois cria uma nova string a cada operação:

```go
s := "Go"
s += "lang" // Cria uma nova string na memória!
```

✅ **Use `strings.Builder` para evitar alocações excessivas:**

```go
var sb strings.Builder
sb.WriteString("Go")
sb.WriteString("lang")

fmt.Println(sb.String()) // "Golang"
```

📌 **`strings.Builder` é a forma mais eficiente de construir strings dinamicamente.**

---

## **5.4.5 Manipulação Avançada com `bytes.Buffer`**

Para modificar grandes quantidades de texto, `bytes.Buffer` pode ser ainda mais eficiente:

```go
var buffer bytes.Buffer

buffer.WriteString("Olá, ")
buffer.WriteString("mundo!")

fmt.Println(buffer.String()) // "Olá, mundo!"
```

📌 **`bytes.Buffer` é útil para grandes strings ou manipulação frequente de texto.**

---

## **5.4.6 Strings vs. `[]byte`: Comparação de Performance**

| Operação               | String (`+`) | `strings.Builder` | `bytes.Buffer` |
|------------------------|-------------|------------------|---------------|
| Concatenar Pequenas   | ✅ Rápido   | ✅ Rápido        | ❌ Desnecessário |
| Concatenar Muitas     | ❌ Ineficiente | ✅ Melhor opção  | ✅ Melhor opção |
| Substituir Textos     | ❌ Ineficiente | ✅ Melhor opção  | ✅ Melhor opção |
| Modificar Caracteres  | ❌ Impossível | ❌ Não aplicável | ✅ Melhor opção |

📌 **Use `strings.Builder` para concatenação e `bytes.Buffer` para manipulação extensa.**

---

## **5.4.7 Comparação com Outras Linguagens**

| Característica       | Go  | C  | Java  | Python |
|----------------------|----|----|-------|--------|
| Strings Imutáveis   | ✅  | ❌ | ✅     | ✅      |
| `StringBuilder`     | ✅  | ❌ | ✅     | ✅ (`join()`) |
| Modificar Strings   | ❌  | ✅ | ❌     | ✅      |
| Suporte UTF-8       | ✅  | ❌ | ✅     | ✅      |

📌 **Go otimiza strings para concorrência e eficiência, evitando modificações diretas.**

---

## **Conclusão**

Go lida com strings de forma segura e eficiente, mas modificá-las requer abordagens otimizadas.  
**Prefira `strings.Builder` e `bytes.Buffer` para manipulação frequente de texto.**

No próximo capítulo, exploraremos **Deep Copy vs. Shallow Copy**, abordando como Go lida com cópias de estruturas de dados! 🚀
