package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type light struct {
	x, y, dx, dy int64
}
type lpoint struct {
	x, y int64
}

func Day10(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
	a := make([]light, 0, 10000)
	for scanner.Scan() {
		g := r.FindStringSubmatch(scanner.Text())
		x, _ := strconv.ParseInt(g[1], 10, 0)
		y, _ := strconv.ParseInt(g[2], 10, 0)
		dx, _ := strconv.ParseInt(g[3], 10, 0)
		dy, _ := strconv.ParseInt(g[4], 10, 0)
		a = append(a, light{x, y, dx, dy})
	}
	gmin := int64(1000000)
	frame := 0
	for {
		miny := a[0].y
		maxy := a[0].y
		minx := a[0].x
		maxx := a[0].x
		for i := range a {
			a[i].x += a[i].dx
			a[i].y += a[i].dy
			if a[i].y < miny {
				miny = a[i].y
			}
			if a[i].y > maxy {
				maxy = a[i].y
			}
			if a[i].x < minx {
				minx = a[i].x
			}
			if a[i].x > maxx {
				maxx = a[i].x
			}
		}
		frame += 1
		fmt.Printf("After %d seconds, Y range: %d+%d X range %d+%d\n", frame, miny, maxy-miny, minx, maxx-minx)
		if maxy-miny < 10 {
			points := make(map[lpoint]int)
			for i := range a {
				points[lpoint{a[i].x, a[i].y}] = 1
			}
			for y := miny; y <= maxy; y++ {
				for x := minx; x <= maxx; x++ {
					if points[lpoint{x, y}] == 1 {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
		if maxy-miny < gmin {
			gmin = maxy - miny
		} else {
			return
		}
	}
}
