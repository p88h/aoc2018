package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type worker struct {
	eta  int
	prod byte
}

func process(deps map[byte][]byte, workers int) (string, int) {
	var sb strings.Builder
	reqs := make(map[byte]int)
	avail := 0
	for _, ds := range deps {
		for _, d := range ds {
			reqs[d]++
		}
	}
	// scan how many items are available
	for s := 'A'; s <= 'Z'; s++ {
		b := byte(s)
		if reqs[b] == 0 {
			avail++
		}
	}
	idle := workers
	work := make([]worker, workers)
	now := 0
	iter := 0
	for s := 'A'; s <= 'Z'; s++ {
		iter++
		// no items are ready to build or no workers are idling - advance time until we can do something
		for idle == 0 || avail == 0 {
			min := -1
			for p, w := range work {
				if w.prod > 0 && (min < 0 || w.eta < work[min].eta) {
					min = p
				}
			}
			now = work[min].eta
			sb.WriteByte(work[min].prod)
			fmt.Printf("[%d] Produced %c by %d NOW %d\n", iter, work[min].prod, min, now)
			// satisfy dependents
			for _, d := range deps[work[min].prod] {
				reqs[d]--
				if reqs[d] == 0 {
					avail++
				}
			}
			// free the worker
			work[min].prod = 0
			idle++
		}
		// add exactly one item, first alphabetically, to the work queue
		for c := 'A'; c <= 'Z'; c++ {
			b := byte(c)
			if reqs[b] == 0 {
				reqs[b] = -1
				for p, w := range work {
					if w.prod == 0 {
						work[p].prod = b
						work[p].eta = now + 60 + int(b-'A'+1)
						idle--
						avail--
						fmt.Printf("[%d] Insert %c at %d ETA %d\n", iter, c, p, work[p].eta)
						break
					}
				}
				break
			}
		}
	}
	for idle < workers {
		min := -1
		for p, w := range work {
			if w.prod > 0 && (min < 0 || w.eta < work[min].eta) {
				min = p
			}
		}
		now = work[min].eta
		sb.WriteByte(work[min].prod)
		fmt.Printf("[%d] Produced %c by %d NOW %d\n", iter, work[min].prod, min, now)
		// free the worker
		work[min].prod = 0
		idle++
	}
	return sb.String(), now
}

func (x Aoc) Day7(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
	deps := make(map[byte][]byte)
	for scanner.Scan() {
		g := r.FindStringSubmatch(scanner.Text())
		deps[g[1][0]] = append(deps[g[1][0]], g[2][0])
	}
	out1, cnt1 := process(deps, 1)
	fmt.Printf("%s %d\n", out1, cnt1)
	out5, cnt5 := process(deps, 5)
	fmt.Printf("%s %d\n", out5, cnt5)
}
