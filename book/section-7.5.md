# **7.5 Anatomia do Garbage Collector do Go**

O **Garbage Collector (GC)** do Go é um dos principais responsáveis pelo gerenciamento automático de memória, garantindo que a memória não utilizada seja liberada sem intervenção manual do programador.

Nesta seção, exploraremos:

- O que é um Garbage Collector e como ele funciona
- Como o GC do Go gerencia memória automaticamente
- Estratégias de otimização e impacto no desempenho
- Como monitorar e ajustar o GC para aplicações de alta performance

---

## **7.5.1 O Que é um Garbage Collector?**

Um **Garbage Collector** é um mecanismo que **automaticamente libera memória** alocada que não está mais sendo utilizada pelo programa.

📌 **Por que usar GC?**

- Evita vazamentos de memória (`memory leaks`)
- Facilita o gerenciamento de memória em tempo de execução
- Reduz a complexidade do código eliminando `malloc()` e `free()` (C-style)

---

## **7.5.2 Como Funciona o Garbage Collector do Go?**

O GC do Go é **concurrent** e **incremental**, minimizando pausas na execução do programa. Ele funciona em três fases:

1️⃣ **Marcação (`Mark`)**: Encontra objetos vivos que ainda estão sendo referenciados.  
2️⃣ **Varredura (`Sweep`)**: Libera a memória ocupada por objetos não utilizados.  
3️⃣ **Compactação (`Compaction`)**: Em algumas situações, reorganiza a memória para melhorar o desempenho.

🔎 **Visualizando o funcionamento do GC:**

```
[Alocação] ----> [Marcação] ----> [Varredura] ----> [Memória Liberada]
```

📌 **Objetos na stack são liberados automaticamente quando saem do escopo.**  
📌 **Objetos na heap são gerenciados pelo GC.**

---

## **7.5.3 Quando o Garbage Collector é Acionado?**

O GC do Go roda de forma **automática** sempre que necessário, mas podemos **forçar sua execução** manualmente:

```go
import "runtime"

runtime.GC() // Força a coleta de lixo
```

📌 **Forçar o GC pode ser útil para liberar memória imediatamente, mas pode impactar o desempenho.**

---

## **7.5.4 Monitorando o Garbage Collector**

Podemos medir o impacto do GC usando `runtime.ReadMemStats`:

```go
var memStats runtime.MemStats
runtime.ReadMemStats(&memStats)

fmt.Println("Memória Alocada:", memStats.Alloc)
fmt.Println("Número de Coletas:", memStats.NumGC)
```

✅ **Também podemos usar `pprof` para analisar o consumo de memória:**

```sh
go tool pprof -alloc_space ./binário
```

---

## **7.5.5 Otimizando o Uso do GC**

✔ **Prefira variáveis de curta duração** para evitar pressão na heap.  
✔ **Evite criar muitos objetos dinâmicos dentro de loops.**  
✔ **Use `sync.Pool` para reutilizar objetos e reduzir alocações.**  
✔ **Se possível, trabalhe com slices pré-alocados (`make([]T, tamanho)`).**  

✅ **Exemplo de uso de `sync.Pool` para reduzir pressão no GC:**

```go
import "sync"

var pool = sync.Pool{
    New: func() any { return new([]byte) },
}

buf := pool.Get().(*[]byte)
pool.Put(buf) // Devolve para o pool
```

📌 **Isso reduz a quantidade de objetos novos criados e melhora o desempenho.**

---

## **7.5.6 Comparação com Outros GCs**

| Característica | Go | Java | C++ (sem GC) |
|---------------|----|------|--------------|
| Coleta Automática | ✅ | ✅ | ❌ |
| Tempo de Pausa | Curto | Longo | N/A |
| Eficiência | Alta | Média | Alta |
| Controle Manual | Parcial | ❌ | ✅ |

📌 **O GC do Go é otimizado para baixa latência, enquanto o do Java pode causar pausas mais longas.**

---

## **Conclusão**

O **Garbage Collector do Go** fornece uma abordagem eficiente para gerenciamento de memória, permitindo que os desenvolvedores foquem na lógica do programa sem se preocupar com alocação manual.  
No próximo capítulo, entraremos em **programação orientada a objetos em Go**, abordando métodos e interfaces! 🚀
