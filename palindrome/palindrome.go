package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "math"
func main() {
       reader := bufio.NewScanner(os.Stdin)
       reader.Scan()
       text := reader.Text()
       palindrome_number := int64(0)
       number_of_lines, _ := strconv.Atoi(text)
       var numbers []int64
       for i := 0; i < number_of_lines; i++ {
               reader.Scan()
               text = reader.Text()
               number, err := strconv.ParseInt(text, 10, 64)
               if err != nil {
    	                fmt.Println("ERRORED ",err)
	       } 
               numbers = append(numbers, number)
       } 
       fmt.Println() 
       for _, val := range numbers {
           digits := convertToDigits(val)
           palindrome_number = convertToNumber(mirrorNumber(digits))
           if palindrome_number <= val {
                   incremented_digits := incrementNumber(digits)
                   palindrome_number = convertToNumber(mirrorNumber(incremented_digits))
           }
           fmt.Println(palindrome_number)
       }
}

func incrementNumber(digits []int) []int {
      n := 0
      if len(digits)%2 == 0 {
          n = len(digits)/2 - 1
      } else {
          n = len(digits)/2
      }
    
      for digits[n] == 9 && n > 0 {
          digits[n] = 0
          n = n-1
      }
      digits[n] = digits[n] + 1
      if(digits[n] > 9) {
         digits1 := make([]int, len(digits)+1)
         digits1[0] = 1
         digits1[1] = 0

         for i := 1; i < len(digits); i++ {
             digits1[i+1] = digits1[i]
         }
         return digits1
      }
      return digits
}

func convertToNumber(digits []int) int64 {
        number := int64(0)
        for i := len(digits)-1; i>=0; i-- {
             number = number + int64(digits[i]) * int64(math.Pow(10, float64(i)))
        }
        return number
}

func mirrorNumber(digits []int) []int {
        for i := 0; i < len(digits)/2; i++ {
            digits[len(digits) - i - 1] = digits[i]
        }
        return digits
}

func convertToDigits(number int64) []int {
        var digits []int
        temp := number
        rem := 0
        for temp > 0 {
                temp = number / 10
                rem = int(number % 10)
                digits = append([]int { rem },digits...)
                number = temp
        }
        return digits
}
