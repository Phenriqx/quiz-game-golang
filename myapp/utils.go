package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type GameState struct {
	Name      string
	Points    string
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Welcome to the game!")
	fmt.Println("Please enter your name:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
        return
	} else if name == "" || len(name) <= 1 {
		fmt.Println("Please enter a valid name.")
		return
	}

	g.Name = name
	fmt.Printf("Let's play %s", g.Name)
	g.Points = "0"
}

func (g *GameState) ProcessCSV () {
	file, err := os.Open("quiz-go.csv")
	if err != nil {
        fmt.Println("Error opening the CSV file:", err)
        return
    }

	defer file.Close() // Close the file when we're done. This will be executed at the end of the enclosing function (main)

	reader := csv.NewReader(file) 
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the CSV file:", err)
		return
	}

	for index, record := range records {
		if index > 0 {
			question := Question {
				Text: record[0],
				Alternatives: record[1:5],
				Answer: record[5],
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

type Question struct {
	Text         string
	Alternatives []string
	Answer       string
}
