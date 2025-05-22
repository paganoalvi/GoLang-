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
	fmt.Println("Ingrese una FRASE:")
	frase, _ := reader.ReadString('\n')
	ejercicio1Entregable(frase)
}

func ejercicio1Entregable(frase string) { // frase → es la versión original, con las mayúsculas reales que necesito copiar
	fraseMin := strings.ToLower(frase) //fraseMin → sirve para buscar de forma insensible a mayúsculas/minúsculas (todo minúscula).
	var builder strings.Builder        // builder
	i := 0

	for i < len(fraseMin) { // Recorrer fraseMin buscando la ocurrencia de "miércoles"

		if i+10 <= len(fraseMin) && fraseMin[i:i+10] == "miércoles" { // con i+9 <= len(fraseMin) antes de mirar si encuentro "miércoles" me fijo que no me pase del final de la frase
			// "miércoles" tiene 9 letras, así que necesito que haya lugar para mirar desde i hasta i+9, sino error
			// con fraseMin[i:i+9] == "miércoles" miro la subcadena de la posición i hasta i+9
			palabraOriginal := frase[i : i+10] // me guardo  la palabra original para mirar mayúsculas

			// construyo "martes" respetando mayúsculas
			target := []rune("automóvil")             // target es un slice de runas (caracteres) que contiene la palabra "automóvil"  así podemos recorrer letra por letra fácilmente
			palabraNueva := make([]rune, len(target)) //Creamos un slice vacío del mismo tamaño que "automóvil" acá vamos a ir construyendo la nueva palabra, letra por letra

			for j, letra := range target { //Vamos a recorrer cada letra de "automóvil" (target)
				// j es la posición (índice): 0, 1, 2, 3, 4, 5, 6, 7, 8
				// letra es el caracter: 'a', 'u', 't', 'o', 'm','ó','v','i','l'
				if unicode.IsUpper(rune(palabraOriginal[j])) { //acá miramos si la letra original en "palabraOriginal es mayuscula"
					palabraNueva[j] = unicode.ToUpper(letra) // si estaba en mayuscula, convierto letra a mayuscula
				} else {
					palabraNueva[j] = letra // sino la dejo como estaba
				}
			}

			builder.WriteString(string(palabraNueva)) // Convierto el slice palabraNueva a string (ya con la mayuscula)
			// string(palabraNueva) convierte ese []rune a un string normal
			// builder.WriteString(...) agrega ese string completo (por ejemplo, "marTes") al resultado final que estamos armando
			i += 10 // Como procesamos 10 (9 mas el acento) letras ("miércoles" tiene 9 letras + el acento), avanzamos 10 posiciones para no volver a procesar.
		} else {
			builder.WriteByte(frase[i]) // Si no encontramos "miércoles", simplemente copiamos la letra original como está
			i++
		}
	}

	fmt.Println("Frase final:", builder.String())
}