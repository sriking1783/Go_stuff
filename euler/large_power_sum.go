package main
import (
     "fmt"
     "math/big"
     "strconv"
)
func main() {
   n := big.NewInt(32768)
   j := 0
   for i := 15; i < 1000; i+=i {
      if 2*i > 1000 {
         j = i
         break;
      }
      temp := n
      temp.Mul(temp, n)
      n = temp
      j = i
   }
   for i := j+1 ; i<=1000; i++ {
       if i > 1000 {
          break;
       }
       temp := n
       temp.Mul(temp, big.NewInt(2))
       n = temp
       j = i
   }

   count := 0
   bigstr := n.String()
   for i, _ := range bigstr {
       //fmt.Printf("%d", int(rune(elem)))
       s := string([]rune(bigstr)[i])
       j,_ := strconv.Atoi(s)
       count += j
   }
   fmt.Println()
   fmt.Println(count)
   fmt.Println()
   fmt.Printf("%c","HELLO"[1])
   fmt.Println()
}

