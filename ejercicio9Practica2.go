package main

import (
	"errors"
	"fmt"
	"strconv"
)

// inciso a
type node struct { // nodo
	dato int   // dato
	sig  *node // sig
}
type List struct {
	head *node // puntero inicial
	tail *node // puntero al ultimo
	size int   // tamano?
}

// OPERACIONES BASICAS

func NewLis() List {
	return List{}
}

func IsEmpty(l List) bool {
	return l.head == nil
}

func LenL(l List) int {
	return l.size
}

func FrontElement(l List) (int, error) {
	if l.head == nil { // si la lista esta vacia,manejo el error
		return 0, errors.New("La lista esta vacia") // devuelvo cero y el mensaje de error
	}
	return l.head.dato, nil // sino retorno el valor y nil para error
}

func Next(l List) List {
	if l.head == nil {
		return List{} // si esta vacia devuelvo lista vacia
	}

	return List{head: l.head.sig, tail: l.tail.sig, size: l.size - 1}
}

func ToString(l List) string {
	s := ""
	act := l.head
	for act != nil { // while (act != nil)
		s += strconv.Itoa(act.dato) + " => " // convierto a string el contenido de act y concateno
		act = act.sig                        // act = act^.sig
	}
	return s + "nil"
}

// FUNCIONES DE INSERCION Y ELIMINACION

func PushFront(l *List, elem int) { // recibe la lista ("x referencia") y un dato(int)
	n := &node{dato: elem, sig: l.head} // var local &nodo => new nodo(elem); nodo=L
	l.head = n                          // L = n
	if l.tail == nil {                  // si esl primer nodo tambien es el ultimo
		l.tail = n
	}
	l.size++
}

func PushBack(l *List, elem int) {
	n := &node{dato: elem}
	if l.tail != nil {
		l.tail.sig = n
		l.tail = n
	} else {
		l.head = n
		l.tail = n
	}
	l.size++
}

func Remove(l *List) (int, error) {
	if l.head == nil {
		return 0, errors.New("La lista esta vacia, no se puede borrrar")
	}
	val := l.head.dato
	l.head = l.head.sig
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return val, nil

}

// ITERADOR CON FUNCION ANONIMA
func Iterate(l List, f func(int) int) {
	act := l.head
	for act != nil { // while actual != nil
		act.dato = f(act.dato) // f retorna el dato contenido en act
		act = act.sig          // act := act^.sig
	}
}

// inciso B
func main9() {
	l := NewLis()
	PushBack(&l, 1)
	PushFront(&l, 2)
	PushBack(&l, 3)

	fmt.Println("Lista: ", ToString(l))
	fmt.Println("Tamano lista: ", LenL(l))

	dato, err := FrontElement(l)
	if err == nil {
		fmt.Println("Primer dato: ", dato)
	}

	Iterate(l, func(n int) int {
		return n * 2
	})
	fmt.Println("Despues de multiplicar x 2: ", ToString(l))

	dato, err = Remove(&l)
	if err == nil {
		fmt.Println("Elemento eliminado: ", dato)
	}
	fmt.Println("Lista final: ", ToString(l))

}

// CONTAINER/LIST
/*Tiene estructuras internas similares (list.List, list.Element).

Soporta operaciones como PushBack, PushFront, Remove, etc.

Es más general, pero no genérica (usa interface{} para elementos).

Ventaja: ya optimizada y probada.

Desventaja: menos tipo-segura que una implementación con genéricos (T any).*/

// inciso d
