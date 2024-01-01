package main

const a = "Hello World!"

var b = "Hello World!"
var c bool
var d int

func main() {
	b = "Olá Mundo!"
	c = true

	// var a string = "X" // long declaration
	a := "X" // short declaration - first	declaration
	// criar, declarar e atribuir == a := "X"

	println(a)
	println(b)
	println(c)
}
