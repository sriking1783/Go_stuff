package main

import "fmt"

func main() {
   num := 0
   for i := 999; i >= 100; i-- {
       for j := i; j >= 100; j-- {
         if num > j* i {
            break
         }
         if isPalidrome(j * i) && num < j*i {
             num = j * i 
         }
       }
   }
   fmt.Println(num)
}

func isPalidrome(number int) bool {
    temp := number
    var digits []int
    for temp > 0 {
        rem := temp % 10
        temp = temp / 10 
        digits = append([]int{rem}, digits...)
    }
    for i := 0; i < len(digits)/2; i++ {
        if digits[i] != digits[len(digits) - i -1] {
            return false
        }
    }
    return true
}
