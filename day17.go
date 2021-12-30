package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func settled(ch byte) bool {
	return ch == '#' || ch == '~'
}

func (x Aoc) Day17(scanner *bufio.Scanner) {
	pat := regexp.MustCompile(`(x|y)=(\d+), (x|y)=(\d+)..(\d+)`)
	mapp := make([][]byte, 2000)
	maxy := int64(0)
	miny := int64(9999)
	for y := 0; y < 2000; y++ {
		mapp[y] = make([]byte, 1000)
	}
	for scanner.Scan() {
		g := pat.FindStringSubmatch(scanner.Text())
		v, _ := strconv.ParseInt(g[2], 10, 0)
		a, _ := strconv.ParseInt(g[4], 10, 0)
		b, _ := strconv.ParseInt(g[5], 10, 0)
		for i := a; i <= b; i++ {
			if g[1] == "x" {
				mapp[i][v] = '#'
				if i > maxy {
					maxy = i
				}
				if i < miny {
					miny = i
				}
			} else {
				mapp[v][i] = '#'
				if v > maxy {
					maxy = v
				}
				if v < miny {
					miny = v
				}
			}
		}
	}
	flows := make([]map[int64]bool, 2000)
	flows[1] = make(map[int64]bool, 1)
	flows[1][500] = true
	deep := int64(1)
	tot := int64(0)
	tot2 := int64(0)
	for deep <= maxy {
		fmt.Printf("@Depth %d Flows: %v\n", deep, flows[deep])
		dir := int64(1)
		if flows[deep+1] == nil {
			flows[deep+1] = make(map[int64]bool, 0)
		}
		for p := range flows[deep] {
			if settled(mapp[deep+1][p]) {
				pl := p - 1
				pr := p + 1
				for settled(mapp[deep+1][pl]) && mapp[deep][pl] != '#' {
					pl--
				}
				for settled(mapp[deep+1][pr]) && mapp[deep][pr] != '#' {
					pr++
				}
				ch := byte('~')
				if !settled(mapp[deep+1][pr]) && mapp[deep][pr] != '#' {
					flows[deep+1][pr] = true
					ch = '|'
				}
				if !settled(mapp[deep+1][pl]) && mapp[deep][pl] != '#' {
					flows[deep+1][pl] = true
					ch = '|'
				}
				if ch == '~' {
					delete(flows[deep], p)
					dir = -1
				}
				for i := pl; i <= pr; i++ {
					if mapp[deep][i] != '#' {
						if mapp[deep][i] == 0 {
							tot += 1
						}
						if ch == '~' && mapp[deep][i] != '~' {
							tot2 += 1
						}
						mapp[deep][i] = ch
					}
				}
			} else {
				if mapp[deep][p] == 0 {
					tot += 1
				}
				mapp[deep][p] = '|'
				flows[deep+1][p] = true
			}
		}
		deep = deep + dir
	}
	fmt.Printf("Part 1: %d\n", tot-miny+1)
	fmt.Printf("Part 2: %d\n", tot2)
}
