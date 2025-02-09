# 📂 13.4 Streaming e Manipulação de Fluxos no Go

## 👋 Introdução

Até agora, aprendemos como manipular arquivos comuns, trabalhar com CSV e JSON. Mas e quando estamos lidando com **grandes volumes de dados** que não podem ser carregados inteiramente na memória? É aí que entra o **streaming** e o uso de fluxos no Go.

Nesta seção, vamos explorar:

- Como **ler e escrever dados de forma eficiente** usando buffers 🚀
- Como **manipular fluxos de entrada e saída** com `io.Reader` e `io.Writer` 🔄
- Como **usar `io.Copy` e `io.TeeReader` para manipulação de fluxos** 📥📤
- Como **processar arquivos grandes sem sobrecarregar a memória** 🏗️

Vamos direto para a prática! 🎯

---

## 📖 Lendo e Escrevendo com Buffers (`bufio`)

O pacote `bufio` permite otimizar a leitura e escrita de dados através de buffers, reduzindo o número de operações de I/O.

### 📝 Lendo um arquivo linha por linha com `bufio.Scanner`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("grande_arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro na leitura:", err)
	}
}
```

🔹 **Explicação:**
- `bufio.NewScanner` permite ler o arquivo linha por linha.
- `scanner.Text()` retorna o conteúdo da linha atual.
- Evita carregar todo o arquivo na memória, ideal para arquivos grandes.

### 📖 Escrevendo em um arquivo usando `bufio.Writer`

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("saida.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(writer, "Linha %d\n", i)
	}
	writer.Flush()
}
```

🔹 **Explicação:**
- `bufio.NewWriter` permite acumular dados antes de gravar no arquivo.
- `Flush()` garante que os dados sejam efetivamente escritos.

---

## 📖 Manipulação de Fluxos (`io.Reader`, `io.Writer`)

No Go, a manipulação de fluxos é baseada nas interfaces `io.Reader` e `io.Writer`, que permitem um processamento eficiente e flexível.

### 📌 Copiando dados de um `Reader` para um `Writer` (`io.Copy`)

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile, err := os.Open("entrada.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de origem:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create("saida.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de destino:", err)
		return
	}
	defer destFile.Close()

	bytesCopied, err := io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Erro ao copiar arquivo:", err)
		return
	}

	fmt.Printf("Cópia concluída! %d bytes copiados.\n", bytesCopied)
}
```

🔹 **Explicação:**
- `io.Copy(dest, src)` lê os dados do `src` (`io.Reader`) e os escreve em `dest` (`io.Writer`).
- Ideal para copiar arquivos sem precisar carregá-los completamente na memória.

### 📖 Duplicando a saída de um `Reader` com `io.TeeReader`

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("entrada.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	tee := io.TeeReader(file, os.Stdout)
	buf := make([]byte, 512)
	_, err = tee.Read(buf)
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}
}
```

🔹 **Explicação:**
- `io.TeeReader` permite ler os dados e enviá-los para múltiplos destinos ao mesmo tempo.
- Ideal para monitorar dados enquanto eles são processados.

---

## 🚀 Streaming e Processamento Eficiente

Ao trabalhar com fluxos de dados, algumas boas práticas são fundamentais:

✅ **Use Buffers**: Reduza chamadas ao sistema operacional com `bufio.Writer` e `bufio.Reader`.
✅ **Prefira Processamento em Streaming**: Evite carregar arquivos inteiros na memória, utilize `io.Reader`.
✅ **Evite `ReadAll()` para arquivos grandes**: Prefira ler pedaços do arquivo.
✅ **Use `io.Copy` para mover dados rapidamente**.

---

## 🎯 Conclusão

Agora você aprendeu:
✅ Como usar `bufio` para ler e escrever arquivos de forma eficiente.
✅ Como manipular fluxos de entrada e saída com `io.Reader` e `io.Writer`.
✅ Como copiar e redirecionar dados de forma otimizada.
✅ Como processar arquivos grandes sem sobrecarregar a memória.

Com isso, estamos prontos para aprofundar ainda mais o uso de entrada e saída no Go! 🚀

