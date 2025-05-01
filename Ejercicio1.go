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
	"fmt"
	"errors"
)

//alta (mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36)
type Celsius float64
type Fahrenheit float64
var temperaturas [10]Celsius

func main() {
	var valorLeido Celsius
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese una temperatura")
		fmt.Scan(&valorLeido)
		pos,err := verificacion(valorLeido)
		if(err == nil){
			temperaturas[pos].suma += valorLeido
			temperaturas[pos].cant++
		}else{
			fmt.Printf("%v\n", err)
		}
	}
	
	cantidad := temperaturas[0].cant + temperaturas[2].cant
	sumaTotal := temperaturas[0].suma + temperaturas[2].suma
	promEnt := int((sumaTotal / Celsius(cantidad)))
	fmt.Print("El promedio entre maxima y minima es: ", promEnt)
}

////alta mayor de 37.5), normal (entre 36 y 37.5) y baja (menor de 36)
func verificacion(num float32) (int,error) {
	if num > 37.5 && num < 50 {
		return 0,nil
	} else if (num >= 36) && (num <= 37.5) {
		return 1,nil
	} else if(num>=20) {
		return 2,nil
	}else{
		return -1, errors.New("temperatura fuera de rango")
	}
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
} 
