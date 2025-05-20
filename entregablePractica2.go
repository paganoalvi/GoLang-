/*
	Objetivo => Diseñar una estructura para representar slices de enteros con muchas

repeticiones consecutivas, de forma óptima en memoria, y manipularla a través
de funciones específicas.

	Run-Length Encoding (RLE) datos con muchas rachas (runs) de elementos repetidos

Ej:
[]int{1,1,1,2,2,3,3,3,3} => [

	{valor: 1, repeticiones: 3},
	{valor: 2, repeticiones: 2},
	{valor: 3, repeticiones: 4},

]

append(slice,valorNuevo) =>
funcion integrada que sirve para agregar un "valorNuevo" o
"valoresNuevos" al final de un "slice"
*/
package main

import "fmt"

type secuencia struct {
	valor        int
	repeticiones int
}

type OptimumSlice struct {
	secuenciasNumeros []secuencia
}

// Recibimos un slice y lo compactamos en un OptimumSlice
func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return OptimumSlice{secuenciasNumeros: []secuencia{}} // retorno slice vacio
	}
	secuenciaNumeros := []secuencia{}
	numAct := s[0]
	repe := 1

	for i := 1; i < len(s); i++ {
		if s[i] == numAct {
			repe++
		} else {
			secuenciaNumeros = append(secuenciaNumeros, secuencia{valor: numAct, repeticiones: repe})
			numAct = s[i]
			repe = 1
		}
	}
	secuenciaNumeros = append(secuenciaNumeros, secuencia{valor: numAct, repeticiones: repe})
	return OptimumSlice{secuenciasNumeros: secuenciaNumeros}
}

// Convertimos de OptimumSlice a slice comun
func SliceArray(os OptimumSlice) []int { // recibimos un os y retornamos un slice
	var resultado []int
	for _, r := range os.secuenciasNumeros { //para todos los valores en os _ indice r valor
		for i := 0; i < r.repeticiones; i++ {
			resultado = append(resultado, r.valor)
		}
	}
	return resultado
}

// len devuelve la dimension logica de lo que seria el slice sin comprimir
func Len(os OptimumSlice) int {
	long := 0
	for _, v := range os.secuenciasNumeros {
		long += v.repeticiones
	}
	return long
}

// isEmpty devuelve true si el os se encuentra vacio
func isEmpty(os OptimumSlice) bool {
	return (len(os.secuenciasNumeros) == 0)
}

func main() {
	sl := []int{}
	os := New(sl)
	fmt.Println(isEmpty(os))

	s := []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 5, 5, 5, 5, 5}
	o := New(s) // primera prueba de new
	fmt.Println(o)
	sliceAgain := SliceArray(o)
	fmt.Println(sliceAgain) // segunda prueba

	fmt.Println(isEmpty(o))

}
