
## **6.3.1 Declarando e Inicializando Structs**

A sintaxe para definir um struct é:

```bnf
StructType     = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl      = (IdentifierList Type | EmbeddedField) [ Tag ] .
IdentifierList = identifier { "," identifier } .
EmbeddedField  = [ "*" ] TypeName .
Tag            = string_lit .
```

💬 Em outras palavras, primeiro usamos a palavra-chave type seguido do nome do struct e seus campos entre chaves `{}`. Cada campo é definido por um nome e um tipo. 

Por exemplo:

```go
type NomeDaStruct struct {
    NomeDoCampo1 TipoDoCampo
    NomeDoCampo2 TipoDoCampo
    ...
    NomeDoCampoN TipoDoCampo
}
```
➡️ Podemos adicionar Tags para os campos de um struct, que são metadados usados para serialização e outras operações. Por exemplo:

```go
type Pessoa struct {
    Nome     string `json:"nome"`
    Idade    int    `json:"idade"`
    Endereco string `json:"endereco,omitempty"`
    Telefone string `json:"telefone,omitempty"`
}
```

As Tags seguem o padrão:

```go
`tag1:"value1" tag2:"value2"`
```

📌 **Os valores precisam estar entre aspas duplas e a string precisa estar entre crases. Tags sem valor são permitidas.**

```go
`tag1:"value1" tag2:"value2" tag3 tag4`
```


📌 **A tag `omitempty` faz com que o campo seja omitido na serialização JSON se estiver vazio.**


Dada a struct Pessoa acima, podemos inicializá-la de várias formas:

```go
// 1. Inicialização explícita
var p1 Pessoa
p1.Nome = "Alice"
p1.Idade = 30

// 2. Inicialização direta
p2 := Pessoa{"Bob", 25} 

// 3. Usando chave-valor (melhor prática)
p3 := Pessoa{Nome: "Carlos", Idade: 40}
```

📌 Os valores não atribuídos são **inicializados com zero** (0 para int, "" para string, nil para ponteiros, etc.).


📌 ** Uma vez que os valores são atribuídos na ordem dos campos, é fácil cometer erros. Assim, o uso de (`{Nome: "Carlos"}`) evita erros caso a ordem dos campos mude no futuro porque informamos o nome do campo explicitamente. **

Podemos iniciar uma struct usando o operador `new`:

```go
p4 := new(Pessoa)
p4.Nome = "Daniel"
p4.Idade = 35
```

A função built-in `new` aloca memória para o struct e retorna um ponteiro para ele. É equivalente a:

```go
p4 := &Pessoa{}
p4.Nome = "Daniel"
p4.Idade = 35
```

📌 **O uso de `new` é menos comum em Go, pois a inicialização direta é mais idiomática.**


Podemos definir valores padrão para os campos de um struct usando uma função construtora:

```go
type Config struct {
    Host string
    Port int
}

func NewConfig() Config {
    return Config{
        Host: "localhost",
        Port: 8080,
    }
}

cfg := NewConfig()
fmt.Println(cfg.Host) // "localhost"
fmt.Println(cfg.Port) // 8080
```

📌 **Isso garante que os structs sejam inicializados com valores sensíveis por padrão.**

Podemos usar funções auxiliares para inicializar structs complexas e encapsular lógica na criação:

```go
type DatabaseConfig struct {
    Username string
    Password string
    Database string
}

func NewDatabaseConfig(username, password, database string) DatabaseConfig {
    return DatabaseConfig{
        Username: username,
        Password: password,
        Database: database,
    }
}

dbConfig := NewDatabaseConfig("user", "pass", "mydb")
fmt.Println(dbConfig)
```

📌 **Funções auxiliares tornam o código mais legível e fácil de manter.**

Além de funções auxiliares normais, podemos usar funções Variádicas e simular o padrão Builder para inicializações altamente configuráveis:

