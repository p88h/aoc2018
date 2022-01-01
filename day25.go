package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type point4d struct {
	x, y, z, v, d int64
	parent        *point4d
	size          int64
}

func (p *point4d) distance(other *point4d) int64 {
	return abs64(p.x-other.x) + abs64(p.y-other.y) + abs64(p.z-other.z) + abs64(p.v-other.v)
}

func (p *point4d) find() *point4d {
	if p.parent != p {
		p.parent = p.parent.find()
	}
	return p.parent
}

func (p *point4d) union(other *point4d) *point4d {
	pp := p.find()
	op := other.find()
	if pp != op {
		if pp.size < op.size {
			pp.parent = op
			op.size += pp.size
		} else {
			op.parent = pp
			pp.size += op.size
			op = pp
		}
	}
	return op
}

func (x Aoc) Day25(scanner *bufio.Scanner) {
	points := make([]*point4d, 0, 10000)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		p := &point4d{}
		p.x, _ = strconv.ParseInt(arr[0], 10, 0)
		p.y, _ = strconv.ParseInt(arr[1], 10, 0)
		p.z, _ = strconv.ParseInt(arr[2], 10, 0)
		p.v, _ = strconv.ParseInt(arr[3], 10, 0)
		p.d = abs64(p.x) + abs64(p.y) + abs64(p.z) + abs64(p.v)
		p.parent = p
		p.size = 1
		points = append(points, p)
	}
	sort.Slice(points, func(i, j int) bool { return points[i].d < points[j].d })
	limit := int64(4)
	for i, p := range points {
		for j := i + 1; j < len(points) && points[j].d-p.d < limit; j++ {
			if p.distance(points[j]) <= 3 {
				// fmt.Printf("merge %d %d\n", i, j)
				p.union(points[j])
			}
		}
	}
	tc := 0
	for _, p := range points {
		if p.find() == p {
			tc += 1
		}
	}
	fmt.Println(tc)
}
