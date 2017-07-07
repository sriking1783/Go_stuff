package main
import "fmt"
var b[]int
var heapSize int
func main() {
        a := []int{6,7,1,2,8,9,4,5,3}
        b = make([]int, len(a)+1)
        b[0] = 12345
        for i := 0; i < len(a); i++ {
                //fmt.Println(i)
                b[i+1] = a[i]
                element := a[i]
                now := i+1
                for now > 0 && b[now/2] < element {
                        b[now] = b[now/2]
                        now = now/2
                }
                b[now] = element
        }
        heapSize = len(b)
        fmt.Println(b)
        size := len(b)
        for i := 0 ; i < size ; i++ {
                fmt.Println(deleteMax())
                //fmt.Println("Size: ",size)
                //size = size-1
        }
}

func deleteMax() int {
        maxElement := b[1]
        now := 0
        lastElement := b[len(b)-1]
        child := 0
        heapSize = heapSize-1
        fmt.Println("Len: ",len(b)) 
        for now = 1; now*2 <= heapSize; now = child {
                child = now * 2
                if(child != len(b) && b[child+1] > b[child]) {
                        child++
                }
                if(child < len(b) && lastElement < b[child]) {
                        b[now] = b[child]
                } else {
                        break
                }
        }
        b[now] = lastElement
        return maxElement
}
