# **5.1 Declaração e Manipulação de Arrays**

Os **arrays** são um dos tipos fundamentais de estrutura de dados em Go. Eles fornecem um bloco de memória contígua, permitindo armazenamento e acesso eficiente a elementos. Embora Go prefira o uso de **slices** na maioria dos casos, entender arrays é essencial para compreender como a linguagem gerencia memória e otimiza operações de dados.

Nesta seção, exploraremos:

- Declaração e inicialização de arrays
- Acessando e modificando elementos
- Arrays fixos vs. slices dinâmicos
- Percorrendo arrays de forma eficiente
- Comparação de arrays com outras linguagens

---

## **5.1.1 Declaração de Arrays**

Um **array** em Go é uma coleção de elementos de mesmo tipo e tamanho fixo. Sua sintaxe é:

```go
var nome [tamanho]tipo
```

### **Exemplos de Declaração**

```go
var numeros [5]int // Array de 5 inteiros
var nomes [3]string // Array de 3 strings
var flags [2]bool // Array de 2 valores booleanos
```

📌 **O tamanho do array faz parte do seu tipo e não pode ser alterado após a declaração!**  

```go
var a [5]int
var b [10]int

// fmt.Println(a == b) // ERRO: arrays de tamanhos diferentes não podem ser comparados
```

### **Inicialização de Arrays**

Podemos inicializar arrays com valores padrão:

```go
var numeros = [3]int{1, 2, 3} // Inicializando diretamente
nomes := [2]string{"Alice", "Bob"} // Forma compacta

// Inicialização parcial (valores ausentes serão zero)
valores := [5]int{1, 2} // [1, 2, 0, 0, 0]
```

📌 **Os arrays em Go são sempre inicializados com valores zero do tipo correspondente.**

Outra forma de declarar sem definir um tamanho fixo (inferido pelo compilador):

```go
numeros := [...]int{10, 20, 30} // O compilador determina o tamanho automaticamente
fmt.Println(len(numeros)) // 3
```

---

## **5.1.2 Acessando e Modificando Elementos**

Os elementos de um array são acessados por índice, começando em `0`:

```go
var nums = [3]int{10, 20, 30}

fmt.Println(nums[0]) // 10
fmt.Println(nums[2]) // 30

// Modificando valores
nums[1] = 50
fmt.Println(nums) // [10, 50, 30]
```

📌 **A tentativa de acessar um índice fora dos limites causará um erro de runtime (`index out of range`).**  

---

## **5.1.3 Arrays e Memória**

Os arrays são armazenados de forma **contígua na memória**, o que permite acesso eficiente:

```go
var a = [4]int{1, 2, 3, 4}
fmt.Printf("Endereço de a[0]: %p\n", &a[0])
fmt.Printf("Endereço de a[1]: %p\n", &a[1]) // Alocado contiguamente na memória
```

📌 **Diferente de slices, arrays ocupam um bloco fixo de memória e não crescem dinamicamente.**

---

## **5.1.4 Comparação de Arrays**

Em Go, arrays **podem ser comparados diretamente** se tiverem o mesmo tamanho e tipo:

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{1, 2, 4}

fmt.Println(a == b) // true
fmt.Println(a == c) // false
```

📌 **Diferente de slices e maps, arrays podem ser comparados diretamente sem precisar de loops.**

---

## **5.1.5 Percorrendo Arrays com `for` e `range`**

### **Usando `for` Clássico**

```go
nums := [3]int{5, 10, 15}

for i := 0; i < len(nums); i++ {
    fmt.Println("Índice:", i, "Valor:", nums[i])
}
```

### **Usando `range`**

O `range` simplifica a iteração:

```go
for i, v := range nums {
    fmt.Println("Índice:", i, "Valor:", v)
}
```

📌 **Se não precisarmos do índice, podemos ignorá-lo usando `_`.**

```go
for _, v := range nums {
    fmt.Println("Valor:", v)
}
```

---

## **5.1.6 Arrays vs. Slices: Por Que Preferimos Slices?**

Os arrays têm um tamanho fixo e não podem crescer. Isso torna seu uso limitado quando não sabemos o tamanho exato dos dados. **Slices são mais flexíveis** e geralmente preferidos em Go.

| Característica | Arrays | Slices |
|--------------|--------|--------|
| Tamanho fixo | ✅ Sim | ❌ Não |
| Redimensionável | ❌ Não | ✅ Sim |
| Eficiência | ✅ Rápido | ✅ Rápido |
| Comparável | ✅ Sim | ❌ Não (apenas com `reflect.DeepEqual`) |

📌 **Na prática, slices são usados 90% das vezes, enquanto arrays são mais comuns para estruturação interna de dados.**

---

## **5.1.7 Quando Usar Arrays?**

✔ **Se o tamanho for conhecido e fixo** (exemplo: matrizes 3x3, buffers fixos).  
✔ **Para garantir que o tamanho não mude acidentalmente** (exemplo: IPv4 `[4]byte`).  
✔ **Em benchmarks ou otimizações específicas** para evitar overheads de slices.  

Caso contrário, **prefira slices**!

---

## **5.1.8 Comparação com Outras Linguagens**

| Recurso       | Go            | C   | Java  | Python |
|--------------|--------------|-----|------|--------|
| Arrays Fixos | ✅ Sim        | ✅ Sim | ✅ Sim | ❌ (listas dinâmicas) |
| Tamanho Dinâmico | ❌ Não | ❌ Não | ✅ Sim (`ArrayList`) | ✅ Sim (`list`) |
| Comparação Direta | ✅ Sim | ❌ Não | ❌ Não | ❌ Não |
| Zero por padrão | ✅ Sim | ❌ Não (lixo de memória) | ✅ Sim | ✅ Sim |

📌 **Go trata arrays como tipos de primeira classe, enquanto C e Java precisam de mais gerenciamento manual.**

---

## **Conclusão**

Os arrays são uma estrutura fundamental em Go, mas raramente usados diretamente em comparação com slices. Compreender seu funcionamento ajuda a **otimizar a manipulação de memória** e evitar alocações desnecessárias.

No próximo capítulo, exploraremos **slices**, uma estrutura poderosa que permite manipulação dinâmica de dados! 🚀
