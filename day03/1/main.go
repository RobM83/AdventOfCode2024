package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type calc struct {
	action string
	value1 int
	value2 int
}

func main() {
	//Read input.txt
	lines := readInput("input.txt")

	calcLines := getMultiplies(lines)
	total := getTotalCalc(calcLines)

	fmt.Println("Total: ", total)
}

func readInput(fileName string) []string {
	//Read input.txt
	file, _ := os.Open(fileName)
	defer file.Close()

	//Split the input into two lists (every line has two numbers, entry list 1, entry list 2)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func getTotalCalc(calcLines []calc) int {
	total := 0
	for _, c := range calcLines {
		switch c.action {
		case "mul":
			total = total + c.value1*c.value2
		}
	}
	return total
}

func getMultiplies(line []string) []calc {
	calcList := []calc{}

	for _, l := range line {
		pattern := `mul\(\d+,\d+\)`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(l, -1)

		for _, match := range matches {
			fmt.Println(match)
			m := strings.Replace(match, "mul(", "", -1)
			m = strings.Replace(m, ")", "", -1)
			nr := strings.Split(m, ",")

			c := calc{
				action: "mul",
				value1: StringToIntWOError(nr[0]),
				value2: StringToIntWOError(nr[1]),
			}
			calcList = append(calcList, c)
		}
	}

	return calcList
}

func StringToIntWOError(str string) int {
	//Convert string to int
	strTrimmed := strings.Trim(str, " ")
	var num int
	num, _ = strconv.Atoi(strTrimmed)
	return num
}
