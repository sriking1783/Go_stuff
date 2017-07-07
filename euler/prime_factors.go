package main
import ( 
       "fmt"
       "math"
)
func main() {
    lowerBound := 2
    upperBound := 775146 //600851475143

    isComposite := runEratosthenesSieve(2, 775146)
    fmt.Println(lowerBound," ", upperBound," ", len(isComposite))
    for i := upperBound-1; i > 0; i-- {
        if !isComposite[i] {
            if (600851475143 % i == 0) {
                fmt.Print(i,", ")
                break
            } 
        }
    }
}

func runEratosthenesSieve(lowerBound int, upperBound int) [] bool {
    upperBoundSquareRoot := int(math.Sqrt(775146))
    fmt.Println(lowerBound," ", upperBound," ",math.Sqrt(775146))
   
    isComposite := make([]bool, upperBound)
    for m := 2; m <= upperBoundSquareRoot; m++ {
        if !isComposite[m] {
            for k := m*m; k < upperBound; k = k+m {
                 isComposite[k] = true
            }
        }
    }
    return isComposite
}
