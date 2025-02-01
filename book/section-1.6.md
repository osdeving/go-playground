# **1.6 O Primeiro Programa: "Hello, World!"**

O clássico programa **"Hello, World!"** é frequentemente o primeiro código que desenvolvedores escrevem ao aprender uma nova linguagem. Em Go, ele é simples, mas ensina os conceitos básicos de estrutura e execução.

---

## **1.6.1 Escrevendo o Primeiro Programa**

Abra um editor de texto e crie um arquivo chamado `main.go` com o seguinte código:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### **Explicação do Código**
1. **`package main`**: Define que este arquivo pertence ao pacote `main`, obrigatório para um programa executável em Go.
2. **`import "fmt"`**: Importa o pacote `fmt`, que contém funções para entrada e saída de texto.
3. **`func main()`**: Define a função `main`, que é o ponto de entrada da aplicação.
4. **`fmt.Println("Hello, World!")`**: Exibe a mensagem `"Hello, World!"` na saída padrão.

---

## **1.6.2 Executando o Programa**

### **Com `go run` (modo desenvolvimento)**
Se quiser testar rapidamente, execute:

```sh
go run main.go
```

A saída esperada será:

```
Hello, World!
```

### **Com `go build` (modo produção)**
Para gerar um binário executável:

```sh
go build main.go
```

No Windows, isso criará `main.exe`, enquanto no Linux/macOS gerará `main`. Para executar:

```sh
./main   # Linux/macOS
main.exe # Windows
```

Isso permite rodar o programa sem precisar do Go instalado.

---

## **1.6.3 Personalizando a Saída**

Podemos modificar o programa para aceitar entrada do usuário:

```go
package main

import (
    "fmt"
)

func main() {
    var nome string
    fmt.Print("Digite seu nome: ")
    fmt.Scanln(&nome)
    fmt.Printf("Hello, %s!\n", nome)
}
```

Agora, ao executar:

```
Digite seu nome: João
Hello, João!
```

---

## **1.6.4 Lidando com Erros**

Se o usuário não inserir um nome, o programa pode falhar. Para tratar isso, podemos verificar erros:

```go
package main

import (
    "fmt"
)

func main() {
    var nome string
    fmt.Print("Digite seu nome: ")
    _, err := fmt.Scanln(&nome)

    if err != nil {
        fmt.Println("Erro ao ler entrada.")
        return
    }

    fmt.Printf("Hello, %s!\n", nome)
}
```

Agora, se houver um erro, o programa informará ao usuário.

---

## **Conclusão**

Agora que você escreveu e executou seu primeiro programa em Go, está pronto para aprender sobre variáveis, tipos de dados e controle de fluxo no próximo capítulo! 🚀
