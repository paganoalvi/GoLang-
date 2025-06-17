package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func isPrime1(n int) bool {
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

func findPrimeSequential(N int) []int {
	var primes []int
	for i := 2; i <= N; i++ {
		if isPrime1(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func mainO1() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run primos_secuencial.go <N>")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil || N < 1 {
		fmt.Println("Por favor ingrese un numero entero positivo valido")
		return
	}

	start := time.Now()
	primes := findPrimeSequential(N)
	elapsed := time.Since(start)
	fmt.Printf("Números primos hasta %d:\n", N)
	if len(primes) > 100 {
		fmt.Printf("Mostrando primeros y últimos 50 primos (total: %d)\n", len(primes))
		fmt.Println(primes[:50])
		println("...")
		fmt.Println(primes[len(primes)-50:])
	} else {
		fmt.Println(primes)
	}
	fmt.Printf("\nTiempo de ejecución: %s\n", elapsed)

}
