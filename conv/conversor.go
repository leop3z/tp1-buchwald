package conv

import (
	"bufio"
	"os"
	"tdas/cola"
	"tdas/pila"
	"unicode"
)

// constantes para caracteres

const (
	SUMA              = "+"
	RESTA             = "-"
	MULTIPLICACION    = "*"
	DIVISION          = "/"
	POTENCIA          = "^"
	ABRE_PARENTESIS   = "("
	CIERRE_PARENTESIS = ")"
)

func LeerArchivo() string {
	scanner := bufio.NewScanner(os.Stdin)
	var exprecion string

	for scanner.Scan() {
		exprecion = scanner.Text()
	}
	return exprecion
}

func procesarLinea(linea string) string {
	postfija := InfijaAPostfija(linea)
	return postfija
}

func esOperador(token string) bool {
	return token == SUMA || token == RESTA || token == MULTIPLICACION || token == DIVISION || token == POTENCIA
}

func precedencia(operador string) int {
	switch operador {
	case SUMA, RESTA:
		return 1
	case MULTIPLICACION, DIVISION:
		return 2
	case POTENCIA:
		return 3
	default:
		return 0
	}
}

func tokenizar(expresion string) []string {
	var tokens []string
	var numero string = ""

	for _, caracter := range expresion {
		if unicode.IsSpace(caracter) {
			continue
		}
		if unicode.IsDigit(caracter) {
			numero += string(caracter)
		} else {
			if numero != "" {
				tokens = append(tokens, numero)
				numero = ""
			}
			tokens = append(tokens, string(caracter))
		}
	}
	if numero != "" {
		tokens = append(tokens, numero)
	}
	return tokens
}

func mostrarPostfija(resultado cola.Cola[string]) string {
	caracter := ""
	for !resultado.EstaVacia() {
		caracter += resultado.Desencolar()
	}
	return caracter
}

func InfijaAPostfija(expresion string) string {
	salida := cola.CrearColaEnlazada[string]()
	operadores := pila.CrearPilaDinamica[string]()
	tokens := tokenizar(expresion)
	for _, token := range tokens {
		procesarToken(token, salida, operadores)
	}
	vaciarOperadores(salida, operadores)
	return mostrarPostfija(salida)
}

func procesarToken(token string, salida cola.Cola[string], operadores pila.Pila[string]) {
	if unicode.IsDigit(rune(token[0])) {
		salida.Encolar(token)
	} else if token == ABRE_PARENTESIS {
		operadores.Apilar(token)
	} else if token == CIERRE_PARENTESIS {
		procesarParentesisCierre(salida, operadores)
	} else if esOperador(token) {
		procesarOperador(token, salida, operadores)
	}
}

func procesarParentesisCierre(salida cola.Cola[string], operadores pila.Pila[string]) {
	for !operadores.EstaVacia() && operadores.VerTope() != ABRE_PARENTESIS {
		salida.Encolar(operadores.Desapilar())
	}
	if !operadores.EstaVacia() {
		operadores.Desapilar() // saco el (
	}
}

func procesarOperador(token string, salida cola.Cola[string], operadores pila.Pila[string]) {
	for !operadores.EstaVacia() && (precedencia(operadores.VerTope()) > precedencia(token) ||
		(precedencia(operadores.VerTope()) == precedencia(token) && token != POTENCIA)) {
		salida.Encolar(operadores.Desapilar())
	}
	operadores.Apilar(token)
}

func vaciarOperadores(salida cola.Cola[string], operadores pila.Pila[string]) {
	for !operadores.EstaVacia() {
		salida.Encolar(operadores.Desapilar())
	}
}
