# **13.1 Manipulação de Arquivos (`os`, `io/ioutil`)**

A manipulação de arquivos é uma tarefa essencial em qualquer linguagem de programação.  
Em Go, a biblioteca padrão fornece pacotes poderosos, como **`os`**, **`io`**, **`ioutil`** e **`bufio`**, para lidar com **leitura, escrita e gerenciamento de arquivos** de maneira eficiente e segura.

Nesta seção, exploraremos:

- Como abrir, criar, ler e escrever arquivos em Go
- Diferenças entre os pacotes `os`, `io`, `ioutil` e `bufio`
- Manipulação de arquivos grandes de forma eficiente
- Tratamento de erros ao lidar com arquivos
- Melhores práticas para garantir segurança e desempenho

---

## **13.1.1 Criando e Abrindo Arquivos**

Para criar ou abrir arquivos, usamos a função `os.OpenFile()`, que permite especificar **permissões e modos de abertura**.

✅ **Exemplo: Criando um novo arquivo**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Erro ao criar o arquivo:", err)
        return
    }
    defer file.Close()

    fmt.Println("Arquivo criado com sucesso!")
}
```

📌 **O arquivo será criado no diretório atual e fechado corretamente ao final do programa.**  

✅ **Exemplo: Abrindo um arquivo existente para leitura**

```go
file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("Erro ao abrir o arquivo:", err)
    return
}
defer file.Close()
```

📌 **Se o arquivo não existir, `os.Open` retornará um erro.**  

✅ **Exemplo: Abrindo um arquivo para leitura e escrita**

```go
file, err := os.OpenFile("example.txt", os.O_RDWR, 0644)
if err != nil {
    fmt.Println("Erro ao abrir o arquivo:", err)
    return
}
defer file.Close()
```

📌 **O modo `os.O_RDWR` permite leitura e escrita no mesmo arquivo.**  

---

## **13.1.2 Escrevendo em Arquivos**

Podemos escrever em arquivos usando `WriteString()`, `Write()`, ou `fmt.Fprint()`.  

✅ **Exemplo: Escrevendo texto em um arquivo**

```go
file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    fmt.Println("Erro ao abrir o arquivo:", err)
    return
}
defer file.Close()

_, err = file.WriteString("Escrevendo no arquivo!
")
if err != nil {
    fmt.Println("Erro ao escrever no arquivo:", err)
}
```

📌 **Usamos `os.O_APPEND` para adicionar texto ao final do arquivo.**  

✅ **Exemplo: Escrevendo bytes diretamente**

```go
data := []byte("Dados binários")
file.Write(data)
```

📌 **Escrever bytes pode ser útil para manipular arquivos binários.**  

---

## **13.1.3 Lendo Arquivos**

✅ **Exemplo: Lendo um arquivo inteiro com `ioutil.ReadFile`**

```go
import "io/ioutil"

data, err := ioutil.ReadFile("example.txt")
if err != nil {
    fmt.Println("Erro ao ler o arquivo:", err)
    return
}
fmt.Println(string(data)) // Converte bytes para string
```

📌 **`ioutil.ReadFile` carrega todo o arquivo na memória, o que pode ser ineficiente para arquivos grandes.**  

✅ **Exemplo: Lendo arquivo linha por linha com `bufio.Scanner`**

```go
import (
    "bufio"
    "os"
)

file, err := os.Open("example.txt")
if err != nil {
    fmt.Println("Erro ao abrir o arquivo:", err)
    return
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text()) // Exibe cada linha do arquivo
}

if err := scanner.Err(); err != nil {
    fmt.Println("Erro ao ler linha:", err)
}
```

📌 **`bufio.Scanner` é eficiente para ler arquivos grandes sem consumir muita memória.**  

✅ **Exemplo: Lendo um arquivo em chunks (blocos)**

```go
buffer := make([]byte, 100) // Lê 100 bytes por vez
for {
    n, err := file.Read(buffer)
    if err != nil {
        break
    }
    fmt.Print(string(buffer[:n])) // Converte bytes para string
}
```

📌 **Essa abordagem é útil para processar arquivos muito grandes.**  

---

## **13.1.4 Removendo e Renomeando Arquivos**

✅ **Exemplo: Excluindo um arquivo**

```go
err := os.Remove("example.txt")
if err != nil {
    fmt.Println("Erro ao deletar o arquivo:", err)
} else {
    fmt.Println("Arquivo removido com sucesso!")
}
```

✅ **Exemplo: Renomeando um arquivo**

```go
err := os.Rename("example.txt", "newname.txt")
if err != nil {
    fmt.Println("Erro ao renomear o arquivo:", err)
}
```

📌 **`os.Remove` e `os.Rename` são úteis para manipular arquivos dinamicamente.**  

---

## **13.1.5 Manipulação Segura e Tratamento de Erros**

✔ **Sempre feche arquivos com `defer file.Close()` para evitar vazamentos de memória.**  
✔ **Verifique sempre erros ao abrir ou manipular arquivos (`if err != nil { ... }`).**  
✔ **Use `bufio` para ler arquivos grandes de forma eficiente.**  
✔ **Prefira `ioutil.ReadFile` apenas para arquivos pequenos.**  
✔ **Evite carregar arquivos enormes na memória, prefira leitura em blocos.**  

---

## **Conclusão**

O **Go fornece diversas formas de manipular arquivos de maneira eficiente**, desde operações básicas de leitura e escrita até manipulação de arquivos grandes com `bufio`.  
No próximo capítulo, exploraremos **leitura e escrita em formatos estruturados como JSON e CSV**, essenciais para integração com bancos de dados e APIs! 🚀
