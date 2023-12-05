package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	destination uint
	length      uint
}

func main() {
	file, err := os.Open("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var seeds []uint
	var maps []map[uint]Mapping

	var current_map map[uint]Mapping

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			for _, p := range strings.Split(line, " ") {
				seed, err := strconv.Atoi(p)
				if err != nil {
					continue
				}
				seeds = append(seeds, uint(seed))
			}
			continue
		}

		if len(strings.TrimSpace(line)) == 0 {
			if len(current_map) > 0 {
				maps = append(maps, current_map)
			}
			continue
		}

		if strings.HasSuffix(line, "map:") {
			current_map = make(map[uint]Mapping)
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			src, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}
			dst, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal(err)
			}
			rng, err := strconv.Atoi(parts[2])
			if err != nil {
				log.Fatal(err)
			}
			current_map[uint(src)] = Mapping{destination: uint(dst), length: uint(rng)}
		}
	}
	if len(current_map) > 0 {
		maps = append(maps, current_map)
	}

	lowestLocation := ^uint(0)

	for _, seed := range seeds {
		loc := seed
		for _, mapping := range maps {
			loc = searchDestination(loc, mapping)
		}
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	fmt.Println(lowestLocation)

	// part 2
	sets := make([][2]uint, 0, len(seeds)/2)
	for i := 1; i < len(seeds); i += 2 {
		sets = append(sets, [2]uint{seeds[i-1], seeds[i]})
	}

	lowestLocation = ^uint(0)

	for _, set := range sets {
		loc := traverse(set[0], set[1], maps)
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	fmt.Println(lowestLocation)
}

func traverse(src uint, size uint, mappings []map[uint]Mapping) uint {
	if len(mappings) == 0 {
		return src
	}
	mapping := mappings[0]
	var overlaps []uint

	for k, m := range mapping {
		if k <= src && src < k+m.length {
			offset := src - k
			if k+size+offset <= k+m.length {
				return traverse(m.destination+offset, size, mappings[1:])
			}
			// split
			loc_a := traverse(m.destination+offset, m.length-offset, mappings[1:])
			loc_b := traverse(src+(m.length-offset), size-(m.length-offset), mappings)
			if loc_a < loc_b {
				return loc_a
			}
			return loc_b
		}
		if src < k && k < src+size {
			overlaps = append(overlaps, k)
		}
	}

	if len(overlaps) == 0 {
		return traverse(src, size, mappings[1:])
	}

	sort.Slice(overlaps, func(i, j int) bool { return overlaps[i] < overlaps[j] })

	loc := ^uint(0)
	overlap := overlaps[0]
	if src < overlap {
		tmp := traverse(src, overlap-src, mappings)
		if tmp < loc {
			loc = tmp
		}
		size -= overlap - src
		src = overlap
	}

	tmp := traverse(src, size, mappings)
	if tmp < loc {
		return tmp
	}
	return loc
}

func searchDestination(src uint, mapping map[uint]Mapping) uint {
	for k, m := range mapping {
		if k <= src && k+m.length >= src {
			return m.destination + (src - k)
		}
		if k > src {
			continue
		}
	}
	return src
}
