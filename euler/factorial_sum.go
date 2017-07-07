package main
import (
        "fmt"
        "math/big"
        "strconv"
)

func main() {
    n := big.NewInt(36288)

    for i := 11; i < 100; i+=1 {
        temp := int64(i)
        if temp%10 == 0 {
                temp = temp/10
        }

        n.Mul(big.NewInt(temp),n)
    }

    bigstr := n.String()
    count := 0
    for i, _ := range bigstr {
          s := string([]rune(bigstr)[i])
          j,_ := strconv.Atoi(s)
          count += j
    }
    fmt.Println(count)
}
