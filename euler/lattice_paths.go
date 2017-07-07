package main

var limita int
var limitb int
var cache
func main() {
    limita = 1
    limitb = 1
    cache = [2][2]int{}
    row = 
    for i := 0; i <= limita; i++ {
        for j := 0; j <= limitb; j++ {
            cache[i][j] = -1
        }
    }  

    fmt.Println(start(limita, limitb))    
}

func start(a int, b int) int {
   i := 0
   if a == limita && b == limitb {
       return 1
   } 

   
}
