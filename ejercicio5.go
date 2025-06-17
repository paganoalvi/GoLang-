package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProd   = 2
	numCons   = 2
	numsXProd = 3
)

func productor(id int, nums chan<- int, wg *sync.WaitGroup, rng *rand.Rand) {
	defer wg.Done()
	for i := 0; i < numsXProd; i++ {
		// espera aleatoria entre 0 y 1 segundo
		delay := time.Duration(rand.Float64() * float64(time.Second))
		time.Sleep(delay)
	}
	// Generar un numero aleatorio entre 0 y 100
	num := rng.Intn(101)
	fmt.Printf("Productor %d produjo: %d\n", id, num)
	nums <- num // envia num
}

func consumidor(id int, nums <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numsXProd; i++ {
		num := <-nums // se queda esperando a recibir num
		fmt.Printf("Consumidor %d consumio: %d\n", id, num)
		time.Sleep(10 * time.Millisecond) // Pequeña pausa para dar chance al otro consumidor
	}
}

func main5() {
	rng1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng2 := rand.New(rand.NewSource(time.Now().UnixNano() + 1)) // Semilla diferente

	var wg sync.WaitGroup
	nums := make(chan int, numProd*numsXProd) // buffer para evitar bloqueos

	// Lanzar porductores
	wg.Add(numProd)
	go productor(1, nums, &wg, rng1) // lanzo productor 1, con numero aleatorio rng1
	go productor(2, nums, &wg, rng2)

	// Lanzar consumidores
	go consumidor(1, nums, &wg)
	go consumidor(2, nums, &wg)

	wg.Wait() // esperamos que todas las go rutines terminen
	close(nums)
}

/*Podria hacer un canal para cada consumidor para asegurarme de que cada uno reciba
problema de scheduling de gorutines, el balanceo no si es correcto(consultar)*/
