/*2. Implementar la función factorial de dos formas, una iterativa y otra
recursiva. Escribir un programa y compilar de forma que utilice una u
otra y la evalúe de 0 a 9. La función factorial se define como:*/

package main

func main(){
	factorialIterativo(5)
}
func factorialIterativo(n int){
	f:=1
	s:=n
	for i := n ; i > 1 ; i--{
		println(f," +  (",s,  " *  ",i-1,")")
		f= (s*(i-1))
		s=f;
	}
	println("factorial es: ",f)
}