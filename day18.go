package main

import (
	"bufio"
	"fmt"
)

func (x Aoc) Day18(scanner *bufio.Scanner) {
	mapp := make([]string, 0, 50)
	for scanner.Scan() {
		mapp = append(mapp, scanner.Text())
	}
	neigh := []point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	prev := make(map[int]int)
	lastcycle := 0
	for round := 0; round < 1000000000; round++ {
		map2 := make([]string, 50)
		for y, l := range mapp {
			line := make([]byte, 50)
			for x, ch := range l {
				tc, lc := 0, 0
				for _, n := range neigh {
					qx, qy := x+n.x, y+n.y
					if qx >= 0 && qy >= 0 && qy < len(mapp) && qx < len(mapp[qy]) {
						switch mapp[qy][qx] {
						case '|':
							tc++
						case '#':
							lc++
						}
					}
				}
				switch ch {
				case '.':
					if tc >= 3 {
						line[x] = '|'
					} else {
						line[x] = '.'
					}
				case '|':
					if lc >= 3 {
						line[x] = '#'
					} else {
						line[x] = '|'
					}
				case '#':
					if lc >= 1 && tc >= 1 {
						line[x] = '#'
					} else {
						line[x] = '.'
					}
				}
			}
			map2 = append(map2, string(line))
		}
		mapp = map2
		lc, tc := 0, 0
		for _, l := range mapp {
			for _, c := range l {
				switch c {
				case '|':
					tc++
				case '#':
					lc++
				}
			}
		}
		key := lc * tc
		if prev[key] > 0 {
			cycle := round + 1 - prev[key]
			rem := (1000000000 - (round + 1)) % cycle
			fmt.Printf("Detected a cycle maybe at round %d from round %d cycle %d rem %d key %d\n", round+1, prev[key], cycle, rem, key)
			if cycle == lastcycle && rem == 0 {
				fmt.Printf("Day 18 Part 2: %d\n", key)
				return
			}
			lastcycle = cycle
		}
		prev[key] = round + 1
		if round+1 == 10 {
			fmt.Printf("Day 18 Part 1: %d\n", key)
		}
	}
}
