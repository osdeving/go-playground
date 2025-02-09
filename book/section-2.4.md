# **2.4 Entrada e Saída com `fmt`**

>🗨️ "The best interface is no interface. The best interaction is no interaction. The best program is the one that requires the least input to produce the most output." - Alan Kay, pioneiro da computação pessoal

ℹ️ **Nota ao Leitor**: Esta seção introduz alguns conceitos que serão explorados em maior profundidade mais adiante no livro, como ponteiros (&), tratamento de erros, interfaces e manipulação de arquivos. Não se preocupe se alguns desses temas parecerem complexos agora - cada um deles será abordado detalhadamente em seus respectivos capítulos. Por enquanto, foque em entender os conceitos básicos de entrada e saída.

O pacote `fmt` é a principal ferramenta de entrada e saída em Go. Ele fornece funções para exibir mensagens na tela e ler entradas do usuário. Além do `fmt`, existem outros pacotes úteis para entrada e saída, como `bufio` e `io`.


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

📌 **Placeholders e Flags de Formatação:**

| Placeholder/Flag | Tipo/Uso | Exemplo |
|-----------------|----------|---------|
| `%d` | Inteiro | `15` |
| `%f` | Float | `3.14` |
| `%s` | String | `"texto"` |
| `%t` | Booleano | `true` |
| `%v` | Valor genérico | `{1 2 3}` |
| `%+v` | Struct com nomes de campos | `{Nome:"João" Idade:25}` |
| `%#v` | Notação Go-syntax | `main.Pessoa{Nome:"João", Idade:25}` |
| `%T` | Tipo da variável | `string` |
| `%.2f` | Float com 2 casas decimais | `3.14` |
| `%q` | String com aspas | `"texto"` |
| `%x` | Hexadecimal | `6A` |
| `%b` | Binário | `1010` |
| `%9.2f` | Largura mínima 9, 2 decimais | `   3.14` |
| `%-9s` | Alinhamento à esquerda | `"texto   "` |
| `%09d` | Padding com zeros | `000000123` |

Exemplo:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

p := Pessoa{Nome: "João", Idade: 25}
num := 123.456

fmt.Printf("Struct %%+v: %+v\n", p)     // {Nome:João Idade:25}
fmt.Printf("Go syntax %%#v: %#v\n", p)   // main.Pessoa{Nome:"João", Idade:25}
fmt.Printf("String com aspas %%q: %q\n", "texto") // "texto"
fmt.Printf("Hexadecimal %%x: %x\n", 106)         // 6a
fmt.Printf("Científico %%e: %e\n", num)          // 1.234560e+02
fmt.Printf("Padding %%09d: %09d\n", 123)         // 000000123
```

### **`println()`** – Função embutida no Go

Além das funções do pacote `fmt`, Go possui a função embutida `println()` que imprime uma linha com uma quebra de linha no final. No entanto, ela é menos flexível e não deve ser usada em produção. Essa função não precisa de importação e pode ser usada diretamente no código.

```go
println("Olá, mundo!")
```

---

## **2.4.2 Lendo Entrada do Usuário (`fmt.Scan`, `fmt.Scanln`, `fmt.Scanf`)**

Go oferece várias funções para capturar entrada do usuário, cada uma com suas particularidades:

### **1. `fmt.Scan()`** – Captura múltiplos valores separados por espaço

```go
var nome string
var idade int

fmt.Print("Digite seu nome e idade: ")
n, err := fmt.Scan(&nome, &idade)
if err != nil {
    fmt.Println("Erro na leitura:", err)
    return
}
fmt.Printf("Lidos %d valores. Nome: %s, Idade: %d\n", n, nome, idade)
```

### **2. `fmt.Scanln()`** – Lê até a quebra de linha

```go
var nome string
var idade int

fmt.Print("Digite seu nome: ")
fmt.Scanln(&nome)
fmt.Print("Digite sua idade: ")
fmt.Scanln(&idade)

fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
```

### **3. `fmt.Scanf()`** – Entrada formatada com padrão específico

```go
var dia, mes, ano int

