# **9.1 Embedding de Structs (Herança Simples)**

Go não possui **herança** no sentido tradicional, como em Java ou C++, mas permite reutilizar código por meio de **embedding de structs**. Isso permite que um struct "herde" comportamentos de outro sem necessidade de hierarquias complexas.

Nesta seção, exploraremos:

- O que é embedding de structs em Go
- Como reutilizar código sem herança tradicional
- Como sobrescrever métodos em structs embutidos
- Benefícios e boas práticas ao usar embedding

---

## **9.1.1 O Que é Embedding de Structs?**

Em Go, podemos **incluir um struct dentro de outro**, permitindo acesso direto aos seus campos e métodos.

```go
type Animal struct {
    Nome string
}

type Cachorro struct {
    Animal // Embedding do struct Animal
    Raca   string
}

dog := Cachorro{Animal{"Rex"}, "Labrador"}
fmt.Println(dog.Nome) // "Rex" (acessando campo da struct embutida)
fmt.Println(dog.Raca) // "Labrador"
```

📌 **Diferente de herança tradicional, `Cachorro` não é uma subclasse de `Animal`, mas pode acessar seus campos diretamente.**

🔎 **Visualização da memória:**

```
+----------------+
| Cachorro       |
|  Nome: "Rex"   |  <--- Herdado de Animal
|  Raca: "Labrador" |
+----------------+
```

✅ **Isso permite reutilizar código sem criar dependências rígidas entre tipos.**

---

## **9.1.2 Chamando Métodos do Struct Embutido**

Se um struct embutido possui métodos, o struct externo pode chamá-los diretamente.

```go
type Animal struct {
    Nome string
}

func (a Animal) Falar() string {
    return "O animal faz um som"
}

type Cachorro struct {
    Animal
    Raca string
}

dog := Cachorro{Animal{"Rex"}, "Labrador"}
fmt.Println(dog.Falar()) // "O animal faz um som"
```

📌 **O struct `Cachorro` herdou o método `Falar()` de `Animal` automaticamente.**

✅ **Também podemos chamar o método explicitamente:**

```go
fmt.Println(dog.Animal.Falar()) // "O animal faz um som"
```

---

## **9.1.3 Sobrescrevendo Métodos em Embeddings**

Podemos sobrescrever métodos simplesmente definindo um novo método com o mesmo nome.

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

📌 **O método `Falar()` da struct `Gato` sobrescreve o método herdado de `Animal`.**

✅ **Chamando o método original:**

```go
fmt.Println(gato.Animal.Falar()) // "O animal faz um som"
```

---

## **9.1.4 Embedding e Interfaces**

Podemos embutir structs que implementam interfaces, tornando a composição ainda mais poderosa.

```go
type Falante interface {
    Falar() string
}

type Humano struct{ Nome string }

func (h Humano) Falar() string {
    return "Oi, meu nome é " + h.Nome
}

type Robo struct {
    Humano
    Modelo string
}

r := Robo{Humano{"X-1000"}, "Androide"}
fmt.Println(r.Falar()) // "Oi, meu nome é X-1000"
```

📌 **O struct `Robo` automaticamente implementa `Falante`, pois `Humano` já implementa.**

---

## **9.1.5 Comparação com Herança Tradicional**

| Característica | Go (Embedding) | Java (Herança) | C++ (Herança) |
|---------------|---------------|---------------|--------------|
| Reutilização de Código | ✅ | ✅ | ✅ |
| Encapsulamento Fraco | ❌ | ✅ | ✅ |
| Acoplamento Baixo | ✅ | ❌ | ❌ |
| Múltipla "Herança" | ✅ (composição) | ❌ | ✅ |
| Sobrescrita de Métodos | ✅ | ✅ (`@Override`) | ✅ (`virtual`) |

📌 **Embedding é mais flexível e evita os problemas de herança tradicional.**

---

## **9.1.6 Boas Práticas**

✔ **Use embedding para reuso de código, mas evite dependências profundas.**  
✔ **Se precisar sobrescrever um método, considere se a composição é realmente necessária.**  
✔ **Evite acessar diretamente a struct embutida dentro de métodos externos.**  
✔ **Prefira composição (`has-a`) em vez de herança rígida (`is-a`).**  

---

## **Conclusão**

O **embedding de structs** permite reutilizar código de forma simples e eficiente, sem os problemas da herança tradicional.  
No próximo capítulo, exploraremos **implementação de múltiplas interfaces em Go**, aumentando a flexibilidade dos nossos tipos! 🚀
