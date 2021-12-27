package main

import (
	"bufio"
	"fmt"
	"strings"
)

func (x Aoc) Day2(scanner *bufio.Scanner) {
	var c2, c3 int = 0, 0
	sl := make([]string, 0, 1024)
	var sb strings.Builder
	for scanner.Scan() {
		var s string = scanner.Text()
		m := make(map[rune]int)
		for _, c := range s {
			m[c] += 1
		}
		for c := range m {
			if m[c] == 2 {
				c2++
				break
			}
		}
		for c := range m {
			if m[c] == 3 {
				c3++
				break
			}
		}
		for _, p := range sl {
			d := 0
			for i := range p {
				if p[i] != s[i] {
					d++
				}
			}
			if d == 1 {
				for i, r := range p {
					if p[i] == s[i] {
						sb.WriteRune(r)
					}
				}
			}
		}
		sl = append(sl, s)
	}
	fmt.Println(c2 * c3)
	fmt.Println(sb.String())
}
