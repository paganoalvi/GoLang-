package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Funcion para determinar si un numero es primo (igual que en version secuencial)()
func esPrimoB(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

/*
encontrarPrimosEnRango ejecuta cada goroutine para buscar primos en un rango especifico
Parametros:
- inicio, fin: rango de numeros a comprobar
- primos: puntero al slice compartido de resultados
- wg: puntero al WaitGroup para sincronizacion
- mu: puntero al Mutex para acceso seguro al slice compartido
*/
func encontrarPrimosEnRango(inicio, fin int, primos *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // Notifica al WaitGroup cuando termine

	var primosLocales []int // Almacena primos encontrados en este rango (slice de primos locales)

	// Buscar primos en el rango asignado
	for i := inicio; i <= fin; i++ {
		if esPrimoB(i) {
			primosLocales = append(primosLocales, i) // agrego al slice de primos locales
		}
	}

	// Seccion critica: agregar resultados al slice compartido (ojo! tene cuidado con recurso compartido)
	mu.Lock()                                   // Bloqueamos el acceso concurrente
	*primos = append(*primos, primosLocales...) // agrego todos los primosLocales en la direccion de memoria de primos
	mu.Unlock()                                 // Liberamos el acceso
}

/*
	encontrarPrimosConcurrent coordina las goroutines

Parametros:
  - N: numero maximo hasta donde buscar primos
  - numWorkers: cantidad de goroutines a utilizar
*/
func encontrarPrimosConcurrent(N, numWorkers int) []int {
	var primos []int      // Slice compartido para resultados
	var wg sync.WaitGroup // Para esperar que terminen todas las goroutines
	var mu sync.Mutex     // Para proteger el acceso concurrente al slice

	// Dividir el trabajo en rangos aproximadamente iguales
	porcion := N / numWorkers      // Tamaño base de cada rango
	recordatorio := N % numWorkers // Números restantes por distribuir

	inicio := 2 // Empezamos desde el primer numero primo
	for i := 0; i < numWorkers; i++ {
		fin := inicio + porcion - 1 // Calcula fin del rango

		// Distribuye el recordatorio entre los primeros workers
		if i < recordatorio {
			fin++
		}

		// Asegura que no nos pasemos de N
		if fin > N {
			fin = N
		}

		wg.Add(1) // Incrementa contador del WaitGroup

		// Lanza goroutine para procesar el rango
		go encontrarPrimosEnRango(inicio, fin, &primos, &wg, &mu) // se lanza una gorutine que ejecuta la funcion encontrarPrimosEnRango

		inicio = fin + 1 // Prepara inicio para el siguiente rango
	}

	wg.Wait() // Espera a que todas las goroutines terminen
	return primos
}

func mainO1B() {
	// Validación de argumentos
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run obligatorio1(b).go <numero> <goroutines>")
		fmt.Println("Se esperan dos parametros:")
		fmt.Println("1.<numero> => Numero hasta donde buscar primos")
		fmt.Println("2. <goroutines> => Cantidad de goroutines a utilizar")
		return
	}

	// Conversión del primer argumento (N)
	N, err := strconv.Atoi(os.Args[1]) //(string => entero)
	if err != nil || N < 1 {
		fmt.Println("Error: El primer parametro debe ser un entero positivo")
		return
	}

	// Conversión del segundo argumento (numWorkers)
	numWorkers, err := strconv.Atoi(os.Args[2]) //(string => entero)
	if err != nil || numWorkers < 1 {
		fmt.Println("Error: La cantidad de gorutines a utilizar debe ser un entero positivo")
		return
	}

	// Ejecución y medición de tiempo
	start := time.Now()
	primes := encontrarPrimosConcurrent(N, numWorkers)
	elapsed := time.Since(start)

	// Mostrar resultados
	fmt.Printf("Números primos hasta %d (usando %d goroutines):\n", N, numWorkers)
	if len(primes) > 100 {
		fmt.Printf("Mostrando primeros y últimos 50 primos (total: %d)\n", len(primes))
		fmt.Println(primes[:50])
		fmt.Println("...")
		fmt.Println(primes[len(primes)-50:])
	} else {
		fmt.Println(primes)
	}
	fmt.Printf("\nTiempo de ejecución: %s\n", elapsed)
}
