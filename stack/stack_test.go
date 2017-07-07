package stack
import (
        _"fmt"
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
         Push("*")
         assert.Equal(t, 1, len(symbol_stack), "")
         symbol_string := Pop()
         assert.Equal(t, "*", symbol_string.Symbol ,"")
         assert.Equal(t, 0, len(symbol_stack), "")
}

func TestPush(t *testing.T) {
         Push("*")
         assert.Equal(t, 1, len(symbol_stack), "")
         Pop()
}

func TestPeek(t *testing.T) {
         Push("*")
         assert.Equal(t, 1, len(symbol_stack), "")
         symbol_string := Peek()
         assert.Equal(t, "*", symbol_string.Symbol ,"")
         assert.Equal(t, 1, len(symbol_stack), "")
}
