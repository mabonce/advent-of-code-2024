package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	filename := "input.txt"

	// lookup file, open file, read contents and assign into 2 slices
	content, err := os.Open(filename)
	if err != nil {
		fmt.Println("There was an issue opening the file:", err)
	}
	defer content.Close()

	// create the slices/arrays for use
	var leftList []int
	var rightList []int

	// now scan the content to back into the slices/arrays
	contentScanner := bufio.NewScanner(content)
	for contentScanner.Scan() {
		line := strings.TrimSpace(contentScanner.Text())
		lineParts := strings.Fields(line)

		firstNum, err := strconv.Atoi(lineParts[0])
		if err != nil {
			fmt.Println("Failed converting first number to an integer:", err)
		}

		secondNum, err := strconv.Atoi(lineParts[1])
		if err != nil {
			fmt.Println("Failed converting second number to an integer:", err)
		}

		// now assign each of the numbers to the respective slice/array
		leftList = append(leftList, firstNum)
		rightList = append(rightList, secondNum)
	}

	// lists should be the same size or we made a mistake
	if len(leftList) != len(rightList) {
		fmt.Println("Lists are not the same length:", len(leftList), len(rightList))
	}

	// sorting the lists for easy comparing of values for the prompt
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	// since both lists should be same length, we can use either for iterator
	for i := 0; i < len(leftList); i++ {
		totalDistance += abs(leftList[i] - rightList[i])
		// fmt.Println("Row:", i, "Diff:", abs(leftList[i]-rightList[i]), "Total:", totalDistance)
	}

	// this should be the number we are looking for
	fmt.Println("Total Distance:", totalDistance)
}
