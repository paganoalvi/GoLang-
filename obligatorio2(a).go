package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mainO2a() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	const (
		numCajas    = 4
		numClientes = 20
		maxAtencion = 1.0 // segundos
	)
	var wg sync.WaitGroup
	colaGlobal := make(chan int, numClientes)
	cajas := make(chan bool, numCajas)

	// Inicializar cajas disponibles
	for i := 0; i < numCajas; i++ {
		cajas <- true
	}
	start := time.Now()

	// Generar clientes
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		colaGlobal <- i
	}

	// Atender clientes
	go func() {
		for cliente := range colaGlobal {
			<-cajas // Ocupar caja
			go func(c int) {
				defer wg.Done()
				defer func() { cajas <- true }() // Liberar caja

				tiempoAtencion := rand.Float64() * maxAtencion
				time.Sleep(time.Duration(tiempoAtencion * float64(time.Second)))
			}(cliente)
		}
	}()

	wg.Wait()
	close(colaGlobal)
	close(cajas)
	elapsed := time.Since(start)
	fmt.Printf("\nTiempo total (Cola Global): %.2f segundos\n", elapsed.Seconds())

}
