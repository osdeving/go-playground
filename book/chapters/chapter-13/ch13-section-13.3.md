# **13.3 Streaming com `bufio`**

Manipular arquivos e fluxos de entrada/saída de maneira eficiente é essencial para aplicações escaláveis.  
O pacote **`bufio`** fornece uma camada de **buffering** que melhora o desempenho de operações de leitura e escrita,  
especialmente ao lidar com **grandes volumes de dados**.

Nesta seção, exploraremos:

- O que é o `bufio` e como ele melhora a performance
- Leitura eficiente de arquivos grandes linha por linha
- Escrita otimizada com `bufio.Writer`
- Uso do `bufio.Reader` para manipular entradas de `os.Stdin`
- Comparação de desempenho entre `bufio` e `os`

---

## **13.3.1 O Que é `bufio` e Por Que Usá-lo?**

O pacote `bufio` cria **buffers internos** que **reduzem o número de chamadas diretas ao sistema operacional**,
evitando operações de I/O excessivas que impactam o desempenho.

📌 **Sem buffering (`os.Open` lê diretamente do disco, o que pode ser ineficiente):**

```go
file, _ := os.Open("largefile.txt")
defer file.Close()

var data []byte
for {
    buffer := make([]byte, 512) // Lê 512 bytes por vez
    n, err := file.Read(buffer)
    if err != nil {
        break
    }
    data = append(data, buffer[:n]...)
}
```

📌 **Com buffering (`bufio` otimiza a leitura e reduz acessos ao disco):**

```go
file, _ := os.Open("largefile.txt")
defer file.Close()

reader := bufio.NewReader(file)
var data []byte

for {
    buffer, err := reader.Peek(512) // Lê 512 bytes sem consumi-los
    if err != nil {
        break
    }
    data = append(data, buffer...)
    reader.Discard(len(buffer)) // Move o ponteiro da leitura
}
```

✅ **`bufio.Reader` reduz o número de chamadas `syscall.Read`, tornando o processo mais rápido.**  

---

## **13.3.2 Leitura Linha por Linha com `bufio.Scanner`**

Para arquivos **grandes**, carregar todo o conteúdo na memória pode ser ineficiente.  
O `bufio.Scanner` permite **ler linha por linha**, processando cada trecho sem sobrecarregar a RAM.

✅ **Exemplo: Lendo um arquivo linha por linha**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("largefile.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // Processa cada linha
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Erro na leitura:", err)
    }
}
```

📌 **`bufio.Scanner` lê arquivos sem carregar tudo na memória.**  
📌 **Se `largefile.txt` tiver 1GB, a memória consumida será mínima.**  

✅ **Use `bufio.Scanner` para processar logs, arquivos CSV e grandes volumes de texto.**  

---

## **13.3.3 Escrita Eficiente com `bufio.Writer`**

O `bufio.Writer` melhora a performance ao escrever em arquivos, pois armazena temporariamente os dados em um buffer interno  
antes de fazer a escrita real no disco.

✅ **Exemplo: Escrita otimizada com `bufio.Writer`**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    writer.WriteString("Linha 1: Escrita otimizada com bufio!
")
    writer.WriteString("Linha 2: Reduzindo operações de I/O...
")

    writer.Flush() // Grava os dados do buffer no arquivo

    fmt.Println("Arquivo salvo com sucesso!")
}
```

📌 **Sem `bufio.Writer`, cada `WriteString()` faria uma chamada ao SO, o que é ineficiente.**  
📌 **Com `bufio.Writer`, os dados são armazenados em memória e escritos em lote.**  

✅ **Use `Flush()` para garantir que os dados sejam gravados no arquivo.**  

---

## **13.3.4 Manipulando `os.Stdin` com `bufio.Reader`**

Podemos usar `bufio.Reader` para ler entrada do usuário de forma eficiente.  
Isso é útil para **aplicações interativas e processamento de logs.**

✅ **Exemplo: Lendo entrada do usuário**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite algo: ")
    input, _ := reader.ReadString('\n')

    fmt.Println("Você digitou:", input)
}
```

📌 **`ReadString('\n')` lê a entrada até o usuário pressionar ENTER.**  
📌 **O buffer evita leituras desnecessárias do teclado, melhorando a performance.**  

✅ **Ideal para CLIs e ferramentas de linha de comando.**  

---

## **13.3.5 Comparação de Desempenho: `os`, `bufio` e `ioutil`**

| Método | Bufferizado? | Uso de Memória | Performance |
|--------|-------------|---------------|-------------|
| `os.Open().Read()` | ❌ Não | Alta (carrega tudo na RAM) | Médio |
| `ioutil.ReadFile()` | ❌ Não | Muito Alta (carrega tudo) | Rápido, mas perigoso |
| `bufio.Reader` | ✅ Sim | Baixa (processa em blocos) | Alto |
| `bufio.Scanner` | ✅ Sim | Baixíssima (linha por linha) | Alto |
| `bufio.Writer` | ✅ Sim | Baixa (buffer interno) | Alto |

📌 **`ioutil.ReadFile()` deve ser evitado para arquivos grandes.**  
📌 **`bufio.Scanner` e `bufio.Reader` são ideais para processamento eficiente.**  

✅ **Sempre escolha a abordagem correta para evitar consumo excessivo de memória!**  

---

## **Conclusão**

O **pacote `bufio` fornece uma forma eficiente de lidar com I/O**, reduzindo chamadas diretas ao SO e melhorando o desempenho.  
No próximo capítulo, exploraremos **tratamento avançado de erros em operações de entrada e saída**, garantindo que aplicações Go sejam resilientes e confiáveis! 🚀
