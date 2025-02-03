# **6.5 Comparação de Structs**

Em Go, **structs podem ser comparados diretamente**, desde que todos os seus campos sejam comparáveis. No entanto, para casos mais complexos, onde há slices, maps ou ponteiros, precisamos de abordagens específicas.

Nesta seção, abordaremos:

- Como comparar structs diretamente
- O impacto de ponteiros e referências na comparação
- Como comparar structs contendo slices e maps
- O uso de `reflect.DeepEqual()` para comparações profundas
- Melhorando eficiência e segurança em comparações

---

## **6.5.1 Comparação Direta de Structs**

Se todos os campos de um struct forem tipos **comparáveis** (inteiros, strings, booleanos, arrays de tamanho fixo), podemos compará-los diretamente:

```go
type Pessoa struct {
    Nome  string
    Idade int
}

p1 := Pessoa{"Alice", 30}
p2 := Pessoa{"Alice", 30}

fmt.Println(p1 == p2) // true
```

📌 **A comparação direta só funciona se todos os campos puderem ser comparados nativamente.**

✅ **Arrays são comparáveis, mas slices não:**

```go
type Dados struct {
    Valores [3]int // Array fixo pode ser comparado
}

d1 := Dados{[3]int{1, 2, 3}}
d2 := Dados{[3]int{1, 2, 3}}

fmt.Println(d1 == d2) // true
```

---

## **6.5.2 Structs com Ponteiros**

Se um struct contém ponteiros, a comparação verifica **os valores apontados**, não apenas os endereços:

```go
type Pessoa struct {
    Nome  string
    Idade *int
}

idade1 := 30
idade2 := 30

p1 := Pessoa{"Alice", &idade1}
p2 := Pessoa{"Alice", &idade2}

fmt.Println(p1 == p2) // true (valores iguais)
```

📌 **Se os ponteiros apontassem para valores diferentes, a comparação falharia.**

✅ **Comparação de ponteiros por referência:**

```go
p3 := Pessoa{"Alice", &idade1}
p4 := Pessoa{"Alice", &idade1}

fmt.Println(p3 == p4) // true (mesmo ponteiro)
```

---

## **6.5.3 Comparação de Structs com Slices e Maps**

Como **slices e maps não podem ser comparados diretamente**, precisamos de abordagens alternativas.

```go
type Pessoa struct {
    Nome  string
    Tags  []string // Slice não é comparável diretamente
}

p1 := Pessoa{"Alice", []string{"Go", "Dev"}}
p2 := Pessoa{"Alice", []string{"Go", "Dev"}}

// fmt.Println(p1 == p2) // ERRO: slices não são comparáveis
```

📌 **Aqui, `reflect.DeepEqual()` é necessário para comparar slices.**

✅ **Comparando structs com `reflect.DeepEqual()`:**

```go
import "reflect"

fmt.Println(reflect.DeepEqual(p1, p2)) // true
```

💡 **Isso compara os valores dentro dos slices, garantindo equivalência correta.**

---

## **6.5.4 Comparação Eficiente de Structs**

Para evitar problemas de performance ao comparar structs grandes:

✔ **Use comparação direta (`==`) sempre que possível.**  
✔ **Para structs contendo slices ou maps, use `reflect.DeepEqual()` apenas quando necessário.**  
✔ **Se possível, converta o struct em uma representação `string` para comparação rápida:**

```go
import "encoding/json"

func structToString(v any) string {
    jsonBytes, _ := json.Marshal(v)
    return string(jsonBytes)
}

p1 := Pessoa{"Alice", []string{"Go", "Dev"}}
p2 := Pessoa{"Alice", []string{"Go", "Dev"}}

fmt.Println(structToString(p1) == structToString(p2)) // true
```

📌 **Isso é mais eficiente que `reflect.DeepEqual()` para grandes estruturas.**

---

## **6.5.5 Comparação com Outras Linguagens**

| Recurso | Go | C | Java | Python |
|---------|----|---|------|--------|
| Comparação direta (`==`) | ✅ | ❌ | ❌ (`equals()`) | ✅ |
| Comparação de slices | ❌ | ❌ | ✅ | ✅ |
| `reflect.DeepEqual()` | ✅ | ❌ | ❌ | ✅ |
| Ponteiros comparáveis | ✅ | ✅ | ✅ | ✅ |

📌 **Diferente de C e Java, Go permite comparar structs diretamente, simplificando verificações de igualdade.**

---

## **6.5.6 Boas Práticas**

✔ **Use `==` para structs com tipos primitivos.**  
✔ **Para slices e maps, utilize `reflect.DeepEqual()` com cautela.**  
✔ **Evite comparação direta entre structs grandes para melhorar performance.**  
✔ **Considere representar structs como strings (`json.Marshal`) para comparações eficientes.**  

---

## **Conclusão Geral**

A comparação de structs em Go é direta para tipos primitivos, mas requer abordagens específicas para slices, maps e ponteiros.  
No próximo capítulo, exploraremos **ponteiros e gerenciamento de memória**, abordando como otimizar o uso da RAM em Go! 🚀


🎉 **Parabéns!** 🎉

🚀 Você está pronto para usar structs em Go! 🎯

---
Cobrimos praticamente tudo que você precisa saber sobre **structs** em Go! Você também pode querer explorar os links da seção a seguir para aprofundar seus conhecimentos.

🕵️ **Para saber mais:**
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Go by Example: JSON](https://gobyexample.com/json)
- [Go by Example: String Formatting](https://gobyexample.com/string-formatting)
- [The Go Blog: JSON and Go](https://blog.golang.org/json-and-go)
- [The Go Blog: Method Sets](https://blog.golang.org/method-sets)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)
- [The Go Blog: Custom JSON Marshalling](https://blog.golang.org/json-and-go)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)
- [The Go Blog: Advanced JSON Handling](https://blog.golang.org/json)
- [The Go Blog: Stringer](https://blog.golang.org/laws-of-reflection#TOC_7.)
- [The Go Blog: JSON and struct composition](https://blog.golang.org/json-and-go)

---
