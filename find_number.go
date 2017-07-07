package main

import "fmt"

func main() {
    for i := 10000; i <= 99999; i++ {
            j := i * 4
            if j == ReverseNumber(i) {
                    fmt.Println(i," ",j)
            }
    }
}

func ReverseNumber(n int) int {
    number := 0
    divider := 10000
    for  divider > 0 {
          rem := n % 10;
          number = number + rem*divider
          divider = divider/10
          n = n/10
    }

    return number
}
