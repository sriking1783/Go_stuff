package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)
func main() {
    grid := make([][]int, 20)
    for i := 0; i < 20; i++ {
        reader := bufio.NewReader(os.Stdin)
        text, _ := reader.ReadString('\n')
        numbers := strings.Split(text, " ")
        grid[i] = make([]int, len(numbers))
        k := 0
        for _, n := range numbers {
              j, _ := strconv.ParseInt(n, 10, 64)
              grid[i][k] = int(j)
              k += 1 
        } 
    }    
    product := 0
    for i := 0; i < 20; i++ {
        for j := 0; j < 16; j++ {
           if grid[i][j] !=0 && grid[i][j+1] != 0 && grid[i][j+2] != 0 && grid[i][j+3] != 0 {
               temp_product := grid[i][j] * grid[i][j+1] * grid[i][j+2] * grid[i][j+3]
               if product < temp_product {
                   product = temp_product
               }
          } 
        }
    } 
    for i := 0; i < 16; i++ {
        for j := 0; j < 20; j++ {
           if grid[i][j] !=0 && grid[i+1][j] != 0 && grid[i+2][j] != 0 && grid[i+3][j] != 0 {
               temp_product := grid[i][j] * grid[i+1][j] * grid[i+2][j] * grid[i+3][j]
               if product < temp_product {
                   product = temp_product
               } 
           }
        }
    } 
    for i := 16; i >= 0; i-- {
        for j := 16; j >= 0; j-- {
           if grid[i][j] !=0 && grid[i+1][j+1] != 0 && grid[i+2][j+2] != 0 && grid[i+3][j+3] != 0 {
               temp_product := grid[i][j] * grid[i+1][j+1] * grid[i+2][j+2] * grid[i+3][j+3]
               if product < temp_product {
                   product = temp_product
               }
          } 
        }
    }

    for i := 16; i >= 0; i-- {
        for j := 19; j >= 3; j-- {
               //fmt.Println(i," ",j)
           if grid[i][j] !=0 && grid[i+1][j-1] != 0 && grid[i+2][j-2] != 0 && grid[i+3][j-3] != 0 {
               temp_product := grid[i][j] * grid[i+1][j-1] * grid[i+2][j-2] * grid[i+3][j-3]
               if product < temp_product {
                   product = temp_product
               }
           }
        }
    }
    fmt.Println("Largest product is ", product)
}
