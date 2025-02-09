# 📂 13.5 Trabalhando com Pipes e Comunicação entre Processos no Go

## 👋 Introdução

Na programação moderna, há situações em que precisamos conectar processos ou manipular fluxos de dados dinamicamente. O Go facilita esse tipo de comunicação por meio de **pipes (`os.Pipe`)** e redirecionamento de entrada e saída.

Nesta seção, você aprenderá:

- Como **criar e usar pipes (`os.Pipe`)** 🔄
- Como **redirecionar a entrada e saída de processos (`exec.Cmd`)** 🖥️
- Como **comunicar goroutines via pipes** 🚀

Vamos explorar esses conceitos na prática! 🎯

---

## 📖 Criando e Usando Pipes (`os.Pipe`)

O Go fornece `os.Pipe()`, que cria um par de conexões de leitura e escrita. Podemos usá-los para transferir dados entre processos ou goroutines.

### 📝 Exemplo de Uso de `os.Pipe()`

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, writer := os.Pipe()

	go func() {
		writer.Write([]byte("Olá, Pipe!\n"))
		writer.Close()
	}()

	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("Erro ao ler do pipe:", err)
		return
	}

	fmt.Println("Recebido do pipe:", string(buf[:n]))
}
```

🔹 **Explicação:**
- `os.Pipe()` retorna um `reader` e um `writer`.
- O `writer` escreve dados e o `reader` lê esses dados.
- Esse mecanismo é útil para comunicação entre processos ou goroutines.

---

## 📖 Redirecionando Entrada e Saída de Processos (`exec.Cmd`)

Às vezes, queremos capturar a saída de um comando externo dentro do Go. Podemos fazer isso com `exec.Cmd` e `os.Pipe`.

### 📝 Capturando a Saída de um Comando

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
	}
}
```

🔹 **Explicação:**
- Criamos um comando `ls -l` com `exec.Command`.
- Redirecionamos a saída padrão (`Stdout`) para `os.Stdout`.
- `cmd.Run()` executa o comando e exibe a saída diretamente.

### 📖 Capturando Saída para Manipulação

Podemos capturar a saída para processamento posterior:

```go
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Go é incrível!")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar comando:", err)
		return
	}

	fmt.Println("Saída do comando:", out.String())
}
```

🔹 **Explicação:**
- Criamos um `bytes.Buffer` para armazenar a saída.
- `cmd.Stdout = &out` direciona a saída do comando para o buffer.
- Após `cmd.Run()`, podemos acessar a saída como string.

---

## 📖 Comunicação entre Goroutines via Pipes

Podemos usar `os.Pipe` para comunicação eficiente entre goroutines.

### 📝 Enviando e Recebendo Dados com Pipes em Goroutines

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	reader, writer := os.Pipe()

	go func() {
		for i := 0; i < 5; i++ {
			writer.Write([]byte(fmt.Sprintf("Mensagem %d\n", i+1)))
			time.Sleep(time.Millisecond * 500)
		}
		writer.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println("Fim da comunicação")
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
```

🔹 **Explicação:**
- Criamos um `os.Pipe()` para comunicação.
- Uma goroutine escreve mensagens periodicamente.
- O processo principal lê do pipe e exibe as mensagens.

---

## 🚀 Benefícios do Uso de Pipes e Processos

✅ **Comunicação Eficiente**: Pipes são leves e eficientes para transferência de dados.
✅ **Integração com o Sistema**: `exec.Cmd` permite interagir com comandos externos facilmente.
✅ **Processamento Paralelo**: Pipes entre goroutines ajudam a otimizar o desempenho.
✅ **Captura de Logs**: Podemos usar pipes para processar logs de forma dinâmica.

---

## 🎯 Conclusão

Agora você aprendeu:
✅ Como usar `os.Pipe` para comunicação entre processos e goroutines.
✅ Como redirecionar entrada e saída de comandos com `exec.Cmd`.
✅ Como capturar e processar a saída de processos externos.
✅ Como criar pipelines eficientes para processar dados dinamicamente.

Com isso, você está pronto para criar sistemas que fazem comunicação eficiente entre processos e goroutines! 🚀

