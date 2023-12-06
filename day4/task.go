package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Card map[string][]int

func main() {
	ptOne := 0
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	scores := parseCards(input)

	for _, score := range scores {
		ptOne += score
	}

	fmt.Println("The answer for part one is: ", ptOne)
}

func parseCards(lines []string) []int {
	var scores []int
	var cards []Card

	for _, line := range lines {
		cards = append(cards, format(line))
	}

  for _, card := range cards {
    scores = append(scores, calculateScore(card))
  }

	return scores
}

func format(line string) Card {
	var temp [][]string
	result := make(Card, 2)
	for i := 0; i < len(line); i++ {
		if line[i] == ':' {
			line = line[i+1:]
      break
		}
	}

	line = strings.TrimSpace(line)
	sides := strings.Split(line, "|")
	temp = append(temp, strings.Split(sides[0], " "))
	temp = append(temp, strings.Split(sides[1], " "))
	temp[0] = strip(temp[0], "")
	temp[1] = strip(temp[1], "")

	for i := 0; i < len(temp[0]); i++ {
    fmt.Println(temp[0][i])
		temp[0][i] = strings.TrimSpace(temp[0][i])
		holder, _ := strconv.Atoi(temp[0][i])
		result["given"] = append(result["given"], holder)
	}

	for i := 0; i < len(temp[1]); i++ {
		temp[1][i] = strings.TrimSpace(temp[1][i])
		holder, _ := strconv.Atoi(temp[1][i])
		result["winning"] = append(result["winning"], holder)
	}
	return result
}

func calculateScore(card Card) int {
	score := 0
	for _, winning := range card["winning"] {
		if slices.Contains(card["given"], winning) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
  fmt.Println(score)
	return score
}

func strip[T comparable](s []T, i T) []T {
	var result []T
	for _, v := range s {
		if v != i {
			result = append(result, v)
		}
	}
	return result
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
