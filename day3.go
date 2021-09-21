package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile("\\#(\\d+) \\@ (\\d+),(\\d+): (\\d+)x(\\d+)")
	var m [1002004]int64
	var clean [4096]bool
	z := 0
	for scanner.Scan() {
		g := r.FindStringSubmatch(scanner.Text())
		id, _ := strconv.ParseInt(g[1], 10, 0)
		x, _ := strconv.ParseInt(g[2], 10, 0)
		y, _ := strconv.ParseInt(g[3], 10, 0)
		w, _ := strconv.ParseInt(g[4], 10, 0)
		h, _ := strconv.ParseInt(g[5], 10, 0)
		clean[id] = true
		for j := int64(0); j < h; j++ {
			o := (y+j)*1000 + x
			for i := int64(0); i < w; i++ {
				if m[o+i] == 0 {
					m[o+i] = id
				} else {
					clean[id] = false
					if m[o+i] > 0 {
						clean[m[o+i]] = false
						m[o+i] = -1
						z++
					}
				}
			}
		}
	}
	fmt.Println(z)
	for p, v := range clean {
		if v {
			fmt.Println(p)
			os.Exit(0)
		}
	}
}
