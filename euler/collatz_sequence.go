package main
import "fmt"
func main() {
    i := 1
    /*for i = 1; i < 1000000;  {
        //fmt.Println(i)
        temp1 := 0
        temp2 := 0
        if((i - 1)%3 == 0 && (i - 1)/3 > 1 && i > 1) {
             temp1 = (i - 1)/3
        }  
        temp2 = i*2
        if(temp1 < temp2 && temp1 > 0) {
            i = temp1                                                                                                                                       
        } else {
            i = temp2
        } 
        fmt.Println(i)
    }*/
    count := 0
    whole_count := 0
    number := 0
    for j:= 999999; j > 1; j-- {
        count = 0
        for i = j; i > 1; {
            if(i % 2 == 0) {
                i = i / 2
            } else {
                i = 3 * i + 1
            }
            count++
            //fmt.Println(i)
        }
        
        if count > whole_count {
            whole_count = count
            number = j
            //fmt.Println(whole_count)
        }
    }
    fmt.Println(number)
}
