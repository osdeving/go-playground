
## **6.3.1 Declarando e Inicializando Structs**

A sintaxe para definir um struct é:

```go
type Pessoa struct {
    Nome  string
    Idade int
}
```

Podemos inicializar structs de várias formas:

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

📌 **O uso de nomeação explícita (`{Nome: "Carlos"}`) evita erros caso a ordem dos campos mude no futuro.**

---

## **6.3.2 Structs Anônimos**

Go permite a criação de **structs anônimos**, úteis para declarações inline:

```go
p := struct {
    Nome  string
    Idade int
}{Nome: "Alice", Idade: 30}

fmt.Println(p.Nome) // "Alice"
```

💡 **Quando usar?**  
- Para **testes rápidos**, sem precisar criar um `type`.  
- Para **objetos temporários** que não precisam ser reutilizados.  

---

## **6.3.3 Acessando e Modificando Campos**

Os campos de um struct podem ser acessados diretamente:

```go
fmt.Println(p1.Nome) // "Alice"

p1.Idade = 31 // Alterando um valor
```

✅ **Os structs em Go são copiados por valor.**  
Isso significa que ao atribuir um struct a uma nova variável, uma cópia será feita:

```go
p4 := p1
p4.Idade = 50

fmt.Println(p1.Idade) // 31 (original não foi alterado)
fmt.Println(p4.Idade) // 50 (cópia modificada)
```

📌 **Se quisermos modificar o struct original, devemos usar ponteiros (`*`).**

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

📌 **O struct `Config` é imutável, pois não há setter público.**

---

## **6.3.5 Métodos Associados a Structs**

Podemos associar **métodos** a structs usando `func` com um **receiver**:

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

---

## **6.3.6 Structs e JSON: Manipulação Avançada**

Além de `omitempty`, podemos usar `json.RawMessage` para armazenar JSON dinâmico:

```go
type Response struct {
    Data json.RawMessage `json:"data"`
}
```

📌 **Isso permite armazenar JSON de diferentes estruturas sem um tipo fixo.**

---

## **6.3.7 Interface `Stringer` para Representação Personalizada**

Podemos definir uma **representação textual customizada** para structs implementando `fmt.Stringer`:

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

## **6.3.8 Structs e Tags Customizadas**

Além de `json`, podemos definir **tags customizadas** para parsear structs de diferentes formas:

```go
type Config struct {
    Host string `env:"APP_HOST"`
    Port int    `env:"APP_PORT"`
}
```

📌 **Isso permite criar pacotes que parseiam configurações de ambiente automaticamente.**  

---

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

Os **structs e métodos** são fundamentais para modelar dados e encapsular comportamento em Go. Agora, com tópicos mais avançados como structs anônimos, mutáveis vs. imutáveis, `Stringer`, manipulação de JSON e tags customizadas, você tem uma visão completa!

No próximo capítulo, veremos como lidar com **campos opcionais e a tag `omitempty`**, permitindo manipular dados de forma mais flexível! 🚀

