package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5",
		"six", "6", "seven", "7", "eight", "8", "nine", "9")

	replace := true
	var numbers []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := []rune{0, 0}

		line := scanner.Text()

		if replace == true {
			line = replacer.Replace(line)
		}

		for _, c := range line {
			if c < '0' {
				continue
			}
			if c > '9' {
				continue
			}
			if current[0] == 0 {
				current[0] = c
			}
			current[1] = c
		}
		numbers = append(numbers, string(current))
	}

	sum := 0
	for _, n := range numbers {
		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		sum += i
	}

	print(sum)
}
