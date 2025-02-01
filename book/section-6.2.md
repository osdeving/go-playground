# **6.2 Operações Comuns (`delete`, `len`, `range`)**

Os **mapas (`map[key]value`)** são estruturas altamente eficientes para armazenar pares **chave-valor**. Além da manipulação básica, existem operações essenciais que tornam os mapas ainda mais úteis, como remoção de elementos, contagem e iteração.

Nesta seção, exploraremos:

- Como remover elementos de um mapa com `delete()`
- Obtendo o número total de elementos com `len()`
- Iterando sobre mapas com `range`
- Boas práticas e uso eficiente de operações em mapas

---

## **6.2.1 Removendo Elementos com `delete()`**

A função `delete()` permite remover uma chave específica de um mapa:

```go
pessoas := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carlos": 40,
}

delete(pessoas, "Bob") // Remove "Bob" do mapa

fmt.Println(pessoas) // map[Alice:25 Carlos:40]
```

📌 **Se a chave não existir, `delete()` não gera erro, apenas não faz nada.**

✅ **Removendo em um loop:**

```go
for k := range pessoas {
    delete(pessoas, k) // Remove todos os elementos
}
fmt.Println(len(pessoas)) // 0
```

📌 **Cuidado ao modificar mapas enquanto os itera, pois isso pode gerar comportamentos inesperados.**

---

## **6.2.2 Obtendo o Tamanho do Mapa com `len()`**

A função `len()` retorna o número total de pares **chave-valor** armazenados no mapa:

```go
fmt.Println(len(pessoas)) // 2
```

✅ **Após remover elementos, `len()` reflete o novo tamanho:**

```go
delete(pessoas, "Alice")
fmt.Println(len(pessoas)) // 1
```

📌 **O tamanho de um mapa pode mudar dinamicamente conforme elementos são adicionados ou removidos.**

---

## **6.2.3 Iterando Sobre Mapas com `range`**

Podemos percorrer um mapa usando `range`, acessando **chaves e valores** diretamente:

```go
pessoas := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carlos": 40,
}

for nome, idade := range pessoas {
    fmt.Println(nome, "tem", idade, "anos")
}
```

Saída:

```
Alice tem 25 anos
Bob tem 30 anos
Carlos tem 40 anos
```

📌 **A ordem de iteração não é garantida!**  
Se precisarmos percorrer um mapa em **ordem alfabética**, devemos **ordenar as chaves manualmente**.

✅ **Ordenando as chaves antes de iterar:**

```go
var chaves []string
for k := range pessoas {
    chaves = append(chaves, k)
}
sort.Strings(chaves) // Ordena alfabeticamente

for _, k := range chaves {
    fmt.Println(k, "->", pessoas[k])
}
```

📌 **Isso é necessário porque mapas em Go são implementados como tabelas de hash, e a ordem dos elementos pode variar.**

---

## **6.2.4 Boas Práticas e Considerações**

✔ **Use `delete()` para remover chaves, mas evite modificar um mapa enquanto o percorre.**  
✔ **Sempre verifique se uma chave existe antes de acessá-la (`val, ok := mapa[chave]`).**  
✔ **Se precisar de ordem nos elementos, armazene as chaves em um slice separado e ordene-as.**  
✔ **Evite mapas muito grandes se precisar de acesso ordenado frequente — slices podem ser mais eficientes nesses casos.**  

---

## **Conclusão**

As operações comuns de mapas permitem manipular dados de forma rápida e eficiente.  
No próximo capítulo, abordaremos **structs e métodos**, que permitem definir tipos complexos e suas operações! 🚀
