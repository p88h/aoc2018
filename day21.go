package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func (x Aoc) Day21(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`\#ip (\d+)`)
	program := make([]Instruction, 0, 1000)
	vm := AocVM{make([]int64, 6), 0, -1}
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			g := r.FindStringSubmatch(scanner.Text())
			r, _ := strconv.ParseInt(g[1], 10, 0)
			vm.rip = int(r)
		} else {
			insn := vm.Parse(scanner.Text())
			program = append(program, insn)
		}
	}
	minr := int64(0)
	maxr := int64(0)
	steps := 0
	stops := make(map[int64]bool)
	for steps < 3000000000 && vm.ip < len(program) {
		steps++
		if program[vm.ip].op == eqrr {
			r3 := vm.regs[3]
			if minr == 0 {
				//fmt.Printf("%d at %d steps\n", r3, steps)
				minr = r3
			} else if !stops[r3] {
				//fmt.Printf("New stop value: %d at step %d\n", r3, steps)
				maxr = r3
			}
			stops[r3] = true
		}
		vm.ExecOne(program[vm.ip])
		vm.ip++
	}
	fmt.Printf("Day 21 Part 1: %d\n", minr)
	fmt.Printf("Day 21 Part 2: %d\n", maxr)
}
