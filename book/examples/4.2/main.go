package main

import "fmt"

func main() {
    name, age := userInfo(1)
    fmt.Println(name, age)

    names := make([]string, 3)
    names[0] = "Alice"
    names[1] = "Bob"
    names[2] = "Charlie"

    logValues(names...)
}

func userInfo(id int) (name string, age int) {
    if id == 1 {
        name, age = "Alice", 30
    } else {
        name, age = "Unknown", 0
    }
    return // Retorno implícito das variáveis nomeadas
}

func logValues[T any](values ...T) {
    for _, value := range values {
        fmt.Print(", value: ", value)
    }
}