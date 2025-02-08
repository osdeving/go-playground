# **1.5 Estrutura de um Programa Go**

Todo programa em Go segue uma estrutura básica que inclui pacotes, importação de módulos, funções e a função `main()`. Esta seção explora os principais componentes da estrutura de um programa Go e suas convenções.

---

## **1.5.1 A Estrutura Básica de um Programa Go**

Abaixo está um exemplo de um programa Go mínimo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### **Explicação do código**:

1. **`package main`**: Define o pacote principal do programa. Todo programa executável em Go deve ter um pacote `main`.
2. **`import "fmt"`**: Importa o pacote `fmt`, utilizado para manipulação de entrada e saída de dados.
3. **`func main()`**: A função `main()` é o ponto de entrada do programa. Quando o programa é executado, essa função será chamada.
4. **`fmt.Println("Hello, Go!")`**: Imprime uma mensagem na saída padrão.

---

## **1.5.2 Pacotes e Organização do Código**

Em Go, todo código-fonte pertence a um **pacote**. Os pacotes ajudam a modularizar e reutilizar código.

### **Pacotes Padrão vs. Pacotes Personalizados**

- **Pacotes padrão**: São fornecidos pela biblioteca padrão do Go (ex.: `fmt`, `math`, `net/http`).
- **Pacotes personalizados**: Criados pelo próprio desenvolvedor para organizar código.

### **Criando um Pacote Personalizado**

1. Crie um diretório chamado `meupacote/`:
   ```sh
   mkdir meupacote
   ```
2. Crie um arquivo `meupacote/util.go` com o seguinte código:
   ```go
   package meupacote

   import "fmt"

   func Ola(nome string) {
       fmt.Printf("Olá, %s!\n", nome)
   }
   ```
3. Agora, no seu `main.go`, importe o pacote e use-o:
   ```go
   package main

   import (
       "meupacote"
   )

   func main() {
       meupacote.Ola("Go Developer")
   }
   ```

📌 **Observação**: Para que o Go reconheça o pacote, ele deve estar no `GOPATH` ou em um módulo (`go mod`).

---

## **1.5.3 Importação de Múltiplos Pacotes**

Podemos importar vários pacotes no mesmo arquivo:

```go
import (
    "fmt"
    "math"
    "time"
)
```

Se um pacote for importado mas não for utilizado, o Go exibirá um erro. Para evitar isso, podemos usar `_` para importação sem uso:

```go
import _ "os"
```

Isso é útil quando apenas queremos inicializar um pacote sem usá-lo diretamente.

---

## **1.5.4 Comentários em Go**

Go suporta dois tipos de comentários:

1. **Comentários de linha única**:
   ```go
   // Isso é um comentário
   fmt.Println("Olá, Go!")
   ```

2. **Comentários de múltiplas linhas**:
   ```go
   /*
      Isso é um comentário
      de múltiplas linhas.
   */
   ```

Comentários são fundamentais para documentar código e são usados em conjunto com `go doc` para gerar documentação automática.

---

## **1.5.5 Convenções de Nomenclatura**

Em Go, a nomenclatura segue algumas regras importantes:

- **Identificadores iniciando com letra maiúscula** são **exportados** (públicos) e podem ser acessados de outros pacotes.
- **Identificadores iniciando com letra minúscula** são **privados** ao pacote.

Exemplo:

```go
package exemplo

var Publico = "Eu sou acessível fora do pacote"
var privado = "Sou acessível apenas dentro do pacote"
```

---

## **1.5.6 Executando e Compilando um Programa Go**

### **Executando um Programa Diretamente**

Podemos executar um programa Go sem compilar manualmente:

```sh
go run main.go
```

Isso compila e executa o código temporariamente.

### **Compilando um Programa**

Para gerar um binário executável:

```sh
go build main.go
```

Isso cria um arquivo executável (`main` no Linux/macOS ou `main.exe` no Windows) que pode ser executado diretamente.

---

## **Conclusão**

Agora que entendemos a estrutura de um programa Go, podemos seguir para conceitos mais avançados, como manipulação de variáveis, tipos e controle de fluxo. 🚀