fmt.Print("Digite uma data (DD/MM/AAAA): ")
n, err := fmt.Scanf("%d/%d/%d", &dia, &mes, &ano)
if err != nil {
    fmt.Println("Formato inválido. Use DD/MM/AAAA")
    return
}
fmt.Printf("Data: %02d/%02d/%04d\n", dia, mes, ano)
```

### **4. Funções `Sscan` para Parsing de Strings**

Além da leitura do teclado, podemos fazer parsing de strings:

```go
var x, y int
input := "123 456"

// Sscanf - parsing com formato específico
fmt.Sscanf(input, "%d %d", &x, &y)
fmt.Printf("x=%d, y=%d\n", x, y)

// Sscan - parsing simples separado por espaços
input2 := "789 012"
fmt.Sscan(input2, &x, &y)
```

---

## **2.4.3 Lidando com Erros de Entrada**

O tratamento de erros é fundamental ao trabalhar com entrada de dados:

```go
var idade int
fmt.Print("Digite sua idade: ")
_, err := fmt.Scan(&idade)

switch {
case err == io.EOF:
    fmt.Println("Entrada terminada pelo usuário")
case err != nil:
    fmt.Println("Erro na leitura:", err)
    return
default:
    if idade < 0 {
        fmt.Println("Idade não pode ser negativa")
        return
    }
    fmt.Println("Idade válida:", idade)
}
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

📌 **Sempre use `defer arquivo.Close()` para garantir que o arquivo seja fechado corretamente.**. Esse tópico e o uso de `defer` será abordado com detalhes em capítulos futuros.

---

## **2.4.5 Usando Cores no Terminal**

Para adicionar cores ao texto no terminal, você pode usar pacotes de terceiros como `github.com/fatih/color`.

```go
package main

import (
    "github.com/fatih/color"
)

func main() {
    color.Red("Este texto é vermelho")
    color.Green("Este texto é verde")
}
```
O uso de importação de pacotes de terceiros será abordado com mais detalhes em capítulos futuros.

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre entrada e saída com `fmt`, tente os seguintes desafios:

🛠️ **Desafios**:

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

<details>
  <summary>5️⃣ Crie um programa que leia nome e notas de um aluno e calcule a média com 2 casas decimais.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var nome string
      var nota1, nota2, nota3 float64
      
      fmt.Print("Nome do aluno: ")
      fmt.Scanln(&nome)
      fmt.Print("Digite as três notas: ")
      fmt.Scan(&nota1, &nota2, &nota3)
      
      media := (nota1 + nota2 + nota3) / 3
      fmt.Printf("Aluno: %s\nMédia: %.2f\n", nome, media)
  }
  ```
</details>

<details>
  <summary>6️⃣ Desenvolva um programa que leia um valor em reais e mostre a formatação em diferentes moedas.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var valor float64
      fmt.Print("Digite um valor em reais: ")
      fmt.Scan(&valor)
      
      fmt.Printf("R$ %9.2f (BRL)\n", valor)
      fmt.Printf("$ %9.2f (USD)\n", valor/5.0)  // taxa fictícia
      fmt.Printf("€ %9.2f (EUR)\n", valor/6.0)  // taxa fictícia
  }
  ```
</details>

<details>
  <summary>7️⃣ Crie um programa que leia dados de um produto e salve em um arquivo.</summary>
  
  ```go
  package main
  import (
      "fmt"
      "os"
  )

  func main() {
      var nome string
      var preco float64
      var quantidade int
      
      fmt.Print("Nome do produto: ")
      fmt.Scanln(&nome)
      fmt.Print("Preço: ")
      fmt.Scanln(&preco)
      fmt.Print("Quantidade: ")
      fmt.Scanln(&quantidade)
      
      arquivo, _ := os.Create("produto.txt")
      defer arquivo.Close()
      
      fmt.Fprintf(arquivo, "Produto: %s\nPreço: R$ %.2f\nQuantidade: %d\n", 
          nome, preco, quantidade)
  }
  ```
