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
		if conNums[0] > conNums[1] {
			for i, _ := range conNums {
				if i+1 >= len(conNums) {
					safeCount++
					break
				}
				diff := conNums[i] - conNums[i+1]
				if diff < 1 || diff > 3 {
					break
				}
			}

		} else {
			for i, _ := range conNums {
				if i+1 >= len(conNums) {
					safeCount++
					break
				}
				diff := conNums[i] - conNums[i+1]
				if diff > -1 || diff < -3 {
					break
				}
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(safeCount)
}

