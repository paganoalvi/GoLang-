package main

import (
	"fmt"
	"time"
)

func enviador(ch chan<- string, id int, intervalo time.Duration, duration time.Duration) {
	timeout := time.After(duration)
	ticker := time.NewTicker(intervalo)
	defer ticker.Stop()
	for i := 1; ; i++ { // no entiendo este for
		select {
		case <-timeout:
			close(ch)
			return
		case <-ticker.C:
			mensaje := fmt.Sprintf("Mensaje %d del canal %d", i, id)
			ch <- mensaje

		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Canal 1 envia cada 500ms (por 5 segundos)
	go enviador(ch1, 1, 500*time.Millisecond, 5*time.Second)
	// Canal 2 envia cada 1500ms por 10 segundos
	go enviador(ch2, 2, 1500*time.Millisecond, 10*time.Second)

	var cant1, cant2 int
	start := time.Now()
loop:
	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				ch1 = nil
				fmt.Println("Canal 1 cerrado (5s alcanzado)")
				if ch2 == nil {
					break loop
				}
				continue
			}
			fmt.Printf("[%.1fs] Recibido de Canal 1: %s\n", time.Since(start).Seconds(), msg)
			cant1++
		case msg, ok := <-ch2:
			if !ok {
				ch2 = nil
				fmt.Println("Canal 2 cerrado (10s alcanzado)")
				if ch1 == nil {
					break loop
				}
				continue
			}
			fmt.Printf("[%.1fs] Recibido de Canal 2: %s\n", time.Since(start).Seconds(), msg)
			cant2++
		}
	}
	fmt.Printf("\nResumen:\nCanal 1: %d mensajes\nCanal 2: %d mensajes\n", cant1, cant2)

}

/*Este programa demuestra cómo usar select con timeouts para recibir datos de dos canales durante períodos de tiempo
diferentes (5 y 10 segundos respectivamente).*/
