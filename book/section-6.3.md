# **6.3 Structs e Métodos**

Os **structs** são a forma como Go define **tipos compostos**, permitindo armazenar múltiplos campos sob um mesmo identificador. Além disso, Go permite associar **métodos** a structs, possibilitando a criação de comportamentos encapsulados sem a necessidade de classes.

Nesta seção, exploraremos:

- Como declarar e inicializar structs
- Acessando e modificando campos
- Métodos associados a structs
- Ponteiros e métodos mutáveis
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

## **6.3.2 Acessando e Modificando Campos**

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

## **6.3.3 Métodos Associados a Structs**

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

## **6.3.4 Comparação de Structs**

Em Go, structs podem ser comparados **se todos os seus campos forem comparáveis**:

```go
p1 := Pessoa{"Alice", 30}
p2 := Pessoa{"Alice", 30}

fmt.Println(p1 == p2) // true (valores iguais)
```

📌 **Se o struct contém slices ou maps, a comparação direta não é possível.**  
Nesses casos, podemos usar `reflect.DeepEqual()`:

```go
import "reflect"

p1 := struct {
    Nomes []string
}{Nomes: []string{"Alice"}}

p2 := struct {
    Nomes []string
}{Nomes: []string{"Alice"}}

fmt.Println(reflect.DeepEqual(p1, p2)) // true
```

---

## **6.3.5 Embedding: Simulando Herança**

Go não tem **herança**, mas permite **composição via embedding**:

```go
type Animal struct {
    Nome string
}

type Cachorro struct {
    Animal
    Raca string
}

dog := Cachorro{Animal: Animal{Nome: "Rex"}, Raca: "Labrador"}
fmt.Println(dog.Nome) // "Rex"
```

📌 **Isso simula herança, mas sem a complexidade de classes.**

---

## **6.3.6 Boas Práticas no Uso de Structs**

✔ **Prefira nomear os campos explicitamente na inicialização.**  
✔ **Use ponteiros para métodos que modificam structs.**  
✔ **Evite structs grandes sendo passados por valor.**  
✔ **Use embedding para reuso de código em vez de herança tradicional.**  

---

## **Conclusão**

Os **structs e métodos** são fundamentais para modelar dados e encapsular comportamento em Go.  
No próximo capítulo, veremos como lidar com **campos opcionais e a tag `omitempty`**, permitindo manipular dados de forma mais flexível! 🚀
