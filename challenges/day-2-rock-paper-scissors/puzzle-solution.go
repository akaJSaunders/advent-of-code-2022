package main

import (
    "bufio"
    "fmt"
    "os"
		"strings"
)

var part1Move = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissor",
	"A": "rock",
	"B": "paper",
	"C": "scissor",
}

var part2Move = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissor",
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

var win = map[string]string{
	"paper": "rock",
	"scissor": "paper",
	"rock": "scissor",
}

var lose = map[string]string{
	"paper": "scissor",
	"scissor": "rock",
	"rock": "paper",
}

var points = map[string]int{
	"rock": 1,
	"paper": 2,
	"scissor": 3,
	"lose": 0,
	"draw": 3,
	"win": 6,
}

func part1RoundScore(player1 string, player2 string) (int, int) {
	player1CurrentScore := points[player1]
	player2CurrentScore := points[player2]

	if player1 == player2 {
		player1CurrentScore += points["draw"]
		player2CurrentScore += points["draw"]
	} else if win[player1] == player2 {
		player1CurrentScore += points["win"]
		player2CurrentScore += points["lose"]
	} else {
		player1CurrentScore += points["lose"]
		player2CurrentScore += points["win"]
	}

	return player1CurrentScore, player2CurrentScore
}

func main() {
	readFile, err := os.Open("challenges/day-2-rock-paper-scissors/puzzle-input.txt")
  if err != nil {
    fmt.Println(err)
  }

  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)

	totalPart1 := 0
	totalPart2 := 0

  for fileScanner.Scan() {

		players := strings.Split(fileScanner.Text(), " ")

		// Part 1
		_, playerScorePart1 := part1RoundScore(
			part1Move[players[0]],
			part1Move[players[1]],
		)

		totalPart1 += playerScorePart1

		// Part 2 
		playerScorePart2 := 0

		player1Move := part2Move[players[0]]
		player2Move := part2Move[players[1]]

		if player2Move == "lose" {
			_, playerScorePart2 = part1RoundScore(
				player1Move,
				win[player1Move],
			)
		} else if player2Move == "win" {
			_, playerScorePart2 = part1RoundScore(
				player1Move,
				lose[player1Move],
			)
		} else {
			_, playerScorePart2 = part1RoundScore(
				player1Move,
				player1Move,
			)
		}

		totalPart2 += playerScorePart2
  }

	fmt.Println("Part 1 - My total score:", totalPart1)
	fmt.Println("Part 2 - My total score:", totalPart2)

  readFile.Close()
}
