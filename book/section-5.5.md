# **5.5 Deep Copy vs. Shallow Copy**

Em Go, a forma como as variáveis são copiadas impacta diretamente a manipulação de memória e o comportamento de estruturas de dados. Existem dois tipos principais de cópias:

- **Shallow Copy** (Cópia Rasa): Copia apenas a referência para os dados originais.
- **Deep Copy** (Cópia Profunda): Copia **todos os dados**, criando uma nova instância independente.

Nesta seção, exploraremos:

- Diferença entre cópias rasas e profundas
- Como Go trata a cópia de arrays, slices e maps
- Como garantir que uma estrutura seja copiada corretamente
- Uso eficiente de `copy()` e `append()`

---

## **5.5.1 O Que é Shallow Copy?**

Uma **shallow copy** copia apenas **referências**, não os dados reais. Isso significa que **modificações no novo valor também afetam o original**.

Exemplo com slices:

```go
original := []int{1, 2, 3}
shallow := original // Apenas copia a referência

shallow[0] = 100

fmt.Println(original) // [100, 2, 3]
fmt.Println(shallow)  // [100, 2, 3]
```

📌 **Ambos os slices apontam para os mesmos dados na memória.**

✅ **Útil quando queremos evitar cópias desnecessárias.**  
❌ **Perigoso se quisermos preservar os dados originais.**

---

## **5.5.2 O Que é Deep Copy?**

Uma **deep copy** copia **todos os dados** para uma nova região de memória, garantindo que o original permaneça inalterado.

```go
original := []int{1, 2, 3}
deep := append([]int{}, original...) // Criando uma cópia independente

deep[0] = 100

fmt.Println(original) // [1, 2, 3] (inalterado)
fmt.Println(deep)     // [100, 2, 3]
```

📌 **`append([]T{}, slice...)` é a maneira recomendada de fazer cópias profundas de slices.**

✅ **Garante que o original não seja alterado.**  
❌ **Pode consumir mais memória.**

---

## **5.5.3 Como Go Trata a Cópia de Diferentes Estruturas?**

### **Arrays: Copiados por Valor (Deep Copy Automática)**

Arrays em Go são copiados por **valor**, ou seja, automaticamente fazem deep copy.

```go
a := [3]int{1, 2, 3}
b := a // Cópia completa dos dados

b[0] = 100

fmt.Println(a) // [1, 2, 3] (inalterado)
fmt.Println(b) // [100, 2, 3]
```

📌 **Cada array ocupa um espaço separado na memória.**

---

### **Slices: Copiados por Referência (Shallow Copy por Padrão)**

Slices são apenas um "ponteiro" para um array subjacente, então a cópia padrão é rasa:

```go
original := []int{1, 2, 3}
copy := original

copy[0] = 100

fmt.Println(original) // [100, 2, 3]
```

✅ **Para deep copy de slices, use `append()`:**

```go
deepCopy := append([]int{}, original...)
```

---

### **Maps: Sempre Shallow Copy**

Maps são copiados **por referência** em Go:

```go
original := map[string]int{"a": 1, "b": 2}
copy := original // Aponta para os mesmos dados

copy["a"] = 100

fmt.Println(original) // map[a:100 b:2]
```

✅ **Para deep copy de maps, devemos iterar manualmente:**

```go
func deepCopyMap(m map[string]int) map[string]int {
    copy := make(map[string]int)
    for k, v := range m {
        copy[k] = v
    }
    return copy
}

original := map[string]int{"a": 1, "b": 2}
copy := deepCopyMap(original)

copy["a"] = 100

fmt.Println(original) // map[a:1 b:2]
fmt.Println(copy)     // map[a:100 b:2]
```

---

### **Structs: Copiados por Valor, Mas Contendo Referências**

Structs são copiados por valor, mas se contiverem slices ou maps, as referências serão copiadas:

```go
type Data struct {
    Numbers []int
}

original := Data{Numbers: []int{1, 2, 3}}
copy := original

copy.Numbers[0] = 100

fmt.Println(original.Numbers) // [100, 2, 3] (original alterado)
```

✅ **Para deep copy, precisamos copiar os slices manualmente:**

```go
func deepCopyStruct(d Data) Data {
    copy := Data{Numbers: append([]int{}, d.Numbers...)}
    return copy
}

original := Data{Numbers: []int{1, 2, 3}}
copy := deepCopyStruct(original)

copy.Numbers[0] = 100

fmt.Println(original.Numbers) // [1, 2, 3] (inalterado)
```

---

## **5.5.4 Comparação de Performance: Shallow vs. Deep Copy**

| Estrutura | Padrão de Cópia | Método para Deep Copy |
|-----------|----------------|----------------------|
| Arrays    | Deep Copy       | Automático |
| Slices    | Shallow Copy    | `append([]T{}, slice...)` |
| Maps      | Shallow Copy    | Iteração manual (`make()`) |
| Structs   | Deep Copy (parcial) | Copiar manualmente slices/maps dentro |

📌 **Shallow copy é mais eficiente em memória, mas deep copy evita efeitos colaterais inesperados.**

---

## **5.5.5 Boas Práticas**

✔ **Use `append([]T{}, slice...)` para cópia profunda de slices.**  
✔ **Para maps, crie um novo e copie os elementos um por um.**  
✔ **Cuidado com struct que contém referências (`[]T` ou `map[T]T`).**  
✔ **Se precisar de alto desempenho, prefira shallow copy sempre que possível.**  

---

## **Conclusão**

A escolha entre **shallow copy e deep copy** depende do contexto. Shallow copies são rápidas e eficientes, mas podem causar efeitos colaterais inesperados. Para evitar isso, Go fornece ferramentas para criar cópias profundas de estruturas de dados quando necessário.

No próximo capítulo, exploraremos **ponteiros e alocação de memória**, abordando como otimizar o uso da RAM em Go! 🚀
