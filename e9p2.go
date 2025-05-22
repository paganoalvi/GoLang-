package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una frase:")
	frase, _ := reader.ReadString('\n')
	ejercicio9(frase)
}

func ejercicio9(frase string) { // frase → es la versión original, con las mayúsculas reales que necesito copiar
	fraseMin := strings.ToLower(frase) //fraseMin → sirve para buscar de forma insensible a mayúsculas/minúsculas (todo minúscula).
	var builder strings.Builder
	i := 0

	for i < len(fraseMin) {  // Si en la posición actual empieza "jueves"
		if i+6 <= len(fraseMin) && fraseMin[i:i+6] == "jueves" {  // con i+6 <= len(fraseMin) antes de mirar si hay "jueves" me fijo que no me pase del final de la frase
		// "jueves" tiene 6 letras, así que necesito que haya lugar para mirar desde i hasta i+6, sino error
		// con fraseMin[i:i+6] == "jueves" miro la subcadena de la posición i hasta i+6 
			palabraOriginal := frase[i : i+6] // me guardo  la palabra original para mirar mayúsculas

			// construyo "martes" respetando mayúsculas
			target := []rune("martes")    // target es un slice de runas (caracteres) que contiene la palabra "martes"  así podemos recorrer letra por letra fácilmente
			palabraNueva := make([]rune, len(target)) //Creamos un slice vacío del mismo tamaño que "martes" acá vamos a ir construyendo la nueva palabra, letra por letra

			for j, letra := range target { //Vamos a recorrer cada letra de "martes" (target) 
			// j es la posición (índice): 0, 1, 2, 3, 4, 5 
			// letra es el caracter: 'm', 'a', 'r', 't', 'e', 's'
				if unicode.IsUpper(rune(palabraOriginal[j])) { //acá miramos si la letra original en "palabraOriginal es mayuscula"
					palabraNueva[j] = unicode.ToUpper(letra) // si estaba en mayuscula, convierto letra a mayuscula
				} else {
					palabraNueva[j] = letra // sino la dejo como estaba
				}
			}

			builder.WriteString(string(palabraNueva)) // Convierto el slice palabraNueva a string (ya con la mayuscula)
			// string(palabraNueva) convierte ese []rune a un string normal
			// builder.WriteString(...) agrega ese string completo (por ejemplo, "marTes") al resultado final que estamos armando			
			i += 6 // Como procesamos 6 letras ("jueves" tiene 6 letras), avanzamos 6 posiciones para no volver a procesar.
		} else {
			builder.WriteByte(frase[i]) // Si no encontramos "jueves", simplemente copiamos la letra original como está
			i++  
		}
	}

	fmt.Println("Frase final:", builder.String())
}
