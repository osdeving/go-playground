package main

import "fmt"


// Testando operadores bit a bit do Go
type WindowStyle uint32

const (
    WS_OVERLAPPED    WindowStyle = 1 << iota  // 0x00000001
    WS_POPUP                                  // 0x00000002  
    WS_CHILD                                  // 0x00000004
    WS_MINIMIZE                               // 0x00000008
    WS_VISIBLE                                // 0x00000010
    WS_DISABLED                               // 0x00000020
    WS_CLIPSIBLINGS                           // 0x00000040
    WS_CLIPCHILDREN                           // 0x00000080
    WS_MAXIMIZE                               // 0x00000100
    WS_CAPTION                                // 0x00000200
)

func main() {
    // Exemplo de combinação de estilos de janela
    var janelaPadrao = WS_OVERLAPPED | WS_CAPTION // | WS_SYSMENU | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
    var janelaPopup = WS_POPUP // | WS_BORDER | WS_SYSMENU
    var janelaFilha = WS_CHILD | WS_VISIBLE // | WS_BORDER

    createWindow(janelaPadrao)
    createWindow(janelaPopup)
    createWindow(janelaFilha)

    // Removendo um estilo específico
    janelaFilhaSemVisibilidade := janelaFilha &^ WS_VISIBLE
    
    createWindow(janelaFilhaSemVisibilidade)

    // Exemplo de arredondamento usando operadores bit a bit
    fmt.Println("\nExemplo de arredondamento para múltiplos:")
    
    // Função que arredonda para o próximo múltiplo usando máscaras de bits
    roundUp := func(n int, multiple int) int {
        mask := multiple - 1
        return (n + mask) &^ mask
    }

    // Demonstração com diferentes números e múltiplos
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println("Arredondando para múltiplos de 4:")
    for _, n := range numbers {
        rounded := roundUp(n, 4)
        fmt.Printf("%2d -> %2d\n", n, rounded)
    }

    fmt.Println("\nArredondando para múltiplos de 8:")
    for _, n := range numbers {
        rounded := roundUp(n, 8)
        fmt.Printf("%2d -> %2d\n", n, rounded)
    }

    // Exemplo prático: alocação de memória em blocos
    fmt.Println("\nSimulando alocação de memória em blocos:")
    memorySizes := []int{3, 7, 12, 15, 17}
    for _, size := range memorySizes {
        aligned := roundUp(size, 8)
        fmt.Printf("Requisição: %2d bytes -> Alocado: %2d bytes (alinhado em 8)\n", 
            size, aligned)
    }


    
    fmt.Println("\nOperações matemáticas com múltiplos retornos:")
    
    calc := func(a,b int) (sum int, sub int, mul int, div float64) {
        sum = a + b
        sub = a - b
        mul = a * b
        div = float64(a) / float64(b)
        return        
    }
    
    a, b := 202, 39
    sum, sub, mul, div := calc(a, b)
    
    fmt.Printf("\nOperações entre %d e %d:\n", a, b)
    fmt.Printf("Soma:          %d\n", sum)
    fmt.Printf("Subtração:     %d\n", sub) 
    fmt.Printf("Multiplicação: %d\n", mul)
    fmt.Printf("Divisão:       %.2f\n", div)
    
}

func createWindow(style WindowStyle) {
    if style & WS_OVERLAPPED != 0 {
        fmt.Println("- Janela sobreposta")
    }

    if style & WS_POPUP != 0 {
        fmt.Println("- Janela popup") 
    }
   
    if style & (WS_CHILD | WS_VISIBLE) == (WS_CHILD | WS_VISIBLE) {
        fmt.Println("- Janela filha visível")
    }
}
