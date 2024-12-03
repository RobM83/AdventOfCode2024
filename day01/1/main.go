package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type lists struct {
	list1     []int
	list2     []int
	distances []int
}

func main() {
	//Read input.txt
	lists := readInput("input.txt")

	// fmt.Println(lists.list1)
	// fmt.Println(lists.list2)

	lists.getDistances()
	// fmt.Println(lists.distances)

	fmt.Println(lists.totalDistance())
}

func (l *lists) totalDistance() int {
	//Return the total distance between the two lists
	var sum int
	for _, distance := range l.distances {
		sum += distance
	}
	return sum
}

func readInput(fileName string) lists {
	//Read input.txt
	file, _ := os.Open(fileName)
	defer file.Close()

	//Split the input into two lists (every line has two numbers, entry list 1, entry list 2)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var list1 []int
	var list2 []int
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")
		list1 = append(list1, StringToIntWOError(lineSplit[0]))
		list2 = append(list2, StringToIntWOError(lineSplit[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return lists{list1, list2, nil}
}

func (l *lists) getDistances() {
	//Get the distances between the two lists

	for i := 0; i < len(l.list1); i++ {
		delta := 0
		num1 := l.list1[i]
		num2 := l.list2[i]

		if num1 > num2 {
			delta = num1 - num2
		} else {
			delta = num2 - num1
		}

		l.distances = append(l.distances, delta)
	}
}

func StringToIntWOError(str string) int {
	//Convert string to int
	strTrimmed := strings.Trim(str, " ")
	var num int
	num, _ = strconv.Atoi(strTrimmed)
	return num
}
