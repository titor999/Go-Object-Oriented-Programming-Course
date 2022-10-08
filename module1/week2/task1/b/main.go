package main

import "fmt"

func main() {
	fmt.Println("Input an year >>>")
	var year int
	_, err1 := fmt.Scan(&year)
	if err1 != nil {
		return
	}
	fmt.Println("Input an month")
	var month string
	_, err2 := fmt.Scan(&month)
	if err2 != nil {
		return
	}

	switch {
	case month == "April" || month == "June" || month == "September" || month == "November":
		fmt.Printf("%s is 30", month)
	case month == "February":
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			fmt.Printf("%s is 29", month)
		} else {
			fmt.Printf("%s is 28", month)
		}
	default:
		fmt.Printf("%s is 31", month)
	}
}
