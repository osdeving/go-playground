# 🛠 **1.4 Instalação e Configuração do Ambiente**

Antes de começar a programar em **Go**, é necessário configurar o ambiente corretamente. Esta seção aborda os passos para **instalar** o Go em diferentes sistemas operacionais, **verificar a instalação** e **configurar variáveis de ambiente**.

---

## 📥 **1.4.1 Instalando o Go**

A instalação do Go pode ser realizada de diferentes formas, dependendo do sistema operacional. A maneira recomendada é utilizar os binários oficiais fornecidos pelo [site oficial do Go](https://go.dev/dl/).

### 🖥 **Windows**
1. Acesse [https://go.dev/dl/](https://go.dev/dl/).
2. Baixe o instalador `.msi` correspondente à sua arquitetura (**x86** ou **x64**).
3. Execute o instalador e siga as instruções na tela.
4. Após a instalação, abra o **Prompt de Comando (cmd)** e digite:
   ```sh
   go version
   ```
   Isso deve exibir a versão instalada do Go.

### 🐧 **Linux**
1. Baixe o binário mais recente para Linux:
   ```sh
   wget https://go.dev/dl/go1.x.x.linux-amd64.tar.gz
   ```
2. Extraia o arquivo para `/usr/local`:
   ```sh
   sudo tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
   ```
3. Adicione o Go ao **PATH** (adicione estas linhas ao `~/.bashrc` ou `~/.zshrc`):
   ```sh
   export PATH=$PATH:/usr/local/go/bin
   ```
4. Verifique a instalação:
   ```sh
   go version
   ```

### 🍏 **macOS**
1. Baixe o pacote `.pkg` da [página oficial](https://go.dev/dl/).
2. Execute o instalador e siga as instruções.
3. Para instalar via Homebrew:
   ```sh
   brew install go
   ```
4. Verifique a instalação:
   ```sh
   go version
   ```

---

## ⚙️ **1.4.2 Configuração do Ambiente**

Após instalar o Go, é necessário configurar corretamente as **variáveis de ambiente**.

### 🌍 **GOPATH e GOROOT**
- **GOROOT**: Aponta para o diretório de instalação do Go (**configurado automaticamente**).
- **GOPATH**: Define o local onde ficarão os projetos Go.

✅ **Adicione ao `.bashrc`, `.zshrc` ou `.profile`:**
```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

🔍 **Verifique se as variáveis foram configuradas corretamente:**
```sh
echo $GOPATH
```

---

## 🚀 **1.4.3 Testando a Instalação**

Para garantir que tudo esteja configurado corretamente, crie um pequeno programa Go:

1. Crie um diretório para seu projeto:
   ```sh
   mkdir -p $GOPATH/src/hello
   cd $GOPATH/src/hello
   ```
2. Crie um arquivo `main.go` com o seguinte conteúdo:
   ```go
   package main
   import "fmt"

   func main() {
       fmt.Println("Hello, Go!")
   }
   ```
3. Compile e execute:
   ```sh
   go run main.go
   ```
   Se a instalação estiver correta, você verá a saída:
   ```
   Hello, Go!
   ```

---

## 🔄 **1.4.4 Mantendo o Go Atualizado**

Sempre que possível, mantenha sua instalação do **Go atualizada** para garantir o suporte a novos recursos e correções de segurança. Para atualizar:

### 🖥 **Windows**
Baixe e execute a versão mais recente do instalador `.msi`.

### 🐧 🍏 **Linux e macOS**
1. Remova a versão antiga:
   ```sh
   sudo rm -rf /usr/local/go
   ```
2. Instale a nova versão seguindo os passos anteriores.

🔍 **Verifique a versão instalada após a atualização:**
```sh
go version
```

---

## 🎯 **Conclusão**

Com o **Go instalado e configurado**, você já pode começar a desenvolver aplicações. No próximo capítulo, veremos a **estrutura básica de um programa Go** e seus principais componentes. 🚀

