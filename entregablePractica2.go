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

	for i := 1; i < len(s); i++ { // for tradicional de 1 hasta len(s)
		if s[i] == numAct { //while
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

// Devolvemos el primer "valor" del Os
func FrontElemen(os OptimumSlice) int {
	if isEmpty(os) {
		panic("OptimumSlice vacio")
	}
	return os.secuenciasNumeros[0].valor
}

// Devolvemos el ultimo elemento del Os
func LastElement(os OptimumSlice) int {
	if isEmpty(os) {
		panic("Optimum Slice vacio")
	}
	return os.secuenciasNumeros[len(os.secuenciasNumeros)-1].valor
	//len(os.secuenciasNumeros) = 3 | -1 porque los indices comienzan en 0
}

// Implementamos Insert => Recibe un ptro a OS,un elemento(int),y una posicion(int) retorna un int
func Insert(os *OptimumSlice, elem int, pos int) int {
	if pos < 0 || pos > len(os.secuenciasNumeros) {
		panic("Posicion invalida para insertar")
	}
	if isEmpty(*os) { //si esta vacio, inserto al principio
		os.secuenciasNumeros = append(os.secuenciasNumeros, secuencia{elem, 1})
	}
	posLog := 0
	for i, r := range os.secuenciasNumeros {
		if posLog+r.repeticiones > pos {
			// insertar en medio de secuenciaNumeros
			offset := pos - posLog
			if r.valor == elem {
				// caso 1: mismo valor => aumentar repeticiones
				os.secuenciasNumeros[i].repeticiones++
				return i
			}
			// caso 2: valor distinto => partir el run en dos y meter el nuevo valor
			antes := secuencia{r.valor, offset}
			nuevaSecuen := secuencia{elem, 1}
			despues := secuencia{r.valor, r.repeticiones - offset}
			// remplazamos secuencia original con antes + nuevaSecuen + despues
			nuevaSecuencias := []secuencia{}
			nuevaSecuencias = append(nuevaSecuencias, os.secuenciasNumeros[:i]...)
			if antes.repeticiones > 0 {
				nuevaSecuencias = append(nuevaSecuencias, antes)
			}
			nuevaSecuencias = append(nuevaSecuencias, nuevaSecuen)
			if despues.repeticiones > 0 {
				nuevaSecuencias = append(nuevaSecuencias, despues)
			}
			nuevaSecuencias = append(nuevaSecuencias, os.secuenciasNumeros[i+1:]...)
			os.secuenciasNumeros = nuevaSecuencias
			return i
		}
		posLog += r.repeticiones
	}
	// ultimo caso: insertar al final
	ult := os.secuenciasNumeros[len(os.secuenciasNumeros)-1]
	if ult.valor == elem {
		os.secuenciasNumeros[len(os.secuenciasNumeros)-1].repeticiones++
	} else {
		os.secuenciasNumeros = append(os.secuenciasNumeros, secuencia{elem, 1})
	}
	return len(os.secuenciasNumeros) - 1
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
	Insert(&o, 4, 2)
	fmt.Println(o)
	s = SliceArray(o)
	fmt.Println(s)

	//fmt.Println(FrontElemen(o))
	//fmt.Println(LastElement(os))

}
