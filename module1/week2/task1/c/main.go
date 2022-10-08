package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Хотите ли Вы сегодня поиграть в бадминтон?")
	fmt.Println("Какой сегодня день недели?")

	var dayWeek string
	_, err := fmt.Scan(&dayWeek)
	if err != nil {
		log.Fatalln(err)
	}
	if dayWeek == "Воскресенье" || dayWeek == "воскресенье" {
		fmt.Println("Какая сегодня погода на улице? (жарко/тепло/холодно)")
		var weather string
		_, err1 := fmt.Scan(&weather)
		if err1 != nil {
			log.Fatalln(err1)
		}
		if weather == "тепло" {
			fmt.Println("Есть ли сегодня осадки? (ясно, облачно, дождь, снег, град)")
			var precipitation string
			_, err2 := fmt.Scan(&precipitation)
			if err2 != nil {
				log.Fatalln(err2)
			}
			if precipitation == "ясно" || precipitation == "облачно" {
				fmt.Println("Есть ли ветер на улице? (есть/нет)")
				var wind string
				_, err3 := fmt.Scan(&wind)
				if err3 != nil {
					log.Fatalln(err3)
				}
				if wind == "нет" {
					fmt.Println("Какая сегодня влажность? (высокая/низкая)")
					var humidity string
					_, err4 := fmt.Scan(&humidity)
					if err4 != nil {
						log.Fatalln(err4)
					}
					if humidity == "низкая" {
						fmt.Println("Да")
						os.Exit(0)
					}
				}
			}
		}
	}
	fmt.Println("нет")
}
