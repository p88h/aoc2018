package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type chocolate struct {
	value int64
	next  *chocolate
}

type factory struct {
	a, b, head, tail *chocolate
}

func step(f *factory) int {
	cnt := 1
	s := f.a.value + f.b.value
	if s > 9 {
		c := &chocolate{value: s / 10}
		f.tail.next = c
		c.next = f.head
		f.tail = c
		cnt = 2
	}
	d := &chocolate{value: s % 10}
	f.tail.next = d
	d.next = f.head
	f.tail = d
	ac := f.a.value + 1
	for j := int64(0); j < ac; j++ {
		f.a = f.a.next
	}
	bc := f.b.value + 1
	for j := int64(0); j < bc; j++ {
		f.b = f.b.next
	}
	return cnt
}

func setup() (f *factory) {
	f = &factory{a: &chocolate{value: 3}, b: &chocolate{value: 7}}
	f.a.next = f.b
	f.b.next = f.a
	f.tail = f.b
	f.head = f.a
	return
}

func Brew(n int64) int64 {
	f := setup()
	tot := int64(2)
	for tot < n+10 {
		tot += int64(step(f))
	}
	v := int64(0)
	for i := int64(0); i < n+10; i++ {
		if i >= n {
			v = v*10 + f.head.value
		}
		f.head = f.head.next
	}
	return v
}

func Brew2(n int64, l int) int {
	f := setup()
	val := f.head.value*10 + f.tail.value
	tot := 2
	tmp := f.tail
	mod := int64(1)
	for j := 0; j < l; j++ {
		mod *= 10
	}
	for val != n {
		k := step(f)
		for j := 0; j < k; j++ {
			tot += 1
			tmp = tmp.next
			val = (val*10 + tmp.value) % mod
			if val == n {
				break
			}
		}
	}
	return tot - l
}

func (x Aoc) Day14(scanner *bufio.Scanner) {
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		fmt.Printf("Day 14 Part 1 (%d): %010d\n", n, Brew(n))
		fmt.Printf("Day 14 Part 2 (%d): %d\n", n, Brew2(n, len(scanner.Text())))
	}
}
