# **6.1 Declaração e Manipulação de Mapas (`map[key]value`)**

Os **mapas (`map[key]value`)** são uma das estruturas de dados mais poderosas e eficientes em Go, permitindo associar chaves a valores de forma rápida. Eles são implementados internamente como **tabelas de hash**, garantindo acessos e atualizações com complexidade média de **O(1)**.

Nesta seção, exploraremos:

- Como declarar e inicializar mapas
- Acesso e modificação de elementos
- Tratamento de valores inexistentes
- Comparação de mapas com arrays e slices
- Eficiência e melhores práticas

---

## **6.1.1 Declaração de Mapas**

Um mapa é declarado usando a seguinte sintaxe:

```go
var nome map[tipo-chave]tipo-valor
```

📌 **Inicialmente, um mapa declarado dessa forma é `nil` e precisa ser inicializado antes do uso.**

Exemplo:

```go
var pessoas map[string]int
fmt.Println(pessoas == nil) // true (mapa ainda não inicializado)
```

✅ **Forma recomendada: inicialização com `make()`.**

```go
pessoas := make(map[string]int) // Cria um mapa vazio
```

📌 **Também podemos inicializar um mapa diretamente com valores:**

```go
idades := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

---

## **6.1.2 Acessando e Modificando Mapas**

Podemos acessar valores no mapa usando a chave correspondente:

```go
fmt.Println(idades["Alice"]) // 25
```

📌 **Se uma chave não existir, o Go retorna o valor zero do tipo:**

```go
fmt.Println(idades["Carlos"]) // 0 (porque o tipo é `int`)
```

✅ **Verificando se uma chave existe:**

```go
idade, existe := idades["Carlos"]
if existe {
    fmt.Println("Idade:", idade)
} else {
    fmt.Println("Carlos não encontrado!")
}
```

📌 **Sempre use essa abordagem para evitar valores inesperados ao acessar mapas.**

✅ **Adicionando e atualizando valores:**

```go
idades["Carlos"] = 40 // Adiciona uma nova entrada
idades["Alice"] = 26  // Atualiza um valor existente
```

---

## **6.1.3 Removendo Elementos de um Mapa**

O Go fornece a função `delete()` para remover chaves de um mapa:

```go
delete(idades, "Bob")
fmt.Println(idades) // map[Alice:26 Carlos:40]
```

📌 **Se a chave não existir, `delete()` não causa erro.**

---

## **6.1.4 Iterando Sobre Mapas**

Podemos percorrer um mapa usando `range`:

```go
for nome, idade := range idades {
    fmt.Println(nome, "tem", idade, "anos")
}
```

📌 **A ordem de iteração não é garantida!**  
Se precisarmos de uma ordem específica, devemos **extrair as chaves, ordená-las e iterar manualmente.**

```go
var chaves []string
for k := range idades {
    chaves = append(chaves, k)
}
sort.Strings(chaves)

for _, k := range chaves {
    fmt.Println(k, "->", idades[k])
}
```

---

## **6.1.5 Mapas vs. Outras Estruturas de Dados**

| Estrutura | Quando Usar |
|-----------|------------|
| **Arrays** | Quando o número de elementos é fixo e acesso por índice for necessário |
| **Slices** | Quando a ordem dos elementos importa e o tamanho pode crescer |
| **Mapas**  | Quando precisamos de acesso rápido baseado em chave |

📌 **Mapas são mais rápidos que slices para busca, mas não possuem ordem definida.**

---

## **6.1.6 Eficiência e Boas Práticas**

✔ **Prefira `make(map[Tipo]Tipo, capacidade)` se souber o tamanho esperado, para otimizar alocações.**  
✔ **Use `delete()` para liberar memória de mapas que crescem dinamicamente.**  
✔ **Evite modificar mapas dentro de loops concorrentes sem `sync.Mutex` ou `sync.Map`.**  
✔ **Se a ordem for importante, use slices como suporte.**  

---

## **Conclusão**

Os mapas são extremamente úteis para armazenar associações chave-valor de forma eficiente.  
No próximo capítulo, veremos **operações comuns com mapas, como `delete`, `len` e `range`**, aprofundando seu uso em cenários reais. 🚀
