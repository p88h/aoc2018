package main

import (
	"bufio"
	"fmt"
	"sort"
)

type unit struct {
	x, y int
	kind byte
	hp   int
	str  int
}

func (p point) neighbors() []point {
	return []point{{p.x, p.y - 1}, {p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y + 1}}
}

func (u *unit) bfs(mapp [][]byte, units []*unit) bool {
	if u.hp <= 0 {
		return true
	}
	found := make([]point, 0, 100)
	queue := make([]point, 0, 10000)
	prev := make(map[point]point, 10000)
	start := point{u.x, u.y}
	queue = append(queue, start)
	prev[start] = start
	dist := make(map[point]int, 10000)
	dist[start] = 0
	pos := 0
	ld := -1
	for pos < len(queue) {
		p := queue[pos]
		d := dist[p]
		if ld >= 0 && d > ld {
			break
		}
		neigh := p.neighbors()
		for i := range neigh {
			np := neigh[i]
			ch := mapp[np.y][np.x]
			_, ok := dist[np]
			if ok {
				continue
			}
			prev[np] = p
			dist[np] = d + 1
			if ch == '.' {
				queue = append(queue, np)
			} else if ch != '#' && ch != u.kind {
				found = append(found, np)
				ld = d
			}
		}
		pos += 1
	}
	sort.Slice(found, func(i, j int) bool {
		return found[i].y < found[j].y || (found[i].y == found[j].y && found[i].x < found[j].x)
	})
	if len(found) == 0 {
		for j := range units {
			if units[j].kind != u.kind && units[j].hp > 0 {
				return true
			}
		}
		// fmt.Printf("Unit %c hp %d at %d,%d found no targets\n", u.kind, u.hp, u.x, u.y)
		return false
	}
	fp := prev[found[0]]
	for prev[fp] != start {
		fp = prev[fp]
	}
	if fp != start {
		// fmt.Printf("Unit %c at %d,%d move to target area %d,%d\n", u.kind, u.x, u.y, fp.x, fp.y)
		mapp[u.y][u.x] = '.'
		mapp[fp.y][fp.x] = u.kind
		u.x, u.y = fp.x, fp.y
	}
	var target *unit = nil
	minhp := 301
	neigh := fp.neighbors()
	for i := range neigh {
		p := neigh[i]
		ch := mapp[p.y][p.x]
		if ch != '.' && ch != '#' && ch != u.kind {
			for j := range units {
				if units[j].x == p.x && units[j].y == p.y && units[j].hp < minhp {
					minhp = units[j].hp
					target = units[j]
				}
			}
		}
	}
	if target != nil {
		//fmt.Printf("Unit %c at %d,%d attacks target unit %c at %d,%d hp %d\n",
		//	u.kind, u.x, u.y, target.kind, target.x, target.y, target.hp)
		target.hp -= u.str
		if target.hp <= 0 {
			//fmt.Printf("Removed unit %c at %d,%d hp %d\n", target.kind, target.x, target.y, target.hp)
			mapp[target.y][target.x] = '.'
			target.x, target.y = -1, -1
		}
	}
	return true
}

func battle(srcmap []string, elfpower int) bool {
	mapp := make([][]byte, 0, 1000)
	units := make([]*unit, 0, 100)
	y := 0
	for l := range srcmap {
		line := srcmap[l]
		mapp = append(mapp, []byte(line))
		for x := range line {
			kind := line[x]
			if kind == 'G' {
				units = append(units, &unit{x, y, kind, 200, 3})
			}
			if kind == 'E' {
				units = append(units, &unit{x, y, kind, 200, elfpower})
			}
		}
		y += 1
	}
	last := 0
	round := 0
	for {
		round += 1
		/*
			fmt.Printf("Round %d begins\n", round)
			for y := range mapp {
				fmt.Println(string(mapp[y]))
			}
		*/
		if last > 0 {
			break
		}
		sort.Slice(units, func(i, j int) bool {
			return units[i].y < units[j].y || (units[i].y == units[j].y && units[i].x < units[j].x)
		})
		for i := range units {
			u := units[i]
			if !u.bfs(mapp, units) && last == 0 {
				last = round - 1
			}
		}
	}
	hp := 0
	elfdeaths := 0
	for i := range units {
		if units[i].hp > 0 {
			hp += units[i].hp
		} else if units[i].kind == 'E' {
			elfdeaths += 1
		}
	}
	fmt.Printf("Last full round %d total HP remaining %d outcome %d\n", last, hp, last*hp)
	fmt.Printf("Dead elves at power %d: %d\n", elfpower, elfdeaths)
	return elfdeaths > 0
}

func (x Aoc) Day15(scanner *bufio.Scanner) {
	lines := make([]string, 0, 1000)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for elfpower := 3; battle(lines, elfpower); elfpower++ {
	}
}
