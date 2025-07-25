package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const workersNum = 4

// Task representa una tarea a procesar con su numero y prioridad
// Las prioridades van de 0 (mas alta) a 3 (mas baja)
type Task struct {
	Numero    int
	Prioridad int // Rangos validos: 0-3
}

// sumDigitos calcula la suma de los digitos de un numero
// Ejemplo: 1234 -> 1+2+3+4 = 10
func sumDigitos(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// numInvertido devuelve el numero invertido
// No maneja ceros a la izquierda en el resultado
func numInvertido(n int) int {
	invertido := 0
	for n > 0 {
		invertido = invertido*10 + n%10
		n = n / 10
	}
	return invertido
}

// worker procesa tareas desde el canal `tasks`
// Usamos un mutex global para prioridad 3 y escritura de archivos
func worker(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		switch task.Prioridad {
		case 0:
			resultado := sumDigitos(task.Numero)
			writeToFile("prioridad0.txt", fmt.Sprintf("(%d, %d)\n", task.Prioridad, resultado))
			fmt.Printf("Prioridad 0: %d = %d\n", task.Numero, resultado)
		case 1:
			resultado := numInvertido(task.Numero)
			// Si el número invertido tiene ceros iniciales, no se muestran :(
			// Ej: 280 -> 82 en lugar de 082
			writeToFile("prioridad1.txt", fmt.Sprintf("(%d, %d)\n", task.Prioridad, resultado))
			fmt.Printf("Prioridad 1: %d = %d\n", task.Numero, resultado)
		case 2:
			resultado := task.Numero * 10
			fmt.Printf("Prioridad 2: %d * 10 = %d\n", task.Numero, resultado)
		case 3:
			mu.Lock()
			acumulador += task.Numero
			fmt.Printf("Prioridad 3: Acumulado actual = %d (añadido %d)\n", acumulador, task.Numero)
			mu.Unlock() // Unlock manual (no defer) para evitar deadlocks
		}
	}
}

// writeToFile escribe en un archivo con locking global
// OBSERVACIÓN: Usa el mismo mutex que el acumulador, podría afectar rendimiento?
func writeToFile(nombreArchivo, contenido string) {
	mu.Lock()
	defer mu.Unlock() // defer en función corta

	arch, err := os.OpenFile(nombreArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error abriendo archivo %s: %v\n", nombreArchivo, err)
		return
	}
	defer arch.Close()

	writer := bufio.NewWriter(arch)
	_, err = writer.WriteString(contenido)
	if err != nil {
		fmt.Printf("Error escribiendo en archivo %s: %v\n", nombreArchivo, err)
	}
	writer.Flush() // Siempre hacer flush para no perder datos
}

// Variables globales compartidas
var (
	mu         sync.Mutex // Protege acumulador y operaciones de archivo
	acumulador int        // Acumulador para prioridad 3
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // inicializacion de rand (seed deprecated)

	// Limpieza inicial de archivos
	os.Remove("prioridad0.txt")
	os.Remove("prioridad1.txt")

	// Canales por prioridad
	priorityChannels := make([]chan Task, 4)
	for i := range priorityChannels {
		priorityChannels[i] = make(chan Task, 100) // Buffered channel(100)
	}

	// Canal principal de tareas
	taskQueue := make(chan Task, 100)

	// Inicio de workers
	var wg sync.WaitGroup
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(taskQueue, &wg) // Puntero a WaitGroup
	}

	// Generador de tareas (goroutine separada)
	go func() {
		for i := 0; i < 50; i++ {
			num := rand.Intn(10000)   // Numeros entre 0-9999
			prioridad := rand.Intn(4) // Prioridad entre 0-3
			tarea := Task{Numero: num, Prioridad: prioridad}

			// Enviamos a canal de prioridad especifica
			priorityChannels[prioridad] <- tarea

			// Retardo aleatorio entre 0-100ms (simulamos carga)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}

		// Cerramos canales de prioridad
		for i := range priorityChannels {
			close(priorityChannels[i])
		}
	}()

	// Scheduler: ordena y despacha tareas se supone que ya estan todas disponibles de 0 a 3
	go func() { // funcion anonima
		// Iteramos de prioridad 0 a 3 (de mayor a menor prioridad)
		for prioridad := 0; prioridad < 4; prioridad++ {
			// Mientras el canal de esta prioridad este abierto, procesamos sus tareas
			for tarea := range priorityChannels[prioridad] {
				fmt.Printf("Scheduler: procesando prioridad %d\n", prioridad)
				taskQueue <- tarea // Enviamos la tarea al canal que consumen los workers
			}
			// Solo despues de agotar todas las tareas de esta prioridad
			// se pasa a la siguiente prioridad
		}
		// Cuando terminan todas las prioridades, se cierra el canal global
		close(taskQueue)
	}()

	wg.Wait() // Esperamos a que todos los workers terminen
	fmt.Println("Procesamiento completado")
}

/* EXPLICACION DE SCHEDULER
 ¿Porque esto garantiza el orden?
    - Go bloquea automaticamente el for tarea := range ch hasta que el canal se cierra.
    - Como cada priorityChannels[i] se cierra en el mismo orden en que fue llenado (dentro del generador), el scheduler
	  solo avanza a la siguiente prioridad cuando se agotaron todas las tareas de la actual.
    - Asi, aunque haya tareas de prioridad baja disponibles(num mayor), no se ejecutaran hasta terminar todas las de
	 mayor prioridad(num menor).
*/
