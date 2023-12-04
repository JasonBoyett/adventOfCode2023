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

type Gear struct {
	row         int
	col         int
	ratio       int
	adjacentOne int
	adjacentTwo int
}

func main() {
	var ptOne uint64
	var ptTwo uint64

	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	schematic := buildScematic(lines)
	gears := generateGears(schematic)
	for _, num := range schematic.validNums {
		ptOne += uint64(num)
	}
	for _, gear := range gears {
		ptTwo += uint64(gear.ratio)
	}
	fmt.Println("The answer for part one is: ", ptOne)
	fmt.Println("The answer for part two is: ", ptTwo)

}

func generateGears(s Schematic) []Gear {
	var gears []Gear
	gearSymbol := byte('*')
	for i := 0; i < len(s.schematic); i++ {
		for j := 0; j < len(s.schematic[i]); j++ {
			if s.schematic[i][j] == gearSymbol {
				gears = append(gears, Gear{
					row:         i,
					col:         j,
					ratio:       0,
					adjacentTwo: 0,
					adjacentOne: 0,
				})
			}
		}
	}

	for i, gear := range gears {
		gears[i] = findAdjacentNumbers(gear, s)
	}

	return gears
}

func findAdjacentNumbers(g Gear, s Schematic) Gear {
	var nums []int
	nums = append(nums, adjcentLeft(g, s))
	nums = append(nums, adjcentRight(g, s))
	nums = append(nums, adjcentUp(g, s))
	nums = append(nums, adjcentDown(g, s))
	nums = append(nums, adjcentUpLeft(g, s))
	nums = append(nums, adjcentUpRight(g, s))
	nums = append(nums, adjcentDownLeft(g, s))
	nums = append(nums, adjcentDownRight(g, s))
	nums = removeDuplicates(nums)
	if len(nums) == 2 {
		g.adjacentOne = nums[0]
		g.adjacentTwo = nums[1]
    g.ratio = nums[0] * nums[1]
		return g
	}
	return g
}

func adjcentLeft(g Gear, s Schematic) int {
	if g.col == 0 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row][g.col-1])) {
		return findFullNumber(s.schematic[g.row], g.col-1)
	}
	return 0
}

func adjcentRight(g Gear, s Schematic) int {
	if g.col == len(s.schematic[g.row])-1 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row][g.col+1])) {
		return findFullNumber(s.schematic[g.row], g.col+1)
	}
	return 0
}

func adjcentUp(g Gear, s Schematic) int {
	if g.row == 0 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row-1][g.col])) {
		return findFullNumber(s.schematic[g.row-1], g.col)
	}
	return 0
}

func adjcentDown(g Gear, s Schematic) int {
	if g.row == len(s.schematic)-1 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row+1][g.col])) {
		return findFullNumber(s.schematic[g.row+1], g.col)
	}
	return 0
}

func adjcentUpLeft(g Gear, s Schematic) int {
	if g.row == 0 || g.col == 0 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row-1][g.col-1])) {
		return findFullNumber(s.schematic[g.row-1], g.col-1)
	}
	return 0
}

func adjcentUpRight(g Gear, s Schematic) int {
	if g.row == 0 || g.col == len(s.schematic[g.row])-1 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row-1][g.col+1])) {
		return findFullNumber(s.schematic[g.row-1], g.col+1)
	}
	return 0
}

func adjcentDownLeft(g Gear, s Schematic) int {
	if g.row == len(s.schematic)-1 || g.col == 0 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row+1][g.col-1])) {
		return findFullNumber(s.schematic[g.row+1], g.col-1)
	}
	return 0
}

func adjcentDownRight(g Gear, s Schematic) int {
	if g.row == len(s.schematic)-1 || g.col == len(s.schematic[g.row])-1 {
		return 0
	}
	if unicode.IsDigit(rune(s.schematic[g.row+1][g.col+1])) {
		return findFullNumber(s.schematic[g.row+1], g.col+1)
	}
	return 0
}

func findFullNumber(line string, point int) int {
	var result string
	var start int
	for i := point; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			start = i
		} else {
			break
		}
	}
	for i := start; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			result = result + string(line[i])
		} else {
			break
		}
	}
  number, err := strconv.Atoi(result)
  if err != nil {
    panic(err)
  }
  return number


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
	return len(line)
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
	if i < 0 || i > len(schematic)-1 {
		return false
	} else if j < 0 || j > len(schematic[i])-1 {
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

func removeDuplicates(input []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for _, value := range input {
		if encountered[value] == false {
			encountered[value] = true
			if value != 0 {
				result = append(result, value)
			}
		}
	}

	return result
}
