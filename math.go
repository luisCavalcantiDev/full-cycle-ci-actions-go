package main

import "fmt"

//metodo que executa soma de dois numeros
func main() {
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}
