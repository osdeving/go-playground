# 📂 13.3 Leitura e Escrita em CSV e JSON no Go

## 👋 Introdução

Depois de aprender a manipular arquivos de forma geral, chegou a hora de trabalhar com **formatos de dados estruturados**. Em muitas aplicações, precisamos lidar com arquivos **CSV** e **JSON**, que são amplamente utilizados para armazenamento e troca de informações.

Nesta seção, você aprenderá:

- Como **ler e escrever arquivos CSV** usando `encoding/csv` 📊
- Como **ler e escrever arquivos JSON** usando `encoding/json` 🛠️
- Como **tratar erros e garantir a integridade dos dados** ⚠️
- Estratégias para **trabalhar com grandes volumes de dados** sem comprometer a performance 🚀

Vamos explorar esses conceitos na prática! 🎯

---

## 📖 Trabalhando com CSV

O formato **CSV (Comma-Separated Values)** é muito utilizado para armazenar dados tabulares. No Go, o pacote `encoding/csv` facilita a leitura e escrita desse tipo de arquivo.

### 📝 Escrevendo em CSV

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("dados.csv")
	if err != nil {
		fmt.Println("Erro ao criar arquivo CSV:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	dados := [][]string{
		{"Nome", "Idade", "Cidade"},
		{"Alice", "30", "São Paulo"},
		{"Bob", "25", "Rio de Janeiro"},
	}

	for _, linha := range dados {
		if err := writer.Write(linha); err != nil {
			fmt.Println("Erro ao escrever no CSV:", err)
		}
	}

	writer.Flush() // Garante que os dados sejam escritos no arquivo
}
```

🔹 **Explicação:**
- Criamos um arquivo CSV com `os.Create`.
- Usamos `csv.NewWriter` para escrever dados.
- `writer.Flush()` garante que os dados sejam persistidos corretamente.

### 📖 Lendo um CSV

```go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("dados.csv")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	registros, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler CSV:", err)
		return
	}

	for _, linha := range registros {
		fmt.Println(linha)
	}
}
```

🔹 **Explicação:**
- `csv.NewReader` é usado para ler os dados do arquivo.
- `reader.ReadAll()` carrega todas as linhas de uma vez.

⚠️ **Cuidado:** Se o arquivo CSV for muito grande, prefira ler linha por linha com `reader.Read()` para economizar memória.

---

## 📖 Trabalhando com JSON

O formato **JSON (JavaScript Object Notation)** é muito usado para APIs e armazenamento de dados estruturados. No Go, utilizamos o pacote `encoding/json` para lidar com esse formato.

### 📝 Convertendo uma Struct para JSON

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Cidade string `json:"cidade"`
}

func main() {
	pessoa := Pessoa{"Alice", 30, "São Paulo"}

	file, err := os.Create("dados.json")
	if err != nil {
		fmt.Println("Erro ao criar arquivo JSON:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pessoa); err != nil {
		fmt.Println("Erro ao escrever JSON:", err)
	}
}
```

🔹 **Explicação:**
- Criamos uma struct `Pessoa`.
- Utilizamos `json.NewEncoder` para gravar o JSON diretamente no arquivo.

### 📖 Lendo JSON e Convertendo para Struct

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Cidade string `json:"cidade"`
}

func main() {
	file, err := os.Open("dados.json")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo JSON:", err)
		return
	}
	defer file.Close()

	var pessoa Pessoa
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pessoa); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	fmt.Printf("Nome: %s, Idade: %d, Cidade: %s\n", pessoa.Nome, pessoa.Idade, pessoa.Cidade)
}
```

🔹 **Explicação:**
- `json.NewDecoder(file)` lê e converte o JSON diretamente para a struct.
- O uso de `&pessoa` permite decodificar os dados corretamente.

---

## 🚀 Dicas para Trabalhar com Grandes Volumes de Dados

🔹 **Usar Buffers:** Prefira `bufio.Writer` e `bufio.Reader` ao processar arquivos grandes.
🔹 **Processamento em Stream:** Evite carregar tudo na memória. Leia e processe linha por linha.
🔹 **Erros e Validações:** Sempre trate erros ao trabalhar com arquivos externos.

---

## 🎯 Conclusão

Agora você aprendeu:
✅ Como ler e escrever arquivos CSV.
✅ Como lidar com JSON no Go.
✅ Estratégias para otimizar a leitura de arquivos grandes.
✅ Como evitar erros e melhorar a eficiência.

Agora que dominamos CSV e JSON, estamos prontos para avançar para **streaming de dados e manipulação de fluxos na próxima seção!** 🚀

