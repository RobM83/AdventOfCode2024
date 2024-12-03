package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read input.txt
	reports := readInput("input.txt")

	safeReports := nrOfSafeReports(reports)

	fmt.Println("Safe reports:", safeReports)
	// for _, report := range reports {
	// 	fmt.Println(report)
	// }
}

func nrOfSafeReports(reports [][]int) int {
	var safeReports int
	for _, report := range reports {
		if isSafe(report) {
			//fmt.Println(report, "is safe")
			safeReports++
		}
	}

	return safeReports
}

func isSafe(report []int) bool {
	//Check the rest
	increasing := false
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			if i != 0 && !increasing {
				return false
			}
			increasing = true
		}
		if report[i] > report[i+1] {
			if i != 0 && increasing {
				return false
			}
			increasing = false
		}
		if report[i] == report[i+1] {
			return false
		}
		if delta(report[i], report[i+1]) > 3 {
			return false
		}
	}
	return true
}

func delta(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func readInput(fileName string) [][]int {
	//Read input.txt
	file, _ := os.Open(fileName)
	defer file.Close()

	//Split the input into two lists (every line has two numbers, entry list 1, entry list 2)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		numbersAsString := strings.Split(line, " ")
		numbers := []int{}
		for _, nr := range numbersAsString {
			numbers = append(numbers, StringToIntWOError(nr))
		}

		reports = append(reports, numbers)
	}

	return reports
}

func StringToIntWOError(str string) int {
	//Convert string to int
	strTrimmed := strings.Trim(str, " ")
	var num int
	num, _ = strconv.Atoi(strTrimmed)
	return num
}
