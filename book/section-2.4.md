# **2.4 Entrada e Saída com `fmt`**

O pacote `fmt` é a principal ferramenta de entrada e saída em Go. Ele fornece funções para exibir mensagens na tela e ler entradas do usuário.

---

## **2.4.1 Imprimindo Dados (`fmt.Print`, `fmt.Println`, `fmt.Printf`)**

Go oferece três formas principais de imprimir dados:

### **1. `fmt.Print()`** – Exibe sem quebra de linha

```go
fmt.Print("Olá, ")
fmt.Print("mundo!")
// Saída: Olá, mundo!
```

### **2. `fmt.Println()`** – Adiciona quebra de linha automática

```go
fmt.Println("Olá, mundo!")
fmt.Println("Aprendendo Go!")
// Saída:
// Olá, mundo!
// Aprendendo Go!
```

### **3. `fmt.Printf()`** – Usa placeholders para formatação

```go
nome := "Alice"
idade := 30
fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
// Saída: Nome: Alice, Idade: 30
```

📌 **Principais Placeholders (`%`):**

| Placeholder | Tipo |
|------------|------|
| `%d` | Inteiro |
| `%f` | Float |
| `%s` | String |
| `%t` | Booleano |
| `%v` | Valor genérico |
| `%T` | Tipo da variável |
| `%.2f` | Float com 2 casas decimais |

Exemplo:

```go
preco := 19.99
fmt.Printf("Preço: %.2f\n", preco) // Preço: 19.99
```

---

## **2.4.2 Lendo Entrada do Usuário (`fmt.Scan`, `fmt.Scanln`, `fmt.Scanf`)**

Go permite capturar entrada do usuário pelo teclado.

### **1. `fmt.Scan()`** – Captura múltiplos valores de uma vez

```go
var nome string
var idade int

fmt.Print("Digite seu nome e idade: ")
fmt.Scan(&nome, &idade)

fmt.Println("Nome:", nome, "Idade:", idade)
```

Entrada:

```
Digite seu nome e idade: João 25
```

Saída:

```
Nome: João Idade: 25
```

### **2. `fmt.Scanln()`** – Lê até a quebra de linha

```go
var nome string

fmt.Print("Digite seu nome: ")
fmt.Scanln(&nome)

fmt.Println("Bem-vindo,", nome)
```

### **3. `fmt.Scanf()`** – Entrada formatada

```go
var nome string
var idade int

fmt.Print("Digite seu nome e idade (ex: João 30): ")
fmt.Scanf("%s %d", &nome, &idade)

fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
```

📌 **Diferenças entre `Scan`, `Scanln` e `Scanf`:**

| Função | Como lê entrada |
|--------|---------------|
| `Scan` | Separa valores por espaço |
| `Scanln` | Lê até a quebra de linha |
| `Scanf` | Usa formatação personalizada |

---

## **2.4.3 Lidando com Erros de Entrada**

Caso a entrada seja inválida, podemos verificar erros:

```go
var idade int
fmt.Print("Digite sua idade: ")
_, err := fmt.Scan(&idade)

if err != nil {
    fmt.Println("Erro ao ler idade. Insira um número válido.")
    return
}

fmt.Println("Idade válida:", idade)
```

---

## **2.4.4 Entrada e Saída com Arquivos**

Além do teclado e da tela, `fmt` pode trabalhar com arquivos:

### **Escrevendo em um Arquivo**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    arquivo, err := os.Create("saida.txt")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer arquivo.Close()

    fmt.Fprintln(arquivo, "Texto salvo em arquivo!")
}
```

### **Lendo um Arquivo**

```go
arquivo, err := os.Open("saida.txt")
if err != nil {
    fmt.Println("Erro ao abrir arquivo:", err)
    return
}
defer arquivo.Close()

var texto string
fmt.Fscanln(arquivo, &texto)
fmt.Println("Conteúdo do arquivo:", texto)
```

📌 **Sempre use `defer arquivo.Close()` para garantir que o arquivo seja fechado corretamente.**

---

## **Conclusão**

O pacote `fmt` fornece métodos simples e poderosos para entrada e saída de dados. No próximo capítulo, veremos como realizar **conversões de tipos** em Go! 🚀
