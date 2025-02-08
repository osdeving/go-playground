# **3.1 Estruturas Condicionais: `if`, `else if`, `switch`**

O controle de fluxo condicional em Go permite executar diferentes blocos de código com base em condições lógicas. Nesta seção, exploraremos **`if`, `else if`, `switch`**, suas particularidades em Go e como podem ser usadas eficientemente.

---

## **3.1.1 O `if` e `else` em Go**

A estrutura `if` em Go segue um formato semelhante ao de outras linguagens, mas possui peculiaridades importantes:

```go
if condição {
    // Bloco executado se a condição for verdadeira
}
```

Exemplo:

```go
x := 10
if x > 5 {
    fmt.Println("x é maior que 5")
}
```

### **Usando `else` e `else if`**

```go
x := 10

if x < 5 {
    fmt.Println("x é menor que 5")
} else if x == 10 {
    fmt.Println("x é exatamente 10")
} else {
    fmt.Println("x é maior que 5 e diferente de 10")
}
```

📌 **Diferente de algumas linguagens, `if` e `else` em Go não exigem parênteses ao redor da condição!**

```go
if (x > 5) { ... } // ℹ️ Opcional

// Sintaxe recomendada:
if x > 5 { ... } // ✅ Sintaxe recomendada
```

### **Declaração de Variáveis no `if`**

Go permite **declarar variáveis dentro da condição do `if`**, tornando o código mais enxuto:

```go
if y := calcular(); y > 0 {
    fmt.Println("y é positivo:", y)
} else {
    fmt.Println("y é negativo:", y)
}
```

📌 **A variável `y` só existe dentro do escopo do `if` e `else`!**

```go
if y := calcular(); y > 0 {
    fmt.Println("y é positivo:", y)
} else {
    fmt.Println("y é negativo:", y)
}

fmt.Println(y) // ERRO: "y" não existe fora do bloco if
```

---

## **3.1.2 `switch`: Alternativa ao `if-else`**

Em Go, `switch` substitui múltiplas comparações `if-else`, tornando o código mais limpo.

### **Forma básica do `switch`**

```go
dia := "segunda"

switch dia {
case "segunda":
    fmt.Println("Início da semana")
case "sexta":
    fmt.Println("Quase fim de semana!")
case "domingo":
    fmt.Println("Descanso!")
default:
    fmt.Println("Dia normal")
}
```

📌 **Diferente de C e Java, `switch` em Go **NÃO** precisa de `break` em cada `case`!**  
Go **não executa os casos seguintes automaticamente**, a menos que usemos `fallthrough`.

### **Usando `fallthrough` para continuar a execução**

Se quisermos **forçar a execução do próximo caso**, usamos `fallthrough`:

```go
x := 1

switch x {
case 1:
    fmt.Println("Caso 1")
    fallthrough
case 2:
    fmt.Println("Caso 2") // Será executado
}
```

📌 **Atenção! `fallthrough` ignora a condição do próximo `case` e o executa incondicionalmente!**

### **`switch` sem Expressão**

Em Go, um `switch` pode funcionar como um **`if-else` simplificado**, sem expressão inicial:

```go
x := 10

switch {
case x > 10:
    fmt.Println("Maior que 10")
case x == 10:
    fmt.Println("Exatamente 10")
default:
    fmt.Println("Menor que 10")
}
```

📌 **Isso é útil para checar múltiplas condições sem usar `if-else`.**

---

## **3.1.3 `switch` com Tipos (`type switch`)**

Go permite verificar o **tipo dinâmico** de uma variável usando `switch`:

```go
var i interface{} = "texto"

switch v := i.(type) {
case int:
    fmt.Println("É um inteiro:", v)
case string:
    fmt.Println("É uma string:", v)
case bool:
    fmt.Println("É um booleano:", v)
default:
    fmt.Println("Tipo desconhecido")
}
```

📌 **Esse recurso é útil em funções genéricas que lidam com diferentes tipos!**

---

## **3.1.4 Melhorando Performance com `switch`**

