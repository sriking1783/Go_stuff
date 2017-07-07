package main
import "fmt"

func main() {
      day_count := make(map[int]int)
      days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
      months := [12]string{"Jan", "Feb", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
      day_count[0] = 31 //Jan
      day_count[1] = 28 //Feb
      day_count[2] = 31 //Mar
      day_count[3] = 30 //Apr
      day_count[4] = 31 //May
      day_count[5] = 30 //Jun
      day_count[6] = 31 //Jul
      day_count[7] = 31 //Aug
      day_count[8] = 30 //Sep
      day_count[9] = 31 //Oct
      day_count[10] = 30 //Nov
      day_count[11] = 31 //Dec
      jan_1st := 1
      current_sunday := (7-jan_1st)+1
      count := 0
      //temp_count := 0
      num_days := 365
      for current_sunday <= 31 {
            if current_sunday+7 <= 31 {
                  current_sunday = current_sunday + 7
                  //count += 1
            } else {
                  break
            }
      }
      fmt.Println("Days ",jan_1st," ",num_days, " Current Sunday ",current_sunday," Count ",count)
      /*for k := 1901; k <= 2000; k++ {
            prev := k-1
            if prev%400 == 0 {
                  num_days = 366
            } else if prev%4 == 0 && prev%400 != 0 && prev%100 != 0 {
                  num_days = 366
            } else {
                  num_days = 365
            }
            jan_1st += (num_days)%7
            for jan_1st > 7 {
                jan_1st = jan_1st%7
            }
            current_sunday = (7-jan_1st)+1 
            //count += 1
            temp_count = 0
            fmt.Print("Jan first sunday: ", current_sunday," for year ",k," ")
            //fmt.Println("Days ",jan_1st," ",num_days, " Current Sunday ",current_sunday," ",k)
            for current_sunday <= 31 {
                  current_sunday = current_sunday + 7
                  
                  count += 1
                  temp_count += 1
            }
            fmt.Println("Num of sundays of year ",k," is ",temp_count," Count: ",count)
      }*/
      for k := 1901; k <= 2000; k++ {
            prev := k-1
            if prev%400 == 0 {
                  num_days = 366
            } else if prev%4 == 0 && prev%400 != 0 && prev%100 != 0 {
                  num_days = 366
            } else {
                  num_days = 365
            }

            if k%400 == 0 {
                  day_count[1] = 29
            } else if k%4 == 0 && k%400 != 0 && k%100 != 0 {
                  day_count[1] = 29
            } else {
                  day_count[1] = 28
            }
          
            jan_1st += (num_days)%7
            for jan_1st > 7 {
                jan_1st = jan_1st%7
            }
            current_sunday = (7-jan_1st)+1 
            if current_sunday == 1 {
                fmt.Println("Jan 1st is on", days[jan_1st%7]," First Sunday ",current_sunday," For year ",k)
                count += 1
            }
            prev_month_days := 31
            prev_month_first := jan_1st
            for i := 0; i < 12; i++ {
                month := i
                number_of_days := day_count[month]
                if month == 0 {
                    continue
                }
                month_first := prev_month_first + (prev_month_days)%7
                current_sunday = (7-month_first%7)+1 
                if current_sunday > 7 {
                    current_sunday = current_sunday-7
                }
                fmt.Println(months[month],"(",month,") 1st is on", days[month_first%7]," First Sunday ",current_sunday," For year ",k)
                if current_sunday == 1 {
                      count += 1
                }
                
                prev_month_days = number_of_days
                prev_month_first = month_first
            }

      }
      fmt.Println(count)
}

/*func get_sunday(first_date int) {
      
}*/