```go
type Option func(*ServerConfig)

func WithAddress(address string) Option {
    return func(cfg *ServerConfig) {
        cfg.Address = address
    }
}

func WithPort(port int) Option {
    return func(cfg *ServerConfig) {
        cfg.Port = port
    }
}

func NewServerConfig(options ...Option) ServerConfig {
    cfg := ServerConfig{
        Address: "localhost",
        Port:    80,
    }
    for _, opt := range options {
        opt(&cfg)
    }
    return cfg
}

config := NewServerConfig(WithAddress("192.168.1.1"), WithPort(8080))
fmt.Println(config)
```


## **6.3.2 Structs Anônimas**

Go permite a criação de **structs anônimas**, que são úteis para declarações inline:

```go
p := struct {
    Nome  string
    Idade int
}{Nome: "Alice", Idade: 30}

fmt.Println(p.Nome) // "Alice"
```

💡 **Vamos usar geralmente nas seguites situações:**  
- Em **testes rápidos**, para não precisar criar um `type`.  
- Em **objetos temporários** que não vamos reutilizar.  

---

## **6.3.3 Acessando e Modificando Campos**

Os campos de uma struct podem ser acessados diretamente:

```go
fmt.Println(p1.Nome) // "Alice"

p1.Idade = 31 // Alterando um valor
```

✅ **as structs em Go são copiadas por valor.**  
Isso significa que ao atribuir uma struct a uma nova variável, uma cópia será feita:

```go
p4 := p1
p4.Idade = 50

fmt.Println(p1.Idade) // 31 (original não foi alterado)
fmt.Println(p4.Idade) // 50 (cópia modificada)
```

📌 **Se quisermos modificar a struct original, devemos usar ponteiros (`*`). _Veremos mais sobre ponteiros em capítulos futuros_**

---

## **6.3.4 Structs Mutáveis vs. Imutáveis**

Go **não tem um sistema nativo de imutabilidade**, mas podemos simular com **campos privados** e métodos getters:

```go
type Config struct {
    timeout int
}

func NewConfig(timeout int) Config {
    return Config{timeout: timeout}
}

func (c Config) Timeout() int {
    return c.timeout
}
```

📌 **A struct `Config` é imutável, pois não há setter público. Note que o campo **timeout** é privado uma vez que inicia com letra minúscula.**

```go
cfg := NewConfig(30)
fmt.Println(cfg.Timeout()) // 30

// cfg.timeout = 40 // Erro: timeout é privado
```

📌 **Nesse exemplo a imutabilidade garante que os valores de configuração permaneçam consistentes.**

---

## **6.3.5 Métodos Associados a Structs**

Conforme já vimos na seção anterior, podemos associar **métodos** a structs usando `func` com um **receiver**:

```go
func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome
}

p := Pessoa{"Alice", 30}
fmt.Println(p.Saudacao()) // "Olá, meu nome é Alice"
```

📌 **Os métodos não podem modificar o struct diretamente, pois ele é passado por valor!**

✅ **Para modificar o struct, devemos usar um ponteiro no receiver:**

```go
func (p *Pessoa) Envelhecer() {
    p.Idade++
}

p.Envelhecer()
fmt.Println(p.Idade) // 31
```

📌 **Com `*Pessoa`, o método pode alterar os campos diretamente.**

💭Os receivers vêm logo após a palavra-chave `func` e antes do nome do método e ficam entre parênteses. O nome do receiver e o tipo são separados por um espaço e o nome pode ser qualquer identificador válido.


---

## **6.3.6 Structs e JSON: Manipulação Avançada**

Já vimos que os campos podem ter Tags. Além de `omitempty`, podemos usar `json.RawMessage` para armazenar JSON dinâmico:

```go
type Response struct {
    Data json.RawMessage `json:"data"`
}
```

📌 **Isso permite armazenar JSON de diferentes estruturas sem um tipo fixo.**

Podemos usar `json.Marshal` para serializar um struct em JSON:

```go
p := Pessoa
jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData))
```

📌 **`json.Marshal` retorna um slice de bytes, que pode ser convertido em uma string para exibição.**

Para desserializar JSON de volta para um struct, usamos `json.Unmarshal`:

```go
var p2 Pessoa
json.Unmarshal(jsonData, &p2)
fmt.Println(p2)
```

📌 **`json.Unmarshal` modifica o struct passado como ponteiro.**

