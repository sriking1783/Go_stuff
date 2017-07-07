package main

import(
       "fmt"
       "bufio"
       "os"
       "strconv"
       "strings"
)

func calculate(operator1 int64, operator2 int64, operand string) int64 {
       switch operand {
               case "+":
                      return operator1 + operator2
               case "-":
                      return operator1 - operator2
               case "*":
                      return operator1 * operator2
       }
       return 0
}

func main() {
       reader := bufio.NewScanner(os.Stdin)
       reader.Scan()
       text := reader.Text()
       text_length := 0
       total_length := 0
       number_of_lines, _ := strconv.Atoi(text)
       var arithmetic_expressions []string
       for i := 0; i < number_of_lines; i++ {
               reader.Scan()
               text = reader.Text()
               arithmetic_expressions = append(arithmetic_expressions, text)
       }
       var w []string
       operator := ""
       for _, val := range arithmetic_expressions {
               text_length = 0
               total_length = 0
               w = strings.FieldsFunc(val, func(r rune) bool {
                          switch r {
                                case '+', '-', '*', '/':
                                      operator = string(r)
                                      return true
                          }
                          return false
               })
               for _, num := range w {
                   total_length += len(num)
                   if text_length < len(num) {
                        text_length = len(num)
                   }
               }
               n1, _ := strconv.ParseInt(w[0], 10, 64)
               n2, _ := strconv.ParseInt(w[1], 10, 64)
               n3 := int64(0)
               if operator == "*" {
                   for j := len(w[1]); j > 0; j-- {
                       temp, _ := strconv.ParseInt(string(w[1][j-1]), 10, 64)
                       n3 = calculate(n1, temp, operator)
                       w = append(w, strconv.FormatInt(n3, 10))
                   }
               } else {
                   n3 = calculate(n1, n2, operator)
                   w = append(w, strconv.FormatInt(n3, 10))
               }
               text_length = text_length + 1
               result_length := text_length
               if operator == "*" {
                   text_length = total_length
               }
               temp_length := text_length
               pad := 0
               for i, num := range w {
                   pad =  text_length - len(num)
                   if i > 1 {
                       pad =  temp_length - len(num)
                       temp_length -= 1
                       if pad < 0 {
                           pad = 0
                       }
                       fmt.Println(strings.Repeat(" ", pad)+num)
                   } else if i == 1 {
                       fmt.Println(strings.Repeat(" ", pad-1)+operator+ num)
                       if operator == "*" {
                           fmt.Print(strings.Repeat(" ", result_length-pad))
                       }
                       for j := 0; j < result_length; j++ {
                             fmt.Print("-")
                       }
                       fmt.Println()
                   } else {
                       fmt.Println(strings.Repeat(" ", pad)+num)
                   }
               }

               if operator == "*" && len(string(w[1])) > 1 {
                   n3 = calculate(n1, n2, operator)
                   for j := 0; j < text_length; j++ {
                       fmt.Print("-")
                   }
                   fmt.Println()
                   fmt.Println(strings.Repeat(" ", pad)+strconv.FormatInt(n3, 10))
               }
               fmt.Println()
       }
}
