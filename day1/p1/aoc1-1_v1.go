package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left := make([]int, 0)
	right := make([]int, 0)
    dist := 0
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
    sort.Slice(left, func(i, j int) bool {
        return left[i] < left[j]
    })
    sort.Slice(right, func(i, j int) bool {
        return right[i] < right[j]
    })

    for i := range left{
        dist += int(math.Abs(float64(left[i]) - float64(right[i])))
    }

    fmt.Printf("Total Distance: %d \n", dist)
}
