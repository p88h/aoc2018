package main

import (
	"bufio"
	"fmt"
)

func react(line string, ignore byte) int {
	buf := make([]byte, 0, len(line))
	l := -1
	for i := 0; i < len(line); i++ {
		if line[i]&31 == ignore&31 {
			continue
		}
		if l >= 0 && buf[l]^line[i] == 32 {
			// reaction
			//fmt.Printf("Cut %c%c at %d\n", buf[l], line[i], i)
			buf = buf[:l]
			l--
		} else {
			// expansion
			buf = append(buf, line[i])
			l++
		}
	}
	return len(buf)
}

func Day5(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		return
	}
	min := react(scanner.Text(), 0)
	fmt.Println(min)
	for ignore := 'a'; ignore <= 'z'; ignore++ {
		tmp := react(scanner.Text(), byte(ignore))
		if tmp < min {
			min = tmp
		}
	}
	fmt.Println(min)
}
