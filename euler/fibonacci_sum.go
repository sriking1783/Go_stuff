package main
import "fmt"

func main() {
    num1 := 1
    num2 := 2
    sum := 0
    for num1 < 4000000 {
      temp := num2
      num2 = num1 + num2
      num1 = temp
      if num1%2 == 0{
          sum = sum + num1
      }
    }
    fmt.Println()
    fmt.Println(sum)
}
