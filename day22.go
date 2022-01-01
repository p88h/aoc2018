package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Tool int

const (
	Torch Tool = iota
	Gear
	None
)

type modenode struct {
	p    point
	tool Tool
}

func genmoves(cd int, current modenode, np point, erosion map[point]int, moves map[modenode]int) {
	t2t := map[int][]Tool{0: {Gear, Torch}, 1: {Gear, None}, 2: {Torch, None}}
	for _, tool := range t2t[erosion[np]%3] {
		invalid := 0
		switch tool {
		case Torch:
			invalid = 1
		case Gear:
			invalid = 2
		case None:
			invalid = 0
		}
		ttype := erosion[current.p] % 3
		if tool == current.tool {
			moves[modenode{np, tool}] = cd + 1
		} else if ttype != invalid {
			moves[modenode{np, tool}] = cd + 8
		}
	}
}

func (x Aoc) Day22(scanner *bufio.Scanner) {
	r1 := regexp.MustCompile(`depth: (\d+)`)
	r2 := regexp.MustCompile(`target: (\d+),(\d+)`)
	scanner.Scan()
	g1 := r1.FindStringSubmatch(scanner.Text())
	depth, _ := strconv.ParseInt(g1[1], 10, 0)
	scanner.Scan()
	g2 := r2.FindStringSubmatch(scanner.Text())
	tx, _ := strconv.ParseInt(g2[1], 10, 0)
	ty, _ := strconv.ParseInt(g2[2], 10, 0)
	target := point{int(tx), int(ty)}
	origin := point{0, 0}
	erosion := make(map[point]int, tx*ty)
	erosion[origin] = int(depth) % 20183
	maxx := target.x * 6
	maxy := target.y * 2
	for x := 1; x <= maxx; x++ {
		g := (x * 16807)
		erosion[point{x, 0}] = (g + int(depth)) % 20183
	}
	for y := 1; y <= maxy; y++ {
		g := (y * 48271)
		erosion[point{0, y}] = (g + int(depth)) % 20183
		for x := 1; x <= maxx; x++ {
			g = erosion[point{x - 1, y}] * erosion[point{x, y - 1}]
			erosion[point{x, y}] = (g + int(depth)) % 20183
		}
	}
	erosion[target] = erosion[origin]
	risk := 0
	segs := map[int]byte{0: '.', 1: '=', 2: '|'}
	for y := 0; y <= target.y; y++ {
		line := make([]byte, target.x+1)
		for x := 0; x <= target.x; x++ {
			rt := erosion[point{x, y}] % 3
			risk += rt
			line[x] = segs[rt]
		}
		//fmt.Println(string(line))
	}
	fmt.Printf("Day 22 Part 1: %d\n", risk)
	start := modenode{origin, Torch}
	dest := modenode{target, Torch}
	distance := make(map[modenode]int, target.x*target.y)
	limit := (target.x + target.y) * 8
	queue := make([][]modenode, limit+1)
	queue[0] = append(queue[0], start)
	distance[start] = 0
	distance[dest] = limit
	dirs := []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for cd := 0; cd <= distance[dest]; cd++ {
		if queue[cd] == nil {
			continue
		}
		for _, mn := range queue[cd] {
			tmp := make(map[modenode]int, 100)
			for _, dir := range dirs {
				np := mn.p.add(dir)
				if np.x < 0 || np.y < 0 || np.x >= maxx || np.y >= maxy {
					continue
				}
				genmoves(cd, mn, np, erosion, tmp)
			}
			for nn, nd := range tmp {
				prev, ok := distance[nn]
				if nd <= limit && (!ok || nd < prev) {
					queue[nd] = append(queue[nd], nn)
					distance[nn] = nd
				}
			}
		}
	}
	fmt.Printf("Day 22 Part 2: %d\n", distance[dest])
}