</details>

<details>
  <summary>8️⃣ Faça um programa que leia uma data no formato DD/MM/AAAA e valide se é uma data válida.</summary>
  
  ```go
  package main
  import "fmt"

  func main() {
      var dia, mes, ano int
      
      fmt.Print("Digite uma data (DD/MM/AAAA): ")
      _, err := fmt.Scanf("%d/%d/%d", &dia, &mes, &ano)
      
      if err != nil || dia < 1 || dia > 31 || mes < 1 || mes > 12 {
          fmt.Println("Data inválida!")
          return
      }
      
      fmt.Printf("Data: %02d/%02d/%04d\n", dia, mes, ano)
  }
  ```
</details>

<details>
  <summary>9️⃣ Desenvolva um programa que leia um texto e conte quantas vogais ele possui.</summary>
  
  ```go
  package main
  import (
      "fmt"
      "strings"
  )

  func main() {
      var texto string
      fmt.Print("Digite um texto: ")
      fmt.Scanln(&texto)
      
      vogais := 0
      for _, c := range strings.ToLower(texto) {
          if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
              vogais++
          }
      }
      
      fmt.Printf("O texto possui %d vogais\n", vogais)
  }
  ```
</details>

<details>
  <summary>🔟 [Avançado] Crie um mini sistema de caixa eletrônico (ATM).</summary>
  
  ```go
  package main
  import (
      "fmt"
      "os"
  )

  // Nota: Este é um desafio mais complexo que utiliza conceitos que serão
  // abordados em capítulos futuros. Recomenda-se voltar a este exercício
  // após estudar estruturas de controle, funções, structs e manipulação
  // de arquivos.

  func main() {
      saldo := 1000.0
      arquivo, _ := os.Create("transacoes.txt")
      defer arquivo.Close()

      for {
          fmt.Println("\n=== CAIXA ELETRÔNICO ===")
          fmt.Println("1. Consultar saldo")
          fmt.Println("2. Fazer depósito")
          fmt.Println("3. Fazer saque")
          fmt.Println("4. Sair")
          
          var opcao int
          fmt.Print("\nEscolha uma opção: ")
          fmt.Scan(&opcao)
          
          switch opcao {
          case 1:
              fmt.Printf("\nSeu saldo é: R$ %.2f\n", saldo)
              fmt.Fprintf(arquivo, "Consulta de saldo: R$ %.2f\n", saldo)
              
          case 2:
              var valor float64
              fmt.Print("Valor do depósito: R$ ")
              fmt.Scan(&valor)
              if valor > 0 {
                  saldo += valor
                  fmt.Printf("Depósito de R$ %.2f realizado com sucesso!\n", valor)
                  fmt.Fprintf(arquivo, "Depósito: R$ %.2f\n", valor)
              } else {
                  fmt.Println("Valor inválido!")
              }
              
          case 3:
              var valor float64
              fmt.Print("Valor do saque: R$ ")
              fmt.Scan(&valor)
              if valor > 0 && valor <= saldo {
                  saldo -= valor
                  fmt.Printf("Saque de R$ %.2f realizado com sucesso!\n", valor)
                  fmt.Fprintf(arquivo, "Saque: R$ %.2f\n", valor)
              } else {
                  fmt.Println("Valor inválido ou saldo insuficiente!")
              }
              
          case 4:
              fmt.Println("Obrigado por usar nosso banco!")
              return
              
          default:
              fmt.Println("Opção inválida!")
          }
      }
  }
  ```
  
  **Nota**: Este último desafio utiliza conceitos como loops, switch-case, manipulação de arquivos
  e estruturas de controle que serão abordados em detalhes nos próximos capítulos. 
  Recomenda-se voltar a este exercício após estudar esses conceitos para melhor compreensão
  e possível implementação de melhorias como:
  - Uso de cores no terminal
  - Validações mais robustas
  - Persistência de dados
  - Múltiplas contas
  - Histórico de transações
  - Transferências entre contas
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

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
