package main

import (
	"fmt"
        "math"
)

func main() {
        var a [100]bool
        sliceA := make([]bool, len(a)) 

        limit := 10
        lower_limit := 1
        upper_limit := 10
       
        segment(lower_limit, upper_limit, &sliceA)
        for i := 0; i < len(a) && upper_limit < len(a); i++ {
                lower_limit = upper_limit + 1
                upper_limit = upper_limit + limit
                if upper_limit >= len(a) {
                       upper_limit = len(a)
                }
                segment(lower_limit, upper_limit, &sliceA)
        }
        for i := 0; i < len(a); i++ {
               fmt.Println(i," ",sliceA[i])
        }

        fmt.Println()
}

func simpleSieve(limit int) []bool {
    sliceA := make([]bool, limit)
}

func segment(start int, n int, a *[]bool) {
        end := int(float64(start) + math.Sqrt(float64(n)))
        
        for i := start; i < end && i < n; i++ {
               for j := 2 ; j < end ; j++ {
               }
               if i%3 == 0 {
                      (*a)[i]=true
               }
               if i%5 == 0 {
                      (*a)[i]=true
               }
        }
}
