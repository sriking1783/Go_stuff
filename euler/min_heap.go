package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    //"math/big"
    "time"
)
var heap [5164]string
var heapSize int
func initMinHeap() {
        heapSize = 0
        heap[0] = ""
        
}

func insert(element string) {
        heapSize++
        //fmt.Println(heapSize)
        heap[heapSize] = element

        now := heapSize
        for heap[now/2] > element {
                heap[now] = heap[now/2]
                now = now/2
        }
        heap[now] = element
}


func deleteMin() string{
        minElement := heap[1]
        heapSize = heapSize-1
        lastElement := heap[heapSize]
        child := 0
        now := 0
        for now = 1; now*2 <= heapSize; now = child {
                child = now * 2
                
                if(child != heapSize && heap[child+1] < heap[child]) {
                        child++
                }
                
                if(lastElement > heap[child]) {
                        heap[now] = heap[child]
                } else {
                        break
                }
        }

        heap[now] = lastElement
        return minElement
}

func main() {
        f, _ := os.Open("/home/vagrant/Go/src/exercises/euler/p022_names.txt")         
        t := time.Now()

        scanner := bufio.NewScanner(f)

        for scanner.Scan() {
                line := scanner.Text()
                names := strings.Split(line, ",")
                //fmt.Println(line)
                initMinHeap()
                for _, val := range names {
                        element := strings.Replace(val,"\"","",-1)
                        element = strings.Replace(element," ","",-1)
                        //fmt.Print(element,",")
                        insert(element)
                }
        }
        result := int64(0)
        elem := ""
        for i := 0 ; i < 5163; i++ {
                temp := deleteMin()
                if elem != temp {
                        elem = strings.TrimSpace(temp)
                        //elem = temp
                        byteArray := []byte(elem)
                        worth := int64(0)
                        for _, val := range byteArray {
                              worth += int64(val) - 64
                              if int64(val) < 65 && int64(val) > 90 {
                                      fmt.Println(elem)
                              }
                        }
                       
                        result = result + int64(int64(i+1) * worth)
                        //score := int64(i+1) * worth
                        //fmt.Println(elem," ",worth," ",i," ",score," ",result)
                        //fmt.Print(score,", ")
                } else {
                        fmt.Println(temp," ",i);
                }
        }
        //temp := deleteMin()
        //fmt.Println(temp)
        

        fmt.Println(result)
        elapsed := time.Since(t)
        fmt.Println(elapsed)

}
