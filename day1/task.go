package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readInput("input2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numbers := processInput(lines)
	result := sum(numbers)
	fmt.Println(numbers)
	fmt.Println(result)

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

func processInput(lines []string) []int {
	var result []int
	for i := 0; i < len(lines); i++ {
		result = append(result, processLine(lines[i]))
	}
	return result
}

func processLine(line string) int {
	first := findFirstDigie(line)
	second := findSecondDigie(line)
	result := ""

	if second == "" {
		result = first + first
	} else {
		result = first + second
	}

	number, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return number
}

func findFirstDigie(line string) string {
	for i := 0; i < len(line); i++ {
		possibleDigit, err := parseDigitName(string(line[:i]))
		if err == nil {
			return possibleDigit
		}
		_, err = strconv.Atoi(string(line[i]))
		if err == nil {
			return string(line[i])
		}
	}
	return ""
}

func findSecondDigie(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		possibleDigit, err := parseDigitName(string(line[i:]))
		if err == nil {
			return possibleDigit
		}
		_, err = strconv.Atoi(string(line[i]))
		if err == nil {
			return string(line[i])
		}
	}
	return ""
}

func sum(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}

func parseDigitName(input string) (string, error) {
	fmt.Println(input)
	if strings.Contains(input, "one") {
		return "1", nil
	} else if strings.Contains(input,"two") {
		return "2", nil
	} else if strings.Contains(input,"three") {
		return "3", nil
	} else if strings.Contains(input,"four") {
		return "4", nil
	} else if strings.Contains(input,"five") {
		return "5", nil
	} else if strings.Contains(input,"six") {
		return "6", nil
	} else if strings.Contains(input,"seven") {
		return "7", nil
	} else if strings.Contains(input,"eight") {
		return "8", nil
	} else if strings.Contains(input,"nine") {
		return "9", nil
	} else if strings.Contains(input,"zero") {
		return "0", nil
	} else {
		return "", fmt.Errorf("Invalid input")
	}
}
