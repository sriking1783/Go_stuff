package main

import ("fmt"
        "math"
)

func main() {
    n := 200
    segmentedSieve(n) 
}

func segmentedSieve(n int) {
    limit := math.Floor(math.Sqrt(float64(n)))
    sum := 0
    fmt.Println(limit)

    var prime []int
    simpleSieve(int(limit), &prime)
    
    for _, num := range prime {
        sum += num
    }
 
    low := int(limit);
    high := int(2 * limit);
    for low < n {
        mark := make([]bool, int(limit))
        mark[0] = true
        for bp := 1; bp < len(mark); bp *= 2 {
            copy(mark[bp:], mark[:bp])
        }

        for i := 0; i < len(prime); i++ {
            loLim := int(math.Floor(float64(low/prime[i])) * float64(prime[i]))
            //fmt.Println("loLim : ",loLim," low : ",low," high : ",high)
            if loLim < int(low) {
                loLim += prime[i]
            }
            
            for j := loLim; j < int(high); j += prime[i] {
                fmt.Println("j : ",j," low : ",low," j-low : ",(j-int(low)))
                mark[j-int(low)] = false
            }
        }
        for i := low; i < int(high); i++ {
            if mark[i - int(low)] == true {
                sum += i
            }
        }
        low  = low + int(limit);
        high = high + int(limit);
        if int(high) >= n {
            high = n
        } 
    }
    fmt.Println("Sum is ",sum)
}

func simpleSieve(limit int, prime *[]int) {
    mark := make([]bool, limit)
    mark[0] = true
    for bp := 1; bp < len(mark); bp *= 2 {
        copy(mark[bp:], mark[:bp])
    }


    for p := 2; p * p < limit; p++ {
        if mark[p] == true {
            for i := p*2 ; i < limit; i+=p {
                mark[i] = false
            }
        }
    } 

    for p := 2; p < limit; p++ {
       if mark[p] == true {
          *prime = append(*prime, p)
          //fmt.Println(p)
       }
    }
}
