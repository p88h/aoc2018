package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parse(a []int64, pos int64) (int64, int64) {
	temp := make([]int64, 0, 1024)
	sp := a[pos]
	mp := a[pos+1]
	//fmt.Printf("pkt start %d sub %d meta %d\n", pos, sp, mp)
	smeta := int64(0)
	pos += 2
	for i := int64(0); i < sp; i++ {
		npos, meta := parse(a, pos)
		temp = append(temp, meta)
		pos = npos
		// smeta += meta
	}
	for i := int64(0); i < mp; i++ {
		if sp == 0 {
			smeta += a[pos]
		} else if a[pos] > 0 && a[pos] <= sp {
			smeta += temp[a[pos]-1]
		}
		pos += 1
	}
	//fmt.Printf("pkt end %d smeta %d\n", pos, smeta)
	return pos, smeta
}

func Day8(scanner *bufio.Scanner) {
	sl := make([]int64, 0, 16384)
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), " ") {
		var x, _ = strconv.ParseInt(v, 10, 64)
		sl = append(sl, x)
	}
	epos, meta := parse(sl, 0)
	fmt.Printf("Day 18 Part 2 %d %d\n", meta, epos)
}
