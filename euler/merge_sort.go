package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    //"math/big"
    "time"
)

func main() {
        array := make([]string, 5164)
        //array := []string{"ant", "harry", "ball", "henry", "chad"}
        f, _ := os.Open("/home/vagrant/Go/src/exercises/euler/p022_names.txt")         
        t := time.Now()

        scanner := bufio.NewScanner(f)
        i := 0
        for scanner.Scan() {
                line := scanner.Text()
                names := strings.Split(line, ",")
                //fmt.Println(line)
                for _, val := range names {
                        element := strings.Replace(val,"\"","",-1)
                        element = strings.Replace(element," ","",-1)
                        //fmt.Print(element,",")
                        array[i] = element
                        i += 1
                }
        }
        //fmt.Println(array)
        array1 := MergeSort(array)
        //fmt.Println(array1)
        result := int64(0)
        elem := ""
        for i, value := range array1 {
                worth := int64(0)
                if value != elem {
                    elem = strings.TrimSpace(value)
                    byteArray := []byte(elem)
                    for _, val := range byteArray {
                            worth += int64(val) - 64
                    }
                    result = result + int64(int64(i) * worth)
                    /*if i > 1000 {
                            break
                    }*/
                }
               
        }
        elapsed := time.Since(t)
        fmt.Println(result)
        fmt.Println(elapsed)
}

func MergeSort(arr[] string) []string {
        if(len(arr) <= 1) {
                return arr
        }

        first := make([]string, len(arr)/2);
        second := make([]string, len(arr) - len(first));
        copy(first, arr[0:len(first)])
        copy(second, arr[len(first):len(first)+len(second)])

        MergeSort(first)
        MergeSort(second)
        Merge(first, second, &arr)
        return arr
}

func Merge(first [] string, second[] string, arr* []string) {
        iFirst := 0
        iSecond := 0
        iMerged := 0

        for iFirst < len(first) && iSecond < len(second) {
                if(first[iFirst] < second[iSecond]) {
                     	(*arr)[iMerged] = first[iFirst]
                        iFirst += 1
                } else {
                	(*arr)[iMerged] = second[iSecond]
                        iSecond += 1
                }
                iMerged += 1
        }
        
        for i := 0; i < len(first) - iFirst; i++ {
        	(*arr)[iMerged] = first[i+iFirst] 
                iMerged += 1
        }

        for i := 0; i < len(second) - iSecond; i++ {
        	(*arr)[iMerged] = second[i+iSecond] 
                iMerged += 1
        }
}

