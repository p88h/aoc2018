package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type nanobot struct {
	x, y, z, r int64
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func (nb nanobot) distance(other nanobot) int64 {
	return abs64(nb.x-other.x) + abs64(nb.y-other.y) + abs64(nb.z-other.z)
}

func (nb nanobot) inrange(other nanobot) bool {
	return nb.distance(other) <= nb.r
}

func (x Aoc) Day23(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`pos=<(\-?\d+),(\-?\d+),(\-?\d+)>, r=(\d+)`)
	bots := make([]nanobot, 0, 10000)
	maxi := 0
	ax, ay, az := int64(0), int64(0), int64(0)
	points := make([]int64, 0, 2000)
	for scanner.Scan() {
		g := r.FindStringSubmatch(scanner.Text())
		x, _ := strconv.ParseInt(g[1], 10, 0)
		y, _ := strconv.ParseInt(g[2], 10, 0)
		z, _ := strconv.ParseInt(g[3], 10, 0)
		r, _ := strconv.ParseInt(g[4], 10, 0)
		d := abs64(x) + abs64(y) + abs64(z)
		points = append(points, (d-r)*2)
		points = append(points, (d+r)*2+1)
		nb := nanobot{x, y, z, r}
		bots = append(bots, nb)
		if nb.r > bots[maxi].r {
			maxi = len(bots) - 1
		}
	}
	tot := 0
	tot2 := 0
	cand := nanobot{ax, ay, az, 0}
	for _, nnb := range bots {
		if bots[maxi].inrange(nnb) {
			tot += 1
		}
		if nnb.inrange(cand) {
			tot2 += 1
		}
	}
	fmt.Printf("Day 23 Part 1: %d\n", tot)
	sort.Slice(points, func(i, j int) bool { return points[i] < points[j] })
	cnt := int64(0)
	mc := int64(0)
	md := int64(0)
	for _, v := range points {
		d := v / 2
		if v%2 == 0 {
			cnt += 1
		} else {
			cnt -= 1
		}
		if cnt > mc {
			mc = cnt
			md = d
		}
	}
	fmt.Printf("Max count %d at %d\n", mc, md)
}
