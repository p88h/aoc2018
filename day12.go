package main

import (
	"bufio"
	"fmt"
	"strings"
)

func (x Aoc) Day12(scanner *bufio.Scanner) {
	scanner.Scan()
	plants := scanner.Text()[15:]
	pad := 4
	ofs := 2 * pad
	work := strings.Repeat(".", ofs) + plants + strings.Repeat(".", ofs)
	fmt.Println(plants)
	rules := make(map[string]byte)
	scanner.Scan()
	for scanner.Scan() {
		pat := strings.Split(scanner.Text(), " => ")
		rules[pat[0]] = pat[1][0]
	}
	ptotal := 0
	for round := 1; round <= 200; round++ {
		total := 0
		b := make([]byte, len(work)+pad*2)
		nofs := ofs + pad
		p0 := len(work) + pad*2
		p1 := 0
		for i := range b {
			if i < 2*pad || i >= len(work) {
				b[i] = '.'
			} else {
				p := work[i-2-pad : i+3-pad]
				b[i] = rules[p]
				if b[i] == '#' {
					total += i - nofs
					if i < p0 {
						p0 = i
					}
					if i > p1 {
						p1 = i
					}
				} else {
					b[i] = '.'
				}
			}
		}
		ofs = nofs - (p0 - 2*pad)
		prev := work
		work = string(b[p0-2*pad : p1+2*pad])
		if round == 20 {
			fmt.Println(round, work[2*pad:(p1-p0)+2*pad+1], total)
		}
		if work == prev {
			fmt.Println(round, work[2*pad:(p1-p0)+2*pad+1], total, total-ptotal)
			if round >= 100 {
				fmt.Println(total + (50000000000-round)*(total-ptotal))
				return
			}
		}
		ptotal = total
	}
}
