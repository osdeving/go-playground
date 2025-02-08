# 📌 Seção 8.2: Receptores (`value receiver` vs `pointer receiver`) em Go

## Introdução

Em Go, as funções podem ser associadas a tipos através de **métodos**. Para isso, usamos **receptores** (receivers), que podem ser:

1. **Value Receivers (`value receiver`)**: O método recebe uma cópia do valor original.
2. **Pointer Receivers (`pointer receiver`)**: O método recebe um ponteiro para o valor original, permitindo modificações no estado do objeto.

Este capítulo explora esses dois tipos de receptores, seus usos e boas práticas.

---

## 🔹 Value Receiver (`value receiver`)

Quando um método tem um **value receiver**, ele recebe uma **cópia** do objeto, o que significa que qualquer alteração feita dentro do método **não afeta o objeto original**.

### 📌 Exemplo:
```go
package main

import "fmt"

// Definição de um tipo struct
type Circulo struct {
    raio float64
}

// Método com value receiver
func (c Circulo) Area() float64 {
    return 3.14 * c.raio * c.raio
}

func main() {
    c := Circulo{raio: 5}
    fmt.Println("Área:", c.Area()) // Área: 78.5
}
```

### 🔥 Características de `value receiver`:
✅ **Seguro para leitura**: Como trabalha com cópias, garante que o objeto original não seja alterado.

✅ **Mais eficiente para tipos pequenos**: Structs pequenas (como `int`, `float64`) são leves para copiar.

❌ **Ineficiente para structs grandes**: Se a struct for grande, cada chamada do método criará uma nova cópia na memória, o que pode impactar o desempenho.

---

## 🔹 Pointer Receiver (`pointer receiver`)

Quando um método tem um **pointer receiver**, ele recebe um **ponteiro para o objeto**, permitindo modificar seu estado original.

### 📌 Exemplo:
```go
package main

import "fmt"

// Definição de um tipo struct
type Contador struct {
    valor int
}

// Método com pointer receiver (modifica o estado do objeto)
func (c *Contador) Incrementar() {
    c.valor++
}

func main() {
    c := Contador{valor: 0}
    c.Incrementar()
    fmt.Println("Valor do contador:", c.valor) // Valor do contador: 1
}
```

### 🔥 Características de `pointer receiver`:
✅ **Permite modificações**: Como trabalha diretamente com o objeto, alterações feitas no método são refletidas no original.

✅ **Mais eficiente para structs grandes**: Em vez de copiar toda a struct, o Go passa um ponteiro, economizando memória e melhorando o desempenho.

❌ **Necessita de um ponteiro na chamada do método**: O Go facilita isso automaticamente em muitos casos, mas pode ser um detalhe importante.

---

## 🎯 Quando usar cada um?

| Situação | Value Receiver | Pointer Receiver |
|------------|---------------|-----------------|
| A struct é pequena e eficiente para copiar | ✅ Sim | ❌ Não |
| O método não altera o estado do objeto | ✅ Sim | ❌ Não |
| O método precisa modificar o estado | ❌ Não | ✅ Sim |
| A struct é grande e custosa para copiar | ❌ Não | ✅ Sim |

### 📌 Exemplo de otimização
Se tivermos uma struct muito grande, usar `value receiver` seria ineficiente. Veja um exemplo com `pointer receiver`:

```go
package main

import "fmt"

type Documento struct {
    conteudo string
}

func (d *Documento) Editar(novoConteudo string) {
    d.conteudo = novoConteudo
}

func main() {
    doc := Documento{conteudo: "Texto inicial"}
    doc.Editar("Texto modificado")
    fmt.Println(doc.conteudo) // Texto modificado
}
```

Como `Documento` pode crescer muito, passar um ponteiro evita a cópia desnecessária.

---

## 📌 Conclusão
1. Use **value receiver** quando não precisar modificar a struct e ela for pequena.
2. Use **pointer receiver** quando precisar alterar o estado ou evitar cópias desnecessárias.
3. Structs que usam **pointer receivers** podem implementar interfaces tanto para valores quanto para ponteiros.

🔹 Dominar `value receiver` e `pointer receiver` é essencial para escrever código eficiente e idiomático em Go! 🚀
