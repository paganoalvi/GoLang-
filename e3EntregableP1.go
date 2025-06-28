package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Falta ingresar la palabra argumento")
		return
	}
	objetivo := os.Args[1]

	reader := bufio.NewReader(os.Stdin) // creamos lector para ingreso de frase desde entrada estandar
	fmt.Println("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println("Frase ingresada: ", frase)

	fraseFinal := helper(frase, objetivo)
	fmt.Println("Frase final: ", fraseFinal)
}

func helper(frase string, objetivo string) string {
	palabras := strings.Fields(frase) // convertimos frase en slice de palabras
	for i, palabra := range palabras {
		prefijo, limpia, sufijo := separar(palabra) // limpiamos de caracteres especiales la palabra objetivo contenida en slice palabras
		if strings.EqualFold(limpia, objetivo) {    // comparamos dos strings (limpia y objetivo) carácter por carácter, no importa si son mayusculas o minusculas
			limpia = invertirMayMin(limpia)
		}
		palabras[i] = prefijo + limpia + sufijo // vuelvo a concatenar(no append, por que es la misma posicion del slice)
	}
	return strings.Join(palabras, " ") // convierto slice a frase con palabras separados por " "
}

// Esta función separa prefijos y sufijos no alfabéticos (simbolos que esten antes de la palabra o despues)
func separar(palabra string) (prefijo, central, sufijo string) {
	runes := []rune(palabra) // conviertimos palabra en runa de letras

	i := 0
	for i < len(runes) && !unicode.IsLetter(runes[i]) { // busco hasta encontrar un caracter no char
		i++
	}
	j := len(runes) - 1
	for j >= 0 && !unicode.IsLetter(runes[j]) {
		j--
	}

	// Validamos rangos por si la palabra no tiene letras
	if i > j {
		return palabra, "", ""
	}

	prefijo = string(runes[:i])
	central = string(runes[i : j+1])
	sufijo = string(runes[j+1:])
	return
}

// Esta funcion invierte las mayusculas o minusculas de la palabra objetivo ya limpia
func invertirMayMin(palabra string) string {
	runes := []rune(palabra)
	for i, letra := range runes {
		if unicode.IsUpper(letra) {
			runes[i] = unicode.ToLower(letra)
		} else {
			runes[i] = unicode.ToUpper(letra)
		}
	}
	return string(runes) // retorno runa de letras convertida a string de chars
}

/* os.Args
os.Args is a variable within the os package that provides access to the command-line arguments passed to a program.
 It is a slice of strings ([]string)
- first element of the os.Args slice, os.Args[0], always contains the path or name of the program itself
- Any arguments provided by the user starting from os.Args[1]
- os.Args provides the raw, unparsed command-line arguments as strings. For more advanced argument parsing with flags
 and options, the flag package is typically used
*/

/* strings.Fields
Split a string around one or more consecutive whitespace characters, returning a slice of substrings
*/

/*strings.EqualFold
Compare two strings for equality in a case-insensitive manner
- It performs a comparison based on Unicode case-folding, which is a more generalized form of case-insensitivity than
 simply converting to lowercase
- It returns true if the strings are equal under this Unicode case-folding, and false otherwise
*/

/*strings.Join
Concatenate elements of a string slice into a single string, using a specified separator between each element
- concatenate elements of a string slice into a single string, using a specified separator between each element
- strings.Join(palabra,"|") | => separetor
*/
