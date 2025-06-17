package main

import (
	"fmt"
	"sync"
	"time"
)

func main6() {
	// Creamos 3 canales
	chan1 := make(chan int) // Tres canales independientes que envían secuencias de números diferentes.
	chan2 := make(chan int)
	chan3 := make(chan int)

	// Variables para contar los valores recibidos de cada canal
	var cant1, cant2, cant3 int
	var wg sync.WaitGroup

	// Funcion para enviar datos a un canal
	enviarACanal := func(ch chan<- int, valores []int) {
		defer wg.Done()
		for _, v := range valores {
			time.Sleep(time.Duration(100) * time.Millisecond) // pequena pausa
			ch <- v                                           // ch recibe v
		}
		close(ch) // cierro canal una vez que termina for range
	}

	// Valores que se enviaran a cada canal
	valores1 := []int{1, 4, 7, 10}
	valores2 := []int{2, 5, 4, 3, 2}
	valores3 := []int{4, 5, 2, 4, 5, 3, 2, 4}

	// Iniciamos las gorutines para enviar datos

	wg.Add(3) // una por cada canal

	go enviarACanal(chan1, valores1)
	go enviarACanal(chan2, valores2)
	go enviarACanal(chan3, valores3)

	// Recibir valores usando select
	var totalRecibidos int
	canalesActivos := 3 // contador de canales activos

	for canalesActivos > 0 {
		select { // Uso de select para recibir valores del primer canal que tenga datos disponibles.
		case val, ok := <-chan1: //  no termino de entender de donde sale el valor de ok
			//fmt.Println(ok)
			//fmt.Println(val)
			if ok {
				fmt.Printf("Recibido del canal 1: %d\n", val)
				cant1++ // Conteo individual de valores recibidos de cada canal.
				totalRecibidos++
			} else {
				chan1 = nil // Manejo de cierre de canales - cuando un canal se cierra, se marca como nil para que select lo ignore
				canalesActivos--
			}
		case val, ok := <-chan2:
			if ok {
				fmt.Printf("Recibido del canal 2: %d\n", val)
				cant2++
				totalRecibidos++
			} else {
				chan2 = nil
				canalesActivos--
			}
		case val, ok := <-chan3:
			if ok {
				fmt.Printf("Recibido del canal 3: %d\n", val)
				cant3++
				totalRecibidos++
			} else {
				chan3 = nil
				canalesActivos--
			}
		}
	}
	wg.Wait() // esperamos que todas las gorutines terminen

	// Mostrar resumen
	fmt.Println("\nResumen:")
	fmt.Printf("Total recibido del Canal 1: %d valores\n", cant1)
	fmt.Printf("Total recibido del Canal 2: %d valores\n", cant2)
	fmt.Printf("Total recibido del Canal 3: %d valores\n", cant3)
	fmt.Printf("Total general de valores recibidos: %d\n", totalRecibidos)

}

/*Este programa utiliza la declaración select para recibir valores de tres canales diferentes de manera concurrente.
Cada canal envía una secuencia de números enteros, y el programa recibe y muestra los valores a medida que están
disponibles,manteniendo un conteo de los valores recibidos de cada canal.*/
