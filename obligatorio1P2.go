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

func Nueva() Lista {
	return Lista{}
}

func (l *Lista) AgregarAtras(unDato ingresante) {
	n := &nodo{dato: unDato}
	if l.finalLista != nil {
		l.finalLista.sig = n
	} else {
		l.incioLista = n
	}
	l.finalLista = n
	l.tamanoLista++
}

func (l *Lista) Iterador(f func(ingresante)) {
	for act := l.incioLista; act != nil; act = act.sig {
		f(act.dato)
	}
}

func (l *Lista) ElimnarSi(pred func(ingresante) bool) {
	var ant *nodo
	act := l.incioLista

	for act != nil {
		if pred(act.dato) {
			if ant == nil {
				l.incioLista = act.sig
			} else {
				ant.sig = act.sig
			}
			if act == l.finalLista {
				l.finalLista = ant
			}
			l.tamanoLista++
		} else {
			ant = act
		}
		act = act.sig
	}
}

// inciso a

func MostrarBarilochenses(l Lista) {
	l.Iterador(func(i ingresante) {
		if i.ciudadOrigen == "Bariloche" {
			fmt.Println(i.nombreCompleto.nombre, i.nombreCompleto.apellido)
		}
	})
}

// inciso b

func AnoConMasIngresantes(l Lista) int {
	count := make(map[int]int)
	l.Iterador(func(i ingresante) { // vector contador
		count[i.fechaNacimiento.ano]++ // map => [2003,2001] => [4,7]
	})
	maxAno := 0
	maxCount := 0

	for ano, c := range count {
		if c > maxCount {
			maxCount = c
			maxAno = ano
		}
	}
	return maxAno
}

// inciso c

func CarreraMasInscriptos(l Lista) string {
	count := make(map[string]int) //
	l.Iterador(func(i ingresante) {
		count[i.codigoCarrera]++
	})
	carreraMax := " "
	maxCount := 0

	for c, n := range count {
		if n > maxCount {
			maxCount = n
			carreraMax = c
		}
	}
	return carreraMax
}

// inciso d

func EliminarSinTitulo(l *Lista) {
	l.ElimnarSi(func(i ingresante) bool {
		return !i.presentoTituloSecundario
	})
}

// Programa Principal
func main() {
	l := Nueva()

	l.AgregarAtras(ingresante{
		nombreCompleto:           nombreCompleto{"Ana", "García"},
		ciudadOrigen:             "Bariloche",
		fechaNacimiento:          fecha{10, 5, 2004},
		presentoTituloSecundario: true,
		codigoCarrera:            "LI",
	})

	l.AgregarAtras(ingresante{
		nombreCompleto:           nombreCompleto{"Alvaro", "Pagano"},
		ciudadOrigen:             "Bariloche",
		fechaNacimiento:          fecha{01, 02, 1996},
		presentoTituloSecundario: false,
		codigoCarrera:            "APU",
	})

	l.AgregarAtras(ingresante{
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
	l.Iterador(func(i ingresante) {
		fmt.Println(i.nombreCompleto)
	})

}
