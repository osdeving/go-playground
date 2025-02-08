# **9.4 Composição vs. Herança em Go**

Em Go, **composição** é a abordagem preferida para reutilização de código, enquanto linguagens como Java e C++ utilizam **herança tradicional**.  
A composição permite combinar comportamentos sem criar dependências rígidas entre tipos, tornando o código mais modular e reutilizável.

Nesta seção, exploraremos:

- Diferenças entre composição e herança
- Como usar composição para compartilhar comportamento
- Quando evitar herança e preferir composição
- Exemplos práticos de uso

---

## **9.4.1 O Que é Herança e Seus Problemas?**

Em linguagens como Java e C++, a herança permite que uma classe **herde** métodos e atributos de outra:

```java
class Animal {
    String nome;
    void falar() {
        System.out.println("O animal faz um som");
    }
}

class Cachorro extends Animal {
    void latir() {
        System.out.println("Au au!");
    }
}
```

📌 **Problemas da Herança Tradicional:**  
- **Acoplamento forte** → Modificar uma classe base pode afetar todas as subclasses.  
- **Herança profunda** → Código difícil de manter e entender.  
- **Problemas de reutilização** → Métodos herdados podem não ser necessários em todas as subclasses.

✅ **Go evita esses problemas usando composição.**

---

## **9.4.2 Como a Composição Resolve Esses Problemas?**

Go permite reutilizar comportamento **sem herança**, simplesmente embutindo structs:

```go
type Animal struct {
    Nome string
}

func (a Animal) Falar() string {
    return "O animal faz um som"
}

type Cachorro struct {
    Animal // Embedding
}

dog := Cachorro{Animal{"Rex"}}
fmt.Println(dog.Falar()) // "O animal faz um som"
```

📌 **`Cachorro` reutiliza o comportamento de `Animal` sem acoplamento rígido.**

✅ **Vantagens da Composição:**  
✔ Maior flexibilidade.  
✔ Código mais modular.  
✔ Permite reuso de comportamento sem dependência hierárquica.  

---

## **9.4.3 Reutilização de Código com Interfaces**

Podemos combinar composição com interfaces para criar código flexível:

```go
type Falante interface {
    Falar() string
}

type Humano struct{ Nome string }

func (h Humano) Falar() string {
    return "Oi, eu sou " + h.Nome
}

type Robo struct {
    Humano
    Modelo string
}

var f Falante = Robo{Humano{"X-1000"}, "Androide"}
fmt.Println(f.Falar()) // "Oi, eu sou X-1000"
```

📌 **O struct `Robo` reutiliza `Falar()` sem precisar de herança.**

✅ **Isso mantém o código desacoplado e modular.**

---

## **9.4.4 Composição Dinâmica: Uso de Campos Embutidos**

Além do embedding de structs, podemos usar **composição dinâmica**:

```go
type Motor struct {
    Potencia int
}

type Carro struct {
    Motor *Motor // Composição via referência
}

c := Carro{Motor: &Motor{200}}
fmt.Println(c.Motor.Potencia) // 200
```

📌 **Isso permite trocar o comportamento dinamicamente sem modificar a estrutura do código.**

---

## **9.4.5 Comparação: Composição vs. Herança**

| Característica | Composição (Go) | Herança (Java, C++) |
|---------------|----------------|----------------------|
| Reutilização de Código | ✅ | ✅ |
| Flexibilidade | ✅ Alta | ❌ Baixa |
| Acoplamento | ✅ Baixo | ❌ Alto |
| Modificação Fácil | ✅ | ❌ Difícil |
| Suporte a Múltiplos Comportamentos | ✅ | ❌ Apenas uma superclasse |

📌 **A composição permite modificar e reutilizar código sem criar dependências rígidas.**

---

## **9.4.6 Boas Práticas**

✔ **Use composição sempre que possível para evitar dependências rígidas.**  
✔ **Se precisar reutilizar comportamento, prefira interfaces ou embedding em vez de herança.**  
✔ **Evite structs muito profundos — mantenha o código modular.**  
✔ **Use composição dinâmica (campos embutidos) quando precisar de maior flexibilidade.**  

---

## **Conclusão**

A **composição é a abordagem preferida em Go**, pois permite reutilizar código sem criar dependências hierárquicas.  
No próximo capítulo, entraremos na programação concorrente com **Goroutines e Channels**, explorando o poder da concorrência em Go! 🚀
