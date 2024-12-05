package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	updateOrder []pagePosition
	updates     []string
}

type pagePosition struct {
	before []int
	after  []int
}

func main() {
	//Read input.txt
	input := readInput("input.txt")
	lookupTable := createLookupTable(input)
	correctOrder := getCorrectUpdates(input, lookupTable)

	//fmt.Println(correctOrder)
	fmt.Println(sumOfTheMiddle(correctOrder))

}

func sumOfTheMiddle(update []string) int {
	//Get the sum of the middle elements
	sum := 0
	for _, u := range update {
		updateArray := strings.Split(u, ",")
		if len(updateArray)%2 == 0 {
			fmt.Println("The update list is not odd, @!$%")
		}
		mid := len(updateArray) / 2
		sum += StringToIntWOError(strings.Join(updateArray[mid:mid+1], ""))
	}

	return sum
}

func getCorrectUpdates(input input, lookupTable map[int]pagePosition) []string {
	//Get the correct order of updates
	var correctOrder []string
	for _, u := range input.updates {
		order := strings.Split(u, ",")
		valid := checkOrder(order, lookupTable)
		if valid {
			//correctOrder = append(correctOrder, u) //READ :P
		} else {
			fixedOrder := fixOrder(order, lookupTable)
			correctOrder = append(correctOrder, fixedOrder)
		}
	}

	return correctOrder
}

func createLookupTable(input input) map[int]pagePosition {
	lookupTable := make(map[int]pagePosition)
	for _, uo := range input.updateOrder {
		//Update after list
		entry := lookupTable[uo.before[0]]               //get the page
		entry.before = append(entry.before, uo.after[0]) //get the page the should come after it
		lookupTable[uo.before[0]] = entry

		//Update before list
		entry = lookupTable[uo.after[0]]
		entry.after = append(entry.after, uo.before[0])
		lookupTable[uo.after[0]] = entry
	}

	return lookupTable
}

func readInput(fileName string) input {
	//Read input.txt
	file, _ := os.Open(fileName)
	defer file.Close()

	//Split the input into two lists (every line has two numbers, entry list 1, entry list 2)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input input
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			pp := pagePosition{
				before: []int{StringToIntWOError(split[0])},
				after:  []int{StringToIntWOError(split[1])},
			}
			input.updateOrder = append(input.updateOrder, pp)
		}
		if strings.Contains(line, ",") {
			input.updates = append(input.updates, line)
		}
	}

	return input
}

func StringToIntWOError(str string) int {
	//Convert string to int
	strTrimmed := strings.Trim(str, " ")
	var num int
	num, _ = strconv.Atoi(strTrimmed)
	return num
}

func checkOrder(order []string, lookupTable map[int]pagePosition) bool {
	for i := 0; i < len(order); i++ {
		page := StringToIntWOError(order[i])
		entry := lookupTable[page]

		//Check if the page is in the correct order
		//page should always before the pages in the after list
		for _, a := range entry.after {
			for fwd := i + 1; fwd < len(order); fwd++ {
				if order[fwd] == strconv.Itoa(a) {
					return false
				}
			}
		}
		//page should always be after the pages in the before list
		for _, b := range entry.before {
			for bck := i - 1; bck >= 0; bck-- {
				if order[bck] == strconv.Itoa(b) {
					return false
				}
			}
		}
	}
	return true
}

func fixOrder(order []string, lookupTable map[int]pagePosition) string {
	//Fixx the order of the updates
	fmt.Println("Fixing order:", order)
	var correctOrder []string
	var checkagain bool
	for i := 0; i < len(order); i++ {
		if i == 0 {
			checkagain = false
		}
		page := StringToIntWOError(order[i])
		entry := lookupTable[page]

		//Check if the page is in the correct order
		//page should always before the pages in the after list
		for _, a := range entry.after {
			for fwd := i + 1; fwd < len(order); fwd++ {
				if order[fwd] == strconv.Itoa(a) {
					//Swap the two elements
					order[i], order[fwd] = order[fwd], order[i]
					fwd = i + 1
					checkagain = true
					//break
				}
			}
		}
		//page should always be after the pages in the before list
		for _, b := range entry.before {
			for bck := i - 1; bck >= 0; bck-- {
				if order[bck] == strconv.Itoa(b) {
					//Swap the two elements
					order[i], order[bck] = order[bck], order[i]
					bck = i - 1
					checkagain = true
					//break
				}
			}
		}

		if checkagain {
			i = -1
		}
	}

	correctOrder = append(correctOrder, strings.Join(order, ","))
	fmt.Println("Fixed order:", correctOrder)
	return strings.Join(correctOrder, "")
}
