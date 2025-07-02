package main

/*Un número primo es un número entero mayor que 1 que solo tiene dos divisores: 1 y el propio número.*/
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Función para determinar si un número es primo
// Saltamos comprobación de múltiplos de 2 y 3
func esPrimo(n int) bool {
	if n <= 1 {
		return false // Los números menores o iguales a 1 no son primos
	}
	if n <= 3 {
		return true // 2 y 3 son primos
	}
	if n%2 == 0 || n%3 == 0 {
		return false // Descartamos múltiplos de 2 y 3
	}
	// Comprobamos divisibilidad desde 5 hasta sqrt(n)(raiz de n)
	// Incrementamos de 6 en 6 (i y i+2)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Función secuencial para encontrar todos los primos hasta N
func encontrarPrimoSecuen(N int) []int {
	var primos []int // Slice para almacenar los números primos encontrados

	// Iteramos desde 2 hasta N (incluido)
	for i := 2; i <= N; i++ {
		if esPrimo(i) {
			primos = append(primos, i) // Agregamos a la lista si es primo
		}
	}
	return primos
}

func maino1() {
	// Validación de argumentos de línea de comandos
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run obligatorio1(b).go <número>")
		fmt.Println("Se espera recibir un número entero como parámetro")
		return
	}

	// Conversión del argumento a entero
	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N < 1 {
		fmt.Println("Por favor ingrese un número entero positivo válido")
		return
	}

	// Medición del tiempo de ejecución
	inicio := time.Now()
	primos := encontrarPrimoSecuen(N)
	lapso := time.Since(inicio)

	// Mostrar resultados
	fmt.Printf("Números primos hasta %d:\n", N)
	if len(primos) > 100 {
		// Para listas largas, mostramos solo los primeros y últimos 50
		fmt.Printf("Mostrando primeros y últimos 50 primos (total: %d)\n", len(primos))
		fmt.Println(primos[:50])
		fmt.Println("...")
		fmt.Println(primos[len(primos)-50:])
	} else {
		// Para listas cortas, mostramos todos
		fmt.Println(primos)
	}
	fmt.Printf("\nTiempo de ejecución: %s\n", lapso)
}

/* package strconv
Atoi(s string) (int, error): Convierte la cadena s a un entero. Si la conversión falla, retorna un error
*/
