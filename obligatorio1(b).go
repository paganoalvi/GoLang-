package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func isPrime(n int) bool {
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

func findPrimesInRange(start, end int, primes *[]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	var localPrimes []int

	for i := start; i <= end; i++ {
		if isPrime(i) {
			localPrimes = append(localPrimes, i)
		}
	}
	mu.Lock()
	*primes = append(*primes, localPrimes...)
	mu.Unlock()
}

func findPrimesConcurrent(N, numWorkers int) []int {
	var primes []int
	var wg sync.WaitGroup
	var mu sync.Mutex

	chunkSize := N / numWorkers
	remainder := N % numWorkers

	start := 2
	for i := 0; i < numWorkers; i++ {
		end := start + chunkSize + -1
		if i < remainder {
			end++
		}
		if end > N {
			end = N
		}
		wg.Add(1)
		go findPrimesInRange(start, end, &primes, &wg, &mu)
		start = end + 1
	}
	wg.Wait()
	return primes
}

func mainO2() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run primos_concurrente.go <N> <num_goroutines>")
		return
	}
	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N < 1 {
		fmt.Println("Por favor ingrese un número entero positivo válido para N")
		return
	}
	numWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil || numWorkers < 1 {
		fmt.Println("Por favor ingrese un número entero positivo válido para el número de goroutines")
		return
	}
	start := time.Now()
	primes := findPrimesConcurrent(N, numWorkers)
	elapsed := time.Since(start)

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
