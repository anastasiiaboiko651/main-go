package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LoginAttempt struct {
	id int

	passwordGuess string

	success bool

	blocked bool

	timestamp time.Time
}

type LoginStats struct {
	total int

	successful int

	failed int

	blocked int
}

func makeGuess(i int, correct string) string {

	return "wrong_" + fmt.Sprint(i)

}

func generateAttempts(correct string, maxAttempts int) []LoginAttempt {

	attempts := make([]LoginAttempt, 0, maxAttempts)

	for i := 0; i < maxAttempts; i++ {

		guess := makeGuess(i, correct)

		attempt := LoginAttempt{

			id: i + 1,

			passwordGuess: guess,

			success: guess == correct,

			blocked: false,

			timestamp: time.Now(),
		}

		attempts = append(attempts, attempt)

	}

	return attempts

}

func simulateLogin(attempts []LoginAttempt, maxFailedInRow int) LoginStats {

	stats := LoginStats{}

	failedInRow := 0

	accountBlocked := false

	for i := 0; i < len(attempts); i++ {

		if accountBlocked {

			stats.blocked++
			continue

		}

		if attempts[i].success {

			failedInRow = 0

			stats.successful++

		} else {

			failedInRow++

			stats.failed++

		}

		if failedInRow >= maxFailedInRow {

			accountBlocked = true

		}

		stats.total++

	}

	return stats

}

func printSummary(stats LoginStats) {

	fmt.Println("=== Simulation summary ===")

	fmt.Println("Total attempts:", stats.total)

	fmt.Println("Successful logins:", stats.successful)

	fmt.Println("Failed attempts:", stats.failed)

	fmt.Println("Blocked times:", stats.blocked)

	if stats.successful > 0 && stats.blocked == 0 {

		fmt.Println("Result: account probably NOT under attack.")

	} else if stats.blocked > 0 {

		fmt.Println("Result: possible brute-force attack, account should be blocked.")

	} else {

		fmt.Println("Result: need more analysis.")

	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	correctPassword := "qwerty123"

	maxAttempts := 5

	maxFailedInRow := 3

	fmt.Println("Simulating login attempts to protected account...")

	attempts := generateAttempts(correctPassword, maxAttempts)

	stats := simulateLogin(attempts, maxFailedInRow)

	printSummary(stats)

}
