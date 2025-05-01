package main

import (
	"fmt"
	"tp1/conv"
)

func main() {
	infija := conv.LeerArchivo()
	postfija := conv.InfijaAPostfija(infija)
	fmt.Println(postfija)
}
