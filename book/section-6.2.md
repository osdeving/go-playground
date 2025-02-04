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

## **Pratique Go**

🎯 Agora que você aprendeu sobre operações comuns em mapas, tente estes desafios de nível sênior:

🛠️ **Desafios Avançados**:

<details>
  <summary>1️⃣ Implemente um sistema de limpeza automática de cache que remove entradas antigas.</summary>
  
  ```go
  type Cache struct {
      data     map[string]interface{}
      lastUsed map[string]time.Time
      maxAge   time.Duration
      mu       sync.RWMutex
  }

  func (c *Cache) cleanup() {
      ticker := time.NewTicker(time.Minute)
      for range ticker.C {
          c.mu.Lock()
          now := time.Now()
          for k, t := range c.lastUsed {
              if now.Sub(t) > c.maxAge {
                  delete(c.data, k)
                  delete(c.lastUsed, k)
              }
          }
          c.mu.Unlock()
      }
  }
  ```
</details>

<details>
  <summary>2️⃣ Desenvolva um contador de referências para gerenciar recursos compartilhados.</summary>
  
  ```go
  type RefCounter struct {
      refs map[string]int
      mu   sync.RWMutex
  }

  func (rc *RefCounter) Increment(key string) int {
      rc.mu.Lock()
      defer rc.mu.Unlock()
      rc.refs[key]++
      return rc.refs[key]
  }

  func (rc *RefCounter) Decrement(key string) int {
      rc.mu.Lock()
      defer rc.mu.Unlock()
      if rc.refs[key] > 0 {
          rc.refs[key]--
          if rc.refs[key] == 0 {
              delete(rc.refs, key)
          }
      }
      return rc.refs[key]
  }
  ```
</details>

<details>
  <summary>3️⃣ Crie um sistema de histórico de alterações com rollback.</summary>
  
  ```go
  type History struct {
      current  map[string]interface{}
      versions []map[string]interface{}
      maxSize  int
  }

  func (h *History) Snapshot() {
      if len(h.versions) >= h.maxSize {
          h.versions = h.versions[1:]
      }
      snapshot := make(map[string]interface{})
      for k, v := range h.current {
          snapshot[k] = v
      }
      h.versions = append(h.versions, snapshot)
  }
  ```
</details>

<details>
  <summary>4️⃣ Implemente um sistema de índice invertido com suporte a remoção eficiente.</summary>
  
  ```go
  type InvertedIndex struct {
      index map[string]map[int]struct{}
      docs  map[int][]string
  }

  func (idx *InvertedIndex) Remove(docID int) {
      for _, word := range idx.docs[docID] {
          delete(idx.index[word], docID)
          if len(idx.index[word]) == 0 {
              delete(idx.index, word)
          }
      }
      delete(idx.docs, docID)
  }
  ```
</details>

<details>
  <summary>5️⃣ Desenvolva um sistema de contagem de frequência com limite de memória.</summary>
  
  ```go
  type BoundedCounter struct {
      counts    map[string]int
      maxItems  int
      minCount  int
  }

  func (bc *BoundedCounter) Increment(key string) {
      if len(bc.counts) >= bc.maxItems && bc.counts[key] == 0 {
          bc.removeLowestCounts()
      }
      bc.counts[key]++
  }
  ```
</details>

<details>
  <summary>6️⃣ Crie um gerenciador de pool de objetos com cleanup automático.</summary>
  
  ```go
  type ObjectPool struct {
      objects map[string]interface{}
      inUse   map[string]time.Time
      timeout time.Duration
  }

  func (op *ObjectPool) Cleanup() {
      now := time.Now()
      for id, lastUse := range op.inUse {
          if now.Sub(lastUse) > op.timeout {
              delete(op.objects, id)
              delete(op.inUse, id)
          }
      }
  }
  ```
</details>

<details>
  <summary>7️⃣ Implemente um sistema de cache em camadas com diferentes tempos de expiração.</summary>
  
  ```go
  type LayeredCache struct {
      l1    map[string]interface{}
      l2    map[string]interface{}
      ttl1  time.Duration
      ttl2  time.Duration
      times map[string]time.Time
  }

  func (lc *LayeredCache) Get(key string) interface{} {
      if v, ok := lc.l1[key]; ok {
          return v
      }
      if v, ok := lc.l2[key]; ok {
          lc.promoteToL1(key, v)
          return v
      }
      return nil
  }
  ```