💡 *json está no pacote .**encoding/json**.  
**json.RawMessage** é um tipo de dados que armazena JSON bruto.  
**json.Marshal** e **json.Unmarshal** são usados para serializar e desserializar JSON.*

Outros usos comuns de Tags incluem **validação de entrada**, **formatação de saída** e **mapeamento de campos**. Por exemplo:

```go
type Pessoa struct {
    Nome  string `json:"name" validate:"required"`
    Idade int    `json:"age" validate:"gte=0,lte=130"`
}
```

Essas Tags são lidas por pacotes de terceiros para **validação de entrada** e **serialização/desserialização JSON**. Para ler as Tags usamos `reflect`:

```go
t := reflect.TypeOf(Pessoa{})
field, _ := t.FieldByName("Nome")
fmt.Println(field.Tag.Get("json")) // "name"
fmt.Println(field.Tag.Get("validate")) // "required"
```

📌 **`reflect` é um pacote poderoso para inspecionar structs e acessar seus metadados.**

---

## **6.3.7 Interface `Stringer` para Representação Personalizada**

Stringer são interfaces que definem um método `String()` que retorna uma representação textual do objeto.

Por exemplo, podemos definir uma **representação textual customizada** para structs implementando `fmt.Stringer`:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func (p Pessoa) String() string {
    return fmt.Sprintf("Pessoa: %s, Idade: %d", p.Nome, p.Idade)
}

p := Pessoa{"Alice", 30}
fmt.Println(p) // "Pessoa: Alice, Idade: 30"
```

💡 **Quando usar?**  
- Para **depuração** e **logs**.  
- Para fornecer uma **saída amigável para o usuário**.  

---

Veremos mais sobre interfaces e métodos em capítulos futuros, mas por enquanto, você já deve ter uma boa compreensão de como usar structs e métodos em Go! 🎉


## **6.3.8 Structs e Tags Customizadas**

Além de `json`, podemos definir **tags customizadas** para parsear structs de diferentes formas:

```go
type Config struct {
    Host string `env:"APP_HOST"`
    Port int    `env:"APP_PORT"`
}
```

Isso permite criar pacotes que parseiam configurações de ambiente automaticamente. Por exemplo:

```go
func ParseConfig(cfg interface{}) {
    t := reflect.TypeOf(cfg).Elem()
    v := reflect.ValueOf(cfg).Elem()

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("env")
        value := os.Getenv(tag)
        if value != "" {
            v.Field(i).SetString(value)
        }
    }
}

