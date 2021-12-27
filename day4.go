package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
)

func lastdayofmonth(month, day int) bool {
	switch month {
	case 2:
		// 1518 is not a leap year. Also, February doesn't actually seem to occur in the input
		return day == 28
	case 4, 6, 9, 11:
		return day == 30
	case 1, 3, 5, 7, 8, 10, 12:
		return day == 31
	}
	return false
}

// return day index & minute of day
func timeparse(s string) (d, i int) {
	var mm, dd, h, m int
	fmt.Sscanf(s, "%d-%d %d:%d", &mm, &dd, &h, &m)
	// clip up
	if h == 23 {
		if lastdayofmonth(mm, dd) {
			mm++
			dd = 1
		} else {
			dd++
		}
		h = 0
		m = 0
	}
	d = mm*100 + dd
	i = h*60 + m
	return
}

func (x Aoc) Day4(scanner *bufio.Scanner) {
	r1 := regexp.MustCompile(`\[1518-(\d+-\d+ \d+:\d+)\] Guard #(\d+) begins shift`)
	r2 := regexp.MustCompile(`\[1518-(\d+-\d+ \d+:\d+)\] (wakes up)?(falls asleep)?`)
	// map day -> guard
	d2g := make(map[int]int)
	// map day -> event list
	d2i := make(map[int][]int)
	for scanner.Scan() {
		if r1.MatchString(scanner.Text()) {
			s := r1.FindStringSubmatch(scanner.Text())
			d, _ := timeparse(s[1])
			var g int
			fmt.Sscan(s[2], &g)
			d2g[d] = g
		} else if r2.MatchString(scanner.Text()) {
			s := r2.FindStringSubmatch(scanner.Text())
			d, m := timeparse(s[1])
			d2i[d] = append(d2i[d], m)
		}
	}
	// guard -> sleep totals
	g2s := make(map[int]int)
	g2t := make(map[int][]int)
	for d, g := range d2g {
		sort.Ints(d2i[d])
		// fmt.Printf("Day %d Guard %d ", d, g)
		// fmt.Println(d2i[d])
		var s int = 0
		var is = d2i[d]
		if _, ok := g2t[g]; !ok {
			g2t[g] = make([]int, 60)
		}
		for i := 0; i < len(is)/2; i++ {
			for j := is[i*2]; j < is[i*2+1]; j++ {
				g2t[g][j]++
				s++
			}
		}
		g2s[g] += s
	}
	var smax, gmax, retg, rets int = 0, 0, 0, 0
	for g, s := range g2s {
		var max, idx int = 0, 0
		for m, v := range g2t[g] {
			if v > max {
				idx = m
				max = v
			}
			if v > gmax {
				retg = g * idx
				gmax = v
			}
		}
		if s > smax {
			smax = s
			rets = g * idx
		}
		fmt.Printf("Guard %d sleep %d max %d idx %d\n", g, s, max, idx)
		// fmt.Println(g2t[g])
	}
	fmt.Println(rets)
	fmt.Println(retg)
}
