package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parse4(strs []string) (ret []int64) {
	ret = make([]int64, len(strs))
	for i := range strs {
		ret[i], _ = strconv.ParseInt(strs[i], 10, 0)
	}
	return
}

func parse_insn(strs []string) Instruction {
	vec := parse4(strs)
	return Instruction{Opcode(vec[0]), vec[1], vec[2], vec[3]}
}

func parse_vmstate(strs []string) AocVM {
	return AocVM{parse4(strs), 0, -1}
}

func (x Aoc) Day16(scanner *bufio.Scanner) {
	r := regexp.MustCompile(`(Before|After):\s+\[(\d+),\s*(\d+),\s*(\d+),\s*(\d+)\]`)
	capture := false
	var before, after AocVM
	var insn Instruction
	program := make([]Instruction, 0, 1000)
	bad := make([][]bool, 16)
	for i := addr; i <= eqrr; i++ {
		bad[i] = make([]bool, 16)
	}
	tot1 := 0
	good := addr

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			g := r.FindStringSubmatch(scanner.Text())
			if capture {
				after = parse_vmstate(g[2:])
				tsame := 0
				opc := int64(insn.op)
				for i := addr; i <= eqrr; i++ {
					insn.op = i
					tmp := before.Copy()
					tmp.ExecOne(insn)
					if tmp.Equals(&after) {
						good = i
						tsame++
					} else {
						bad[opc][i] = true
					}
				}
				if tsame >= 3 {
					tot1++
				}
				capture = false
			} else {
				before = parse_vmstate(g[2:])
				capture = true
			}
		} else {
			nums := strings.Split(scanner.Text(), " ")
			if len(nums) == 4 {
				insn = parse_insn(nums)
				if !capture {
					program = append(program, insn)
				}
			}
		}
	}
	ooo := make(map[int64]Opcode, 16)
	for i := int64(0); i < 16; i++ {
		_, done := ooo[i]
		if done {
			continue
		}
		gcnt := 0
		for j := addr; j <= eqrr; j++ {
			if !bad[i][j] {
				gcnt++
				good = j
			}
		}
		if gcnt == 1 {
			ooo[i] = good
			for j := 0; j < 16; j++ {
				bad[j][good] = true
			}
			i = -1
		}
	}
	fmt.Printf("Mapped %d opcodes\n", len(ooo))
	vm := AocVM{make([]int64, 4), 0, -1}
	for _, insn := range program {
		insn.op = ooo[int64(insn.op)]
		// fmt.Printf("[%02d] %s %d %d %d | %v\n", op[0], descr[ooo[op[0]]], op[1], op[2], op[3], regs)
		vm.ExecOne(insn)
	}
	fmt.Printf("Final registers: %v\n", vm.regs)
	fmt.Printf("Day 16 Part1: %d\n", tot1)
	fmt.Printf("Day 16 Part2: %d\n", vm.regs[0])
}
