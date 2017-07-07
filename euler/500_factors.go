package main
import (
     "fmt"
)
var m map[int][]int
func main() {
  sum := 0;
  m = make(map[int][]int)
  for i := 1; i <= 50000; i++ {
      sum = sum + i
      count := num_factors(sum)
      if count >= 500 {
          fmt.Println(i," ",sum," ",count)
          //fmt.Println("---------------------------")
          break
      }
      go func(num int){
          if num % 1000 == 0 {
              fmt.Println("Parsed ",num, " ", count)
          }
      }(i)
  }
}

func num_factors(num int) int {
    factors := make([]int, 1)
    keys := make([]int, 0)
    for k := range m {
        keys = append([]int{k}, keys...)
    }

    factors[0] = 1 
    for _, v := range keys {
        if v > 1 && num % v == 0 {
            value1, ok1 := m[v]
            if ok1 {
                    for _, n := range value1 {
                        if contains(factors, n) != true {
                            factors = append(factors, n)
                        }
                    }
            }
            divisor := num / v
            value, ok := m[divisor]
            if ok {
                for _, n := range value {
                    if contains(factors, n) != true {
                        factors = append(factors, n)
                    }
                } 
                m[num]=factors
                return len(factors)
            }
        }
    }

    go func(number int){
        for i := 2; i <= number; i++ {
            if num % i == 0 {
                if contains(factors, i) != true {
                    factors = append(factors, i)
                }
            }
        }
    }(num)

    m[num] = factors
    return len(factors)
}

func contains(a []int, e int) bool {
    for _, n := range a{
        if e == n {
            return true
        }
    }
    return false
}
