package main

import "fmt"

// agrupando declarações de variáveis
var (
	i int     = 42
	f float64 = 2.3
	b bool    = true
	c uint8
	s string = "Hello, World!"
)

// agrupando declarações de variáveis
var x, y int = 30, 50

// infere o tipo da variável
var x1 = 30
var y1 = 50

const (
	// declaração de constantes
	Const1 = "Constante 1"
	Const2 = 3.14
)

// declaração de constante com tipo
const Constante int = 10

// define um tipo chamado TextoPuro para ser usado como string
type TextoPuro string

func VarTest() {
	// declara e inicializa uma variável do tipo TextoPuro em escopo local
	var a TextoPuro = "Hello"

	// declara e inicializa uma variável em escopo local usando inferência de tipo (não precisa, nesse caso, de var e do tipo)
	w := 3.4

	println(w)
	println(a)
	println(i)
	println(f)
	println(b)
	println(c)
	println(s)

	fmt.Printf("%v, %T\n", a, a)

}
