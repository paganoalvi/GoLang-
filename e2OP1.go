package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainO2() {
	// creo lector para leer frase

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase) // strings.TrimSpace quita el salto de linea final
	fraseFinal := ejercicio2(frase)
	fmt.Println(fraseFinal)
}

func ejercicio2(frase string) string {
	// divido la frase en un slice de palabras
	palabras := strings.Fields(frase)

	// armo slice nuevo para guardar palabras modificadas
	var resultado []string

	for i, palabra := range palabras {
		if i%2 == 0 { // las posiciones impares "reales" (1,3,5..) son indices pares en Go(indice empiezan en 0)
			palabraInvertida := InvertirPalabra(palabra)
			resultado = append(resultado, palabraInvertida)
		} else {
			resultado = append(resultado, palabra)
		}
	}

	return strings.Join(resultado, " ")
}

// funcion auxiliar invertirPalabra
func InvertirPalabra(palabra string) string {
	runes := []rune(palabra) // convierto palabra a runa para manejar correctamente caracteres especiales
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/* explicacion libreria strings:
strings.Fields(frase) split a string into a slice of substrings splitting identifying whitespaces
[] strings (slice of strings)


Apend([]slice,string)
append() function is used to add elements to a slice. When appending a string to a slice of strings ([]string),
 the string is treated as a single element to be added to the end of the slice.
*/
