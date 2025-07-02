package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mai2c() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // incializo random

	const (
		numCajas    = 4   // Numero de cajas
		numClientes = 20  // Total de clientes
		maxAtencion = 1.0 // Máximo tiempo de atencion en segundos
	)

	var wg sync.WaitGroup
	colas := make([]chan int, numCajas) // slice de canales, con buffer = numCajas

	// Crear una cola (canal) para cada caja
	for i := range colas {
		colas[i] = make(chan int, numClientes)
	}

	incio := time.Now() // Marcar inicio del procesamiento

	// Asignación de clientes a la cola más corta
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		// Buscar la cola con menos elementos
		masCorta := 0
		for j := 1; j < numCajas; j++ {
			if len(colas[j]) < len(colas[masCorta]) {
				masCorta = j
			}
		}
		// Asignar el cliente a esa cola
		colas[masCorta] <- i
	}

	// Procesar la atencion de clientes en cada caja
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

	wg.Wait() // Esperar a que todos los clientes sean atendidos

	// Cerrar las colas (buena practica)
	for _, c := range colas {
		close(c)
	}

	lapso := time.Since(incio)
	fmt.Printf("\nTiempo total (Cola Más Corta): %.2f segundos\n", lapso.Seconds())
}