</details>

<details>
  <summary>8️⃣ Desenvolva um sistema de agregação de eventos com janela deslizante.</summary>
  
  ```go
  type WindowAggregator struct {
      events   map[time.Time][]Event
      window   time.Duration
      callback func([]Event)
  }

  func (wa *WindowAggregator) Add(e Event) {
      now := time.Now()
      wa.events[now] = append(wa.events[now], e)
      wa.cleanup(now)
  }
  ```
</details>

<details>
  <summary>9️⃣ Crie um sistema de roteamento de mensagens com prioridade.</summary>
  
  ```go
  type MessageRouter struct {
      routes    map[string][]Handler
      priority  map[Handler]int
      mu        sync.RWMutex
  }

  func (mr *MessageRouter) Route(msg Message) {
      handlers := mr.routes[msg.Type]
      sort.Slice(handlers, func(i, j int) bool {
          return mr.priority[handlers[i]] > mr.priority[handlers[j]]
      })
  }
  ```
</details>

<details>
  <summary>🔟 Implemente um sistema de cache distribuído com invalidação.</summary>
  
  ```go
  type DistCache struct {
      local     map[string]interface{}
      peers     map[string]*Client
      versions  map[string]int64
      mu        sync.RWMutex
  }

  func (dc *DistCache) Invalidate(key string) {
      dc.mu.Lock()
      defer dc.mu.Unlock()
      delete(dc.local, key)
      for _, peer := range dc.peers {
          go peer.Notify(InvalidateEvent{Key: key})
      }
  }
  ```
</details>

## **Perguntas e Respostas**

❓ **Teste seus conhecimentos:**

<details>
  <summary>1️⃣ Qual é o comportamento do delete() quando a chave não existe no mapa?</summary>
  O delete() é seguro para usar mesmo quando a chave não existe, não gerando nenhum erro ou pânico.
</details>

<details>
  <summary>2️⃣ Como o len() se comporta com mapas nil?</summary>
  len() retorna 0 para mapas nil, assim como para mapas vazios inicializados.
</details>

<details>
  <summary>3️⃣ Por que devemos ter cuidado ao modificar mapas durante iteração?</summary>
  Modificar um mapa durante a iteração pode resultar em comportamento indefinido, podendo pular ou repetir elementos.
</details>

<details>
  <summary>4️⃣ Como podemos garantir uma ordem consistente ao iterar sobre um mapa?</summary>
  Extraindo as chaves para um slice, ordenando-o e então usando-o para acessar os valores do mapa em ordem.
</details>

<details>
  <summary>5️⃣ Qual é a complexidade de tempo do delete() em mapas?</summary>
  O delete() tem complexidade O(1) em média, mas pode variar dependendo das colisões de hash.
</details>

<details>
  <summary>6️⃣ Como o garbage collector lida com elementos deletados de um mapa?</summary>
  O GC pode recuperar a memória dos valores deletados, mas o mapa mantém sua capacidade interna até ser realocado.
</details>

<details>
  <summary>7️⃣ Qual é a diferença entre deletar uma chave e atribuir nil ao seu valor?</summary>
  Deletar remove a entrada completamente do mapa, enquanto atribuir nil mantém a chave com um valor nil.
</details>

<details>
  <summary>8️⃣ Como podemos otimizar operações de delete em massa em um mapa?</summary>
  Para muitas deleções, pode ser mais eficiente criar um novo mapa e copiar apenas os elementos desejados.
</details>

<details>
  <summary>9️⃣ Por que range retorna uma cópia dos valores do mapa?</summary>
  Para evitar modificações acidentais dos valores originais e garantir consistência durante a iteração.
</details>

<details>
  <summary>🔟 Como podemos implementar um contador de referências usando delete?</summary>
  Decrementando o contador e usando delete quando chegar a zero, removendo completamente a entrada do mapa.
</details>

---

## **Conclusão**

As operações comuns de mapas permitem manipular dados de forma rápida e eficiente.  
No próximo capítulo, abordaremos **structs e métodos**, que permitem definir tipos complexos e suas operações! 🚀
