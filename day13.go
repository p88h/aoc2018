package main

import (
	"bufio"
	"fmt"
	"sort"
)

type cart struct {
	x, y, dx, dy, t int
	crashed         bool
}

func (x Aoc) Day13(scanner *bufio.Scanner) {
	mapp := make([]string, 0, 1000)
	carts := make([]*cart, 0, 100)
	chars := map[byte]cart{'<': cart{dx: -1}, '>': cart{dx: 1}, '^': cart{dy: -1}, 'v': cart{dy: 1}}
	y := 0
	occ := make(map[point]*cart)
	for scanner.Scan() {
		line := scanner.Text()
		mapp = append(mapp, line)
		for x := range line {
			car, ok := chars[line[x]]
			if ok {
				car.x = x
				car.y = y
				carts = append(carts, &car)
				occ[point{car.x, car.y}] = &car
			}
		}
		y += 1
	}
	riding := len(carts)
	first := point{-1, -1}
	for riding > 1 {
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].y < carts[j].y || carts[i].x < carts[j].x
		})
		for c := range carts {
			car := carts[c]
			if car.crashed {
				continue
			}
			delete(occ, point{car.x, car.y})
			car.x += car.dx
			car.y += car.dy
			p := point{car.x, car.y}
			other, crash := occ[p]
			if crash {
				//fmt.Printf("crash@%d,%d | ", car.x, car.y)
				delete(occ, p)
				car.crashed = true
				other.crashed = true
				riding -= 2
				if first.x < 0 {
					first = p
				}
				continue
			}
			occ[p] = car
			switch ch := mapp[car.y][car.x]; ch {
			case '+':
				switch car.t {
				case 0:
					car.dx, car.dy = car.dy, -car.dx
					car.t = 1
				case 1:
					car.t = 2
				case 2:
					car.dx, car.dy = -car.dy, car.dx
					car.t = 0
				}
			case '/':
				car.dx, car.dy = -car.dy, -car.dx
			case '\\':
				car.dx, car.dy = car.dy, car.dx
			default:
			}
			// fmt.Printf("%d@%d,%d+%d,%d | ", c+1, car.x, car.y, car.dx, car.dy)
		}
		// fmt.Println(riding)
	}
	last := carts[0]
	for c := range carts {
		if !carts[c].crashed {
			last = carts[c]
		}
	}
	fmt.Printf("First crash at %d,%d\n", first.x, first.y)
	fmt.Printf("Last car at %d,%d\n", last.x, last.y)
}
