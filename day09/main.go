package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func main() {
	file, err := os.Open("./day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	sum := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var numbers []int

		for _, p := range strings.Split(line, " ") {
			num, err := strconv.Atoi(p)
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}
		newNum := numbers[len(numbers)-1] + calc(numbers)
		reverse(numbers)
		newNum2 := numbers[len(numbers)-1] + calc(numbers)
		sum += newNum
		sum2 += newNum2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func calc(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	distances := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		distance := numbers[i+1] - numbers[i]
		distances[i] = distance
	}
	return distances[len(distances)-1] + calc(distances)
}

func travers(numbers []int, n int) int {
	num := get(numbers, n, 0)
	if num == 0 {
		return 0
	}
	return num + travers(numbers, n+1)
}

func get(numbers []int, n int, m int) int {
	if n == 0 {
		return numbers[m]
	}

	val := get(numbers, n-1, m) - get(numbers, n-1, m+1)
	if val < 0 {
		return -val
	}
	return val
}
