package main

import ("fmt" ; "time")

func main() {
	todaysTime := time.Now()
	theHour := todaysTime.Hour()
	fmt.Println("hello" , theHour)
	
	switch {
	case theHour <12:
		fmt.Println("It is " , theHour,"oclock. The MORNING")
	case theHour >12:
		fmt.Println("It is " , theHour,"oclock. The Evening")
	}

}