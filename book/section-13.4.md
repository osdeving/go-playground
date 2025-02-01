# **13.4 Tratamento de Erros (`errors`, `fmt.Errorf`)**

O tratamento de erros é uma parte essencial do desenvolvimento em Go.  
Diferente de linguagens que utilizam exceções (`try/catch`), o Go usa um modelo baseado em **valores de erro explícitos**,  
o que torna o código mais previsível e seguro.

Nesta seção, exploraremos:

- O modelo de tratamento de erros em Go
- Como usar o pacote `errors` para criar e comparar erros
- Uso de `fmt.Errorf` para formatar mensagens de erro
- Como encapsular erros e adicionar contexto
- Estratégias para escrever código Go robusto

---

## **13.4.1 O Modelo de Erros em Go**

Diferente de linguagens como Java e Python, onde erros são tratados com exceções (`throw/catch`),  
Go trata erros **como valores de retorno convencionais**.

✅ **Exemplo básico de tratamento de erro:**

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero não é permitida")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    fmt.Println("Resultado:", result)
}
```

📌 **O erro é retornado como o segundo valor e deve ser sempre verificado antes de prosseguir.**  
📌 **Se `err == nil`, significa que a operação foi bem-sucedida.**  

---

## **13.4.2 Criando Erros com `errors.New()`**

O pacote `errors` fornece a função `errors.New()` para criar erros simples.  

✅ **Exemplo: Criando um erro e comparando com `errors.Is()`**

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("registro não encontrado")

func findUser(id int) error {
    if id != 1 {
        return ErrNotFound
    }
    return nil
}

func main() {
    err := findUser(2)
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Usuário não encontrado!")
    }
}
```

📌 **Criar erros como variáveis globais (`var Err...`) facilita comparações e evita erros duplicados.**  
📌 **O método `errors.Is()` permite verificar a causa raiz do erro.**  

---

## **13.4.3 Formatando Erros com `fmt.Errorf()`**

A função `fmt.Errorf()` permite criar erros formatados, adicionando contexto ao erro original.

✅ **Exemplo: Formatando mensagens de erro**

```go
package main

import (
    "fmt"
)

func openFile(filename string) error {
    return fmt.Errorf("erro ao abrir o arquivo %s: arquivo não encontrado", filename)
}

func main() {
    err := openFile("data.txt")
    fmt.Println(err)
}
```

📌 **O erro contém contexto útil sobre a operação falha.**  

✅ **Adicionando erro original com `%w` (error wrapping)**

```go
package main

import (
    "errors"
    "fmt"
)

var ErrPermissionDenied = errors.New("permissão negada")

func openRestrictedFile() error {
    return fmt.Errorf("erro crítico: %w", ErrPermissionDenied)
}

func main() {
    err := openRestrictedFile()
    if errors.Is(err, ErrPermissionDenied) {
        fmt.Println("Ação não permitida!")
    }
}
```

📌 **O `%w` permite que `errors.Is()` identifique a causa raiz do erro encapsulado.**  

---

## **13.4.4 Lidando com Erros em Funções Encadeadas**

Em funções que chamam outras funções, é comum **propagar erros** em vez de tratá-los imediatamente.

✅ **Exemplo: Propagando erros corretamente**

```go
package main

import (
    "fmt"
    "os"
)

func readFile(name string) error {
    file, err := os.Open(name)
    if err != nil {
        return fmt.Errorf("erro ao abrir arquivo: %w", err)
    }
    defer file.Close()
    return nil
}

func main() {
    err := readFile("inexistente.txt")
    if err != nil {
        fmt.Println("Erro detectado:", err)
    }
}
```

📌 **Os erros são propagados com `return fmt.Errorf()`, mantendo o contexto.**  

✅ **Usando `errors.Unwrap()` para obter a causa raiz**

```go
origErr := fmt.Errorf("erro original")
wrappedErr := fmt.Errorf("erro adicional: %w", origErr)

fmt.Println(errors.Unwrap(wrappedErr)) // Retorna o erro original
```

📌 **`errors.Unwrap()` ajuda a depurar erros encadeados.**  

---

## **13.4.5 Estratégias para Boas Práticas**

✔ **Sempre retorne erros em operações que possam falhar.**  
✔ **Use variáveis de erro globais (`var ErrSomething = errors.New(...)`).**  
✔ **Encapsule erros para adicionar contexto (`fmt.Errorf("erro ao carregar: %w", err)`).**  
✔ **Evite panics, a menos que seja realmente um erro crítico.**  
✔ **Documente os erros retornados pelas funções (`// Retorna ErrNotFound se não existir`).**  

---

## **Conclusão**

O **tratamento de erros em Go é explícito e previsível**, garantindo **código mais seguro e testável**.  
No próximo capítulo, exploraremos **programação de redes com TCP e UDP**, aplicando tratamento de erros em comunicações distribuídas! 🚀
