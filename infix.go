package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/conv"
)

func main() {
<<<<<<< HEAD
	infija := conv.LeerArchivo()
	postfija := conv.InfijaAPostfija(infija)
	fmt.Println(postfija)
=======
	scanner := bufio.NewScanner(os.Stdin)
	var linea string
	for scanner.Scan() {
		linea = scanner.Text()
		postfija := conv.InfijaAPostfija(linea)
		fmt.Println(postfija)
	}
>>>>>>> a8fd911 (lectura y procesamiento correcto de los archivos)
}