Em **casos de múltiplas comparações**, `switch` pode ser **mais rápido** que `if-else`, pois algumas implementações otimizam a avaliação de `case` com tabelas de salto (jump tables). As tabelas de salto são estruturas que mapeiam diretamente os valores dos cases para os endereços de memória do código a ser executado, evitando múltiplas comparações sequenciais como acontece no `if-else`.

**Evite essa sintaxe:**

```go
if x == 1 {
    ...
} else if x == 2 {
    ...
} else if x == 3 {
    ...
}
```

**E dê preferência para essa:**

```go
switch x {
case 1:
    ...
case 2:
    ...
case 3:
    ...
}
```

📌 **Além de mais rápido, `switch` torna o código mais legível.**

---

## **3.1.5 Casos Especiais e Armadilhas**

1. **A comparação entre tipos diferentes causa um erro de compilação**:

```go
var x int = 10
var y float64 = 10.0

// if x == y { ... } // ERRO: Tipos diferentes
```

Nesse caso, converta antes de comparar:

```go
if float64(x) == y { ... } // Correto
```

2. **A exemplo de Java, valores booleanos não são convertidos implicitamente**:

```go
if 1 { ... } // ERRO!
```

Use:

```go
if 1 != 0 { ... } // Correto
```

3. **Não é uma boa prática omitir `default`**:

Se não houver `default`, um `switch` pode não executar nenhum bloco, e isso pode ser indesejado. Por exemplo:

```go
switch dia {
case "segunda":
    fmt.Println("Início da semana")
}
```

Sempre que possível, forneça um `default`:

```go
switch dia {
case "segunda":
    fmt.Println("Início da semana")
default:
    fmt.Println("Dia qualquer")
}
```

---

## **3.1.6 `switch` com Inicialização**

O `switch` em Go permite uma instrução de inicialização, similar ao `if`:

```go
switch x := 10; x {
case 1:
    fmt.Println("x é 1")
case 2:
    fmt.Println("x é 2")
default:
    fmt.Println("x não é 1 nem 2")
}
```

📌 **A variável inicializada (`x`) só existe dentro do escopo do `switch`.**

Você também pode usar funções na inicialização:

```go
switch valor := calcularAlgo(); valor {
case 1:
    fmt.Println("resultado é 1")
case 2:
    fmt.Println("resultado é 2")
default:
    fmt.Println("resultado é outro valor")
}
```

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre estruturas condicionais em Go, tente os seguintes desafios:

🛠️ **Desafios**:

<details>
  <summary>✅ Crie um programa que converte notas numéricas em conceitos usando `switch`.</summary>
  
  ```go
  func converteNota(nota int) string {
      switch {
      case nota >= 90 && nota <= 100:
          return "A"
      case nota >= 80 && nota < 90:
          return "B"
      case nota >= 70 && nota < 80:
          return "C"
      case nota >= 60 && nota < 70:
          return "D"
      case nota >= 0 && nota < 60:
          return "F"
      default:
          return "Nota inválida"
      }
  }
  ```
</details>

<details>
  <summary>✅ Desenvolva uma função que recebe uma interface{} e retorna uma descrição do tipo e valor.</summary>
  
  ```go
  func descreverValor(v interface{}) string {
      switch x := v.(type) {
      case int:
          if x > 0 {
              return "Inteiro positivo"
          }
          return "Inteiro não positivo"
      case string:
          if len(x) == 0 {
              return "String vazia"
          }
          return "String com conteúdo"
      case bool:
          if x {
              return "Booleano verdadeiro"
          }
          return "Booleano falso"
      default:
          return "Tipo desconhecido"
      }
  }
  ```
</details>

<details>
  <summary>✅ Implemente uma calculadora simples usando estruturas condicionais.</summary>
  
  ```go
  func calcular(a, b float64, op string) (float64, error) {
      switch op {
      case "+":
          return a + b, nil
      case "-":
          return a - b, nil
      case "*":
          return a * b, nil
      case "/":
          if b == 0 {
              return 0, fmt.Errorf("divisão por zero")
          }
          return a / b, nil
      default:
          return 0, fmt.Errorf("operação inválida")
      }
  }
  ```
