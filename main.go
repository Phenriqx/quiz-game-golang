package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (g *GameState) RunGame() {
	fmt.Println("You have 18 questions to answer. You have 1:30 minutes to complete the quiz.")
	done := make(chan bool)

	// Starts a 90-second countdown in a separate Goroutine.
	go func() {
		countdown := 15
		for countdown > 0 {
			time.Sleep(1 * time.Second)
			countdown--
			fmt.Printf("\033[31m%d seconds remaining\033[31m\n", countdown)
		}
		done <- true
	} ()

	// Print the questions and alternatives
	for i, question := range g.Questions {
		select {
		case <-done:
			fmt.Println("\n⏳ Time's up! Quiz over.")
			return
		default:
		}
		
		fmt.Printf("\033[33m %d - %s \033[33m\n", i+1, question.Text)

		for j, alternative := range question.Alternatives {
			fmt.Printf("\033[36m %d. %s \033[36m\n", j+1, alternative)
		}

		fmt.Println("Enter your answer: ")

		var answer int
		var err error

		for {
			select {
			case <-done: // If time runs out, stop the quiz
				fmt.Println("\n⏳ Time's up! Quiz over.")
				return
			default:
				// Reading user input
				reader := bufio.NewReader(os.Stdin)
				read, _ := reader.ReadString('\n')

				answer, err = toInt(read[:len(read)-1])
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			break
		}
		
		select {
			case <-done:
                fmt.Println("\n��� Time's up! Quiz over.")
                return
            default:
                if question.Answer == answer {
					fmt.Println("Correct!")
					g.Points += 10
					fmt.Printf("Your current score is: %d\n", g.Points)
				} else {
					fmt.Println("Wrong!")
					fmt.Printf("Your current score is: %d\n", g.Points)
					fmt.Println("------------------------------")
				}
		}
	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
	game.RunGame()

	fmt.Printf("Your final score is: %d points!\n", game.Points)
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid input. Please enter a number")
	}
	return i, nil
}