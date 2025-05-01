package main

import (
	"fmt"
	"tp1/conv"
)

func main() {
	expresiones := conv.LeerArchivo()

	for _, expresion := range expresiones {
		postfija := conv.InfijaAPostfija(expresion)
		for !postfija.EstaVacia() {
			fmt.Print(postfija.Desencolar(), " ")
		}
		fmt.Println()
	}
}
