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

	valid, pos := validityCheck(report)
	fmt.Println(report)

	if !valid { //dampener
		newReport := make([]int, len(report))
		copy(newReport, report)
		newReport = append(newReport[:pos], newReport[pos+1:]...)
		valid, _ = validityCheck(newReport)
		fmt.Println(newReport, valid)
		if !valid && pos == 2 { //stupid fix
			newReport := make([]int, len(report))
			copy(newReport, report)
			newReport = append(newReport[:1], newReport[2:]...)
			valid, _ = validityCheck(newReport)
			fmt.Println(newReport, valid)
		}
		if !valid { //stupid fix
			newReport := make([]int, len(report))
			copy(newReport, report)
			newReport = append(newReport[:0], newReport[1:]...)
			valid, _ = validityCheck(newReport)
			fmt.Println(newReport, valid)
		}
	}

	return valid
}

// If false return the index of the first invalid number
func validityCheck(report []int) (bool, int) {
	increasing := false

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			if i != 0 && !increasing {
				return false, i + 1
			}
			increasing = true
		}
		if report[i] > report[i+1] {
			if i != 0 && increasing {
				return false, i + 1
			}
			increasing = false
		}
		if report[i] == report[i+1] {
			return false, i + 1
		}
		if delta(report[i], report[i+1]) > 3 {
			return false, i + 1
		}
	}
	return true, 0
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
