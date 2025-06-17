package main

import (
	"fmt"
	//"time"
	//"sync"
)

func main3() {
	//var wg sync.WaitGroup
	done := make(chan bool)
	fmt.Println("Inicia Goroutine del main")
	//wg.Add(1) // agrego una gorutine al waitGroup
	go func() {
		//defer wg.Done() // marco gorutine completada al finalizar
		hello()      // se ejecuta gorutine
		done <- true // se marca que gorutine termino
	}()
	//time.Sleep(1 * time.Second) // le doy tiempo a go rutine a terminar(solucion poco robusta)
	//wg.Wait() // espera a que terminen todas las gorutines del waitGroup

	<-done // Espera a recibir senal de termino del canal
	fmt.Println("Termina Goroutine del main")
}

func hello() {
	fmt.Println("Inicia Goroutine de hello")
	for i := 0; i < 3; i++ {
		fmt.Println(i, " Hello world")
	}
	fmt.Println("Termina Goroutine de hello")
}

/*

a)​ ¿Cuántas veces se imprime Hello world? => 0?
    Go no espera automáticamente a que las goroutines terminen(termina el main antes que hello)
    El scheduler de Go no garantiza el orden de ejecución entre goroutines
    La goroutine principal (main) tiene privilegio y su terminación mata el proceso completo

b)​ ¿Cuántas Goroutines tiene el programa? => 2?

c)​ ¿Cómo cambiaría el programa (con la misma cantidad de
Goroutines) para que imprima 3 veces Hello world? => agregando o time.sleep o waitgroup

i)​ Hágalo usando time.Sleep
ii)​ Hágalo usando Channel Synchronization



Trabajar con goroutines:

    Sincronización: Siempre uso mecanismos de sincronización (WaitGroup, canales) para coordinar goroutines

    Planificación: No asumo un orden de ejecución específico

    Comunicación: Uso canales para comunicar entre goroutines en lugar de variables compartidas

    Terminación: Me aaseguro de que el programa principal espere a las goroutines cuando sea necesario
*/
