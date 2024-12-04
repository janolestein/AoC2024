package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		conNums := make([]int, 0)
		for _, v := range nums {
			conv, _ := strconv.Atoi(v)
			conNums = append(conNums, conv)
		}
		if isRowSave(conNums) {
			safeCount++
		} else if isRowWithRemovalSafe(conNums) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(safeCount)
}

func isRowWithRemovalSafe(conNums []int) bool {
	for i := range conNums {
		isSafe := checkAfterIndexRemoval(conNums, i)
		if isSafe {
			return true
		}
	}
	return false
}

func checkAfterIndexRemoval(conNums []int, index int) bool {
	copyNums := make([]int, len(conNums))
	copy(copyNums, conNums)

	if index == len(copyNums)-1 {
		copyNums = copyNums[:index]
	} else {
		copyNums = append(copyNums[:index], copyNums[index+1:]...)
	}
	return isRowSave(copyNums)
}

func isRowSave(conNums []int) bool {
	if conNums[0] > conNums[1] {
		for i, _ := range conNums {
			if i+1 >= len(conNums) {
				return true
			}
			diff := conNums[i] - conNums[i+1]
			if diff < 1 || diff > 3 {
				break
			}
		}

	} else {
		for i, _ := range conNums {
			if i+1 >= len(conNums) {
				return true
			}
			diff := conNums[i] - conNums[i+1]
			if diff > -1 || diff < -3 {
				break
			}
		}

	}
	return false
}

