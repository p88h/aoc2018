package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	sl := make([]int64, 0, 1024)
	for scanner.Scan() {
		var x, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		sl = append(sl, x)
	}
	fs := make(map[int64]bool)
	var sum int64 = 0
	var iter = 0
	for {
		iter++
		for i := range sl {
			sum += sl[i]
			if fs[sum] {
				fmt.Println(sum)
				os.Exit(0)
			}
			fs[sum] = true
		}
		if iter == 1 {
			fmt.Println(sum)
		}
	}
}
