package main

import(
       "fmt"
       "bufio"
       "os"
       "strconv"
 s "exercises/stack"
)

const (
        lparen = iota
        minus
        plus
        divide
        multiply
        power
        rparen
        operand
)

func getPrecedence(symbol string) int {
       switch symbol {
               case "(":
                   return lparen
               case ")":
                   return rparen
               case "+", "-":
                   return plus
               case "*","/":
                   return divide
               case "^":
                   return power
               default:
                   return operand
       }
}

func isOperator(operator string) bool {
        operators := [7]string {"(", ")", "+", "-", "*", "/", "^"}
        for _, val := range operators {
            if val == operator {
                return true
            }
        }
        return false
}

func canInsert(symbol string) bool {
    if s.Peek() == nil && symbol != ")"{
        return true
    } else if symbol == "(" {
        return true
    } else {
        return false
    }
    return false
}

func postfix(expression string) {
       for _, c := range expression {
             if isOperator(string(c)) {
                if canInsert(string(c)) {
                    s.Push(string(c))
                } else if string(c) == ")" {
                    prev_symbol := s.Peek()
                    for prev_symbol.Symbol != "(" {
                        s.Pop()
                        fmt.Print(prev_symbol.Symbol)
                        prev_symbol = s.Peek()
                    }
                    s.Pop()
		} else {
                    prev_symbol := s.Peek()
                    for prev_symbol != nil {
                      if getPrecedence(string(c)) < getPrecedence(prev_symbol.Symbol) {
                          s.Pop()
                          fmt.Print(prev_symbol.Symbol)
                      } else {
                         s.Push(string(c))
                         break
                      }
                      prev_symbol = s.Peek()
                    }
                }
             } else {
                     fmt.Print(string(c))
             }
       }
}

func popStack() {
    symbol := s.Pop()
    for symbol != nil {
        fmt.Print(symbol.Symbol)
        symbol = s.Pop()
    }
    fmt.Println()
}

func main() {
       reader := bufio.NewScanner(os.Stdin)
       reader.Scan()
       text := reader.Text()
       number_of_lines, _ := strconv.Atoi(text)
       var infix_expressions []string

       for i := 0; i < number_of_lines; i++ {
               reader.Scan()
               text = reader.Text()
               infix_expressions = append(infix_expressions, text)
               postfix(text)
               popStack()
       }
}
