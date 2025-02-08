# **9.3 Métodos em Embeddings**

Em Go, quando usamos **embedding de structs**, os métodos do struct embutido são automaticamente promovidos para o struct que o contém. Isso permite reutilizar funcionalidades sem precisar reescrevê-las, evitando dependências rígidas.

Nesta seção, exploraremos:

- Como os métodos do struct embutido são acessados
- Como sobrescrever métodos herdados
- Como chamar métodos do struct embutido explicitamente
- Diferença entre métodos promovidos e métodos sobrescritos

---

## **9.3.1 Métodos Promovidos pelo Embedding**

Quando um struct embute outro struct, ele herda automaticamente seus métodos:

```go
type Animal struct{ Nome string }

func (a Animal) Falar() string {
    return "O animal faz um som"
}

type Cachorro struct {
    Animal // Embedding
}

dog := Cachorro{Animal{"Rex"}}
fmt.Println(dog.Falar()) // "O animal faz um som"
```

📌 **O método `Falar()` de `Animal` foi promovido para `Cachorro`.**

✅ **Não precisamos redefinir o método em `Cachorro` para usá-lo.**

---

## **9.3.2 Sobrescrevendo Métodos do Struct Embutido**

Podemos sobrescrever um método simplesmente definindo um novo método com o mesmo nome:

```go
type Gato struct {
    Animal
}

func (g Gato) Falar() string {
    return "Miau!"
}

gato := Gato{Animal{"Whiskers"}}
fmt.Println(gato.Falar()) // "Miau!" (método sobrescrito)
```

📌 **O método `Falar()` de `Gato` sobrescreveu o de `Animal`.**  
📌 **Os métodos do struct mais externo têm prioridade.**

✅ **Ainda podemos chamar o método original do struct embutido:**

```go
fmt.Println(gato.Animal.Falar()) // "O animal faz um som"
```

---

## **9.3.3 Chamando Métodos do Struct Embutido**

Mesmo quando sobrescrevemos um método, podemos chamar o original explicitamente:

```go
func (g Gato) Falar() string {
    return g.Animal.Falar() + " mas também diz Miau!"
}

fmt.Println(gato.Falar()) // "O animal faz um som mas também diz Miau!"
```

📌 **Isso é útil para reutilizar comportamento sem descartar a implementação original.**

---

## **9.3.4 Quando um Método do Struct Embutido Não é Promovido?**

Os métodos do struct embutido **não são promovidos** se houver um conflito de nome com um campo:

```go
type Carro struct {
    Marca string
}

func (c Carro) Nome() string {
    return "Carro da marca " + c.Marca
}

type Eletrico struct {
    Carro
    Nome string // Conflito!
}

ev := Eletrico{Carro{"Tesla"}, "Modelo X"}

fmt.Println(ev.Nome)       // "Modelo X"
// fmt.Println(ev.Nome())  // Erro! Nome é um campo, não um método
```

📌 **Se um campo e um método compartilharem o mesmo nome, o campo tem prioridade.**

---

## **9.3.5 Embedding e Interfaces**

Se um struct embutido implementa uma interface, o struct externo também a implementa:

```go
type Falante interface {
    Falar() string
}

type Papagaio struct {
    Animal
}

var f Falante = Papagaio{Animal{"Loro"}}
fmt.Println(f.Falar()) // "O animal faz um som"
```

📌 **Isso permite que um struct automaticamente implemente uma interface ao embutir outro struct.**

✅ **É uma forma eficiente de reutilizar comportamento sem herança tradicional.**

---

## **9.3.6 Boas Práticas**

✔ **Use embedding para reaproveitar código sem herança rígida.**  
✔ **Evite sobrescrever métodos sem necessidade — prefira chamar o método original.**  
✔ **Se precisar sobrescrever um método, garanta que ele mantém a lógica esperada.**  
✔ **Evite conflitos de nome entre métodos e campos.**  

---

## **Conclusão**

O **embedding de structs** promove métodos automaticamente, tornando Go uma linguagem poderosa para composição de código.  
No próximo capítulo, compararemos **composição vs. herança tradicional**, destacando quando cada abordagem deve ser utilizada! 🚀
