package main
import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
       reader := bufio.NewScanner(os.Stdin)
       reader.Scan()
       text := reader.Text()
       palindrome_number := ""
       number_of_lines, _ := strconv.Atoi(text)
       var numbers []string
       for i := 0; i < number_of_lines; i++ {
               reader.Scan()
               text = reader.Text()
               j := 0
               for text[0] == 48 {
                   j = j+1
                   text = text[j:]
               }
               numbers = append(numbers, text)
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

func incrementNumber(digits []string) []string {
      n := 0
      if len(digits)%2 == 0 {
          n = len(digits)/2 - 1
      } else {
          n = len(digits)/2
      }
    
      for digits[n] == "9" && n > 0 {
          digits[n] = "0"
          n = n-1
      }
      temp_n, _  := strconv.Atoi(digits[n])
      digits[n] =  strconv.Itoa(temp_n + 1)
      if(temp_n + 1 > 9) {
         digits1 := make([]string, len(digits)+1)
         digits1[0] = "1"
         digits1[1] = "0"

         for i := 1; i < len(digits); i++ {
             digits1[i+1] = digits1[i]
         }
         return digits1
      }
      return digits
}

func convertToNumber(digits []string) string {
        number := ""
        for i := 0; i < len(digits); i++ {
             number = number + digits[i]
        }
        return number
}

func mirrorNumber(digits []string) []string {
        for i := 0; i < len(digits)/2; i++ {
            digits[len(digits) - i - 1] = digits[i]
        }
        return digits
}

func convertToDigits(number string) []string {
        digits := make([]string, len(number))
        for i, c := range number {
             digits[i] = fmt.Sprintf("%c", c)
        }
        return digits
}
