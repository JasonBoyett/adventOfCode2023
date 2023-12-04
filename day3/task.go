package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type SchematicNumber struct {
	number int
	start  int
	end    int
	row    int
}

type Schematic struct {
	schematic    []string
	scematicNums []SchematicNumber
	validNums    []int
	invalidNums  []int
}

func main() {
	var ptOne uint64 

	lines, err := readInput("example1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	schematic := buildScematic(lines)
	schematic.printSchematic()
	for _, num := range schematic.validNums {
		ptOne += uint64(num)
	}
	fmt.Println("The answer for part one is: ", ptOne)

}

func (s *Schematic) printSchematic() {
	fmt.Println("Valid numbers:", s.validNums)
	fmt.Println("-------------")
	fmt.Println("Invalid numbers:", s.invalidNums)
	fmt.Println("-------------")
}

func buildScematic(scematic []string) Schematic {
	s := Schematic{schematic: scematic}
	s.scematicNums = findSchematicNumbers(s)
	for _, num := range s.scematicNums {
		if validate(num, scematic) {
			s.validNums = append(s.validNums, num.number)
		} else {
			s.invalidNums = append(s.invalidNums, num.number)
		}
	}
	return s
}

func findSchematicNumbers(s Schematic) []SchematicNumber {
	var result []SchematicNumber
	for i := 0; i < len(s.schematic); i++ {
		for j := 0; j < len(s.schematic[i]); j++ {
			if unicode.IsDigit(rune(s.schematic[i][j])) {
				end := findEndOfNumber(s.schematic[i], j)
				value, _ := strconv.Atoi(s.schematic[i][j:end])
				result = append(result,
					SchematicNumber{
						number: value,
						start:  j,
						end:    end,
						row:    i,
					})
				j = end
			}
		}
	}
	return result
}

func findEndOfNumber(line string, start int) int {
	for i := start; i < len(line); i++ {
		if !unicode.IsDigit(rune(line[i])) {
			return i
		}
	}
	return len(line) - 1
}

func validate(sNum SchematicNumber, schematic []string) bool {
	for i := sNum.start; i < sNum.end; i++ {
		if stepAround(schematic, sNum.row, i) {
			return true
		}
	}
  return false
}

func stepAround(schematic []string, i, j int) bool {
	if stepUp(schematic, i, j) {
		return true
	}
	if stepDown(schematic, i, j) {
		return true
	}
	if stepLeft(schematic, i, j) {
		return true
	}
	if stepRight(schematic, i, j) {
		return true
	}
  if stepUpLeft(schematic, i, j) {
    return true
  }
  if stepUpRight(schematic, i, j) {
    return true
  }
  if stepDownLeft(schematic, i, j) {
    return true
  }
  if stepDownRight(schematic, i, j) {
    return true
  }

	return false
}

func stepUp(schematic []string, i, j int) bool {
	return step(schematic, i-1, j) 
}

func stepDown(schematic []string, i, j int) bool {
	return step(schematic, i+1, j) 
}

func stepLeft(schematic []string, i, j int) bool {
	return step(schematic, i, j-1) 
}

func stepRight(schematic []string, i, j int) bool {
	return step(schematic, i, j+1) 
}

func stepUpLeft(schematic []string, i, j int) bool {
  return step(schematic, i-1, j-1) 
}

func stepUpRight(schematic []string, i, j int) bool {
  return step(schematic, i-1, j+1) 
}

func stepDownLeft(schematic []string, i, j int) bool {
  return step(schematic, i+1, j-1) 
}

func stepDownRight(schematic []string, i, j int) bool {
  return step(schematic, i+1, j+1) 
}

func step(schematic []string, i, j int) bool {
	if i < 0 || i > len(schematic) - 1 {
		return false
	} else if j < 0 || j > len(schematic[i]) - 1 {
		return false
	} else if !isSimbol(schematic[i][j]) {
		return false
	}
	return true
}

func isSimbol(char byte) bool {
	if unicode.IsDigit(rune(char)) {
		return false
	}
	if char == '.' {
		return false
	}
	return true
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
