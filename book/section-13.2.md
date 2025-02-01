# **13.2 Leitura e Escrita em CSV e JSON**

Os formatos **CSV** (Comma-Separated Values) e **JSON** (JavaScript Object Notation) são amplamente utilizados para **armazenamento e transferência de dados estruturados**.  
Go oferece suporte nativo para manipulação desses formatos através dos pacotes `encoding/csv` e `encoding/json`.

Nesta seção, exploraremos:

- Como ler e escrever arquivos **CSV** e **JSON** em Go
- Diferenças entre **serialização** e **desserialização**
- Uso de **tags em structs** para personalizar a formatação
- Tratamento de **erros comuns ao manipular dados estruturados**
- Comparação de **desempenho e eficiência**

---

## **13.2.1 Trabalhando com CSV**

O **CSV** é um formato de dados baseado em texto onde cada linha representa um registro e os valores são separados por vírgulas.  

✅ **Exemplo de um arquivo `data.csv`**:

```
id,nome,email
1,Alice,alice@example.com
2,Bob,bob@example.com
3,Charlie,charlie@example.com
```

📌 **Podemos ler e escrever arquivos CSV utilizando o pacote `encoding/csv`.**  

---

### **Lendo Arquivos CSV**

Para ler arquivos CSV, usamos o `csv.Reader`.  
Cada linha do arquivo é convertida em um slice (`[]string`).

✅ **Exemplo: Lendo um arquivo CSV linha por linha**

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Erro ao ler o arquivo CSV:", err)
        return
    }

    for _, row := range records {
        fmt.Println(row) // Cada linha é um slice de strings
    }
}
```

📌 **Esse método carrega todas as linhas na memória, o que pode ser ineficiente para arquivos muito grandes.**  

✅ **Para leitura eficiente linha por linha, use `Read()` em vez de `ReadAll()`.**  

```go
for {
    record, err := reader.Read()
    if err != nil {
        break
    }
    fmt.Println(record)
}
```

📌 **Isso evita carregamento excessivo de memória.**  

---

### **Escrevendo Arquivos CSV**

Para gravar dados em CSV, usamos o `csv.Writer`.  
Cada linha é representada por um slice de strings (`[]string`).

✅ **Exemplo: Criando um novo arquivo CSV**

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("output.csv")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush() // Garante que os dados sejam escritos

    data := [][]string{
        {"id", "nome", "email"},
        {"1", "Alice", "alice@example.com"},
        {"2", "Bob", "bob@example.com"},
    }

    for _, row := range data {
        writer.Write(row)
    }

    fmt.Println("Arquivo CSV gerado com sucesso!")
}
```

📌 **O método `Flush()` força a escrita dos dados no arquivo.**  
📌 **Os dados devem ser passados como slices (`[]string`).**  

---

## **13.2.2 Trabalhando com JSON**

O **JSON** é um formato de dados baseado em chave-valor e é muito utilizado em APIs e aplicações web.  
O Go possui suporte nativo ao JSON através do pacote `encoding/json`.  

✅ **Exemplo de JSON:**

```json
{
    "id": 1,
    "nome": "Alice",
    "email": "alice@example.com"
}
```

📌 **Em Go, o JSON pode ser convertido para structs ou mapas (`map[string]interface{}`).**  

---

### **Lendo Arquivos JSON**

Para ler arquivos JSON, usamos `json.Unmarshal()` para converter os dados em structs.  

✅ **Exemplo: Lendo JSON para uma struct**

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"nome"`
    Email string `json:"email"`
}

func main() {
    file, err := os.ReadFile("data.json")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }

    var user User
    err = json.Unmarshal(file, &user)
    if err != nil {
        fmt.Println("Erro ao converter JSON:", err)
        return
    }

    fmt.Printf("Usuário: %+v\n", user)
}
```

📌 **`json.Unmarshal()` converte JSON em uma struct Go.**  
📌 **As tags `json:"nome"` mapeiam os campos corretamente.**  

✅ **Para JSONs grandes, use `json.Decoder()` para evitar carregar tudo na memória.**  

```go
decoder := json.NewDecoder(file)
decoder.Decode(&user)
```

---

### **Escrevendo Arquivos JSON**

Para salvar dados em JSON, usamos `json.Marshal()`.  
Podemos converter structs diretamente para JSON.  

✅ **Exemplo: Escrevendo JSON em um arquivo**

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"nome"`
    Email string `json:"email"`
}

func main() {
    user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

    file, err := os.Create("output.json")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(user)
    if err != nil {
        fmt.Println("Erro ao escrever JSON:", err)
    }

    fmt.Println("Arquivo JSON salvo com sucesso!")
}
```

📌 **`json.Marshal()` converte structs para JSON.**  
📌 **O `json.NewEncoder()` escreve diretamente no arquivo.**  

---

## **13.2.3 Comparação de Desempenho: CSV vs. JSON**

| Característica | CSV | JSON |
|---------------|-----|------|
| Formato | Estruturado, baseado em colunas | Estruturado, baseado em chave-valor |
| Legibilidade | Média | Alta |
| Tamanho do Arquivo | Pequeno | Pode ser maior |
| Performance | Rápido para leitura | Mais lento que CSV |
| Uso | Planilhas, bancos de dados | APIs, comunicação web |

📌 **Use CSV para grandes volumes de dados tabulares.**  
📌 **Use JSON quando precisar de estrutura hierárquica e comunicação entre sistemas.**  

---

## **Conclusão**

O **Go fornece suporte nativo para manipulação de CSV e JSON**, facilitando a integração de aplicações com bancos de dados, APIs e processamento de dados.  
No próximo capítulo, veremos **como manipular grandes volumes de dados usando `bufio` para otimizar leitura e escrita!** 🚀
