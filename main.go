package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	for true {
		fmt.Printf("Shake the Magic 8-Ball? (press 'Y' or 'N'): ")

		var userResp string
		fmt.Scan(&userResp)

		if strings.ToUpper(userResp) == "Y" {
			fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
			continue
		} else if strings.ToUpper(userResp) == "N" {
			fmt.Println("Magic 8-Ball says: Thank you for playing")
			break
		} else {
			fmt.Println("Please use a valid response ('Y' or 'N': ")
			continue
		}
	}
}
