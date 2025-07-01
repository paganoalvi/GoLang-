package main

import "fmt"

type fecha struct {
	dia int
	mes int
	ano int
}
type nombreCompleto struct {
	nombre   string
	apellido string
}
type ingresante struct {
	nombreCompleto           nombreCompleto
	ciudadOrigen             string
	fechaNacimiento          fecha
	presentoTituloSecundario bool
	codigoCarrera            string // puede ser ("APU","LI" o LS)
}

type nodo struct {
	dato ingresante
	sig  *nodo
}

type Lista struct {
	incioLista  *nodo
	finalLista  *nodo
	tamanoLista int
}

func new() Lista {
	return Lista{} // retorna lista incializada vacia
}

func (l *Lista) agregarAtras(unDato ingresante) {
	n := &nodo{dato: unDato}
	if l.finalLista != nil {
		l.finalLista.sig = n
	} else {
		l.incioLista = n
	}
	l.finalLista = n
	l.tamanoLista++
}

func (l *Lista) iterador(f func(ingresante)) { // Receiver => (l *Lista) , Input Parameter: f of type func(ingresante)
	for act := l.incioLista; act != nil; act = act.sig {
		f(act.dato)
	}
}

func (l *Lista) ElimnarSi(pred func(ingresante) bool) { // metodo sobre l *lista
	var ant *nodo
	act := l.incioLista

	for act != nil {
		if pred(act.dato) { // funcion logica que se utiliza para determinar un criterio de eliminacion
			if ant == nil {
				l.incioLista = act.sig // nodo a eliminar es el primero de la lista ("desconecto" nodo 1 de la lista)
			} else {
				ant.sig = act.sig // nodo no es el primero, saltamos  actual
			}
			if act == l.finalLista { // nodo a eliminar es el ultimo, actualizo finlista
				l.finalLista = ant
			}
			l.tamanoLista-- // decremento tamano lista
		} else { // si no hay que eliminar el nodo actualizo anterior, antes de avanzar
			ant = act
		}
		act = act.sig // avanzo en la lista
	}
	/* no se hace explicitamente dispose(act) ya que en Go el garbage colector se encarga
	de eliminar los nodos "desconectados"
	[nodo1] -> [nodo2] -> [nodo3]
	nodo1.sig = nodo2.sig
	[nodo1] ------> [nodo3]
	[nodo2] ya no esta referenciado por nadie (dispose nodo2)
	*/
}

// inciso a

func MostrarBarilochenses(l Lista) { // recibo lista por valor
	l.iterador(func(i ingresante) { // funcion anonima para recorrer la lista utilizando iterator
		if i.ciudadOrigen == "Bariloche" { // si es de Bariloche, imprimo sus nombre completo
			fmt.Println(i.nombreCompleto.nombre, i.nombreCompleto.apellido)
		}
	})
}

// inciso b

func AnoConMasIngresantes(l Lista) int { // recibo lista por valor, devuelvo un int
	count := make(map[int]int)      // count = map de  clave(entero)|valor (entero)
	l.iterador(func(i ingresante) { // funcion anonima para iterar contando ingresantes por ano
		count[i.fechaNacimiento.ano]++ // map => [2003,2001] => [4,7] (map tipo vector contador(ano = indice))
	})
	maxAno := 0   // indice max
	maxCount := 0 // max

	for ano, c := range count { // recorro map "contador de ingresantes por ano" (ano = key | c = value)
		if c > maxCount {
			maxCount = c // actualizo el max
			maxAno = ano // actualizo el indice donde se encuentra el maximo
		}
	}
	return maxAno // devuelvo el ano (indice del map) con mas ingresantes
}

// inciso c

func CarreraMasInscriptos(l Lista) string { // recibo lista por valor, retorno string
	count := make(map[string]int)   // map => key(string) | value(int)
	l.iterador(func(i ingresante) { // recorro toda la lista, cargando map "contador" por codigo de carrera
		count[i.codigoCarrera]++
	})
	carreraMax := " " // incializo carrear con mas inscriptos con string vacio(indice del map)
	maxCount := 0     // max

	for c, n := range count { // para cada key|value en count
		if n > maxCount {
			maxCount = n
			carreraMax = c
		}
	}
	return carreraMax
}

