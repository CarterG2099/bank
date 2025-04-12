package main

import (
	"fmt"
	"sort"

	"bank/bank"
	"bank/example"
	"bank/carter"
)

func main() {
	// PlayGame()

	// Run 100 games and summarize results
	RunSimulations(100)
}

func PlayGame() {
	game := bank.NewGame(
		&example.ExamplePlayer{RoundLimit: 20},
		&example.ExamplePlayer{RoundLimit: 100},
		&example.ExamplePlayer{RoundLimit: 200},
		&example.AnotherExamplePlayer{},
		&example.BankAfter{BankAfter: 5},
		&example.BankAfter{BankAfter: 6},
		&carter.Carter{},
	)
	game.Play()
}

func RunSimulations(numGames int) {
	winCount := make(map[string]int)
	highScores := make(map[string]int)

	for i := 0; i < numGames; i++ {
		game := bank.NewGame(
			&example.ExamplePlayer{RoundLimit: 20},
			&example.ExamplePlayer{RoundLimit: 100},
			&example.ExamplePlayer{RoundLimit: 200},
			&example.AnotherExamplePlayer{},
			&example.BankAfter{BankAfter: 5},
			&example.BankAfter{BankAfter: 6},
			&carter.Carter{},
		)
		game.Play()

		highest := -1
		var winner string

		for _, p := range game.Players {
			name := p.Name
			if p.Score > highScores[name] {
				highScores[name] = p.Score
			}
			if p.Score > highest {
				highest = p.Score
				winner = name
			}
		}

		winCount[winner]++
	}

	// Display results
	fmt.Printf("\nAfter %d games:\n", numGames)
	type stats struct {
		name      string
		wins      int
		winPct    float64
		highScore int
	}
	var resultList []stats
	for name, wins := range winCount {
		resultList = append(resultList, stats{
			name:      name,
			wins:      wins,
			winPct:    float64(wins) / float64(numGames) * 100,
			highScore: highScores[name],
		})
	}

	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].wins > resultList[j].wins
	})

	for _, r := range resultList {
		fmt.Printf("%-25s | Wins: %3d (%.1f%%) | Highest Score: %d\n", r.name, r.wins, r.winPct, r.highScore)
	}
}

