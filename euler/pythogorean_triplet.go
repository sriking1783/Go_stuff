package main
import (
      "fmt"
)

func main() {
    for a := 1; a < 998; a++ {
        for b := 998; b > a ; b-- {
           c := 1000 - a - b
           if a*a + b*b == c*c {
               fmt.Println(a," ",b," ",c)
               fmt.Println(a*b*c)
           }
           
        }
    }
}
