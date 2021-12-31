package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func fast_prog(r0 int64) int64 {
	r1 := int64(2)
	r1 = r1 * r1
	r1 *= 19
	r1 *= 11
	r4 := int64(7)
	r4 *= 22
	r4 += 20
	r1 += r4
	if r0 == 1 {
		r4 = 27
		r4 *= 28
		r4 += 29
		r4 *= 30
		r4 *= 14
		r4 *= 32
		r1 += r4
	}
	fmt.Printf("r1 is %d\n", r1)

	r0 = int64(0)
	for r3 := int64(1); r3 <= r1; r3++ {
		if r1%r3 == 0 {
			r0 += r3
		}
	}
	return r0
}

func (x Aoc) Day19(scanner *bufio.Scanner) {
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
	for vm.ip < len(program) {
		//fmt.Println(vm.Describe(program[vm.ip]))
		vm.ExecOne(program[vm.ip])
		vm.ip++
	}
	fmt.Printf("Final registers: %v\n", vm.regs)
	fmt.Printf("Fast prog 1: %d\n", fast_prog(0))
	fmt.Printf("Fast prog 2: %d\n", fast_prog(1))
}
