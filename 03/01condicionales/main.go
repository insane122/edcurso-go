package main

import "fmt"


// || Ok
// && And
//  ! Not

func main() {
//  isvalid := false

if edad, Nombre := 17, "alvarito"; edad < 14 {
  fmt.Println(Nombre," Es un niño ")
} else if edad < 18 {

  fmt.Println(Nombre," Es un adolescente")
} else if edad < 30 {
  fmt.Println(Nombre," Es un adulto")
} else {  fmt.Println("Es un adulto mayor")
}
/* if isvalid {

  fmt.Println("esto esta en el bloque de verdadero")
  fmt.Println(isvalid)
  if 10 < 5 {
  fmt.Println("5 es menor a 10")
  } else {
  fmt.Println("5 no es menor a 10")}
  } else {
    fmt.Println("Esto esta en el bloque de falso")
    fmt.Println(isvalid)
  }
  */
  fmt.Println("Fin del programa")

}
