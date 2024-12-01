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
	left := make([]int, 0)
	right := make([]int, 0)
    simScore := 0
	file, err := os.Open("../../input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pairs := strings.Fields(scanner.Text())
        leftNumber, _ := strconv.Atoi(pairs[0])
        rightNumber, _ := strconv.Atoi(pairs[1])
        left = append(left, leftNumber)
        right = append(right, rightNumber)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    freq := make(map[int]int)
    for _, num := range right{
        freq[num] = freq[num] + 1
    }

    for _, num := range left{
       simScore += num * freq[num] 
    }

    fmt.Printf("Sim Score: %d \n", simScore)
}
