package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Red int = iota
	Green
	Blue
)

type Game struct {
	sets      []Set
	id        int
	maxGreens int
	maxReds   int
	maxBlues  int
}

type Set struct {
	red   int
	green int
	blue  int
}

func (g *Game) printGame() {
	fmt.Println("Game id:", g.id)
	fmt.Println("Max greens:", g.maxGreens)
	fmt.Println("Max reds:", g.maxReds)
	fmt.Println("Max blues:", g.maxBlues)
	fmt.Println("-------------")
}

func (g *Game) getPower() int {
  return g.maxGreens * g.maxReds * g.maxBlues
}

func main() {
	var games []Game
	var possibleGames []Game
	ptOne := 0
  ptTwo := 0
	lines, err := readInput("input1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	for line := range lines {
		games = append(games, parseGame(lines[line]))
	}

	for game := range games {
		if isPossible(games[game], 12, 13, 14) {
			possibleGames = append(possibleGames, games[game])
		}
	}

  for _, g := range games {
    ptTwo += g.getPower()
  }

	for game := range possibleGames {
		ptOne += possibleGames[game].id
	}

  fmt.Println("The answer for part one is: ", ptOne)
  fmt.Println("The answer for part two is: ", ptTwo)

}

func parseGame(gameString string) Game {
	var sets []Set
	gameId, mark := identifyGame(gameString)
	gameString = gameString[mark:]
	setStrings := strings.Split(gameString, ";")
	for set := range setStrings {
		sets = append(sets, generateSet(setStrings[set]))
	}
	game := generateGame(sets, gameId)
	return game
}

func identifyGame(gameString string) (int, int) {
	var id, mark int
  var err error
	for i := range gameString {
		if gameString[i] == ':' {
			mark = i + 1
      gameString = gameString[:i]
			gameString = strings.Trim(gameString, "Game")
			gameString = strings.Trim(gameString, " ")
			gameString = strings.Trim(gameString, ":")
      id, err = strconv.Atoi(gameString)
      if err != nil {
        fmt.Println("Error converting string to int")
        break
      }
			break
		}
	}
	return id, mark
}

func generateSet(setString string) Set {
	colorStrings := strings.Split(setString, ",")
	set := Set{
		red:   0,
		green: 0,
		blue:  0,
	}

	for i := range colorStrings {
		color, _ := matchColor(colorStrings[i])
		switch color {
		case Red:
			{
				colorStrings[i] = strings.Trim(colorStrings[i], "red")
				colorStrings[i] = strings.Trim(colorStrings[i], " ")
        colorStrings[i] = strings.Trim(colorStrings[i], ",")
				val, _ := strconv.Atoi(colorStrings[i])
				set.red += val
			}
		case Green:
			{
				colorStrings[i] = strings.Trim(colorStrings[i], "green")
				colorStrings[i] = strings.Trim(colorStrings[i], " ")
        colorStrings[i] = strings.Trim(colorStrings[i], ",")

				val, _ := strconv.Atoi(colorStrings[i])
				set.green += val
			}
		case Blue:
			{
				colorStrings[i] = strings.Trim(colorStrings[i], "blue")
				colorStrings[i] = strings.Trim(colorStrings[i], " ")
        colorStrings[i] = strings.Trim(colorStrings[i], ",")
				val, _ := strconv.Atoi(colorStrings[i])
				set.blue += val
			}
		}
	}
	return set

}

func generateGame(sets []Set, id int) Game {
	game := Game{
		sets:      sets,
		id:        id,
		maxGreens: 0,
		maxReds:   0,
		maxBlues:  0,
	}

	for set := range sets {
		if sets[set].green > game.maxGreens {
			game.maxGreens = sets[set].green
		}
		if sets[set].red > game.maxReds {
			game.maxReds = sets[set].red
		}
		if sets[set].blue > game.maxBlues {
			game.maxBlues = sets[set].blue
		}
	}

	return game
}

func matchColor(match string) (int, error) {
	match = strings.ToLower(match)
	if strings.Contains(match, "green") {
		return Green, nil
	}
	if strings.Contains(match, "red") {
		return Red, nil
	}
	if strings.Contains(match, "blue") {
		return Blue, nil
	}
	return -1, errors.New("No color match")
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func isPossible(game Game, red, green, blue int) bool {
  if game.maxReds > red {
    return false
  }
  if game.maxGreens > green {
    return false
  }
  if game.maxBlues > blue {
    return false
  }
  return true
}

