package main

import (
	"bufio"
	"fmt"
)

func plot(s string, p *int, heads map[point]bool, mapp map[point]byte) (ret map[point]bool) {
	dirs := map[byte]point{'N': {0, -1}, 'S': {0, 1}, 'W': {-1, 0}, 'E': {1, 0}}
	doors := map[byte]byte{'N': '-', 'S': '_', 'W': '|', 'E': '|'}
	for s[*p] != '|' && s[*p] != ')' && s[*p] != '$' {
		ch := s[*p]
		ret = make(map[point]bool, len(heads))
		if ch == '(' {
			for s[*p] != ')' {
				*p++
				for p, _ := range plot(s, p, heads, mapp) {
					ret[p] = true
				}
			}
			*p++
		} else {
			dir := dirs[ch]
			for p, _ := range heads {
				q := p.add(dir)
				mapp[q] = doors[ch]
				q = q.add(dir)
				mapp[q] = '.'
				ret[q] = true
			}
			*p++
			heads = ret
		}
	}
	return
}

func explore(start point, mapp map[point]byte) (maxp int64, nump int) {
	queue := make([]point, 1, len(mapp))
	dist := make(map[point]int64, len(mapp))
	queue[0] = start
	dist[start] = 0
	pos := 0
	dirs := []point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for pos < len(queue) {
		p := queue[pos]
		d := dist[p]
		for _, m := range dirs {
			q := p.add(m)
			_, ok := mapp[q]
			if ok {
				q = q.add(m)
				_, done := dist[q]
				if !done {
					dist[q] = d + 1
					queue = append(queue, q)
				}
			}
		}
		pos++
	}
	nump = 0
	for l := 0; l < len(queue); l++ {
		maxp = dist[queue[l]]
		if dist[queue[l]] >= 1000 {
			nump++
		}
	}
	return
}

func (x Aoc) Day20(scanner *bufio.Scanner) {
	for scanner.Scan() {
		path := scanner.Text()
		pos := 1
		start := point{0, 0}
		mapp := make(map[point]byte, 10000)
		mapp[start] = '.'
		plot(path, &pos, map[point]bool{start: true}, mapp)
		p1, p2 := explore(start, mapp)
		fmt.Printf("Day 20 Part 1: %d\n", p1)
		fmt.Printf("Day 20 Part 1: %d\n", p2)
	}
}
