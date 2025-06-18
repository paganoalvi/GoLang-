// ANALIZAR Y REANALIZAR Y RE CONTRA RE ANALIZAR EL PROGRAMA

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	numWorkers = 4
)

type Task struct {
	Number   int
	Priority int
}

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func worker(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		switch task.Priority {
		case 0:
			result := sumDigits(task.Number)
			writeToFile("prioridad0.txt", fmt.Sprintf("(%d, %d)\n", task.Priority, result))
			fmt.Printf("Prioridad 0: %d = %d\n", task.Number, result)
		case 1:
			result := reverseNumber(task.Number)
			writeToFile("prioridad1.txt", fmt.Sprintf("(%d, %d)\n", task.Priority, result))
			fmt.Printf("Prioridad 1: %d = %d\n", task.Number, result)

		case 2:
			result := task.Number * 10
			fmt.Printf("Prioridad 2: %d * 10 = %d\n", task.Number, result)
		case 3:
			// Acumulador global para prioridad 3
			mu.Lock()
			accumulator += task.Number
			fmt.Printf("Prioridad 3: Acumulado actual = %d (añadido %d)\n", accumulator, task.Number)
			mu.Unlock()
		}
	}
}

func writeToFile(filename, content string) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error abriendo archivo %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Printf("Error escribiendo en archivo %s: %v\n", filename, err)
		return
	}
	writer.Flush()
}

var (
	mu          sync.Mutex
	accumulator int
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Limpiar archivos de salida
	os.Remove("prioridad0.txt")
	os.Remove("prioridad1.txt")

	// Crear canales para cada prioridad
	priorityChannels := make([]chan Task, 4)
	for i := range priorityChannels {
		priorityChannels[i] = make(chan Task, 100)
	}

	// Canal para tareas a procesar
	taskQueue := make(chan Task, 100)

	// Iniciar workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(taskQueue, &wg)
	}

	// Generar 50 tareas aleatorias
	go func() {
		for i := 0; i < 50; i++ {
			num := rand.Intn(10000)
			priority := rand.Intn(4) // random entre 0 y 3
			task := Task{Number: num, Priority: priority}
			priorityChannels[priority] <- task
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
		// Cerrar canales de prioridad cuando no haya más tareas
		for i := range priorityChannels {
			close(priorityChannels[i])
		}
	}()

	// Scheduler
	go func() {
		// Procesar en orden de prioridad
		for priority := 0; priority < 4; priority++ {
			for task := range priorityChannels[priority] {
				taskQueue <- task
			}
		}
		close(taskQueue)
	}()

	wg.Wait()
	fmt.Println("Procesamiento completado")
}
