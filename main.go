package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() { /*holaMundo() ejercicio3()ejercicio4()*/
	//fmt.Println("Ingrese valor de x: ")
	//x := 0
	//fmt.Scan(&x) /*fmt.Scan recibe la direccion de memoria de x (&x)*/
	//ejercicio5(x)

	reader := bufio.NewReader(os.Stdin) // creo un lector con bufio para pedir ingreso de una frase
	fmt.Println("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n') // porque el _ ??
	ejercicio9(frase)
}

func holaMundo() {
	fmt.Println("Hello world")
}
func ejercicio3() { /* integers */
	/*var zz int = 0*/ // se comenta esta linea porque zz nunca es usada
	x := 10
	z := x
	fmt.Println("z: ", z)
	fmt.Println("x: ", x)
	var y int = x + 1
	fmt.Println("y: ", y)
	const n = 5001
	fmt.Println("n: ", n)
	const c = 5001
	fmt.Println("c: ", c)
	/* float */
	var e float32 = 6
	fmt.Println("e: ", e)
	var f float32 = e
	fmt.Println("f: ", f)
}

func ejercicio4() {
	const tope = 250 /*Inciso a*/
	fmt.Println("Inciso a: ")
	suma := 0
	for i := 0; i <= tope; i = i + 2 {
		suma = suma + i
	}
	fmt.Println("La suma de todos los numeros pares entre 0 y 250 (inlcusive) es:", suma)
	/*Inciso b*/
	fmt.Println("Inciso b: ")
	sumab := 0
	for i := 250; i >= 0; i = i - 2 { /*alternativa => i -= 2*/
		sumab = sumab + i
	}
	fmt.Println("La suma de todos los numeros pares entre 0 y 250 (inclsuive) es:", sumab)
}

func ejercicio5(x int) {
	if x > -9999 && x < -18 {
		x = x * (-1)
	} else if x >= -18 && x <= -1 {
		x = x % 4
	} else if x >= 1 && x < 20 {
		x = x * x
	} else if x >= 20 && x < 9999 {
		x = -x
	}
	fmt.Println("x= ", x)
}

func ejercicio9(frase string) {
	fraseOriginal := frase
	fmt.Println("Frase original: ", fraseOriginal) // la imprimo coomo entra
	fraseMin := strings.ToLower(frase)             // transformo frase a minuscula y la guardo en fraseMin
	pos := strings.Index(fraseMin, "jueves")       // busco la ocurrencia de la palabra 'jueves' en la frase convertida en minuscula fraseMin

	println("primera letra de pos: ", string(fraseOriginal[pos]))

	posiciones := make([]int, 0, 0)
	posiciones = append(posiciones, pos)
	if pos != -1 {
		fraseMin = fraseMin[pos+6:]
		fmt.Println("Frase acortada: ", fraseMin)
		for range fraseMin {
			println("pos vale: ", pos)
			pos = strings.Index(fraseMin, "jueves")
			if pos != -1 {
				fraseMin = fraseMin[pos+6:]
				fmt.Println("Frase acortada: ", fraseMin)
				posiciones = append(posiciones, pos)
			}
		}
	}
	fmt.Println("Posiciones donde aparece jueves: ", posiciones)

	println("letras de pos: ", string(fraseOriginal[posiciones[0]]))
	println("letras de pos: ", string(fraseOriginal[posiciones[1]]))
	println("letras de pos: ", string(fraseOriginal[posiciones[2]]))

	// fmt.Println("Frase transformada en minuscula: ", fraseMin)                                                                                // Imprimo frase en minuscula
	// fmt.Println("Las posiciones de la frase completa en la que tengo que cambiar jueves por martes son: ", strings.Index(fraseMin, "jueves")) // posicion en la que tengo que insertar la palabra    // imprime la posicion de la primera vez que encuentra la palabra jueves, indice basado en las letras, por ejemplo en este caso imprime 3 porque encuentra donde empieza la palabra jueves(el espacio lo cuenta)

	// if pos == -1 {
	// 	fmt.Println("No se encontro la palabra 'jueves' contenida en la frase ingresada")
	// }

	// var palabraJueves = fraseOriginal[pos : pos+6] // palabraJueves = fraseOriginal desde posicion donde empieza jueves hasta la cantidad de letras
	// fmt.Println("Palabra encontrada: ", palabraJueves)

	// letras := []rune(palabraJueves)
	// for i, letra := range letras {
	// 	fmt.Printf("%c es mayuscula: %v\n", letra, unicode.IsUpper(letra))
	// 	if unicode.IsUpper(letra) {
	// 		fmt.Println("La letra mayuscula de la frase original esta en la posicion: ", i)
	// 	}
	// }

	// // construyo 'martes' respetando las mayusculas que traia juevEs

	// target := []rune("martes")                // Convierte el string "martes" en un slice de rune. => ['m', 'a', 'r', 't', 'e', 's']
	// palabraNueva := make([]rune, len(target)) // Crea un slice vacío de tipo rune, con el mismo tamaño que target (6 letras en "martes"). make([]rune, 6) = slice de 6 lugares listo para que pongamos letras.

	// for i, letra := range target { //Recorre target ("martes") letra por letra.
	// 	if unicode.IsUpper(rune(palabraJueves[i])) { // Pregunta: ¿Esta letra (posicion i) de "jueves" está en mayúscula?(palabraJueves[0] => false)
	// 		palabraNueva[i] = unicode.ToUpper(letra) // Si la letra en "jueves" era mayúscula, entonces convierto la letra de "martes" a mayúscula.
	// 	} else {
	// 		palabraNueva[i] = letra // de lo contrario la dejo igual que como estaba
	// 	}
	// }

	// fmt.Println("Palabra nueva: ", string(palabraNueva))
	// // remplazo primera aparicion
	// fraseResultante := fraseOriginal[:pos] + string(palabraNueva) + fraseOriginal[pos+6:]

	// fmt.Println("Frase final: ", fraseResultante)

}
