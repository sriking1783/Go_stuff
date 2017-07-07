package main
import ( 
       "fmt"
)
func main() {
    num := 1
    count := 0
    i := 0
    nth := 10001
    for count < nth {
        num = num+1;
        for i = 2; i <= num; i++ {
             if num % i == 0 {
                break;
             }
        }
        if i == num {
            count = count+1
        }
    }
    fmt.Println(num)
}

