# **9.2 Implementação de Múltiplas Interfaces**

Go não suporta **herança múltipla**, mas permite que um tipo implemente **múltiplas interfaces** simultaneamente. Isso torna a linguagem mais flexível e evita problemas comuns da herança tradicional.

Nesta seção, exploraremos:

- Como um struct pode implementar várias interfaces
- Benefícios da implementação implícita
- Composição de interfaces para reutilização de código
- Melhores práticas ao trabalhar com múltiplas interfaces

---

## **9.2.1 Como um Struct Implementa Múltiplas Interfaces**

Diferente de linguagens como Java e C++, onde precisamos declarar explicitamente quais interfaces uma classe implementa, **Go usa implementação implícita**:

```go
type Falante interface {
    Falar() string
}

type Trabalhador interface {
    Trabalhar()
}

type Pessoa struct {
    Nome string
}

func (p Pessoa) Falar() string {
    return "Olá, meu nome é " + p.Nome
}

func (p Pessoa) Trabalhar() {
    fmt.Println(p.Nome, "está trabalhando")
}

p := Pessoa{"Alice"}
var f Falante = p
var t Trabalhador = p

fmt.Println(f.Falar()) // "Olá, meu nome é Alice"
t.Trabalhar()          // "Alice está trabalhando"
```

📌 **`Pessoa` implementa `Falante` e `Trabalhador` automaticamente, sem precisar declarar.**

✅ **Isso reduz o acoplamento e melhora a flexibilidade.**

---

## **9.2.2 Criando Interfaces Compostas**

Podemos combinar várias interfaces em uma única, criando **interfaces compostas**:

```go
type SerHumano interface {
    Falante
    Trabalhador
}

var sh SerHumano = p
sh.Trabalhar() // "Alice está trabalhando"
fmt.Println(sh.Falar()) // "Olá, meu nome é Alice"
```

📌 **Isso permite agrupar funcionalidades comuns sem criar dependências desnecessárias.**

✅ **Uso prático em bibliotecas e APIs:**

```go
type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(dados string)
}

type Arquivo interface {
    Leitor
    Escritor
}
```

📌 **Isso permite que qualquer tipo que implemente `Ler()` e `Escrever()` seja tratado como um `Arquivo`.**

---

## **9.2.3 Interfaces e Ponteiros**

Quando usamos um **struct por valor**, apenas métodos com **value receiver** são chamados:

```go
func (p Pessoa) Falar() string {
    return "Olá, sou " + p.Nome
}

var f Falante = Pessoa{"Bob"}
fmt.Println(f.Falar()) // "Olá, sou Bob"
```

📌 **Se um método usa `pointer receiver`, precisamos passar um ponteiro para a interface.**

```go
func (p *Pessoa) Trabalhar() {
    fmt.Println(p.Nome, "trabalhando duro!")
}

var t Trabalhador = &Pessoa{"Bob"} // Agora funciona!
t.Trabalhar()
```

✅ **Isso evita cópias desnecessárias e permite modificar o estado do struct.**

---

## **9.2.4 Comparação com Outras Linguagens**

| Característica | Go | Java | C++ | Python |
|---------------|----|------|-----|--------|
| Herança Múltipla | ❌ (Usa interfaces) | ✅ | ✅ | ✅ |
| Implementação Implícita | ✅ | ❌ | ❌ | ✅ |
| Composição de Interfaces | ✅ | ✅ | ✅ | ✅ |
| Interface Segura por Design | ✅ | ❌ | ❌ | ✅ |

📌 **Go evita os problemas de herança múltipla ao permitir que structs implementem múltiplas interfaces de forma independente.**

---

## **9.2.5 Boas Práticas**

✔ **Use interfaces pequenas e focadas em um único propósito.**  
✔ **Prefira composição em vez de herança tradicional.**  
✔ **Evite definir interfaces desnecessárias — implemente-as apenas quando precisar.**  
✔ **Use ponteiros para modificar o estado do struct dentro da interface.**  

---

## **Conclusão**

A implementação de **múltiplas interfaces** em Go permite criar código flexível e desacoplado, sem os problemas da herança múltipla.  
No próximo capítulo, exploraremos **métodos em embeddings**, aprofundando como Go lida com a reutilização de código! 🚀
