package main

import (
	"errors"
	"fmt"
)

type secuencia struct {
	valor       int
	ocurrencias int
}

type OptimumSlice struct {
	secuenciasNumeros []secuencia
}

// Recibimos un slice y lo compactamos en un OptimumSlice
func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return OptimumSlice{secuenciasNumeros: []secuencia{}} // retorna slice vacio
	}
	secuenciaNumeros := []secuencia{}
	numAct := s[0]
	repe := 1

	for i := 1; i < len(s); i++ {
		if s[i] == numAct {
			repe++
		} else {
			secuenciaNumeros = append(secuenciaNumeros, secuencia{valor: numAct, ocurrencias: repe})
			numAct = s[i]
			repe = 1
		}
	}
	secuenciaNumeros = append(secuenciaNumeros, secuencia{valor: numAct, ocurrencias: repe})
	return OptimumSlice{secuenciasNumeros: secuenciaNumeros}
}

// Convertimos de OptimumSlice a slice comun
func SliceArray(os OptimumSlice) []int { // recibimos un os y retornamos un slice
	var resultado []int
	for _, r := range os.secuenciasNumeros {
		for i := 0; i < r.ocurrencias; i++ {
			resultado = append(resultado, r.valor)
		}
	}
	return resultado
}

func Len(os OptimumSlice) int {
	long := 0
	for _, v := range os.secuenciasNumeros {
		long += v.ocurrencias
	}
	return long
}

func isEmpty(os OptimumSlice) bool {
	return (len(os.secuenciasNumeros) == 0)
}

func FrontElement(os OptimumSlice) (int, error) {
	if isEmpty(os) {
		return -1, errors.New("No hay frontElement,optimumSlice vacio")
	}
	return os.secuenciasNumeros[0].valor, nil
}

func LastElement(os OptimumSlice) (int, error) {
	if isEmpty(os) {
		return -1, errors.New("No hay LastElement,optimumSlice vacio")
	}
	return os.secuenciasNumeros[len(os.secuenciasNumeros)-1].valor, nil
}

func Insert(os *OptimumSlice, elem int, pos int) (int, error) {
	pos = pos - 1
	if pos < 0 || pos > Len(*os) { // chequeo pos
		return -1, errors.New("Posicion no valida para insertar")
	}
	if isEmpty(*os) { //si esta vacio, inserto al principio
		os.secuenciasNumeros = append(os.secuenciasNumeros, secuencia{elem, 1})
		return 0, nil
	}
	posLog := 0
	for i, r := range os.secuenciasNumeros {
		if posLog+r.ocurrencias > pos {
			// offset define en que parte del bloque insertar
			offset := pos - posLog
			if r.valor == elem { //caso 1: mismo valor => aumentar repeticiones
				os.secuenciasNumeros[i].ocurrencias++
				return pos, nil
			}
			// caso 2: valor distinto => partir el run en dos y meter el nuevo valor
			antes := secuencia{r.valor, offset}
			nuevaSecuen := secuencia{elem, 1}
			despues := secuencia{r.valor, r.ocurrencias - offset}
			// remplazamos secuencia original con antes + nuevaSecuen + despues
			nuevaSecuencias := []secuencia{} //OSlice vacio
			nuevaSecuencias = append(nuevaSecuencias, os.secuenciasNumeros[:i]...)
			if antes.ocurrencias > 0 {
				nuevaSecuencias = append(nuevaSecuencias, antes)
			}
			nuevaSecuencias = append(nuevaSecuencias, nuevaSecuen)
			if despues.ocurrencias > 0 {
				nuevaSecuencias = append(nuevaSecuencias, despues)
			}
			nuevaSecuencias = append(nuevaSecuencias, os.secuenciasNumeros[i+1:]...)
			os.secuenciasNumeros = nuevaSecuencias
			return pos, nil
		}
		posLog += r.ocurrencias
	}
	// ultimo caso: insertar al final
	ult := os.secuenciasNumeros[len(os.secuenciasNumeros)-1]
	if ult.valor == elem {
		os.secuenciasNumeros[len(os.secuenciasNumeros)-1].ocurrencias++
	} else {
		os.secuenciasNumeros = append(os.secuenciasNumeros, secuencia{elem, 1})
	}
	return pos, nil
}

func main() {
	//sl := []int{}
	//os := New(sl)
	//fmt.Println(isEmpty(os))
	//fmt.Println(FrontElemen(o))
	//fmt.Println(LastElement(os))

	s := []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 4, 4, 4, 4, 3, 33, 33, 8888, 8888, 9, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 1, 1, 1, 1, 1, 1, 15, 15}
	o := New(s)
	fmt.Println("Dimension logica =>", Len(o))
	fe, erro := FrontElement(o)
	if erro == nil {
		fmt.Println("Front Element => ", fe)
	} else {
		fmt.Println(erro)
	}
	le, erro := LastElement(o)
	if erro == nil {
		fmt.Println("Last Element => ", le)
	} else {
		fmt.Println(erro)
	}

	fmt.Println("OptimumSlice original=> ", o)
	sliceAgain := SliceArray(o)
	fmt.Println("OptimumSlice desempaquetado =>", sliceAgain) // segunda prueba

	i, err := Insert(&o, 4, 60) // valor 4, posicion dl
	if err == nil {
		fmt.Println("Elemento insertado correctamente en la posicion ", i+1)
	} else {
		fmt.Println(err)
	}
	fmt.Println("OptimumSlice luego de la insercion=> ", o)
	s = SliceArray(o)
	fmt.Println("Oslice desempaquetado luego de la insercion=> ", s)
	fmt.Println("Dimension logica =>", Len(o))

}

/*
BREVE DESCRIPCION DE VARIABLES DE INSERT
---------------------------------------------------------------------------------
| NOMBRE	=>								SIGNIFICADO                         |

|posLog      |		=>		indice logico simulado sobre los valores expandidos |
|offset      |		=>		cuantos elementos dentro de una secuencia hasta pos |
|antes       |		=>		fragmento izquierdo de la secuencia original        |
|despues     |		=>		fragmento derecho de la secuencia original          |
|nuevaSecuen |		=>		el nuevo valor a insertar como secuencia propia     |
---------------------------------------------------------------------------------
*/
