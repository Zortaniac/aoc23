package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	x int
	y int
}

func main() {
	file, err := os.Open("./day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, "."+line+".")
	}

	lines = append([]string{strings.Repeat(".", len(lines[0]))}, lines...)
	lines = append(lines, lines[0])

	gears := make(map[Key][]int)
	sum := 0
	for i := 1; i < len(lines)-1; i++ {
		current := ""
		valid := false
		var gearMatch []Key
		for p, c := range lines[i] {
			if c < '0' || c > '9' {
				if current == "" {
					continue
				}

				number, err := strconv.Atoi(current)
				if err != nil {
					log.Fatal(err)
				}

				if c == '*' {
					gearMatch = append(gearMatch, Key{i, p})
				}
				if lines[i-1][p] == '*' {
					gearMatch = append(gearMatch, Key{i - 1, p})
				}
				if lines[i+1][p] == '*' {
					gearMatch = append(gearMatch, Key{i + 1, p})
				}

				for _, match := range gearMatch {
					if val, ok := gears[match]; ok {
						gears[match] = append(val, number)
					} else {
						gears[match] = []int{number}
					}
				}

				if valid || test(lines[i][p]) || test(lines[i-1][p]) || test(lines[i+1][p]) {
					// add number
					sum += number
					fmt.Printf("%d\n", number)
				}
				current = ""
				valid = false
				gearMatch = nil
				continue
			}

			current += string(c)

			if lines[i-1][p] == '*' {
				gearMatch = append(gearMatch, Key{i - 1, p})
			}
			if lines[i+1][p] == '*' {
				gearMatch = append(gearMatch, Key{i + 1, p})
			}
			if len(current) == 1 {
				if lines[i][p-1] == '*' {
					gearMatch = append(gearMatch, Key{i, p - 1})
				}
				if lines[i-1][p-1] == '*' {
					gearMatch = append(gearMatch, Key{i - 1, p - 1})
				}
				if lines[i+1][p-1] == '*' {
					gearMatch = append(gearMatch, Key{i + 1, p - 1})
				}
			}

			if valid {
				continue
			}

			if test(lines[i-1][p]) || test(lines[i+1][p]) {
				valid = true
				continue
			}

			if len(current) != 1 {
				continue
			}

			if test(lines[i][p-1]) || test(lines[i-1][p-1]) || test(lines[i+1][p-1]) {
				valid = true
				continue
			}
		}
	}
	fmt.Printf("\n%d\n", sum)

	gearProducts := 0
	for _, nums := range gears {
		if len(nums) != 2 {
			continue
		}
		gearProducts += nums[0] * nums[1]
	}

	fmt.Printf("\n%d\n", gearProducts)
}

func test(c uint8) bool {
	if c == '.' {
		return false
	}
	if c > '0' && c < '9' {
		return false
	}
	if c > 'A' && c < 'Z' {
		return false
	}
	if c > 'a' && c < 'z' {
		return false
	}
	return true
}
