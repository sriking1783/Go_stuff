package stack


type Stack struct {
    Symbol string
}
var symbol_stack []Stack
//var symbol_stack []string

func Push(symbol string) {
    symbol_string := new(Stack)
    symbol_string.Symbol = symbol
    symbol_stack = append(symbol_stack, *symbol_string)
}

func Peek() *Stack {
    i := len(symbol_stack)
    if i < 1 {
       return nil
    }
    symbol := symbol_stack[i-1]
    return &symbol
}

func Pop() *Stack {
    i := len(symbol_stack)
    if i < 1 {
       return nil
    }
    symbol := symbol_stack[i-1]
    symbol_stack = symbol_stack[:i-1]
    return &symbol
}

