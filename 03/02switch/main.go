package main

import "fmt"


func main() {


//  var a uint8

//  a = 98
  /*
  switch a {
  case 1:

    fmt.Println(" lunes ")

  case 2:

    fmt.Println("martes")

  case 3:

      fmt.Println("miercoles")

    case 4:

        fmt.Println("jueves")

      case 5:

          fmt.Println("viernes")

        case 6:

            fmt.Println("sabado")

          case 7:

              fmt.Println("domingo")

          default :

    fmt.Println("no es un dia de la semana")


  }
  */



/*
  switch a {
  case 1:
    fallthrough
  case 2:
      fallthrough

    case 3:
        fallthrough

      case 4:
          fallthrough

        case 5:
            fmt.Println("estas entre semana")

          case 6:
             fallthrough

            case 7:
                fmt.Println("estas en el fin de semana")

              default:
                fmt.Println("no es un dia valido")

  }
  */

switch a:= 5; {
case a >= 0 && a <= 5:
  fmt.Println("estas entre semana")

case a >= 6 && a <= 7:
  fmt.Println("estas en el fin de semana")

  default :
  fmt.Println("no es un dia valido")

}
}
