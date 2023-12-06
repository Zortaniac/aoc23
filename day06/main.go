package main

import "fmt"

type Race struct {
	distance int
	time     int
}

func main() {
	// test data part 1
	//races := []Race{{distance: 9, time: 7}, {distance: 40, time: 15}, {distance: 200, time: 30}}
	// data part 1
	//races := []Race{{distance: 222, time: 51}, {distance: 2031, time: 92}, {distance: 1126, time: 68}, {distance: 1225, time: 90}}
	// test data part 2
	//races := []Race{{distance: 940200, time: 71530}}
	// data part 2
	races := []Race{{distance: 222203111261225, time: 51926890}}

	score := 1
	for _, race := range races {
		count := 0
		for i := 1; i < race.time; i++ {
			if i*(race.time-i) > race.distance {
				count++
			}
		}
		score *= count
		fmt.Println(count)
	}

	fmt.Println(score)
}
