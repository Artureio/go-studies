package main

import (
	"fmt"
	"strconv"
)

// TODO: 1- pedir os numeros para o usuario:
func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println("Soma: " + strconv.Itoa(soma(i, j)))
			fmt.Println("Divisão: " + strconv.Itoa(divisao(i, j)))
			fmt.Println("Subtração: " + strconv.Itoa(subtracao(i, j)))
			fmt.Println("Multiplicação: " + strconv.Itoa(multiplicacao(i, j)))
		}
	}
}

func soma(i, j int) int {
	return i + j
}
func divisao(i, j int) int {
	if j == 0 {
		return j
	}
	return i / j
}
func subtracao(i, j int) int {
	return i - j
}
func multiplicacao(i, j int) int {
	return i * j
}