// inciso d

func EliminarSinTitulo(l *Lista) { // paso lista por referencia
	l.ElimnarSi(func(i ingresante) bool { // defino el criterio de funcion anonima(retorna true si nodo actual de la lista tiene tituloSecundario = false)
		return !i.presentoTituloSecundario
	})
}

// Programa Principal
func mainO1() {
	l := new() // creo lista

	l.agregarAtras(ingresante{ // agrego ingresante al final
		nombreCompleto:           nombreCompleto{"Ana", "García"},
		ciudadOrigen:             "Bariloche",
		fechaNacimiento:          fecha{10, 5, 2004},
		presentoTituloSecundario: true,
		codigoCarrera:            "LI",
	})

	l.agregarAtras(ingresante{
		nombreCompleto:           nombreCompleto{"Alvaro", "Pagano"},
		ciudadOrigen:             "Bariloche",
		fechaNacimiento:          fecha{01, 02, 1996},
		presentoTituloSecundario: false,
		codigoCarrera:            "APU",
	})

	l.agregarAtras(ingresante{
		nombreCompleto:           nombreCompleto{"Agustin", "Servin"},
		ciudadOrigen:             "Bariloche",
		fechaNacimiento:          fecha{22, 9, 2002},
		presentoTituloSecundario: true,
		codigoCarrera:            "LS",
	})

	fmt.Println("Ingresantes de bariloche: ")
	MostrarBarilochenses(l)

	EliminarSinTitulo(&l)

	fmt.Println("Lista luego de eliminar ingresantes sin titulos secundarios: ")
	l.iterador(func(i ingresante) {
		fmt.Println(i.nombreCompleto)
	})

}

/* ESTRUCTURAS DE DATOS
struct
composite data type that allows for the grouping of related values of different types into a single unit. It
servesas a blueprint for creating custom data structures,similar to classes in object-oriented programming languages, but
without supporting inheritance in the same way (COMO UN REGISTRO)(o class)

- Methods:
 Methods can be associated with a struct type, allowing for the definition of behavior specific to instances of that
 struct.These methods are called using a receiver, which is an instance of the struct.
- Visibility:
 Field names (and the struct name itself) starting with an uppercase letter are public (exported) and accessible outside
 the package, while those starting with a lowercase letter are private (unexported) and only accessible within the same
 package.
- Initialization:
 Structs can be initialized using struct literals, either by listing all field values in order or by explicitly naming
 specific fields. Omitted fields will be assigned their zero value

*nodo
Declaring a Pointer to nodo: * signifies that the variable is a pointer to a
specific type (struct in this case). It indicates that the variable will hold the memory address of a value of that type,
rather than the value itself
* is fundamental to working with pointers in Go, allowing you to declare variables that store memory addresses and then
access the underlying values at those addresses

map
data structure that represents an unordered collection of key(unique)-value pairs.
- Unordered: Maps do not guarantee any specific order for their elements.
- Dynamic Size: Maps can grow or shrink dynamically as elements are added or removed.
- Comparable Keys: Keys must be of a comparable type (e.g., int, string, float64,
 struct where all fields are comparable). Slices and non-comparable arrays/structs cannot be used as keys
- Any Value Type: Values can be of any type, including other maps, slices, or structs.
- Zero Value: The zero value for a map is nil. A nil map cannot be used to store key-value pairs; it must be
 initialized first
*/

/* ESTRUCTURAS DE CONTROL O PARTICULARIDADES DE GO
function type parameter
func (l *Lista) iterador(f func(ingresante)) {
- (l *Lista) iterador esto indica que iterador es un metodo perteneciente al struct Lista
 (iterador is a method with receiver *Lista)

- f func(ingresante)
 Function parameter — a callback function that takes an ingresante

 f(act.dato): “Call the function f and pass it the current node’s data (act.dato) as the argument.”
*/

/*	FUNCION ANONIMA
La siguiente funcion por ejemplo:
l.ElimnarSi(func(i ingresante) bool {
	return !i.presentoTituloSecundario
})
Dado un ingresante, devolveme true si no presento el titulo secundario.

*/
