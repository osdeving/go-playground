# 📌 Seção 8.4: Interface `io.Reader` e `io.Writer` em Go

## Introdução

Go possui um poderoso sistema de interfaces que facilita a manipulação de entradas e saídas de dados. Entre as interfaces mais importantes da linguagem, destacam-se `io.Reader` e `io.Writer`, que são fundamentais para leitura e escrita de fluxos de dados.

Este capítulo explora o funcionamento dessas interfaces, como usá-las e como implementá-las em seus próprios tipos.

---

## 🔹 A Interface `io.Reader`

A interface `io.Reader` define um método único para leitura de dados:
```go
package io

type Reader interface {
    Read(p []byte) (n int, err error)
}
```
### 📌 Como funciona?
- O método `Read` lê **até** `len(p)` bytes em `p` e retorna o número real de bytes lidos (`n`).
- Se `Read` atingir o final da entrada, ele retorna `io.EOF`.

### 📌 Exemplo de uso:
```go
package main

import (
    "fmt"
    "strings"
    "io"
)

func main() {
    r := strings.NewReader("Exemplo de leitura com io.Reader")
    buf := make([]byte, 8)
    
    for {
        n, err := r.Read(buf)
        fmt.Printf("Lido: %s\n", buf[:n])
        if err == io.EOF {
            break
        }
    }
}
```

Neste exemplo:
- `strings.NewReader` cria um `io.Reader` a partir de uma string.
- O loop lê 8 bytes por vez até atingir `io.EOF`.

---

## 🔹 A Interface `io.Writer`

A interface `io.Writer` permite a escrita de dados em um destino:
```go
package io

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
### 📌 Como funciona?
- O método `Write` grava `len(p)` bytes do slice `p`.
- Retorna o número real de bytes gravados (`n`).
- Em caso de erro, ele retorna um valor `error`.

### 📌 Exemplo de uso:
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f := os.Stdout
    f.Write([]byte("Escrevendo no stdout usando io.Writer\n"))
}
```

Neste exemplo:
- `os.Stdout` implementa `io.Writer`.
- A função `Write` escreve diretamente no console.

---

## 🔹 Criando um `io.Reader` Personalizado

Podemos criar nosso próprio tipo que implementa `io.Reader`:

```go
package main

import (
    "fmt"
    "io"
)

type MeuReader struct{}

func (MeuReader) Read(p []byte) (n int, err error) {
    for i := range p {
        p[i] = 'A'
    }
    return len(p), nil
}

func main() {
    r := MeuReader{}
    buf := make([]byte, 5)
    r.Read(buf)
    fmt.Println("Lido:", string(buf))
}
```

Aqui, `MeuReader` sempre retorna uma sequência de `'A'`.

---

## 🔹 Criando um `io.Writer` Personalizado

Assim como `io.Reader`, podemos criar um `io.Writer` personalizado:
```go
package main

import (
    "fmt"
    "io"
)

type MeuWriter struct{}

func (MeuWriter) Write(p []byte) (n int, err error) {
    fmt.Println("Recebido:", string(p))
    return len(p), nil
}

func main() {
    w := MeuWriter{}
    w.Write([]byte("Testando io.Writer"))
}
```

Este exemplo:
- Implementa `Write`, imprimindo os dados na tela.
- Retorna o número de bytes escritos.

---

## 🔹 Combinando `io.Reader` e `io.Writer`

Um exemplo prático de como `io.Reader` e `io.Writer` podem ser combinados é a função `io.Copy`, que copia dados de um `Reader` para um `Writer`:

```go
package main

import (
    "io"
    "os"
    "strings"
)

func main() {
    r := strings.NewReader("Copiando de um Reader para um Writer")
    io.Copy(os.Stdout, r)
}
```

Isso copia os dados da string diretamente para `os.Stdout`.

---

## 📌 Conclusão

1. `io.Reader` e `io.Writer` são essenciais para manipulação de dados em Go.
2. Interfaces permitem flexibilidade e abstração na leitura e escrita de dados.
3. Criar implementações personalizadas dessas interfaces pode facilitar a construção de aplicações modulares e reutilizáveis.

🔹 **Dominar `io.Reader` e `io.Writer` é fundamental para desenvolver aplicações eficientes em Go!** 🚀

