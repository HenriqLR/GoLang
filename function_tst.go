package main

import "fmt"

//Uso de Constantes
func constante() string {
	const World = "abc"
	return World
}

//Calculos Simples
func soma(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(soma(10, 30))
	fmt.Println(constante())
}
