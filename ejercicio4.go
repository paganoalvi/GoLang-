package main

import (
	"fmt"
	"sync"
)

func ping(pingC chan<- string, pongC <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	pingC <- "PING"  // Envía "PING" (se almacena en el buffer)
	msg := <-pongC   // Espera recibir "PONG"
	fmt.Println(msg) // Imprime "PONG" (para verificar)
}

func pong(pongC chan<- string, pingC <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-pingC   // Espera recibir "PING"
	fmt.Println(msg) // Imprime "PING" (para verificar)
	pongC <- "PONG"  // Envía "PONG" (se almacena en el buffer)
}

func main4() {
	const cant = 5
	pingChan := make(chan string, cant) // creo canal y agrego Buffer para "PING"
	pongChan := make(chan string, cant) // creo canal y agrego Buffer para "PONG"
	var wg sync.WaitGroup               // creo variable waitGroup

	wg.Add(2 * cant) // 2 goroutines por iteración (ping + pong)

	// Lanzamos las goroutines
	for i := 0; i < cant; i++ {
		go ping(pingChan, pongChan, &wg)
		go pong(pongChan, pingChan, &wg)
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen
	close(pingChan)
	close(pongChan)
}
