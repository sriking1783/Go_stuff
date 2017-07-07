package main
import (
     "fmt"
     "math"
)
func main() {
    number := 0
    i := 1
 
    for(numberOfDivisors(number)< 500) {
       number = number + i
       i++;
    }
    fmt.Println(i)
    fmt.Println(number)
}

func numberOfDivisors(number int) int {
    nod := 0
    sqrt := int(math.Sqrt(float64(number)))
  
    for i := 1; i <= sqrt; i++ {
        if number % i == 0 {
            nod += 2
        }
    }
    if sqrt * sqrt == number {
        nod += 2
    }
    return nod;
}

