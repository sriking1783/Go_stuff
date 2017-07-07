package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)


func main() {
    var number_of_examples int
    _, err := fmt.Scanf("%d", &number_of_examples)  

    if err != nil {
        fmt.Println(number_of_examples)
    }
    numbers := make([]int, number_of_examples)
    for i := 0; i < number_of_examples; i++ {
        reader := bufio.NewReader(os.Stdin)
        text, _ := reader.ReadString('\n')
        line := strings.Split(text, " ")
        a, _ := strconv.Atoi(line[0])
        b, _ := strconv.Atoi(line[1][:1])
        
        numbers[i] = get_pow(a, b)
    }

    fmt.Println()

    for i, _ := range numbers {
        fmt.Println(numbers[i])
    }
}

func get_pow(a int, b int) int {
    prod := 1
    if a > 0 && b > 0  {
        n := a%10
        if n == 1 || n == 5 || n == 6 {
            return n
        }

        if n == 2 || n == 3 || n == 7 || n == 8{
            if b % 4 == 1 {
                return n
            }

            if b % 4 == 2 {
                return n *n % 10
            }
            if b % 4 == 3 {
                return n*n*n % 10
            }
            if b % 4 == 4 {
                return n*n*n*n % 10
            } 
             
        }

        if n == 4  || n == 9{
            if b % 2 == 0 {
                return n*n %10
            }else {
                return n % 10
            }
        }
        for b >= 1 {
            prod = (prod*a)%10
            b--
        }
    } 

    return prod
}
