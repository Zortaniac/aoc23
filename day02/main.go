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
	file, err := os.Open("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	id_sum := 0
	set_power := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		game, err := strconv.Atoi(strings.Replace(parts[0], "Game ", "", 1))
		if err != nil {
			log.Fatal(err)
		}

		sets := strings.Split(parts[1], "; ")

		red, green, blue := 0, 0, 0
		sucess := true
		for _, s := range sets {
			if !testSet(s) {
				sucess = false
			}
			cubes := strings.Split(s, ", ")
			for _, c := range cubes {
				cp := strings.Split(c, " ")
				num, err := strconv.Atoi(cp[0])
				if err != nil {
					log.Fatal(err)
				}
				switch cp[1] {
				case "red":
					if red < num {
						red = num
					}
					break
				case "green":
					if green < num {
						green = num
					}
					break
				case "blue":
					if blue < num {
						blue = num
					}
					break
				}
			}
		}

		set_power += red * green * blue
		if sucess {
			id_sum += game
		}
	}

	fmt.Printf("%d\n", id_sum)
	fmt.Printf("%d\n", set_power)
}

func testSet(set string) bool {
	cubes := strings.Split(set, ", ")
	for _, c := range cubes {
		cp := strings.Split(c, " ")
		num, err := strconv.Atoi(cp[0])
		if err != nil {
			log.Fatal(err)
		}
		switch cp[1] {
		case "red":
			if num > 12 {
				return false
			}
			break
		case "green":
			if num > 13 {
				return false
			}
			break
		case "blue":
			if num > 14 {
				return false
			}
			break
		}
	}
	return true
}
