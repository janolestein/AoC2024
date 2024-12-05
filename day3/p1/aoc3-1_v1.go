package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    total := 0
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    allHits := re.FindAllStringSubmatch(text, -1)
    for _, v := range allHits{
        num1, _ := strconv.Atoi(v[1])
        num2, _ := strconv.Atoi(v[2])
        total += num1 * num2
    }
    fmt.Println(total)


}