cfg := &Config{}
ParseConfig(cfg)
fmt.Println(cfg.Host, cfg.Port)
```

O código acima lê as variáveis de ambiente e as atribui aos campos correspondentes do struct `Config` com base nas tags `env`.

📌 **Essa técnica permite criar pacotes que parseiam configurações de ambiente automaticamente.**  

📌 **`reflect` é um pacote poderoso que permite inspecionar structs dinamicamente.**  

📌 **Tags customizadas são amplamente usadas para serialização e validação de dados.**

---

## **Pratique Go**

🎯 Agora que você aprendeu sobre Structs e Métodos, tente os seguintes desafios:

🧐 **Desafios Avançados:**

<details>
  <summary>1⃣ Crie uma struct imutável em Go e implemente um construtor seguro.</summary>
  
  ```go
  package main
  import "fmt"

  type Config struct {
      timeout int
  }

  func NewConfig(timeout int) Config {
      return Config{timeout: timeout}
  }

  func (c Config) Timeout() int {
      return c.timeout
  }

  func main() {
      cfg := NewConfig(30)
      fmt.Println("Timeout:", cfg.Timeout())
  }
  ```
  
</details>

<details>
  <summary>2⃣ Implemente um método em um struct que retorne um ponteiro para ele mesmo e permita chamadas encadeadas (method chaining).</summary>
  
  ```go
  package main
  import "fmt"

  type Pessoa struct {
      Nome string
      Idade int
  }

  func (p *Pessoa) SetNome(nome string) *Pessoa {
      p.Nome = nome
      return p
  }

  func (p *Pessoa) SetIdade(idade int) *Pessoa {
      p.Idade = idade
      return p
  }

  func main() {
      p := &Pessoa{}
      p.SetNome("Alice").SetIdade(30)
      fmt.Println(p)
  }
  ```
  
</details>

<details>
  <summary>3⃣ Crie uma struct que implemente a interface `Stringer` e exiba uma representação personalizada do objeto.</summary>
  
  ```go
  package main
  import "fmt"

  type Produto struct {
      Nome  string
      Preco float64
  }

  func (p Produto) String() string {
      return fmt.Sprintf("Produto: %s, Preco: R$%.2f", p.Nome, p.Preco)
  }

  func main() {
      p := Produto{"Notebook", 3599.90}
      fmt.Println(p)
  }
  ```
  
</details>

<details>
  <summary>4⃣ Utilize `json.Marshal` para serializar um struct e `json.Unmarshal` para desserializá-lo de volta.</summary>
  
  ```go
  package main
  import (
      "encoding/json"
      "fmt"
  )

  type Pessoa struct {
      Nome  string `json:"nome"`
      Idade int    `json:"idade"`
  }

  func main() {
      p := Pessoa{"Alice", 30}
      jsonData, _ := json.Marshal(p)
      fmt.Println(string(jsonData))

      var p2 Pessoa
      json.Unmarshal(jsonData, &p2)
      fmt.Println(p2)
  }
  ```
  
</details>

<details>
  <summary>5⃣ Crie uma struct com um campo `sync.Mutex` e implemente um método seguro para concorrência.</summary>
  
  ```go
  package main
  import (
      "fmt"
      "sync"
  )

  type Contador struct {
      mu sync.Mutex
      valor int
  }

  func (c *Contador) Incrementar() {
      c.mu.Lock()
      defer c.mu.Unlock()
      c.valor++
  }

  func main() {
      c := Contador{}
      c.Incrementar()
      fmt.Println("Valor:", c.valor)
  }
  ```
  
</details>

<details>
  <summary>6⃣ Implemente um struct que use `sync.Once` para garantir que um método seja executado apenas uma vez.</summary>
  
  ```go
  package main
  import (
      "fmt"
      "sync"
  )

  type Singleton struct {
      once sync.Once
  }

  func (s *Singleton) Executar() {
      s.once.Do(func() {
          fmt.Println("Executando apenas uma vez")
      })
  }

  func main() {
      s := &Singleton{}
      s.Executar()
      s.Executar()
  }
  ```
  
</details>

<details>
  <summary>7⃣ Defina uma struct aninhada (struct dentro de struct) e acesse os campos internos.</summary>
  
  ```go
  package main
  import "fmt"

  type Endereco struct {
      Rua  string
      Cidade string
  }

  type Pessoa struct {
      Nome     string
      Endereco Endereco
  }

  func main() {
      p := Pessoa{
          Nome: "Alice",
          Endereco: Endereco{
              Rua: "Rua das Flores",
              Cidade: "São Paulo",
          },
      }
      fmt.Println(p.Nome, "mora em", p.Endereco.Cidade)
  }
  ```
  
</details>

<details>
  <summary>8⃣ Crie uma struct que implemente múltiplas interfaces.</summary>
  
  ```go
  package main
  import "fmt"

  type Animal interface {
      EmitirSom()
  }

  type Movel interface {
      Mover()
  }

  type Cachorro struct {}

  func (c Cachorro) EmitirSom() {
      fmt.Println("Au au")
  }

  func (c Cachorro) Mover() {
      fmt.Println("Correndo...")
  }

  func main() {
      var c Cachorro
      c.EmitirSom()
      c.Mover()
  }
  ```
  
</details>

<details>
  <summary>9⃣ Escreva um método que retorne uma interface vazia e faça type assertion.</summary>
  
  ```go
  package main
  import "fmt"

  func retornaAlgo() interface{} {
      return "Texto"
  }

  func main() {
      valor := retornaAlgo()
      if str, ok := valor.(string); ok {
          fmt.Println("String recebida:", str)
      }
  }
  ```
  
</details>

<details>
  <summary>🏢 10⃣ Utilize `reflect` para inspecionar um struct dinamicamente.</summary>
  
  ```go
  package main
  import (
      "fmt"
      "reflect"
  )

  type Pessoa struct {
      Nome  string
      Idade int
  }

  func main() {
      p := Pessoa{"Alice", 30}
      t := reflect.TypeOf(p)
      fmt.Println("Nome do tipo:", t.Name())
      for i := 0; i < t.NumField(); i++ {
          field := t.Field(i)
          fmt.Printf("Campo: %s, Tipo: %s\n", field.Name, field.Type)
      }
  }
  ```
  
</details>

---

---

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>1⃣ Qual a diferença entre um método com receiver por valor e um método com receiver por ponteiro?</summary>
  Um método com receiver por valor trabalha com uma cópia do struct, sem modificar o original. Um método com receiver por ponteiro permite modificar os campos do struct original.
</details>

<details>
  <summary>2⃣ Como Go trata herança e como podemos simular esse comportamento?</summary>
  Go não suporta herança, mas permite a reutilização de código através de **embedding**, onde um struct pode conter outro struct e acessar seus métodos diretamente.
</details>

<details>
  <summary>3⃣ O que acontece ao comparar structs diretamente?</summary>
  Dois structs podem ser comparados diretamente se todos os seus campos forem comparáveis. Se contiverem slices, maps ou funções, a comparação direta resultará em erro de compilação.
</details>

<details>
  <summary>4⃣ Como evitar cópias desnecessárias ao passar structs para funções?</summary>
  Para evitar cópias, passe um **ponteiro para o struct** em vez do struct por valor.
</details>

<details>
  <summary>5⃣ O que acontece se usarmos um método com receiver por valor em um ponteiro?</summary>
  O compilador automaticamente **desreferencia** o ponteiro e chama o método normalmente.
</details>

<details>
  <summary>6⃣ Como representar campos opcionais dentro de um struct?</summary>
  Usando **ponteiros** para os campos opcionais ou a tag `omitempty` para ignorar campos vazios na serialização JSON.
</details>

<details>
  <summary>7⃣ Como podemos garantir que um struct seja imutável?</summary>
  Declarando os campos como **privados** (letra minúscula) e fornecendo apenas métodos de leitura.
</details>

<details>
  <summary>8⃣ Como Go trata a inicialização padrão de structs?</summary>
  Se um struct for declarado sem inicialização explícita, seus campos assumem os **valores zero** de seus respectivos tipos.
</details>

<details>
  <summary>9⃣ Qual é a vantagem de implementar a interface `Stringer` para um struct?</summary>
  A interface `Stringer` permite definir um **método de formatação personalizado** quando o struct for impresso, tornando a saída mais legível.
</details>

<details>
  <summary>💼 10⃣ Como podemos usar `reflect` para inspecionar um struct dinamicamente?</summary>
  Utilizando `reflect.TypeOf()` para obter os metadados do struct e `reflect.ValueOf()` para acessar seus valores.
</details>

---




## **Conclusão**

🎉 **Parabéns!** 🎉

Agora você sabe que: 

- **Structs** são tipos de dados compostos que armazenam campos nomeados.
- **Tags** são metadados associados a campos de struct.
- **Structs anônimas** são úteis para declarações inline.
- **Métodos** são funções associadas a structs e usam receivers para acessar campos.
- **JSON** é um formato comum para serialização de dados. `json.Marshal` e `json.Unmarshal` são usados para converter structs em JSON e vice-versa.
- **Tags customizadas** são amplamente usadas para serialização e validação de dados.
- A interface `Stringer` permite definir uma representação textual personalizada para um objeto.

🚀 E você está pronto para usar structs e métodos em Go! 🎯

---

🕵️ **Para saber mais:**
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: JSON](https://gobyexample.com/json)
- [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
- [The Go Blog: JSON and Go](https://blog.golang.org/json-and-go)
- [The Go Blog: Method Sets](https://blog.golang.org/method-sets)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)
- [The Go Blog: Custom JSON Marshalling](https://blog.golang.org/json-and-go)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)
- [The Go Blog: Advanced JSON Handling](https://blog.golang.org/json)
- [The Go Blog: Stringer](https://blog.golang.org/laws-of-reflection#TOC_7.)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)

---