package main

import (
	"fmt"
)

func main() {
	var (
		num1, num2     float32
		operator, stop string
	)

	for {
		fmt.Println("введите первое число")
		fmt.Scanln(&num1)
		fmt.Println("введите второе число")
		fmt.Scanln(&num2)
		fmt.Println("выберите действие +,-,*,/")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Printf("ваш ответ: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("ваш ответ: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("ваш ответ: %.2f\n", num1*num2)
		case "/":
			fmt.Printf("ваш ответ: %.2f\n", num1/num2)
		default:
			fmt.Println("неизвестный оператор")
		}

		fmt.Println("завершить программу? (Y/N)")
		fmt.Scanln(&stop)

		if stop == "Y" || stop == "y" {
			break
		}
	}
}