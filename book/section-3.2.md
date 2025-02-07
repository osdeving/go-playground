# **3.2 Laços de Repetição: `for`, `range`**

Go utiliza apenas uma estrutura de repetição: **`for`**. No entanto, sua sintaxe é flexível o suficiente para cobrir diferentes cenários, incluindo loops tradicionais, iterações sobre coleções e loops infinitos.

---

## **3.2.1 Estrutura Básica do `for`**

A forma mais comum do `for` em Go segue o padrão de três expressões presentes em outras linguagens:

```go
for inicialização; condição; incremento {
    // Código a ser repetido
}
```

Exemplo:

```go
for i := 0; i < 5; i++ {
    fmt.Println("Iteração:", i)
}
```

📌 **Diferente de C e Java, Go não suporta `while` e `do-while`, pois `for` cobre todos esses casos.**

### **O formato `while` em Go**

Podemos usar `for` sem a inicialização e incremento, criando um loop estilo `while`:

```go
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}
```


### **Loop Infinito**

Se omitirmos todas as expressões, teremos um loop infinito:

```go
for {
    fmt.Println("Rodando para sempre...")
}
```

📌 **Loop infinito é útil para servidores e processos que nunca devem encerrar.**

---

Os 3 formatos de `for` em Go são:

* for inicialização; condição; incremento { }
* for condição { } (estilo while)
* for { } (loop infinito)



## **3.2.2 Iterando sobre Arrays, Slices e Mapas com `range`**

Go fornece a palavra-chave `range` para percorrer **arrays, slices, strings, mapas e canais** de forma simplificada.

### **Iterando sobre um Slice**

```go
numeros := []int{10, 20, 30}

for indice, valor := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
}
```

📌 **Se o índice não for necessário, use `_` para ignorá-lo:**

```go
for _, valor := range numeros {
    fmt.Println("Valor:", valor)
}
```

### **Iterando sobre um Mapa**

```go
alunos := map[string]int{"Alice": 20, "Bob": 25}

for nome, idade := range alunos {
    fmt.Printf("%s tem %d anos\n", nome, idade)
}
```

### **Iterando sobre uma String (`rune` por `rune`)**

Strings em Go são codificadas em **UTF-8**. Usando `range`, podemos percorrer os caracteres:

```go
s := "GoLang"

for i, r := range s {
    fmt.Printf("Índice: %d, Caractere: %c\n", i, r)
}
```

📌 **Podemos usar `range` para manipulação correta de strings Unicode!**

---

## **3.2.3 Uso de `break` e `continue`**

### **Interrompendo o Loop com `break`**

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Sai do loop quando i == 5
    }
    fmt.Println(i)
}
```

### **Pulando uma Iteração com `continue`**

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue // Pula a iteração quando i == 2
    }
    fmt.Println(i)
}
```

📌 **`break` e `continue` funcionam tanto em loops normais quanto com `range`.**

---

## **3.2.4 Rotulando Loops para Controle Avançado**

Go permite **rotular loops** para controlar `break` e `continue` em loops aninhados:

```go
externo:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break externo // Sai do loop externo
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

📌 **Os rótulos são úteis para sair de loops aninhados sem usar flags booleanas.**

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre laços de repetição em Go, tente os seguintes desafios:

🛠️ **Desafios**:

<details>
  <summary>✅ Crie um programa que imprima a tabuada de um número usando `for` tradicional.</summary>
  
  ```go
  func tabuada(n int) {
      for i := 1; i <= 10; i++ {
          fmt.Printf("%d x %d = %d\n", n, i, n*i)
      }
  }
  ```
</details>

<details>
  <summary>✅ Implemente um programa que conte quantas vogais existem em uma string usando `for range`.</summary>
  
  ```go
  func contaVogais(texto string) int {
      vogais := 0
      for _, c := range texto {
          switch c {
          case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
              vogais++
          }
      }
      return vogais
  }
  ```
</details>

<details>
  <summary>✅ Desenvolva uma função que encontre o maior valor em um slice usando loop.</summary>
  
  ```go
  func maiorValor(nums []int) int {
      if len(nums) == 0 {
          return 0
      }
      maior := nums[0]
      for _, num := range nums {
          if num > maior {
              maior = num
          }
      }
      return maior
  }
  ```
</details>

<details>
  <summary>✅ Crie um programa que imprima um padrão de asteriscos usando loops aninhados.</summary>
  
  ```go
  func imprimirPadrao(n int) {
      for i := 1; i <= n; i++ {
          for j := 1; j <= i; j++ {
              fmt.Print("*")
          }
          fmt.Println()
      }
  }
  ```
</details>

<details>
  <summary>✅ Implemente um programa que verifique se uma palavra é palíndromo usando `for range`.</summary>
  
  ```go
  func ehPalindromo(palavra string) bool {
      runes := []rune(palavra)
      for i := 0; i < len(runes)/2; i++ {
          if runes[i] != runes[len(runes)-1-i] {
              return false
          }
      }
      return true
  }
  ```
</details>

<details>
  <summary>✅ Desenvolva uma função que calcule o fatorial de um número usando loop.</summary>
  
  ```go
  func fatorial(n int) int {
      resultado := 1
      for i := 2; i <= n; i++ {
          resultado *= i
      }
      return resultado
  }
  ```
</details>

<details>
  <summary>✅ Crie um programa que itere sobre um mapa e imprima chaves e valores ordenados.</summary>
  
  ```go
  func imprimirMapaOrdenado(m map[string]int) {
      var chaves []string
      for k := range m {
          chaves = append(chaves, k)
      }
      sort.Strings(chaves)
      for _, k := range chaves {
          fmt.Printf("%s: %d\n", k, m[k])
      }
  }
  ```
</details>

<details>
  <summary>✅ Implemente uma função que encontre números primos até N usando loop.</summary>
  
  ```go
  func primosAteN(n int) []int {
      primos := []int{}
      for i := 2; i <= n; i++ {
          ehPrimo := true
          for j := 2; j*j <= i; j++ {
              if i%j == 0 {
                  ehPrimo = false
                  break
              }
          }
          if ehPrimo {
              primos = append(primos, i)
          }
      }
      return primos
  }
  ```
</details>

<details>
  <summary>✅ Desenvolva um programa que simule um jogo de adivinhação com número limitado de tentativas.</summary>
  
  ```go
  func jogoAdivinhacao(maxTentativas int) {
      numero := rand.Intn(100) + 1
      for i := 0; i < maxTentativas; i++ {
          fmt.Printf("Tentativa %d/%d: ", i+1, maxTentativas)
          var palpite int
          fmt.Scan(&palpite)
          
          if palpite == numero {
              fmt.Println("Parabéns! Você acertou!")
              return
          }
          if palpite < numero {
              fmt.Println("Tente um número maior")
          } else {
              fmt.Println("Tente um número menor")
          }
      }
      fmt.Printf("Game over! O número era %d\n", numero)
  }
  ```
</details>

<details>
  <summary>✅ Crie uma função que remova elementos duplicados de um slice usando loops.</summary>
  
  ```go
  func removerDuplicados(slice []int) []int {
      encontrados := make(map[int]bool)
      resultado := []int{}
      
      for _, valor := range slice {
          if !encontrados[valor] {
              encontrados[valor] = true
              resultado = append(resultado, valor)
          }
      }
      return resultado
  }
  ```
</details>

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>💡 Por que Go tem apenas a estrutura `for` e não possui `while` ou `do-while`?</summary>
  
  Go prioriza simplicidade e clareza. A estrutura `for` é flexível o suficiente para cobrir todos os casos de uso de loops, eliminando redundância na linguagem.
</details>

<details>
  <summary>💡 Qual a diferença entre usar `for range` e um loop tradicional ao iterar sobre uma string?</summary>
  
  `for range` em strings itera sobre runes (caracteres Unicode), enquanto o loop tradicional itera sobre bytes. `for range` é mais seguro para strings UTF-8.
</details>

<details>
  <summary>💡 Como o `range` se comporta com diferentes tipos de dados (slice, map, string)?</summary>
  
  Com slices/arrays: retorna índice e valor
  Com maps: retorna chave e valor
  Com strings: retorna índice e rune (caractere Unicode)
</details>

<details>
  <summary>💡 O que acontece se modificarmos a variável de iteração dentro de um loop `for range`?</summary>
  
  A modificação não afeta a iteração, pois a variável é uma cópia do valor original em cada iteração.
</details>

<details>
  <summary>💡 Como podemos interromper um loop infinito de forma segura?</summary>
  
  Usando `break` quando uma condição específica é atingida, ou usando `os.Exit()` em casos extremos.
</details>

<details>
  <summary>💡 Qual a diferença entre `break` e `continue`?</summary>
  
  `break` sai completamente do loop, enquanto `continue` pula para a próxima iteração.
</details>

<details>
  <summary>💡 Como funcionam os rótulos (labels) em loops aninhados?</summary>
  
  Rótulos permitem especificar qual loop deve ser afetado por `break` ou `continue` em loops aninhados.
</details>

<details>
  <summary>💡 Por que devemos evitar loops infinitos sem condição de saída?</summary>
  
  Loops infinitos sem condição de saída podem consumir recursos do sistema indefinidamente e tornar o programa irresponsivo.
</details>

<details>
  <summary>💡 Como o garbage collector lida com variáveis declaradas dentro de um loop?</summary>
  
  Variáveis declaradas dentro do loop são elegíveis para coleta de lixo quando saem do escopo em cada iteração.
</details>

<details>
  <summary>💡 Qual a forma mais eficiente de iterar sobre um slice grande?</summary>
  
  `for range` é geralmente a forma mais eficiente, pois o compilador pode otimizar a iteração e evitar cópias desnecessárias.
</details>

---

## **Conclusão**

🚀 Os laços de repetição em Go são simples mas poderosos, oferecendo uma única estrutura `for` que cobre todos os casos de uso comuns. O `range` torna a iteração sobre coleções mais segura e idiomática. No próximo capítulo, exploraremos funções em Go! 
