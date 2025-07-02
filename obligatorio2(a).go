package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func mainO2a() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Inicializacion para generar numeros aleatorios

	const (
		numCajas    = 4   // Cantidad de cajas disponibles
		numClientes = 20  // Cantidad total de clientes a atender
		maxAtencion = 1.0 // Maximo tiempo de atencion por cliente (en segundos)
	)

	var wg sync.WaitGroup                     // WaitGroup para esperar a que todos los clientes sean atendidos
	colaGlobal := make(chan int, numClientes) // Cola global donde esperan los clientes
	cajas := make(chan bool, numCajas)        // Canal para gestionar cajas disponibles

	// Inicializarcion de todas las cajas como disponibles (true)
	for i := 0; i < numCajas; i++ {
		cajas <- true
	}

	inicio := time.Now() // Registrar el momento de inicio para medir el tiempo total

	// Generar clientes: simplemente agregamos los IDs de los clientes a la cola
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)       // Aumentamos el contador del WaitGroup por cada cliente
		colaGlobal <- i // Enviamos el cliente a la cola global
	}

	// Goroutine para consumir la cola y asignar clientes a cajas
	go func() {
		for cliente := range colaGlobal {
			<-cajas // Espera hasta que una caja este disponible (bloquea si no hay ninguna)
			go func(c int) {
				defer wg.Done()                  // Marca al cliente como atendido al final
				defer func() { cajas <- true }() // Libera la caja al finalizar

				// Simular atención con un tiempo aleatorio entre 0 y 1 segundo
				tiempoAtencion := rand.Float64() * maxAtencion
				fmt.Printf("Cliente %d está siendo atendido...\n", c)
				time.Sleep(time.Duration(tiempoAtencion * float64(time.Second)))
				fmt.Printf("Cliente %d atendido en %.2f segundos\n", c, tiempoAtencion)
			}(cliente)
		}
	}()

	wg.Wait()         // Esperamos que todos los clientes hayan sido atendidos
	close(colaGlobal) // Cerramos la cola global (buena practica, aunque no estrictamente necesario aca)
	close(cajas)      // Cerramos el canal de cajas tambien

	lapso := time.Since(inicio) // Medimos el tiempo total transcurrido
	fmt.Printf("\nTiempo total (Cola Global): %.2f segundos\n", lapso.Seconds())
}

/* BREVE RESUMEN DE SINTAXYS DE CANALES
- ENVIAR un valor al canal:  canal <- valor goroutine bloquea (espera) hasta que otra goroutine estE recibiendo del canal
 Ejemlo :
ch := make(chan int)
go func() {
    ch <- 5 // Enviar 5 al canal
}()
valor := <-ch // Recibir el 5

- RECIBIR un valor del canal: valor := <-canal La goroutine bloquea hasta que haya un valor disponible para recibir.

- Canal de solo LECTURA : func leer(c <-chan int)
- Canal de solo ESCRITURA: func escribir(c chan<- int)
 (Esto es util para limitar el uso del canal y evitar errores)
*/

/*	RESUMEN VISUAL
-----------------------------------------------------------------------
|Operación			 |	Sintaxis	|	Significado                   |
----------------------------------------------------------------------|
|Enviar al canal	 |	ch <- x		|	Manda x al canal ch       	  |
|Recibir del canal	 |	x := <- ch	|	Recibe valor desde ch         |
|Canal solo lectura	 |	<-chan T	|	Solo se puede leer de él      |
|Canal solo escritura|	chan<- T	|	Solo se puede escribir en él  |
-----------------------------------------------------------------------
*/
/*
Analogia simple
Un canal es una cinta transportadora:
    ch <- 10 pone el valor 10 en la cinta.
    <- ch toma el valor 10 de la cinta.
La cinta no avanza si no hay alguien del otro lado que
reciba o envie. Por eso se dice que los canales son bloqueantes por defecto.
*/
