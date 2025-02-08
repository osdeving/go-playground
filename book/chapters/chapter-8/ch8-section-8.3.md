# 📌 Seção 8.3: Interfaces e Polimorfismo em Go

## Introdução

Go é uma linguagem que suporta polimorfismo através do uso de **interfaces**. Interfaces permitem definir conjuntos de comportamentos sem especificar como eles são implementados. Em Go, a implementação de uma interface é **implícita**, ou seja, um tipo satisfaz uma interface automaticamente se ele implementa seus métodos.

Este capítulo explora o conceito de interfaces, como utilizá-las para criar código reutilizável e como o polimorfismo é aplicado em Go.

---

## 🔹 O que são Interfaces em Go?

Uma interface em Go define um conjunto de métodos que um tipo precisa implementar. Qualquer tipo que implementar esses métodos será considerado como compatível com a interface.

### 📌 Exemplo básico de interface:
```go
package main

import "fmt"

// Definição de uma interface
 type Forma interface {
    Area() float64
}

// Struct que implementa a interface
type Retangulo struct {
    largura, altura float64
}

// Implementação do método Area() para Retangulo
func (r Retangulo) Area() float64 {
    return r.largura * r.altura
}

func main() {
    var f Forma = Retangulo{largura: 10, altura: 5}
    fmt.Println("Área do retângulo:", f.Area()) // Área do retângulo: 50
}
```

### 🔥 Características importantes:
✅ **Implementação implícita**: Não é necessário declarar explicitamente que um tipo implementa uma interface.
✅ **Permite polimorfismo**: O mesmo código pode manipular diferentes tipos que implementam a mesma interface.
✅ **Flexibilidade**: Qualquer tipo pode implementar uma interface, desde que possua os métodos exigidos.

---

## 🔹 Polimorfismo com Interfaces

O polimorfismo em Go permite que diferentes tipos sejam tratados de maneira uniforme ao implementarem a mesma interface. Isso possibilita a escrita de código mais genérico e modular.

### 📌 Exemplo com múltiplos tipos:
```go
package main

import "fmt"

type Forma interface {
    Area() float64
}

type Circulo struct {
    raio float64
}

type Quadrado struct {
    lado float64
}

// Implementação do método Area() para Circulo
func (c Circulo) Area() float64 {
    return 3.14 * c.raio * c.raio
}

// Implementação do método Area() para Quadrado
func (q Quadrado) Area() float64 {
    return q.lado * q.lado
}

func CalcularArea(f Forma) {
    fmt.Println("Área:", f.Area())
}

func main() {
    c := Circulo{raio: 5}
    q := Quadrado{lado: 4}
    
    CalcularArea(c) // Área: 78.5
    CalcularArea(q) // Área: 16
}
```

Neste exemplo, `CalcularArea` pode receber qualquer tipo que implemente a interface `Forma`, demonstrando o polimorfismo.

---

## 🔹 Interfaces embutidas e composição

Go permite que interfaces sejam compostas através da **embutida (embedding)**, o que facilita a criação de interfaces mais complexas.

### 📌 Exemplo de interface composta:
```go
package main

import "fmt"

type Leitor interface {
    Ler() string
}

type Escritor interface {
    Escrever(texto string)
}

type Dispositivo interface {
    Leitor
    Escritor
}

type Notebook struct {
    conteudo string
}

func (n *Notebook) Ler() string {
    return n.conteudo
}

func (n *Notebook) Escrever(texto string) {
    n.conteudo = texto
}

func main() {
    var d Dispositivo = &Notebook{}
    d.Escrever("Olá, Go!")
    fmt.Println(d.Ler()) // Olá, Go!
}
```

Aqui, `Dispositivo` combina as interfaces `Leitor` e `Escritor`, criando um tipo mais poderoso e modular.

---

## 🔹 Interface vazia (`interface{}`) e `any`

Em Go, a interface vazia (`interface{}`) pode ser usada para representar **qualquer tipo**. No Go 1.18+, o alias `any` foi introduzido para facilitar a leitura do código.

### 📌 Exemplo:
```go
package main

import "fmt"

func MostrarValor(v any) {
    fmt.Println("Valor recebido:", v)
}

func main() {
    MostrarValor(42)
    MostrarValor("Texto genérico")
    MostrarValor(3.14)
}
```

**⚠️ Cuidado:** Como `interface{}` aceita qualquer tipo, pode ser necessário **fazer type assertion** para recuperar o valor original.

```go
valor, ok := v.(int) // Tenta converter v para int
if ok {
    fmt.Println("É um int com valor:", valor)
}
```

---

## 📌 Conclusão

1. Interfaces em Go são uma ferramenta poderosa para modelar comportamento.
2. A implementação implícita facilita a flexibilidade e modularidade do código.
3. O polimorfismo permite que métodos genéricos manipulem diferentes tipos sem conhecer seus detalhes internos.
4. A interface vazia (`interface{}`) pode representar qualquer tipo, mas deve ser usada com cautela.

🔹 **Dominar interfaces é essencial para escrever código escalável e reutilizável em Go!** 🚀

