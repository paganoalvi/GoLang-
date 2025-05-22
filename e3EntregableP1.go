// Tengo que recibir una palabra, pedir el ingreso de una frase por teclado,
// Luego tengo que imprimir la frase ingresada pero cambiando todas las ocurrencias de la palabra
// invirtiendo todas las letras de minusculas a mayusculas y de mayusculas a minusculas
// paso 1: convertir la palabra a runa para manejar bien caracteres especiales
// paso 2: buscar la ocurrencia de la palabra en la frase e invertir de mayuscula a minuscula y de minuscula a mayuscula
// paso 3: remplazar en la frase
// por ultimo imprimir la frase

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	principal("PeQueÑo") // envio la palabra PeQueÑo a la funcion principal
}

func principal(palabra string) {
	reader := bufio.NewReader(os.Stdin) // creo reader para solicitar ingreso de frase
	fmt.Println("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n') // guardo la frase ingresada en frase
	

	fraseFinal := helper(frase, palabra) // guardo el resultado de ejecutar helper en fraseFinal
	fmt.Println("Frase final: ", fraseFinal) // imprimo fraseFinal
}

func helper(frase string, objetivo string) string {
	palabras := strings.Fields(frase) // divide a la frase en palabras dentro de un slice
	fmt.Println("palabras: ",palabras) // veo el contenido de palabras luego de dividir en palabras
	for i, palabra := range palabras { // i → es el índice (0, 1, 2…) , palabra → es el valor actual en esa posición del slice (palabra = palabras[i])
		if strings.EqualFold(palabra, objetivo) { //strings.EqualFold compara dos cadenas ignorando mayúsculas y minúsculas asi me ahorro convertir toda la frase a minuscula
			palabras[i] = invertirMayMin(palabra)
		}
	}
	return strings.Join(palabras, " ")  //Une todas las palabras con un espacio entre cada una
}


func invertirMayMin(palabra string) string {
	runes := []rune(palabra)
	for i, letra := range runes {
		if unicode.IsUpper(letra) {
			runes[i] = unicode.ToLower(letra)
		} else {
			runes[i] = unicode.ToUpper(letra)
		}
	}
	return string(runes)
}

