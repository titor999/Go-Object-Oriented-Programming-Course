package main

import "fmt"

func main() {
	fmt.Println("Input your age >>>")
	var age int
	_, err := fmt.Scan(&age)
	if err != nil {
		return
	}
	fmt.Println(Name(age))
}

func Name(age int) string { // rename
	var name string
	if age == 0 {
		name = fmt.Sprintf("%v лет", age)
	} else if age == 1 {
		name = fmt.Sprintf("%v год", age)
	} else if age >= 2 && age <= 4 {
		name = fmt.Sprintf("%v года", age)
	} else if age >= 5 && age <= 20 {
		name = fmt.Sprintf("%v лет", age)
	} else if age%10 >= 1 && age%10 <= 4 {
		name = fmt.Sprintf("%v год", age)
	} else {
		name = fmt.Sprintf("%v лет", age)
	}
	return name
}
