# **8.1 Métodos Associados a Structs**

Em Go, métodos são funções associadas a **structs**, permitindo encapsular comportamento dentro de um tipo.  
Embora Go não tenha **classes** como em linguagens orientadas a objetos tradicionais, **métodos** e **interfaces** fornecem uma abordagem equivalente.

Nesta seção, exploraremos:

- Como declarar e usar métodos em Go
- Diferença entre métodos com **value receiver** e **pointer receiver**
- Encapsulamento e boas práticas no uso de métodos

---

## **8.1.1 O Que São Métodos em Go?**

Um **método** em Go é uma função associada a um tipo **struct**:

```go
type Pessoa struct {
    Nome string
}

// Método associado ao struct Pessoa
func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome
}

p := Pessoa{"Alice"}
fmt.Println(p.Saudacao()) // "Olá, meu nome é Alice"
```

📌 **O método `Saudacao()` pertence ao struct `Pessoa` e pode ser chamado em qualquer instância.**

✅ **O que diferencia um método de uma função normal?**  
- Um método recebe um **receiver**, que representa a instância do struct.  
- Isso permite associar comportamento diretamente a um tipo.

---

## **8.1.2 Métodos com Value Receiver vs. Pointer Receiver**

Os métodos podem receber **cópias da struct** (**value receiver**) ou um **ponteiro para a struct** (**pointer receiver**).  

📌 **Value Receiver:** O método recebe **uma cópia** da struct, sem modificar o original.

```go
func (p Pessoa) Saudacao() string {
    return "Olá, meu nome é " + p.Nome
}
```

✅ **Pointer Receiver:** O método recebe um **ponteiro** para a struct, permitindo modificar o valor original.

```go
func (p *Pessoa) AlterarNome(novoNome string) {
    p.Nome = novoNome
}

p := Pessoa{"Alice"}
p.AlterarNome("Bob")

fmt.Println(p.Nome) // "Bob" (alterado!)
```

📌 **Regra geral:**  
- **Use `value receiver`** se o método não precisa modificar a struct.  
- **Use `pointer receiver`** se o método modifica o estado da struct ou se a struct for grande (evita cópias desnecessárias).

---

## **8.1.3 Métodos vs. Funções Normais**

Podemos definir funções normais que operam em structs:

```go
func saudacao(p Pessoa) string {
    return "Olá, " + p.Nome
}
```

Mas a abordagem **com métodos** é mais idiomática e mantém a lógica organizada:

```go
func (p Pessoa) Saudacao() string {
    return "Olá, " + p.Nome
}
```

📌 **Métodos fazem sentido quando o comportamento pertence ao struct e não a outro conceito.**

---

## **8.1.4 Encapsulamento e Visibilidade**

Go não possui modificadores de acesso (`private`, `public`), mas usa **convenções de capitalização**:

- **Começando com maiúscula (`Exportado`)** → Público (acessível fora do pacote).
- **Começando com minúscula (`interno`)** → Privado ao pacote.

```go
type ContaBancaria struct {
    saldo float64 // Privado ao pacote
}

func (c *ContaBancaria) Depositar(valor float64) {
    c.saldo += valor // Método pode acessar campo privado
}

func (c *ContaBancaria) Saldo() float64 {
    return c.saldo
}
```

✅ **Mesmo sendo privado, o saldo pode ser acessado indiretamente via método público `Saldo()`.**

---

## **8.1.5 Métodos em Structs Embutidos**

Go permite **reutilizar métodos via composição** (embedding).

```go
type Animal struct {
    Nome string
}

func (a Animal) Falar() string {
    return "O animal faz um som"
}

type Cachorro struct {
    Animal // Embedded struct
}

dog := Cachorro{Animal{"Rex"}}
fmt.Println(dog.Falar()) // "O animal faz um som"
```

📌 **Os métodos da struct embutida (`Animal`) são herdados pelo struct que a contém (`Cachorro`).**

---

## **8.1.6 Comparação com Outras Linguagens**

| Recurso | Go | Java | Python | C++ |
|---------|----|------|--------|-----|
| Métodos em Structs | ✅ | ✅ (Classes) | ✅ (Classes) | ✅ |
| Encapsulamento via visibilidade | ✅ (por convenção) | ✅ (`private`, `public`) | ✅ (`_nome`) | ✅ |
| Pointer Receiver (`*`) | ✅ | ❌ | ❌ | ✅ |
| Herança | ❌ (Usa composição) | ✅ | ✅ | ✅ |

📌 **Diferente de Java e Python, Go usa **composição** em vez de herança para reutilizar código.**

---

## **8.1.7 Boas Práticas**

✔ **Use métodos quando o comportamento estiver ligado a um struct.**  
✔ **Use `pointer receiver` (`*T`) para modificar o struct e evitar cópias desnecessárias.**  
✔ **Prefira composição (embedding) em vez de herança para reuso de código.**  
✔ **Evite métodos muito grandes — divida lógica complexa em funções auxiliares.**  

---

## **Conclusão**

Os **métodos em structs** permitem encapsular comportamento de forma organizada, tornando o código mais legível e eficiente.  
No próximo capítulo, exploraremos **value receivers vs. pointer receivers**, entendendo seu impacto na performance! 🚀
