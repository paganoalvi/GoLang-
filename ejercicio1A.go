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

type registroTemps struct {
	suma float32
	cant int
}

const cantTempLeidas = 10

var maperValoresIncorrectos map[string]int

var maperTemperaturas = map[string]registroTemps{
	"ALTA":   {0, 0},
	"NORMAL": {0, 0},
	"BAJA":   {0, 0},
}

var porcPacientesXGrupo map[string]int

var tempMax = -999

var tempMin = 999

func main1A() {
	var valorLeido float32
	maperValoresIncorrectos = make(map[string]int) // incializo maper valores incorrectos
	for range cantTempLeidas {
		//fmt.Println("Ingrese una temperatura")
		fmt.Scan(&valorLeido)
		if valorLeido > float32(tempMax) {
			tempMax = int(valorLeido)
		}
		if valorLeido < float32(tempMin) {
			tempMin = int(valorLeido)
		}
		clave, err := verificacionMaper(valorLeido)
		if err == nil {
			if clave != "FUERA_DE_RANGO" {
				registro := maperTemperaturas[clave]
				registro.cant++
				registro.suma += valorLeido
				maperTemperaturas[clave] = registro // hay que volver a asignar
			} else {
				maperValoresIncorrectos["FUERA_DE_RANGO"]++
			}
		} else {
			fmt.Printf("%v\n", err)
		}
	}

	porcPacientesXGrupo = make(map[string]int)                                // incializo maper
	calcularPorcentajePacientesxGrupo(maperTemperaturas, porcPacientesXGrupo) //& direccion de memoria de porcentajesPacientesXGrupo

	fmt.Println(porcPacientesXGrupo["ALTA"], "%", " de los pacientes se encuentra en el grupo TEMPERATURA ALTA")
	fmt.Println(porcPacientesXGrupo["NORMAL"], "%", " de los pacientes se encuentra en el grupo TEMPERATURA NORMAL")
	fmt.Println(porcPacientesXGrupo["BAJA"], "%", " de los pacientes se encuentra en el grupo TEMPERATURA BAJA")

	promEnt := int((float32(tempMax+tempMin) / 2)) // se debe imprimir el promedio entero entre la temperatura máxima y la temperatura mínima
	fmt.Print("Promedio entero entre la temperatura MAXIMA y la MINIMA es: ", promEnt)
	fmt.Printf("Cantidad de valores fuera del rango: %v\n", maperValoresIncorrectos["FUERA_DE_RANGO"])
}

// calculo de porcentaje de pacientes por grupo
// porcentPxG map[string]int => maper (ya se pasa por referencia, en cambio un arreglo si se pesa el puntero a ese arreglo
func calcularPorcentajePacientesxGrupo(mapTemp map[string]registroTemps, porcentPxG map[string]int) { //recibo el ptro como parametro y lo modifico
	porcentPxG["ALTA"] = int(float32((mapTemp["ALTA"].cant)) / (float32(cantTempLeidas)) * 100)
	porcentPxG["NORMAL"] = int(float32((mapTemp["NORMAL"].cant)) / (float32(cantTempLeidas)) * 100)
	porcentPxG["BAJA"] = int(float32((mapTemp["BAJA"].cant)) / (float32(cantTempLeidas)) * 100)
	// convierto resultado a int(cant a float3 y total de temp leidas tambien a float32)

}

// //alta mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36)
func verificacionMaper(num float32) (string, error) {
	switch {
	case num > 50, num < 20:
		return "FUERA_DE_RANGO", nil // temperaturas incorrectas
	case num > 37.5 && num <= 50:
		return "ALTA", nil
	case num >= 36 && num <= 37.5:
		return "NORMAL", nil
	case num >= 20 && num < 36:
		return "BAJA", nil
	default:
		return " ", errors.New("Temperatura fuera de rango")
	}

}

func pasarDeCelciusAFahrenheit(unaTemp float32) float32 {
	unaTemp = (unaTemp)*float32(9/5) + float32(32)
	return unaTemp

}
