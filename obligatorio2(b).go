package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main2b() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	const (
		numCajas    = 4
		numClientes = 20
		maxAtencion = 1.0 // segundos
	)

	var wg sync.WaitGroup
	colas := make([]chan int, numCajas)
	for i := range colas {
		colas[i] = make(chan int, numClientes)
	}
	start := time.Now()

	// Asignacion round-robin
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		colas[i%numCajas] <- i
	}

	//Atender clientes
	for i, cola := range colas {
		go func(caja int, c <-chan int) {
			for cliente := range c {
				tiempoAtencion := rand.Float64() * maxAtencion
				time.Sleep(time.Duration(tiempoAtencion * float64(time.Second)))
				fmt.Printf("Cliente %d atendido en caja %d (%.2f segundos)\n", cliente, caja+1, tiempoAtencion)
				wg.Done()
			}

		}(i, cola)
	}
	wg.Wait()
	for _, c := range colas {
		close(c)
	}

	elapsed := time.Since(start)

	fmt.Printf("\nTiempo total (Round-Robin): %.2f segundos\n", elapsed.Seconds())

}
