# **7.4 Alocação Dinâmica com `new` e `make`**

Go gerencia a memória automaticamente, mas oferece duas funções fundamentais para **alocação dinâmica**:  
- **`new`**: Aloca memória para um único valor e retorna um ponteiro para ele.
- **`make`**: Cria e inicializa slices, maps e channels.

Nesta seção, exploraremos:

- Diferença entre `new` e `make`
- Quando usar cada um
- Como funciona a alocação de memória em Go
- O impacto na performance e boas práticas

---

## **7.4.1 `new`: Alocação de Memória para Valores Únicos**

A função `new` aloca espaço na memória para um valor do tipo especificado e retorna um **ponteiro para ele**.

```go
p := new(int) // Aloca um inteiro e retorna um ponteiro

fmt.Println(*p) // 0 (valor padrão de int)
```

🔎 **Visualização da memória**:

```
+--------------------+
|  Endereço: 0xc0000 |
|  Valor: 0         |
+--------------------+
        ^
        |
        p (ponteiro)
```

📌 **`new` apenas aloca memória, mas não inicializa slices, maps ou channels.**

✅ **Exemplo com struct:**

```go
type Pessoa struct {
    Nome  string
    Idade int
}

p := new(Pessoa)
p.Nome = "Alice"

fmt.Println(p) // &{Alice 0}
```

---

## **7.4.2 `make`: Criando e Inicializando Estruturas Dinâmicas**

Diferente de `new`, a função `make` **inicializa** slices, maps e channels.  

📌 **Usamos `make` para esses tipos porque eles possuem metadados internos.**

```go
s := make([]int, 5) // Cria um slice de tamanho 5
fmt.Println(s) // [0 0 0 0 0]
```

✅ **Exemplo com mapas e canais:**

```go
m := make(map[string]int) // Inicializa um mapa
m["Alice"] = 25

ch := make(chan int) // Cria um canal
```

📌 **Se tentarmos usar `new` com slices, maps ou channels, teremos um ponteiro para um valor `nil`.**

```go
s := new([]int)
fmt.Println(s == nil) // false, mas `*s` ainda é `nil`!
```

✅ **Por isso, sempre use `make` para esses tipos!**

---

## **7.4.3 Diferença Entre `new` e `make`**

| Função | Para Que Serve? | Retorna |
|--------|---------------|---------|
| `new`  | Aloca memória para um valor único | Ponteiro (`*T`) |
| `make` | Inicializa slices, maps e channels | Valor inicializado (`T`) |

📌 **Resumo:**
- **Use `new`** para alocar memória para valores únicos (ex: `structs`, `int`, `bool`).
- **Use `make`** para criar **slices, maps e channels**.

---

## **7.4.4 Como o Go Gerencia a Memória?**

Go usa **gerenciamento automático de memória**, sem necessidade de `malloc` ou `free`.  
A linguagem possui um **Garbage Collector (GC)** que libera memória automaticamente.

📌 **Quando uma variável é alocada:**

1️⃣ Se for um valor local pequeno, ele pode ser armazenado na **stack** (rápido).  
2️⃣ Se for uma estrutura maior ou um ponteiro, pode ser alocado na **heap** (mais lento).  
3️⃣ O **Garbage Collector** libera memória quando os objetos não são mais referenciados.

✅ **Exemplo de alocação stack vs heap:**

```go
func criarValor() int {
    v := 42  // Alocado na stack
    return v
}

func criarPonteiro() *int {
    p := new(int) // Alocado na heap
    *p = 42
    return p
}
```

📌 **Ponteiros podem fazer com que um valor escape da stack para a heap.**

---

## **7.4.5 Impacto na Performance e Boas Práticas**

✔ **Prefira valores por cópia para tipos pequenos (`int`, `bool`).**  
✔ **Use `make` para inicializar slices e maps corretamente.**  
✔ **Evite criar ponteiros desnecessários (`*int` ao invés de `int`).**  
✔ **Use perfis (`pprof`) para identificar alocações excessivas.**  

---

## **Conclusão**

As funções `new` e `make` são essenciais para gerenciar memória em Go, mas devem ser usadas corretamente.  
No próximo capítulo, exploraremos **o funcionamento interno do Garbage Collector do Go**! 🚀