</details>

<details>
  <summary>✅ Crie um programa que determine se um ano é bissexto.</summary>
  
  ```go
  func ehBissexto(ano int) bool {
      switch {
      case ano%400 == 0:
          return true
      case ano%100 == 0:
          return false
      case ano%4 == 0:
          return true
      default:
          return false
      }
  }
  ```
</details>

<details>
  <summary>✅ Desenvolva um validador de senhas que verifique as seguintes regras: mínimo de 8 caracteres, pelo menos 1 número, pelo menos 1 letra, pelo menos 1 caractere especial e não pode conter espaços.</summary>
  
  ```go
  func validarSenha(senha string) bool {
      if len(senha) < 8 {
          return false
      }
      
      temNumero := false
      temLetra := false
      temEspecial := false
      
      for _, c := range senha {
          switch {
          case unicode.IsNumber(c):
              temNumero = true
          case unicode.IsLetter(c):
              temLetra = true
          case unicode.IsPunct(c) || unicode.IsSymbol(c):
              temEspecial = true
          case unicode.IsSpace(c):
              return false // não permite espaços
          }
      }
      
      return temNumero && temLetra && temEspecial
  }
  ```
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>💡 Por que Go não exige `break` em cada `case` do `switch`?</summary>
  Go foi projetada para evitar erros comuns de fallthrough acidental. Cada `case` automaticamente termina a execução do `switch`.
</details>

<details>
  <summary>💡 Como funciona o `type switch` em Go e quando devemos usá-lo?</summary>
  O `type switch` permite verificar o tipo dinâmico de uma interface{} usando a sintaxe v.(type). É útil quando precisamos tratar diferentes tipos de forma específica.
</details>

<details>
  <summary>💡 Qual a diferença entre usar múltiplos `if-else` e `switch`?</summary>
  `switch` geralmente resulta em código mais limpo e pode ser mais eficiente devido à otimização do compilador com tabelas de salto.
</details>

<details>
  <summary>💡 Como Go lida com a comparação entre tipos diferentes?</summary>
  Go não permite comparação direta entre tipos diferentes. É necessário fazer conversão explícita antes de comparar.
</details>

<details>
  <summary>💡 O que acontece com variáveis declaradas dentro de um `if`?</summary>
  Variáveis declaradas na condição do `if` só existem dentro do escopo do `if` e seus blocos `else`.
</details>

<details>
  <summary>💡 Como o `fallthrough` funciona e quando devemos usá-lo?</summary>
  `fallthrough` força a execução do próximo case independentemente de sua condição. Deve ser usado com cautela e apenas quando necessário.
</details>

<details>
  <summary>💡 Por que é importante incluir um `default` no `switch`?</summary>
  O `default` garante um comportamento definido quando nenhum case corresponde à condição, evitando comportamentos indefinidos.
</details>

<details>
  <summary>💡 Como Go trata a avaliação de condições em estruturas `if`?</summary>
  Go avalia as condições de forma estrita, exigindo que resultem em um booleano. Não há conversão implícita de outros tipos para bool.
</details>

<details>
  <summary>💡 Qual a vantagem de declarar variáveis no `if`?</summary>
  Limita o escopo da variável apenas onde é necessária, evitando poluição do namespace e tornando o código mais seguro.
</details>

<details>
  <summary>💡 Como otimizar múltiplas comparações em Go?</summary>
  Use `switch` em vez de múltiplos `if-else`, pois o compilador pode otimizar usando tabelas de salto.
</details>

---

## **Conclusão**

🚀 As estruturas condicionais em Go são projetadas para serem simples, seguras e eficientes. O `switch` é especialmente poderoso, oferecendo funcionalidades além das encontradas em outras linguagens. No próximo capítulo, exploraremos os laços de repetição em Go!
