---

## 6.3 Structs e Métodos {#6.3-structs-e-métodos}

# **6.3 Structs e Métodos**

Os **structs** são a forma como Go define **tipos compostos**, permitindo armazenar múltiplos campos sob um mesmo identificador. Além disso, Go permite associar **métodos** a structs, possibilitando a criação de comportamentos encapsulados sem a necessidade de classes.

Nesta seção, exploraremos:

- Como declarar e inicializar structs
- Acessando e modificando campos
- Métodos associados a structs
- Ponteiros e métodos mutáveis
- Structs anônimos
- Structs mutáveis vs. imutáveis
- Manipulação avançada de JSON
- Interface `Stringer`
- Uso de tags customizadas
- Boas práticas no uso de structs

---

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

## **Conclusão**

Os **structs e métodos** são fundamentais para modelar dados e encapsular comportamento em Go. Agora, com tópicos mais avançados como structs anônimos, mutáveis vs. imutáveis, `Stringer`, manipulação de JSON e tags customizadas, você tem uma visão completa!

No próximo capítulo, veremos como lidar com **campos opcionais e a tag `omitempty`**, permitindo manipular dados de forma mais flexível! 🚀

