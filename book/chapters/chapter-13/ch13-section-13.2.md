# 📂 13.2 Manipulação de Arquivos em Go

## 👋 Introdução

Agora que já entendemos como o Go lida com entrada e saída de dados por meio das interfaces `io.Reader` e `io.Writer`, é hora de colocar isso em prática manipulando arquivos. Seja para criar logs, armazenar configurações ou processar grandes volumes de dados, trabalhar com arquivos é uma tarefa essencial em muitas aplicações.

Nesta seção, vamos explorar **tudo** sobre arquivos em Go, cobrindo:

- Como **abrir, criar, escrever e ler arquivos** 📜
- Modos de abertura de arquivos 🛠️
- Como **trabalhar com diretórios** 📂
- Como **tratar erros corretamente** ⚠️
- Como **manipular permissões de arquivos** 🔐
- Como **trabalhar com arquivos temporários** 📌
- Como **truncar arquivos e modificar seu tamanho** 🔧
- Como **lidar com concorrência na manipulação de arquivos** ⚡
- Como **copiar arquivos usando `io.Copy`** 📥📤
- Como **trabalhar com `io.Reader` e `io.Writer` diretamente** 🔄
- Como **ler arquivos grandes em chunks** 🗂️
- Como **monitorar arquivos para detectar alterações** 🔍
- Como **trabalhar com arquivos JSON e CSV** 📊
- Como **compactar e descompactar arquivos (`gzip`, `zip`)** 📦
- Como **trabalhar com hard links e symbolic links** 🔗

Bora mergulhar nesse assunto? 🚀

---

## 📖 Criando e Escrevendo em Arquivos

Podemos criar arquivos em Go usando o pacote `os`. Vamos ver um exemplo de como criar e escrever em um arquivo:

```go
file, err := os.Create("exemplo.txt")
if err != nil {
    log.Fatal("Erro ao criar arquivo:", err)
}
defer file.Close()

file.WriteString("Hello, Go!\nEsta é uma linha de exemplo.")
```

### 🛠️ Modos de Abertura de Arquivos

```go
file, err := os.OpenFile("exemplo.txt", os.O_APPEND|os.O_WRONLY, 0644)
```

- `os.O_RDONLY` → Apenas leitura.
- `os.O_WRONLY` → Apenas escrita.
- `os.O_RDWR` → Leitura e escrita.
- `os.O_APPEND` → Acrescenta ao final do arquivo.
- `os.O_CREATE` → Cria o arquivo se ele não existir.
- `os.O_TRUNC` → Apaga o conteúdo se o arquivo já existir.

---

## 📖 Trabalhando com Arquivos Temporários

```go
file, err := os.CreateTemp("./", "tempfile-*.txt")
```

🔹 **Explicação:**
- `os.CreateTemp` cria um arquivo temporário com nome aleatório.
- `defer os.Remove(file.Name())` remove o arquivo ao final do programa.

---

## 📖 Monitoramento de Arquivos (Detectando Alterações)

Podemos monitorar mudanças em arquivos usando a biblioteca `fsnotify`:

```go
watcher, err := fsnotify.NewWatcher()
watcher.Add("arquivo.txt")
for {
    event := <-watcher.Events
    fmt.Println("Arquivo modificado:", event.Name)
}
```

🔹 **Explicação:**
- Detecta modificações em tempo real.
- Útil para sistemas de log e configurações dinâmicas.

---

## 📖 Trabalhando com JSON e CSV

```go
type Pessoa struct {
    Nome string `json:"nome"`
    Idade int    `json:"idade"`
}

json.NewEncoder(file).Encode(pessoa)
```

🔹 **Explicação:**
- `json.Marshal` serializa structs para JSON.
- `json.Unmarshal` lê JSON para structs.
- `encoding/csv` permite manipulação eficiente de arquivos CSV.

---

## 📖 Compactação e Descompactação

```go
file, _ := os.Open("arquivo.txt")
gzipWriter := gzip.NewWriter(destFile)
io.Copy(gzipWriter, file)
gzipWriter.Close()
```

🔹 **Explicação:**
- `gzip.NewWriter` cria um arquivo compactado.
- `zip.Writer` pode ser usado para compactação de múltiplos arquivos.

---

## 📖 Trabalhando com Links Simbólicos e Hard Links

```go
os.Symlink("original.txt", "atalho.txt")
os.Link("original.txt", "copia.txt")
```

🔹 **Explicação:**
- `os.Symlink` cria um link simbólico.
- `os.Link` cria um hard link, que aponta diretamente para os mesmos dados.

---

## 🎯 Conclusão

Agora cobrimos **tudo** sobre arquivos em Go! 🚀
✅ Criação, leitura, escrita.
✅ Modos de abertura e permissões.
✅ Arquivos temporários.
✅ Monitoramento de alterações.
✅ JSON, CSV, compactação e descompactação.
✅ Links simbólicos e hard links.

Com esse conhecimento, você está pronto para lidar com qualquer operação de arquivos em Go! 🔥

