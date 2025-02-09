# 📂 13.6 Tratamento de Erros em Entrada e Saída no Go

## 👋 Introdução

Ao trabalhar com **entrada e saída de dados (I/O)**, erros são inevitáveis. Seja ao tentar abrir um arquivo inexistente, ler de um pipe fechado ou escrever em um local sem permissão, precisamos de boas práticas para lidar com esses cenários corretamente.

Nesta seção, vamos explorar:

- Como **identificar e tratar erros em operações de I/O** ⚠️
- Uso do pacote `errors` para manipulação avançada de erros 🚀
- Como **recuperar de falhas sem comprometer o programa** 🛠️
- Estratégias para **logs e debugging eficazes** 🔍

Vamos entender como tornar nosso código mais robusto! 🎯

---

## 📖 Identificando e Tratando Erros em I/O

Em Go, quase todas as operações de I/O retornam um `error`. O tratamento adequado desses erros evita comportamentos inesperados.

### 📝 Exemplo Simples de Tratamento de Erro ao Abrir Arquivo

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("arquivo_inexistente.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Println("Arquivo aberto com sucesso!")
}
```

🔹 **Explicação:**
- `os.Open` retorna um erro caso o arquivo não exista.
- Verificamos `if err != nil` antes de prosseguir.
- `defer file.Close()` fecha o arquivo corretamente caso a abertura tenha sido bem-sucedida.

---

## 📖 Manipulando Erros com `errors.Is` e `errors.As`

O pacote `errors` fornece formas elegantes de manipular erros específicos.

### 📌 Verificando Tipos de Erros com `errors.Is`

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("arquivo_inexistente.txt")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("O arquivo não existe!")
	} else if err != nil {
		fmt.Println("Erro desconhecido:", err)
	}
}
```

🔹 **Explicação:**
- `errors.Is(err, os.ErrNotExist)` verifica se o erro é de arquivo inexistente.
- Isso permite tratarmos erros específicos de forma mais granular.

### 📖 Extraindo Informações de Erros com `errors.As`

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("arquivo_inexistente.txt")
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("Erro relacionado a caminho de arquivo:", pathErr.Path)
	} else if err != nil {
		fmt.Println("Erro desconhecido:", err)
	}
}
```

🔹 **Explicação:**
- `errors.As(err, &pathErr)` verifica se o erro é um `os.PathError`.
- Isso permite capturar mais informações específicas do erro.

---

## 📖 Recuperação de Falhas e Resiliência

Em alguns casos, podemos tentar recuperar de falhas e continuar a execução.

### 📝 Repetindo Tentativas de Leitura

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var file *os.File
	var err error
	for i := 0; i < 3; i++ {
		file, err = os.Open("arquivo_temporario.txt")
		if err == nil {
			break
		}
		fmt.Println("Tentativa", i+1, "falhou, tentando novamente...")
		time.Sleep(time.Second)
	}

	if err != nil {
		fmt.Println("Falha ao abrir o arquivo após múltiplas tentativas:", err)
		return
	}
	defer file.Close()
	fmt.Println("Arquivo aberto com sucesso!")
}
```

🔹 **Explicação:**
- Tentamos abrir o arquivo até **3 vezes** antes de desistir.
- `time.Sleep(time.Second)` adiciona um intervalo entre tentativas.

---

## 📖 Logs e Debugging Eficazes

Para depuração e diagnóstico, podemos usar `log` em vez de `fmt.Println`.

### 📝 Exemplo de Uso do Pacote `log`

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("arquivo.txt")
	if err != nil {
		log.Fatalf("Erro crítico: %v", err)
	}
	defer file.Close()
	log.Println("Arquivo aberto com sucesso")
}
```

🔹 **Explicação:**
- `log.Fatalf` exibe uma mensagem e encerra o programa.
- `log.Println` registra eventos normais sem interromper a execução.

---

## 🚀 Melhores Práticas para Tratamento de Erros

✅ **Sempre verifique erros retornados** ao trabalhar com arquivos e I/O.
✅ **Use `errors.Is` e `errors.As`** para tratamento refinado de erros.
✅ **Evite `panic`** para erros esperados; prefira retornar e lidar com eles.
✅ **Utilize logs (`log.Fatalf`, `log.Println`)** para melhor rastreamento de falhas.
✅ **Implemente estratégias de recuperação** quando possível, como tentativas repetidas.

---

## 🎯 Conclusão

Agora você aprendeu:
✅ Como capturar e tratar erros corretamente em operações de I/O.
✅ Como usar `errors.Is` e `errors.As` para identificar erros específicos.
✅ Como implementar recuperação de falhas para evitar interrupções desnecessárias.
✅ Como utilizar logs para debugging eficaz.

Com esse conhecimento, você está pronto para escrever código mais **resiliente e confiável** em Go! 🚀

