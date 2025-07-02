package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main2b() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Semilla aleatoria

	const (
		numCajas    = 4   // Numero de cajas disponibles
		numClientes = 20  // Total de clientes
		maxAtencion = 1.0 // Tiempo maximo de atencion por cliente (en segundos)
	)

	var wg sync.WaitGroup               // WaitGroup para esperar a que todos los clientes sean atendidos
	colas := make([]chan int, numCajas) // Slice de canales (simula una cola por caja)

	// Inicializar colas individuales para cada caja
	for i := range colas {
		colas[i] = make(chan int, numClientes) // Canal con buffer suficiente para no bloquear
	}

	inicio := time.Now() // Inicio de la medicion de tiempo

	// Asignar los clientes a las cajas (round-robin)
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)              // Por cada cliente, incrementamos el contador del WaitGroup
		colas[i%numCajas] <- i // Asignamos el cliente a una caja segun round-robin
	}

	// Atender clientes: una goroutine por caja que consume su propia cola
	for i, cola := range colas {
		go func(caja int, c <-chan int) {
			for cliente := range c {
				// Simulamos la atencion con un sleep aleatorio entre 0 y 1 segundo
				tiempoAtencion := rand.Float64() * maxAtencion
				time.Sleep(time.Duration(tiempoAtencion * float64(time.Second)))
				fmt.Printf("Cliente %d atendido en caja %d (%.2f segundos)\n", cliente, caja+1, tiempoAtencion)
				wg.Done() // Cliente atendido
			}
		}(i, cola) // Pasamos el indice de la caja y su canal
	}

	wg.Wait() // Esperamos que todos los clientes hayan sido atendidos

	// Cerramos las colas una vez terminado (importante para que las goroutines terminen correctamente)
	for _, c := range colas {
		close(c)
	}

	// Medimos el tiempo total
	lapso := time.Since(inicio)
	fmt.Printf("\nTiempo total (Round-Robin): %.2f segundos\n", lapso.Seconds())
}

/*
 Round-robin es una forma de repartir tareas de manera equitativa y cíclica entre un conjunto de recursos disponibles.
 Caja 1 → Caja 2 → Caja 3 → Caja 4 → Caja 1 → Caja 2 → ....
*/
