/* Las temperaturas de los pacientes de un hospital se dividen en 3
grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de
	36). Se deben leer 10 temperaturas de pacientes e informar el
	porcentaje de pacientes de cada grupo. Luego se debe imprimir el
	promedio entero entre la temperatura máxima y la temperatura mínima.
	Resolver cargando primero todos los valores usando un arreglo y
	almacenar los datos en variables escalares como acumuladores y
	contadores. Probar generar archivos de entrada con los valores y
	ejecutar, por ejemplo, de la siguiente forma:

	go run p2-1.go < input2-1.txt

	a) Volver a resolver pero usando un arreglo o un Map de tres
	posiciones donde se acumulan los valores de cada grupo.

	b) Modificar la solución para incluir grupo de valores incorrectos, como
	pueden ser los mayores a 50◦ y los menores a 20◦.

	c) Escribir una función que pasa de grados Celsius a Fahrenheit
	utilizando nuevos tipos y aplicarla al arreglo de los valores leídos. La
	conversión se realiza de acuerdo a la siguiente ecuación: */
// un vector 3 posiciones (alta,media,baja)de registros donde cada registro tenemos 2 campos
// uno para acumular la temperatura y otro para llevar cuenta de la cantidad leida
package main

import (
	"errors"
	"fmt"
)

//alta (mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36)

type registroTemperatura struct {
	suma float32
	cant int
}

const cantidadTemperaturasLeidas = 10

var temperaturas [3]registroTemperatura
var porcentajesPacientesXGrupo [3]int
var max = -999
var min = 999

func main() {
	var valorLeido float32
	for i := 0; i < cantidadTemperaturasLeidas; i++ {
		fmt.Println("Ingrese una temperatura")
		fmt.Scan(&valorLeido)
		if valorLeido > float32(max) {
			max = int(valorLeido)
		}
		if valorLeido < float32(min) {
			min = int(valorLeido)
		}
		pos, err := verificacion(valorLeido)
		if err == nil {
			temperaturas[pos].suma += valorLeido
			temperaturas[pos].cant++
		} else {
			fmt.Printf("%v\n", err)
		}
	}
	fmt.Println(temperaturas[0].cant)
	fmt.Println(temperaturas[1].cant)
	fmt.Println(temperaturas[2].cant)

	//fmt.Println("Cantidad de temperatuas leidas: ",cantidad) // pa debug

	//sumaTotal := temperaturas[0].suma + temperaturas[1].suma + temperaturas[2].suma
	calcularPorcentajePacientesPorGrupo(temperaturas, &porcentajesPacientesXGrupo) //& direccion de memoria de porcentajesPacientesXGrupo
	//fmt.Println("Suma de todas las temperaturas leidas:‌ ",sumaTotal) // pa debug
	fmt.Println(porcentajesPacientesXGrupo[0], "%", " de los pacientes se encuentra en el grupo TEMPERATURA ALTA")
	fmt.Println(porcentajesPacientesXGrupo[1], "%", " de los pacientes se encuentra en el grupo TEMPERATURA NORMAL")
	fmt.Println(porcentajesPacientesXGrupo[2], "%", " de los pacientes se encuentra en el grupo TEMPERATURA BAJA")

	//fmt.Println("Max:‌ ", max)
	//fmt.Println("Min: ", min)

	promEnt := int((float32(min+max) / 2)) // se debe imprimir el promedio entero entre la temperatura máxima y la temperatura mínima
	fmt.Print("Promedio entero entre la temperatura MAXIMA y la MINIMA es: ", promEnt)
}

// inciso a
// porcentPxG *[3]int => Puntero a arreglo porcentajesPacientesXGrupo
func calcularPorcentajePacientesPorGrupo(temp [3]registroTemperatura, porcentPxG *[3]int) { //recibo el ptro como parametro y lo modifico
	porcentPxG[0] = int(float32((temp[0].cant)) / (float32(cantidadTemperaturasLeidas)) * 100)
	porcentPxG[1] = int(float32((temp[1].cant)) / (float32(cantidadTemperaturasLeidas)) * 100)
	porcentPxG[2] = int(float32((temp[2].cant)) / (float32(cantidadTemperaturasLeidas)) * 100)
	// convierto resultado a int(cant a float3 y total de temp leidas tambien a float32)

}

// //alta mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36)
func verificacion(num float32) (int, error) {
	if num > 37.5 && num < 50 {
		return 0, nil
	} else if (num >= 36) && (num <= 37.5) {
		return 1, nil
	} else if (num >= 20) && (num < 36) {
		return 2, nil
	} else {
		return -1, errors.New("temperatura fuera de rango")
	}
}
