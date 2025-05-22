package main
import "fmt"

func main() {
    /*holaMundo()
		ejercicio3()
    ejercicio4()*/
    fmt.Println("Ingrese valor de x: ")
    x := 0
    fmt.Scan(&x)
    ejercicio5(x)
 }


func holaMundo() {    
  fmt.Println("Hello world")
}
func ejercicio3() {
  /* integers */    
		/*var zz int = 0*/ // se comenta esta linea porque zz nunca es usada
    x := 10    
    z := x
    fmt.Println("z: ", z)    
    fmt.Println("x: ", x)    
    var y int = x + 1    
    fmt.Println("y: ", y)    
    const n = 5001    
    fmt.Println("n: ", n)    
    const c = 5001    
    fmt.Println("c: ", c)    
    /* float */    
    var e float32 = 6    
    fmt.Println("e: ", e)    
    var f float32 = e    
    fmt.Println("f: ", f)

}



func ejercicio4() {
  const tope = 250 
  /*Inciso a*/
  fmt.Println("Inciso a: ")    
  suma := 0
  for i := 0; i <= tope; i = i + 2 {        
    suma = suma + i    
  }    
  fmt.Println("La suma de todos los numeros pares entre 0 y 250 (inlcusive) es:", suma)

  /*Inciso b*/
  fmt.Println("Inciso b: ")
	sumab := 0 
	for i := 250; i >= 0; i = i - 2 { /*alternativa => i -= 2*/        
	  sumab = sumab + i    
  }    
  fmt.Println("La suma de todos los numeros pares entre 0 y 250 (inclsuive) es:", sumab)
}

func ejercicio5(x int) {
  if x > -9999 && x < -18 {
    x = x *(-1)
  } else if  x >= -18 && x <= -1 {
      x = x % 4
  } else if  x >= 1 && x < 20 {
      x =  x ^= 2 
  } else if x >= 20 && x < 9999 {
      x = -x
  }
  fmt.Println("x= ",x)
}
  
