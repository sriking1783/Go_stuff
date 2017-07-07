package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
   
  number_array := [100][]int64{}
  for i:=0 ; i < 100; i++ {
      reader := bufio.NewReader(os.Stdin)
      input, _ := reader.ReadString('\n')
      numbers_string := make([]string, i+1, i+1)
      numbers_string = strings.Split(input, " ")
      numbers := make([]int64, i+1, i+1)
      for j, number := range numbers_string {
          number = strings.Replace(number,"\n","",-1)
          nums, err := strconv.ParseInt(number, 10, 64)
          if err != nil {
              panic(err)
          }
          numbers[j] = int64(nums)
      }
      number_array[i] = make([]int64, i+1)
      number_array[i] = numbers
      
  }

  for i:=len(number_array)-2; i >=0; i-- {
      temp_array := number_array[i+1]
      for j := 0 ; j<len(number_array[i]); j++ {
          number_array[i][j] = max_of_numbers(number_array[i][j]+temp_array[j], number_array[i][j]+temp_array[j+1])
          //fmt.Print(number_array[i][j]," ")
      }
      //fmt.Println()
  }
  fmt.Println(number_array[0][0])

}

func max_of_numbers(num1, num2 int64) int64 {
    if num1 > num2 {
        return num1
    }
    return num2
}
