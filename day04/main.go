package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	total_score := 0
	var card int
	cards := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		card, err = strconv.Atoi(strings.TrimLeft(strings.Replace(parts[0], "Card ", "", 1), " "))
		if err != nil {
			log.Fatal(err)
		}

		sets := slices.DeleteFunc(strings.Split(parts[1], " | "), func(e string) bool {
			return e == ""
		})

		winning := slices.DeleteFunc(strings.Split(sets[0], " "), func(e string) bool {
			return e == ""
		})
		my := strings.Split(sets[1], " ")
		sort.Strings(winning)
		sort.Strings(my)
		score := 1
		matching := 0
		for i, j := 0, 0; i < len(winning) && j < len(my); {
			if winning[i] == my[j] {
				score = score << 1
				matching++
				i++
				j++
			} else if winning[i] < my[j] {
				i++
			} else {
				j++
			}
		}
		score = score >> 1
		total_score += score
		fmt.Printf("Card %d: %d\n", card, score)
		if val, ok := cards[card]; ok {
			cards[card] = val + 1
		} else {
			cards[card] = 1
		}
		for i := 1; i <= matching; i++ {
			if val, ok := cards[card+i]; ok {
				cards[card+i] = val + cards[card]
			} else {
				cards[card+i] = cards[card]
			}
		}

	}

	fmt.Printf("%d\n", total_score)

	total_count := 0
	for i, count := range cards {
		if i > card {
			break
		}
		total_count += count
	}

	fmt.Printf("%d\n", total_count)
}

func convert(input string) []int {
	var out []int
	for _, n := range strings.Split(input, " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}
	return out
}
