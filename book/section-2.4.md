# **2.4 Entrada e Saída com `fmt`**

>🗨️ "Uma boa comunicação começa com uma boa formatação. Em Go, a biblioteca fmt lhe dá controle total sobre como os dados são exibidos e lidos. Entender suas nuances fará de você um programador mais eficiente e expressivo." — Filosofia Go

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

## **Pratique Go**

🎯 Agora que você aprendeu sobre entrada e saída com `fmt`, tente os seguintes desafios:

🧐 **Desafios**:

<details>
  <summary>1️⃣ Escreva um programa que leia um nome e exiba uma saudação personalizada.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var nome string
      fmt.Print("Digite seu nome: ")
      fmt.Scanln(&nome)
      fmt.Printf("Olá, %s! Seja bem-vindo.\n", nome)
  }
  ```
  
</details>

<details>
  <summary>2️⃣ Leia dois números do usuário e exiba a soma, subtração, multiplicação e divisão.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var a, b float64
      fmt.Print("Digite dois números: ")
      fmt.Scan(&a, &b)
      fmt.Printf("Soma: %.2f\nSubtração: %.2f\nMultiplicação: %.2f\nDivisão: %.2f\n", a+b, a-b, a*b, a/b)
  }
  ```
  
</details>

<details>
  <summary>3️⃣ Formate um número `float64` para exibir apenas duas casas decimais ao imprimir.</summary>
  
  ```go
  var num float64 = 3.141592
  fmt.Printf("%.2f\n", num)
  ```
  
</details>

<details>
  <summary>4️⃣ Utilize `fmt.Scanf` para capturar múltiplos valores em uma única linha.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var nome string
      var idade int
      fmt.Print("Digite seu nome e idade: ")
      fmt.Scanf("%s %d", &nome, &idade)
      fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
  }
  ```
  
</details>

---

## **Perguntas e Respostas**

🎡 **Teste seus conhecimentos:**

<details>
  <summary>1️⃣ Qual a diferença entre `fmt.Print`, `fmt.Println` e `fmt.Printf`?</summary>
  `fmt.Print` imprime sem adicionar nova linha, `fmt.Println` adiciona uma nova linha no final, e `fmt.Printf` permite formatação avançada.
</details>

<details>
  <summary>2️⃣ Como capturar a entrada do usuário usando `fmt.Scan`?</summary>
  `fmt.Scan` lê valores separados por espaço e os armazena nas variáveis passadas como ponteiros.
</details>

<details>
  <summary>3️⃣ Qual o formato correto para exibir um número decimal, hexadecimal e binário usando `fmt.Printf`?</summary>
  `%d` para decimal, `%x` para hexadecimal e `%b` para binário.
</details>

<details>
  <summary>4️⃣ Como formatar um número `float64` para exibir apenas duas casas decimais?</summary>
  Usando `fmt.Printf("%.2f", valor)`.
</details>

<details>
  <summary>5️⃣ Para que serve `fmt.Errorf` e como usá-lo?</summary>
  `fmt.Errorf` cria erros formatados com strings personalizadas.
</details>

<details>
  <summary>6️⃣ Qual a vantagem de `fmt.Sprintf` sobre `fmt.Printf`?</summary>
  `fmt.Sprintf` retorna a string formatada sem imprimir diretamente no console.
</details>

<details>
  <summary>7️⃣ Como capturar múltiplos valores de uma única linha de entrada?</summary>
  Usando `fmt.Scanf("%s %d", &nome, &idade)`.
</details>

<details>
  <summary>8️⃣ O que acontece se `fmt.Scan` não conseguir converter a entrada para o tipo esperado?</summary>
  O programa retorna um erro e pode não armazenar corretamente os valores lidos.
</details>

<details>
  <summary>9️⃣ Como redirecionar a saída formatada para um arquivo em vez do terminal?</summary>
  Usando `fmt.Fprint(arquivo, "mensagem formatada")`.
</details>

<details>
  <summary>🔢 Como imprimir um valor dentro de uma string sem usar `fmt.Printf`?</summary>
  Usando `fmt.Sprint("O valor é " + valorString)` ou `fmt.Sprintf("O valor é %d", valor)`.
</details>

---



## **Conclusão**

🚀 **Resumo Final:**

O pacote `fmt` fornece métodos simples e poderosos para entrada e saída de dados. No próximo capítulo, veremos como realizar **conversões de tipos** em Go! 🚀
