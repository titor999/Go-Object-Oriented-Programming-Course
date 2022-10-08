package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Input your name >>>")
	var userName string

	_, err := fmt.Scan(&userName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Hello, %s!", userName)
}
