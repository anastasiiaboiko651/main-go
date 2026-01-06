package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var input string
	var length int

	for {
		fmt.Print("Введіть довжину пароля (4..128): ")
		fmt.Scan(&input)

		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Помилка: введіть число.")
			continue
		}
		if n < 4 || n > 128 {
			fmt.Println("Помилка: довжина має бути від 4 до 128.")
			continue
		}

		length = n
		break
	}

	var pool string

	for {
		pool = ""

		fmt.Print("Малі літери? (1/0): ")
		fmt.Scan(&input)
		if input == "1" {
			pool += "abcdefghijklmnopqrstuvwxyz"
		}

		fmt.Print("Великі літери? (1/0): ")
		fmt.Scan(&input)
		if input == "1" {
			pool += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}

		fmt.Print("Цифри? (1/0): ")
		fmt.Scan(&input)
		if input == "1" {
			pool += "0123456789"
		}

		fmt.Print("Спецсимволи? (1/0): ")
		fmt.Scan(&input)
		if input == "1" {
			pool += "!@#$%"
		}

		if pool == "" {
			fmt.Println("Помилка: оберіть хоча б одну категорію.")
			continue
		}
		break
	}

	for i := 0; i < length; i++ {
		fmt.Print(string(pool[rand.Intn(len(pool))]))
	}
	fmt.Println()
}
