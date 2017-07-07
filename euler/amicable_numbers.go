package main
import ( 
       "fmt"
)

func main() {
  numbers :=[]int{220, 284}
  sum := 0
  for i:= 9999 ; i >= 2; i--{
            sum1 := sumDivisors(i)
            sum2 := sumDivisors(sum1)
            //fmt.Println(i,", ",sum1,", ",sum2)
            if(i == sum2 && i != sum1 && !contains(numbers, i) && !contains(numbers, sum1)) {
                numbers = append(numbers, i)
                numbers = append(numbers, sum1)
            }
  }
  
  k := 0
  
  for _, j := range numbers {
      k++
      //fmt.Println(i,", ",j)
      sum = sum + j
  }

  fmt.Println("Sum: ", sum)
  fmt.Println("k: ", k)
  
}

func contains(s []int, e int) bool {
      for _, a := range s {
            if a == e {
                  return true
            }
      }
    
      return false
}

func sumDivisors(number int) int {
      sum := 0
      for i:= 1; i <= number/2; i++ {
            if number%i == 0 {
                  sum = sum + i
            }
      }


      return sum;
}
