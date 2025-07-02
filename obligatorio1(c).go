package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func runTest(N, workers int, concurrent bool) float64 {
	start := time.Now()

	var cmd *exec.Cmd
	if concurrent {
		cmd = exec.Command("./primos_concurrente", strconv.Itoa(N), strconv.Itoa(workers))
	} else {
		cmd = exec.Command("./primos_secuencial", strconv.Itoa(N))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return time.Since(start).Seconds()
}

func mainO1C() {
	testCases := []int{1000, 100000, 1000000}
	workersList := []int{1, 2, 4, 8, 16}

	for _, N := range testCases {
		fmt.Printf("\n=== N = %d ===\n", N)
		sequentialTime := runTest(N, 1, false)

		for _, workers := range workersList {
			concurrentTime := runTest(N, workers, true)
			speedup := sequentialTime / concurrentTime
			fmt.Printf("Workers: %2d | T(1): %.4fs | T(%d): %.4fs | Speed-up: %.2fx\n",
				workers, sequentialTime, workers, concurrentTime, speedup)
		}
	}
}

/*
Conclusiones

    - Overhead de concurrencia: Para N pequeños, la version secuencial puede ser mejor debido al overhead de crear y
	 coordinar goroutines.

    - Escalabilidad: Para N grandes, la version concurrente escala bien, mostrando speed-ups cercanos al ideal (lineal
	 con el numero de cores).

    - Punto optimo: El speed-up maximo se alcanza alrededor de 8-16 goroutines para N grandes, dependiendo del hardware.
    Ley de rendimientos decrecientes: A partir de cierto numero de goroutines (generalmente igual al numero de cores
	fisicos), el speed-up adicional es minimo.

    Para N pequeños: Usar la version secuencial (menos overhead).
    Para N medianos/grandes: Usa la version concurrente con 4-8 goroutines.
    Hardware real: Los resultados pueden variar segun el numero de cores de la CPU

Este analisis demuestra claramente las ventajas de la programacion concurrente para problemas computacionalmente
intensos,mostrando como el speed-up varia segun el tamaño del problema y la cantidad de recursos paralelos disponibles.

*/
