package main

import (
  "fmt"
)
var number_words = make(map[int]string)
func main() {
    number_words[1] = "one"
    number_words[2] = "two"
    number_words[3] = "three"
    number_words[4] = "four"
    number_words[5] = "five"
    number_words[6] = "six"
    number_words[7] = "seven"
    number_words[8] = "eight"
    number_words[9] = "nine"
    number_words[10] = "ten"
    number_words[11] = "eleven"
    number_words[12] = "twelve"
    number_words[13] = "thirteen"
    number_words[14] = "fourteen"
    number_words[15] = "fifteen"
    number_words[16] = "sixteen"
    number_words[17] = "seventeen"
    number_words[18] = "eighteen"
    number_words[19] = "nineteen"
    number_words[20] = "twenty"
    number_words[30] = "thirty"
    number_words[40] = "forty"
    number_words[50] = "fifty"
    number_words[60] = "sixty"
    number_words[70] = "seventy"
    number_words[80] = "eighty"
    number_words[90] = "ninety"
    number_words[100] = "hundred"
    number_words[1000] = "thousand"
    word :=""
    count := 0
    for i := 1; i <= 1000; i++ {
        word = ""
        temp_word, num := convertToWords(i)
        word = temp_word
        for num != 0 {
            temp_word, num = convertToWords(num)
            word += temp_word
        }
        //fmt.Println(word+"---",len(word))
        count = count + len(word)
    }
    fmt.Println(count)
}

func convertToWords(number int) (string, int) {
    temp_number := 0
    ret_string := ""
    ret_number := 0
    if(number == 1000) {
        return "one"+number_words[1000],0
    } 
    if number >= 100 && number < 1000 {
        temp_number = number/100
        if val, ok := number_words[temp_number]; ok {
              ret_string = val+number_words[100]
              ret_number = number%100
              if ret_number != 0 {
                  ret_string = ret_string +"and"
              }
        }
    }
    if number >= 10 && number < 100 {
        temp_number = number/10
        if val, ok := number_words[number]; ok {
            ret_string = val
            ret_number = 0
        } else if val, ok := number_words[temp_number*10]; ok {
              ret_string = val
              ret_number = number%10
        }
    }
    if number < 10 {
        if val, ok := number_words[number]; ok {
           ret_string =  val
        }
    }
    return ret_string, ret_number
}
