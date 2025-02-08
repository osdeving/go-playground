# **4.5 Funções Anônimas e Closures**

Em Go, **funções anônimas** são funções sem um nome explícito, geralmente usadas para lógica rápida e temporária. Já os **closures** permitem capturar variáveis do escopo externo, tornando-as úteis para encapsular estados e criar funções mais dinâmicas.

Nesta seção, abordaremos:

- Como declarar e usar funções anônimas
- Passagem de parâmetros e retornos em funções anônimas
- O conceito de closures e sua aplicação prática
- Uso avançado de closures para encapsulamento de estado

---

## **4.5.1 O Que São Funções Anônimas?**

Uma função anônima é simplesmente uma função sem nome:

```go
func() {
    fmt.Println("Função anônima executada!")
}()
```

📌 **Note que a função foi chamada imediatamente com `()`.**

### **Atribuindo a uma Variável**

```go
mensagem := func() {
    fmt.Println("Olá, mundo!")
}

mensagem() // Chama a função
```

📌 **Funções anônimas podem ser armazenadas em variáveis e chamadas posteriormente.**

---

## **4.5.2 Funções Anônimas com Parâmetros e Retorno**

Funções anônimas podem receber parâmetros e retornar valores:

```go
soma := func(a, b int) int {
    return a + b
}

resultado := soma(10, 20)
fmt.Println(resultado) // 30
```

📌 **Elas seguem a mesma sintaxe de funções normais, apenas sem nome.**

---

## **4.5.3 Closures: Funções que Capturam Variáveis Externas**

Um **closure** é uma função que **captura variáveis do escopo externo**, permitindo criar funções dinâmicas e encapsular estados.

```go
func contador() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

incrementa := contador()

fmt.Println(incrementa()) // 1
fmt.Println(incrementa()) // 2
fmt.Println(incrementa()) // 3
```

📌 **A variável `i` é mantida na memória mesmo após `contador` ter retornado.**

---

## **4.5.4 Encapsulamento de Estado com Closures**

Closures são úteis para encapsular estados e evitar variáveis globais:

```go
func novoContador(nome string) func() string {
    contador := 0
    return func() string {
        contador++
        return fmt.Sprintf("%s: %d", nome, contador)
    }
}

contadorA := novoContador("A")
contadorB := novoContador("B")

fmt.Println(contadorA()) // A: 1
fmt.Println(contadorA()) // A: 2
fmt.Println(contadorB()) // B: 1
```

📌 **Cada closure mantém seu próprio estado independentemente.**

---

## **4.5.5 Closures e Funções de Ordem Superior**

Closures podem ser usados para criar **funções de ordem superior**, que retornam ou recebem funções:

```go
func multiplicador(fator int) func(int) int {
    return func(x int) int {
        return x * fator
    }
}

dobrar := multiplicador(2)
triplicar := multiplicador(3)

fmt.Println(dobrar(10)) // 20
fmt.Println(triplicar(10)) // 30
```

📌 **Isso permite reutilizar lógica de forma eficiente.**

---

## **4.5.6 Comparação com Outras Linguagens**

| Recurso               | Go  | JavaScript | Python | C |
|----------------------|----|------------|--------|---|
| Funções Anônimas     | ✅  | ✅ (`()=>{}`) | ✅ (`lambda`) | ❌ |
| Closures             | ✅  | ✅ | ✅ | ❌ |
| Captura de Variáveis | ✅  | ✅ | ✅ | ❌ |
| Encapsulamento       | ✅  | ✅ | ✅ | ❌ |

📌 **Go tem suporte nativo para closures, mas sem `this` como em JavaScript.**

---

## **Conclusão**

Funções anônimas e closures são ferramentas poderosas para manipular funções dinamicamente. No próximo capítulo, exploraremos **recursão**, um conceito fundamental na programação! 🚀
