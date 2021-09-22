package main

import (
	"bufio"
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type point struct {
	x, y int
}

func (p *point) add(d point) point {
	return point{x: p.x + d.x, y: p.y + d.y}
}

func (p *point) distance(a point) int {
	return abs(p.x-a.x) + abs(p.y-a.y)
}

func Day6(scanner *bufio.Scanner) {
	// map day -> guard
	start := make([]point, 0, 1000)
	points := make(map[point]int)
	sizes := make(map[int]int)
	queue := make([]point, 0, 1000)
	var id, sx, sy, mx, my int = 1, 1000, 1000, 0, 0
	for scanner.Scan() {
		var p point
		fmt.Sscanf(scanner.Text(), "%d, %d", &p.x, &p.y)
		points[p] = id
		sizes[id] = 1
		queue = append(queue, p)
		start = append(start, p)
		mx = max(mx, p.x)
		my = max(my, p.y)
		sx = min(sx, p.x)
		sy = min(sy, p.y)
		id++
	}
	sx--
	sy--
	mx++
	my++
	md := max(mx-sx, my-sy)
	moves := [4]point{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}
	invalid := make(map[int]bool)
	for d := 1; d < md && len(queue) > 0; d++ {
		fmt.Printf("distance %d queue size %d\n", d, len(queue))
		nqueue := make([]point, 0, 1000)
		ncnt := make(map[point]int)
		nid := make(map[point]int)
		// take all the points at range d
		for _, p := range queue {
			cid := points[p]
			for _, m := range moves {
				t := p.add(m)
				//fmt.Printf("From %d,%d (%d) to %d,%d\n", p.x, p.y, cid, t.x, t.y)
				if t.x < sx || t.x > mx || t.y < sy || t.y > my {
					invalid[cid] = true
					//fmt.Printf("Invalid\n")
					continue
				}
				if _, ok := points[t]; !ok {
					if oid, ok := nid[t]; !ok {
						ncnt[t] = 1
						nid[t] = cid
						nqueue = append(nqueue, t)
					} else if oid != cid {
						ncnt[t]++
					}
				}
			}
		}
		queue = make([]point, 0, len(nqueue))
		for _, p := range nqueue {
			cid := nid[p]
			if ncnt[p] == 1 {
				points[p] = cid
				queue = append(queue, p)
				sizes[cid]++
				//fmt.Printf("id %d + %dx%d\n", cid, p.x, p.y)
			} else if ncnt[p] > 0 {
				points[p] = 0
				queue = append(queue, p)
				ncnt[p] = 0
			}
		}
	}
	ss := 0
	for y := sy; y <= my; y++ {
		for x := sx; x <= mx; x++ {
			td := 0
			for _, p := range start {
				td += p.distance(point{x: x, y: y})
				if td >= 10000 {
					break
				}
			}
			if td < 10000 {
				ss++
			}
			// fmt.Printf("%d,%d td %d\n", x, y, td)
			/*
				if v, ok := points[point{x: x, y: y}]; ok && v > 0 {
					fmt.Printf("%d", v)
				} else {
					fmt.Printf(".")
				}
			*/
		}
		//fmt.Println()
	}
	//fmt.Println()
	maxsize := 0
	for id, size := range sizes {
		fmt.Printf("id %d size %d infinite %t\n", id, size, invalid[id])
		if !invalid[id] {
			maxsize = max(size, maxsize)
		}
	}
	fmt.Println(maxsize)
	fmt.Println(ss)
}
