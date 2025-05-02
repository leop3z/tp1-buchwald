package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/conv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var linea string
	for scanner.Scan() {
		linea = scanner.Text()
		postfija := conv.InfijaAPostfija(linea)
		fmt.Println(postfija)
	}
}
