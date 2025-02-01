# **6.4 Campos Opcionais e `omitempty`**

Em Go, **structs não possuem campos opcionais nativamente**, já que todos os campos são inicializados com um valor padrão. No entanto, a linguagem fornece maneiras eficientes de lidar com **dados ausentes** e **otimizar a serialização**.

Nesta seção, exploraremos:

- Como representar campos opcionais em structs
- O uso da tag `omitempty` para JSON e outras codificações
- Como diferenciar valores padrão de valores não definidos
- Uso de ponteiros para representar nulabilidade
- Estratégias para manipular dados opcionais corretamente

---

## **6.4.1 Campos Opcionais em Go**

Diferente de outras linguagens como Python ou JavaScript, Go **não suporta valores `nil` diretamente em tipos primitivos**. Isso significa que todos os campos de um struct sempre terão um valor inicial.

Exemplo:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

p := Pessoa{}
fmt.Println(p.Nome)  // "" (string vazia)
fmt.Println(p.Idade) // 0 (inteiro inicializado)
```

📌 **Não há como distinguir um campo "não definido" de um campo "definido com o valor zero".**

✅ **Para representar a ausência de um valor, usamos ponteiros (`*`).**

```go
type Pessoa struct {
    Nome  *string
    Idade *int
}
```

Agora podemos diferenciar valores não definidos de valores explicitamente definidos:

```go
var idade int = 30
p := Pessoa{Nome: nil, Idade: &idade}

if p.Nome == nil {
    fmt.Println("Nome não definido!")
}
```

📌 **O uso de ponteiros em structs permite representar valores opcionais corretamente.**

---

## **6.4.2 Serialização com `omitempty`**

Ao trabalhar com JSON, podemos omitir campos vazios usando a tag `omitempty`:

```go
type Pessoa struct {
    Nome  string `json:"nome,omitempty"`
    Idade int    `json:"idade,omitempty"`
}
```

Agora, se um campo tiver o valor **zero**, ele não será incluído na saída JSON:

```go
p := Pessoa{Nome: "", Idade: 0}

jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData)) // "{}"
```

📌 **A tag `omitempty` remove automaticamente valores vazios do JSON.**

✅ **Isso reduz o tamanho da resposta e evita valores irrelevantes.**

---

## **6.4.3 Quando Usar Ponteiros vs. `omitempty`?**

| Estratégia         | Vantagens | Desvantagens |
|--------------------|-----------|--------------|
| **Usar `omitempty`** | JSON mais limpo, evita valores irrelevantes | Não permite diferenciar `0` de "não definido" |
| **Usar Ponteiros** | Permite `nil` para detectar valores não definidos | Aumenta a complexidade e uso de memória |
| **Criar Tipos Customizados** | Maior controle sobre valores opcionais | Mais código e complexidade extra |

📌 **Use `omitempty` para JSON, e ponteiros quando precisar diferenciar valores nulos.**

---

## **6.4.4 Estratégias Avançadas para Campos Opcionais**

### **1. Criando Tipos Customizados**

Podemos criar **tipos personalizados** para representar valores opcionais:

```go
type NullableInt struct {
    Valor int
    Definido bool
}
```

Agora podemos representar a ausência de valor:

```go
var idade NullableInt

if idade.Definido {
    fmt.Println("Idade:", idade.Valor)
} else {
    fmt.Println("Idade não definida")
}
```

📌 **Isso evita o uso excessivo de ponteiros e mantém segurança de tipos.**

### **2. Métodos para Campos Opcionais**

Podemos adicionar métodos para facilitar a manipulação:

```go
func (n NullableInt) String() string {
    if !n.Definido {
        return "N/A"
    }
    return fmt.Sprintf("%d", n.Valor)
}

idade := NullableInt{Valor: 30, Definido: true}
fmt.Println(idade.String()) // "30"
```

📌 **Isso melhora a legibilidade do código e encapsula a lógica de nulabilidade.**

---

## **6.4.5 Comparação com Outras Linguagens**

| Recurso | Go | JavaScript | Python | Java |
|---------|----|------------|--------|------|
| Valores `nil` | ❌ (exceto ponteiros) | ✅ | ✅ | ✅ |
| Campos opcionais | ✅ (`omitempty`) | ✅ | ✅ | ✅ (`Optional<T>`) |
| Serialização flexível | ✅ | ✅ | ✅ | ❌ |
| Segurança de tipos | ✅ | ❌ | ❌ | ✅ |

📌 **Go não possui `null` nativo, mas fornece estratégias eficientes para lidar com valores ausentes.**

---

## **6.4.6 Melhores Práticas**

✔ **Use `omitempty` para JSON quando valores padrão não forem necessários.**  
✔ **Use ponteiros para distinguir valores `0` de valores indefinidos.**  
✔ **Crie tipos customizados quando precisar representar nulabilidade de forma clara.**  
✔ **Evite ponteiros para tipos pequenos (`int`, `bool`) para não aumentar a complexidade.**  

---

## **Conclusão**

Go trata campos opcionais de maneira eficiente usando **`omitempty`**, **ponteiros** e **tipos customizados**.  
No próximo capítulo, exploraremos **comparação de structs**, abordando como verificar igualdade corretamente! 🚀
