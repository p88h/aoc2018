package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func energy_value(x, y, r int64) int64 {
	return (((((x+10)*y + r) * (x + 10)) / 100) % 10) - 5
}

func energy_line(y, r, b int64) []int64 {
	a := make([]int64, 0, 300-b+1)
	e := int64(0)
	for x := int64(1); x < b; x++ {
		e += energy_value(x, y, r)
	}
	for x := int64(b); x <= 300; x++ {
		e += energy_value(x, y, r)
		a = append(a, e)
		e -= energy_value(x-b+1, y, r)
	}
	return a
}

func find_best(n, b int64) (max, tx, ty int64) {
	l := energy_line(1, n, b)
	for y := int64(2); y < b; y++ {
		t := energy_line(y, n, b)
		for x := range t {
			l[x] += t[x]
		}
	}
	max = int64(0)
	for y := int64(b); y <= 300; y++ {
		p := energy_line(y-b+1, n, b)
		t := energy_line(y, n, b)
		for x := range t {
			l[x] += t[x]
			if l[x] > max {
				max = l[x]
				tx = int64(x) + 1
				ty = y - b + 1
			}
			l[x] -= p[x]
		}
	}
	return
}

func Day11(scanner *bufio.Scanner) {
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	fmt.Printf("Day 11 Part 0: %d\n", n)
	max, tx, ty := find_best(n, 3)
	fmt.Printf("Day 11 Part 1: %d @ %d,%d\n", max, tx, ty)
	var amax, atx, aty, abs int64 = 0, 0, 0, 0
	for b := int64(1); b <= 300; b++ {
		max, tx, ty = find_best(n, b)
		if max > amax {
			amax = max
			atx = tx
			aty = ty
			abs = b
		}
	}
	fmt.Printf("Day 11 Part 1: %d @ %d,%d,%d\n", amax, atx, aty, abs)
}
