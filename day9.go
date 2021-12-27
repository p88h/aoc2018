package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type ball struct {
	value int64
	prev  *ball
	next  *ball
}

func Game(n int64, l int64) int64 {
	b := &ball{value: 0}
	// z := b
	b.next = b
	b.prev = b
	w := make([]int64, n+1, 1024)
	p := int64(1)
	for i := int64(1); i <= l; i++ {
		if i%23 != 0 {
			b = b.next
			t := &ball{value: i}
			t.next = b.next
			t.next.prev = t
			t.prev = b
			b.next = t
			b = b.next
		} else {
			w[p] += i
			for j := 0; j < 7; j++ {
				b = b.prev
			}
			b.prev.next = b.next
			b.next.prev = b.prev
			w[p] += b.value
			b = b.next
		}
		/*
			fmt.Printf("[%2d] %2d", i, z.value)
			for t := z.next; t != z; t = t.next {
				if t == b {
					fmt.Printf("(%2d) ", t.value)
				} else {
					fmt.Printf(" %2d  ", t.value)
				}
			}
			fmt.Println()
		*/
		p = (p % n) + 1
	}
	max := int64(0)
	for p := int64(1); p <= n; p++ {
		if w[p] > max {
			max = w[p]
		}
	}
	return max
}

func (x Aoc) Day9(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`(\d+) players; last marble is worth (\d+) points`)
	scanner.Scan()
	g := r.FindStringSubmatch(scanner.Text())
	n, _ := strconv.ParseInt(g[1], 10, 0)
	l, _ := strconv.ParseInt(g[2], 10, 0)
	fmt.Printf("Day 9 Part 1: %d\n", Game(n, l))
	fmt.Printf("Day 9 Part 2: %d\n", Game(n, l*100))
}
